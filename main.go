package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"magireco/htmltemplate"
	"magireco/model"
	"net/http"
	"net/url"
	"strings"
)

var (
	useProxy   = false
	SocksProxy = "socks5://127.0.0.1:10808"
)

func main() {
	//获取公告json串
	domain := "https://android.magi-reco.com/magica/json/announcements/announcements.json"
	resp, body := httpGet(domain)
	infos := make([]model.InfoStruct, 0)
	err := json.Unmarshal(body, &infos)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	//取最新时间
	timeArr := make([]string, 0)
	for _, info := range infos {
		timeArr = append(timeArr, info.StartAt)
	}
	//获取并格式化最新公告
	htmlArr := make([]model.HtmlStruct, 0)
	for _, info := range infos {
		timeStr := timeArr[len(timeArr)-1]
		if info.StartAt == timeStr {
			temp := model.HtmlStruct{}
			info.Text = strings.ReplaceAll(info.Text, "data-src='",
				"data-src='https://android.magi-reco.com/magica/resource/download/asset/master/")
			temp.SubText = info.Subject
			temp.Text = info.Text
			htmlArr = append(htmlArr, temp)
		}
	}
	//生成html
	htmlStr := ""
	htmlStr += htmltemplate.Header
	for _, html := range htmlArr {
		htmlStr += fmt.Sprintf(htmltemplate.ContentFormat, html.SubText, html.Text)
	}
	htmlStr += htmltemplate.Footer
	//写入文件
	err = ioutil.WriteFile("index.html", []byte(htmlStr), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished!Please open index.html")
}

func httpGet(domain string) (*http.Response, []byte) {
	httpclient := generateHttpClent(useProxy)

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		panic(err)
	}
	resp, err := httpclient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	return resp, body
}

func generateHttpClent(useProxy bool) *http.Client {
	if useProxy {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(SocksProxy)
		}
		httpTransport := &http.Transport{
			Proxy: proxy,
		}
		httpclient := &http.Client{
			Transport: httpTransport,
		}
		return httpclient
	} else {
		client := &http.Client{}
		return client
	}
}
