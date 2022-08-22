package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	first  = iota
	second = iota + 6  //iota value is here 1 , output will be 6
	third  = 3 << iota //iota vaule is here 2 , output is 2 bitshifted by 3
)

const (
	fourth = iota // output is 0 because iota resets on new constant block !!
	fifth         // out is 1
)

func main() {
	//primitive_data_types()
	//pointers()
	//constants()
	//iota_constant_expressions()
	//arrays()
	//slices()
	//maps()
	//structs()
	//functionWithInput("kokeloo")

	//plus, maal := functionWithReturn(3, 5)
	//fmt.Println("som:", plus, "maal", maal)
	//looping()
	//panicFunction()
	//ifstatement()
	//switchStatement()
	//readLogFile()
	//readLogFileWithArguments()
	//myWebservice()
	//myWebservice2() //webservice with query parameters
	//myWebservice3() //webservice that returns a json
	//printing_in_format()
	//read_buffer()
	//using_flags()
	//eading_input()
	//bufioScanner()
	//bufioScannerFile()

	args := os.Args[1:]                       //filename is argemunt  [1:] to exlcude arg 1 who is the filename
	if len(args) == 0 || args[0] == "/help" { //check is ther is argement, and it is [0] because we exclude the program name above
		fmt.Println("Usage: main.go <input file>")
	} else {
		//do stuff
	}

}

func primitive_data_types() {
	//how to declare variables in golang

	//classic way
	var i int
	i = 42
	fmt.Println("variable i =", i)

	//declaration and initialization on the same line
	var f float32 = 3.14
	fmt.Println("variable f =", f)

	//the go way
	g := "Pascal"
	fmt.Println("variable g =", g)

	//boolean datatype
	b := true
	fmt.Println("vairbale b =", b)

	//complex data type
	/*
		A complex number is a basic data type of GoLang. It is a combination of the real and imaginary part,
		where both parts of the complex number will be float type in the go program.
		Complex numbers are further categorized into two parts in go language
	*/
	c := complex(3, 4)
	fmt.Println("the complex variable contains :", c)
	r, im := real(c), imag(c)
	fmt.Println("the real part is", r, "the imaginary part is", im)

}

func pointers() {

	//the long way
	var firstname *string = new(string)

	fmt.Println("adress pointer firstname :", firstname)

	*firstname = "Pascal"
	fmt.Println("adress pointer firstname after assigning a value :", firstname)
	fmt.Println("the value of the pointer firstname is", *firstname)

	//declaring a pointer after you declared a variable

	lastname := "Ingelrest"

	fmt.Println("the variable lastname contains :", lastname)
	ptr := &lastname
	fmt.Println("the adress of prt is", ptr, "and the value is", *ptr)

}

func constants() {

	//when you declare a constant you have to initialize it on the same line
	const pi = 3.1415
	fmt.Println("pi =", pi)

	const cr = 3
	println(cr + 3)
	println(cr + 3.4) // the will work in go

	/*
		but is you inplicite like const cr int = 3 then it will not work and you
		than have to convert it
	*/
	const crr int = 3
	println(crr + 3)
	println(float32(crr) + 3.5)
}

func iota_constant_expressions() {
	println(first, second, third)
	println(fourth, fifth)
}

func arrays() {
	//arrays are fixed sized

	//the long way
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println(arr)
	fmt.Println(arr[1])

	//the short way
	arr2 := [3]int{40, 50, 60}
	fmt.Println(arr2)

}

func slices() {
	//kind like an array but not fixed size

	//creating an slice from an existing array, could be using
	//in case you want to work with an existing or importerd array
	arr3 := [3]int{100, 200, 300}

	slice := arr3[:]

	fmt.Print(arr3, slice)

	/*
		if you change the value of arr3[1] to 42 and slice[2] to 27
		both arr3 qnd slice contains 100,42,27 because slices acting like a pointer
	*/

	//creating a slice from scratch
	sl := []int{1000, 2000, 3000}

	fmt.Println(sl)

	sl = append(sl, 4000, 5000)
	fmt.Println(sl)

	sl2 := sl[1:] //sl2 constains 2000 3000 4000
	fmt.Println(sl2)
	sl3 := sl[:2] //sl3 contains 1000 2000  so till en NOT including index 2
	fmt.Println(sl3)
	sl4 := sl[1:3] //contains from index 1 to AND NOT including index 3, 2000 3000
	fmt.Println(sl4)
}

func maps() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m["foo"])

	m["fox"] = 54
	m["foy"] = 154
	fmt.Println(m)
	delete(m, "fox")
	fmt.Print(m)

}

func structs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u)
	u.ID = 0
	u.FirstName = "Pascal"
	u.LastName = "Ingelrest"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	//create with initialization
	u2 := user{ID: 1,
		FirstName: "Mieke",
		LastName:  "Van Kerkchove",
	}
	fmt.Println(u2)
}

func functionWithInput(data string) {
	fmt.Println("data input :", data)
}

func functionWithReturn(n1, n2 int) (add int, multi int) {
	return (n1 + n2), (n1 * n2)
}

func looping() {

	var i int

	//loop till condition
	for i < 5 {
		println(i) //println and not fmt.Println, this is more used for debugging
		i++
		/*  break out early of the loop
		if i == 3 {
			break
		}
		use continue in stead break to terminate a SINGLE interation
		*/
	}

	//loop till condition with post clause
	// in this example t is only kwon in the loop not outside it
	// if you the value further t outside the loop you need
	// to declare outside the loop as var t = 0
	// and then loop as 'for , t < 10, t++
	for t := 0; t < 10; t++ {
		println(t) //prints out 0 to 9
	}

	//infite loops

	i = 100
	for {
		if i == 105 {
			break
		}
		println(i)
		i++
	}

	//loop over collections

	slice := []string{"Pascal", "Mieke", "Erick"}

	for i, v := range slice {
		println(i, v)
	}

	//example for a map

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for k, v := range wellKnownPorts {
		println(k, v)
	}

	//what if you need only the k
	for k := range wellKnownPorts {
		println(k)
	}
	//what if you need only the v
	for _, v := range wellKnownPorts {
		println(v)
	}
	//this could be important because you get an compile error
	//if you don't use a variable !!!!
}

func panicFunction() {
	println("Starting the webserver")
	//do some things

	//panic("something bad happend")

	println("Web server started")

}

func ifstatement() {
	type User struct {
		ID        int
		FirstName string
		LastName  string
	}
	u1 := User{
		ID:        1,
		FirstName: "Pascal",
		LastName:  "Ingelrest",
	}
	u2 := User{
		ID:        2,
		FirstName: "Mieke",
		LastName:  "Van Kerkchove",
	}

	if u1.ID == u2.ID {
		println("Same User!!")
	} else {
		println("Different users!")
	}

	//you can also net with elsif !!!!!
}

func switchStatement() {
	type HTTPRequest struct {
		Method string
	}

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get request")
		//fallthrough
	case "DELETE":
		println("Delete request")
		//fallthrough
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}

	//go doesn,t expliciet break, so after a case matches it breaks
	//you can add fallthrough to undo that
}

func readLogFile() {
	f, err := os.Open("app.log")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}

func readLogFileWithArguments() {

	//access the arguments : go run programname -level "tekst" or -path "path" or both....
	path := flag.String("path", "app.log", "the path to the logfile")
	level := flag.String("level", "ERROR", "Log level to search")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

func myWebservice() {
	//register the handler that we are going the use
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//start up de webserser
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func myWebservice2() {
	//register the handler that we are going the use
	//reponsewriter to write to te website, http.request to get data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"] //there could be more than one names, so its a collection
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		w.Write([]byte("Hello " + name))
		//to execute on the browser surf to localhost:3000/?name=Pascal
	})

	//start up de webserser
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func myWebservice3() {
	//register the handler that we are going the use
	//reponsewriter to write to te website, http.request to get data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"] //there could be more than one names, so its a collection
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	//start up de webserser
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func printing_in_format() {
	os := runtime.GOOS

	fmt.Printf("opertion system = %v\n", os)

	n1 := 10
	n2 := 100
	n3 := 1000

	fmt.Printf("Nummer 1 : %4d\n", n1)
	fmt.Printf("Nummer 2 : %4d\n", n2)
	fmt.Printf("Nummer 3 : %4d\n", n3)

}

func read_buffer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is you name ?")
	fname, _ := reader.ReadString('\n')

	fmt.Printf("Your name is %s", fname)

}

func using_flags() {
	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "IA64":
		fmt.Println("Remember IA64")
	}

	fmt.Println("Process Complete")

	//run as main.go and x86 is printed out because is default
	// run as main.go -arch AMD64 or main.go -arch IA64

}

func reading_input() {
	var name string
	fmt.Println("What is your name")
	fmt.Scanf("%s", &name)
	fmt.Printf("My name is %s", name)

	//read thge documentation because scanf can do a lot

}

func bufioScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("Quitting.....")
			os.Exit(3)
		} else {
			fmt.Println("You typed " + scanner.Text())
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func bufioScannerFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
