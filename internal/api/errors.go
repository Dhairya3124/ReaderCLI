package api

import "net/http"

func (app *Application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}
