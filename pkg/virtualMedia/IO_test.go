package virtualMedia

import (
	"math/rand"
	"os"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/vInfo"

	"github.com/fanap-infra/archiverMedia/pkg/media"
	"github.com/fanap-infra/archiverMedia/pkg/utils"
	"github.com/fanap-infra/fsEngine"
	"github.com/fanap-infra/log"
	"github.com/stretchr/testify/assert"
)

const (
	blockSizeTest = 5120
	vfID          = 1

	fsPath       = "fs.beh"
	headerPath   = "Header.Beh"
	fileSizeTest = blockSizeTest * 128
	fsID = 11
)

type ArchMock struct {
	log       *log.Logger
	openFiles map[uint32]*VirtualMedia
	fs        *fsEngine.FSEngine
	tst       *testing.T
}

func (archMock *ArchMock) Closed(fileID uint32) error {
	return nil
}

func (archMock *ArchMock) VirtualFileDeleted(fileID uint32, message string) {
	archMock.log.Warnv("Media file deleted", "fileID", fileID, "message", message)
}

func (archMock *ArchMock) Close() error {
	return archMock.fs.Close()
}

func NewVBufMock(t *testing.T, path string) (*ArchMock, error) {
	arch := &ArchMock{
		openFiles: make(map[uint32]*VirtualMedia),
		tst:       t,
		log:       log.GetScope("test"),
	}
	// arch.log.Infov("create archiver ", "path", path)
	fs, err := fsEngine.CreateFileSystem(fsID, path, fileSizeTest, blockSizeTest, arch, arch.log, nil)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	return arch, nil
}

func TestIO_WR(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)

	archMock, err := NewVBufMock(t, homePath)
	assert.Equal(t, nil, err)
	vf, err := archMock.fs.NewVirtualFile(vfID, "test2")
	assert.Equal(t, nil, err)
	vm := NewVirtualMedia("test", vfID, blockSizeTest, vf, archMock, log.GetScope("test2"))
	archMock.openFiles[vfID] = vm
	var packets []*media.Packet

	size := 0
	VFSize := int(3.5 * blockSizeTest)
	MaxByteArraySize := int(blockSizeTest * 0.1)

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
	vfInfo := &vInfo.Info{}
	vm2 := OpenVirtualMedia("test", vfID, blockSizeTest, vf, archMock, vfInfo, log.GetScope("test2"))

	for i, packet := range packets {
		pkt, err := vm2.ReadFrame()
		assert.Equal(t, nil, err)
		if err != nil {
			assert.Equal(t, i+1, len(packets))
			break
		}
		assert.Equal(t, packet.Data, pkt.Data)
	}

	err = archMock.Close()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)
}
