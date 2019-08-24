package main

import (
	"github.com/overcls/blog/router"
)

func main() {

	r :=router.InitRouter()
	_ = r.Run(":3000")
}
