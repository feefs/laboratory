package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	strManipulator weaver.Ref[StrManipulator]
	listener       weaver.Listener
}

func serve(ctx context.Context, app *app) error {

	fmt.Printf("Listener: %v\n", app.listener)

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")

		manipulator := app.strManipulator.Get()
		reversed, err := manipulator.Reverse(ctx, word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, reversed)
	})

	http.HandleFunc("/capitalize", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")

		manipulator := app.strManipulator.Get()
		capitalized, err := manipulator.Capitalize(ctx, word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, capitalized)
	})

	return http.Serve(app.listener, nil)
}
