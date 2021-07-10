package virtualMedia

import (
	"github.com/fanap-infra/archiverMedia/internal/media"
	"github.com/fanap-infra/fsEngine/internal/virtualFile"
	"github.com/fanap-infra/log"
)

type VirtualMedia struct {
	vFile      *virtualFile.VirtualFile
	log        *log.Logger
	frameChunk *media.PacketChunk
	readOnly   bool
}
