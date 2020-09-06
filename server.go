/* package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Todo Structure
type Todo struct {
	userID    int
	id        int
	title     string
	completed bool
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Todo</h1>")

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Weather</h1><ul>")

	})

	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/* func h() {


	fmt.Printf("Starting server at port 8080\n")
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
} */
