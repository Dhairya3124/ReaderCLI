package api

import (
	"net/http"

	"github.com/Dhairya3124/ReaderCLI/internal/store"
)

type CreateArticlePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func (app *Application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload CreateArticlePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	article := store.Article{
		Title:       payload.Title,
		Description: payload.Description,
		URL:         payload.URL,
	}

	if err := app.storage.Create(ctx, &article); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, article); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
