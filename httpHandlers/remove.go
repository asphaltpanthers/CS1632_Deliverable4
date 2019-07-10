package httpHandlers

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"../storage"
	"../structs"
	"./httpUtils"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var id structs.ID

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		httpUtils.HandleError(&w, 400, "Bad Request", "Error unmarshalling", err)
		return
	}

	if id.ID == 0 {
		httpUtils.HandleError(&w, 500, "Bad Request", "ID not provided", nil)
		return
	}

	if storage.Remove(id.ID) {
		httpUtils.HandleSuccess(&w, structs.ID{ID: id.ID})
	} else {
		httpUtils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}
