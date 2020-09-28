package main

import (
	"strconv"
	"syscall"
	"unsafe"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}



func taskdic() map[string]string{
	tasks:=make(map[string]string)
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
	}
	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			tasks[strconv.Itoa(int(proc.th32ProcessID))]=string(proc.szExeFile[0:])
		} else {
			break
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)
	return tasks
}