// TODO:videoList配列にVideo型データをセットする関数を作成する
// package site
package main

import (
	"fmt"
	"time"
    "github.com/go-rod/rod"
	"strconv"
)

type Video struct {
	Name string
	Url string
	Image string
}

func IsNumberInRange(getNum int,min int,max int) (b bool){
	if getNum < min && getNum > max {
		fmt.Println("Please enter 1 or more and 20 or less")
		return false
	} else {
		return true
	}
}

func FetchTokyomotionVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://www.tokyomotion.net"
    Page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/search?search_query=" + searchQuery + "&search_type=videos")
	defer Page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := Page.MustElement("#wrapper > div.container > div.row > div > div.row > div:nth-child(" + numStr + ") > div")
		name := item.MustElement("a > span").MustText()
		path := item.MustElement("a").MustAttribute("href")
		url := target + *path
		image := item.MustElement("a > div > img").MustAttribute("src")

		videoList = append(videoList,Video{name,url,*image})
		time.Sleep(1 * time.Second)
	}

	return videoList
}

func FetchtktubeVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://tktube.com/ja"
    Page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/search/" + searchQuery)
	defer Page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := Page.MustElement("body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(" + numStr + ")")
		name := item.MustElement("a > strong").MustText()
		url := item.MustElement("a").MustAttribute("href")
		image := item.MustElement("a > div.img > img").MustAttribute("src")

		videoList = append(videoList,Video{name,*url,*image})
		time.Sleep(1 * time.Second)
	}

	return videoList
}

func FetchsupjavVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://supjav.com/ja"
    Page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/?s=" + searchQuery)
	fmt.Println(Page.MustElement("body > div > div > div > a").MustText())
	defer Page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := Page.MustElement("body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(" + numStr + ")")
		name := item.MustElement("a > strong").MustText()
		url := item.MustElement("a").MustAttribute("href")
		image := item.MustElement("a > div.img > img").MustAttribute("src")

		videoList = append(videoList,Video{name,*url,*image})
		time.Sleep(1 * time.Second)
	}

	return videoList
}

func FetchjavbangersVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://www.javbangers.com"
    Page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/search/" + searchQuery)
	fmt.Println(Page.MustElement("body > div.container > div.content > div > div.main-container > div > div.porntrex-box > div > div:nth-child(1)"))
	defer Page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := Page.MustElement("body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(" + numStr + ")")
		name := item.MustElement("a > strong").MustText()
		url := item.MustElement("a").MustAttribute("href")
		image := item.MustElement("a > div.img > img").MustAttribute("src")

		videoList = append(videoList,Video{name,*url,*image})
		time.Sleep(1 * time.Second)
	}

	return videoList
}

func main() {
	items := FetchjavbangersVideos("fc2",5)
	fmt.Println(items[0].Image)
}
