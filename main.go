package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// PrintfURLAndPlace ...
func PrintfURLAndPlace(all []byte) {
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/kunming" data-v-5e878427>昆明</a>
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-648a6cbc>北京</a>
	reg, err := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)
	if err != nil {
		fmt.Printf("regexp error: %v", err)
	}
	urlCity := reg.FindAllSubmatch(all, -1)
	for _, v := range urlCity {
		fmt.Printf("url= %s and city= %s\n", v[2], v[1])
		fmt.Println()
	}
	fmt.Println(len(urlCity))
}
func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status error", resp.StatusCode)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	PrintfURLAndPlace(all)
}

/*
1:通用爬虫
2:聚焦爬虫
市面上受欢迎的库
henrylee2cn/pholcus
gocrawl
colly
hu17889/go_spider
本项目使用ElasticSearch作为数据存储
使用Go语言标准模板库实现http数据展示部分

总体的算法

golang.org/x/test
golang.org/x/net/html

1：css选择器
$('#__nuxt')
$('#__nuxt>dd>a')

2：使用xpath
3：使用正则表达式
reg,err:=regexp.Compile("匹配的信息")
匹配的信息可以这样写
.+/.*
.代表任何一个字符
+一个或者多个
*0个或者多个








*/
