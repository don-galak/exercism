package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet() string
}

type Italian struct{}
type Portuguese struct{}

func (l Italian) LanguageName() string {
	return "Italian"
}
func (l Italian) Greet() string {
	return "Ciao"
}
func (l Portuguese) LanguageName() string {
	return "Portuguese"
}
func (l Portuguese) Greet() string {
	return "Olá"
}

func SayHello(name string, lang Greeter) string {
	return fmt.Sprintf("I can speak %s: %s %s!", lang.LanguageName(), lang.Greet(), name)
}
