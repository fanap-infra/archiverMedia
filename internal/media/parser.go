package media

import (
	"github.com/fanap-infra/log"
	"google.golang.org/protobuf/proto"
)

func Parse(data []byte) (*Packet, error) {
	packet := &Packet{}
	if err := proto.Unmarshal(data, packet); err != nil {
		log.Warn("Parse data is wrong")
		return packet, err
	}
	return packet, nil
}
