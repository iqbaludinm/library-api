package main

import (
	"github.com/iqbaludinm/library-api/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
