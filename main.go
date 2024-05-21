package main

import (
        "log"
        "net/http"
        "text/template"
        "github.com/gorilla/sessions"
        "sync"
        "my-app-go/fetcher"
	"my-app-go/scrapers"
)

// session variable. (not used)
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-1234"))

// Template for no-template.
func notemp() *template.Template {
        src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
        tmp, _ := template.New("index").Parse(src)
        return tmp
}

// get target Temlate.
func page(fname string) *template.Template {
        tmps, _ := template.ParseFiles("templates/"+fname+".html",
                "templates/head.html", "templates/foot.html")
        return tmps
}

// index handler.
func index(w http.ResponseWriter, rq *http.Request,videos map[string][]fetcher.Video) {
        er := page("index").Execute(w, videos)
        if er != nil {
                log.Fatal(er)
        }
}

// hello handler.
func hello(w http.ResponseWriter, rq *http.Request) {
        data := []string{
                "One", "Two", "Three",
        }

        item := struct {
                Title string
                Data  []string
        }{
                Title: "Hello",
                Data:  data,
        }

        er := page("hello").Execute(w, item)
        if er != nil {
                log.Fatal(er)
        }
}

// main program.
func main() {

        searchWord := "fc2"
        getNum := 5
        VideoPages := scrapers.GetVideoPages()
    
        ch := make(chan []fetcher.Video)
        var wg sync.WaitGroup
        videos := make(map[string][]fetcher.Video)
    
        wg.Add(len(VideoPages))
        for pageName, vp := range VideoPages {
            go func(pageName string, vp scrapers.VideoPage) {
                defer wg.Done()
                fetcher.FetchTargetPageVideos(searchWord, getNum, vp, ch, &wg)
            }(pageName, vp)
        }
    
        go func() {
            wg.Wait()
            close(ch)
        }()
    
        for videoList := range ch {
            for pageName := range VideoPages {
                videos[pageName] = append(videos[pageName], videoList...)
            }
        }
        // index handling.
        http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
                index(w, rq,videos)
        })
        // hello handling
        http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
                hello(w, rq)
        })

        http.ListenAndServe(":8080", nil)
}
