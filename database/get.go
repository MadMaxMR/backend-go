package database

import (
	"errors"
	"strconv"
)

func GetAll(modelo interface{}, page string) (mod interface{}, err error) {
	type model struct {
		modelo struct{}
	}
	db := GetConnection()
	defer db.Close()
	pageInt, _ := strconv.Atoi(page)
	if page == "" {
		db.Order("id ASC").Find(modelo)
		return modelo, nil
	}
	if page == "1" {
		result := db.Limit(10).Order("id ASC").Find(modelo)
		if result.RowsAffected != 0 {
			return modelo, nil
		} else {
			return nil, errors.New("No se encontro datos en la página: " + page)
		}
	} else {
		result := db.Limit(10).Offset((pageInt - 1) * 10).Order("id ASC").Find(modelo)
		if result.RowsAffected != 0 {
			return modelo, nil
		} else {
			return nil, errors.New("No se encontro datos en la página: " + page)
		}
	}
	return nil, nil
}

func Get(modelo interface{}, id string) (mod interface{}, err error) {
	type model struct {
		modelo struct{}
	}
	db := GetConnection()
	defer db.Close()
	idInt, _ := strconv.Atoi(id)
	result := db.Find(modelo, idInt)
	if result.RowsAffected != 0 {
		return modelo, nil
	} else {
		return nil, errors.New("No se encontro datos con el ID: " + id)
	}
}
