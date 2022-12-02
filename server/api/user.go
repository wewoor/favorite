package api

import (
	"favorite/db"
	"favorite/ent/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	client := db.GetDBClient()

	result, err := client.User.Query().
		Where(user.UUID(uid)).
		Only(c)

	if err != nil {
		log.Print("query user failed: ", err)
	}

	c.JSON(http.StatusOK, result)
}
