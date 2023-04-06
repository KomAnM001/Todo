package models

import (
	"encoding/json"
	"os"
)

func WriteJson(file *os.File, users []User) error {
	//	fmt.Println(users)
	usersByte, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.WriteAt(usersByte, 0)
	if err != nil {
		return err
	}
	return nil
}
