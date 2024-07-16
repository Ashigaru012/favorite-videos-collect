package scrapers

// ページ情報を格納する構造体
type VideoPage struct {
	TargetURL string
    Item      string
    Name      string
    Url       string
    Image     string
    Logo      string
}

// VideoPage 構造体に値をセットするメソッド
func (vp *VideoPage) SetValue(Item string,Name string,Url string,Image string) {
	vp.Item = Item
	vp.Name = Name
	vp.Url = Url
	vp.Image = Image
}

// 動画ページの情報を取得する関数
func GetVideoPages() map[string]VideoPage {
    VideoPages := make(map[string]VideoPage)

    VideoPages["tokyomotion"] = VideoPage{
        "https://www.tokyomotion.net/search?search_query=%s&search_type=videos&type=public",
        "#wrapper > div.container > div.row > div > div.row > div:nth-child(%s) > div",
        "a > span",
        "a",
        "a > div > img",
        "https://cdn.tokyo-motion.net/img/logo.gif",
    }

    VideoPages["tktube"] = VideoPage{
        "https://tktube.com/ja/search/%s/",
        "body > div.container > div.content > div > div.main-container > div > div > div > div > div:nth-child(%s)",
        "a > strong",
        "a",
        "a > div.img > img",
        "https://tktube.com/static/images/logo.png",
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