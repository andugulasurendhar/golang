package main

import (
	"fmt"
)

func main(){
var i int
	for{
		fmt.Println(`
*****************************************************		
		Please select an option
		1.call GoRoutinesWaitGroup print numbers
		2.call GoRoutinesMutex Transaction example
		3. Linked List
		0. Exit
*****************************************************`)			
		_,err := fmt.Scan(&i)
		if(err != nil){
			fmt.Println("Invalid Input", err)
			continue
		}
		if(i==0){
			break
		}
		switch i{
		case 1:GoRoutinesWaitGroup()
		case 2:GoRoutinesMutex()
		case 3:LinkedListfunc()
		}
	}
}