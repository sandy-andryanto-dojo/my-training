package main

// import ("fmt"; "container/list")

// func main(){
// 	var x list.List
// 	x.PushBack(1)
// 	x.PushBack(2)
// 	x.PushBack(3)
// 	for e := x.Front(); e != nil; e=e.Next(){
// 		fmt.Println(e.Value.(int))
// 	}
// }

import ("fmt"; "sort")

type Person struct{
	Name string
	Age int
}

type ByName []Person
type ByAge []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool{
	return this[i].Name < this[j].Name
}

func (this ByName) swap(i,j int){
	this[i], this[j] = this[j], this[i]
}

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool{
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}

func main(){
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}