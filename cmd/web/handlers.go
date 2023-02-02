package main

import (
	"errors"
	"fmt"
	"github.com/Mike-95/blog_web/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.html", &templateData{
		Posts: s,
	})
}

func (app *application) createPost(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		app.clientError(writer, http.StatusBadRequest)
		return
	}

	title := request.PostForm.Get("title")
	content := request.PostForm.Get("content")
	category := request.PostForm.Get("category")

	id, err := app.posts.Insert(title, content, category)
	if err != nil {
		app.serverError(writer, err)
		return
	}

	http.Redirect(writer, request, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)

}

func (app *application) showPost(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	s, err := app.posts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}
		return
	}

	app.render(writer, request, "show.page.html", &templateData{
		Post: s,
	})
}

func (app *application) createPostForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", nil)
}
