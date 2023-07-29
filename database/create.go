package database

import (
	"errors"
)

func Create(modelo interface{}) (mod interface{}, err error) {
	db := GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	err = db.Create(modelo).Error
	if err != nil {
		return nil, errors.New("Error al guardar - " + err.Error())
	}
	return modelo, nil
}
