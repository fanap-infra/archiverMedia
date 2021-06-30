package ArchiverEngine

import (
	"ArchiverEngine/internal/virtualMedia"
	"fmt"
)

// Create new virtual file and add opened files
func (arch *Archiver) NewVirtualFile(id uint32, fileName string) (*virtualMedia.VirtualMedia, error) {
	fse.crudMutex.Lock()
	defer fse.crudMutex.Unlock()
	if fse.header.CheckIDExist(id) {
		return nil, fmt.Errorf("this ID: %v, had been taken", id)
	}
	blm := blockAllocationMap.New(fse.log, fse, fse.maxNumberOfBlocks)

	vf := virtualFile.NewVirtualFile(fileName, id, fse.blockSize-BlockHeaderSize, fse, blm,
		int(fse.blockSize-BlockHeaderSize)*VirtualFileBufferBlockNumber, fse.log)
	err := fse.header.AddVirtualFile(id, fileName)
	if err != nil {
		return nil, err
	}
	fse.openFiles[id] = vf
	return vf, nil
}

func (arch *Archiver) OpenVirtualFile(id uint32) (*virtualMedia.VirtualMedia, error) {
	fse.crudMutex.Lock()
	defer fse.crudMutex.Unlock()
	_, ok := fse.openFiles[id]
	if ok {
		return nil, fmt.Errorf("this ID: %v is opened before", id)
	}
	fileInfo, err := fse.header.GetFileData(id)
	if err != nil {
		return nil, err
	}
	blm, err := blockAllocationMap.Open(fse.log, fse, fse.maxNumberOfBlocks, fileInfo.GetLastBlock(),
		fileInfo.GetRMapBlocks())
	if err != nil {
		return nil, err
	}
	vf := virtualFile.OpenVirtualFile(fileInfo, fse.blockSize-BlockHeaderSize, fse, blm,
		int(fse.blockSize-BlockHeaderSize)*VirtualFileBufferBlockNumber, fse.log)
	//err = fse.header.AddVirtualFile(id, fileInfo.GetName())
	//if err != nil {
	//	return nil, err
	//}
	return vf, nil
}

func (arch *Archiver) RemoveVirtualFile(id uint32) error {
	fse.crudMutex.Lock()
	defer fse.crudMutex.Unlock()
	_, ok := fse.openFiles[id]
	if ok {
		return fmt.Errorf("virtual file id : %d is opened", id)
	}
	return nil
}



