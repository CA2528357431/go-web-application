package main

import (
	"database/sql"
	"fmt"
)


func (p *Post)createPost(db *sql.DB)  {
	sql := "insert into post (username, content) values ($1,$2) returning id"
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}

	err = stmt.QueryRow(p.Username,p.Content).Scan(&(p.Id))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(p)
}

func getPost(db *sql.DB,id int) Post  {
	sql := "select id, username, content from post where id = $1"
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}
	p := Post{}
	err = stmt.QueryRow(id).Scan(&p.Id,&p.Username,&p.Content)
	if err!=nil{
		fmt.Println(err)
	}
	return p
}

func getPosts(db *sql.DB,name string) []Post  {
	sql := "SELECT id, username, content FROM post WHERE username = $1 "
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}
	posts := []Post{}
	rows,err := stmt.Query(name)
	if err!=nil{
		fmt.Println(err)
	}
	// rows是迭代器
	// next进行迭代

	for rows.Next(){
		p := Post{}
		err = rows.Scan(&p.Id,&p.Username,&p.Content)
		if err!=nil{
			fmt.Println(err)
		}
		posts = append(posts,p)
	}

	rows.Close()

	return posts
}

func (p *Post)changePost(db *sql.DB) {
	sql := "update post set username = $2, content = $3 where id = $1"
	_,err := db.Exec(sql,p.Id,p.Username,p.Content)
	// 不关心结果，用exec
	if err!=nil{
		fmt.Println(err)
	}
}

func removePost(db *sql.DB,id int) {
	sql := "delete from post where id = $1"
	_,err := db.Exec(sql,id)
	// 不关心结果，用exec
	if err!=nil{
		fmt.Println(err)
	}
}

func connect() *sql.DB {
	db,err := sql.Open("postgres","user=postgres dbname=post_test password=131128287 sslmode=disable")
	if err!=nil{
		fmt.Println(err)
	}
	return db
}
