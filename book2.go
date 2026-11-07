package main
import "fmt"

func main(){
	intgers([]int{1,2,3,4,5},procces)
}

func intgers(a []int , f func(int)){
	for _, s := range a {
		f(s) 
	}
}

func procces(pass int){
	fmt.Print(pass)
}


############################################

func main(){
	fmt.Println(slice([]int{1,2,3,4,5,6},procces2))
}

func slice(a []int,f func(int) bool)[]bool {
	var t []bool
	for _, s := range a{
		// result := f(s) // 1 2 3 4 5 6
		t = append(t,f(s)) // [false]
	//  t = append(t,f(2)) // [true]
	//  t = append(t,f(3)) // [true]
	//  t = append(t,f(4)) // [false]
	}
	return t
}

func procces2(pass int) bool {
	if pass < 2 {
		return false
	}
	for a := 2 ; a * a <= pass ; a++{
		if pass % a == 0 {
			return false
		}
	}
	return true 
}

using make 


func slice(a []int , f func(int) bool )[]bool {
	box := make([]bool,len(a)) 
	for i , s := range a {
		box[i] = f(s)
	}
	return box
}

func procces2(s int)bool {
	if s < 2 {
		return false
	}

	for a := 2 ; a * a <= s ; a++{
		if s % a == 0 {
			return false 
		}
	}
	return true 
}


doing the same but without f func(int)bool

func main(){
	fmt.Println(slice([]int{1,2,3,4,5,36}))
}

func slice(a []int) []bool { // false or true depending on a 
	box := make([]bool , len(a)) // [false , false , false , false , false , false]
	// [1,2,3,4,5,6]
	// 123456
	for d , i := range a {
		test := true // it's true whatever i have now 

		if i < 2 { // find this hit false
			test = false
		}else {
			for s := 2 ; s * s <= i ; s++{ // it starts here from 2 right ?
				if i % s == 0 { // if it whatever the number is less than = i apply this and see if it's composite or prime 
					test = false // but it does that one try with every number 
					break // show me in picture how it works
				}
			}
		}
		box[d] = test
	}
	return box
}


##################################

multidimansionlal

func main(){
	two := [3] [3]string { // does that mean 3 {} and 3 elements
		{"ahmed","omar","hamo"},
		{"gana","hanan","nore"},
		{"homos","kamel","mohamed"},
	}
	for a := 0 ; a < len(two); a++{ // one to the first [3]
		for c := 0 ; c < len(two[a]) ; c++{ // this to the socend
			fmt.Println(two[a][c])
		}
	}
}



func main(){
	numbers := [3][3]int{
		{1,2,3},
		{4,5,6},
		{6,7,8}, // why should i put comma here in the last one
	}
	
	for a := 0 ; a < len(numbers) ; a++{
		for s := 2 ; s < len(numbers[a]) ; s++{
			fmt.Println(numbers[a][s])
		}
	}
}



func main(){
	fmt.Println(slice([]string{"ahme5d","omar","mohamed"},procces))
}

func slice(s []string, f func(string) bool )[]bool {
	box := make([]bool , len(s))
	for i , a := range s {
		box[i] = f(a)
	}
	return box
}



func procces(s string) bool {
	for _, a := range s {
		if a >= '0' && a <= '9' {
			return false
		}
	}
	return true 
}




func main(){
	a := "mohamed"
	box := make([]string, len(a))
	for i, s := range a {
		box[i] = string(s)
	}
	fmt.Println(box)

	var t []string
	for _, r := range a {
		t = append(t,string(r))
	}
	fmt.Println(t)
}