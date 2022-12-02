package api

import (
	"favorite/service"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

func GetContainers(c *gin.Context) {

	containers, err := service.GetContainers()

	if err != nil {
		log.Print("GetContainers failed: ", err)
	}

	for _, containers := range containers {
		log.Print(containers.Names)
	}

	c.JSON(http.StatusOK, containers)
}

func ContainerStart(c *gin.Context) {

	containerID := c.Param("containerID")

	err := service.ContainerStart(containerID, types.ContainerStartOptions{})

	res := ResData{
		Code:    200,
		Message: "Container start success",
	}

	if err != nil {
		log.Print("Container start failed: ", err)
		res.Code = 500
		res.Message = "Container start failed" + err.Error()
	}

	c.JSON(http.StatusOK, res)
}

func ContainerStop(c *gin.Context) {

	containerID := c.Param("containerID")

	err := service.ContainerStop(containerID, nil)

	res := ResData{
		Code:    200,
		Message: "Container stop success",
	}

	if err != nil {
		log.Print("Container stop failed: ", err)
		res.Code = 500
		res.Message = "Container stop failed" + err.Error()
	}

	c.JSON(http.StatusOK, res)
}
