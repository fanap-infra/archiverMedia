package archiverMedia

import (
	"fmt"

	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"

	"github.com/fanap-infra/fsEngine"
	"github.com/fanap-infra/log"
)

type Provider struct {
	log            *log.Logger
	openedArchiver map[string]*Archiver
}

func NewProvider() *Provider {
	return &Provider{
		log:            log.GetScope("Archiver"),
		openedArchiver: make(map[string]*Archiver),
	}
}

func (p *Provider) CreateFileSystem(path string, size int64, blockSize uint32, eventsHandler Events,
	log *log.Logger) (*Archiver, error) {
	_, ok := p.openedArchiver[path]
	if ok {
		return nil, fmt.Errorf("archiver created before")
	}

	arch := &Archiver{
		log: log, EventsHandler: eventsHandler, openFiles: make(map[uint32]*virtualMedia.VirtualMedia),
		blockSize: blockSize,
	}
	fs, err := fsEngine.CreateFileSystem(path, size, blockSize, arch, log)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	p.openedArchiver[path] = arch
	return arch, nil
}

func (p *Provider) ParseFileSystem(path string, eventsHandler Events, log *log.Logger) (*Archiver, error) {
	arch, ok := p.openedArchiver[path]
	if ok {
		return arch, nil
	}

	arch = &Archiver{log: log, EventsHandler: eventsHandler, openFiles: make(map[uint32]*virtualMedia.VirtualMedia)}
	fs, err := fsEngine.ParseFileSystem(path, arch, log)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	p.openedArchiver[path] = arch
	return arch, nil
}
