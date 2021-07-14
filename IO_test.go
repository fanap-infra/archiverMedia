package archiverMedia

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"

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
	err = vm2.Close()
	assert.Equal(t, nil, err)
	err = arch.Close()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}

func TestIO_MultipleVirtualMediaFileConsecutively(t *testing.T) {
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

	MaxID := 1000
	MaxByteArraySize := int(blockSizeTest * 0.5)
	VFSize := int(3.5 * blockSizeTest)

	virtualMediaFiles := make([]*virtualMedia.VirtualMedia, 0)
	numberOfVMs := 5
	packets := make([][]*media.Packet, numberOfVMs)
	vmIDs := make([]uint32, 0)
	for i := 0; i < numberOfVMs; i++ {
		vmID := uint32(rand.Intn(MaxID))
		if utils.ItemExists(vmIDs, vmID) {
			i = i - 1
			continue
		}
		vmIDs = append(vmIDs, vmID)
		vm, err := arch.NewVirtualMediaFile(vmID, "test"+strconv.Itoa(i))
		if assert.Equal(t, nil, err) {
			virtualMediaFiles = append(virtualMediaFiles, vm)
		}
	}
	if len(virtualMediaFiles) != numberOfVMs {
		return
	}

	for j, vm := range virtualMediaFiles {
		size := 0
		for {
			token := make([]byte, uint32(rand.Intn(MaxByteArraySize)))
			m, err := rand.Read(token)
			assert.Equal(t, nil, err)

			size = size + m
			pkt := &media.Packet{Data: token, PacketType: media.PacketType_PacketVideo, IsKeyFrame: true}
			packets[j] = append(packets[j], pkt)
			err = vm.WriteFrame(pkt)
			assert.Equal(t, nil, err)

			if size > VFSize {
				break
			}
		}
		err = vm.Close()
		assert.Equal(t, nil, err)
	}

	for i, pckts := range packets {
		vf2, err := arch.OpenVirtualMediaFile(vmIDs[i])
		assert.Equal(t, nil, err)

		for _, v := range pckts {
			frame, err := vf2.ReadFrame()
			assert.Equal(t, nil, err)
			if err != nil {
				break
			}
			assert.Equal(t, v.Data, frame.Data)
			assert.Equal(t, v.Index, frame.Index)
		}
		err = vf2.Close()
		assert.Equal(t, nil, err)
	}

	err = arch.Close()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}

func TestIO_MultipleVirtualMediaFileConcurrency(t *testing.T) {
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

	MaxID := 1000
	MaxByteArraySize := int(blockSizeTest * 0.5)
	VFSize := int(3.5 * blockSizeTest)

	virtualMediaFiles := make([]*virtualMedia.VirtualMedia, 0)
	numberOfVMs := 5
	packets := make([][]*media.Packet, numberOfVMs)
	vmIDs := make([]uint32, 0)
	for i := 0; i < numberOfVMs; i++ {
		vmID := uint32(rand.Intn(MaxID))
		if utils.ItemExists(vmIDs, vmID) {
			i = i - 1
			continue
		}
		vmIDs = append(vmIDs, vmID)
		vm, err := arch.NewVirtualMediaFile(vmID, "test"+strconv.Itoa(i))
		if assert.Equal(t, nil, err) {
			virtualMediaFiles = append(virtualMediaFiles, vm)
		}
	}
	if len(virtualMediaFiles) != numberOfVMs {
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	for j, vm := range virtualMediaFiles {
		wg.Add(1)
		go func(j int, vm *virtualMedia.VirtualMedia) {
			defer wg.Done()
			size := 0
			for {
				token := make([]byte, uint32(rand.Intn(MaxByteArraySize)))
				m, err := rand.Read(token)
				assert.Equal(t, nil, err)

				size = size + m
				pkt := &media.Packet{Data: token, PacketType: media.PacketType_PacketVideo, IsKeyFrame: true}
				mu.Lock()
				packets[j] = append(packets[j], pkt)
				mu.Unlock()
				err = vm.WriteFrame(pkt)
				assert.Equal(t, nil, err)

				if size > VFSize {
					break
				}
			}
			err = vm.Close()
			assert.Equal(t, nil, err)
		}(j, vm)
	}

	wg.Wait()

	for i, pckts := range packets {
		vf2, err := arch.OpenVirtualMediaFile(vmIDs[i])
		wg.Add(1)
		assert.Equal(t, nil, err)
		go func(pckts []*media.Packet, vf2 *virtualMedia.VirtualMedia) {
			defer wg.Done()
			for _, v := range pckts {
				frame, err := vf2.ReadFrame()
				assert.Equal(t, nil, err)
				if err != nil {
					break
				}
				assert.Equal(t, v.Data, frame.Data)
				assert.Equal(t, v.Index, frame.Index)
			}
			err = vf2.Close()
		}(pckts, vf2)

		assert.Equal(t, nil, err)
	}
	wg.Wait()
	err = arch.Close()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}