package api

import (
	"favorite/db"
	"favorite/ent/project"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProject(c *gin.Context) {
	uid := c.Param("uid")
	client := db.GetDBClient()

	result, err := client.Project.Query().
		Where(project.UUID(uid)).
		Only(c)

	if err != nil {
		log.Print("query project failed: ", err)
	}

	c.JSON(http.StatusOK, result)
}
