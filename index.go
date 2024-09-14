package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	age      int
	position string
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	width int
	height int
}
 type Shape interface{
	Area()
	Perimeter()
 }
 type Vehicle interface {
	Start()
	Stop()
 }
 type Car struct {
	model string
 }
 type Bike struct {
	model string }
func (c Car) Start() {
	fmt.Printf("%s dang khoi dong...\n", c.model)
}
func (c Car) Stop() {
	fmt.Printf("%s dang ngat dong...\n", c.model)
}

func (b Bike) Start() {
	fmt.Printf("%s dang khoi dong...\n", b.model)
}
func (b Bike) Stop() {
	fmt.Printf("%s dang ngat dong...\n", b.model)
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
 }
func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
	}
func (r Rectangle) Area() float64 {
	return float64(r.width) * float64(r.height)
	}
func (r Rectangle) Perimeter() float64 {
		return 2 * (float64(r.width) + float64(r.height))
	}
func maxAge(employees []Employee) Employee {
	max := employees[0]
	for _, employee := range employees {
		if employee.age > max.age {
			max = employee
		}
	}
	return max
}

func chuVi (radius float64) float64 {
	return 2 * math.Pi * radius
}
func dienTich(radius float64) float64 {
	return math.Pi * radius * radius
}
func hoandoi (a,b *int) {
	tam := *a;
	*a = *b;
	*b = tam;
}
func tang (a *int) {
	*a += 1
}
func giaithua (n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n*giaithua(n-1)
}
func fibonacci(n int) {
    if n <= 1 {
        return
    }
    printFibonacci(n, 1, 1)
}
func printFibonacci(n int, a, b int) {
	if n > 0 {
        fmt.Print(a)
        if n > 1 {
            fmt.Print(", ")
        }
        printFibonacci(n-1, b, a+b)
    }
}
func chia(a int, b int) int {
	defer func () {
		if r:=recover(); r != nil {
			fmt.Println()
			fmt.Println("Error:", r)
		}
	}()
    if b == 0 {
        panic("Lỗi: không thể chia cho 0")
    }
    return a / b
}
func sayHello() {
    fmt.Println("Hello!")
}

func sayGoodbye() {
    fmt.Println("Goodbye!")
}
func tong(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
func max(nums ...int) int {
    if len(nums) == 0 {
        return 0 
    }
    highest := nums[0]
    for _, num := range nums {
        if num > highest {
            highest = num
        }
    }
    return highest
}
func main() {
	employees := []Employee{
		{"Nam", 20, "DEV C"},
		{"Khoa", 25, "BA"},
		{"Huy", 30, "DEV PHP"},
	}
	fmt.Println("Information of user oldest: ", maxAge(employees))
	radius := 5.0;
	circle := Circle{radius: radius}
	fmt.Printf("Chu vi cua hinh tron la: %.1f\n", chuVi(circle.radius))
	fmt.Printf("Dien tich cua hinh tron la: %.1f\n", dienTich(circle.radius))
	rectangle := Rectangle{width: 5, height: 7}
	fmt.Printf("Chu vi cua hinh tron (interface) la: %.1f\n", circle.Perimeter())
	fmt.Printf("Dien tich cua hinh tron (interface) la: %.1f\n", circle.Area())
	fmt.Println("Chu vi cua hinh vuong (interface) la: ", rectangle.Perimeter())
	fmt.Println("Dien tich cua hinh vuong (interface) la: ",rectangle.Area())
	car := Car{model: "Tesla Model 3"}
	bike := Bike{model: "Yamaha"}
	car.Start()
	car.Stop()
	bike.Start()
	bike.Stop()
	a := 4
	b := 7
	fmt.Println("Gia tri truoc khi doi a: ", a, " ; b: ",b)
	hoandoi(&a,&b)
	fmt.Println("Gia tri sau khi doi a: ",a," ; b: ", b)
	x := 20
	fmt.Println("Gia tri ban dau x: ", x)
	tang(&x)
	fmt.Println("Gia tri khi tang len 1 don vi x: ", x)
	fmt.Println("Giai thua cua 7: ", giaithua(7))
	fibonacci(6)
	chia(5, 0)
	//defer sayGoodbye()
	//sayHello()
	//panic("Panic xay ra")
	fmt.Println("Tong cua 1 2 3 4 6 la: ", tong(1,2,3,4,6))
	fmt.Println("So lon nhat trong day 2 4 7 8 11 -2 -23 la: ", max(2,4,7,8,11,-2,-23))
}