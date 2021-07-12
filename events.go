package archiverMedia

type Events interface {
	DeleteFile(fileID uint32)
}
