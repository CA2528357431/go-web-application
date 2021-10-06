package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type Post struct {
	Content string
	Username string
	Id int
}


var db *sql.DB

func main() {
	db = connect()

	http.HandleFunc("/main",mux)

	http.ListenAndServe(":8888",nil)

}

func mux(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.Method)
	switch r.Method{
	case http.MethodGet:
		HandleGet(w,r)
	case http.MethodPut:
		HandlePut(w,r)
	case http.MethodPost:
		HandlePost(w,r)
	case http.MethodDelete:
		HandleDelete(w,r)
	}
}

func HandleGet(w http.ResponseWriter,r *http.Request)  {
	query := r.URL.Query()
	name := query.Get("name")
	ids := query.Get("id")
	if name==""{
		id,err := strconv.Atoi(ids)
		if err!=nil{
			fmt.Println(err)
		}
		p := getPost(db,id)
		if p.Id!=0 {
			data, err := json.Marshal(p)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(data)
		}else {
			str := "not found"
			data, err := json.Marshal(str)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(data)

		}
	}else {
		p := getPosts(db,name)
		if len(p)!=0{
			data, err := json.Marshal(p)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(data)
		}else {
			str := "not found"
			data, err := json.Marshal(str)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(data)

		}
	}

}

func HandlePut(w http.ResponseWriter,r *http.Request)  {
	data := make([]byte,r.ContentLength)
	r.Body.Read(data)
	p := Post{}
	err := json.Unmarshal(data,&p)
	if err != nil {
		fmt.Println(err)
	}
	p.createPost(db)

	str := "put it successfully"
	strdata, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(strdata)
}

func HandlePost(w http.ResponseWriter,r *http.Request)  {
	data := make([]byte,r.ContentLength)
	r.Body.Read(data)
	p := Post{}
	err := json.Unmarshal(data,&p)
	if err != nil {
		fmt.Println(err)
	}
	p.changePost(db)

	str := "post it successfully"
	strdata, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(strdata)

}
func HandleDelete(w http.ResponseWriter,r *http.Request)  {
	query := r.URL.Query()
	ids := query.Get("id")
	id,err := strconv.Atoi(ids)
	if err!=nil{
		fmt.Println(err)
	}
	removePost(db,id)

	str := "delete it successfully"
	strdata, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(strdata)
}