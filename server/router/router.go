package router

import (
	"favorite/api"

	"github.com/gin-gonic/gin"
)

// func InitWebsocket(r *gin.Engine) {
// 	server := socketIo.NewServer(nil)

// 	server.OnConnect("/", func(s socketIo.Conn) error {
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID())
// 		return nil
// 	})

// 	server.OnError("/", func(s socketIo.Conn, e error) {
// 		log.Println("connection error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketIo.Conn, reason string) {
// 		log.Println("connection closed", reason)
// 	})

// 	go func() {
// 		if err := server.Serve(); err != nil {
// 			log.Fatalf("socket.io listen error %s\n", err)
// 		}
// 	}()
// 	defer server.Close()

// 	r.GET("/socket.io/*any", gin.WrapH(server))
// 	r.POST("/socket.io/*any", gin.WrapH(server))
// }

func Init(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World, Favorite!")
	})
	r.GET("/user/:uid", api.GetUser)
	r.GET("/project/:uid", api.GetProject)
	r.GET("/containers", api.GetContainers)
	r.GET("/containers/start/:containerID", api.ContainerStart)
	r.GET("/containers/stop/:containerID", api.ContainerStop)
	r.GET("/catalogue/:folderID", api.GetFolderCatalogue)
	r.POST("/catalogue/getFile", api.GetFile)
	r.POST("/catalogue/writeFile", api.WriteFile)
}
