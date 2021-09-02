package archiverMedia

import (
	"fmt"
	"github.com/fanap-infra/archiverMedia/pkg/media"
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

func (p *Provider) RecoverHeaderFileSystem(id uint32,path string, blockSize uint32, eventsHandler Events, log *log.Logger, redisDB Header_.RedisDB) (*Archiver, error) {
	p.crudMutex.Lock()
	defer p.crudMutex.Unlock()
	arch, ok := p.openedArchiver[path]
	if ok {
		return arch, nil
	}

	arch = &Archiver{log: log, EventsHandler: eventsHandler, openFiles: make(map[uint32][]*virtualMedia.VirtualMedia)}
	fs, err := fsEngine.RecoverHeaderFileSystem(id, path, blockSize, arch, log, redisDB)
	if err != nil {
		return nil, err
	}
	arch.fs = fs
	arch.blockSize = fs.GetBlockSize()
	p.openedArchiver[path] = arch
	filesIndex := fs.GetFileList()
	for i, fIndex := range filesIndex {
		vm, err :=arch.OpenVirtualMediaFileForHeaderRecovery(fIndex.Id)
		if err != nil {
			p.log.Errorv("can not open virtual media file", "err", err.Error())
			continue
		}

		vInfo := vm.GetInfo()
		vInfo.StartTime = 0
		vInfo.EndTime = 0
		for {
			frame, err := vm.ReadFrame()
			if err == virtualMedia.EndOfFile {
				log.Infov("end of file", "id", fIndex.Id, "i",i)
				break
			} else if err != nil {
				log.Errorv("can not read frame", "id", fIndex.Id, "i",i, "err", err.Error())
				break
			}
			if frame.PacketType == media.PacketType_PacketVideo {
				vInfo.EndTime = frame.Time
			}
		}

		err = vm.UpdateFileOptionalData()
		if err != nil {
			log.Errorv("can not update file optional data",
				"id", fIndex.Id, "i",i, "err", err.Error())
		}
		err =vm.Close()
		if err != nil {
			log.Errorv("can not close virtual media file",
				"id", fIndex.Id, "i",i, "err", err.Error())
		}
	}
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
