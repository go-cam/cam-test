package controllers

import (
	"github.com/go-cam/cam"
)

type FileController struct {
	cam.HttpController
}

func (ctrl *FileController) Upload() {
	ctrl.GetFile("file")
	//if file != nil {
	//	dir := camUtils.File.GetRunPath() + "/runtime/test"
	//	if !camUtils.File.Exists(dir) {
	//		camUtils.File.Mkdir(dir)
	//	}
	//	absFilename := dir + "/test.txt"
	//	var fileBytes []byte
	//	_, err := file.Read(fileBytes)
	//	if err != nil {
	//		cam.App.Error("FileController.Save", err.Error())
	//	}
	//	err = camUtils.File.WriteFile(absFilename, fileBytes)
	//	if err != nil {
	//		cam.App.Error("FileController.Save", err.Error())
	//	}
	//}
	//
	//cam.App.Info("FileController.Save", "done")
}
