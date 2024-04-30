package main

import (
    "fmt"
    "net/url"
)

func main() {
    // a := []string{
    //     "https://www.tokyomotion.net/001/002.html", // ベースとなるWebページのURL
    //     "/foo/bar/",                        // base タグの href に指定してあるURL
    //     "baz/003.html",                     // アンカータグの href に指定されている相対URL
    // }
	searchQuery := "fc2"
	baseURL, _ := url.Parse(fmt.Sprintf("https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",searchQuery))
	url,_ := url.Parse("/foo/bar/")
	urlQuery := baseURL.RequestURI()
	fmt.Println(baseURL.ResolveReference(url),urlQuery)
}