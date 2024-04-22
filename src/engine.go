package Groupie

import "fmt"

type Engine struct {

}

func (E *Engine) Init() { }

func (E *Engine) Run() {
	E.Init()
	fmt.Printf("grzdv")
}