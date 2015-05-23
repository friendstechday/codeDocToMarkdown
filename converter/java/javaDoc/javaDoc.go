package javaDoc

import (    
	"bufio"
	"bytes"
	"regexp"    
)

type JavaDoc struct {
	Tags []Tag
}

func (j *JavaDoc) appendTag(tag Tag) {
	j.Tags = append(j.Tags, tag)
}

func (j *JavaDoc) appendTextLastElement(text string) {

	var size int = len(j.Tags)

	if size > 0 {
		size--
	}

	j.Tags[size].AppendText(text)

}

func (j *JavaDoc) ToMarkdown() string{
    if len(j.Tags) <= 0 {
        return ""
    }
    
     var buff bytes.Buffer
    
    for _, t := range j.Tags {        
        buff.WriteString(t.Markdown())
        buff.WriteString("\n")
    }
    
    return buff.String()
}

func NewJavaDoc(strJavaDoc string) *JavaDoc {

	rComment, _ := regexp.Compile(`(^\s+.\*|^\*)`)
	rTags, _ := regexp.Compile(`(@param|@)`)

	jDoc := &JavaDoc{}

	jDoc.appendTag(NewComment(""))

	r := bytes.NewReader([]byte(strJavaDoc))
	scan := bufio.NewScanner(r)

	for scan.Scan() {
		li := scan.Text()

		if !rComment.MatchString(li) {
			continue
		}

		tag := rTags.FindString(li)

		switch tag {            
            case "@param", "@":        
                jDoc.appendTag(NewParam(li))
            default:
                jDoc.appendTextLastElement(li)

		}

	}

	return jDoc
}
