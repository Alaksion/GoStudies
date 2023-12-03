package main

type Person struct {
	name string
	age  int
}

func main() {

	const x = 10
	const y = 5

	var person1 Person

	person1.name = "Hello"
	person1.age = 10

	println(operation(x, y, sum))
	print(operation(x, y, subtract))

}

func operation(value1 int, value2 int, transformation func(int, int) int) int {
	return transformation(value1, value2)
}

func sum(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
