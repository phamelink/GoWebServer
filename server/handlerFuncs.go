package server

import (
	pdb "PhilMessage/database"
	"fmt"
	"html/template"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	t:= template.Must(template.ParseFiles("./static/html/layout.gohtml",
		"./static/html/homepage.gohtml"))
	allUsers, err := pdb.GetAllUsers()
	if err != nil {
		w.WriteHeader(400)
		panic(err)
	}
	

	fmt.Println(allUsers[0].Name)
	check(t.ExecuteTemplate(w, "layout", allUsers))
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/layout.gohtml",
		"./static/html/login.gohtml"))
	check(t.ExecuteTemplate(w, "layout", ""))
}

func signup(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/layout.gohtml",
		"./static/html/signup.gohtml"))
	check(t.ExecuteTemplate(w, "layout", ""))
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	user := pdb.User{}
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	email := r.FormValue("email")
	name := r.FormValue("name")
	password := r.FormValue("password")
	user = pdb.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err = user.Create()
	if err != nil {
		fmt.Println("user no created")
	}

	http.Redirect(w, r, "/homepage", 303)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}