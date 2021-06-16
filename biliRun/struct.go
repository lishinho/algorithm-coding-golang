package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	time2 "time"
)

type Hot struct {
	Code    int    `json:"code,omitempty"`
	SeID    string `json:"seid,omitempty"`
	TrackID string `json:"trackid"`
	Ffffff  struct {
		Keyword   string `json:"keyword,omitempty"`
		IsExisted bool   `json:"is_ex,omitempty"`
	} `json:"fff,omitempty"`
}

type ReserveActExtra struct {
	Skin   ReserveActSkin `json:"skin,omitempty"`
	ActUrl string         `json:"act_url,omitempty"`
}

type HiddenAttribute struct {
	IsSpaceHidden  bool   `json:"is_space_hidden,omitempty"`
	Text           string `json:"text,omitempty"`
	AiExtraMetrics map[int]map[string]interface{}
}

type ReserveActSkin struct {
	Svga      string `json:"svga,omitempty"`
	LastImg   string `json:"last_img,omitempty"`
	PlayTimes bool   `json:"play_times,omitempty"`
}

type PopularWatchTime struct {
	Aid  int64 `json:"aid" form:"aid"`
	Time int64 `json:"time" form:"time"`
}

var _thumbUpSlice = []string{"story_like_combo_22", "story_like_combo_33", "story_like_combo_tv"}

const (
	_authKey    = "%s:32368a9369d6571b85cb4f6075d591daf77da09bc338b8154c5769932ab20b82"
	_huaweiLink = "bilibili%3A%2F%2Fbangumi%2Fseason%2F31778%3Fh5awaken%3Db3Blbl9hcHBfZnJvbV90eXBlPWRlZXBsa"
)

func randomPick() {
	time := time2.Now()
	for i := 0; i < 10; i++ {
		fmt.Printf("_thumbUpSlice %s\n", _thumbUpSlice[rand.Intn(len(_thumbUpSlice))])
	}
	fmt.Printf("time %v", time2.Now().Sub(time))
}

func postRequestBody() {
	str := make([]string, 3)
	fmt.Println(len(str))
	str = append(str, "ddd")
	fmt.Println(str)
	str = []string{"ddd", "ggg", "ddd", "eee"}
	bt, err := json.Marshal(str)
	if err != nil {
		return
	}
	fmt.Println(http.NewRequest(http.MethodPost, "https://ad-dp-drcn.cloud.huawei.com.cn/ddl/deeplink", bytes.NewReader(bt)))

}

func whenMapPanic() {
	var mmap map[int64][]string
	fmt.Println(mmap[3])
	// call panic
	//mmap[3] = append(mmap[3], "ddd")
	var data []string
	data = append(data, "ddd")
	fmt.Println(data[0])
}

func whyStructNotNil() {
	extra := &ReserveActExtra{}
	fmt.Println(extra.ActUrl)
	extra.ActUrl = "ddd"
	fmt.Println(extra)
	extra.Skin.Svga = "ddd"
	fmt.Println(extra.Skin.Svga)
	//extra = nil
	//fmt.Println(extra.ActUrl)
}

func whenArrayPanic() {
	var hots []*Hot
	hot := &Hot{
		Code: 101,
	}
	fmt.Printf("ddd %+v\n", hots)
	hots = append(hots, hot)
	fmt.Printf("ddd %+v\n", hots)
}

func calculateTime() {
	current := time2.Now()
	time2.Sleep(100)
	fmt.Println(time2.Now().Sub(current))
}

// authorize returns authorization for upload file to bfs
func authorize() (authorization string) {
	bundle := "tv.danmaku.bili"
	nonce := time2.Now().Unix()
	return authorizeHW(bundle, nonce)
}

func authorizeHW(bundle string, nonce int64) string {
	key := fmt.Sprintf(_authKey, bundle)
	data := fmt.Sprintf("%d:%s", nonce, url.QueryEscape(_huaweiLink))
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	signature := strings.ToLower(hex.EncodeToString(mac.Sum(nil)))
	return fmt.Sprintf("Authorization:Digest username=%s,nonce=%d,response=%s,algorithm=HmacSHA256", bundle, nonce, signature)
}

// bytes to hex string
func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

type Param struct {
	name string
}

func updateMapValueMap() {
	s := map[int]Param{
		1: {"lx"},
	}
	// s[1].name = "ok"  这里这样赋值是错误的！！！
	r := s[1]
	r.name = "ok"
}
