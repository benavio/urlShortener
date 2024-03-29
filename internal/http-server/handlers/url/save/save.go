package save

import (
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias, omitempty"`
}
type Response struct {
	Status string `json: "status"`
	Error  string `json: "error, omitempty"`
	Alias  string `json: "alias, omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
