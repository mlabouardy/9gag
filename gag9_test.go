package gag9

import (
	"bytes"
	"reflect"
	"testing"
)

const RAW_HTML = `
<a class="badge-evt badge-track"
   data-evt="PostList,TapPost,Tag,,PostTitle"
   data-track="post,v,,,d,aKDjw1N,l"
   data-entry-id="aKDjw1N"
   data-position="1"
   href="/gag/aKDjw1N"
   target="_blank">
  When you mix MJ with Got
</a>

<a class="badge-evt badge-track"
   data-evt="PostList,TapPost,Tag,,PostTitle"
   data-track="post,v,,,d,anjMq4b,l"
   data-entry-id="anjMq4b"
   data-position="4"
   href="/gag/anjMq4b"
   target="_blank">
    Dammit Gilly !!! Tell your brother to step down ....
</a>

<a class="badge-evt badge-track"
   data-evt="PostList,TapPost,Tag,,PostTitle"
   data-track="post,v,,,d,aQe3PXK,l"
   data-entry-id="aQe3PXK"
   data-position="3"
   href="/gag/aQe3PXK"
   target="_blank">
    [SPOILER] Nothing more to say
</a>
`

var gag9 *Gag9

func init() {
	gag9 = New()
}

func TestFind(t *testing.T) {
	memes := gag9.Find()

	if memes == nil || len(memes) == 0 {
		t.Error("expected list of memes")
	}
}

func TestFindByTag(t *testing.T) {
	memes := gag9.FindByTag("love")

	if memes == nil || len(memes) == 0 {
		t.Error("expected list of memes")
	}
}

func TestCrawl(t *testing.T) {
	expectedMemes := []Meme{
		Meme{
			Description: "When you mix MJ with Got",
			Image:       "https://img-9gag-fun.9cache.com/photo/aKDjw1N_460s.jpg",
		},
		Meme{
			Description: "Dammit Gilly !!! Tell your brother to step down ....",
			Image:       "https://img-9gag-fun.9cache.com/photo/anjMq4b_460s.jpg",
		},
		Meme{
			Description: "[SPOILER] Nothing more to say",
			Image:       "https://img-9gag-fun.9cache.com/photo/aQe3PXK_460s.jpg",
		},
	}

	gotMemes := crawl(bytes.NewBufferString(RAW_HTML))

	if !reflect.DeepEqual(expectedMemes, gotMemes) {
		t.Error(
			"expected", expectedMemes,
			"got", gotMemes,
		)
	}

}
