package db

import "database/sql"

func ConnectToDB() (*sql.DB,error){
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=todoye sslmode=disable"

	db,err := sql.Open("postgress",dsn)

	if err != nil{
		return nil,err
	};
	return db,nil;
	}