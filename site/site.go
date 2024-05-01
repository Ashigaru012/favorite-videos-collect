// TODO:javbangersVideosがエラーになってる原因を調べる。Mustじゃなくてerr返ってくる方使ってみるとよいかも(element関数など)

// package site
package main

import (
	"fmt"
	"time"
    "github.com/go-rod/rod"
	"strconv"
	"net/url"
	"log"
)

type Video struct {
	Name string
	Url string
	Image string
}

func (v *Video) SetValue(Name string,Url string,Image string,Target string) {
	v.Name = Name
	v.Url = Url
	v.Image = Image
}

// Baseを親として別で構造体用意してもよいかも
type VideoPageScraper struct {
	TargetURL string
    Item      string
    Name      string
    Url       string
    Image     string
}

func (vs *VideoPageScraper) SetValue(Item string,Name string,Url string,Image string) {
	vs.Item = Item
	vs.Name = Name
	vs.Url = Url
	vs.Image = Image
}

func IsNumberInRange(getNum int,min int,max int) bool {
	if getNum < min && getNum > max {
		fmt.Println("Please enter 1 or more and 20 or less")
		return false
	} else {
		return true
	}
}


// ターゲットページを引数に入れて、動画をサイトから取ってくるようにする
// name,url,imageを事前にセットする関数を作る。baseセレクタは一旦ハードコードで置いておく
// TODO:urlQuery と itemElement はサイトによって変わるので、引数に入れておくようにする

func FetchTargetPageVideos(searchQuery string,getNum int,vs VideoPageScraper,ch chan []Video) {
	var videoList []Video
	second := 1
	baseURL, err := url.Parse(fmt.Sprintf(vs.TargetURL,searchQuery))
	if err != nil {
		log.Fatal(err)
	}
	
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(baseURL.String())

	defer page.MustClose()

	// if !IsNumberInRange(getNum,min,max) {
	// 	return
	// }

	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := page.MustElement(fmt.Sprintf(vs.Item, numStr))

		// このあたりVideo構造体に値をセットするメソッド作ってもよいかも
		name := item.MustElement(vs.Name).MustText()
		u,err := url.Parse(string(*(item.MustElement(vs.Url).MustAttribute("href"))))
		if err != nil {
			log.Fatal(err)
		}
		image := *(item.MustElement(vs.Image).MustAttribute("src"))

		videoList = append(videoList,Video{name,baseURL.ResolveReference(u).String(),image})
		time.Sleep(time.Duration(second) * time.Second)
	}

	ch <- videoList
}

func main() {
	tokyomotionPage := VideoPageScraper{
		"https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",
		"#wrapper > div.container > div.row > div > div.row > div:nth-child(%s) > div",
		"a > span",
		"a",
		"a > div > img",
	}

	// tktubePage := VideoPageScraper{
	// 	"https://tktube.com/ja/search/%s/",
	// 	"body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
	// 	"a > strong",
	// 	"a",
	// 	"a > div.img > img",
	// }

	// このあたり並列処理にしてもよいかも
	// tokyomotionVideos := FetchTargetPageVideos("fc2",5,tokyomotionPage)
	// tktubeVideos := FetchTargetPageVideos("fc2",5,tktubePage)

	ch := make(chan []Video, 2)
	go FetchTargetPageVideos("fc2",5,tokyomotionPage,ch)
	tokyomotionVideos := <- ch

	fmt.Println(tokyomotionVideos[0])
}



