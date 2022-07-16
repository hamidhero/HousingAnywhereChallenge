package libs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const SectorID float64 = 1

type Error struct {
	ErrorCode int
	ErrorMsg  string
}

type Output struct {
	Timestamp time.Time `json:"TimeStamp"`
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Error     []Error   `json:"error"`
}

//NewOutput produces a basic response
func NewOutput() Output {
	output := Output{}
	output.Timestamp = time.Now()
	output.Status = http.StatusOK
	output.Message = "operation successful!"
	return output
}

//SetError returns appropriate error output
func SetError(e interface{}, c *gin.Context, output Output, internalErrorCode int, externalErrorCode int) {
	var err Error
	var errorMsg string
	switch e.(type) {
	case error:
		errorMsg = e.(error).Error()
	case string:
		errorMsg = e.(string)
	}

	err.ErrorCode = internalErrorCode
	err.ErrorMsg = errorMsg
	output.Error = append(output.Error, err)

	output.Status = externalErrorCode
	output.Message = errorMsg
	c.JSON(output.Status, output)
	return
}

//CalculateLocation calculates location of databank according to its coords
func CalculateLocation(x, y, z, vel string) (*float64, error) {

	xFloat, e := strconv.ParseFloat(x, 64)
	if e != nil {
		return nil, e
	}
	yFloat, e := strconv.ParseFloat(y, 64)
	if e != nil {
		return nil, e
	}
	zFloat, e := strconv.ParseFloat(z, 64)
	if e != nil {
		return nil, e
	}
	velFloat, e := strconv.ParseFloat(vel, 64)
	if e != nil {
		return nil, e
	}

	loc := SectorID *( xFloat+yFloat+zFloat) + velFloat
	return &loc, nil
}
