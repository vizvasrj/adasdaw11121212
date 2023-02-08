package main

import "fmt"

// i have converted your complex structure to simpler structer for
// better code understandablity
func absListner(cnp chan func() string) {
	cnp <- func() string {
		return "Hea !"
	}
}

func _main() {
	// i have returning string here so output can be viewed
	cnp := make(chan func() string, 10)
	// ^ at this code we are creating buffer channel with limit of 10
	// so we can cnp can hold upto 10

	// at this stage we are creating 4 threads
	// to run job concurently with 4 workers
	for i := 0; i < 4; i++ {
		go absListner(cnp)
	}

	// at this stage we are inputing to cnp channel
	cnp <- func() string {
		fmt.Println("HERE1")
		return fmt.Sprintf("What?")
	}

	// at this stage we are recieving from channels
	// and outputing it
	for x := 1; x <= 2; x++ {
		select {
		case rfx := <-cnp:
			msg := rfx()
			fmt.Println("msg", msg)
			fmt.Println("Recieved")
		}
	}
	fmt.Println("Hello")
}

// application and use case of this when we want to limit
// the number of thread to run concurrently
// so that it can prevent it from heavy resource utilization
