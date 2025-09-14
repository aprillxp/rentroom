package properties

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func CreateProperty(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		var req models.Property
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
