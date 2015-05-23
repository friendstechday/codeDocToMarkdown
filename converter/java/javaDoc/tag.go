package javaDoc

import (
    "strings"
    "regexp"
)

type Tag interface {	
	Markdown() string
	AppendText(string)
}


func cleanTag (strTag string) string {
    
    regex := regexp.MustCompile("(\\*|/)")
    str := strings.TrimSpace(regex.ReplaceAllLiteralString(strTag,""))  
    
    str += "\n"
    
    return str 
}