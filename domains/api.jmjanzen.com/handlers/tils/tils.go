package tils

import (
	"jm/internal/db"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type Til struct {
	Slug           string   `bson:"slug" json:"slug"`
	Name           string   `bson:"name" json:"name"`
	Summary        string   `bson:"summary" json:"summary"`
}

// GetTils godoc
//
//	@Summary		Get all tils
//	@Description	get a JSON array of strings
//	@Product		json
//	@Success		200	{array} string
//	@Router			/tils [get]
//
// Returns JSON array of tils in slug form like
//
//	["cats-purr", "soup-is-hot", "etc"]
//
func GetTils(c *gin.Context) {
	collection, cancel := db.GetCollection("tils")
	defer cancel()

	results, err := collection.Find(c, bson.M{})
	if err != nil {
		log.Println(err)
	}

	var (
		til Til
		slugs    []string
	)
	defer results.Close(c)
	for results.Next(c) {
		if err = results.Decode(&til); err != nil {
			log.Println(err)
		}

		slugs = append(slugs, til.Slug)
	}

	c.JSON(http.StatusOK, slugs)
}

// GetTil godoc
//
//	@Summary		Get til
//	@Description	get til, based on provided slug
//	@Produce		json
//	@Param			slug	path		string	true	"slug of specific til"
//	@Success		200		{object}	tils.Til
//	@Router			/tils/{slug} [get]
func GetTil(c *gin.Context) {
	var til Til
	collection, cancel := db.GetCollection("tils")
	defer cancel()

	result := collection.FindOne(
		c, bson.M{"slug": c.Param("slug")},
	)

	result.Decode(&til)
	if til.Slug == "" {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "slug not found"},
		)
		return
	}

	c.JSON(http.StatusOK, til)
}

