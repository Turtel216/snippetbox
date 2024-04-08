package main

import (
  "flag"
  "log"
  "net/http"
  "os"
)

func main() {
  
  // You can run go run ./cmd/web -help to get the flag info e.g. "HTTP network adress" and the conversion type e.g. String
  addr := flag.String("addr", ":4000", "HTTP network adress")

  flag.Parse()

  // You can redirect the stdout and stderr streams to on disk files when starting the application like so 
  // go run ./cmd/web >>/tmp/info.log 2 >>/tmp/error.log

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

  mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  srv := &http.Server {
    Addr:     *addr,
    ErrorLog:  errorLog,
    Handler:  mux,
  }
  infoLog.Printf("Starting server on %d", *addr)
  err := srv.ListenAndServe()
  errorLog.Fatal(err)
}
