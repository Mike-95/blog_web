package main

import (
	"errors"
	"fmt"
	"github.com/Mike-95/blog_web/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, post := range s {
		fmt.Fprintf(w, "%v\n", post)
	}

	//files := []string{
	//	"./ui/html/home.page.html",
	//	"./ui/html/base.layout.html",
	//	"./ui/html/footer.partial.html",
	//}
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	http.Error(w, "Internal Server Error", 500)
	//	return
	//}
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	app.serverError(w, err)
	//	http.Error(w, "Internal Server Error", 500)
	//}
}

func (app *application) createPost(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	title := "title test"
	content := "content test"
	category := "category test"

	id, err := app.posts.Insert(title, content, category)
	if err != nil {
		app.serverError(writer, err)
		return
	}

	http.Redirect(writer, request, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)

}

func (app *application) showPost(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
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
	fmt.Fprintf(writer, "%v", s)
}
