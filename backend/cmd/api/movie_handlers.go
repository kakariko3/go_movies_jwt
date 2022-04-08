package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからIDを取得
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	// 映画を一つ取得
	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// JSONの生成
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	// 映画を全て取得
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// JSONの生成
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {

// }
