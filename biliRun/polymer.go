package main

import (
	"context"
	"fmt"
	"strings"
)

type AvToBvRequest struct {
	Avid int64 `form:"avid"`
}

type AvToBvResp struct {
	Bvid string `json:"bvid"`
}

func AvToBv2(ctx context.Context, request *AvToBvRequest) (resp *AvToBvResp, err error) {
	resp = &AvToBvResp{}
	if request.Avid < _minAid {
		err = fmt.Errorf("avid lte zero")
		return
	}
	if request.Avid >= _maxAid {
		err = fmt.Errorf("avid is too big")
		return
	}
	var (
		tmpbv string
		ids   []string
	)
	tmp := ((1 << 51) | request.Avid) ^ _xorcode
	for tmp != 0 {
		tmpbv = fmt.Sprintf("%s%s", string(_alphabet[tmp%_base]), tmpbv)
		tmp = tmp / _base
	}
	ids = strings.Split(tmpbv, "")
	// 位置转换
	tmpstr := ids[0]
	ids[0] = ids[6]
	ids[6] = tmpstr
	tmpstr = ids[1]
	ids[1] = ids[4]
	ids[4] = tmpstr
	resp.Bvid = _prefix + strings.Join(ids, "")
	return
}
