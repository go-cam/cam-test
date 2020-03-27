package controllers

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camUtils"
)

type FileController struct {
	cam.HttpController
}

func (ctrl *FileController) Upload() {
	uploadFile := ctrl.GetFile("file")
	if uploadFile == nil {
		cam.App.Fatal("FileController.Upload", "no upload file")
		return
	}

	absFilename := camUtils.File.GetRunPath() + "/runtime/upload/tmp.jpg"
	err := uploadFile.Save(absFilename)
	if err != nil {
		cam.App.Fatal("FileController.Upload", err.Error())
		return
	}

	cam.App.Trace("FileController.Upload", absFilename)
}
