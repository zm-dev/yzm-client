package handler

import (
	"net/http"
	"github.com/zm-dev/yzm-client/src/backend/pkg/httputils"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/zm-dev/yzm-client/src/backend/distinguish"
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"io"
)

const mappingsDir = "./mappings/"

func CreateHTTPAPIHandler(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	{
		api.Handle("/batch_upload", httputils.APPHandler(batchUpload)).Methods("POST")
		api.Handle("/upload", httputils.APPHandler(upload)).Methods("POST")
	}
}

func batchUpload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	batchImageFile, batchImageFileHeader, err := r.FormFile("image")
	if err != nil {
		return httputils.InternalServerError("图片压缩包上传失败！").WithError(err)
	}
	defer batchImageFile.Close()

	if distinguish.NeedAutoCategory(category) {
		category, err = distinguish.ZipAutoCategory(batchImageFile, batchImageFileHeader.Size)
		if err != nil {
			return httputils.BadRequest(err.Error()).WithError(err)
		}
	}

	mappings, err := distinguish.BatchProcess(category, batchImageFile, batchImageFileHeader.Size, distinguish.BatchDistinguish)

	if err != nil {
		return httputils.InternalServerError("图片压缩包处理失败！").WithError(err)
	}
	content := distinguish.LoadMappingLines(mappings).ToSortedString()

	u := uuid.NewV4().String()
	err = ioutil.WriteFile(mappingsDir+u, []byte(content), 0644)
	if err != nil {
		return httputils.InternalServerError("生成" + mappingsFileName + "文件失败！").WithError(err)
	}

	jsonBytes, _ := json.Marshal(struct {
		Category    int    `json:"category"`
		DownloadUrl string `json:"download_url"`
	}{Category: category + 1, DownloadUrl: "download?id=" + u})
	w.Write(jsonBytes)
	return nil
}

func upload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {

	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	imageFile, _, err := r.FormFile("image")
	defer imageFile.Close()

	if distinguish.NeedAutoCategory(category) {
		category, err = distinguish.AutoCategory(imageFile)
		imageFile.Seek(io.SeekStart, io.SeekStart)
		if err != nil {
			return httputils.BadRequest(err.Error()).WithError(err)
		}
	}
	yzmStr, err := distinguish.Process(category, imageFile)
	if err != nil {
		return httputils.InternalServerError("图片识别出错!").WithError(err)
	}

	b, err := json.Marshal(struct {
		Category int    `json:"category"`
		Res      string `json:"res"`
	}{Category: category + 1, Res: yzmStr})

	if err != nil {
		return httputils.InternalServerError("").WithError(err)
	}
	w.Write(b)
	return nil
}
