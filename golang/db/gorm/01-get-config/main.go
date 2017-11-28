package main

import (
	"flag"
	"fmt"
	"runtime"
)

// getConfigFilePath: 获取config file 的地址
func getConfigFilePath(configFlag string) string {

	if configFlag == "" {
		if runtime.GOOS == "darwin" {
			return "~/.jindowin/config.toml"
		}
		return "/root/.jindowin/config.toml"
	}

	return configFlag
}

func main() {

	configPtr := flag.String("c", "", "config file dest")
	dbinitPtr := flag.Bool("d", false, "db init flag default false")
	flag.Parse()

	configFlag := getConfigFilePath(*configPtr)

	fmt.Println("config file:", configFlag)
	fmt.Println("dbinit flag:", *dbinitPtr)
}
