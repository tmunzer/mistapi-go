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

// OrgsStatsDevices represents a controller struct.
type OrgsStatsDevices struct {
    baseController
}

// NewOrgsStatsDevices creates a new instance of OrgsStatsDevices.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsDevices.
func NewOrgsStatsDevices(baseController baseController) *OrgsStatsDevices {
    orgsStatsDevices := OrgsStatsDevices{baseController: baseController}
    return &orgsStatsDevices
}

// ListOrgDevicesStats takes context, orgId, mType, status, siteId, mac, evpntopoId, evpnUnused, fields, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.ListOrgDevicesStatsResponse data and
// an error if there was an issue with the request or response.
// Get List of Org Devices stats
// This API renders some high-level device stats, pagination is assumed and returned in response header (as the response is an array)
func (o *OrgsStatsDevices) ListOrgDevicesStats(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.DeviceStatusEnum,
    siteId *string,
    mac *string,
    evpntopoId *string,
    evpnUnused *string,
    fields *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.ListOrgDevicesStatsResponse],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/stats/devices", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if evpntopoId != nil {
        req.QueryParam("evpntopo_id", *evpntopoId)
    }
    if evpnUnused != nil {
        req.QueryParam("evpn_unused", *evpnUnused)
    }
    if fields != nil {
        req.QueryParam("fields", *fields)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result []models.ListOrgDevicesStatsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ListOrgDevicesStatsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
