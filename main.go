package main

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"BlogApp/DataAccess"
	"html/template"
)

type Env struct {
	blogDB data.BlogStore
}


func main() {
	db, _ := data.NewMongoDB("mongodb+srv://phelps:1995@cluster0.ddb3v.mongodb.net/blog?retryWrites=true&w=majority")
	env := &Env{db}
	//fmt.Println("MY ADDED ID: " + fmt.Sprint(id))
	http.HandleFunc("/", env.IndexHandler)
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
	p := indexStruct{"Home","Home","I am home"}
	err := tpl.ExecuteTemplate(w,"index.gohtml",p)
	if err != nil {
		log.Fatalln(err)
	}
}






