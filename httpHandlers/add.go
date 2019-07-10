package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"../storage"
	"../structs"
	"./httpUtils"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var message structs.Message

	err = json.Unmarshal(byteData, &message)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if message.Message == "" || message.Sender == "" {
		httpUtils.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	id := storage.Add(message)

	log.Println("Added message:", message)

	httpUtils.HandleSuccess(&w, structs.ID{ID: id})
}
