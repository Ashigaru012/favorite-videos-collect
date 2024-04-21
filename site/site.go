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

// Baseを親として別で構造体用意してもよいかも
type VideoSelector struct {
    Base   string
    Name   string
    URL    string
    Image  string
}

func IsNumberInRange(getNum int,min int,max int) (b bool){
	if getNum < min && getNum > max {
		fmt.Println("Please enter 1 or more and 20 or less")
		return false
	} else {
		return true
	}
}

// ターゲットページを引数に入れて、動画をサイトから取ってくるようにする
// name,url,imageを事前にセットする関数を作る。baseセレクタは一旦ハードコードで置いておく
func FetchTargetPageVideos() {

}




func FetchTokyomotionVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://www.tokyomotion.net"
	urlQuery := target + "/search?search_query=" + searchQuery + "&search_type=videos"
    page := rod.New().NoDefaultDevice().MustConnect().MustPage(urlQuery)
	defer page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		itemElement := "#wrapper > div.container > div.row > div > div.row > div:nth-child(" + numStr + ") > div"
		nameElement := "a > span"
		urlElement := "a"
		imageElement := "a > div > img"
		second := 1
		item := page.MustElement(itemElement)
		name := item.MustElement(nameElement).MustText()
		url := string(*(item.MustElement(urlElement).MustAttribute("href")))
		url = target + url
		image := item.MustElement(imageElement).MustAttribute("src")

		videoList = append(videoList,Video{name,url,*image})
		time.Sleep(time.Duration(second) * time.Second)
	}

	return videoList
}

func FetchtktubeVideos(searchQuery string,getNum int) (v []Video) {
	var videoList []Video
	target := "https://tktube.com/ja"
    page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/search/" + searchQuery)
	defer page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := page.MustElement("body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(" + numStr + ")")
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
    page := rod.New().NoDefaultDevice().MustConnect().MustPage( target + "/search/" + searchQuery)
	defer page.MustClose()

	if getNum < 1 && getNum > 20 {
		fmt.Println("Please enter 1 or more and 20 or less")
		return
	}

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := page.MustElement("body > div.container > div.content > div > div.main-container > div > div.porntrex-box > div  > div:nth-child(" + numStr + ")")
		name := item.MustElement("a > strong").MustText()
		url := item.MustElement("a").MustAttribute("href")
		image := item.MustElement("a > div.img > img").MustAttribute("src")

		videoList = append(videoList,Video{name,*url,*image})
		time.Sleep(1 * time.Second)
	}

	return videoList
}

func main() {
	items := FetchTokyomotionVideos("fc2",5)
	fmt.Println(items[0])
}
