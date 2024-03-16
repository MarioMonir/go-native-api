package main

import (
	"api/utils/http_server_utils"
	"api/utils/load_env_utils"
)

func main() {
	load_env_utils.LoadEnv()
	http_server_utils.LaunchHttpServer()
}
