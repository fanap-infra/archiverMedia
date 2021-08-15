package archiverMedia
//
//import (
//	"math/rand"
//	"os"
//	"strconv"
//	"testing"
//
//	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"
//
//	"github.com/fanap-infra/archiverMedia/pkg/utils"
//	"github.com/fanap-infra/log"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestVirtualMedia_Remove_Redis(t *testing.T) {
//	homePath, err := os.UserHomeDir()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//	eventListener := EventsListener{t: t}
//	provider := NewProvider()
//	arch, err := provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest, &eventListener,
//		log.GetScope("test"), redisOptions)
//	assert.Equal(t, nil, err)
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+fsPath))
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+headerPath))
//	var testIDs []uint32
//	var testNames []string
//	TestSize := 5
//	MaxID := 1000
//	var vms []*virtualMedia.VirtualMedia
//	for i := 0; i < TestSize; i++ {
//		tmp := uint32(rand.Intn(MaxID))
//		if utils.ItemExists(testIDs, tmp) {
//			i = i - 1
//			continue
//		}
//		testIDs = append(testIDs, tmp)
//		testNames = append(testNames, "test"+strconv.Itoa(i))
//		vm, err := arch.NewVirtualMediaFile(testIDs[i], testNames[i])
//		assert.Equal(t, nil, err)
//		vms = append(vms, vm)
//	}
//
//	for i := 0; i < TestSize; i++ {
//		_, err := arch.NewVirtualMediaFile(testIDs[i], testNames[i])
//		assert.NotEqual(t, nil, err)
//	}
//
//	// cna not remove opened virtual files
//	for i := 0; i < TestSize; i++ {
//		err := arch.RemoveVirtualMediaFile(testIDs[i])
//		assert.NotEqual(t, nil, err)
//	}
//
//	for i := 0; i < TestSize/2; i++ {
//		err := vms[i].Close()
//		assert.Equal(t, nil, err)
//	}
//
//	for i := 0; i < TestSize; i++ {
//		if i < TestSize/2 {
//			err := arch.RemoveVirtualMediaFile(testIDs[i])
//			assert.Equal(t, nil, err)
//		} else {
//			err := arch.RemoveVirtualMediaFile(testIDs[i])
//			assert.NotEqual(t, nil, err)
//			err = vms[i].Close()
//			assert.Equal(t, nil, err)
//			err = arch.RemoveVirtualMediaFile(testIDs[i])
//			assert.Equal(t, nil, err)
//		}
//	}
//
//	err = arch.Close()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//}
//
//func TestVirtualMedia_Open_Redis(t *testing.T) {
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
//	var testIDs []uint32
//	var testNames []string
//
//	TestSize := 1
//	MaxID := 1000
//	var vms []*virtualMedia.VirtualMedia
//	for i := 0; i < TestSize; i++ {
//		tmp := uint32(rand.Intn(MaxID))
//		if utils.ItemExists(testIDs, tmp) {
//			i = i - 1
//			continue
//		}
//		testIDs = append(testIDs, tmp)
//		testNames = append(testNames, "test"+strconv.Itoa(i))
//		vm, err := arch.NewVirtualMediaFile(testIDs[i], testNames[i])
//		assert.Equal(t, nil, err)
//		vms = append(vms, vm)
//	}
//
//	for i := 0; i < len(vms); i++ {
//		err := vms[i].Close()
//		assert.Equal(t, nil, err)
//	}
//
//	for i := 0; i < len(testIDs); i++ {
//		_, err := arch.OpenVirtualMediaFile(testIDs[i])
//		assert.Equal(t, nil, err)
//	}
//
//	err = arch.Close()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//}
