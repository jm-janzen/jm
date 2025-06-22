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

type Me struct {
	ID   int    `json:"id"`
	Mode string `json:"mode"`
	Data interface{}
}

type NotFound struct {
	Message string `json:"message"`
}
type Error struct {
	Message string `json:"message"`
}

// GetMe godoc
// @Summary      Get me
// @Description  get a representation of me by mode id
// @Produce      json
// @Param        id   path      int  false  "which mode"
// @Success      200  {object}  me.Me
// @Failure      404  {object}  me.NotFound
// @Failure      500  {object}  me.Error
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
			Error{errMessage},
		)
		return
	}

	for _, me := range modes {
		if me.ID == id {
			c.JSON(http.StatusOK, me)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		NotFound{fmt.Sprint("id not found: ", queryParam)},
	)
}

var modes []Me

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
}
