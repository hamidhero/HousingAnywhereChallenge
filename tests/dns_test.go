package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"housingAnywherChallenge/core/requests"
	"housingAnywherChallenge/core/resources"
	"housingAnywherChallenge/libs"
	"housingAnywherChallenge/routers"
	"net/http/httptest"
	"testing"
)

func TestNewOutput(t *testing.T)  {
	output := libs.NewOutput()
	assert.Equal(t, 200, output.Status)
}

func TestCalculateLocation(t *testing.T)  {
	x := "23.4"
	y := "24.1"
	z := "12.4"
	vel := "33.0"

	loc, _ := libs.CalculateLocation(x, y, z, vel)
	assert.Equal(t, 92.9, *loc)
}

func TestGetDataBankLocation(t *testing.T)  {
	route := routers.GetRouter()

	request := requests.GetDataBankLocationRequest{
		X:   "23.4",
		Y:   "13.4",
		Z:   "14.1",
		Vel: "20.0",
	}
	requestByte, _ := json.Marshal(request)


	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/location", bytes.NewReader(requestByte))
	route.ServeHTTP(w, req)

	defer req.Body.Close()

	var output resources.GetDataBankLocationResource
	err := json.Unmarshal(w.Body.Bytes(), &output)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 70.9, output.Loc)

}