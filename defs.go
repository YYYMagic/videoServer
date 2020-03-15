package main

type video struct {
	Vid string
	Name string
	Url string
}
var videoMap map[string]video

func init() {

	videoMap = make(map[string]video)
	videoMap ["kobe"] = video{Vid: "kobe", Name: "唠嗑集锦",Url:  "http://47.101.41.7:8000/videos/kobe.mp4"}
	videoMap ["akb48"] = video{Vid: "akb48", Name: "初级福利",Url:  "http://47.101.41.7:8000/videos/akb48.mp4"}

}
