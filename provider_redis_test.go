package archiverMedia

import (
	"github.com/fanap-infra/fsEngine/mocks"
	"os"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/utils"
	"github.com/fanap-infra/log"
	"github.com/stretchr/testify/assert"
)


func TestCreateFS_Redis(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	redisMock := mocks.NewRedisMock()
	_, err = provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"), &redisMock)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, utils.FileExists(homePath+"/"+fsPath))
	assert.Equal(t, true, utils.FileExists(homePath+"/"+headerPath))
	size, err := utils.FileSize(homePath + "/" + fsPath)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(fileSizeTest), size)
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)
}

func TestParseFS_Redis(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	redisMock := mocks.NewRedisMock()
	_, err = provider.CreateFileSystem(fsID,homePath, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"), &redisMock)
	assert.Equal(t, nil, err)
	fs, err := provider.ParseFileSystem(fsID,homePath, &eventListener, log.GetScope("test"), &redisMock)
	assert.Equal(t, nil, err)
	assert.Equal(t, fs.blockSize, uint32(blockSizeTest))
	_ = utils.DeleteFile(homePath + "/" + fsPath)
	_ = utils.DeleteFile(homePath + "/" + headerPath)
}
