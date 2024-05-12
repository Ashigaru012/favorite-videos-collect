// TODO:判定系やテストしてみても良いかも
// TODO:VideoPages をパッケージ化


package fetcher

import (
	"fmt"
	"time"
    "github.com/go-rod/rod"
	"strconv"
	"net/url"
	"log"
	"sync"
	"my-app-go/scrapers"
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

func IsNumberInRange(getNum int,min int,max int) bool {
	if getNum < min && getNum > max {
		fmt.Println("Please enter 1 or more and 20 or less")
		return false
	} else {
		return true
	}
}

func FetchTargetPageVideos(searchQuery string,getNum int,vp scrapers.VideoPage,ch chan []Video,wg *sync.WaitGroup) {
	var videoList []Video
	second := 1
	baseURL, err := url.Parse(fmt.Sprintf(vp.TargetURL,searchQuery))
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
		item := page.MustElement(fmt.Sprintf(vp.Item, numStr))

		// このあたりVideo構造体に値をセットするメソッド作ってもよいかも
		name := item.MustElement(vp.Name).MustText()
		u,err := url.Parse(string(*(item.MustElement(vp.Url).MustAttribute("href"))))
		if err != nil {
			log.Fatal(err)
		}
		image := *(item.MustElement(vp.Image).MustAttribute("src"))

		videoList = append(videoList,Video{name,baseURL.ResolveReference(u).String(),image})
		time.Sleep(time.Duration(second) * time.Second)
	}

	ch <- videoList
	wg.Done()
}





