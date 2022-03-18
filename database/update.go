package database

import (
	"API/auth"
	"errors"
	"net/http"
)

func Update(req *http.Request, modelo interface{}, id string) (mod interface{}, err error) {
	type model struct {
		modelo interface{}
	}
	db := GetConnection()
	defer db.Close()
	result := db.Find(modelo, id)

	if result.RowsAffected != 0 {
		err = auth.ValidateBody(req, modelo)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		err = db.Save(modelo).Error
		if err != nil {
			return nil, errors.New("Error al actualizar - " + err.Error())
		}
		return modelo, nil
	} else {
		return nil, errors.New("No se encontro datos con el ID: " + id)
	}
}
