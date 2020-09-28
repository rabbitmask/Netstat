package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)


func netstat()string{
	cmd := exec.Command("netstat","-anop","tcp")
	buf, _ := cmd.Output()
	return string(buf)
}


func main() {
	ipadrs:=make(map[string]string)
	tasks:=make(map[string]string)
	tasks=taskdic()
	res := strings.Split(strings.Replace(netstat(),"\r\n","",-1), "TCP")
	var nw [][] string
	var gw [][] string
	for _,tcp:=range res[1:]{
		if strings.Contains(tcp, "0.0.0.0") || strings.Contains(tcp, "127.0.0.1"){

		}else{
			detail:=strings.Split((delete_extra_space(tcp)), " ")
			re := regexp.MustCompile(`^(127\.0\.0\.1)|(localhost)|(10\.\d{1,3}\.\d{1,3}\.\d{1,3})|(172\.((1[6-9])|(2\d)|(3[01]))\.\d{1,3}\.\d{1,3})|(192\.168\.\d{1,3}\.\d{1,3})$`)
			if re.MatchString(detail[2]) {
				var s[]string
				s=append(s,detail[1],detail[2],detail[4],tasks[detail[4]],"本地局域网")
				nw= append(nw, s)
			}else{
				var s[]string
				var adr string
				if apiConfig=="tb"{
					adr=GetAdr_TB(ipadrs,strings.Split(detail[2],":")[0])
				}else if apiConfig=="zz"{
					adr=GetAdr_ZZ(ipadrs,strings.Split(detail[2],":")[0])
				}else {
					fmt.Println("接口配置失败")
					break
				}

				s=append(s,detail[1],detail[2],detail[4],tasks[detail[4]],adr)
				//fmt.Println(s)
				gw= append(gw, s)
			}
		}

	}

	fmt.Println("# 内网")
	for _,v:=range nw{
		fmt.Println(v[0]+"\t"+v[1]+"\t"+v[2]+"\t"+v[3]+"\t"+v[4])
	}
	fmt.Println("\n# 公网")
	for _,v:=range gw{
		//fmt.Println(v)
		fmt.Println(v[0]+"\t"+v[1]+"\t"+v[2]+"\t"+v[3]+"\t"+v[4])
	}
	fmt.Println("\n"+"w(ﾟДﾟ)w！！！有内鬼，终止交易！！！\n一个简单的netstat + tasklist + ipwhois 反入侵检测小工具.\n项目地址：https://github.com/rabbitmask/Netstat")
}

