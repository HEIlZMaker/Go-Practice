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
       for range 3 {
           fmt.Scan(&firstname)
           if len(firstname) > 2 {
						break
           }
       }
       fmt.Println("Type your lastname")
       for range 3 {
           fmt.Scan(&lastname)
					 if len(lastname) > 2 {
						break
           }
       }
       fmt.Println("Enter your email")
			tlds := TLD()
			var validTLD bool
       	for range 3 {
           fmt.Scan(&email)
				  containsAt := strings.Contains(email, "@")
					for i := range tlds { 
					if strings.Contains(email, tlds[i]) {
						validTLD = true
						break
					} else {
						validTLD = false
					}
					if containsAt && validTLD {
						break
					}
						
				}
       }
      fmt.Println("Enter your phonenumber")
      for range 3 {
        fmt.Scan(&phonenumber)
			for i := range phonenumber {
				if !unicode.IsLetter(phonenumber) {
					break
				}
			}
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