package archiverMedia


//type EventHandlerMock struct {
//}
//
//func (eventHandler *EventHandlerMock) DeleteFile(fileID uint32) {
//	log.Infov("delete file", "fileID", fileID)
//}
//
//func Test_CheckRecordFile(t *testing.T) {
//	ArchiverProvider := NewProvider()
//	var evHandler EventHandlerMock
//	recFile := "/media/mohammad/ccd5ecc0-957f-469c-8d52-f3ed9339773d/Archivers/Archiver#62"
//	arch, err := ArchiverProvider.ParseFileSystem(recFile, &evHandler, log.GetScope("Archiver"))
//	assert.Equal(t, nil, err)
//	files := arch.fs.GetFileList()
//	//arch.fs.
//	var blocks []uint32
//	for i, file := range  files {
//		info, err := vInfo.Parse(file.GetOptional())
//		if err != nil {
//			log.Errorv("can not parse optional data", "id", file.Id, "err", err.Error())
//			continue
//		}
//		log.Infov("files list", "i",i,"ID", file.Id, "StartTime", info.StartTime, "End Time", info.EndTime)
//		blm, err := arch.fs.GetFileBLM(file.GetId())
//		assert.Equal(t, nil, err)
//		log.Infov("files blocks","ID", file.Id, "blocks", blm.ToArray())
//		assert.Equal(t, nil, err)
//		if hasCommonValue(blocks, blm.ToArray()) {
//			log.Error("There is common value")
//		}
//		blocks = append(blocks, blm.ToArray()...)
//
//		vm, err :=arch.OpenVirtualMediaFile(file.GetId())
//		assert.Equal(t, nil, err)
//		fc, err := vm.NextFrameChunk()
//		assert.Equal(t, nil, err)
//		endTime := fc.EndTime
//		ID := fc.Index
//		for {
//			fc, err := vm.NextFrameChunk()
//			if err == virtualMedia.EndOfFile {
//				log.Infov("end of file", "ID FC", ID)
//				break
//			}
//			assert.Equal(t, nil, err)
//			assert.Equal(t, endTime, fc.StartTime)
//			endTime = fc.EndTime
//			ID = fc.Index
//		}
//	}
//	err = arch.Close()
//	assert.Equal(t, nil, err)
//	assert.Equal(t, true, hasCommonValue([]uint32{1,2,3}, []uint32{5,2,6}))
//}
//
//func Test_CheckFile(t *testing.T) {
//	ArchiverProvider := NewProvider()
//	var evHandler EventHandlerMock
//	recFile := "/media/mohammad/ccd5ecc0-957f-469c-8d52-f3ed9339773d/Archivers/Archiver#62"
//	arch, err := ArchiverProvider.ParseFileSystem(recFile, &evHandler, log.GetScope("Archiver"))
//	assert.Equal(t, nil, err)
//
//	vm, err :=arch.OpenVirtualMediaFile(1399)
//	assert.Equal(t, nil, err)
//	fc, err := vm.NextFrameChunk()
//	assert.Equal(t, nil, err)
//	endTime := fc.EndTime
//	ID := fc.Index
//	for {
//		fc, err := vm.NextFrameChunk()
//		if err == virtualMedia.EndOfFile {
//			log.Infov("end of file", "ID FC", ID)
//			break
//		}
//		assert.Equal(t, nil, err)
//		assert.Equal(t, endTime, fc.StartTime)
//		endTime = fc.EndTime
//		ID = fc.Index
//	}
//
//
//	err = arch.Close()
//	assert.Equal(t, nil, err)
//}
//
//
//func hasCommonValue(first, second []uint32) bool {
//	//if len(first) != len(second) {
//	//	return false
//	//}
//	exists := make(map[uint32]bool)
//	for _, value := range first {
//		exists[value] = true
//	}
//	for _, value := range second {
//		if exists[value] {
//			return true
//		}
//	}
//	return false
//}
//
