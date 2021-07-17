package archiverMedia

import (
	"sync"

	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"

	"github.com/fanap-infra/fsEngine"
	"github.com/fanap-infra/log"
)

type Archiver struct {
	log           *log.Logger
	EventsHandler Events
	fs            *fsEngine.FSEngine
	blockSize     uint32
	crudMutex     sync.Mutex
	openFiles     map[uint32]*virtualMedia.VirtualMedia
}

func (arch *Archiver) GetPath() string {
	// ToDo:: add get file path
	// return arch.fs
	return ""
}
