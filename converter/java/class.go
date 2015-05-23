package java

import (    
	"regexp"
    "strings"
    "github.com/friendstechday/codeDocToMarkdown/utils"    
)

type Class struct {
	Body    string
	Name    string
	Package string
	Methods []*Method
}

func NewClass(strClass string) *Class {
    
    class := &Class{}

	reMethod, _ := regexp.Compile(`(public)([^{;/]+)`)
	idxMatch := reMethod.FindAllStringSubmatchIndex(strClass, -1)
    
    countMethods := len(idxMatch) - 1
    
    for i := 0 ; i <= countMethods ; i++ {
        
        var initDoc int
        
        if i == 0 {
            initDoc = 0
        }else{
            initDoc = idxMatch[i-1][5]
        }
        
        endDoc := idxMatch[i][0]
        
        
        initName := idxMatch[i][0]
        endName := idxMatch[i][5]
        
        method := NewMethod(strClass[initName:endName],strClass[initDoc:endDoc])
        
        class.Methods = append(class.Methods,method)

    }
    
    class.Name = getClassName(strClass)
    class.Package = getPackage(strClass)    


	return class

}


func getClassName(str string) string{
    idxClass := strings.Index(str,"class")
    idxNewLine := strings.Index(str[idxClass:],"{") + idxClass
    
    return strings.TrimSpace(str[idxClass:idxNewLine])
}

func getPackage(str string) string{
    idxPackage := strings.Index(str,"package")
    idxNewLine := strings.Index(str[idxPackage:],"\n") + idxPackage
    
    return strings.TrimSpace(str[idxPackage:idxNewLine])
}

func (c *Class) ToMarkdown() string{
    mark := utils.Concat("#",c.Package,"\n\n") 
    
    for _, m := range c.Methods {
        mark = utils.Concat(mark,m.ToMarkdown())
    }
    
    return mark   
}