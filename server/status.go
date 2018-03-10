package server

import (
	"net/http"
	"os"
	"runtime"
)

func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, Response{
		Meta: makeMeta(r, s.now()),
		Data: Data{
			"dyno":     os.Getenv("DYNO"),
			"language": "go",
			"go": map[string]interface{}{
				"version":        runtime.Version(),
				"num_cpu":        runtime.NumCPU(),
				"num_goroutines": runtime.NumGoroutine(),
				"goarch":         runtime.GOARCH,
				"goos":           runtime.GOOS,
			},
			"heroku": map[string]string{
				"app_id":           os.Getenv("HEROKU_APP_ID"),
				"app_name":         os.Getenv("HEROKU_APP_NAME"),
				"dyno_id":          os.Getenv("HEROKU_DYNO_ID"),
				"release_version":  os.Getenv("HEROKU_RELEASE_VERSION"),
				"slug_commit":      os.Getenv("HEROKU_SLUG_COMMIT"),
				"slug_description": os.Getenv("HEROKU_SLUG_DESCRIPTION"),
			},
		},
	})
}
