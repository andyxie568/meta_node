package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/

func numPlus(num *int) {
	*num += 10
}

func pointerPlus() {
	num := 0
	numPlus(&num)
	fmt.Println(num)
}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func multSlice(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
}

// Goroutine
/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func goroutinePrintNum() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i = i + 2 {
			fmt.Println("Goroutine1:", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i = i + 2 {
			fmt.Println("Goroutine2:", i)
		}
	}()
	wg.Wait()
}

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

type Task struct {
	Name string
	Func func()
}

func taskScheduler(taskList []*Task) {
	var wg sync.WaitGroup
	scheduler := func(taskList []*Task) {
		wg.Add(len(taskList))
		for _, task := range taskList {
			go func(t *Task) {
				defer wg.Done()
				start := time.Now()
				fmt.Printf("开始执行任务: %s\n", t.Name)
				t.Func()
				duration := time.Since(start)
				fmt.Printf("任务 %s 执行时间: %v\n", t.Name, duration)
			}(task)
		}
	}
	scheduler(taskList) // 启动任务调度器
	wg.Wait()
}

// 面向对象
/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，
并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func calcShape(shape Shape) {
	fmt.Println(shape.Area(), shape.Perimeter())
	fmt.Println(shape.Area(), shape.Perimeter())
}

func Calc() {
	r1 := &Rectangle{Width: 10, Height: 20}
	c1 := &Circle{Radius: 5}
	calcShape(r1)
	calcShape(c1)
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println(e.Name, e.Age, e.EmployeeID)
}

// Channel
/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func channelPrintNum() {
	wg := sync.WaitGroup{}
	numChan := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if v, ok := <-numChan; ok {
				fmt.Println(v)
			} else {
				return
			}
		}
	}()
	wg.Wait()
}

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func buffChannel() {
	wg := sync.WaitGroup{}
	numChan := make(chan int, 100)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if v, ok := <-numChan; ok {
				fmt.Println(v)
			} else {
				return
			}
		}
	}()
	wg.Wait()
}

// 锁机制
/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func mutexAdd() {
	var num int
	lock := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				num++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func atomicAdd() {
	var num int64
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
