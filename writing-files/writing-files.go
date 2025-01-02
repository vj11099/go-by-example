package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/media/vishvam-joshi/Data/projects/go-by-example/dat", d1, 0644)
	check(err)

	f, err := os.Create("/media/vishvam-joshi/Data/projects/go-by-example/dat1")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Println(string(d2))
	fmt.Printf("wrote %d bytes\n", n2)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("write %d bytes\n", n4)
	w.Flush()
}
