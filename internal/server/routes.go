package server

import (
	"LaTeXCompilationService/internal/archive"
	"LaTeXCompilationService/internal/compiler"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	UploadDir       = "upload"
	DefaultFileName = "main.tex"
)

func CompilationHandler(ctx *gin.Context) {
	mainFileName := ctx.DefaultPostForm("filename", DefaultFileName)
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Printf("error while getting file: %v", err)
		ctx.String(http.StatusBadRequest, "could not retrieve the file from the request")
		return
	}

	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		log.Printf("error while creating upload directory: %v", err)
		ctx.String(http.StatusInternalServerError, "an error occurred while creating the upload dir")

	}

	tempDir, err := ioutil.TempDir(UploadDir, "upload")
	if err != nil {
		log.Printf("error while creating temp directory: %v", err)
		ctx.String(http.StatusInternalServerError, "an error occurred while creating the temp dir")
		return
	}
	defer os.RemoveAll(tempDir)

	destZipFile, err := ioutil.TempFile(tempDir, "*.zip")
	if err != nil {
		log.Printf("error while creating temp file: %v", err)
		ctx.String(http.StatusInternalServerError, "an error occurred while creating the temp file")
		return
	}

	err = ctx.SaveUploadedFile(file, destZipFile.Name())
	if err != nil {
		log.Printf("could not save uploaded file: %v", err)
		ctx.String(http.StatusBadRequest, "an error occurred while saving the file")
		return
	}

	dirName, err := archive.Unzip(destZipFile.Name(), tempDir)
	if err != nil {
		log.Printf("error while unzipping: %v", err)
		ctx.String(http.StatusBadRequest, "an error occurred while unzipping the file")
		return
	}

	fileName, err := compiler.CompileWithToC(filepath.Join(tempDir, dirName), mainFileName)
	if err != nil {
		log.Printf("error while compiling: %v", err)
		ctx.String(http.StatusBadRequest, "an error occurred during the compilation")
		return
	}

	pdf, err := ioutil.ReadFile(filepath.Join(tempDir, dirName, fileName))
	if err != nil {
		log.Printf("error while sending file: %v", err)
		ctx.String(http.StatusInternalServerError, "error occurred while sending back the file")
		return
	}

	ctx.Data(http.StatusOK, "application/pdf", pdf)
}
