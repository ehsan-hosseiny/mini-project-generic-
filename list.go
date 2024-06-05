package main

import (
	"fmt"
	"genericList/generics"
)

type Person struct{
	Name string
	Age int
}

func main(){
	
	genericInt()
	genericString()
	genericPerson()
}

func genericInt(){
	list1 := generics.List[int]{Items:[]int{}}
	
	list1.Add(10)
	list1.Add(20)
	list1.Add(30)

	fmt.Printf("%v\n",list1.Items)

	list1.InsertAt(1,50)

	fmt.Printf("%v\n",list1.Items)

	list1.RemoveAt(2)

	fmt.Printf("%v\n",list1.Items)

	list1.Remove(30)

	fmt.Printf("%v\n",list1.Items)
}

func genericString(){
	list1 := generics.List[string]{Items: []string{}}

	list1.Add("John")
	list1.Add("Sara")
	list1.Add("Joe")
	list1.Add("Kate")
	fmt.Printf("%v\n",list1.Items)

	list1.Remove("Sara")
	fmt.Printf("%v\n",list1.Items)

	list1.InsertAt(2,"k2")
	fmt.Printf("%v\n",list1.Items)

}

func genericPerson(){
	list1 := generics.List[Person]{Items: []Person{}}

	list1.Add(Person{Name:"Brad",Age:40})
	list1.Add(Person{Name:"Nicole",Age:20})
	list1.Add(Person{Name:"Rose",Age:35})

	fmt.Printf("%v\n",list1.Items)

	list1.Remove(Person{Name:"Nicole",Age:20})
	fmt.Printf("%v\n",list1.Items)


	list1.InsertAt(2,Person{Name:"Mike",Age:30})



}


