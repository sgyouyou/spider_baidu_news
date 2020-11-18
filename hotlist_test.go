package main

import (
	"fmt"
	"testing"
)

func TestHotList(t *testing.T) {
	url := hotList("北方暴雪南方多地气温超31度")
	//url := hotList("孙正义出售软银800亿美元资产")
	fmt.Println(url)
}
