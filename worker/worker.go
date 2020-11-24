package worker

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Worker struct {}

/**
 * 请求worker
 */
func (w Worker) Worker(uri string) ([]byte, error) {
	// 设置客户端: 请求头等
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
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
