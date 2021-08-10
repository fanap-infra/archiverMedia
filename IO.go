package archiverMedia

// It is event handler
func (arch *Archiver) Close() error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()

	//for _, vms := range arch.openFiles {
	//	for _, vm := range vms {
	//		err := vm.CloseWithNotifyArchiver()
	//		if err != nil {
	//			arch.log.Warnv("Can not close virtual media", "err", err.Error())
	//			return err
	//		}
	//	}
	//}
	err := arch.fs.Close()
	if err != nil {
		arch.log.Warnv("Can not close arch", "err", err.Error())
		return err
	}
	return nil
}

// It is called from virtual file
func (arch *Archiver) Closed(fileID uint32) error {
	arch.crudMutex.Lock()
	defer arch.crudMutex.Unlock()
	//arch.log.Infov("virtual media file ")
	//vms, ok := arch.openFiles[fileID]
	//if ok {
	//	if len(vms) == 1 {
	//		delete(arch.openFiles, fileID)
	//	} else {
	//		// ToDo: get index in addition to fileID
	//		arch.openFiles[fileID] = vms[:len(vms)-1]
	//	}
	//}
	return nil
}
