package migration

import (
	"net/http"

	"template/datasource"
	"template/entity"
	"template/function"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type post entity.Post

func Migrate(w http.ResponseWriter, r *http.Request) {
	autoMigrate()

	function.SendResponse(w, http.StatusOK, "Success", nil)
}

// autoMigrate ...
func autoMigrate() {
	db := datasource.OpenDB()
	db.AutoMigrate(
		&post{},
	)
}
