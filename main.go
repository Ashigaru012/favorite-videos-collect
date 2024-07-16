package main

import (
        // "log"
        "net/http"
        // "text/template"
        // "github.com/gorilla/sessions"
        "strconv"
        "sync"
        "my-app-go/fetcher"
	"my-app-go/scrapers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
        // インスタンスを作成
        e := echo.New()
      
        // ミドルウェアを設定
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
      
        // ルートを設定
        e.GET("/", IndexPage)
        e.GET("/fetcher",GetVideos)
      
        // サーバーをポート番号1323で起動
        e.Logger.Fatal(e.Start(":1323"))
      }
      
      // ハンドラーを定義
      func IndexPage(c echo.Context) error {
        return c.File("public/index.html")
      }

      func GetVideos(c echo.Context) error {
        searchWord := c.QueryParam("searchWord")
	getNum,_ := strconv.Atoi(c.QueryParam("getNum"))
        pageName := c.QueryParam("pageName")
	VideoPages := scrapers.GetVideoPages()

	// チャンネル作成
	ch := make(chan []fetcher.Video, 2)
	var wg sync.WaitGroup
	videos := make(map[string][]fetcher.Video)
	

        if pageName != "ALL" {
                wg.Add(1)
                go fetcher.FetchTargetPageVideos(searchWord,getNum,VideoPages[pageName],ch,&wg)
                videos[pageName] = <- ch
                close(ch)
                return c.JSON(http.StatusOK, videos[pageName])
        } 

	// 動画情報を videos 配列に格納する処理
	// goroutine 使って FetchTargetPageVideos 関数の処理を並行化
        wg.Add(len(VideoPages))
	for p, _:= range VideoPages {
		go fetcher.FetchTargetPageVideos(searchWord,getNum,VideoPages[p],ch,&wg)
		videos[p] = <- ch
	}
	close(ch)
        return c.JSON(http.StatusOK, videos)
      }