package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../storage"
	"../structs"
	"./httpUtils"
)

func List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, _ := ioutil.ReadAll(r.Body)

	var message structs.TanRequest

	json.Unmarshal(byteData, &message)

	result := storage.Get(message.Input)
	httpUtils.HandleSuccess(&w, structs.TanResult{Result: result})
}
