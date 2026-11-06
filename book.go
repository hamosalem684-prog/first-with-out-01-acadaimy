package main 
import "fmt"


func main(){
	fmt.Printf("%#v\n",str("Hello how are you?"))
}

func str(s string)[]string{
	var t []string 
	var r string
	for _, i := range s{ // hello how are you?
		if i != ' ' && i != '\n' && i != '\t' {
			r += string(i) // untill here it takes the whole word and print it like this "hellohowareyou"
		}else { // when it hits a space or something from these it turns to else 
			if r != "" { 
				t = append(t , r)
				r = ""
			}
		}
	}
	if r != "" { // the left word which is you? . you can also add space after it in func main
		t = append(t,r)
	}
	return t
}
########################################
func main(){
	day := 111
	switch  day { // without this day what happens
	case 1 :
		fmt.Println("sunday")
	case 2 :
		fmt.Println("monday")
	case 3 :
		fmt.Println("tuesday")
	case 4 :
		fmt.Println("wesdensday")
	case 5 :
		fmt.Println("thursday")
	case 6 :
		fmt.Println("friday")
	case 7 :
		fmt.Println("saturday")
		default :
		fmt.Println("invalid")
	}
}

#######################################################

func main(){
	color := "red"
	switch color{
	case "red":
		fmt.Println("stop")
	case "green" :
		fmt.Println("go")
	case "blue" :
		fmt.Println("sky")
		default :
		fmt.Println("not exists")
	}
}

################################################################

func main(){
	command := "important"
	switch command {
	case "not important" :
		fmt.Println("delet")
	case "important" :
		fmt.Println("save")
		default :
		fmt.Println("remove")
	}
}

#################################################################

func main(){ // which one is this type switch or switch expresstion
	number := 360 // but it seem a litter weired for me I am the one who puts the number and before printing so why and for what // i mean when will i be using it and for what ?
	switch {
	case number > 500 :
		fmt.Println("very big")
	case number < 400 :
		fmt.Println("big")
	case number < 300 :
		fmt.Println("small")
		default :
		fmt.Println("isn't exist")
	}
}

##############################################

func main(){
	grades := 101
	switch {
	case grades < 0 || grades > 100 :
		fmt.Println("invalid grades")
	case grades >= 90 && grades <= 100 :
		fmt.Println("Exelent")
	case grades >= 80 && grades <= 90 :
		fmt.Println("cool")
	case grades >= 70 && grades <= 80 :
		fmt.Println("good")
	case grades < 70 :
		fmt.Println("failed")
		default :
		fmt.Println("nice")
	}
}

########################################################

func main(){ // the bottom line about switch is that switch can be replaced as conditions like if and else if 
	sw("heilthy") // also when we wanna compare somthing aganist many passible values
	sw("strong") // easeir to read and write than if and else if 
	sw("week") // this's what i know now 
}

func sw(s string){
	switch s {
	case "strong" :
		fmt.Println("acceptble")
	case "week" :
		fmt.Println("go hamo")
	case "big" :
		fmt.Println("welcame")
		default :
		fmt.Println("not acceptble")
	}
}

#################################################


func num(a,b int,div *int , mod *int)(int,int){
	// *div , *mod = a / b , a % b
	*div = a / b
	*mod = a % b
	return *div,*mod
}


func main(){
	a , b := 13 , 2 
	var (
		one int
		two int
	)
	num(a,b,&one,&two)
	f := one 
	h := two
	fmt.Println(one,two)
}

###############################################


the one before the last exercise in qiest 7


func main(){
	fmt.Printf("%#v\n",str("hello how are you"))
}

func str(s string)[]string{
	var t []string
	var a string
	for _, i := range s {
		if i == ' ' || i == '\n' || i == '\t' { // hello how are you 
			if a != "" { // hello is not "" right ? yes so it appends hello 
				t = append(t,string(a))
				a = "" // and than this starts again with anew one to hold the word how 
			}
		}else{ // how this gives me something ? as long as i == ' ' is false h is not a space so ? how or wait because it's not a space it prints it ?
			a += string(i) // hello how are you and in fact like this h he hel hell helle and so on 
		} // but how it gets back to if a != "" and so on how ?
	}
	return t

}
