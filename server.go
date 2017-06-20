package astrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	srv "github.com/astromechio/astro-go/servlib"
)

type AServer struct {
	HubAddress  string
	ServiceName string
}

func DefaultEnvServer(service string) (*AServer, error) {
	hubAddr, ok := os.LookupEnv(srv.AstroHubAddrEnvKey)
	if !ok {
		return nil, srv.AServerError(srv.HubAddrNoExistErr)
	}

	newServ := AServer{
		HubAddress:  hubAddr,
		ServiceName: service,
	}

	return &newServ, nil
}

func (s *AServer) ListenAndServe(handler http.Handler) {
	next := make(chan bool)

	for true {
		go s.handleReq(handler, next)

		ok := <-next
		if !ok {
			break
		}
	}
}

func (s *AServer) handleReq(handler http.Handler, next chan (bool)) {
	req, err := s.getNextRequest()
	if err != nil {
		timer := time.After(time.Second)
		_ = <-timer
		next <- true
		return
	}

	next <- true

	resp := srv.AResponse{
		StatusCode: 200,
	}

	handler.ServeHTTP(resp, req.UserRequest)

	respJSON, _ := json.Marshal(resp)

	jobRespURL := fmt.Sprintf("http://%s/jobs/%s", s.HubAddress, req.ID)
	answerReq, _ := http.NewRequest("POST", jobRespURL, bytes.NewReader(respJSON))
	answerRes, _ := http.DefaultClient.Do(answerReq)

	log.Printf("Job (id %s) response dispatched with status code %d", req.ID, answerRes.StatusCode)
}

func (s *AServer) getNextRequest() (*srv.ARequest, error) {
	jobReqURL := fmt.Sprintf("http://%s/service/%s/jobs", s.HubAddress, s.ServiceName)
	req, _ := http.NewRequest("GET", jobReqURL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Couldn't get job:", err)
		return nil, err
	}

	log.Println("Job recived with status code:", res.StatusCode)

	if res.StatusCode > 204 {
		return nil, srv.AServerError(srv.NoJobsExistForService)
	}

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))

	var job srv.ARequest
	err = json.Unmarshal(resBody, &job)
	if err != nil {
		log.Fatal("Unable to unmarshal job JSON", err)
	}

	return &job, err
}
