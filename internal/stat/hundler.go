package stat

import (
	"apigo/configs"
	"apigo/pkg/middlewere"
	"fmt"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHundlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHundler struct {
	StatRepository *StatRepository
}

func NewStatHundler(router *http.ServeMux, deps StatHundlerDeps) {
	handler := &StatHundler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middlewere.IsAuthed(handler.GetAll(), deps.Config))
}

func (handler *StatHundler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}

		fmt.Println(by, from, to)
	}
}
