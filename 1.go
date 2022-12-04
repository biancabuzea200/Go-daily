package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func menu(menu int ) {

	if(menu == 1){
		fmt.Println("Hello. Zour in the main menu. What do you want to do?Select one of the follwoing")
		fmt.Println("1, option 1")
		fmt.Println("2, option 2")
		fmt.Println("3, option 3")
		fmt.Println("4, exit")
		fmt.Print("Enter Num,  ")
	}else if(menu == 2){ 
		fmt.Println("Hello.Zour in the option 2 Menu.  What do you want to do?Select one of the follwoing")
		fmt.Println("1, menu 2.1")
		fmt.Println("2, menu 2.2")
		fmt.Println("4, exit")

	}

}

func read() (int, bool){
	reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = input[0:1]

		inputI, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("option invalid")
			return 0, false
		}

		if inputI < 1 || inputI > 4 {
			fmt.Println("choose option between 1-4")
			return 0, false
		} else {
			return inputI, true
		}

}

func main() {

	for {
		menu(1)
		inputI, ok := read()
		if !ok{
			continue
		}
		

		if inputI == 1 {
			fmt.Println("1 was chosen")
		} else if inputI == 2 {
			for {
				menu(2)

				
				inputI, ok := read()
				if !ok{
					continue
				}

				fmt.Print("You have selected ", inputI)
				if inputI == 4 {
					break
				}


			}
		} else if inputI == 3 {
			fmt.Println("3 was chosen")
		} else if inputI == 4 {
			fmt.Println("exiting")
			break
		}
	}

}
