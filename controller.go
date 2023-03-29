package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateUsers(c *gin.Context) {
	var (
		users []User
		user  User
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Fatal(err)
	}

	if json.Valid(file) {
		err = json.Unmarshal(file, &users)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	users = append(users, user)

	usersByte, err := json.Marshal(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("dsadfjbsfjbsdjkbsvcdsvsd")
	err = os.WriteFile("todo.json", usersByte, 0644)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetByID(c *gin.Context) {
	var (
		users []User
	)

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return
	}

	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Println(err)
		return
	}

	if json.Valid(file) {
		err = json.Unmarshal(file, &users)
		if err != nil {
			return
		}
	}
	for i, a := range users {
		fmt.Println("users", users)
		if a.Id == id {
			c.JSON(http.StatusOK, users[i])
			break
		} else {
			log.Println("id not found")
			c.JSON(http.StatusBadRequest, "id not found")
			return
		}
	}
}

func GetAll(c *gin.Context) {
	var (
		users []User
	)
	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Fatal(err)
	}

	if json.Valid(file) {
		err = json.Unmarshal(file, &users)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c.JSON(http.StatusOK, users)

}

//type UpdateUserInput struct {
//	Id          int    `json:"id"`
//	Name        string `json:"name"`
//	Description string `json:"description"`
//}

func UpdateUser(c *gin.Context) {

	var (
		users []User
		user  User
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return
	}

	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Println(err)
		return
	}

	if json.Valid(file) {
		err = json.Unmarshal(file, &users)
		if err != nil {
			return
		}
	}

	for i, a := range users {
		if a.Id == id {
			//users = append(users[:i], users[i+1:]...)
			users[i] = user
			break
		} else {
			log.Println("invailed id ")
		}
	}
	//users = append(users, user)
	newUserBytes, err := json.Marshal(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = os.WriteFile("todo.json", newUserBytes, 0644)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUSer(c *gin.Context) {
	var (
		users []User
	)
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return
	}

	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Println(err)
		return
	}

	if json.Valid(file) {
		err = json.Unmarshal(file, &users)
		if err != nil {
			return
		}
	}

	for i, a := range users {
		if a.Id == id {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	usersByte, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("todo.json", usersByte, 0644)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "deleted successfully")
}
