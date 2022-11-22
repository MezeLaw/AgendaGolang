package main

import "agenda/modelos"

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

	return nil, nil
}

/*
	Pre: recibe el dni a consultar
	Post: retorna el contacto consultado y/o un error
*/
func (a *Agenda) VerContacto(dni string) (*modelos.Contacto, error) {
	return nil, nil
}

/*
	Pre: recibe el dni del contacto a eliminar
	Post: Retorna el doc del contacto eliminado y/o un error segun corresponda
*/
func (a *Agenda) BorrarContacto(dni string) (*string, error) {
	return nil, nil
}
