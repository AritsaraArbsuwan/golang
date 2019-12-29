package main

import "fmt"

type student student {
	name string
	age int
	email string
}

func main(){
	var a student
	a.name = "Goku"
	a.age  = 18
	a.email = "Goku@super.saiya"

	b :=student{"Gohan", 3, "Gohan@super.saiya"}

	c :=student{name: "videl",email: "videl@daughter.satan"}

	d :=student{age: 20}

	fmt.Println(a)
	f
}
