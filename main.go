/*
>main.exe --help
Usage of main.exe:
-env string
		environment to be compiled (default "dev")
>
>main.exe -env=prod
Environment: prod
*/
package main

import (
	"flag"
	"fmt"
)

func main() {

	env := flag.String("env", "dev", "environment to be compiled")

	flag.Parse()

	fmt.Println("Environment:", *env)
}
