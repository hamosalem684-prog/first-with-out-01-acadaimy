
//####################################################   makeing them in slices

// func main(){
// 	fmt.Printf("%#v\n",str("ahmed  mohamed abdo ali hamos")) 
// }

// func str(s string)[]string{
// 	var t []string
// 	var i string
// 	for _, a := range s{
// 		if a != ' ' && a != '\n' && a != '\t' { // i don't think i can get what a has unless i take what it has and put them in another varaible
// 			i += string(a)
// 		}else{
// 			if i != "" {
// 				t = append(t,i)
// 				i = ""
// 			} 
// 		}
// 	}
// 	if i != ""{
// 		t = append(t,i)
// 	}
// 	return t
// }

//############################################################# newline though the difrence is slices take a look important

// func main(){
// 	fmt.Println(str([]string{"hello","how","are","you"}))
// }

// func str(s []string)string{
// 	var r string
//     for i , a := range s { // i has every position to every element in the slice but len(s) has the 3. 0 to the first word and 1 to the socend and so on  
// 		r += string(a)
// 		if i < len(s)-1{ // len(s) looks 4 things but the index are 3 
// 			r += "\n"
// 		}
// 	}
// 	return r
// }

//####################################################### // to put newline after each word 

// func main(){
// 	fmt.Println(str("how are you"))
// }

// func str(s string)string{
// 	var i string
// 	v := true 
// 	for o , a := range s{
// 		if a == ' ' {
// 			v = false
// 		}else {
// 			if !v {//alse
// 				if o < len(s){
// 					i += "\n"
// 					v = true // !!!!!!
// 				}
// 			}
// 		}
// 		i += string(a)
// 	}
// 	return i 
// }


//############################################################################### using passed function

// func main(){
//     slices(
// 		[]string{"ahmed","ali","omar","ramadan"},
// 		change,
// 	)
// }

// func slices(names []string,f func(string,int)){ 
// 	for i := 0 ; i < len(names) ; i++{
// 		f(names[i],i)
// 	}
// }

// func change(s string,a int){
// 	fmt.Printf("the %dth name is %v\n",a,s) // takse the index 0 that has also on the other side a name
// }

// func change(s string,a int){ // using switch exeprision
// 	switch a {
// 	case 0 :
// 		fmt.Println("the first name",s) // it's going to print all of these
// 	case 1 :
// 		fmt.Println("the socend name",s)
// 	case 2 :
// 		fmt.Println("the third name",s)
// 	case 3 :
// 		fmt.Println("the fourth name",s)
// 	}
// }

//########################################################## with numbers 

// func main(){
// 	slice(
// 		[]int{20,14,30},
// 		change,
// 	)
// }

// func slice(s []int,f func(int,int)){
// 	for i , a := range s{
// 		f(i,a)
// 	}
// }

// func change(s,a int){
// 	fmt.Printf("this is %dth number %v\n",s+1,a)
// }


//#################################################################

// func change(s int){ // change f func(int) just one is enough up here 
// 	switch s { // this takse more time but it's usfell
// 	case 20 :
// 		fmt.Println("the first number",s)
// 	case 14 :
// 		fmt.Println("the second",s)
// 	case 30 :
// 		fmt.Println("the third",s)
// 	}
// }
//########################################################################

// type test struct {
// 	sqrt int
// 	fibnacci int
// }

// func main(){
// 	type1 := test {
// 		sqrt: 36,
// 		fibnacci:10,
// 	}
// 	fmt.Printf(
// 		"the number is %d and this is the sqrt [%v]\n",
// 		type1.sqrt,
// 		recursion2(type1.sqrt,test3),
// 	)
// 	fmt.Printf(
// 		"fibnacci numbers from 0 to %v are %v\n",
// 		type1.fibnacci,
// 		recursion(type1.fibnacci,test2),
// 	)
// 	fmt.Println(number1(20,mines5)) // i wanna say number 1 is false and 2 is true i wanna show every number
// }

// //############################################## sqrt with for now and leter learn how to do it with recursion
// func recursion2(s int,f func(int)int)int{
// 	return f(s)
// }

// func test3(s int)int{ // i have to try it with recursion
// 	for a := 2 ; a * a <= s ; a++{
// 		if a * a == s {
// 			return a
// 		}
// 	}
// 	return 0
// }

// //############################################# fibnacci with [] slices next time should be with no slices

// func recursion(s int,f func(int)int)[]int{
// 	box:= make([]int,s)
// 	for a := 0 ; a < s ; a ++{
// 		box[a] = f(a)
// 	}
// 	return box
// }

// func test2(s int)int{
// 	if s == 0 {
// 		return 0
// 	}

// 	if s == 1 {
// 		return 1
// 	}
// 	return test2(s - 1) + test2(s - 2)
// }

// //################################################## numbers from 0 to 19 chicking whether prime or composite the more imporntly is the way

// func number1(s int,f func(int) bool )[]bool{
// 	var t []bool
// 	for a := 0 ; a < s ; a++ {
// 		t = append(t,f(a))
// 	}
// 	return t
// }

// func mines5(s int)bool{
// 	if s < 2 {
// 		return false
// 	}
// 	for a := 2 ; a < s ; a++{
// 		if s % a == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

