package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
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
	conn, err := grpc.Dial("10.230.33.191:9000", opts...)
	var reply Reply
	// 调用 + &reply
	err = grpc.Invoke(context.Background(), "/onip.item_asset.service.v1.ItemAssetService/GetItemSeriesAssetSingle", []byte("{\"item_id\": 12}"), &reply, conn)
	if err != nil {
		panic(err)
	}
	// 打印prettyJSON
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, reply.res, "", "\t🐱")
	fmt.Println(prettyJSON.String())
}
