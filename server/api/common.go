package api

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FolderTree struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Kind   int    `json:"kind"`
	Parent string `json:"parent"`
	IsDir  bool   `json:"isDir"`
}
