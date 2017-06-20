package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		reqJSON, _ := ioutil.ReadAll(r.Body)

		var problem map[string]int
		err := json.Unmarshal(reqJSON, &problem)
		if err != nil {
			w.Write([]byte("Fuck you json.Unmarshal"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := make(map[string]int)
		resp["answer"] = problem["first"] + problem["second"]

		respJSON, _ := json.Marshal(resp)

		w.Write(respJSON)
	}
}
