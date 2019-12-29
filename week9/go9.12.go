package main

import "fmmt"

ty I interface{
	F()
}

func desc(i I)  {
	fmt.Println("%v , 5T \n",i,i)	
}

type T1 struct {
	text string
}