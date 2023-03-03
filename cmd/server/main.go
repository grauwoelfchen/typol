package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"git.sr.ht/grauwoelfchen/typol/typol/service"
)

type ServerAddr string

const maxMemory = 5 << 20

const kServerAddr ServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/convert", convert)

	ctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, kServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server closed")
		} else if err != nil {
			log.Fatal(err)
		}
		cancel()
	}()

	<-ctx.Done()
}

func convert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		log.Fatal(err)
	}

	text := r.PostFormValue("text")
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIXME
	from := "Dvorak"
	to := "Qwerty"

	log.Printf(
		"addr: %s from: %s to: %s text: %s\n",
		ctx.Value(kServerAddr), from, to, text,
	)

	args := []string{"convert", "-in", from, "-out", to, text}
	out, err := service.Run(args)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.WriteString(w, fmt.Sprintf("%s\n", out)); err != nil {
		log.Fatal(err)
	}
}
