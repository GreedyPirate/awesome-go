package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMysql(t *testing.T) {
	db,err := sql.Open("mysql","root:0817@tcp(127.0.0.1:3306)/dq?charset=utf8");
	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}

	result, err := db.Query("select * from t1")

	for result.Next() {
		var (
			id int
			name string
			age int
		)
		result.Scan(&id, &name, &age)
		t.Log(id, name, age)
	}
}
