package gag9

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	GAG9_URL   = "https://9gag.com"
	IMG_URL    = "https://img-9gag-fun.9cache.com/photo/%s_460s.jpg"
	ANCHOR_TAG = "a"
)

type Meme struct {
	Description string
	Image       string
}

type Gag9 struct {
	client *http.Client
}

func New() *Gag9 {
	gag9 := &Gag9{
		client: &http.Client{},
	}
	return gag9
}

func (g *Gag9) Find() []Meme {
	req, err := http.NewRequest("GET", GAG9_URL, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := g.client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return crawl(resp.Body)
}

func (g *Gag9) FindByTag(tag string) []Meme {
	req, err := http.NewRequest("GET", GAG9_URL+"/tag/"+tag, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := g.client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return crawl(resp.Body)
}

func crawl(body io.Reader) []Meme {
	memes := make([]Meme, 0)
	tokenizer := html.NewTokenizer(body)

	found := false
	content := ""
	meme := Meme{}
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			return memes
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == ANCHOR_TAG {
				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "badge-evt badge-track" {
						img := fmt.Sprintf(IMG_URL, t.Attr[3].Val)
						meme.Image = img
						found = true
					}
				}
			}
		case html.TextToken:
			if found {
				t := tokenizer.Token()
				content = content + t.String()
			}
		case html.EndTagToken:
			t := tokenizer.Token()
			if t.Data == ANCHOR_TAG && found {
				meme.Description = strings.TrimSpace(content)
				memes = append(memes, meme)
				content = ""
				found = false
			}
		}
	}
}
