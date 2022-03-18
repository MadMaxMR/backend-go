package database

import (
	"errors"
)

func Delete(modelo interface{}, id string) (message string, err error) {
	type model struct {
		modelo interface{}
	}
	db := GetConnection()
	defer db.Close()
	result := db.Find(modelo, id)
	if result.RowsAffected != 0 {
		err := db.Delete(modelo, id).Error
		if err != nil {
			return "", errors.New("Error al eliminar - " + err.Error())
		}
		return "Elemento eliminado correctamente", nil
	} else {
		return "", errors.New("No se encontro datos con el ID: " + id)
	}
}
