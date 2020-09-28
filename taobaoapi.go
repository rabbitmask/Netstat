package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type IPInfo struct {
	Data IP  `json:"data"`
}

type IP struct {
	Country   string `json:"country"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Isp       string `json:"isp"`
}

func TabaoAPI(ip string) *IPInfo {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://ip.taobao.com/outGetIpInfo", strings.NewReader("ip="+ip+"&accessKey=alibaba-inc"))
	checkErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:67.0) Gecko/20100101 Firefox/67.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: "test",Value: "test"})
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()


	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return nil
	}

	return &result
}

// 基于淘宝api的物理地址及运营商查询
func GetAdr_TB(ipadrs map[string]string,ip string)string {
	if _, ok :=ipadrs[ip];ok{
		return ipadrs[ip]
	}else {
		result:=TabaoAPI(ip)
		adr:= result.Data.Country+" "+result.Data.Region+" "+result.Data.City+" "+result.Data.Isp
		ipadrs[ip]=adr
		return adr
	}
}
