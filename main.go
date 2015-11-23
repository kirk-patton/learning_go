package main

import "fmt"
import "os"
import "bufio"
import go_util "github.com/kirk-patton/go_util" //Syntax to rename the imported pkg.

//import "io"
//import "io/ioutil"

func main() {
	fmt.Println("This is main")

	//This is a commet

	/*
		          This is a multi line comment
			  See, another line
	*/

	//Parsing arguments
	args := os.Args[1:] //Read all arguments except the program name
	//Also a foreach style loop
	for a := range args {
		fmt.Printf("Arg: %s \n", args[a])
	}

	//A standard for loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("Example for loop %v \n", i)
	}

	//While loop
	loop_var := 5
	for loop_var > 0 {
		fmt.Println("Decrement the loop var")
		loop_var--
	}

	//Reading a file into memory
	//data, err := ioutil.ReadFile("/etc/passwd")
	//go_util.check(err)

	//fmt.Print(string(data)) //Dump the file to stdout

	//opening a file handle and itterating it
	fh, err := os.Open("/etc/passwd")
	go_util.Check(err)
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	//open file for append
	w_fh, err := os.OpenFile("/tmp/go_write", os.O_WRONLY, os.ModeAppend)
	go_util.Check(err)
	writer := bufio.NewWriter(w_fh)

	//This will copy the text
	for scanner.Scan() {
		text := scanner.Text() + "\n"
		writer.WriteString(text)
	}
	writer.Flush()

	//Freak out
	freak_out := false
	if freak_out {
		panic("Oh shit!!!")
	}

	//Creating a list aka slice
	strings := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(strings[:1]) //Slice up to. Start a element 0. Same as [0]
	fmt.Println(strings[0])  //Exactly element [0]
	fmt.Println(strings[0:]) //Slice from 0 through end of slice
	fmt.Println(strings[:])  //Also Slice from 0 through end of slice

	//Appending to the slice
	strings = append(strings, "Z", "Y", "X")
	fmt.Println(strings)

	//Maps/Hashes/Dictionaries
	var m map[string]int
	m = make(map[string]int)

	//Add key value
	m["foo"] = 1
	fmt.Printf("Map: m[foo] = %v\n", m["foo"])
	fmt.Printf("Map: NumKeys %v\n", len(m))
	value, tf := m["foo"]
	fmt.Printf("Map KeyExists: value: %v exists: %v\n", value, tf)

	//Delete key
	fmt.Println("Map Delete Key\n")
	delete(m, "foo")

	//Get the number of key in a map
	fmt.Printf("Map: NumKeys %v\n", len(m))

	//Test key exists
	value, tf = m["foo"]
	fmt.Printf("Map KeyExists: value: %v exists: %v\n", value, tf)

	//Iterate over a map
	m = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for k, v := range m {
		fmt.Printf("Map Interation Key: %v Value:%v\n", k, v)
	}
}
