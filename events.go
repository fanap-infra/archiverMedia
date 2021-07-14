package archiverMedia

type Events interface {
	DeleteFile(fileID uint32)
}

func (arch *Archiver) VirtualFileDeleted(fileID uint32, message string) {
	arch.log.Warnv("Media file deleted", "fileID", fileID, "message", message)
	arch.EventsHandler.DeleteFile(fileID)
}
