package database

import (
	"errors"
	"net/http"
)

func Update(req *http.Request, modelo interface{}, id string) (mod interface{}, err error) {
	db := GetConnection()
	defer db.Close()
	err = db.Model(modelo).Where("id = ?", id).Update(modelo).Error

	if err != nil {
		return nil, errors.New("Error al actualizar - " + err.Error())
	}
	return modelo, nil

}
