package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MadMaxMR/backend-go/modelos"
)

func ValidateBody(req *http.Request, modelo interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("ValidateBOdy ReadAll")
		return err
	}
	err = json.Unmarshal(body, modelo)
	if err != nil {
		fmt.Println("ValidateBOdy Unmarshal")
		return err
	}

	return nil
}

func ValidateBody2(req *http.Request, modelo1, modelo2 interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("ValidateBOdy ReadAll")
		return err
	}
	err = json.Unmarshal(body, modelo1)
	if err != nil {
		fmt.Println("ValidateBOdy1 Unmarshal")
		return err
	}
	err = json.Unmarshal(body, modelo2)
	if err != nil {
		fmt.Println("ValidateBody2 Unmarshal")
		return err
	}
	return nil
}

func ValidateCurso(curso *modelos.Cursos) error {

	if curso.Nombre_Curso == "" {
		return errors.New("required field 'nombre_curso'")
	}

	return nil
}

func ValidateTema(tema *modelos.Temas) error {
	if tema.Id_Curso == 0 {
		return errors.New("required field 'Id_Curso'")
	}
	if tema.Nombre_Tema == "" {
		return errors.New("required field 'Nombre_Tema'")
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

func ValidateRecovery(usuario *modelos.Usuarios) error {
	if usuario.Email == "" {
		return errors.New("required field 'email'")
	}
	return nil
}

func ValidateReset(usuario *modelos.Usuarios) error {
	if usuario.Password == "" {
		return errors.New("required field 'password'")
	}
	return nil
}

func ValidateVideo(video *modelos.Videos) error {
	if video.TemasId == 0 {
		return errors.New("required field 'TemasId'")
	}
	if video.Link == "" {
		return errors.New("required field 'Link'")
	}
	if video.Titulo == "" {
		return errors.New("required field 'Titulo'")
	}
	if video.Duracion == "" {
		return errors.New("required field 'Duracion'")
	}
	return nil
}

func ValidateEvaluaciones(eval *modelos.Evaluaciones) error {
	if eval.TemasID == 0 {
		return errors.New("required field 'TemasId'")
	}
	if eval.Nivel == 0 {
		return errors.New("required field 'Nivel'")
	}
	if eval.PdfLink == "" {
		return errors.New("required field 'Pdf_link'")
	}
	return nil
}

func ValidateExamen(examen *modelos.Examens) error {
	if examen.Id_Uni == "" {
		return errors.New("required field 'id_uni'")
	}
	if examen.AreasId == "" {
		return errors.New("required field 'id_area'")
	}
	if examen.Descripcion == "" {
		return errors.New("required field 'descripcion'")
	}
	return nil
}

func ValidatePreguntaExamen(pregunta *modelos.PreguntaExamens) error {
	if pregunta.Enunciado1 == "" {
		return errors.New("required field 'enunciado1'")
	}
	if pregunta.RespuestaExs[0].Respuesta == "" {
		return errors.New("required field 'respuestasExs.Respuesta'")
	}
	return nil
}
