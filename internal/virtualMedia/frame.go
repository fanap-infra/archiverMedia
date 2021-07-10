package virtualMedia

import (
	"encoding/binary"

	"github.com/fanap-infra/archiverMedia/internal/media"
	"google.golang.org/protobuf/proto"
)

func generateFrameChunk(med *media.PacketChunk) ([]byte, error) {
	b, err := proto.Marshal(med)
	if err != nil {
		return nil, err
	}
	binSize := make([]byte, 4)
	binary.BigEndian.PutUint32(binSize, uint32(len(b)+4))
	// b = append(b, binSize...)
	b = append(binSize, b...)
	b = append([]byte(FrameChunkIdentifier), b...)
	return b, nil
}
