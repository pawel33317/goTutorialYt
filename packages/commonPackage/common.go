package commonPackage

import (
	"fmt"
	"runtime"
)

func PrintThreads() {
	fmt.Println("Number of CPU threadse: ",
		runtime.GOMAXPROCS(-1))
}
