package main

import (
  "errors"
  "net/http"
  "strconv"

  "snippetbox.dimitrios_papakonstantinou.com/internal/models"
  "github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

  snippets, err := app.snippets.Latest()
  if err != nil {
    app.serverError(w, err)
    return
  }

  data := app.newTemplateData()
  data.Snippets = snippets

  app.render(w, http.StatusOK, "home.html", data)
}


func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
  params := httprouter.ParamsFromContext(r.Context())

  id, err := strconv.Atoi(params.ByName("id"))

  if err != nil || id < 1 {
    app.serverError(w, err)
    return
  }

  snippet, err := app.snippets.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      app.notFound(w)
    } else {
      app.serverError(w, err)
    }
    return
  }

  data := app.newTemplateData()
  data.Snippet = snippet

  app.render(w, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request)  {
  data := app.newTemplateData(r)

  app.render(w, http.StatusOK, "create.html", data)
}
