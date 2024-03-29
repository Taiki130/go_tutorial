package main

import "fmt"

func Eat(name string) bool {
	if name == "" {
		return false
	}
	// <nameが空白以外ならば、`fmt.Println(name)`を実行し、`return true`を行う>
	 fmt.Println(name)
	 return true
}

func main() {
	var name1 string = "GYUDON"
	if ok := Eat(name1); !ok {
		fmt.Println("cannt eat: ", name1)
	}

	var name2 string = ""
	if ok := Eat(name2); !ok {
		fmt.Println("cannt eat: ", name2)
	}
}
