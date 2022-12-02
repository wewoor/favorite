package api

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetFolderCatalogue(c *gin.Context) {

	pid := c.Param("folderID")
	workspace := os.DirFS(os.Getenv("WORKSPACE_PATH"))
	entities, err := fs.ReadDir(workspace, pid)

	if err != nil {
		log.Print("GetFolderCatalogue: ", err.Error())
		c.JSON(http.StatusOK, ResData{Code: 200, Message: "GetFolderCatalogue failed" + err.Error()})
		return
	}

	folderTree := []FolderTree{}

	for _, entity := range entities {
		folderTree = append(folderTree, FolderTree{
			Name:   entity.Name(),
			Kind:   int(entity.Type()),
			ID:     entity.Name(),
			Parent: pid,
			IsDir:  entity.IsDir(),
		})
	}

	c.JSON(http.StatusOK, folderTree)
}

func GetFile(c *gin.Context) {
	fileID := c.PostForm("fileID")
	path := c.PostForm("path")
	fileTarget := os.DirFS(os.Getenv("WORKSPACE_PATH") + "/" + path)

	entity, err := fs.ReadFile(fileTarget, fileID)

	if err != nil {
		log.Print("Get Filed Failed: ", err)
		c.JSON(http.StatusOK, ResData{Code: 200, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResData{Code: 200, Message: "GetFile success", Data: string(entity)})
}

func WriteFile(c *gin.Context) {
	fileID := c.PostForm("fileID")
	path := c.PostForm("path")
	content := c.PostForm("content")

	err := os.WriteFile(os.Getenv("WORKSPACE_PATH")+"/"+path+"/"+fileID, []byte(content), 0644)

	if err != nil {
		log.Print("Write File Failed: ", err)
		c.JSON(http.StatusOK, ResData{Code: 200, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResData{Code: 200, Message: "WriteFile success"})
}
