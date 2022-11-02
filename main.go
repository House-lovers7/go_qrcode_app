package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("./hello,txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	fmt.Println("Done!")
}


// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	flag.Parse()
// 	arg := flag.Arg(0)
// 	fmt.Printf("Hello %s\n", arg)
// }
