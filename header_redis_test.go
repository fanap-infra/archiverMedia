package archiverMedia
//
//import (
//	"os"
//	"testing"
//
//	"github.com/fanap-infra/archiverMedia/pkg/media"
//	"github.com/fanap-infra/archiverMedia/pkg/utils"
//	"github.com/fanap-infra/log"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestHeader_VF_Redis(t *testing.T) {
//	homePath, err := os.UserHomeDir()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//	var eventListener EventsListener
//	provider := NewProvider()
//	arch, err := provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest, &eventListener,
//		log.GetScope("test"), redisOptions)
//	assert.Equal(t, nil, err)
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+fsPath))
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+headerPath))
//
//	packetTime := int64(0)
//	token := make([]byte, 100)
//	vm, err := arch.NewVirtualMediaFile(1, "test1")
//	assert.Equal(t, nil, err)
//	err = vm.WriteFrame(&media.Packet{
//		Data: token, PacketType: media.PacketType_PacketVideo, IsKeyFrame: true,
//		Time: packetTime,
//	})
//	assert.Equal(t, nil, err)
//	packetTime = packetTime + 30
//	err = vm.WriteFrame(&media.Packet{
//		Data: token, PacketType: media.PacketType_PacketVideo, IsKeyFrame: true,
//		Time: packetTime,
//	})
//	assert.Equal(t, nil, err)
//
//	err = vm.Close()
//	assert.Equal(t, nil, err)
//
//	vm2, err := arch.OpenVirtualMediaFile(1)
//	assert.Equal(t, nil, err)
//	vInfo := vm2.GetInfo()
//	assert.Equal(t, int64(0), vInfo.StartTime)
//	assert.Equal(t, packetTime, vInfo.EndTime)
//	err = vm2.Close()
//	assert.Equal(t, nil, err)
//
//	err = arch.Close()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//}
