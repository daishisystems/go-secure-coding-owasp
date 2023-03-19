package char

import "html"

type Char struct{}

func Newchar() *Char {
	return &Char{}
}

func (c *Char) Escape(text string) string {
	return html.EscapeString(text)
}
