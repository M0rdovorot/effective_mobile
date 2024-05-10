package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
	"github.com/gorilla/mux"
)

const (
	maxBytesToRead = 1024 * 1024 * 1024
)

type IValidatable interface {
	// IsValide() bool
	IsEmpty() bool
}

type IMarshable interface {
	// MarshalJSON() ([]byte, error)
	// MarshalEasyJSON(w *jwriter.Writer)
}

type Wrapper[Req IValidatable, Resp IMarshable] struct {
	fn func(ctx context.Context, req Req) (Resp, int, error)
}

func NewWrapper[Req IValidatable, Resp IMarshable](fn func(ctx context.Context, req Req) (Resp, int, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}

func (wrapper *Wrapper[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ctx := ctxusecase.AddWriter(r.Context(), w)
	ctx = ctxusecase.AddVars(ctx, mux.Vars(r))

	queryMap := map[string]string{}

	if r.URL.Query().Has("regNum") {
		queryMap["regNum"] = r.URL.Query().Get("regNum")
	}
	if r.URL.Query().Has("mark") {
		queryMap["mark"] = r.URL.Query().Get("mark")
	}
	if r.URL.Query().Has("model") {
		queryMap["model"] = r.URL.Query().Get("model")
	}
	if r.URL.Query().Has("year") {
		queryMap["year"] = r.URL.Query().Get("year")
	}
	if r.URL.Query().Has("name") {
		queryMap["name"] = r.URL.Query().Get("name")
	}
	if r.URL.Query().Has("surname") {
		queryMap["surname"] = r.URL.Query().Get("surname")
	}
	if r.URL.Query().Has("patronymic") {
		queryMap["patronymic"] = r.URL.Query().Get("patronymic")
	}
	if r.URL.Query().Has("page") {
		queryMap["page"] = r.URL.Query().Get("page")
	}
	ctx = ctxusecase.AddQueryVars(ctx, queryMap)

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	var request Req
	if !request.IsEmpty() {
		err := json.NewDecoder(body).Decode(&request)
		if err != nil {
			WriteHttpError(w, ErrDecoding)
			log.Println(err)
			return
		}
	}

	log.Printf("Endpoint: %s\nQuery vars: %v\nPath vars: %v\nPayload: %v\n", r.URL, queryMap, mux.Vars(r), request)
	response, SuccesCode, err := wrapper.fn(ctx, request)
	if err != nil {
		log.Printf("Endpoint: %s\nError: %v\n", r.URL, err)
		WriteHttpError(w, err)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		WriteHttpError(w, ErrEncoding)
		return
	}

	w.WriteHeader(SuccesCode)
	log.Printf("Endpoint: %s\nCode: %d\nResponse:%v\n", r.URL, SuccesCode, string(rawJSON))
	if SuccesCode != http.StatusNoContent {
		_, err = w.Write(rawJSON)
		if err != nil {
			log.Println(err)
		}
	}
}
