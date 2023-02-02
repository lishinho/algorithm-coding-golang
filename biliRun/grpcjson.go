package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"io"
	"net/http"
)

type JSON struct {
	jsonpb.Marshaler
	jsonpb.Unmarshaler
}

// Marshal is json marshal
func (j JSON) Marshal(v interface{}) (out []byte, err error) {
	return v.([]byte), nil
}

// Reply for test
type Reply struct {
	res []byte
}

// Unmarshal is json unmarshal
func (j JSON) Unmarshal(data []byte, v interface{}) (err error) {
	v.(*Reply).res = data
	return nil
}

// Name is name of JSON
func (j JSON) Name() string {
	return "json"
}

func grpcJsonGet() {
	// grpc初始化 注册codec
	// go的实现暴露了一个codec接口去实现payload的marshal成不同格式
	// The Codec will be stored and looked up by result of its Name() method,
	// which should match the content-subtype of the encoding handled by the Codec.
	encoding.RegisterCodec(JSON{
		Marshaler: jsonpb.Marshaler{
			EmitDefaults: true,
			OrigName:     true,
		},
	})
	// grpc客户端的拨号选择器
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype(JSON{}.Name())),
	}
	opts = append(opts, grpc.WithInsecure()) // 不用tls
	// grpc客户端指定ip端口连接
	conn, err := grpc.Dial("172.22.17.120:9000", opts...)
	var reply Reply
	// 调用 + &reply
	err = grpc.Invoke(context.Background(), "/main.community.emote_service.v1.EmoteService/UserPanelPackageBiz", []byte("{\n    \"mid\": \"382719467\",\n    \"up_mid\": \"423895\",\n    \"business\": \"reply\"\n}"), &reply, conn)
	if err != nil {
		panic(err)
	}
	// 打印prettyJSON
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, reply.res, "", "\t🐱")
	fmt.Println(prettyJSON.String())
}

func httpJsonGet() {
	// 填想要请求的路径
	addr := "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=50dda777-efd3-4ab2-a738-c740416dace9&type=file"
	// http 请求
	var Body io.Reader
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, addr, Body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Disposition", "form-data; name=\"media\";filename=\"wework.txt\"; filelength=6")
	// http do
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// 读body
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	// pretty json
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, b, "", "    ")
	fmt.Println(prettyJSON.String())
}
