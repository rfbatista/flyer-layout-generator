package entities

import (
	"algvisual/database"
	"errors"
)

type ComponentType uint8

const (
	ComponentTypeProduto ComponentType = iota
	ComponentTypeCallToAction
	ComponentTypeMarca
	ComponentTypeLogotipo
	ComponentTypeModelo
	ComponentTypeCelebridade
	ComponentTypePlanoDeFundo
	ComponentTypeGrafismo
	ComponentTypeOferta
	ComponentTypeUnknown
	ComponentTypePackshot
	ComponentTypeIllustration
	ComponentTypeTextoLegal
	ComponentTypeIcone
	ComponentTypeContorno
	ComponentTypeTitulo
	ComponentTypePreco
	ComponentTypeBotao
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
	case ComponentTypeContorno:
		return "contorno"
	case ComponentTypeTitulo:
		return "titulo"
	case ComponentTypePreco:
		return "preco"
	case ComponentTypeBotao:
		return "botao"
	case ComponentTypeLogotipo:
		return "logotipo"
	}

	return "desconhecido"
}

func StringToComponentType(s string) ComponentType {
	switch s {
	case "Produto":
		return ComponentTypeProduto
	case "CTA":
		return ComponentTypeCallToAction
	case "marca":
		return ComponentTypeMarca
	case "Modelo":
		return ComponentTypeModelo
	// case "grafismo":
	// 	return ComponentTypeGrafismo
	// case "celebridade":
	// 	return ComponentTypeCelebridade
	case "Oferta":
		return ComponentTypeOferta
	// case "packshot":
	// 	return ComponentTypePackshot
	case "Ilustração":
		return ComponentTypeIllustration
	case "Texto-legal":
		return ComponentTypeTextoLegal
	case "Fundo":
		return ComponentTypePlanoDeFundo
	case "Ícone":
		return ComponentTypeIcone
	case "Contorno":
		return ComponentTypeContorno
	case "Título":
		return ComponentTypeTitulo
	case "Preço":
		return ComponentTypePreco
	case "Botão":
		return ComponentTypeBotao
	case "Logotipo":
		return ComponentTypeLogotipo
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
	case string(database.ComponentTypeIcone):
		return database.ComponentTypeIcone, nil
	case string(database.ComponentTypeContorno):
		return database.ComponentTypeContorno, nil
	case string(database.ComponentTypeTitulo):
		return database.ComponentTypeTitulo, nil
	case string(database.ComponentTypePreco):
		return database.ComponentTypePreco, nil
	case string(database.ComponentTypeBotao):
		return database.ComponentTypeBotao, nil
	case string(database.ComponentTypeLogotipo):
		return database.ComponentTypeLogotipo, nil
	default:
		return "", errors.New("invalid ComponentType")
	}
}

// TODO: Mover para o pacote mapper
func ComponentTypeToDatabaseComponentType(s ComponentType) database.ComponentType {
	switch s {
	case ComponentTypePlanoDeFundo:
		return database.ComponentTypeBackground
	case ComponentTypeMarca:
		return database.ComponentTypeLogotipoMarca
	case ComponentTypeProduto:
		return database.ComponentTypeLogotipoProduto
	case ComponentTypePackshot:
		return database.ComponentTypePackshot
	case ComponentTypeCelebridade:
		return database.ComponentTypeCelebridade
	case ComponentTypeModelo:
		return database.ComponentTypeModelo
	case ComponentTypeIllustration:
		return database.ComponentTypeIlustracao
	case ComponentTypeOferta:
		return database.ComponentTypeOferta
	case ComponentTypeTextoLegal:
		return database.ComponentTypeTextoLegal
	case ComponentTypeGrafismo:
		return database.ComponentTypeGrafismo
	case ComponentTypeCallToAction:
		return database.ComponentTypeTextoCta
	case ComponentTypeIcone:
		return database.ComponentTypeIcone
	case ComponentTypeContorno:
		return database.ComponentTypeContorno
		// TODO criar tipo desconhecido
	case ComponentTypeTitulo:
		return database.ComponentTypeTitulo
	case ComponentTypePreco:
		return database.ComponentTypePreco
	case ComponentTypeLogotipo:
		return database.ComponentTypeLogotipo
	default:
		return database.ComponentTypeGrafismo
	}
}

// TODO: Mover para o pacote mapper
func DatabaseComponentTypeToDomain(s database.ComponentType) ComponentType {
	switch s {
	case database.ComponentTypeBackground:
		return ComponentTypePlanoDeFundo
	case database.ComponentTypeLogotipoMarca:
		return ComponentTypeMarca
	case database.ComponentTypeLogotipoProduto:
		return ComponentTypeProduto
	case database.ComponentTypePackshot:
		return ComponentTypePackshot
	case database.ComponentTypeCelebridade:
		return ComponentTypeCelebridade
	case database.ComponentTypeModelo:
		return ComponentTypeModelo
	case database.ComponentTypeIlustracao:
		return ComponentTypeIllustration
	case database.ComponentTypeOferta:
		return ComponentTypeOferta
	case database.ComponentTypeTextoLegal:
		return ComponentTypeTextoLegal
	case database.ComponentTypeGrafismo:
		return ComponentTypeGrafismo
	case database.ComponentTypeTextoCta:
		return ComponentTypeCallToAction
	case database.ComponentTypeIcone:
		return ComponentTypeIcone
	case database.ComponentTypeContorno:
		return ComponentTypeContorno
	case database.ComponentTypeTitulo:
		return ComponentTypeTitulo
	case database.ComponentTypePreco:
		return ComponentTypePreco
	case database.ComponentTypeLogotipo:
		return ComponentTypeLogotipo
	default:
		return ComponentTypeUnknown
	}
}
