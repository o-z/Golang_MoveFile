package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./a")
	if err != nil {
		log.Fatal(err)
	}
	var i int
	for _, f := range files {
		File, err := os.Open("./a/" + f.Name())
		if strings.Contains(f.Name(), ".xml") {
			fmt.Println(f.Name())
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			fmt.Println(i)
			i++
			defer File.Close()

			decoder := xml.NewDecoder(File)
			if decoder == nil {
				fmt.Println("hata")
				return
			}

			fmt.Println(decoder)
			File.Close()
			err := os.Rename("./a/"+f.Name(), "./b/"+f.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
