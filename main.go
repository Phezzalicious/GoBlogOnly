package main

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"BlogApp/DataAccess"
	"html/template"
	"github.com/satori/go.uuid"
	
	
)

type Env struct {
	blogDB data.BlogStore
}
type user struct {
	UserName string
	First    string
	Last     string
}
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func main() {
	db, _ := data.NewMongoDB("mongodb+srv://phelps:1995@cluster0.ddb3v.mongodb.net/blog?retryWrites=true&w=majority")
	env := &Env{db}
	//fmt.Println("MY ADDED ID: " + fmt.Sprint(id))
	http.HandleFunc("/", env.IndexHandler)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/blog",env.BlogHomeHandler)
	http.HandleFunc("/createBlog", env.CreateBlogHandler)

	http.ListenAndServe(":8000", nil)
}


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

type cBlogStruct struct{
	Title string
	Heading string
	Message string
}
func(env *Env) CreateBlogHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		c := cBlogStruct{"Create a blog today!", "Blog form", "Make a blog"}

	
	err := tpl.ExecuteTemplate(w,"createblog.gohtml",c)

	if err != nil {
		log.Fatalln(err)
	}
		break
	case "POST":
		c := cBlogStruct{"Create a blog today!", "Blog form", "Make a blog"}
		title := r.FormValue("title")
		author := r.FormValue("author")
		topic := r.FormValue("topic")
		body := r.FormValue("body")
		fmt.Println("My form data: ")
		fmt.Println(title,author,topic,body)
		newPost := data.BlogPost {Title: title, Topic: topic,Body: body,Author: author}
		toPrint := env.blogDB.Create(newPost)
		fmt.Println(toPrint)
		err := tpl.ExecuteTemplate(w,"createblog.gohtml",c)
	
		if err != nil {
			log.Fatalln(err)
		}
		break
	}
	

}
func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}


type blogStruct struct{
	Title string
	Heading string
	Message string
	BlogTitle string
	Author string
	Body string
	Topic string
}
func (env *Env) BlogHomeHandler(w http.ResponseWriter, r *http.Request){

	//all := env.blogDB.ReadAll()
	env.blogDB.ReadAll()
	p := blogStruct{"Blog","Blog Posts","Posts for the blog","Title","Author","Body","Topic"}
	err := tpl.ExecuteTemplate(w,"blog.gohtml",p)
	if err != nil {
		log.Fatalln(err)
	}
}




type indexStruct struct {
	Title string
	Heading string
	Message string
}

func (env *Env) IndexHandler(w http.ResponseWriter, r *http.Request) {
	//p := indexStruct{"Home","Home","I am home"}



	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l,}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}






