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
	EndpointCreateTemplate          RoutePath = "/api/v1/templates"
	EndpointListTemplate            RoutePath = "/api/v1/templates"
	EndpointListImagesGenerated     RoutePath = "/api/v1/images"
	EndpointRemoveComponentElements RoutePath = "/api/v1/photoshop/:photoshop_id/components/remove"
	EndpointCreateComponent         RoutePath = "/api/v1/photoshop/:photoshop_id/component"
	EndpointSetPhotoshopBackground  RoutePath = "/api/v1/photoshop/:photoshop_id/background"

	ListComponentByFileIDEndpoint RoutePath = "/api/v1/file/:photoshop_id/components"
	CreateNewDesignEndpoint       RoutePath = "/api/v1/design"
	UploadImageEndpoint           RoutePath = "/api/v1/images"
	DownloadImageEndpoint         RoutePath = "/api/v1/images/:image_name"
	DownloadDesignFileEndpoint    RoutePath = "/api/v1/design/:id/file"

	PageHome                 RoutePath = "/"
	PageHomeCreateRequest    RoutePath = "/request/generation"
	PageUploadDesign         RoutePath = "/upload"
	PageJobs                 RoutePath = "/jobs"
	PageProccessDesing       RoutePath = "/process"
	PageGenerateImage        RoutePath = "/generate"
	PageCreateTemplate       RoutePath = "/templates"
	PageDefineElements       RoutePath = "/elements"
	PageRequestUploadFile    RoutePath = "/request/file-upload"
	PageRequestProcessDesign RoutePath = "/request/file/:id/process"
	PageRequestElements      RoutePath = "/request/design/:id/elements"

	PageRequestElementsCreateComponent    RoutePath = "/request/design/:design_id/component"
	PageRequestElementsDefineBackground   RoutePath = "/request/design/:design_id/background"
	PageRequestElementsRemoveElement      RoutePath = "/request/design/:design_id/component/remove"
	PageRequestUploadSheet                RoutePath = "/request/design/:design_id/sheet-upload"
	PageRequestUploadSheetCreateTemplates RoutePath = "/request/design/:design_id/sheet-upload/create-templates"
	PageRequestTemplatesCreated           RoutePath = "/request/design/:design_id/request/:request_id/templates"
	PageRequestGenerateImages             RoutePath = "/request/design/:design_id/request/:request_id/generate"

	PageRequestResult       RoutePath = "/request/results"
	PageRequestAdjustImages RoutePath = "/request/adjust-image"

	WebEndpointUploadPhotoshop RoutePath = "/api/v1/web/design/file"
	WebEndpointProccessDesign  RoutePath = "/api/v1/web/design/:id/proccess"
)

func (m RoutePath) String() string {
	return string(m)
}

func (m RoutePath) Replace(p []string) string {
	regex := regexp.MustCompile(`:[a-zA-Z0-9\_]+`)
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
