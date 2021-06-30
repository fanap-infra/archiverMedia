package virtualMedia

import "ArchiverEngine/internal/media"

func (vm *VirtualMedia) WriteFrame(frame *media.Packet)  {
}

func (vm *VirtualMedia) ReadFrame() *media.Packet {
}

func (vm *VirtualMedia) GotoTime(frameTime int64) (int64, error) {
}

func (vm *VirtualMedia) Close() error {

}
