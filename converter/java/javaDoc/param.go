package javaDoc

type Param struct {
	text string
}

func (p *Param) Markdown() string {
    mk := "* "
    mk += p.text
	return mk
}

func (p *Param) AppendText(text string) {     
	p.text += cleanTag(text)
    
}

func NewParam(strParam string) *Param {
    p := &Param{}
    p.AppendText(strParam)
	return p
}
