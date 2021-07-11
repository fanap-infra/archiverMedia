package virtualMedia

import (
	"bytes"
	"encoding/binary"

	"github.com/fanap-infra/archiverMedia/internal/media"
	"google.golang.org/protobuf/proto"
)

func generateFrameChunk(med *media.PacketChunk) ([]byte, error) {
	b, err := proto.Marshal(med)
	if err != nil {
		return nil, err
	}
	binSize := make([]byte, 4)
	binary.BigEndian.PutUint32(binSize, uint32(len(b)+4))
	b = append(binSize, b...)
	b = append([]byte(FrameChunkIdentifier), b...)
	return b, nil
}

func (vm *VirtualMedia) FrameChunk() (*media.PacketChunk, error) {
	tmpBuf := make([]byte, vm.blockSize)
	frameChunkDataSize := uint32(0)
	nextFrameChunk := -1
	for {
		if frameChunkDataSize != 0 && nextFrameChunk != -1 {
			if uint32(len(vm.vfBuf[nextFrameChunk+FrameChunkHeader:])) >= frameChunkDataSize {
				fc := &media.PacketChunk{}
				err := proto.Unmarshal(vm.vfBuf[nextFrameChunk+FrameChunkHeader:nextFrameChunk+int(frameChunkDataSize)+FrameChunkIdentifierSize], fc)
				if err != nil {
					vm.vfBuf = vm.vfBuf[nextFrameChunk+int(frameChunkDataSize)+FrameChunkIdentifierSize:]
					vm.log.Errorv("FrameChunk proto.Unmarshal", "err", err.Error(),
						"nextFrameChunk", nextFrameChunk, "frameChunkDataSize", frameChunkDataSize)
					return nil, err
				}
				vm.frameChunkRX = fc
				vm.vfBuf = vm.vfBuf[nextFrameChunk+int(frameChunkDataSize)+FrameChunkIdentifierSize:]
				return fc, nil
			}
		}

		if nextFrameChunk == -1 {
			nextFrameChunk = bytes.Index(vm.vfBuf, []byte(FrameChunkIdentifier))
			if nextFrameChunk != -1 {
				frameChunkDataSize = binary.BigEndian.Uint32(
					vm.vfBuf[nextFrameChunk+FrameChunkIdentifierSize : nextFrameChunk+FrameChunkHeader])
				continue
			}
		}

		n, err := vm.vFile.Read(tmpBuf)
		if n == 0 {
			return nil, err
		}
		vm.vfBuf = append(vm.vfBuf, tmpBuf[:n]...)
	}
}
