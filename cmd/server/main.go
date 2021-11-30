package main

import (
	"bytes"
	"context"
	"flag"
	"io"
	"io/ioutil"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bilal-bhatti/echo/pkg/services"
	"github.com/go-chi/chi/v5"
)

type copyWriter struct {
	http.ResponseWriter
	statusCode int
	capture    *bytes.Buffer
}

func (cw *copyWriter) Write(p []byte) (int, error) {
	i, err := cw.ResponseWriter.Write(p)
	if err != nil {
		return i, err
	}
	return cw.capture.Write(p)
}

func (cw *copyWriter) WriteHeader(code int) {
	cw.statusCode = code
	cw.ResponseWriter.WriteHeader(code)
}

func Inspector(router chi.Router) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log.Println("request url", r.URL)

			if r.ContentLength > 0 {
				buf := &bytes.Buffer{}
				tee := io.TeeReader(r.Body, buf)
				r.Body = ioutil.NopCloser(buf)

				capture, _ := ioutil.ReadAll(tee)
				request := bytes.NewBuffer(capture)

				log.Println("request", request.String())
			}

			copy := &copyWriter{
				ResponseWriter: w,
				capture:        &bytes.Buffer{},
			}

			next.ServeHTTP(copy, r)

			rctx := r.Context().Value(chi.RouteCtxKey).(*chi.Context)
			log.Println("matched", rctx.RoutePatterns)
			log.Println("path params", rctx.URLParams)

			log.Println("query params", r.URL.Query())
			log.Println("status code", copy.statusCode)
			log.Println("response", copy.capture.String())
			log.Println("headers", w.Header())

		}
		return http.HandlerFunc(fn)
	}
}

func main() {
	listenAddr := flag.String("listen", ":5678", "spcifiy port to listen on")
	flag.Parse()

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	// Listen for OS signals
	signal.Notify(quit, os.Interrupt)

	// Configure a new logger
	logger := log.New(os.Stdout, "ECHO: ", log.LstdFlags)

	mux := chi.NewRouter()

	mux.Use(Inspector(mux))

	mux = services.NewRouter(mux)

	// Init http server
	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Start shutdown go routine
	go gracefullShutdown(server, logger, quit, done)

	// Ready to serve
	logger.Println("Server is ready to handle requests at", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", *listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func gracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
