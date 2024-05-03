// TODO:判定系やテストしてみても良いかも

// package site
package main

import (
	"fmt"
	"time"
    "github.com/go-rod/rod"
	"strconv"
	"net/url"
	"log"
	"sync"
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

func FetchTargetPageVideos(searchQuery string,getNum int,vs VideoPageScraper,ch chan []Video,wg *sync.WaitGroup) {
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
	wg.Done()
}

func main() {
	// start := time.Now()
	searchWord := "fc2"
	getNum := 5
	videoPageScrapers := make(map[string]VideoPageScraper)
	videos := make(map[string][]Video)

	videoPageScrapers["tokyomotion"] = VideoPageScraper{
		"https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",
		"#wrapper > div.container > div.row > div > div.row > div:nth-child(%s) > div",
		"a > span",
		"a",
		"a > div > img",
	}

	videoPageScrapers["tktube"] = VideoPageScraper{
		"https://tktube.com/ja/search/%s/",
		"body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
		"a > strong",
		"a",
		"a > div.img > img",
	}

	ch := make(chan []Video, 2)
	var wg sync.WaitGroup
	wg.Add(len(videoPageScrapers))
	for pageName, _:= range videoPageScrapers {
		go FetchTargetPageVideos(searchWord,getNum,videoPageScrapers[pageName],ch,&wg)
		videos[pageName] = <- ch
	}
	close(ch)

	fmt.Println(videos["tktube"])
	// fmt.Println("took: ", time.Since(start))
}



