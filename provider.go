package archiverMedia

import (
	"fmt"

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

func (p *Provider) CreateFileSystem(path string, size int64, blockSize uint32, eventsHandler fsEngine.Events,
	log *log.Logger) (*Archiver, error) {
	_, ok := p.openedArchiver[path]
	if ok {
		return nil, fmt.Errorf("archiver created before")
	}
	fs, err := fsEngine.CreateFileSystem(path, size, blockSize, log)
	if err != nil {
		return nil, err
	}

	arch := &Archiver{log: log, fs: fs, EventsHandler: eventsHandler, blockSize: blockSize}
	p.openedArchiver[path] = arch
	return arch, nil
}

func (p *Provider) ParseFileSystem(path string, eventsHandler fsEngine.Events, log *log.Logger) (*Archiver, error) {
	arch, ok := p.openedArchiver[path]
	if ok {
		return arch, nil
	}

	fs, err := fsEngine.ParseFileSystem(path, log)
	if err != nil {
		return nil, err
	}

	arch = &Archiver{log: log, fs: fs, EventsHandler: eventsHandler}
	p.openedArchiver[path] = arch
	return arch, nil
}
