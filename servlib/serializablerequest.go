package servlib

import "net/http"
import "io/ioutil"

type SerializableRequest map[string]interface{}

func SerializableReqFromRequest(r *http.Request) (SerializableRequest, error) {
	newReq := make(map[string]interface{})
	newReq["Method"] = r.Method
	newReq["URL"] = r.URL
	newReq["Proto"] = r.Proto
	newReq["ProtoMajor"] = r.ProtoMajor
	newReq["ProtoMinor"] = r.ProtoMinor
	newReq["Header"] = r.Header
	newReq["ContentLength"] = r.ContentLength
	newReq["TransferEncoding"] = r.TransferEncoding
	newReq["Close"] = r.Close
	newReq["Host"] = r.Host
	newReq["Form"] = r.Form
	newReq["PostForm"] = r.PostForm
	newReq["MultipartForm"] = r.MultipartForm
	newReq["Trailer"] = r.Trailer
	newReq["RemoteAddr"] = r.RemoteAddr
	newReq["RequestURI"] = r.RequestURI
	newReq["TLS"] = r.TLS
	newReq["Response"] = r.Response

	origBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	newReq["Body"] = origBody

	return newReq, nil
}
