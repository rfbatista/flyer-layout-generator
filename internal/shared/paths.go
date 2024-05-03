package shared

type RoutePath string

const (
	EndpointUploadPhotoshop         RoutePath = "/api/v1/photoshop"
	EndpointGetPhotoshopByID        RoutePath = "/api/v1/photoshop/:photoshop_id"
	EndpointListPhotoshop           RoutePath = "/api/v1/photoshop"
	EndpointListPhotoshopElements   RoutePath = "/api/v1/photoshop/:photoshop_id/elements"
	EndpointCreateTemplate          RoutePath = "/api/v1/template"
	EndpointListTemplate            RoutePath = "/api/v1/template"
	EndpointListImagesGenerated     RoutePath = "/api/v1/images"
	EndpointRemoveComponentElements RoutePath = "/api/v1/photoshop/:photoshop_id/components/remove"
	EndpointCreateComponent         RoutePath = "/api/v1/photoshop/:photoshop_id/component"
	EndpointSetPhotoshopBackground  RoutePath = "/api/v1/photoshop/:photoshop_id/background"

	ListComponentByFileIDEndpoint RoutePath = "/api/v1/file/:photoshop_id/components"
	CreateNewDesignEndpoint       RoutePath = "/api/v1/design"
	UploadImageEndpoint           RoutePath = "/api/v1/images"
	DownloadImageEndpoint         RoutePath = "/api/v1/images/:image_name"
)

func (m RoutePath) String() string {
	return string(m)
}
