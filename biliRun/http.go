package main

import (
	"fmt"
	"log"
	"net/url"
)

type ShowHdMap struct {
	Key    int32
	Heads  []string
	Limits []*Limit
}

type Limit struct {
	ID        int64 `json:"id"`
	Rid       int64
	Build     int    `json:"build"`
	Condition string `json:"condition"`
}

//func main() {
//	examp := &ShowHdMap{
//		Limits: make([]*Limit, 2),
//	}
//	examp.Limits[0] = &Limit{
//		ID:        1,
//		Rid:       3,
//		Build:     3,
//		Condition: "",
//	}
//
//	fmt.Printf("show hd %+v", examp)
//}

// url encode test
func urlEncodeTest() {
	deeplink, _ := url.QueryUnescape(_huaweiLink)
	fmt.Printf("url: %s\n", deeplink)
	fmt.Printf("url: %s\n", _huaweiLink[2:])
	fmt.Printf("https://live.bilibili.com/p/html/live-fansmedal-wall/index.html?is_live_webview=1\u0026tId=%d#/medal", 336)
}

func urlTest() {
	u, err := url.Parse("bilibili://bangumi/season/39311?h5awaken=b3Blbl9hcHBfZnJvbV90eXBlPWRlZXBsaW5rX3R0bGgtYW5kLXJkLTMxeWF6LTE2Nzc5NTY3Mzc5MjgxOTkmb3Blbl9hcHBfdXJsPXNzMzkzMTE=&from_spmid=out_open_deeplink_ttlh-and-rd-31yaz-1677956737928199&")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("before url: %+v\n", u)
	fmt.Printf("before path: %+v, %+v\n", u.Path, u.Opaque)
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
}
