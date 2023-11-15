package main

import (
	"fmt"
	"math"
	"strings"

)

// tipos de argumentos podem ser omitidos se forem o mesmo
func adicionar_sub(a , b int) (x, y int) {
	x = a+b
	y = a-b
	return
}



func main() {
	fmt.Println("Hello World")

	// variables and types
	var verdadeiro, falso, verdadeiro2 bool = false, true, true

	verdadeiro = true
	palavra, idade, eh_pessoa := "jose", 20, true

	// const
	const Pi = 3.1415

	fmt.Println(Pi)

	// variables inside blocks
	var (
		variavel1 string  = "variavel1"
		maxInt    int     = 12
		floatVar  float32 = 15.132
		floatSqrt float64 = math.Sqrt(64)
	)

	// conversion (casting)
	maxIntUint := uint(maxInt)

	// fmt.Println(variavel1, maxInt, floatVar) // %v for generic type
	fmt.Printf("Type: %T  value: %v\n", variavel1,variavel1)
	fmt.Printf("Type: %T  value: %v\n", maxInt,maxInt)
	fmt.Printf("Type: %T  value: %v\n", floatVar,floatVar)
	fmt.Printf("Type: %T  value: %v\n", maxIntUint,maxIntUint)
	fmt.Printf("Type: %T  value: %v\n", floatSqrt,floatSqrt)

	// condicoes
	if (verdadeiro) {
		fmt.Println("verdadeiro!")
		if falso {
			if verdadeiro2  {
				fmt.Println("nossa!")
				fmt.Println(palavra, idade, eh_pessoa)

			}
		}
		
	}

	// auto typing
	soma, sub := adicionar_sub(10, 5)

	var a int =10;
	a= 18;

	fmt.Printf("ad: %d sub: %d a: %d\n", soma, sub, a)

	// structs

	type Vertice struct {
		X int
		Y int
	}

	ponto1 := Vertice{1, 2}
	var ponto2 Vertice = Vertice{X: 4 } // Y=0 implicity
	var ponto3 Vertice

	ponto3.X = 10
	ponto3.Y = 15

	fmt.Printf("Pontos: %v %v %v\n", ponto1, ponto2, ponto3)


	// pointers to struct
	var pPonto1 *Vertice

	pPonto1 = &ponto1

	pPonto1.X = 90 // changing values from pointer without having to explicit dereference
	fmt.Printf("changing structures : %v\n", ponto1)

	// Arrays

	var array_of_string [10]string

	var int_array = [4]int{1, 2, 3, 4}
	var slice_of_int []int = int_array[1:3]
	fmt.Printf("Array of ints %v\n" +
	"slice of arrays: %v\n", int_array, slice_of_int)

	array_of_string[0] = "string 1 "
	array_of_string[1] = "string 2\n"

	fmt.Print(array_of_string[0], array_of_string[1])

	// slices
	var s []int // slice s of int (default is nil)
	if s == nil {
		fmt.Println(s, cap(s), len(s))
	}

	var slice_of_struct = []struct{
		X int 
		Y int; b bool
		}  {
			{0, 1,true},
			{13,145,false},
			{1,-46, true},
		}
	
	fmt.Println(slice_of_struct)

	// using make
	make_a := make ([]string, 1, 5 ) // len(make_a) = 1, cap(make_a) = 5 (capacity = length of the array, len = length of the slice)
	make_a[0] = "aaaa"

	// example of a tic tac toe board

	board := [][]string {
		{"_", "_", "_"},// each
		{"_", "_", "_"},// element
		{"_", "_", "_"},// here
	}
	/*It could be done like: (redundant)
	{
		[]string{"string"}
		[]string{"definition", "here"}
		[]string{"hehe", "."}
	}
		*/

	for i:= 0 ; i < len(board) ; i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " ")) // joins each element of board in a row separated by " "
	}

	fmt.Println("Printing using range:")

	for i, v := range board {
		for j, u := range v {
			fmt.Printf("i: %d, j: %d, string: %s ", i, j ,u)
		}
		fmt.Println()
	}

	// for only the index:
	for i := range board {
		fmt.Println("lines: ", i)
	}

	array := make([]string, 5)

	array = append(array, "apendando")

	fmt.Printf("%s\n", array)

	
	r := make([][]int , 10)

	for i := range r {
		// fmt.Println(a)
		r[i] = make([]int, 3)
		fmt.Println(r[i])
	}

	// Maps
	var m map[string]int // so far is a nil map, we have to use the make to create and intialize the map
	m = make(map[string]int)

	m["inteiro1"] = 0
	m["inteiro2"] = 5
	
	fmt.Printf("Inteiro1: %d; Inteiro2: %d\n", m["Inteiro1"], m["Inteiro2"])

	type StructMap struct {
		segredo, segredo2 uint8
	}
	// var mapa map[int]StuctMap

	mapa := make(map[int]StructMap)

	mapa[0] = StructMap{1, 2}
	mapa[3] = StructMap{1, 4}

	for i, v := range mapa { // i returns the keys
		fmt.Printf("i: %v, struct: %v\n", i, v)
	}

	// map literals
	var map_literal = map[string]StructMap{
		"aaa string1" : StructMap{2, 3},
		"asdlfasdfasdklf" : {7,8}, // we doesnt need to tell its a StructMap since we already defined
	}

	for i, v := range map_literal { 
		fmt.Printf("i: %v, struct: %v\n", i, v)
	}

	fmt.Println(map_literal)

	// testing if the key is present

	elem, ok := map_literal["key qye nao existe"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("Key nao existe")
	}


	// messing with functions and closures
	function_as_var := func (x, y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}

	fmt.Println(function_as_var(10, 12))

	func_var := closure_function()

	func_var(1)
	func_var(2)
	i := func_var(3)

	fmt.Println(i) // prints 6
	
	// printing fibonacci sequence using closures
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

	// Methods and interfaces

	

	// func aaaaaaaa(aa int) (x, y int) {
	// 	x = 1
	// 	y = 1
	// 	return
	// }

	// func split(sum int) (x, y int) {
	// 	x = sum * 4 / 9
	// 	y = sum - x
	// 	return
	// }

	obj := Objeto5{10, 56}

	fmt.Println()
	x, y := obj.MethodObj5()
	fmt.Printf("Metodo obj 5: %d %d\n",x, y )
	
	obj.MethodObj5Pointer()
	x, y = obj.MethodObj5()
	fmt.Printf("Metodo obj ref 5: %d %d\n",x, y )

	// interfaces (see code below main function)
	

	var GenericGameObject1 GameObject // GameObject is an interface
	GenericGameObject1 = &Mario{1, 2, true}

	var GenericGameObject2 GameObject
	GenericGameObject2 = &Koopa{10, 56.23 , false}

	for i:= 0; i < 4; i++ {
		fmt.Println()
		GenericGameObject1.GameMethod()
		GenericGameObject2.GameMethod()
	}

	// but what are the advantages????
	// we can pass interfaces as function arguments so it keeps a generic type

	for i:= 0; i < 4; i++ {
		fmt.Println()
		InterfacMethod(GenericGameObject1)
		InterfacMethod(GenericGameObject2) // using the same function to pass mario and koopa types
	}

	// getting the concrete type T from interfae

	mario, ok := GenericGameObject1.(*Mario) // getting the concrete value of mario and the ok=true if success
	fmt.Println(mario, ok) // &{1 2 true} true

	// switch cases using interface types

	switch v := GenericGameObject1.(type) {
	case *Mario:
		fmt.Println("GenericGameObject1 is mario !")
		
	
	case *Koopa:
		fmt.Println("GenericGameObject1 is koopa !")
		
	default:
		
		fmt.Println("GenericGameObject1 is neither mario nor koopa !")
		// we used v:= so we can use v again, but if we dont need, just switch the type is enough
		v.GameMethod()
	}

	// showing the stringer interface
	fmt.Println()
	fmt.Println(GenericGameObject1) // Mario with x: 1, y: 2 and state alive: true true (has the Strint method defined from Stringer interface)
	fmt.Println(GenericGameObject2) // &{10 56.23 false} (does not have String method defined)

	// errors
	// we can define our own error type by implementing the Error() string interface
	// example defined after main
	

	newPosition, error := MoveMario(11) 

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Success! New position at ", newPosition)
	}
	
}// main

// errors ////////////////////////////////////////////////////////////
type InvalidMarioPosition struct {
	X int
	Y int
	message string
}

// implementing error
func (i *InvalidMarioPosition) Error() string {
	return fmt.Sprintf("Invalid position: x: %d, y: %d, message: %s", i.X, i.Y, i.message)
}

func MoveMario(x int) (int, error) {
	if x < 10 {
		return x+1, nil // error = nil means success
	} else {
		return 0, &InvalidMarioPosition{x, 10, "muito longe!"} // returns an error 
	}
} 

/////////////////////////////////////////////////////////////////////////
// messing with interfaces (like abstractions)
type GameObject interface { // interfce (like a virtual stuff)
	GameMethod() bool
}

type Mario struct {
	X int
	Y int
	isAlive bool
}

// we HAVE to implement exactly like the interface definition (in this case, no argument and bool return)
func (m *Mario) GameMethod() bool { // implicitly implementing interface GameObject on *Mario (it has to be a pointer since its implemented here as a pointer)
	fmt.Printf("Printing mario and changing its state: x: %d y: %d isalive: %v\n", m.X, m.Y, m.isAlive)
	m.isAlive = !m.isAlive
	return !m.isAlive
}

func (m *Mario) String() string { // Stringer interface, most used to print its values. Allows: fmt.Println(mario)
	return fmt.Sprintf("Mario with x: %d, y: %d and state alive: %v", m.X, m.Y, m.isAlive)
}

type Koopa struct {
	killCount int
	distanceToMario float64
	isAlive bool
}

func (k *Koopa) GameMethod() bool {
	fmt.Printf("Printing koopa information and changing its state: kc: %d dtm: %f isalive: %v\n", k.killCount, k.distanceToMario, k.isAlive)
	return k.isAlive
}

func InterfacMethod(i GameObject) bool {
	i.GameMethod()
	return true
}
// implementing on ip
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
	
}
// when using fmt.Println(ip) -> comes.on.this.format

/// other stuff


type Objeto5 struct {
	X int
	Y int
}

// method with value receiver
func (o Objeto5) MethodObj5() (x, y int) {
	x = o.X
	y = o.Y
	return
}

// method with pointer receiver
func (o *Objeto5) MethodObj5Pointer() {
	o.X *= 2
	o.Y *= 2
	
}

func closure_function() func (int) int { // closer function that returns a functions with int parameter and int return value

	variable := 0

	return func(x int) int { // function that is returndes
		variable += x
		return variable
	}

}

 // function make on go tour (PASSED)
func WordCount(s string) map[string]int {
	map_count := make(map[string]int)
	
	splitted := strings.Fields(s)
	
	for _, v := range splitted {
		_, ok := map_count[v]
		
		if ok {
			map_count[v] += 1
		} else {
			map_count[v] = 1
		}
	}
	
	return map_count
	//return map[string]int{"x": 1}

	

}

// fibonacci using closures
func fibonacci() func(int) int {
	x_1 := 0
	x_2 := 0

	//x_2 := 0
	return func(x int) int {
		if x == 0 {
			x_2 = 0
			return 0
		} else if x==1 {
			x_1 = 1
			
			return 1
		}
		r := x_1 + x_2
		x_2 = x_1
		x_1 = r
		return x_1
	}
}