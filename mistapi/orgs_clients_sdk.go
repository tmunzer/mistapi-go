package mistapi

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
)

// OrgsClientsSDK represents a controller struct.
type OrgsClientsSDK struct {
	baseController
}

// NewOrgsClientsSDK creates a new instance of OrgsClientsSDK.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsSDK.
func NewOrgsClientsSDK(baseController baseController) *OrgsClientsSDK {
	orgsClientsSDK := OrgsClientsSDK{baseController: baseController}
	return &orgsClientsSDK
}

// UpdateSdkClient takes context, orgId, sdkclientId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Update SDK Client
func (o *OrgsClientsSDK) UpdateSdkClient(
	ctx context.Context,
	orgId uuid.UUID,
	sdkclientId uuid.UUID,
	body *models.NameString) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/sdkclients/%v", orgId, sdkclientId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
