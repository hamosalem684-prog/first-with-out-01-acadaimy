package main
import "fmt"



func main(){
	fmt.Println("______________________\n")
	fmt.Println("THE PLAYERS")
	players([]string{"mohamed salem","abdo alngar","mahmoud adel","ahmed kamel","mostafa sobhy"},[]string{"omar","bedaye","alsayed","kaled","ali"},number,number2)
	match(4)
	fmt.Printf("the match at [%d]pm\n",7)
	fmt.Println("______________________\n")
	
}
func players(a []string,i []string , f func(string),r func(string)){
	for _, s := range a {
		f(s)
	}
	for _, t := range i {
		r(t)
	}
}

func number(a string){
	fmt.Printf("the player [%v] is ready\n",a)
	fmt.Println("______________________\n")
}
func number2(a string){
	fmt.Printf("the player [%v] is ready\n",a)
	fmt.Println("______________________\n")
}


func match(s int){
	switch s {
	case 1 :
		fmt.Println("the match on sunday")
	case 2 :
		fmt.Println("monday")
	case 3 :
		fmt.Println("tousday")
	case 4 :
		fmt.Println("the match on wednesday")
	case 5 :
		fmt.Println("thuorsday")
	case 6 :
		fmt.Println("friday")
	case 7 :
		fmt.Println("suterday")
		default :
		fmt.Println("there is no match today")
	}
}

######################################################################

func main(){
	p1 := []players{
		{"ahmed","cool",20,"completly ready"}, // i use for to itrate over them first by using  
		{"abdo","vary good",22,"ceirtnly ready"},
		{"mohamed","good",18,"ready for sure"},
		{"mostafa sophy","Gadamn good",19,"fucing ready"},
		{"mahmoud adel","good",20,"ok ready"},
	}
	m1 := thematch{
		time: 9,
		day: 4,
		prize: 3,
	}

	for _, a := range p1 {
		fmt.Printf("the player %v is %v and his Health is [%v]\n" ,a.Name,a.Ready,a.Health)
	}

    


	fmt.Print("\n")
	fmt.Println("the match and the time")
	fmt.Println("_______________________\n")
	fmt.Printf("the match at [%d]pm\n",m1.time)
	fmt.Printf("the match on [%v]\n",check(m1.day)) 
	fmt.Printf("the winner takes [%v]\n",gift(m1.prize))
}

// the players 
type players struct {
	Name string
	Health string
	Age int
	Ready string
}

type thematch struct {
	time int
	day int
	prize int
}

func check(s int)string{
	switch s {
	case 1 :
		return "sunday"
	case 2 :
		return "monday"
	case 3 :
		return "touthday"
	case 4 :
		return "wendesday"
	case 5 :
		return "thuorsday"
	case 6 :
		return "friday"
	case 7 :
		return "saturday"
	}
	return "unknown day"
}

func gift(s int)string{
	switch  {
	case s < 3 :
		return "ball"
	case s == 3 :
		return "monay"
	case s > 3 :
		return "women"
		default :
		return "snakers" 
	}
	return "no prize"
}

##################################################################

my naturel way

func Capitalize(s string)string{
	var thebag string // HAMO salem hamo salem 
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			thebag += string(a + 32)
		}else {
			thebag += string(a)
		}

	} 
	var thebag2 string
	for i := 0 ; i < len(thebag) ; i++{
		if (i == 0 || thebag[i - 1] == ' ')&&
		   (thebag[i] >= 'a' && thebag[i] <= 'z'){
			thebag2 += string(thebag[i]-32)
		}else {
			thebag2 += string(thebag[i])
		}
	}
	return thebag2
}


#######################################################################

using arreys

func main(){
	names := [2][]string{
		{"ahmed","omar","salem"},
		{"eslam","kamel","tamer"},
	}
	names[0] = append(names[0],"gana")
	for a := 0 ; a <len(names[0]) ; a++{
		fmt.Printf("this is the number [%d] and the nams is [%v]\n",a,names[0][a])
	} 
	fmt.Println("\n")
	names1 := []string{"ahmed","omar","salem","hamo"}
	choose := names1[len(names1)-1]
	fmt.Println(choose)

}


#####################################################################

suppose i can't see a ceirten number and i wanna find it once it's exist so // s = true
func main(){
	numbers := []int{1,2,3,4,5,6,7,8,9,11,22,33,44,55,66,77,665,675,23,56,45,76,78,98,23,12}
	s := false // can't see something 
	for _, a := range numbers {
		if a == 10 {
			s = true 
		}
	}
	if s {
		fmt.Println("exist")
	}else {
		fmt.Println("no exist")
	}
}


##########################################################################################


using make with []bool

func numbers(s []int) []bool{
		box := make([]bool, len(s)) // [false , false , false , false , false , false]
	for  o , a := range s {
			if a < 2 {
					box[o] = false // handling the first number 
			continue
		}
		box[o] = true // does that mean [all of them now are true exept 1 when you find something like composite number make it false]
		for i := 2 ; i * i <= a ; i++{
				if a % i == 0 {
						box[o] = false
					}
					// checking if the num
			// ber prime or composite 
		}
	}
	return box
}


###########################################################


func main(){
	fmt.Println(numbers([]int{1,2,3,4,5,6}))
}

func numbers(s []int) []bool {
	var t []bool
	for _, a := range s {
		if a < 2 {
			t = append(t,false)
			continue
		}
		t = append(t,true)
		for i := 2 ; i * i <= a ; i++{
			if a % i == 0 {
				t[len(t)-1] = false 
			}
		}
	}
	return t
}

#############################################################


func main(){
	fmt.Println(numbers([]int{1,2,3,4,5,6},procces))
}
func  numbers(s []int,f func(int) bool)[]bool {
	box := make([]bool, len(s))
	for i , a := range s {
		box[i] = f(a)
	}
	return box
}

func procces(s int)bool{
	if s < 2 {
		return false
	}

	for a := 2 ; a * a <= s ; a++ {
		if s % a == 0 {
			return false
		}
	}
	return true 
}

##########################################################  using flags

func main(){
	// starting with false means i assume this is wrong and i wanna find it  in another way i soppuse i don't hava this or see anything 
	list_of_names := []string{"ahmed","hamo","omar","mahmoud"}
	find := false 

	for _, a := range list_of_names {
		if a == "omar" {
			find = true 
		}
	}
	if find {
		fmt.Println("exist")
	}else {
		fmt.Println("not exist")
	}
}

#################################################################


func main(){
	what := 5 // 4 it's going to be false 
	check := true // here before i even start i assume that what i have is true and i wanna make sure
//                   i suppose it's true now and i'll find that out	
	for i := 2 ; i * i <= what ; i++{
		if what % i == 0 {
			check = false
		}
	}

	fmt.Println(check)
}