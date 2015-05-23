package converter

import (
    "fmt"    
    "github.com/friendstechday/codeDocToMarkdown/utils"
    "github.com/friendstechday/codeDocToMarkdown/converter/java"
)

type Java struct{
}

func (j Java) Process(src string){
    
    strClass := utils.ReadFile(src)
    
    class := java.NewClass(strClass)
    fmt.Println(class.ToMarkdown())
}