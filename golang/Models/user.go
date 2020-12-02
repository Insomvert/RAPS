package Models

import (
	"golang/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetListUser(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	var cols []string
	if user.FirstName != `` {
		cols = append(cols, `FirstName`)
	}
	if user.LastName != `` {
		cols = append(cols, `LastName`)
	}
	if err := Config.DB.Model(user).Where("id = ?", id).Select(cols).Update(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
