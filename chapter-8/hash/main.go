package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("Test"))
	v := h.Sum32()
	fmt.Println(v)

	h.Write([]byte("test"))
	cLength := h.Size()
	c := h.Sum32()
	fmt.Println(c, cLength)

	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	io.Copy(h, file)

	fmt.Println("File hash:", h.Sum32())

	hash := sha1.New()
	hash.Write([]byte("test"))
	fmt.Println("SHA1 hash:", hash.Sum(nil))
	bs := hash.Sum([]byte{})
	fmt.Println(bs)
}
