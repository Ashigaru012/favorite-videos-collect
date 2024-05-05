// scrapers.go
package scrapers

import "my-app-go/site"

func GetVideoPageScrapers() map[string]site.VideoPageScraper {
    videoPageScrapers := make(map[string]site.VideoPageScraper)

    videoPageScrapers["tokyomotion"] = site.VideoPageScraper{
        "https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",
        "#wrapper > div.container > div.row > div > div.row > div:nth-child(%s) > div",
        "a > span",
        "a",
        "a > div > img",
    }

    videoPageScrapers["tktube"] = site.VideoPageScraper{
        "https://tktube.com/ja/search/%s/",
        "body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
        "a > strong",
        "a",
        "a > div.img > img",
    }

	videoPageScrapers["javmix.tv"] = VideoPageScraper{
		"https://javmix.tv/%s/",
		"body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
		"a > strong",
		"a",
		"a > div.img > img",
	}

    return videoPageScrapers
}