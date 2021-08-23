package archiverMedia

import (
	"fmt"
	Header_ "github.com/fanap-infra/fsEngine/pkg/Header"
	"sync"

	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"

	"github.com/fanap-infra/fsEngine"
	"github.com/fanap-infra/log"
)

type Provider struct {
	crudMutex      sync.Mutex
	log            *log.Logger
	openedArchiver map[string]*Archiver
}

func NewProvider() *Provider {
	return &Provider{
		log:            log.GetScope("Archiver"),
		openedArchiver: make(map[string]*Archiver),
	}
}

func (p *Provider) CreateFileSystem(id uint32,path string, size int64, blockSize uint32, eventsHandler Events,
	log *log.Logger, redisDB Header_.RedisDB) (*Archiver, error) {
	p.crudMutex.Lock()
	defer p.crudMutex.Unlock()
	_, ok := p.openedArchiver[path]
	if ok {
		return nil, fmt.Errorf("archiver created before")
	}

	arch := &Archiver{
		log: log, EventsHandler: eventsHandler, openFiles: make(map[uint32][]*virtualMedia.VirtualMedia),
		blockSize: blockSize,
	}
	fs, err := fsEngine.CreateFileSystem(id,path, size, blockSize, arch, log, redisDB)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	p.openedArchiver[path] = arch
	return arch, nil
}

func (p *Provider) ParseFileSystem(id uint32,path string, eventsHandler Events, log *log.Logger, redisDB Header_.RedisDB) (*Archiver, error) {
	p.crudMutex.Lock()
	defer p.crudMutex.Unlock()
	arch, ok := p.openedArchiver[path]
	if ok {
		return arch, nil
	}

	arch = &Archiver{log: log, EventsHandler: eventsHandler, openFiles: make(map[uint32][]*virtualMedia.VirtualMedia)}
	fs, err := fsEngine.ParseFileSystem(id,path, arch, log, redisDB)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	arch.blockSize = fs.GetBlockSize()
	p.openedArchiver[path] = arch
	return arch, nil
}

func (p *Provider) CloseArchiver(path string) error {
	p.crudMutex.Lock()
	defer p.crudMutex.Unlock()
	arch, ok := p.openedArchiver[path]
	if !ok {
		return fmt.Errorf("archiver with path: %v is not opened", path)
	}
	delete(p.openedArchiver, path)
	return arch.Close()
}
