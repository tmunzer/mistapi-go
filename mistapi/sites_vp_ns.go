package mistapi

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesVPNs represents a controller struct.
type SitesVPNs struct {
	baseController
}

// NewSitesVPNs creates a new instance of SitesVPNs.
// It takes a baseController as a parameter and returns a pointer to the SitesVPNs.
func NewSitesVPNs(baseController baseController) *SitesVPNs {
	sitesVPNs := SitesVPNs{baseController: baseController}
	return &sitesVPNs
}

// ListSiteVpnsDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with []models.Vpn data and
// an error if there was an issue with the request or response.
// VPN object represents an overlay network where gateways can participate in and optionally expose routes to.
func (s *SitesVPNs) ListSiteVpnsDerived(
	ctx context.Context,
	siteId uuid.UUID,
	resolve *bool) (
	models.ApiResponse[[]models.Vpn],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/vpns/derived", siteId),
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
	if resolve != nil {
		req.QueryParam("resolve", *resolve)
	}

	var result []models.Vpn
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Vpn](decoder)
	return models.NewApiResponse(result, resp), err
}
