package main

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutinesWaitGroup(){
	fmt.Print("Hello World")
	var mu sync.WaitGroup
	
	for i:=0; i<5;i++{
		mu.Add(1)
		go createGoRoutine(&mu,i+1)
	}

	mu.Wait()
	
}

func createGoRoutine(mu *sync.WaitGroup, n int){
	if(n<3){
		time.Sleep(2*time.Second)
	}
	for i:=0;i<20;i++{
		fmt.Println(i,n)
	}
	defer mu.Done()
}