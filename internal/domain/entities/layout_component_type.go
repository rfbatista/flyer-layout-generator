package entities

import (
	"algvisual/internal/infrastructure/database"
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
	ComponentTypeFoto
	ComponentTypeTexto
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
	case ComponentTypeIcone:
		return "icone"
	case ComponentTypeFoto:
		return "foto"
	case ComponentTypeTexto:
		return "texto"
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
	// case "grafismo":
	// 	return ComponentTypeGrafismo
	// case "celebridade":
	// 	return ComponentTypeCelebridade
	case "oferta":
		return ComponentTypeOferta
	// case "packshot":
	// 	return ComponentTypePackshot
	case "ilustracao":
		return ComponentTypeIllustration
	case "text-legal":
		return ComponentTypeTextoLegal
	case "plano-de-fundo":
		return ComponentTypePlanoDeFundo
	case "icone":
		return ComponentTypeIcone
	case "contorno":
		return ComponentTypeContorno
	case "titulo":
		return ComponentTypeTitulo
	case "preco":
		return ComponentTypePreco
	case "botao":
		return ComponentTypeBotao
	case "logotipo":
		return ComponentTypeLogotipo
	case "foto":
		return ComponentTypeFoto
	case "texto":
		return ComponentTypeTexto
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
	case string(database.ComponentTypeFoto):
		return database.ComponentTypeFoto, nil
	case string(database.ComponentTypeTexto):
		return database.ComponentTypeTexto, nil
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
	case ComponentTypeFoto:
		return database.ComponentTypeFoto
	case ComponentTypeTexto:
		return database.ComponentTypeTexto
	case ComponentTypeBotao:
		return database.ComponentTypeBotao
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
	case database.ComponentTypeFoto:
		return ComponentTypeFoto
	case database.ComponentTypeTexto:
		return ComponentTypeTexto
	case database.ComponentTypeBotao:
		return ComponentTypeBotao
	default:
		return ComponentTypeUnknown
	}
}
