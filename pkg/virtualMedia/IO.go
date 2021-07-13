package virtualMedia

import "github.com/fanap-infra/archiverMedia/pkg/media"

func (vm *VirtualMedia) WriteFrame(frame *media.Packet) error {
	if vm.readOnly {
		return ErrFileIsReadOnly
	}
	vm.fwMUX.Lock()
	defer vm.fwMUX.Unlock()
	if vm.frameChunk.Packets == nil {
		vm.log.Warnv("frame chunk is nil", "name", vm.name)
		vm.frameChunk.Packets = []*media.Packet{}
	}

	if frame.PacketType == media.PacketType_PacketVideo && frame.IsKeyFrame {
		if len(vm.frameChunk.Packets) >= FrameChunkMinimumFrameCount {
			b, err := generateFrameChunk(vm.frameChunk)
			if err != nil {
				return err
			}
			l, err := vm.vFile.Write(b)
			vm.fileSize += uint32(l)
			if err != nil {
				return err
			}
			vm.log.Infov("packet chunk is written", "Index", vm.frameChunk.Index,
				"packets number", len(vm.frameChunk.Packets), "size frame chunk", len(b))
			vm.frameChunk = &media.PacketChunk{Index: vm.frameChunk.Index + 1}
		}
	}

	vm.frameChunk.Packets = append(vm.frameChunk.Packets, frame)
	if frame.PacketType == media.PacketType_PacketVideo {
		if vm.frameChunk.StartTime == 0 {
			vm.frameChunk.StartTime = frame.Time
		}
		vm.frameChunk.EndTime = frame.Time
	}
	return nil
}

func (vm *VirtualMedia) ReadFrame() (*media.Packet, error) {
	vm.rxMUX.Lock()
	defer vm.rxMUX.Unlock()
	if vm.frameChunk == nil || int(vm.currentFrameInChunk) >= len(vm.frameChunk.Packets) {
		fc, err := vm.NextFrameChunk()
		if err != nil {
			return nil, err
		}
		vm.frameChunk = fc
		vm.currentFrameInChunk = 0
	} else if uint32(len(vm.frameChunk.Packets)) <= (vm.currentFrameInChunk) {
		fc, err := vm.NextFrameChunk()
		if err != nil {
			return nil, err
		}
		vm.frameChunk = fc
		vm.currentFrameInChunk = 0
	}
	vm.currentFrameInChunk++
	return vm.frameChunk.Packets[vm.currentFrameInChunk-1], nil
}

func (vm *VirtualMedia) GotoTime(frameTime int64) (int64, error) {
	return 0, nil // vm.vFile.
}

func (vm *VirtualMedia) Close() error {
	vm.fwMUX.Lock()
	defer vm.fwMUX.Unlock()
	vm.rxMUX.Lock()
	defer vm.rxMUX.Unlock()

	if len(vm.frameChunk.Packets) > 0 {
		b, err := generateFrameChunk(vm.frameChunk)
		if err != nil {
			return err
		}
		l, err := vm.vFile.Write(b)
		vm.fileSize += uint32(l)
		if err != nil {
			return err
		}
		vm.log.Infov("packet chunk is written in close", "Index", vm.frameChunk.Index,
			"packets number", len(vm.frameChunk.Packets), "size frame chunk", len(b))
		vm.frameChunk = &media.PacketChunk{Index: vm.frameChunk.Index + 1}
	}

	err := vm.vFile.Close()
	if err != nil {
		vm.log.Errorv("virtual media can not close", "err", err.Error())
	}
	vm.vfBuf = vm.vfBuf[:0]
	vm.frameChunkRX = nil
	vm.frameChunk = nil
	return vm.archiver.Closed(vm.fileID)
}
