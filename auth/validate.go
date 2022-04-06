package auth

import (
	"encoding/json"
	"errors"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/MadMaxMR/backend-go/models"
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
		return errors.New("required field 'nombre'")
	}
	if curso.Description == "" {
		return errors.New("required field 'description'")
	}
	if curso.Image == "" {
		return errors.New("required field 'image'")
	}
	return nil
}

func ValidateTema(tema *models.Temas) error {
	if tema.Title == "" {
		return errors.New("required field 'Title'")
	}
	if tema.Description == "" {
		return errors.New("required field 'Description'")
	}
	if tema.CursosID == 0 {
		return errors.New("required field 'CursoID'")
	}
	return nil
}
func ValidateUsuario(usuario *modelos.Usuarios) error {

	if usuario.Nombres == "" {
		return errors.New("required field 'nombres'")
	}
	if usuario.Apellidos == "" {
		return errors.New("required field 'apellidos'")
	}
	if usuario.Email == "" {
		return errors.New("required field 'email'")
	}
	if usuario.Password == "" {
		return errors.New("required field 'password'")
	}
	if usuario.Dni == 0 {
		return errors.New("required field 'dni'")
	}
	if usuario.Dni > 100000000 || usuario.Dni < 10000000 {
		return errors.New("field 'dni' must be 8 digits")
	}
	if usuario.Direccion == "" {
		return errors.New("required field 'Dirección'")
	}
	if usuario.Celular == 0 {
		return errors.New("required field 'celular'")
	}
	if usuario.Celular > 1000000000 || usuario.Celular < 100000000 {
		return errors.New("field 'celular' must be 9 digits")
	}
	return nil
}

func ValidateStudent(estudiante *modelos.Estudiante) error {
	if estudiante.Uni_Pref == "" {
		return errors.New("required field 'uni_pref'")
	}
	if estudiante.Carr_Pref == "" {
		return errors.New("required field 'carr_pref'")
	}
	if estudiante.Nick == "" {
		return errors.New("required field 'nick'")
	}
	if estudiante.UsuariosId == 0 {
		return errors.New("required field 'usuarios_id'")
	}
	return nil
}

func ValidateLogin(usuario *modelos.Usuarios) error {
	if usuario.Email == "" {
		return errors.New("required field 'email'")
	}
	if usuario.Password == "" {
		return errors.New("required field 'password'")
	}
	return nil
}
