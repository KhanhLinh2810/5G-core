package smf_service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(csrJSON []byte) (*http.Response, error) {
	smfURL := "http://localhost:8081/nsmf-pdusession/v1/sm-contexts"
	resp, err := http.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
	if err != nil {
		return nil, err
	}
	return resp, nil
}