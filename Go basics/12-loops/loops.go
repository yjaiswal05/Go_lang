package main

import "fmt"

func main() {
	days := []string{"Sun", "Mon", "Tues", "Wed", "Thu", "Fri", "Sat"}
	j := 1
	k := 0

	for i := 0; i < len(days); i++ {	//normal for loop
		fmt.Println(days[i])
	}

	for i := range days{				//for loop with range
		fmt.Println(days[i])
	}

	for index, days := range days{		//for each loop
		fmt.Printf("index is %v and value is %v\n", index,days)
	}

	for j < 10{							//while loop but written in for
		fmt.Println(j)
		j++
	}

	for k < 10{							//use of goto statement
		fmt.Println(k)
		if k==7 {
			goto alpha
		}
		k++
	}
	alpha : fmt.Println("Hello there")	//if this goto statement is written at the top then whatever is below it, it starts to run again

}
