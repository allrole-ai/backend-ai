package allrole

import (
	"github.com/allrole-ai/backend-ai/routes"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("WebHook", routes.URL)
}
