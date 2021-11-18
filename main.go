package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	SocksProxy = "socks5://127.0.0.1:10808"
)

func main() {
	domain := "https://android.magi-reco.com/magica/json/announcements/announcements.json"
	resp, body := httpGet(domain)
	infos := make([]InfoStruct, 0)
	json.Unmarshal(body, &infos)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	fmt.Println(infos[len(infos)-1])
	timeArr := make([]string, 0)
	for _, info := range infos {
		timeArr = append(timeArr, info.StartAt)
	}
	htmlArr := make([]htmlStruct, 0)
	for _, info := range infos {
		timeStr := timeArr[len(timeArr)-1]
		if info.StartAt == timeStr {
			temp := htmlStruct{}
			temp.subText = info.Subject
			temp.text = info.Text
			htmlArr = append(htmlArr, temp)
		}
	}
	for _, html := range htmlArr {
		str := fmt.Sprintf("<p>%v</p>", html.text)
		content := fmt.Sprintf("<hr><br/><div>%v</div>", html.subText)
		fmt.Println(str + content)
	}
}

func httpGet(domain string) (*http.Response, []byte) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(SocksProxy)
	}
	httpTransport := &http.Transport{
		Proxy: proxy,
	}
	httpclient := &http.Client{
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		panic(err)
	}
	resp, err := httpclient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return resp, body
}
