package main

import "fmt"

type Old interface {
	OldPr(text string)
}
type New interface {
	NewPr(text string)
}
type Oldin struct{}

type Adapter struct {
	Old Old
}

func (o *Oldin) OldPr(text string) {
	fmt.Println("Oldin", text)
}

func (a *Adapter) NewPr(text string) {
	fmt.Println("NewPr", text)
}
func main() {
	oldpr := &Oldin{}
	newpr := &Adapter{Old: oldpr}
	var printer New = newpr
	printer.NewPr("Hello, World!")

}
