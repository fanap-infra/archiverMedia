package archiverMedia

import (
	"github.com/fanap-infra/fsEngine"
	"github.com/fanap-infra/log"
)

type Archiver struct {
	log           *log.Logger
	EventsHandler Events
	fs            *fsEngine.FSEngine
	blockSize     uint32
}
