package main

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(AvToBv(602109295))
	fmt.Println(BvToAv("BV1bg411A7pa"))
	//前闭后开()
	fmt.Println(len(base64.RawURLEncoding.EncodeToString([]byte(uuid.NewString()))))
	fmt.Println(len(uuid.NewString()))
	hash()
}
