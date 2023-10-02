package routes

import "net/http"

func getIndicadores(w http.ResponseWriter, r *http.Request) {

	return
}

/*
select e.anio,e.id_uni,e.areas_id ,pe.cursos_id ,pe.temas_id,count(*) total from examens e
		inner join examen_preguntas ep on ep.examens_id  = e.id
		inner join pregunta_examens pe on pe.id = ep.pregunta_examens_id
		where e.tipo_examen  = 'Admision' and e.cantidad_preguntas = e.limite_preguntas
		group by e.anio, e.id_uni,e.areas_id,pe.cursos_id,pe.temas_id
		order by  e.anio, pe.cursos_id,pe.temas_id
*/
