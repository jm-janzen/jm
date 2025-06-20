package me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Modes struct {
	Modes []Me `json:"modes"`
}

// FIXME Make structs for each mode, too
// Probably put these in their own modes/ dir...
type Me struct {
	ID   int    `json:"id"`
	Name *Name  `json:"name,omitempty"`
	Mode string `json:"mode"`
}

type Name struct {
	First *string `json:"first,omitempty"`
	Nick  *string `json:"nick,omitempty"`
	Last  *string `json:"last,omitempty"`
}

// GetMe godoc
// @Summary      Get me
// @Description  get a representation of me by mode id
// @Produce      json
// @Param        id   path      int  false  "which mode"
// @Success      200  {object}  me.Me
// @Router       /me/{id} [get]
// Get associated me from modes. If id is not provided, default to 1
// FIXME Add flag to make dates human readable?
func GetMe(c *gin.Context) {
	var queryParam string
	if queryParam = c.Param("id"); queryParam == "" {
		queryParam = "1"
	}

	var id, err = strconv.Atoi(queryParam)
	if err != nil {
		var errMessage = fmt.Sprintf("Couldn't parse id: %s", queryParam)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": errMessage},
		)
		return
	}

	for _, me := range modes {
		if me.ID == id {
			// Merge match with defaults
			data, err := json.Marshal(me)
			if err != nil {
				fmt.Println("Error marshalling:", err)
			}
			json.Unmarshal(data, &defaultMode)
			c.JSON(http.StatusOK, me)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{"message": fmt.Sprint("id not found: ", queryParam)},
	)
}

var modes []Me
var defaultMode Me

// FIXME Make... better. Much better.
func init() {
	file, err := os.Open("./assets/modes.json")
	if err != nil {
		panic(err)
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &modes)

	// for default mode
	file, err = os.Open("./assets/default-mode.json")
	if err != nil {
		panic(err)
	}

	byteValue, err = io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &defaultMode)
}
