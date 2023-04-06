package controllers

import (
	"atodo/models"
	"atodo/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type todoService struct {
	strg storage.StoreI
}

func NewServiceTodo(strg storage.StoreI) todoService {
	return todoService{
		strg: strg,
	}
}

func (s *todoService) CreateUsers(c *gin.Context) {
	var (
		user *models.User
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.strg.Todo().CreateUser(user)
	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, user)
}

func (s *todoService) GetByID(c *gin.Context) {
	var (
		user models.User
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("here::::1")
	user.Id = id
	fmt.Println("here::::2")

	user1, err := s.strg.Todo().GetByID(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user1)

}

func (s *todoService) GetAll(c *gin.Context) {

	user1, err := s.strg.Todo().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user1)

}

func (s *todoService) UpdateUser(c *gin.Context) {

	var (
		user models.User
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = s.strg.Todo().UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, user)

}

func (s *todoService) DeleteUSer(c *gin.Context) {
	var (
		user models.User
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	err = s.strg.Todo().DeleteUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "deleted successfully")
}
