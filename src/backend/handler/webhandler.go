package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"github.com/zm-dev/yzm-client/src/backend/pkg/httputils"
	"os"
)

func CreateWebHandler(r *mux.Router) {
	r.Handle("/download", httputils.APPHandler(download)).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
}

const mappingsFileName = "mappings.txt"

func download(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
	id := r.FormValue("id")
	if id == "" {
		return httputils.NotFound("")
	}
	filename := mappingsDir + id
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 文件不存在
		return httputils.NotFound("文件不存在")
	} else if err != nil {
		return httputils.InternalServerError("").WithError(err)
	}
	w.Header().Add("Content-Disposition", "attachment; filename="+url.QueryEscape(mappingsFileName))
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Expires", "0")
	w.Header().Add("Cache-Control", "must-revalidate")
	w.Header().Add("Pragma", "public")
	http.ServeFile(w, r, filename)
	return nil
}
