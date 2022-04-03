package main

import (
	"fmt"
	"os"
)

func ev(key string) string {
	return os.Getenv(key)
}
func main() {

	token := ev("token")
	fmt.Println(token)

}
