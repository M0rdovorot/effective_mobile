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

	featureId := r.URL.Query().Get("feature_id")
	tagId := r.URL.Query().Get("tag_id")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	use_last_revision := r.URL.Query().Get("use_last_revision")
	ctx = ctxusecase.AddQueryVars(ctx, map[string]string{
		"feature_id": featureId,
		"tag_id": tagId,
		"limit": limit,
		"offset": offset,
		"use_last_revision": use_last_revision,
	})

	token := r.Header.Get("token")
	ctx = ctxusecase.AddToken(ctx, token)
	
	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)
	
	var request Req
	if !request.IsEmpty(){
		err := json.NewDecoder(body).Decode(&request)
		if err != nil {
			WriteHttpError(w, ErrDecoding)
			log.Println(err)
			return
		}
	}

	response, SuccesCode, err := wrapper.fn(ctx, request)
	if err != nil {
		log.Printf("%s: error: %v\n", r.URL, err)
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
	log.Printf("%s: %d", r.URL, SuccesCode)
	if SuccesCode != http.StatusNoContent {
		_, err = w.Write(rawJSON)
		if err != nil {
			log.Println(err)
		}
	}
}