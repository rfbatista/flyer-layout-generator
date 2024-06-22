package entities

import (
	"algvisual/internal/database"
	"errors"
)

type ComponentType uint8

const (
	ComponentTypeProduto ComponentType = iota
	ComponentTypeCallToAction
	ComponentTypeMarca
	ComponentTypeModelo
	ComponentTypeCelebridade
	ComponentTypePlanoDeFundo
	ComponentTypeGrafismo
	ComponentTypeOferta
	ComponentTypeUnknown
)

func (c ComponentType) ToString() string {
	switch c {
	case ComponentTypeProduto:
		return "produto"
	case ComponentTypeCallToAction:
		return "cta"
	case ComponentTypeMarca:
		return "marca"
	case ComponentTypeModelo:
		return "modelo"
	case ComponentTypeGrafismo:
		return "grafismo"
	case ComponentTypeCelebridade:
		return "celebridade"
	case ComponentTypeOferta:
		return "oferta"
	}

	return "desconhecido"
}

func StringToComponentType(s string) ComponentType {
	switch s {
	case "produto":
		return ComponentTypeProduto
	case "cta":
		return ComponentTypeProduto
	case "marca":
		return ComponentTypeProduto
	case "modelo":
		return ComponentTypeProduto
	case "grafismo":
		return ComponentTypeProduto
	}
	return ComponentTypeUnknown
}

func StringToDatabaseComponentType(s string) (database.ComponentType, error) {
	switch s {
	case string(database.ComponentTypeBackground):
		return database.ComponentTypeBackground, nil
	case string(database.ComponentTypeLogotipoMarca):
		return database.ComponentTypeLogotipoMarca, nil
	case string(database.ComponentTypeLogotipoProduto):
		return database.ComponentTypeLogotipoProduto, nil
	case string(database.ComponentTypePackshot):
		return database.ComponentTypePackshot, nil
	case string(database.ComponentTypeCelebridade):
		return database.ComponentTypeCelebridade, nil
	case string(database.ComponentTypeModelo):
		return database.ComponentTypeModelo, nil
	case string(database.ComponentTypeIlustracao):
		return database.ComponentTypeIlustracao, nil
	case string(database.ComponentTypeOferta):
		return database.ComponentTypeOferta, nil
	case string(database.ComponentTypeTextoLegal):
		return database.ComponentTypeTextoLegal, nil
	case string(database.ComponentTypeGrafismo):
		return database.ComponentTypeGrafismo, nil
	case string(database.ComponentTypeTextoCta):
		return database.ComponentTypeTextoCta, nil
	default:
		return "", errors.New("invalid ComponentType")
	}
}
