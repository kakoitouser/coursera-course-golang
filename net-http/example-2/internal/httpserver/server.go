package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/models"
	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/store"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	Address     string
	store       store.BookStore
}

func NewServer(ctx context.Context, address string, store store.BookStore) *Server {
	return &Server{
		ctx:         ctx,
		Address:     address,
		store:       store,
		idleConnsCh: make(chan struct{}),
	}
}

func mainHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("response to addr ", r.URL.Path)
	rw.Header().Set("Content-type", "text/html; charset=utf-8")
	rw.Write([]byte("<h1>Main page</h1>"))
}

func (srv *Server) basicHandler() chi.Router {
	r := chi.NewRouter()
	r.Post("/books", func(rw http.ResponseWriter, r *http.Request) {
		book := new(models.Book)
		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			fmt.Fprintf(rw, "Error %v", err)
			return
		}
		srv.store.Create(r.Context(), book)
	})
	r.Get("/books", func(rw http.ResponseWriter, r *http.Request) {
		books, err := srv.store.GetAll(r.Context())
		if err != nil {
			fmt.Fprintf(rw, "Error ", err)
			return
		}

		rw.Header().Add("Content-type", "application/json")
		if err := json.NewEncoder(rw).Encode(books); err != nil {
			fmt.Fprintf(rw, "Error ", err)
			return
		}
	})
	r.Get("/books/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fmt.Fprintf(rw, "Error %v", err)
			return
		}
		book, err := srv.store.GetById(r.Context(), id)
		render.JSON(rw, r, book)
	})
	r.Put("/books", func(rw http.ResponseWriter, r *http.Request) {
		book := new(models.Book)
		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			fmt.Fprintf(rw, "Error %v", err)
			return
		}
		srv.store.Update(r.Context(), book)
	})
	r.Delete("/books/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fmt.Fprintf(rw, "Error %v", err)
			return
		}
		if err := srv.store.Delete(r.Context(), id); err != nil {
			fmt.Fprintf(rw, "Error %v", err)
			return
		}
	})
	return r
}
func (srv *Server) Run() error {
	s := &http.Server{
		Addr:         srv.Address,
		Handler:      srv.basicHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go srv.ListenCtxForGT(s)
	log.Println("server runned on ", srv.Address)
	return s.ListenAndServe()
}

func (srv *Server) ListenCtxForGT(s *http.Server) {
	<-srv.ctx.Done()

	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down %v", err)
	}
	log.Println("[HTTP] Proccessed all idle connections")
	close(srv.idleConnsCh)
}

func (srv *Server) WaitForGracefulTermination() {
	<-srv.idleConnsCh
}
