package agenda

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Service
}

func NewHandler(s Service) Handler {
	return &handler{
		s,
	}
}

func (h *handler) Create(c *gin.Context) {
	var agenda CreateAgenda
	if err := c.ShouldBindJSON(&agenda); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	res, err := h.Service.Create(c, &agenda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) Show(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.Show(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) Index(c *gin.Context) {
	res, err := h.Service.Index(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) Update(c *gin.Context) {
	id := c.Param("id")
	var agenda CreateAgenda
	if err := c.ShouldBindJSON(&agenda); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	res, err := h.Service.Update(c, id, &agenda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete agenda",
	})
}
