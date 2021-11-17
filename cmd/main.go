package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	fileName := flag.String("fname", "cmd/baby.txt", "Please specify a file to stream")
	port := flag.String("port", "3000", "Please specify a port number")

	flag.Parse()

	r.Get("/stream", func(w http.ResponseWriter, r *http.Request) {

		src, err := os.Open(*fileName)
		if err != nil {
			fmt.Printf("Fail to open file")
		}

		w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
		w.Header().Set("Transfer-encoding", "chunked")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		p := make([]byte, 5)
		enc := json.NewEncoder(w)
		for {
			n, err := src.Read(p)
			if err == io.EOF {
				fmt.Println("---End of File---")
				break
			} else if err != nil {
				fmt.Println("Error in reading strings", err)
				break
			}

			fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

			f, ok := w.(http.Flusher)
			if ok {
				f.Flush()
			} else {
				fmt.Println("streaming unsupported")
			}

			enc.Encode(p[:n])

			time.Sleep(1 * time.Second)
		}

	})

	http.ListenAndServe(":"+*port, r)

}
