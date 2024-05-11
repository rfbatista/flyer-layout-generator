package shared

import (
	"fmt"
	"regexp"
)

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
	DownloadDesignFileEndpoint    RoutePath = "/api/v1/design/:id/file"

	PageHome           RoutePath = "/"
	PageUploadDesign   RoutePath = "/upload"
	PageProccessDesing RoutePath = "/proccess"
	PageGenerateImage  RoutePath = "/generate"
	PageCreateTemplate RoutePath = "/template"
	PageDefineElements RoutePath = "/elements"

	WebEndpointUploadPhotoshop RoutePath = "/api/v1/web/photoshop"
	WebEndpointProccessDesign  RoutePath = "/api/v1/web/design/:id/proccess"
)

func (m RoutePath) String() string {
	return string(m)
}

func (m RoutePath) Replace(p []string) string {
	regex := regexp.MustCompile(`:[a-zA-Z0-9]+`)
	matches := regex.FindAllString(m.String(), -1)
	replacements := make([]string, len(matches))
	for i := range matches {
		replacements[i] = fmt.Sprintf("%s", p[i])
	}
	result := regex.ReplaceAllStringFunc(m.String(), func(match string) string {
		fmt.Println("ids", p)
		return replacements[findIndex(matches, match)]
	})
	fmt.Println(result)
	return result
}

func findIndex(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
