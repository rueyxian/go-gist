package main

import "fmt"

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

*/

type person struct {
	name string
	age  int
}

func main() {

	// h1 := func(w ResponseWriter, r *Request){
	// 	io.WriteString(w, )
	// }

	// fooHandler := func(w ResponseWriter, r *Request) {
	// }

	// http.Handle("/foo", fooHandler)

	p1 := new(person)
	p2 := &person{}
	p3 := person{}
	var p4 person

	fmt.Printf("%T  %+v \n ", p1, p1)
	fmt.Printf("%T  %+v \n ", p2, p2)
	fmt.Printf("%T  %+v \n ", p3, p3)
	fmt.Printf("%T  %+v \n ", p4, p4)

	fmt.Println()

	// fmt.Printf()
}

/*
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
*/
