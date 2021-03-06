package database

import (
	"errors"
)

func Create(modelo interface{}) (mod interface{}, err error) {
	db := GetConnection()
	defer db.Close()
	err = db.Create(modelo).Error
	if err != nil {
		return nil, errors.New("Error al guardar - " + err.Error())
	}
	return modelo, nil
}
