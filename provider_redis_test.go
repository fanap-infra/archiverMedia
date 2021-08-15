package archiverMedia
//
//import (
//	"github.com/fanap-infra/fsEngine/pkg/redisClient"
//	"os"
//	"testing"
//
//	"github.com/fanap-infra/archiverMedia/pkg/utils"
//	"github.com/fanap-infra/log"
//	"github.com/stretchr/testify/assert"
//)
//
//var redisOptions = &redisClient.RedisOptions{
//	Addr:     "127.0.0.1:6379",
//	Password: "",
//	DB:       0,
//}
//
//func TestCreateFS_Redis(t *testing.T) {
//	homePath, err := os.UserHomeDir()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//	eventListener := EventsListener{t: t}
//	provider := NewProvider()
//	_, err = provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest,
//		&eventListener, log.GetScope("test"), redisOptions)
//	assert.Equal(t, nil, err)
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+fsPath))
//	assert.Equal(t, true, utils.FileExists(homePath+"/"+headerPath))
//	size, err := utils.FileSize(homePath + "/" + fsPath)
//	assert.Equal(t, nil, err)
//	assert.Equal(t, int64(fileSizeTest), size)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//}
//
//func TestParseFS_Redis(t *testing.T) {
//	homePath, err := os.UserHomeDir()
//	assert.Equal(t, nil, err)
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//	eventListener := EventsListener{t: t}
//	provider := NewProvider()
//	_, err = provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest,
//		&eventListener, log.GetScope("test"), redisOptions)
//	assert.Equal(t, nil, err)
//	fs, err := provider.ParseFileSystem(fsID,homePath, &eventListener, log.GetScope("test"), redisOptions)
//	assert.Equal(t, nil, err)
//	assert.Equal(t, fs.blockSize, uint32(blockSizeTest))
//	_ = utils.DeleteFile(homePath + "/" + fsPath)
//	_ = utils.DeleteFile(homePath + "/" + headerPath)
//}
