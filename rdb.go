package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang?charset=utf8")
	if err != nil {
		fmt.Println("オープンできていません")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("疎通できてません")
	} else {
		fmt.Println("疎通できました")
	}

	type User struct {
		UserID    string
		UserName  string
		CreatedAt time.Time
	}

	var (
		ctx context.Context
	)

	rows, err := db.QueryContext(ctx, `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        time.Time
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", err)
		}
		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users: %v", err)
	}
}
