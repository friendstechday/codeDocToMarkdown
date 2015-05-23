package java

import(
    "strings"    
    "github.com/friendstechday/codeDocToMarkdown/converter/java/javaDoc"
    "github.com/friendstechday/codeDocToMarkdown/utils"
)

type Method struct {
    Name string    
    JavaDoc *javaDoc.JavaDoc
}


func NewMethod(name string, strDocMethod string) *Method {
    init := strings.Index(strDocMethod, "/**")
    
    var jDoc *javaDoc.JavaDoc
    
    if init != -1 {        
        jDoc = javaDoc.NewJavaDoc(strDocMethod[init:])
    }
    
    return &Method{name,jDoc}
    
}


func (m *Method) ToMarkdown() string{
    var strJDoc string
    
    if m.JavaDoc != nil {
        strJDoc = m.JavaDoc.ToMarkdown()
        
    }
    
    
    return utils.Concat("## ",m.Name,"\n\n",strJDoc,"\n\n")
}