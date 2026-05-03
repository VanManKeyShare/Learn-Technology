package handler

import (
	"vmk-gin-app-docker/module/response"
	MySQL "vmk-gin-app-docker/module/service/db"

	"github.com/gin-gonic/gin"
)

func Check_Health_Database(c *gin.Context) {

	if err := MySQL.DB.Ping(); err != nil {
		response.InternalError(c, "DATABASE UNREACHABLE", err.Error())
		return
	}

	rows, err := MySQL.DB.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		response.InternalError(c, "QUERY FAILED", err.Error())
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		response.InternalError(c, "FAILED TO GET COLUMNS", err.Error())
		return
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			response.InternalError(c, "FAILED TO SCAN ROW", err.Error())
			return
		}
		row := make(map[string]interface{})
		for i, col := range columns {
			if b, ok := values[i].([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = values[i]
			}
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		response.InternalError(c, "ROW ITERATION ERROR", err.Error())
		return
	}

	// KÈM META PAGINATION NẾU CẦN
	response.Success_With_Meta(c, "OK", result, gin.H{
		"total": len(result),
	})
}
