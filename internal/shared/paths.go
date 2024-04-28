package shared

type RoutePath string

const (
	EndpointUploadPhotoshop        RoutePath = "/api/v1/photoshop"
	EndpointGetPhotoshopByID       RoutePath = "/api/v1/photoshop/{photoshop_id}"
	EndpointListPhotoshop          RoutePath = "/api/v1/photoshop"
	EndpointListPhotoshopElements  RoutePath = "/api/v1/photoshop/elements"
	EndpointCreateTemplate         RoutePath = "/api/v1/template"
	EndpointListTemplate           RoutePath = "/api/v1/template"
	EndpointListImagesGenerated    RoutePath = "/api/v1/images"
	EndpointRemoveComponent        RoutePath = "/api/v1/photoshop/{photoshopID}/components"
	EndpointCreateComponent        RoutePath = "/api/v1/photoshop/{photoshopID}/component"
	EndpointSetPhotoshopBackground RoutePath = "/api/v1/photoshop/{photoshopID}/background"
)

func (m RoutePath) String() string {
	return string(m)
}
