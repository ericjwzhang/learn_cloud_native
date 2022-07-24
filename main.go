package main

import (
	"fmt"
	"time"
)

func test() {
	var myArray [6]string = [6]string{"i", "am", "stupid", "and", "weak"}
	fmt.Printf("org: %s\n", myArray)
	for index, value := range myArray {
		if index == 2 {
			fmt.Print("smart ")
		} else if index == 4 {
			fmt.Print("strong!")
		} else {
			fmt.Printf("%s ", value)
		}
	}
}

func main() {
	myChan := make(chan int, 10)
	done := make(chan bool)
	defer close(myChan)
	go func() {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("subprocess interrupt!")
				return
			case a := <-myChan:
				fmt.Printf("subprocess recive msg: %d\n", a)
			}
		}

	}()
	mainTicker := time.NewTicker(time.Second)
	cnt := 0
	for _ = range mainTicker.C {
		if cnt < 10 {
			myChan <- cnt
			cnt++
		} else {
			close(done)
		}
	}
}
