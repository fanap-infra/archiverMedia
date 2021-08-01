package virtualMedia

import (
	"github.com/fanap-infra/archiverMedia/pkg/media"
	"github.com/fanap-infra/archiverMedia/pkg/vInfo"
	"github.com/fanap-infra/fsEngine/pkg/virtualFile"
	"github.com/fanap-infra/log"
)

func OpenVirtualMedia(fileName string, fileID uint32, blockSize uint32, vFile *virtualFile.VirtualFile, archiver Arch,
	info *vInfo.Info, log *log.Logger) *VirtualMedia {
	return &VirtualMedia{
		vfBuf:      make([]byte, 0),
		frameChunk: &media.PacketChunk{Packets: []*media.Packet{}},
		vFile:      vFile,
		name:       fileName,
		fileID:     fileID,
		blockSize:  blockSize,
		log:        log,
		archiver:   archiver,
		info:       info,
	}
}

func NewVirtualMedia(fileName string, fileID uint32, blockSize uint32, vFile *virtualFile.VirtualFile, archiver Arch,
	log *log.Logger) *VirtualMedia {
	return &VirtualMedia{
		vfBuf:      make([]byte, 0),
		frameChunk: &media.PacketChunk{Packets: []*media.Packet{}, Index: 1, PreviousChunkSize: 0},
		vFile:      vFile,
		name:       fileName,
		fileID:     fileID,
		blockSize:  blockSize,
		log:        log,
		archiver:   archiver,
		info:       &vInfo.Info{},
	}
}
