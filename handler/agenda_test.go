package handler

import (
	"agenda/modelos"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgendaHandler_Agendar_Vacio_DNI_Error(t *testing.T) {
	mockedContacto := modelos.Contacto{
		Nombre:   "Fabian",
		Apellido: "Cerchi",
		DNI:      "",
		Telefono: "1121694009",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestAgendaHandler_Agendar_Caracter_Invalido_DNI_Error(t *testing.T) {
	mockedContacto := modelos.Contacto{
		Nombre:   "Fabian",
		Apellido: "Cerchi",
		DNI:      "3832499A",
		Telefono: "1121694009",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)
	assert.Error(t, err)
	assert.EqualValues(t, mockedContacto.DNI, *resp)
}
func TestAgendaHandler_Agendar_Vacio_Nombre_Error(t *testing.T) {
	mockedContacto := modelos.Contacto{
		Nombre:   "",
		Apellido: "Cerchi",
		DNI:      "38324991",
		Telefono: "1121694009",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAgendaHandler_Agendar_Vacio_Telefono_Error(t *testing.T) {
	mockedContacto := modelos.Contacto{
		Nombre:   "Fabian",
		Apellido: "Cerchi",
		DNI:      "38324991",
		Telefono: "",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestAgendaHandler_Agendar_Caracter_Invalido_Telefono_Error(t *testing.T) {
	mockedContacto := modelos.Contacto{
		Nombre:   "Fabian",
		Apellido: "Cerchi",
		DNI:      "38324991",
		Telefono: "11216-94009",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)
	assert.Error(t, err)
	assert.EqualValues(t, mockedContacto.Telefono, *resp)
}

/*
func TestAgendaHandler_Agendar_Contacto_Valido(t *testing.T) {
	//TODO: Ver error en este test  Unexpected call to *handler.MockServicioDeLaAgenda.AgregarContacto
	mockedContacto := modelos.Contacto{
		Nombre:   "Fabian",
		Apellido: "Cerchi",
		DNI:      "38324991",
		Telefono: "1121694009",
	}
	mockService := NewMockServicioDeLaAgenda(gomock.NewController(t))
	handler := NewAgenda(mockService)
	resp, err := handler.Agendar(mockedContacto)
	assert.NoError(t, err)
	assert.EqualValues(t, mockedContacto.DNI, *resp)
}
*/
