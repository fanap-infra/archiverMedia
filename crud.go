package archiverMedia

import (
	"fmt"

	"github.com/fanap-infra/archiverMedia/pkg/virtualMedia"
)

// Create new virtual file and add opened files
func (arch *Archiver) NewVirtualMediaFile(id uint32, fileName string) (*virtualMedia.VirtualMedia, error) {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()
	vf, err := arch.fs.NewVirtualFile(id, fileName)
	if err != nil {
		return nil, err
	}
	vm := virtualMedia.NewVirtualMedia(fileName, id, arch.blockSize, vf, arch, arch.log)
	arch.openFiles[id] = vm
	return vm, nil
}

func (arch *Archiver) OpenVirtualMediaFile(id uint32) (*virtualMedia.VirtualMedia, error) {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()
	_, ok := arch.openFiles[id]
	if ok {
		return nil, fmt.Errorf("this ID: %v is opened before", id)
	}
	vf, err := arch.fs.OpenVirtualFile(id)
	if err != nil {
		return nil, err
	}
	// ToDO: get file name from virtual file
	vm := virtualMedia.OpenVirtualMedia("fileName", id, arch.blockSize, vf, arch, arch.log)
	return vm, nil
}

func (arch *Archiver) RemoveVirtualMediaFile(id uint32) error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()
	_, ok := arch.openFiles[id]
	if ok {
		return fmt.Errorf("virtual media id : %d is opened", id)
	}
	err := arch.fs.RemoveVirtualFile(id)
	if err != nil {
		return err
	}
	return nil
}

func (arch *Archiver) VirtualFileDeleted(fileID uint32, message string) {
	arch.log.Warnv("Media file deleted", "fileID", fileID, "message", message)
	arch.EventsHandler.DeleteFile(fileID)
}
