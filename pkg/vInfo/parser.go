package vInfo

import (
	"github.com/fanap-infra/log"
	"google.golang.org/protobuf/proto"
)

func Parse(data []byte) (*Info, error) {
	info := &Info{}
	if err := proto.Unmarshal(data, info); err != nil {
		log.Warn("info data is wrong")
		return info, err
	}
	return info, nil
}
