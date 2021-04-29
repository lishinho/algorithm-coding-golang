package bili

import (
	"fmt"
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
	fmt.Printf("url: %s\n", _huaweiLink)
	fmt.Printf("url after: %s\n", url.QueryEscape(_huaweiLink))
}
