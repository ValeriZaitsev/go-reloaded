package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main(){
	txt, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	slova := (string(txt[:]))
	slices := strings.Fields(slova)
	//spaceXY(punctuation(upLowCap(numberVersion(slices))))
	
	answr := (strings.Join(Skips(spaceXY(letters(upLowCap(numberVersion(slices))))), " "))
	
	ioutil.WriteFile("result.txt", []byte(answr), 0644)
}

func numberVersion(pizza []string) []string {
	for i := 0; i < len(pizza); i++ {
		if pizza[i] == "(hex)"{
			number, err := strconv.ParseInt(pizza[i-1], 16, 32)
			if err != nil {
				fmt.Println(err) 
			}
			pizza[i-1] = strconv.FormatInt(number, 10)
			pizza[i] = ""
	
		}else if pizza[i] == "(bin)"{
				number, err := strconv.ParseInt(pizza[i-1], 2, 32)
				if err != nil {
					fmt.Println(err) 
				}
				pizza[i-1] = strconv.FormatInt(number, 10)
				pizza[i] = ""
		}
	}
	return  pizza
}

func upLowCap(pizza []string) []string {	
	for i := 0; i < len(pizza); i++ {
		switch pizza[i] {
		case "(up)" :
			pizza[i-1] = strings.ToUpper(pizza[i-1])
			pizza[i] = ""
		case "(up," :
			pizza[i+1] = strings.Trim(pizza[i+1], ")")
			amount, err := strconv.Atoi(pizza[i+1])
				if err != nil{
					fmt.Println(err)
				}
				
				for w := 1; w <= amount; w++ {
					pizza[i-w] = strings.ToUpper(pizza[i-w])
				}
				
				pizza[i], pizza[i+1] = "",""
				
		case "(low)" :
			pizza[i-1] = strings.ToLower(pizza[i-1])
			pizza[i] = ""
		case "(low," :
			pizza[i+1] = strings.Trim(pizza[i+1], ")")
			amount, err := strconv.Atoi(pizza[i+1])
				if err != nil{
					fmt.Println(err)
				}
				
				for w := 1; w <= amount; w++ {
					pizza[i-w] = strings.ToLower(pizza[i-w])
				}
				pizza[i], pizza[i+1] = "",""
		
			case "(cap)" :
				pizza[i-1] = strings.Title(pizza[i-1])
				pizza[i] = ""
			case "(cap," :
				pizza[i+1] = strings.Trim(pizza[i+1], ")")
				amount, err := strconv.Atoi(pizza[i+1])
					if err != nil{
						fmt.Println(err)
					}
					for w := 1; w <= amount; w++ {
						pizza[i-w] = strings.Title(pizza[i-w])
					}
					pizza[i], pizza[i+1] = "",""
	}
}
return pizza
}

func letters(pizza []string) []string {
	for i := 0; i < len(pizza); i++ {
		for _, letter := range "aeiouhAEIOUH"{
			if (pizza[i] == "a" || pizza[i] == "A") && rune(pizza[i+1][0]) == letter{
				pizza[i] += "n"
			}
		}
	}
	return pizza
}
	
	func spaceXY(pizza []string) []string {
		str1 := strings.Join(pizza, " ")
		var resStr string
		char := regexp.MustCompile(`[.,\?!:; ]+`)
		text := char.Split(str1, -1)
		punct := char.FindAllString(str1, -1)
		for i := 0; i < len(punct); i++ {
			resStr += strings.TrimSpace(text[i])
			resStr += strings.TrimSpace(punct[i])
			resStr += " "
		}
		resStr += text[len(text)-1] // print last symbol
		String1 := strings.Fields(resStr)
		return String1
	}

	func Skips(pizza []string) []string { // delete spaces
		var q bool = true
		for i := 0; i < len(pizza); i++ {
			if pizza[i] == "'" {
				if q {
					pizza[i+1] = "'" + pizza[i+1]
					pizza[i] = ""
					q = false
				} else {
					pizza[i-1] = pizza[i-1] + "'"
					pizza[i] = ""
					q = true
				}
			}
		}
		return pizza
	}