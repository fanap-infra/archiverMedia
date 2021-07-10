package virtualMedia

import "github.com/fanap-infra/archiverMedia/internal/media"

func (vm *VirtualMedia) WriteFrame(frame *media.Packet) {
	if vm.readOnly {
		return 0, ErrFileIsReadOnly
	}
}

func (vm *VirtualMedia) ReadFrame() *media.Packet {
}

func (vm *VirtualMedia) GotoTime(frameTime int64) (int64, error) {
}

func (vm *VirtualMedia) Close() error {
}
