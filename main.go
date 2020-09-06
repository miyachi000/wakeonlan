package main

import (
	mydns "github.com/miyachi000/myddns/app/myddns"
)

func main() {
	srv := mydns.NewServer()

	srv.Run(":8082")
}
