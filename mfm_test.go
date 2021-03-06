package mfm

import (
	"testing"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Test_FormatRows(t *testing.T){


	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/x_000?multiStatements=true")
	if err != nil {
		fmt.Println(err)
		return
	}
	sql := "select uid,lat,lng from x_member_active_0 order by uid asc limit 10"
	stmt,err := db.Prepare(sql)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	rows,err := stmt.Query()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rows.Close()

	record := FormatRows(rows)

	for _,v := range record{
		fmt.Println(v)
	}
}
