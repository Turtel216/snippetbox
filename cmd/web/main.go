package main

import (
  "database/sql"
  "flag"
  "log"
  "net/http"
  "os"
  "snippetbox.dimitrios_papakonstantinou.com/internal/models"
  
  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  errorLog  *log.Logger
  infoLog   *log.Logger
  snippets  *models.SnippetModel
}

func main() {
  
  // You can run go run ./cmd/web -help to get the flag info e.g. "HTTP network adress" and the conversion type e.g. String
  addr := flag.String("addr", ":4000", "HTTP network adress")
  dsn := flag.String("dsn", "web:7777@/snippetbox?parseTime=true", "MySQL data source name")

  flag.Parse()

  // You can redirect the stdout and stderr streams to on disk files when starting the application like so 
  // go run ./cmd/web >>/tmp/info.log 2 >>/tmp/error.log

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }

  defer db.Close()

  app := &application{
    errorLog: errorLog,
    infoLog:  infoLog,
    snippets: &models.SnippetModel{DB: db},
  }

  srv := &http.Server {
    Addr:     *addr,
    ErrorLog:  errorLog,
    Handler:  app.routes(),
  }
  infoLog.Printf("Starting server on %d", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
