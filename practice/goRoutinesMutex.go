package main

import (
	"fmt"
	"sync"
	"time"
)

var bal int = 0

func GoRoutinesMutex(){
	var mx sync.RWMutex
	var n sync.WaitGroup

	trans := []int{
		0:100,
		1:200,
		2:-500,
		3:-400,
		4:0,
		5:1000,
		6:-1000,
		7:0,
		8:200,
		9:600,
		10:0,
	}

	
	for len(trans)>0{
		n.Add(1)
		go doTransaction(&n, &mx,&trans)
	}

	n.Wait()
	
}

func doTransaction(n *sync.WaitGroup, mu *sync.RWMutex ,trans *[]int){
	time.Sleep(time.Second) 
	defer n.Done()
	defer mu.Unlock()
	mu.Lock()//we cant use only read lock because the trans is getting updated even when checking the balance
	if len(*trans) == 0 {
		return // Nothing to process
	}
	if(len(*trans)>0) {
		if((*trans)[0] ==0){
			getbalance(trans)
		} else if(len(*trans)>0) {
			calculateTransaction(trans)	
		} 
	} else{
			fmt.Println("inSufficient length")
	}
}

func calculateTransaction(trans *[]int) {
	val := (*trans)[0]
	if bal+val > 0 {
		bal = bal + val
		
		(*trans) = (*trans)[1:]
		fmt.Println(val," Transaction done balance : ", bal)
	} else {
		(*trans) = (*trans)[1:]
		(*trans) = append((*trans), val)
		fmt.Println("inSufficient balance")
	}
}

func getbalance(trans *[]int) {
	(*trans) = (*trans)[1:]
	fmt.Println("balance : ", bal)
}