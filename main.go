package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Details struct {
	Title   string
	Author  string
	Time    string
	Content string
}

// 首页标题正则
const titleRegexp = `"pure_title": "([^"]+)"`

func main() {
	// 获取新闻内页链接
	urlList := topHot()
	details := make([]Details, 30)

	// 获取每一则新闻的详细信息
	for _, url := range urlList {
		if url == "" {
			continue
		}
		byte, err := worker(url)
		if err != nil {
			log.Printf("Get detail error: %v\n", err)
			continue
		}
		reader := bytes.NewReader(byte)
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Printf("Get doc error: %v\n", err)
			continue
		}
		title 		:= doc.Find(".article-title h2").Text()
		author 		:= doc.Find(".author-name a").Text()
		time 		:= doc.Find(".article-source").Text()
		content, _ 	:= doc.Find(".article-content").Html()
		details = append(details, Details{
			Title: title,
			Author: author,
			Time: time,
			Content: content,
		})
		fmt.Printf("result: %v\n\n", details)
	}
}

func worker(url string) ([]byte, error) {
	// 设置客户端: 请求头等
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	// 解析内容
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("fetch error")
	}
	return byte, nil
}

/**
 * 百度热点列表
 */
func topHot() []string {
	// 解析内容
	url := "https://www.baidu.com"
	byte, err := worker(url)
	if err != nil {
		panic(err)
	}

	// 提取标题
	urlList := make([]string, 30)
	reGe := regexp.MustCompile(titleRegexp)
	res := reGe.FindAllSubmatch(byte, -1)
	for _, temp := range res {
		urlList = append(urlList, hotList(string(temp[1])))
	}
	return urlList
}

/**
 * 搜索热点详情
 */
func hotList(title string) string {
	url := "https://www.baidu.com/s?wd=" + title
	byte, err := worker(url)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(byte)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
	// 提取搜索页第一个新闻提取url
	uri, exist := doc.Find(".op-timeliness-new-union .op-timeliness-new-title a").Attr("href")
	if !exist {
		uri = ""
		log.Printf("Not match url:; %s\n", url)
	}
	return uri
}