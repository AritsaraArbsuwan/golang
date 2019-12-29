package main

import "fmt"

func hello(t interface{}) {
	fmt.Printf("Hello %T \n", t)
}
func hi(t ...interface{})  {
	fmt.Printf("Hi ")
	for _, v := rande t {	
		fmt.Println(" %T,", v)
}
