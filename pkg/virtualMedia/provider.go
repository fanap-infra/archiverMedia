package virtualMedia

import (
	"github.com/fanap-infra/fsEngine/pkg/virtualFile"
	"github.com/fanap-infra/log"
)

func OpenVirtualMedia(fileName string, fileID uint32, blockSize uint32, vFile *virtualFile.VirtualFile, archiver Arch,
	log *log.Logger) *VirtualMedia {
	return &VirtualMedia{
		vfBuf:     make([]byte, 0),
		vFile:     vFile,
		name:      fileName,
		fileID:    fileID,
		blockSize: blockSize,
		log:       log,
		archiver:  archiver,
	}
}

func NewVirtualMedia(fileName string, fileID uint32, blockSize uint32, vFile *virtualFile.VirtualFile, archiver Arch,
	log *log.Logger) *VirtualMedia {
	return &VirtualMedia{
		vfBuf:     make([]byte, 0),
		vFile:     vFile,
		name:      fileName,
		fileID:    fileID,
		blockSize: blockSize,
		log:       log,
		archiver:  archiver,
	}
}
