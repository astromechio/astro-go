package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"encoding/json"
	"fmt"
	"io/ioutil"

	srv "github.com/astromechio/astro-go/servlib"
)

func main() {
	hubAddr, ok := os.LookupEnv(srv.AstroHubAddrEnvKey)
	if !ok {
		log.Fatal("Unable to get hub addr")
	}

	body := make(map[string]int)
	body["first"] = 5
	body["second"] = 12
	bodyJSON, _ := json.Marshal(body)

	userReq, _ := http.NewRequest("POST", "/", bytes.NewBuffer(bodyJSON))

	userReqJSON, err := json.Marshal(srv.SerializableReqFromRequest(userReq))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(userReqJSON))

	jobReq, reqErr := http.NewRequest("POST", fmt.Sprintf("http://%s/service/add", hubAddr), bytes.NewBuffer(userReqJSON))
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	jobResp, respErr := http.DefaultClient.Do(jobReq)
	if respErr != nil {
		log.Fatal("Error adding job:", err)
	}

	defer jobResp.Body.Close()
	rawBody, readErr := ioutil.ReadAll(jobResp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(rawBody))
	fmt.Println(jobResp.Status)
}
