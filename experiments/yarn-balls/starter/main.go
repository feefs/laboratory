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
	listenerA      weaver.Listener
	listenerB      weaver.Listener
}

func serve(ctx context.Context, app *app) error {

	fmt.Printf("listenerA: %v\n", app.listenerA)
	fmt.Printf("listenerB: %v\n", app.listenerB)

	muxA := http.NewServeMux()
	muxB := http.NewServeMux()

	muxA.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")

		manipulator := app.strManipulator.Get()
		reversed, err := manipulator.Reverse(ctx, word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, reversed)
	})

	muxB.HandleFunc("/capitalize", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("word")

		manipulator := app.strManipulator.Get()
		capitalized, err := manipulator.Capitalize(ctx, word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, capitalized)
	})

	errs := make(chan error)
	go func() {
		fmt.Println("Serving HTTP requests on listenerA...")
		if err := http.Serve(app.listenerA, muxA); err != nil {
			errs <- err
		}
	}()
	go func() {
		fmt.Println("Serving HTTP requests on listenerB...")
		if err := http.Serve(app.listenerB, muxB); err != nil {
			errs <- err
		}
	}()

	return <-errs
}
