package main

import (
	"fmt"
	"ginproject/app/myfile"
	"ginproject/app/user"
	"ginproject/routers"
)

func main() {
	routers.Include(user.Routers, myfile.Routers)
	r := routers.Init()

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	// listen and serve on 0.0.0.0:8080
}
