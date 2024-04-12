package main

import "snippetbox.dimitrios_papakonstantinou.com/internal/models"

type templateData struct {
  Snippet *models.Snippet
  Snippets []*models.Snippet
}
