package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(file *os.File) ([]User, error) {
	var users []User

	buf := new(bytes.Buffer)
	file.Seek(0, 0)
	_, err := buf.ReadFrom(file)

	if err != nil {
		return nil, err
	}

	if json.Valid(buf.Bytes()) {
		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			fmt.Println(err)
		}
	}
	return users, nil
}
