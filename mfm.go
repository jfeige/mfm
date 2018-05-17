package mfm

import (
	"database/sql"
)

func FormatRows(rows *sql.Rows)[]map[string]interface{}{
	columns,_ := rows.Columns()
	cnt := len(columns)
	pts := make([]interface{},cnt)
	containers := make([]interface{},cnt)
	data := make(map[string]interface{})
	record := make([]map[string]interface{},0)

	for i := range pts{
		pts[i] = &containers[i]
	}
	for rows.Next(){
		//把结果指针放入pts，而pts本身保存的就是containers的指针地址
		err := rows.Scan(pts...)
		if err != nil{
			continue
		}
		for pos,title := range columns{
			data[title] = containers[pos]
		}
		record = append(record,format(data))
	}

	return record
}

func format(src map[string]interface{})map[string]interface{}{
	ret := make(map[string]interface{})
	for k,v := range src{
		switch v := v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			ret[k] = v
		case []uint8:
			ret[k] = string(v)
		}
	}
	return ret
}
