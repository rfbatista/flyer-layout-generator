package shared

type RoutePath string

const (
	EndpointUploadPhotoshop RoutePath = "/api/v1/photoshop"
	EndpointListPhotoshop   RoutePath = "/api/v1/photoshop"
)

func (m RoutePath) String() string {
	return string(m)
}
