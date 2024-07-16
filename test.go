// TODO:判定系やテストしてみても良いかも
// TODO:videoPageScrapers をパッケージ化

// package fetcher
package main

import (
	"fmt"
	"sync"
	"my-app-go/fetcher"
	"my-app-go/scrapers"
)

func main() {
	// start := time.Now()
	searchWord := "fc2"
	getNum := 5
	VideoPages := scrapers.GetVideoPages()

	// チャンネル作成
	ch := make(chan []fetcher.Video, 2)
	var wg sync.WaitGroup
	videos := make(map[string][]fetcher.Video)
	wg.Add(len(VideoPages))

	// 動画情報を videos 配列に格納する処理
	// goroutine 使って FetchTargetPageVideos 関数の処理を並行化
	for pageName, _:= range VideoPages {
		go fetcher.FetchTargetPageVideos(searchWord,getNum,VideoPages[pageName],ch,&wg)
		videos[pageName] = <- ch
	}
	close(ch)

	fmt.Println(videos)
	// fmt.Println("took: ", time.Since(start))
}



