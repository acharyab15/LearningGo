package main

import ("fmt"
	"net/http")


func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `<h1>Whoa, Go is neat!</h1>
		<p> Go is fast </p>
		<p> ... and simple </p>
		`)
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Expert Blockchain")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
