package decorator

import "fmt"

func HandlerSayGoodBye() *HandlerSayGoodBye {
	return &HandlerSayGoodBye{}
}

func (h *HandlerSayGoodBye) Process() error {
	fmt.Println("bye")
	return nil
}
