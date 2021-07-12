package archiverMedia

import (
	"os"
	"testing"

	"github.com/fanap-infra/archiverMedia/pkg/utils"
	"github.com/fanap-infra/log"
	"github.com/stretchr/testify/assert"
)

const (
	fsPathTest     = "/fsTest.beh"
	headerPathTest = "/Header.Beh"
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
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	_, err = provider.CreateFileSystem(homePath+fsPathTest, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, utils.FileExists(homePath+fsPathTest))
	assert.Equal(t, true, utils.FileExists(homePath+headerPathTest))
	size, err := utils.FileSize(homePath + fsPathTest)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(fileSizeTest), size)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}

func TestParseFS(t *testing.T) {
	homePath, err := os.UserHomeDir()
	assert.Equal(t, nil, err)
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
	eventListener := EventsListener{t: t}
	provider := NewProvider()
	_, err = provider.CreateFileSystem(homePath+fsPathTest, fileSizeTest, blockSizeTest,
		&eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	fs, err := provider.ParseFileSystem(homePath+fsPathTest, &eventListener, log.GetScope("test"))
	assert.Equal(t, nil, err)
	assert.Equal(t, fs.blockSize, uint32(blockSizeTest))
	_ = utils.DeleteFile(homePath + fsPathTest)
	_ = utils.DeleteFile(homePath + headerPathTest)
}
