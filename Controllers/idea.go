package Controllers

import (
	"fmt"
	"net/http"
    "github.com/labstack/echo"
	"example.com/niceidea/Utilities"
)

func IdeaListPage(c echo.Context) error {
	db := Utilities.Connect()
	defer db.Close()
	
	//データの検索
    rows, err := db.Query("SELECT title FROM ideas_idea")
    Utilities.CheckErr(err)

	for rows.Next() {
        var title string
        err = rows.Scan(&title)
        Utilities.CheckErr(err)
    }
	fmt.Println(rows)
	return c.String(http.StatusOK, "hello")
}

