package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"housingAnywherChallenge/core/requests"
	"housingAnywherChallenge/core/resources"
	"housingAnywherChallenge/libs"
	"net/http"
)

//GetDataBankLocation is the endpoint that accepts json as input and provides loc as output in json format
func GetDataBankLocation(c *gin.Context)  {
	output := libs.NewOutput()

	//bind input to the suitable datatype
	var request requests.GetDataBankLocationRequest
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		libs.SetError(err, c, output, http.StatusInternalServerError, http.StatusBadRequest)
		return
	}

	//call CalculateLocation to get the location or error if any
	location, e := libs.CalculateLocation(request.X, request.Y, request.Z, request.Vel)
	if e != nil {
		libs.SetError(e, c, output, http.StatusInternalServerError, http.StatusBadRequest)
		return
	}

	//put location in a desired datatype
	response := resources.GetDataBankLocationResource{Loc:*location}

	//provide json output
	c.JSON(http.StatusOK, response)
	return
}
