package scrapers


type VideoPage struct {
	TargetURL string
    Item      string
    Name      string
    Url       string
    Image     string
}

func (vp *VideoPage) SetValue(Item string,Name string,Url string,Image string) {
	vp.Item = Item
	vp.Name = Name
	vp.Url = Url
	vp.Image = Image
}

func GetVideoPages() map[string]VideoPage {
    VideoPages := make(map[string]VideoPage)

    VideoPages["tokyomotion"] = VideoPage{
        "https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",
        "#wrapper > div.container > div.row > div > div.row > div:nth-child(%s) > div",
        "a > span",
        "a",
        "a > div > img",
    }

    VideoPages["tktube"] = VideoPage{
        "https://tktube.com/ja/search/%s/",
        "body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
        "a > strong",
        "a",
        "a > div.img > img",
    }

	// VideoPages["javmix.tv"] = VideoPage{
	// 	"https://javmix.tv/%s/",
	// 	"body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
	// 	"a > strong",
	// 	"a",
	// 	"a > div.img > img",
	// }

    return VideoPages
}