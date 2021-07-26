package archiverMedia

import (
	"os"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/utils"
	"github.com/fanap-infra/log"
	"github.com/stretchr/testify/assert"
)

const (
	fsPath               = "fs.beh"
	headerPath           = "Header.Beh"
	blockSizeTest  = 5120
	fileSizeTest   = blockSizeTest * 128
)

type EventsListener struct {
	t      *testing.T
	fileID uint32
}

func (el *EventsListener) DeleteFile(fileID uint32) {
	assert.Equal(el.t, el.fileID, fileID)
}

func TestCreateFS(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + "/"+fsPath)
	_ = utils.DeleteFile(homePath + "/"+headerPath)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	_, err = provider.CreateFileSystem(homePath, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, utils.FileExists(homePath+ "/"+fsPath))
	assert.Equal(t, true, utils.FileExists(homePath+"/"+headerPath))
	size, err := utils.FileSize(homePath +  "/"+fsPath)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(fileSizeTest), size)
	_ = utils.DeleteFile(homePath +  "/"+fsPath)
	_ = utils.DeleteFile(homePath + "/"+headerPath)
}

func TestParseFS(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath +  "/"+fsPath)
	_ = utils.DeleteFile(homePath + "/"+headerPath)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	_, err = provider.CreateFileSystem(homePath, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	fs, err := provider.ParseFileSystem(homePath, &eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	assert.Equal(t, fs.blockSize, uint32(blockSizeTest))
	_ = utils.DeleteFile(homePath +  "/"+fsPath)
	_ = utils.DeleteFile(homePath + "/"+headerPath)
}
