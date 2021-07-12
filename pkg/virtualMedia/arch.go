package virtualMedia

type Arch interface {
	Closed(fileID uint32) error
}
