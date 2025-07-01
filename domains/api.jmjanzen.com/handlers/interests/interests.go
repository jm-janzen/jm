package interests

import (
	"jm/internal/db"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type Interest struct {
	Slug           string   `bson:"slug" json:"slug"`
	Name           string   `bson:"name" json:"name"`
	Passion        float64  `bson:"passion" json:"passion"`
	Summary        string   `bson:"summary" json:"summary"`
	Aliases        []string `bson:"aliases" json:"aliases"`
	ElaborationUrl *string  `bson:"elaborationUrl" json:"elaborationUrl,omitempty"`
}

// GetInterests godoc
//
//	@Summary		Get all interests
//	@Description	get a JSON array of strings
//	@Product		json
//	@Success		200	{array} string
//	@Router			/interests [get]
//
// Returns JSON array of interests in slug form like
//
//	["reading", "coding", "etc"]
//
// which may then be used to look up more detail using
//
//	GET /interest/:slug
func GetInterests(c *gin.Context) {
	collection, cancel := db.GetCollection("interests")
	defer cancel()

	results, err := collection.Find(c, bson.M{})
	if err != nil {
		log.Println(err)
	}

	var (
		interest Interest
		slugs    []string
	)
	defer results.Close(c)
	for results.Next(c) {
		if err = results.Decode(&interest); err != nil {
			log.Println(err)
		}

		slugs = append(slugs, interest.Slug)
	}

	c.JSON(http.StatusOK, slugs)
}

// GetInterest godoc
//
//	@Summary		Get interest
//	@Description	get interest, based on provided slug
//	@Produce		json
//	@Param			slug	path		string	true	"slug of specific interest"
//	@Success		200		{object}	interests.Interest
//	@Router			/interest/{slug} [get]
func GetInterest(c *gin.Context) {
	var interest Interest
	collection, cancel := db.GetCollection("interests")
	defer cancel()

	result := collection.FindOne(
		c, bson.M{"slug": c.Param("slug")},
	)

	result.Decode(&interest)
	if interest.Slug == "" {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "slug not found"},
		)
		return
	}

	c.JSON(http.StatusOK, interest)
}
