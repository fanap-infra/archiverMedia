package virtualMedia

import (
	"sync"

	"github.com/fanap-infra/archiverMedia/pkg/vInfo"

	"github.com/fanap-infra/archiverMedia/pkg/media"
	"github.com/fanap-infra/fsEngine/pkg/virtualFile"
	"github.com/fanap-infra/log"
)

type VirtualMedia struct {
	vfBuf               []byte
	vFile               *virtualFile.VirtualFile
	log                 *log.Logger
	frameChunk          *media.PacketChunk
	currentFrameInChunk uint32
	frameChunkRX        *media.PacketChunk
	readOnly            bool
	fwMUX               sync.Mutex
	rxMUX               sync.Mutex
	name                string
	fileSize            uint32
	blockSize           uint32
	fileID              uint32
	archiver            Arch
	info                *vInfo.Info
}
