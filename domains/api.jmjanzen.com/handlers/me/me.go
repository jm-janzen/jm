package me

import (
	"fmt"
	"jm/internal/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Modes struct {
	Modes []Me `json:"modes"`
}

type Me struct {
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

// GetMe godoc
//
//	@Summary		Get me
//	@Description	get a representation of me by mode id
//	@Produce		json
//	@Param			id	path		int	false	"which mode"
//	@Success		200	{object}	me.Me
//	@Failure		404	{object}	me.NotFound
//	@Failure		500	{object}	me.Error
//	@Router			/me/{id} [get]
//
// Get associated me from modes. If id is not provided, default to 1
// FIXME Add flag to make dates human readable?
func GetMe(c *gin.Context) {

	client, cancel := db.Connect()
	defer cancel()

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

	var me Me
	collection := db.GetCollection(client, "modes")
	result := collection.FindOne(c, bson.M{"id": id})
	result.Decode(&me)

	c.JSON(http.StatusOK, me)
}
