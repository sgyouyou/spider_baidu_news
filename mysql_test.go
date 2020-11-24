package main

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkDb(t *testing.T) {
	//mysqlDriver.WorkDb()
	_, err := time.Parse("2006-01-02 15:04", "2020-11-24 15:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now().Unix())

}
