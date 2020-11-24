package main

import (
	"baidu_news/model"
	"baidu_news/worker"
	_ "baidu_news/worker"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	url2 "net/url"
	"regexp"
	"time"
)

// 首页标题正则
const titleRegexp = `"pure_title": "([^"]+)"`

func main() {
	// 获取新闻内页链接
	urlList := topHot()

	// 获取每一则新闻的详细信息
	for _, url := range urlList {
		if url == "" {
			continue
		}
		byte, err := worker.Worker{}.Worker(url)
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
		//date 		:= doc.Find(".article-source .date").Text() + " " + doc.Find(".article-source .time").Text()
		content, _ 	:= doc.Find(".article-content").Html()

		//timeRune := []rune(date)
		//date = string(timeRune[5:])
		model.Post{}.InsertNew(model.News{
			Title: 	 	title,
			Author:  	author,
			Time: 	 	time.Now().Unix(),
			Content: 	content,
		})
	}
}


/**
 * 百度热点列表
 */
func topHot() []string {
	// 解析内容
	url := "https://www.baidu.com"
	byte, err := worker.Worker{}.Worker(url)
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
	url := "https://www.baidu.com/s?wd=" + url2.QueryEscape(title)
	byte, err := worker.Worker{}.Worker(url)
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