package javaDoc

type Comment struct {
	text string
}

func (c *Comment) Markdown() string {	     
	return c.text
}

func (c *Comment) AppendText(text string) {
	c.text += cleanTag(text)
}

func NewComment(strComment string) *Comment {
    c := &Comment{}
    c.AppendText(strComment)
	return c
}
