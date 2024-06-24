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
	ComponentTypePackshot
	ComponentTypeIllustration
	ComponentTypeTextoLegal
)

func (c ComponentType) ToString() string {
	switch c {
	case ComponentTypeProduto:
		return "produto"
	case ComponentTypeCallToAction:
		return "cta"
	case ComponentTypeMarca:
		return "marca"
	case ComponentTypePlanoDeFundo:
		return "plano-de-fundo"
	case ComponentTypeModelo:
		return "modelo"
	case ComponentTypeGrafismo:
		return "grafismo"
	case ComponentTypeCelebridade:
		return "celebridade"
	case ComponentTypeOferta:
		return "oferta"
	case ComponentTypePackshot:
		return "packshot"
	case ComponentTypeIllustration:
		return "ilustracao"
	case ComponentTypeTextoLegal:
		return "texto-legal"
	}

	return "desconhecido"
}

func StringToComponentType(s string) ComponentType {
	switch s {
	case "produto":
		return ComponentTypeProduto
	case "cta":
		return ComponentTypeCallToAction
	case "marca":
		return ComponentTypeMarca
	case "modelo":
		return ComponentTypeModelo
	case "grafismo":
		return ComponentTypeGrafismo
	case "celebridade":
		return ComponentTypeCelebridade
	case "oferta":
		return ComponentTypeOferta
	case "packshot":
		return ComponentTypePackshot
	case "ilustracao":
		return ComponentTypeIllustration
	case "texto-legal":
		return ComponentTypeTextoLegal
	case "plano-de-fundo":
		return ComponentTypePlanoDeFundo
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

func ComponentTypeToDatabaseComponentType(s ComponentType) (database.ComponentType, error) {
	switch s {
	case ComponentTypePlanoDeFundo:
		return database.ComponentTypeBackground, nil
	case ComponentTypeMarca:
		return database.ComponentTypeLogotipoMarca, nil
	case ComponentTypeProduto:
		return database.ComponentTypeLogotipoProduto, nil
	case ComponentTypePackshot:
		return database.ComponentTypePackshot, nil
	case ComponentTypeCelebridade:
		return database.ComponentTypeCelebridade, nil
	case ComponentTypeModelo:
		return database.ComponentTypeModelo, nil
	case ComponentTypeIllustration:
		return database.ComponentTypeIlustracao, nil
	case ComponentTypeOferta:
		return database.ComponentTypeOferta, nil
	case ComponentTypeTextoLegal:
		return database.ComponentTypeTextoLegal, nil
	case ComponentTypeGrafismo:
		return database.ComponentTypeGrafismo, nil
	case ComponentTypeCallToAction:
		return database.ComponentTypeTextoCta, nil
	default:
		return "", errors.New("invalid ComponentType")
	}
}
