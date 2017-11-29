package main

import (
	"fmt"
	"log"
	"os/user"
	"path"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)
	fmt.Println(path.Join(usr.HomeDir, "/.jindowin/config.toml"))
}
