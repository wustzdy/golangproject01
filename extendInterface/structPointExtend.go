package main

import "fmt"

type Usber interface {
	start()
	stop()
}
type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println(p.name, "自动")

}
func (p Phone) stop() {
	fmt.Println(p.name, "关机")

}

func main() {
	p := Phone{
		name: "华为手机",
	}

	var p1 Usber

	p1 = p
	p1.start()
	p1.stop()

}
