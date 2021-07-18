package archiverMedia

// It is event handler
func (arch *Archiver) Close() error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()

	for _, vm := range arch.openFiles {
		err := vm.CloseWithNotifyArchiver()
		if err != nil {
			arch.log.Warnv("Can not close virtual media", "err", err.Error())
			return err
		}
	}
	err := arch.fs.Close()
	if err != nil {
		arch.log.Warnv("Can not close arch", "err", err.Error())
		return err
	}
	return nil
}

func (arch *Archiver) Closed(fileID uint32) error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()

	delete(arch.openFiles, fileID)

	return nil
}
