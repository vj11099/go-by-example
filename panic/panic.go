package main

import "os"

func main() {
	panic("a problem occured")

	_, err := os.Create("/media/vishvam-joshi/Data/projects/go-by-example/file")
	if err != nil {
		panic(err)
	}
}
