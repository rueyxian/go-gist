package main

import (
	"fmt"
	"net/http"
)

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

https://golang.org/pkg/net/http/#Request
*/

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	// fmt.Printf("  type: %T \n", w)
	fmt.Printf("**************************************** \n")
	fmt.Println()
	fmt.Printf("********** http.Request ********** \n")
	fmt.Printf("          Method: %v \n", r.Method)
	fmt.Printf("             URL: %v \n", r.URL)
	fmt.Printf("           Proto: %v \n", r.Proto)
	fmt.Printf("      ProtoMajor: %v \n", r.ProtoMajor)
	fmt.Printf("      ProtoMinor: %v \n", r.ProtoMinor)
	fmt.Printf("          Header: %v \n", r.Header)
	fmt.Printf("            Body: %v \n", r.Body)
	// fmt.Printf("      GetBody: %v \n", r.GetBody)
	fmt.Printf("   ContentLength: %v \n", r.ContentLength)
	fmt.Printf("TransferEncoding: %v \n", r.TransferEncoding)
	fmt.Printf("           Close: %v \n", r.Close)
	fmt.Printf("            Host: %v \n", r.Host)
	fmt.Printf("            Form: %v \n", r.Form)
	fmt.Printf("        PostForm: %v \n", r.PostForm)
	fmt.Printf("   MultipartForm: %v \n", r.MultipartForm)
	fmt.Printf("         Trailer: %v \n", r.Trailer)
	fmt.Printf("      RequestURI: %v \n", r.RequestURI)
	fmt.Printf("             TLS: %v \n", r.TLS)
	fmt.Printf("          Cancel: %v \n", r.Cancel)
	fmt.Printf("        Response: %v \n", r.Response)

}
