package main

import (
	"os"

	"github.com/RyuseiAndy/go-vue-pay-pj/backapi/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
