package main

import (
	"bufio"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"strings"
	"text/template"
	"time"

	"git.sr.ht/grauwoelfchen/typol/typol/service"
)

type ServerAddr string

const maxMemory = 5 << 20

const serverAddr ServerAddr = "serverAddr"

//go:embed "home.tmpl"
var tmplHome string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routeHome)
	mux.HandleFunc("/convert", routeConvert)

	ctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:              ":3000",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      7 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverAddr, l.Addr().String())
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

func routeHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").Parse(tmplHome)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]string{}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func routeConvert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var text string

	// multipart/form-data
	contentType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(contentType, "multipart/") {
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			log.Fatal(err)
		}

		for k := range r.MultipartForm.File {
			file, _, err := r.FormFile(k)
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text += scanner.Text()
			}

			if err = scanner.Err(); err != nil {
				log.Println(err)
			}
		}
	}

	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		text = r.PostFormValue("text")
	}

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIXME: input validations
	from := r.PostFormValue("in")
	to := r.PostFormValue("out")

	log.Printf(
		"addr: %s from: %s to: %s text: %s\n",
		ctx.Value(serverAddr), from, to, text,
	)

	out := convert(from, to, text)
	if _, err = io.WriteString(w, fmt.Sprintf("%s\n", out)); err != nil {
		log.Println(err)
	}
}

func convert(from, to, text string) string {
	args := []string{"convert", "-in", from, "-out", to, text}

	out, err := service.Run(args)
	if err != nil {
		log.Fatal(err)
	}

	return out
}
