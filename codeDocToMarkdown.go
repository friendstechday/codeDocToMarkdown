package main

import (
    "flag"
    "github.com/friendstechday/codeDocToMarkdown/converter"
)

func main() {
    
    file := flag.String("f", "", "Source file")
	flag.Parse()
    
    java := &converter.Java{}
    java.Process(*file)

}
