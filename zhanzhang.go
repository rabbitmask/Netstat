package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// 基于站长之家api的物理地址及运营商查询
func GetAdr_ZZ(ipadrs map[string]string,ip string)string {
	if _, ok :=ipadrs[ip];ok{
		return ipadrs[ip]
	}else {
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://ip.tool.chinaz.com", strings.NewReader("ip="+ip))
		checkErr(err)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:67.0) Gecko/20100101 Firefox/67.0")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//req.Header.Set("Cookie", "name=anny")
		req.AddCookie(&http.Cookie{Name: "BAIDUID",Value: "00A1B1EC9FF50D09E8740C2BB49A2120"})
		resp, err := client.Do(req)
		checkErr(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)
		adr := regexp.MustCompile(`<span class="Whwtdhalf w50-0">(.*?)</span>`).FindAllStringSubmatch(string(body), -1)
		//fmt.Println(ip)
		//fmt.Println(adr)
		ipadrs[ip]=adr[1][1]
		time.Sleep(1)
		return adr[1][1]
	}
}