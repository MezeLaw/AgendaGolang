package main

import (
	"agenda/modelos"
	"errors"
	"strconv"
	"strings"
)

type ServicioDeLaAgenda interface {
	/*
			Pre: recibe un contacto
		    Post: retorna el contacto agendado y/o un error si la transaccion fallo
	*/
	AgregarContacto(contacto modelos.Contacto) (*modelos.Contacto, error)
	/*
			  Pre: recibe un numero de DNI
		 	  Post: retorna el contacto deseado y/o un error
	*/
	ConsultarContato(dni string) (*modelos.Contacto, error)

	/*
	  Pre: recibe un numero de DNI
	  Post: retorna el documento del contacto eliminado y/o un error segun corresponda
	*/

	EliminarContacto(dni string) (*string, error)
}

type Agenda struct {
	servicio ServicioDeLaAgenda
}

func NewAgenda(servicio ServicioDeLaAgenda) Agenda {
	return Agenda{servicio: servicio}
}

/*
	Pre: recibe el contacto a agendar
	Post: retorna el id del contacto agendado y/o un error segun corresponda
*/
func (a *Agenda) Agendar(contacto modelos.Contacto) (*string, error) {
	if contacto.DNI == "" {
		return nil, errors.New("dni perdido")
	}
	dni, err := validadorNumerosStr(contacto.DNI)
	if err != nil {
		return &dni, err
	}
	if contacto.Nombre == "" {
		return nil, errors.New("nombre perdido")
	}
	if contacto.Telefono == "" {
		return nil, errors.New("telefono perdido")
	}
	telefono, err := validadorNumerosStr(contacto.Telefono)
	if err != nil {
		return &telefono, err
	}
	c, err := a.servicio.AgregarContacto(contacto)
	if err != nil {
		return &c.DNI, err
	}
	return &c.DNI, nil
}

/*
	Pre: recibe el dni a consultar
	Post: retorna el contacto consultado y/o un error
*/
func (a *Agenda) VerContacto(dni string) (*modelos.Contacto, error) {
	if dni == "" {
		return nil, errors.New("DNI perdido")
	}
	dni, err := validadorNumerosStr(dni)
	if err != nil {
		return nil, err
	}
	c, err := a.servicio.ConsultarContato(dni)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*
	Pre: recibe el dni del contacto a eliminar
	Post: Retorna el doc del contacto eliminado y/o un error segun corresponda
*/
func (a *Agenda) BorrarContacto(dni string) (*string, error) {
	if dni == "" {
		strings := "DNI Invalido"
		return &strings, errors.New("dni perdido")
	}
	dni, err := validadorNumerosStr(dni)
	if err != nil {
		return &dni, err
	}
	dniEliminado, err := a.servicio.EliminarContacto(dni)
	if err != nil {
		return nil, err
	}
	return dniEliminado, nil
}

func validadorNumerosStr(numero string) (string, error) {
	if strings.Contains(numero, ".") || strings.Contains(numero, ",") {
		return numero, errors.New("No debe contener ',' o '.' ")
	}
	_, err := strconv.Atoi(numero)
	if err != nil {
		return numero, errors.New("DNI ingresado con caracter invalido")
	}
	return numero, nil
}
