package todo

import (
	"atodo/models"
	"atodo/storage"
	"errors"
	"fmt"
	"log"
	"os"
)

type todo struct {
	file *os.File
}

func NewTodo(file *os.File) storage.Todo {
	return &todo{
		file: file,
	}
}

func (c *todo) CreateUser(in *models.User) (*models.User, error) {

	user1, err := models.ReadJson(c.file)
	if err != nil {
		return nil, err
	}
	user1 = append(user1, *in)
	err = models.WriteJson(c.file, user1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return in, nil
}

func (c *todo) GetByID(in *models.User) (*models.User, error) {

	user, err := models.ReadJson(c.file)
	if err != nil {
		return in, err
	}
	flag := true
	for _, a := range user {
		if a.Id == in.Id {
			flag = false
			return &a, nil
		}
	}
	if flag {
		err := errors.New("id not found")
		fmt.Println(err.Error())
		return in, err
	}

	return in, nil
}
func (c *todo) GetAll() ([]models.User, error) {
	user, err := models.ReadJson(c.file)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (c *todo) UpdateUser(in *models.User) (*models.User, error) {
	//flag := true
	users1, err := models.ReadJson(c.file)
	fmt.Println(users1)
	if err != nil {
		return nil, err
	}
	for i, a := range users1 {
		if a.Id == in.Id {
			//flag = false
			users1[i] = *in
		}
	}

	err = models.WriteJson(c.file, users1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return in, err

}
func (c *todo) DeleteUser(in *models.User) (error error) {

	user, err := models.ReadJson(c.file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	flag := true
	for i, a := range user {
		if a.Id == in.Id {
			user = append(user[:i], user[i+1:]...)
			flag = false
			break
		}
	}
	err = models.WriteJson(c.file, user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if flag {
		log.Println("id not found")
	}

	return nil
}
