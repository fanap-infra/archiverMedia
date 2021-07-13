package archiverMedia

import (
	"math/rand"
	"os"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/media"

	"github.com/fanap-infra/archiverMedia/pkg/utils"
	"github.com/fanap-infra/log"
	"github.com/stretchr/testify/assert"
)

func TestIO_OneVirtualMediaFile(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	arch, err := provider.CreateFileSystem(homePath+fsPathTest, fileSizeTest, blockSizeTest, &eventListener,
		log.GetScope("test"))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, utils.FileExists(homePath+fsPathTest))
	assert.Equal(t, true, utils.FileExists(homePath+headerPathTest))
	var packets []*media.Packet

	MaxID := 1000
	MaxByteArraySize := int(blockSizeTest * 0.1)
	VFSize := int(3.5 * blockSizeTest)
	vfID := uint32(rand.Intn(MaxID))
	vm, err := arch.NewVirtualMediaFile(vfID, "test")
	assert.Equal(t, nil, err)
	size := 0

	for {
		token := make([]byte, uint32(rand.Intn(MaxByteArraySize)))
		m, err := rand.Read(token)
		assert.Equal(t, nil, err)
		pkt := &media.Packet{Data: token, PacketType: media.PacketType_PacketVideo, IsKeyFrame: true}
		packets = append(packets, pkt)
		size = size + m
		err = vm.WriteFrame(pkt)
		assert.Equal(t, nil, err)

		if size > VFSize {
			break
		}
	}

	err = vm.Close()
	assert.Equal(t, nil, err)

	vm2, err := arch.OpenVirtualMediaFile(vfID)
	assert.Equal(t, nil, err)

	for i, packet := range packets {
		pkt, err := vm2.ReadFrame()
		assert.Equal(t, nil, err)
		if err != nil {
			assert.Equal(t, i+1, len(packets))
			break
		}
		assert.Equal(t, packet.Data, pkt.Data)
	}

	err = arch.Close()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}
