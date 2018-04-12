package main

import ("fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say(s string) {
	// defer statements work in LIFO order 
	// defers running of this statement until other statements are run
	defer cleanup()
	for i:=0; i<3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
	// defer fmt.Println("Done!")
	// defer fmt.Println("Are we done?")
	// fmt.Println("Doing some stuff")
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()


}
