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

// 動画情報を格納する構造体
type Video struct {
	Name string
	Url string
	Image string
}

// Video 構造体に値をセットするメソッド
func (v *Video) SetValue(Name string,Url string,Image string,Target string) {
	v.Name = Name
	v.Url = Url
	v.Image = Image
}

// 入力を判定する関数
func IsNumberInRange(getNum int,min int,max int) bool {
	if getNum < min && getNum > max {
		fmt.Println("Please enter 1 or more and 20 or less")
		return false
	} else {
		return true
	}
}

// scrapers.VideoPageの情報をもとにサイトから動画をスクレイピングする関数
func FetchTargetPageVideos(searchQuery string,getNum int,vp scrapers.VideoPage,ch chan []Video,wg *sync.WaitGroup) {
	var videoList []Video
	second := 1

	// 検索ワードを含んだURLに変換
	baseURL, err := url.Parse(fmt.Sprintf(vp.TargetURL,searchQuery))
	if err != nil {
		log.Fatal(err)
	}
	
	// URL先のページを取得
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(baseURL.String())

	defer page.MustClose()

	// if !IsNumberInRange(getNum,min,max) {
	// 	return
	// }

	// 構造体 Video に値をセットする処理
	for i:=1; i<=getNum; i++ {
		numStr := strconv.Itoa(i)
		item := page.MustElement(fmt.Sprintf(vp.Item, numStr))
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
