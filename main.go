package main

import (
	"fmt"
	"jtkolean/config"
	"jtkolean/task"
	"net/http"
)

func main() {
	c := config.NewConfig("./resources/application.yaml")
	db := c.ConnectDB()

	http.HandleFunc("/task", task.New(db).Handle)

	addr := fmt.Sprintf(":%v", c.Server.Port)
	http.ListenAndServe(addr, nil)
}
