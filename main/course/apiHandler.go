package course

import "net/http"

func apis() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
}
