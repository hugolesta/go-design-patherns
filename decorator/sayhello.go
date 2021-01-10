package decorator

import "fmt"

type HanderSayHello struct{}

func NewHanderSayHello() *HandlerSayHello {
	return &handlerSayHello{}
}

func (h *handlerSayHello) Process() error {
	fmt.Println("Hello")
	return nil
}
