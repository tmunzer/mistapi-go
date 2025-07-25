// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsOtherDevices represents a controller struct.
type OrgsStatsOtherDevices struct {
    baseController
}

// NewOrgsStatsOtherDevices creates a new instance of OrgsStatsOtherDevices.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsOtherDevices.
func NewOrgsStatsOtherDevices(baseController baseController) *OrgsStatsOtherDevices {
    orgsStatsOtherDevices := OrgsStatsOtherDevices{baseController: baseController}
    return &orgsStatsOtherDevices
}

// GetOrgOtherDeviceStats takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with models.StatsDeviceOther data and
// an error if there was an issue with the request or response.
// Get Otherdevice Stats
func (o *OrgsStatsOtherDevices) GetOrgOtherDeviceStats(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.StatsDeviceOther],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/otherdevices/%v")
    req.AppendTemplateParams(orgId, deviceMac)
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
    
    var result models.StatsDeviceOther
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.StatsDeviceOther](decoder)
    return models.NewApiResponse(result, resp), err
}
