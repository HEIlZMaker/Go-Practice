package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"unicode"
)

func main() {
    //firstname , lastname, email, phone number, amountoftickets into matrix
		var firstname, lastname, email, phonenumber string
    tickets := 0
  for i := 1; i <= 50; i = i + tickets {
    for {
    	fmt.Scan(&firstname)
    	if len(firstname) > 2 {
				break
    }
					 
  }
       
  for {
		fmt.Println("Type your lastname")
  	fmt.Scan(&lastname)
		if len(lastname) > 2 {
			break
   }
  }
       
	tlds := TLD()
	validTLD := true
	containsAt := true
  for {
		fmt.Println("Enter your email")
    fmt.Scan(&email)
	  containsAt = strings.Contains(email, "@")
		for i := range tlds {
			if strings.Contains(email, tlds[i]) {
				break
			} 
			if i == len(tlds) {
				
			}
			}
		if containsAt && validTLD {
			break
		}
					
		}
  }
validNumber := true
 for {
		fmt.Println("Enter your phonenumber")
		fmt.Scan(&phonenumber)
		for _, i := range phonenumber {
			if !unicode.IsNumber(i) {
				validNumber = false
			}
		}
		if validNumber {
			break
		}
	}
	validTicketcount := true
	for {
		fmt.Println("Enter how many tickets you want to order")
		fmt.Scan(&tickets)
		for i, unicodeIndex := range tickets {
			if unicode.IsLetter(unicodeIndex) {
				validTicketcount = false			
			}
			if 
		}
	}
}



func TLD() []string {
	file, err := os.Open("TLD")
  checkErr(err)
    defer file.Close() 


    scanner := bufio.NewScanner(file)

		var tlds []string
    for scanner.Scan() {
      tlds = append(tlds, scanner.Text())
    }
		return tlds
	}
	


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}