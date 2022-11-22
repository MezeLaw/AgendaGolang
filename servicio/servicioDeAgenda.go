package servicio

import (
	"agenda/modelos"
	"errors"
)

type ServicioAgendador struct {
	Contactos map[string]modelos.Contacto
}

func New() ServicioAgendador {
	agenda := map[string]modelos.Contacto{}
	agenda["37869099"] = modelos.Contacto{
		Nombre:   "Meze",
		Apellido: "Abogado",
		DNI:      "37869099",
		Telefono: "1561814200",
	}
	return ServicioAgendador{Contactos: agenda}
}

func (s *ServicioAgendador) AgregarContacto(c modelos.Contacto) (*modelos.Contacto, error) {
	for dni, contacto := range s.Contactos {
		if dni == c.DNI {
			println("El contacto ya existe en la agenda! Su telefono es: ", contacto.Telefono)
			return nil, errors.New("El contacto ya esta agendado")
		}
	}

	s.Contactos[c.DNI] = c
	return &c, nil
}

func (s *ServicioAgendador) ConsultarContato(doc string) (*modelos.Contacto, error) {
	for dni, contacto := range s.Contactos {
		if dni == doc {
			return &contacto, nil
		}
	}
	return nil, errors.New("El contacto no existe en la agenda")
}

func (s *ServicioAgendador) EliminarContacto(doc string) (*string, error) {
	for dni, contacto := range s.Contactos {
		if dni == doc {
			delete(s.Contactos, doc)
			return &contacto.DNI, nil
		}
	}
	return nil, errors.New("El contacto no existe en la agenda")
}
