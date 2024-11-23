package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Thing interface {
	getName() string
}

type Person struct {
	FirstName string
	LastName  string
	age       int
}

func (p Person) getName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) getNames() (string, string) {
	return p.FirstName, p.LastName
}

func (p Person) ageString() string {
	return string(p.age)
}

type Student struct {
	person Person
}

func main() {
	variables()
	types()
	arrays()
	// io()
	math()
	comparison(11, 12)
	logical(10, 12)
	conditions(12, 17)
	forLoop()
	forLoop2()
	switchSample()
	everything()
	printEverything()
	fibo(10)
	pointers()
	goroutines()
}

func types() {
	fmt.Println("\n\nGo Types")
	var getLastNameFirst = func(p Person) string {
		return p.LastName + ", " + p.FirstName
	}

	var person1 = Person{}
	person1.LastName = "Aguinaldo"
	person1.FirstName = "Emilio"
	var t Thing = person1

	var s = Student{person1}

	var x, y = person1.getNames()

	fmt.Println(person1.getName())
	fmt.Println(getLastNameFirst(person1))
	fmt.Println(x, y)
	fmt.Println(s.person.getName())
	fmt.Println(t)
}

func arrays() {
	fmt.Println("\n\nGo Arrays")
	var x [5]int
	x = [5]int{1, 2, 3, 4, 5}
	var y = [3]int{1, 2, 3}
	var p = [2]Person{}
	p[0] = Person{"Manuel", "Quezon", 40}
	p[1] = Person{"Manuel", "Roxas", 40}
	var s = y[:1] // slice
	s = append(s, 10)
	s = append(s, 11)
	s = append(s, 12)
	s[1] = 200
	fmt.Println("S is", s)
	fmt.Println("y is", y)
	var t = y[1:]  // slice
	var u = x[1:1] // slice
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("p", p)
	fmt.Println("s", s)
	fmt.Println("t", t)
	fmt.Println("u", u)

	i, str := multipleReturn()
	fmt.Println("multiple return values", i, str)
}

func multipleReturn() (int, string) {
	return 42, "the universe"
}

func io() {
	var i, j string
	n, err := fmt.Scanln(&i, &j)
	fmt.Println("n:", n, "err:", err)
	fmt.Println("i:", i)
	fmt.Println("j:", j)

	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	fmt.Printf("The input is i:%s and j:%s\n", i, j)
	fmt.Println("ReadString returned:", line)
}

func math() {
	fmt.Println("\n\nGo Math Operations")
	var x, y int = 10, 27
	sum := x + y
	diff := x - y
	product := x * y
	quotient := x / y
	mod := x % y
	fmt.Println("sum:", sum, "diff:", diff)
	sum++
	fmt.Println("sum++", sum)
	sum--
	fmt.Println("sum--", sum)
	fmt.Println("product:", product, "quotient:", quotient, "mod:", mod)
	fmt.Println()

	var a, b float64 = 12.5, 23.1
	fSum := a + b
	fDiff := a - b
	fProduct := a * b
	fQuotient := a / b
	a++
	fmt.Println("a:", a, "b:", b)
	fmt.Println("f_sum:", fSum, "f_diff:", fDiff, "f_product:", fProduct, "f_quotient:", fQuotient)

	sumXA := float64(x) + a // error without explicit type conversion
	fmt.Println("sum_x_a", sumXA)

	var aa, bb byte = 66, 67
	cc := aa & bb
	dd := aa | bb
	fmt.Printf("aa: %b bb: %b cc: %b dd: %b", aa, bb, cc, dd)
}

func comparison(a, b int) {
	fmt.Println("\n\nGo comparison operators")
	var ab = a == b
	var neq = a != b
	var gt = a > b
	var lt = a < b
	var gte = a >= b
	var lte = a <= b

	fmt.Println("a:", a, "b:", b)
	fmt.Println("a == b", ab)
	fmt.Println("a != b", neq)
	fmt.Println("a > b ", gt)
	fmt.Println("a < b ", lt)
	fmt.Println("a >= b", gte)
	fmt.Println("a <= b", lte)
}

func logical(a, b int) {
	fmt.Println("\n\nGo logical operators")
	var and = a < 5 && b%2 == 0
	var or = a > 0 || a%2 == 0
	var not = !(a < 0)

	fmt.Println("a:", a, "b:", b)
	fmt.Println("a < 5 && b % 2: ", and)
	fmt.Println("a > 0 || a % 2 == 0: ", or)
	fmt.Println("!(a < 0): ", not)
}

func conditions(a, b int) {
	fmt.Println("\n\nGo conditional operators")
	fmt.Println("a:", a, "b:", b)
	if a > b {
		fmt.Println("a is greater than b")
	}
	if a == 0 {
		fmt.Println("a is zero")
	} else {
		fmt.Println("a is not zero")
	}
	if b > a {
		fmt.Println("b is greater than a")
	} else if b < a {
		fmt.Println("b is less than a")
	} else {
		fmt.Println("b is equal to a")
	}
}

func forLoop() {
	fmt.Println("\n\nGo For Loop")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}

func forLoop2() {
	fmt.Println("\n\nGo For Loop that looks like a while loop")
	number := 15
	for number > 0 {
		if (number%3 == 0) && (number%5 == 0) {
			fmt.Print("FizzBuzz")
		} else if number%3 == 0 {
			fmt.Print("Fizz")
		} else if number%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(number)
		}
		number--
		fmt.Println()
	}
}

func switchSample() {
	fmt.Println("\n\nGo Switch Sample")
	var x = 1

	switch x {
	case 1:
		fmt.Println("uno")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}

func printEverything() {
	fmt.Println(42)
}

func everything() int {
	return 42
}

func variables() {
	fmt.Println("\n\nGo Variables")
	var one int = 1
	var isDeclared bool = true
	var ascii byte = 255
	var three string = "three"
	var pi float64
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	var r rune = 123
	pi = 3.1415926
	var two = 2 // inferred type
	uno := one  // short assignment
	dos := 2.0
	four := "four" // short assignment
	const e float64 = 2.71828

	var f float64 = 1

	var double = func(x int) int {
		return x * 2
	}

	doubled := double(10)

	isDeclared = false
	fmt.Println(one, isDeclared, ascii, e, two, three, pi, four, x, y, r, uno, f, dos, doubled)

	name := "Juan dela Cruz"

	fmt.Println("Hello and welcome, %s!", name)
}

func fibo(n int) []int {
	var f []int

	f = append(f, 1)
	f = append(f, 1)
	for i := 2; i < n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	fmt.Println(f)
	return f
}

func pointers() {
	fmt.Println("\n\nGo Pointers")
	var i int = 10
	var address_i *int = &i

	fmt.Println(address_i, *address_i)
	change(address_i)
	fmt.Println(*address_i)
}

func change(i *int) {
	var x = 17
	*i = x
}

func double(ch chan int, x int, wg *sync.WaitGroup) {
	var ans = x * 2
	ch <- ans
	wg.Done()
}

func times2(x int, ch chan int) {
	ch <- x * 2
}

func goroutines() {
	fmt.Println("\n\nGo Goroutines")
	var wg sync.WaitGroup

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	for _, i := range tasks {
		wg.Add(1)
		go double(ch, i, &wg)
	}

	results := make([]int, len(tasks))
	for i := range results {
		results[i] = <-ch
	}
	wg.Wait()

	fmt.Println("results", results)

	go times2(123, ch)

	z := <-ch
	fmt.Println(z)

}
