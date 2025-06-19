package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// func main() {
// 	// var conferenceName string = "booking app"
// 	// var age int

// 	// var bookings []string
// 	// var mapper = map[string]string{"name": "gbasa"}

// 	// mapper["asas"] ="Sxasx"

// 	// bookings = append(bookings, "nana")

// 	// fmt.Println("Enter your age")
// 	// fmt.Scan(&age)

// 	// fmt.Printf("The content is %v and the type is %T \n",conferenceName, conferenceName)
// 	// fmt.Printf("The user entered age is %v\n", age)

// 	// for index, booking := range(bookings){

// 	// }

// 	a, b := sayHello(5)
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("The number is %v\n", i)
// 	}

// 	hide.SaySomething()

// 	var c = []int{1, 2, 3, 4, 5}
// 	// c = append(c, 6)
// 	// c = append(c, 6)
// 	c[3] = 6

// 	fmt.Println(c)

// 	type Person struct {
// 		name string
// 	}

// 	var user1 = Person{name: "Gbenga"}
// 	fmt.Println(user1.name)

// 	// user2 := struct {
// 	// 	name string
// 	// 	age  int
// 	// }{
// 	// 	name: "Ade",
// 	// 	age:  50,
// 	// }

// 	// var i = 42

// 	///////
// 	///// Pointers in GO
// 	//////
// 	var x = 5
// 	var y *int = &x

// 	fmt.Println(y)

// 	firstNum := 6
// 	secondNum := 7

// 	fmt.Println(calculate(&firstNum, &secondNum))

// 	//////
// 	///// Channels and Go routines
// 	//////
// }

// func calculate(firstNumber *int, secondNumber *int) int {
// 	return *firstNumber + *secondNumber
// }

// func sayHello(number int) (age int, school string) {
// 	return number, "hasas"
// }

// // /////
// // //// Generics in Go
// // /////
// func SumInts(m map[string]int64) int64 {
// 	var result int64

// 	for _, value := range m {
// 		result += value
// 	}

// 	return result
// }

// func SumFloats(m map[string]float64) float64 {
// 	var result float64

// 	for _, value := range m {
// 		result += value
// 	}

// 	return result
// }

// type Number interface {
// 	int64 | float64
// }

// func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
// 	var result V

// 	for _, value := range m {
// 		result += value
// 	}

// 	return result
// }

// func main2() {
// 	value1 := map[string]int64{
// 		"gbenga": 65,
// 	}

// 	value2 := map[string]float64{
// 		"gbenga": 65.67,
// 	}

// 	SumInts(value1)

// 	SumFloats(value2)

// 	SumIntsOrFloats(value1)
// 	SumIntsOrFloats(value2)
// }

/////
///// End of Generics
////

// ///
// //// Go routines
// ////
func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMsg(msg string, wg *sync.WaitGroup) {
	pause()
	fmt.Println(msg)
	defer wg.Done()
}

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go func(msg string) {
// 		pause()
// 		fmt.Println(msg)
// 		defer wg.Done()
// 	}("test1")
// 	go sendMsg("test2", &wg)
// 	go sendMsg("test3", &wg)

// 	wg.Wait()
// }

// func main() {
// 	// msgChan := make(chan string)

// 	// go func() {
// 	// 	time.Sleep(2 * time.Millisecond)

// 	// 	msgChan <- "Hello"
// 	// 	msgChan <- "World"
// 	// }()

// 	// msg1 := <-msgChan
// 	// msg2 := <-msgChan

// 	// fmt.Printf("Hey this are the messages %v, %v", msg1, msg2)

// 	msgChan := make(chan int, 3)
// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	go func() {
// 		for {
// 			a := <-msgChan

// 			fmt.Printf("Read %v\n", a)
// 			time.Sleep(2 * time.Second)

// 			if a == 5 {
// 				wg.Done()

// 				break
// 			}
// 		}
// 	}()

// 	for i := 0; i < 6; i++ {
// 		msgChan <- i

// 		fmt.Println(i)
// 	}

// 	wg.Wait()
// }

func reader(channel <-chan string) {
	message := <-channel

	fmt.Printf("The message is %v \n", message)
}

func writer(channel chan<- string, msg string) {
	pause()
	channel <- msg
}

// select
// func main() {
// 	var channel1 = make(chan string)
// 	var channel2 = make(chan string)

// 	go writer(channel1, "Hello")
// 	go writer(channel2, "World")

// 	for {
// 		var done bool
// 		select {
// 		case msg1 := <-channel1:
// 			fmt.Printf("The message is '%v'\n", msg1)
// 		case msg2 := <-channel2:
// 			fmt.Printf("The message is '%v'\n", msg2)

// 		case <-time.After(2 * time.Second):
// 			fmt.Println("I had to exited")

// 			done = true
// 		}

// 		if done {
// 			break
// 		}
// 	}

// 	channelIterator := make(chan string, 3)

// 	go func() {
// 		channelIterator <- "hello"
// 		channelIterator <- "hello2"
// 		channelIterator <- "hello3"
// 		close(channelIterator)
// 	}()

// 	for msg := range channelIterator {
// 		fmt.Println(msg)
// 	}

// 	var a = []interface{}{4, "Hello"}

// 	for _, b := range a {
// 		fmt.Println(b)
// 	}

// 	func(something interface{}) {
// 		// type switches only works on interfaces
// 		switch v := something.(type) {
// 		case int:
// 			fmt.Printf("Twice %v is %v\n", v, v*2)
// 		case string:
// 			fmt.Printf("%q is %v bytes long\n", v, len(v))
// 		default:
// 			fmt.Printf("I don't know about type %T!\n", v)
// 		}
// 	}(6) // You can pass anything as the args. All types in Go are a superset of the interface{} type

// }

//
//

func isDigit(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}

	return false
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}

	x := 0

	for ; i < len(b) && isDigit(b[i]); i++ {
		// This is one way to convert a byte slice to an integer
		x = x*10 + int(b[i]) - '0'
	}

	return x, i
}

func main() {
	// array := [4]byte{1, 2, 3, 4}

	// slice := array[1:3]
	// arrayPtr := unsafe.SliceData(slice)

	// fmt.Println(slice, len(slice), cap(slice))

	// println(&array)
	// println(arrayPtr)
	// println(slice)

	defer func() {
		fmt.Println("I am leaving this function. I run before function returns")
	}()

	integers, n := nextInt([]byte("This is just me 123 and so my case 78"), 0)

	fmt.Printf("The integers is %d and the stuff is %d \n", integers, n)

	fmt.Println(reflect.TypeOf("Gbenga"[0:3]))

	// var p *[]int = new([]int)

}

// The function is always called if present in a module
// func init() {
// 	fmt.Println("The init function is called")
// }

// func init() {
// 	fmt.Println("The init function is called ======> 2")
// }
