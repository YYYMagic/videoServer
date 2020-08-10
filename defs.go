package main

type video struct {
	Vid string
	Name string
	Url string
}
var videoMap map[string]video

func init() {
	videoMap = make(map[string]video)
	videoMap ["黑夜问白天"] = video{Vid: "hei", Name: "黑夜问白天",Url:  "https://yyyvideos.oss-cn-beijing.aliyuncs.com/%E9%BB%91%E5%A4%9C%E5%95%8F%E7%99%BD%E5%A4%A9%20cover.mp4"}
	videoMap ["只要有你的地方"] = video{Vid: "zhi", Name: "只要有你的地方",Url:  "https://yyyvideos.oss-cn-beijing.aliyuncs.com/%E5%8F%AA%E8%A6%81%E6%9C%89%E4%BD%A0%E7%9A%84%E5%9C%B0%E6%96%B9%20cover.mp4"}
	videoMap ["灰色と青"] = video{Vid: "hui", Name: "灰色と青",Url:  "https://yyyvideos.oss-cn-beijing.aliyuncs.com/%E7%81%B0%E8%89%B2%E3%81%A8%E9%9D%92%20cover.mp4"}
}
