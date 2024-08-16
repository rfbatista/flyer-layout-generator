// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ComponentType string

const (
	ComponentTypeBackground      ComponentType = "background"
	ComponentTypeLogotipoMarca   ComponentType = "logotipo_marca"
	ComponentTypeLogotipoProduto ComponentType = "logotipo_produto"
	ComponentTypePackshot        ComponentType = "packshot"
	ComponentTypeCelebridade     ComponentType = "celebridade"
	ComponentTypeModelo          ComponentType = "modelo"
	ComponentTypeIlustracao      ComponentType = "ilustracao"
	ComponentTypeOferta          ComponentType = "oferta"
	ComponentTypeTextoLegal      ComponentType = "texto_legal"
	ComponentTypeGrafismo        ComponentType = "grafismo"
	ComponentTypeIcone           ComponentType = "icone"
	ComponentTypeContorno        ComponentType = "contorno"
	ComponentTypeTitulo          ComponentType = "titulo"
	ComponentTypePreco           ComponentType = "preco"
	ComponentTypeBotao           ComponentType = "botao"
	ComponentTypeTextoCta        ComponentType = "texto_cta"
)

func (e *ComponentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ComponentType(s)
	case string:
		*e = ComponentType(s)
	default:
		return fmt.Errorf("unsupported scan type for ComponentType: %T", src)
	}
	return nil
}

type NullComponentType struct {
	ComponentType ComponentType `json:"component_type"`
	Valid         bool          `json:"valid"` // Valid is true if ComponentType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullComponentType) Scan(value interface{}) error {
	if value == nil {
		ns.ComponentType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ComponentType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullComponentType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ComponentType), nil
}

type DesignAssetType string

const (
	DesignAssetTypeText        DesignAssetType = "text"
	DesignAssetTypeSmartobject DesignAssetType = "smartobject"
	DesignAssetTypeShape       DesignAssetType = "shape"
	DesignAssetTypePixel       DesignAssetType = "pixel"
	DesignAssetTypeGroup       DesignAssetType = "group"
	DesignAssetTypeUnknown     DesignAssetType = "unknown"
)

func (e *DesignAssetType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DesignAssetType(s)
	case string:
		*e = DesignAssetType(s)
	default:
		return fmt.Errorf("unsupported scan type for DesignAssetType: %T", src)
	}
	return nil
}

type NullDesignAssetType struct {
	DesignAssetType DesignAssetType `json:"design_asset_type"`
	Valid           bool            `json:"valid"` // Valid is true if DesignAssetType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDesignAssetType) Scan(value interface{}) error {
	if value == nil {
		ns.DesignAssetType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DesignAssetType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDesignAssetType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.DesignAssetType), nil
}

type Roles string

const (
	RolesAdmin Roles = "admin"
	RolesColab Roles = "colab"
	RolesGm    Roles = "gm"
)

func (e *Roles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Roles(s)
	case string:
		*e = Roles(s)
	default:
		return fmt.Errorf("unsupported scan type for Roles: %T", src)
	}
	return nil
}

type NullRoles struct {
	Roles Roles `json:"roles"`
	Valid bool  `json:"valid"` // Valid is true if Roles is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoles) Scan(value interface{}) error {
	if value == nil {
		ns.Roles, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Roles.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoles) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Roles), nil
}

type TemplateType string

const (
	TemplateTypeSlots      TemplateType = "slots"
	TemplateTypeDistortion TemplateType = "distortion"
)

func (e *TemplateType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TemplateType(s)
	case string:
		*e = TemplateType(s)
	default:
		return fmt.Errorf("unsupported scan type for TemplateType: %T", src)
	}
	return nil
}

type NullTemplateType struct {
	TemplateType TemplateType `json:"template_type"`
	Valid        bool         `json:"valid"` // Valid is true if TemplateType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTemplateType) Scan(value interface{}) error {
	if value == nil {
		ns.TemplateType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TemplateType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTemplateType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TemplateType), nil
}

type Advertiser struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	CompanyID pgtype.Int4      `json:"company_id"`
}

type Client struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	CompanyID pgtype.Int4      `json:"company_id"`
}

type CompaniesApiCredential struct {
	ID        int64            `json:"id"`
	Name      pgtype.Text      `json:"name"`
	ApiKey    string           `json:"api_key"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	CompanyID pgtype.Int4      `json:"company_id"`
}

type Company struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Enabled   pgtype.Bool      `json:"enabled"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

type Design struct {
	ID             int32            `json:"id"`
	Name           string           `json:"name"`
	ImageUrl       pgtype.Text      `json:"image_url"`
	LayoutID       pgtype.Int4      `json:"layout_id"`
	ProjectID      pgtype.Int4      `json:"project_id"`
	ImageExtension pgtype.Text      `json:"image_extension"`
	FileUrl        pgtype.Text      `json:"file_url"`
	FileExtension  pgtype.Text      `json:"file_extension"`
	Width          pgtype.Int4      `json:"width"`
	Height         pgtype.Int4      `json:"height"`
	IsProccessed   pgtype.Bool      `json:"is_proccessed"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	CompanyID      pgtype.Int4      `json:"company_id"`
}

type DesignAsset struct {
	ID            int32               `json:"id"`
	ProjectID     pgtype.Int4         `json:"project_id"`
	DesignID      pgtype.Int4         `json:"design_id"`
	AlternativeTo pgtype.Int4         `json:"alternative_to"`
	Name          string              `json:"name"`
	Width         pgtype.Int4         `json:"width"`
	Type          NullDesignAssetType `json:"type"`
	AssetUrl      pgtype.Text         `json:"asset_url"`
	AssetPath     pgtype.Text         `json:"asset_path"`
	Height        pgtype.Int4         `json:"height"`
	CreatedAt     pgtype.Timestamp    `json:"created_at"`
	UpdatedAt     pgtype.Timestamp    `json:"updated_at"`
}

type DesignAssetsProperty struct {
	ID        int32            `json:"id"`
	AssetID   pgtype.Int4      `json:"asset_id"`
	Key       string           `json:"key"`
	Value     string           `json:"value"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Image struct {
	ID          int64            `json:"id"`
	Url         string           `json:"url"`
	PhotoshopID pgtype.Int4      `json:"photoshop_id"`
	TemplateID  pgtype.Int4      `json:"template_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type Layout struct {
	ID         int64            `json:"id"`
	DesignID   pgtype.Int4      `json:"design_id"`
	RequestID  pgtype.Int4      `json:"request_id"`
	IsOriginal pgtype.Bool      `json:"is_original"`
	ImageUrl   pgtype.Text      `json:"image_url"`
	Width      pgtype.Int4      `json:"width"`
	Height     pgtype.Int4      `json:"height"`
	Data       pgtype.Text      `json:"data"`
	Stages     []string         `json:"stages"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
	CompanyID  pgtype.Int4      `json:"company_id"`
}

type LayoutComponent struct {
	ID         int32             `json:"id"`
	LayoutID   int32             `json:"layout_id"`
	DesignID   int32             `json:"design_id"`
	Width      pgtype.Int4       `json:"width"`
	Height     pgtype.Int4       `json:"height"`
	IsOriginal pgtype.Bool       `json:"is_original"`
	Color      pgtype.Text       `json:"color"`
	Type       NullComponentType `json:"type"`
	Xi         pgtype.Int4       `json:"xi"`
	Xii        pgtype.Int4       `json:"xii"`
	Yi         pgtype.Int4       `json:"yi"`
	Yii        pgtype.Int4       `json:"yii"`
	BboxXi     pgtype.Int4       `json:"bbox_xi"`
	BboxXii    pgtype.Int4       `json:"bbox_xii"`
	BboxYi     pgtype.Int4       `json:"bbox_yi"`
	BboxYii    pgtype.Int4       `json:"bbox_yii"`
	Priority   pgtype.Int4       `json:"priority"`
	InnerXi    pgtype.Int4       `json:"inner_xi"`
	InnerXii   pgtype.Int4       `json:"inner_xii"`
	InnerYi    pgtype.Int4       `json:"inner_yi"`
	InnerYii   pgtype.Int4       `json:"inner_yii"`
	CreatedAt  pgtype.Timestamp  `json:"created_at"`
}

type LayoutElement struct {
	ID             int32            `json:"id"`
	DesignID       int32            `json:"design_id"`
	LayoutID       int32            `json:"layout_id"`
	ComponentID    pgtype.Int4      `json:"component_id"`
	AssetID        int32            `json:"asset_id"`
	Name           pgtype.Text      `json:"name"`
	LayerID        pgtype.Text      `json:"layer_id"`
	Text           pgtype.Text      `json:"text"`
	Xi             pgtype.Int4      `json:"xi"`
	Xii            pgtype.Int4      `json:"xii"`
	Yi             pgtype.Int4      `json:"yi"`
	Yii            pgtype.Int4      `json:"yii"`
	InnerXi        pgtype.Int4      `json:"inner_xi"`
	InnerXii       pgtype.Int4      `json:"inner_xii"`
	InnerYi        pgtype.Int4      `json:"inner_yi"`
	InnerYii       pgtype.Int4      `json:"inner_yii"`
	Width          pgtype.Int4      `json:"width"`
	Height         pgtype.Int4      `json:"height"`
	IsGroup        pgtype.Bool      `json:"is_group"`
	GroupID        pgtype.Int4      `json:"group_id"`
	Level          pgtype.Int4      `json:"level"`
	Kind           pgtype.Text      `json:"kind"`
	ImageUrl       pgtype.Text      `json:"image_url"`
	ImageExtension pgtype.Text      `json:"image_extension"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}

type LayoutRequest struct {
	ID        int64            `json:"id"`
	LayoutID  pgtype.Int4      `json:"layout_id"`
	DesignID  pgtype.Int4      `json:"design_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	Log       pgtype.Text      `json:"log"`
	Config    pgtype.Text      `json:"config"`
	Done      int32            `json:"done"`
	Total     pgtype.Int4      `json:"total"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	CompanyID pgtype.Int4      `json:"company_id"`
}

type LayoutRequestsJob struct {
	ID         int64            `json:"id"`
	LayoutID   pgtype.Int4      `json:"layout_id"`
	DesignID   pgtype.Int4      `json:"design_id"`
	RequestID  pgtype.Int4      `json:"request_id"`
	TemplateID pgtype.Int4      `json:"template_id"`
	Status     pgtype.Text      `json:"status"`
	ImageUrl   pgtype.Text      `json:"image_url"`
	StartedAt  pgtype.Timestamp `json:"started_at"`
	FinishedAt pgtype.Timestamp `json:"finished_at"`
	ErrorAt    pgtype.Timestamp `json:"error_at"`
	StoppedAt  pgtype.Timestamp `json:"stopped_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	Config     pgtype.Text      `json:"config"`
	Log        pgtype.Text      `json:"log"`
}

type Project struct {
	ID           int64            `json:"id"`
	ClientID     pgtype.Int4      `json:"client_id"`
	AdvertiserID pgtype.Int4      `json:"advertiser_id"`
	Briefing     pgtype.Text      `json:"briefing"`
	UseAi        pgtype.Bool      `json:"use_ai"`
	Name         string           `json:"name"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
	CompanyID    pgtype.Int4      `json:"company_id"`
}

type Template struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	RequestID pgtype.Text      `json:"request_id"`
	ProjectID pgtype.Int4      `json:"project_id"`
	Width     pgtype.Int4      `json:"width"`
	Height    pgtype.Int4      `json:"height"`
	SlotsX    pgtype.Int4      `json:"slots_x"`
	SlotsY    pgtype.Int4      `json:"slots_y"`
	MaxSlotsX pgtype.Int4      `json:"max_slots_x"`
	MaxSlotsY pgtype.Int4      `json:"max_slots_y"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	CompanyID pgtype.Int4      `json:"company_id"`
}

type TemplatesDistortion struct {
	ID         int32            `json:"id"`
	X          pgtype.Int4      `json:"x"`
	Y          pgtype.Int4      `json:"y"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
	TemplateID int32            `json:"template_id"`
}

type TemplatesSlot struct {
	ID         int32            `json:"id"`
	Xi         pgtype.Int4      `json:"xi"`
	Yi         pgtype.Int4      `json:"yi"`
	Width      pgtype.Int4      `json:"width"`
	Height     pgtype.Int4      `json:"height"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
	TemplateID int32            `json:"template_id"`
}

type User struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Email     pgtype.Text      `json:"email"`
	Username  pgtype.Text      `json:"username"`
	Role      NullRoles        `json:"role"`
	CompanyID pgtype.Int4      `json:"company_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}
