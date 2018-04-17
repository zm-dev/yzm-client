package handler

import (
	"net/http"
	"github.com/zm-dev/yzm-client/pkg/httputils"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/zm-dev/yzm-client/distinguish"
)

func CreateHTTPAPIHandler() (http.Handler) {
	r := mux.NewRouter()
	r.Handle("/batch_upload", httputils.APPHandler(batchUpload))
	r.Handle("/upload", httputils.APPHandler(upload))
	return r
}

func batchUpload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	batchImageFile, batchImageFileHeader, err := r.FormFile("batch_image")
	defer batchImageFile.Close()

	if err != nil {
		return httputils.InternalServerError("图片压缩包上传失败！").WithError(err)
	}

	yzmChan, err := distinguish.BatchProcess(category, batchImageFile, batchImageFileHeader.Size, distinguish.BatchDistinguish)

	if err != nil {
		return httputils.InternalServerError("图片压缩包处理失败！").WithError(err)
	}

	for yzm := range yzmChan {
		fmt.Println(yzm.ImageFilename + ", " + yzm.Yzm)
	}

	return nil
}

func upload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {

	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	imageFile, imageFileHeader, err := r.FormFile("image")
	defer imageFile.Close()

	label, err := distinguish.Process(category, imageFileHeader.Filename, imageFile)
	if err != nil {
		return httputils.InternalServerError("图片识别出错!").WithError(err)
	}
	fmt.Println(label.ImageFilename, ",", label.Yzm)
	return nil

}
