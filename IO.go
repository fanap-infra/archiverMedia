package archiverMedia

// It is event handler
func (arch *Archiver) Close() error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()
	err := arch.fs.Close()
	if err != nil {
		arch.log.Warnv("Can not close arch", "err", err.Error())
		return err
	}

	for _, vm := range arch.openFiles {
		err := vm.Close()
		if err != nil {
			arch.log.Warnv("Can not close virtual media", "err", err.Error())
			return err
		}
	}

	return nil
}
