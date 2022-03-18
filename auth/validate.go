package auth

import (
	"API/modelos"
	"API/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ValidateBody(req *http.Request, modelo interface{}) error {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return errors.New("error al leer los datos del body")
	}
	err = json.Unmarshal(body, modelo)

	if err != nil {
		return errors.New("error al guardar los datos del body")
	}
	return nil
}

func ValidateCurso(curso *models.Cursos) error {

	if curso.Nombre == "" {
		return errors.New("Required field 'nombre'")
	}
	if curso.Description == "" {
		return errors.New("Required field 'description'")
	}
	if curso.Image == "" {
		return errors.New("Required field 'image'")
	}
	return nil
}

func ValidateTema(tema *models.Temas) error {
	if tema.Title == "" {
		return errors.New("Required field 'Title'")
	}
	if tema.Description == "" {
		return errors.New("Required field 'Description'")
	}
	if tema.CursosID == 0 {
		return errors.New("Required field 'CursoID'")
	}
	return nil
}
func ValidateUsuario(usuario *modelos.Usuarios) error {
	fmt.Printf("%+v\n", usuario)
	/*if usuario.Nombres == "" {
		return errors.New("Required field 'nombres'")
	}
	if usuario.Apellidos == "" {
		return errors.New("Required field 'apellidos'")
	}
	if usuario.Email == "" {
		return errors.New("Required field 'email'")
	}
	if usuario.Password == "" {
		return errors.New("Required field 'password'")
	}
	if usuario.Dni == 0 {
		return errors.New("Required field 'dni'")
	}
	if usuario.Dni > 100000000 || usuario.Dni < 10000000 {
		return errors.New("Field 'dni' must be 8 digits")
	}
	if usuario.Universidad == "" {
		return errors.New("Required field 'universidad'")
	}
	if usuario.Celular == 0 {
		return errors.New("Required field 'celular'")
	}
	if usuario.Celular > 1000000000 || usuario.Celular < 100000000 {
		return errors.New("Field 'celular' must be 9 digits")
	}*/
	return nil
}

func ValidateLogin(usuario *models.Usuarios) error {
	if usuario.Email == "" {
		return errors.New("Required field 'email'")
	}
	if usuario.Password == "" {
		return errors.New("Required field 'password'")
	}
	return nil
}
