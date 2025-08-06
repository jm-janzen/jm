package modes

import (
	"fmt"
	"jm/internal/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Modes struct {
	Modes []Modes `json:"modes"`
}

type Mode struct {
	ID   int            `bson:"id" json:"id"`
	Mode string         `bson:"mode" json:"mode"`
	Data map[string]any `bson:"data" json:"data"`
}

type NotFound struct {
	Message string `bson:"message" json:"message"`
}
type Error struct {
	Message string `bson:"message" json:"message"`
}

// GetModes godoc
//
//	@Summary		Get my modes
//	@Description	get a representation of my mode by mode id
//	@Produce		json
//	@Param			id	path		int	false	"which mode"
//	@Success		200	{object}	modes.Mode
//	@Failure		404	{object}	modes.NotFound
//	@Failure		500	{object}	modes.Error
//	@Router			/modes/{id} [get]
//
// Get associated modes associated with given id.
// If id is not provided, default to 1
func GetModes(c *gin.Context) {

	var queryParam string
	if queryParam = c.Param("id"); queryParam == "" {
		queryParam = "1"
	}

	id, err := strconv.Atoi(queryParam)
	if err != nil {
		var errMessage = fmt.Sprintf("Couldn't parse id: %s", queryParam)
		c.JSON(
			http.StatusInternalServerError,
			Error{errMessage},
		)
		return
	}

	var mode Mode

	collection, cancel := db.GetCollection("modes")
	defer cancel()

	result := collection.FindOne(c, bson.M{"id": id})
	result.Decode(&mode)

	c.JSON(http.StatusOK, mode)
}
