package httpHandlers

import (
	"net/http"

	"../storage"
	"./httpUtils"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}
