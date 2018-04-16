package handler

import (
	"net/http"
	"github.com/3tnet/yzm-client/pkg/httputils"
	"github.com/3tnet/yzm-client/batch"
	"fmt"
	"github.com/gorilla/mux"
)

func CreateHTTPAPIHandler() (http.Handler) {
	r := mux.NewRouter()
	r.Handle("/batch", httputils.APPHandler(batchUpload))
	return r
}

func batchUpload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
	batchImageFile, batchImageFileHeader, err := r.FormFile("batch_image")
	defer batchImageFile.Close()

	if err != nil {
		return httputils.InternalServerError("图片压缩包上传失败！").WithError(err)
	}

	yzmChan, err := batch.Process(0, batchImageFile, batchImageFileHeader.Size, batch.BatchDistinguish)

	if err != nil {
		return httputils.InternalServerError("图片压缩包处理失败！").WithError(err)
	}

	for yzm := range yzmChan {
		fmt.Println(yzm.ImageFilename + ", " + yzm.Yzm)
	}

	return nil
}
