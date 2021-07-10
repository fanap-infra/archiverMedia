package archiverMedia

type Events interface {
	DeleteFile(fileID uint32)
	DeleteFileByArchiver(archiverFile string)
}

