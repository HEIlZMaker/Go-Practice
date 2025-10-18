package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"

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
				tlds = TLD()
        for range 3 {
            fmt.Scan(&email)
					validMail :=
					if {
						strings.Contains(email, "@")
						
					}

        }
        fmt.Println("Enter your phonenumber")
        for range 3 {
            phonenumber := fmt.scan()
   
        i = i + tickets
    }

    booking := [5]string{firstname, lastname, email, phonenumber, tickets}
    i := 0
    people := []string{}
    people = append(people, booking[i])
}

}
func TLD() []string {
	file, err := os.Open("TLD")
    if err != nil {
        log.Fatalf("failed to open file: %s", err)
    }
    defer file.Close() 


    scanner := bufio.NewScanner(file)

		var tlds []string
    for scanner.Scan() {
      tlds = append(tlds, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("error reading file: %s", err)
    }
		return tlds
	}
	


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}