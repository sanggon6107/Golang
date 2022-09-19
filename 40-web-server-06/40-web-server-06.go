package main

import "net/http"

// HTML can include files such as iamges and music. It doesn't include files as themselves, but as URL indicating the path of them.

func main() {
	// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)

}
