package api

import (
	"db"
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
	store db.Storage
	port  string
}

func MakeApiServer(port string, store db.Storage) *ApiServer {
	return &ApiServer{
		store: store,
		port:  port,
	}
}

func (s *ApiServer) HandleEndpoints() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", makehttpHandlerFunc(s.handleTest))
	mux.HandleFunc("/test", makehttpHandlerFunc(s.handleAddFile))
	log.Printf("Running on port %s", s.port)
	http.ListenAndServe(s.port, mux)
	return nil
}

func (s *ApiServer) handleTest(w http.ResponseWriter, r *http.Request) error {
	WriteJson(w, http.StatusOK, "hey, its S3 speaking")
	return nil
}

func (s *ApiServer) handleAddFile(w http.ResponseWriter, r *http.Request) error {
	
	WriteJson(w, http.StatusOK, "file added")
	return nil
}

type RespWriterRequest struct {
	w http.ResponseWriter
	r *http.Request
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makehttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJson(w, http.StatusBadRequest, "err")
		}
	}
}

func WriteJson(w http.ResponseWriter, code int, msg any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(msg)
}
