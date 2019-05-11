package main

import (
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

type song struct {
	Title      string
	ArtistName string
	AlbumName  string
}

var songs = []song{
	{"オレンジ", "Greeeen ", "オレンジ"},
	{"オレンジ", "Greeeen ", "C、Dですと!?"},
	{"オレンジ", "逢坂大河(釘宮理恵)/櫛枝実乃梨(堀江由衣)/川嶋亜美(喜多村英梨) ", "オレンジ - EP"},
	{"オレンジ", "逢坂大河(釘宮理恵)/櫛枝実乃梨(堀江由衣)/川嶋亜美(喜多村英梨) ", "√HAPPYEND とらドラ!BEST ALBUM"},
	{"オレンジ", "クリープハイプ ", "オレンジ/愛の標識 - EP"},
	{"オレンジ", "クリープハイプ ", "死ぬまで一生愛されてると思ってたよ"},
	{"オレンジ", "クリープハイプ ", "クリープハイプ名作選"},
	{"オレンジ", "トーマ ", "アザレアの心臓"},
	{"オレンジ", "sumika ", "SALLY e.p"},
	{"オレンジ", "安全地帯 ", "オレンジ/恋の予感(2010ヴァージョン)(通常版)"},
	{"オレンジ", "安全地帯 ", "ALL TIME BEST"},
	{"オレンジ", "パスピエ ", "ネオンと虎"},
}

func main() {
	idx, err := fuzzyfinder.FindMulti(songs, func(i int) string {
		return songs[i].Title
	}, fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
		if i == -1 {
			return ""
		}
		return fmt.Sprintf("Title:  %s\nArtist: %s\nAlbum:  %s\n", songs[i].Title, songs[i].ArtistName, songs[i].AlbumName)
	}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find: %s\n", err)
		os.Exit(1)
	}
	for _, i := range idx {
		fmt.Println(fmt.Sprintf("%s / %s / %s", songs[i].Title, songs[i].ArtistName, songs[i].AlbumName))
	}
}
