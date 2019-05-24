package main

import (
	"testing"
	"ericivan/crawler/fetcher"
	"fmt"
	"ericivan/crawler/zhenai/parser"
	"time"
	"github.com/chromedp/chromedp"
	"log"
	"context"
)

func TestProfile(t *testing.T) {

	//content, err := fetcher.Fetch("http://album.zhenai.com/u/1275316631")
	content, err := fetcher.Fetch("http://m.zhenai.com/u/1172365166")

	//content, _ := fetcher.Fetch("http://zhenai.com/zhenghun")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(content))
	//parser.ParseProfile(content)
}

func TestParseCity(t *testing.T) {

	content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/anyang")

	if err != nil {

		fmt.Errorf(err.Error())
	}

	parser.ParseCity(content)
}

func text(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		// 访问页面
		chromedp.Navigate(`https://movie.douban.com/tag/`),
		// 等待列表渲染
		chromedp.Sleep(5 * time.Second),
		// 获取获取服务列表HTML
		chromedp.OuterHTML("#content", res, chromedp.ByID),
	}
}

func TestChrome(t *testing.T) {
	var res string

	text(&res)

	fmt.Printf("result is %s", res)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建 chrome 实例
	cdp, err := chromedp.NewContext(ctx, func(ctx *chromedp.Context) {

	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cdp)
}

func TestParseMoveCategory(t *testing.T) {


	url := "https://www.xl720.com/category/dongzuopian"

	content, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	parser.ParseMoveCategory(content)
}
