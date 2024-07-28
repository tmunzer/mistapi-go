package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsGuests represents a controller struct.
type OrgsGuests struct {
    baseController
}

// NewOrgsGuests creates a new instance of OrgsGuests.
// It takes a baseController as a parameter and returns a pointer to the OrgsGuests.
func NewOrgsGuests(baseController baseController) *OrgsGuests {
    orgsGuests := OrgsGuests{baseController: baseController}
    return &orgsGuests
}

// ListOrgGuestAuthorizations takes context, orgId as parameters and
// returns an models.ApiResponse with []models.Guest data and
// an error if there was an issue with the request or response.
// Get List of Org Guest Authorizations
func (o *OrgsGuests) ListOrgGuestAuthorizations(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Guest],
    error) {
    req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v/guests", orgId))
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
    
    var result []models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgGuestAuthorizations takes context, orgId, distinct, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org Authorized Guest
func (o *OrgsGuests) CountOrgGuestAuthorizations(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgGuestsCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/guests/count", orgId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgGuestAuthorization takes context, orgId, wlanId, authMethod, ssid, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseGuestSearch data and
// an error if there was an issue with the request or response.
// Search Authorized Guest
func (o *OrgsGuests) SearchOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId *string,
    authMethod *string,
    ssid *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseGuestSearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/guests/search", orgId),
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
    if wlanId != nil {
        req.QueryParam("wlan_id", *wlanId)
    }
    if authMethod != nil {
        req.QueryParam("auth_method", *authMethod)
    }
    if ssid != nil {
        req.QueryParam("ssid", *ssid)
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
    
    var result models.ResponseGuestSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseGuestSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgGuestAuthorization takes context, orgId, guestMac as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Guest Authorization
func (o *OrgsGuests) DeleteOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/guests/%v", orgId, guestMac),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrgGuestAuthorization takes context, orgId, guestMac as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Get Guest Authorization
func (o *OrgsGuests) GetOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string) (
    models.ApiResponse[models.Guest],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/guests/%v", orgId, guestMac),
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
    
    var result models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgGuestAuthorization takes context, orgId, guestMac, body as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Update Guest Authorization
func (o *OrgsGuests) UpdateOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string,
    body *models.Guest) (
    models.ApiResponse[models.Guest],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/guests/%v", orgId, guestMac),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}
