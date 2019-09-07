package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type roverRequest struct {
	Plateau    Plateau    `json:"plateau"`
	Direction  string     `json:"direction"`
	Coordinate Coordinate `json:"coordinates"`
}

type instructionsRequest struct {
	Instructions string `json:"instruction"`
}

var rover Rover

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/rover", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendErrorResponse(c, err)
		}
		var roverRequest roverRequest
		err = json.Unmarshal(body, &roverRequest)
		if err != nil {
			sendErrorResponse(c, err)
		}
		rover.Coordinate = roverRequest.Coordinate
		rover.Plateau = roverRequest.Plateau
		rover.Direction = getDirection(roverRequest.Direction)
		sendSuccessResponse(c, "Rover successfully initialised")
	})

	r.POST("/rover/instructions", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendErrorResponse(c, err)
		}
		var instructions instructionsRequest
		err = json.Unmarshal(body, &instructions)
		if err != nil {
			sendErrorResponse(c, err)
		}
		position, errorMsg := rover.Run(instructions.Instructions)
		if errorMsg != nil {
			sendErrorResponse(c, errorMsg)
		} else {
			sendSuccessResponse(c, position)
		}
	})

	r.GET("/rover/left", func(c *gin.Context) {
		unitInstructionHandler(c, "L")
	})

	r.GET("/rover/right", func(c *gin.Context) {
		unitInstructionHandler(c, "R")
	})

	r.GET("/rover/move", func(c *gin.Context) {
		unitInstructionHandler(c, "M")
	})

	r.GET("/", func(c *gin.Context) {
		sendSuccessResponse(c, rover.toString())
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func unitInstructionHandler(c *gin.Context, instruction string) {

	position, err := rover.Run("R")
	if err != nil {
		sendErrorResponse(c, err)
	} else {
		sendSuccessResponse(c, position)
	}
}

func sendSuccessResponse(c *gin.Context, result string) {
	c.JSON(200, gin.H{
		"roverposition": result,
	})
}

func sendErrorResponse(c *gin.Context, errorMsg error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": errorMsg.Error(),
	})
}

func getDirection(s string) Direction {
	switch s {
	case "N":
		return Enum.N
	case "S":
		return Enum.S
	case "E":
		return Enum.E
	case "W":
		return Enum.W
	}
	return nil
}
