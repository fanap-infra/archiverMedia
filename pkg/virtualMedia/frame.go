package virtualMedia

import (
	"bytes"
	"encoding/binary"
	"fmt"

	errPkg "github.com/fanap-infra/archiverMedia/pkg/err"
	"github.com/fanap-infra/archiverMedia/pkg/media"
	"github.com/fanap-infra/fsEngine/pkg/virtualFile"
	"google.golang.org/protobuf/proto"
)

const EndOfFile = errPkg.Error("no more frames")

func generateFrameChunk(med *media.PacketChunk) ([]byte, error) {
	b, err := proto.Marshal(med)
	if err != nil {
		return nil, err
	}
	binSize := make([]byte, 4)
	binary.BigEndian.PutUint32(binSize, uint32(len(b)))
	b = append(binSize, b...)
	b = append([]byte(FrameChunkIdentifier), b...)
	return b, nil
}

func (vm *VirtualMedia) NextFrameChunk() (*media.PacketChunk, error) {
	//ToDo: do not expose this functions , it is not safe thread
	frameChunkDataSize := uint32(0)
	nextFrameChunk := -1
	errorCounter := 0
	for {
		tmpBuf := make([]byte, vm.blockSize)
		if frameChunkDataSize != 0 && nextFrameChunk != -1 {
			if uint32(len(vm.vfBuf[nextFrameChunk+FrameChunkHeader:])) >= frameChunkDataSize {
				fc := &media.PacketChunk{}
				vm.frameChunkRXSize = FrameChunkHeader + frameChunkDataSize
				err := proto.Unmarshal(vm.vfBuf[nextFrameChunk+FrameChunkHeader:nextFrameChunk+FrameChunkHeader+int(frameChunkDataSize)], fc)
				if err != nil {
					vm.log.Errorv("FrameChunk proto.Unmarshal", "err", err.Error(),
						"nextFrameChunk", nextFrameChunk, "frameChunkDataSize", frameChunkDataSize,
						"len(vm.vfBuf)", len(vm.vfBuf), "BLMArray", vm.vFile.GetBLMArray())
					vm.vfBuf = vm.vfBuf[nextFrameChunk+int(frameChunkDataSize)+FrameChunkIdentifierSize:]
					errorCounter++
					if errorCounter > 1 {
						vm.log.Errorv("can not parse proto file multiple times",
							"nextFrameChunk", nextFrameChunk, "frameChunkDataSize", frameChunkDataSize,
							"len(vm.vfBuf)", len(vm.vfBuf), "BLMArray", vm.vFile.GetBLMArray())
						return nil, err
					}
				} else {
					vm.frameChunkRX = fc
					vm.currentFrameInChunk = 0
					vm.vfBuf = vm.vfBuf[nextFrameChunk+FrameChunkIdentifierSize+int(frameChunkDataSize):]
					return fc, nil
				}
			}
			nextFrameChunk = -1
		}

		if nextFrameChunk == -1 {
			nextFrameChunk = bytes.Index(vm.vfBuf, []byte(FrameChunkIdentifier))
		}

		if nextFrameChunk != -1 && frameChunkDataSize == 0 && len(vm.vfBuf) >= nextFrameChunk+FrameChunkHeader {
			frameChunkDataSize = binary.BigEndian.Uint32(
				vm.vfBuf[nextFrameChunk+FrameChunkIdentifierSize : nextFrameChunk+FrameChunkHeader])
			continue
		}

		n, err := vm.vFile.Read(tmpBuf)
		if n == 0 {
			if err == virtualFile.EndOfFile {
				return nil, EndOfFile
			}
			vm.log.Warnv("can not read data", "id", vm.fileID,
				"nextFrameChunk", nextFrameChunk, "frameChunkDataSize", frameChunkDataSize,
				"len(vm.vfBuf)", len(vm.vfBuf), "err", err.Error())
			return nil, err
		}
		vm.vfBuf = append(vm.vfBuf, tmpBuf[:n]...)
	}
}

func (vm *VirtualMedia) PreviousFrameChunk() (*media.PacketChunk, error) {
	//ToDo: do not expose this functions , it is not safe thread
	if vm.frameChunkRX != nil {
		if vm.frameChunkRX.Index == 1 || vm.frameChunkRX.PreviousChunkSize == 0 {
			return nil, fmt.Errorf("there is no previous frame chunk")
		}
	} else {
		return vm.NextFrameChunk()
	}
	currentFrameChunkIndex := vm.frameChunkRX.Index
	//seekPointer := vm.vFile.GetSeek() - int(vm.frameChunkRX.PreviousChunkSize) - int(vm.frameChunkRXSize)
	//if seekPointer < 0 {
	//	seekPointer = 0
	//}
	tmpBuf := make([]byte, vm.frameChunkRX.PreviousChunkSize)
	vm.vfBuf = vm.vfBuf[:0]
	// counter := 0
	n, err := vm.vFile.ReadAt(tmpBuf, int64(vm.frameChunkRX.PreviousChunkStartAddress))
	if n == 0 {
		return nil, err
	}
	vm.vfBuf = append(tmpBuf[:n], vm.vfBuf...)
	fc, err := vm.NextFrameChunk()
	if err != nil {
		vm.log.Errorv("can not get previous frame chunk", "n", n, "len(tmpBuf)", len(tmpBuf),
			"currentFrameChunkIndex", currentFrameChunkIndex,
			"err", err.Error())
		return nil, err
	}
	return fc, nil

	//for {
	//	counter++
	//	if counter > 5 {
	//		vm.log.Errorv("break PreviousFrameChunk loop")
	//		return vm.NextFrameChunk()
	//	}
	//	n, err := vm.vFile.ReadAt(tmpBuf, int64(seekPointer))
	//
	//	if n == 0 {
	//		return nil, err
	//	}
	//	vm.vfBuf = append(tmpBuf[:n], vm.vfBuf...)
	//fc, err := vm.NextFrameChunk()
	//if err != nil {
	//	return nil, err
	//}
	//	if fc.Index+1 == currentFrameChunkIndex {
	//		return fc, nil
	//	} else if fc.Index > currentFrameChunkIndex {
	//		seekPointer = seekPointer - int(vm.blockSize*2)
	//		if seekPointer < 0 {
	//			seekPointer = 0
	//		}
	//	} else {
	//		return vm.NextFrameChunk()
	//	}
	//}
}
