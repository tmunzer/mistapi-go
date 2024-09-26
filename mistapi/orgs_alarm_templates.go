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

// OrgsAlarmTemplates represents a controller struct.
type OrgsAlarmTemplates struct {
    baseController
}

// NewOrgsAlarmTemplates creates a new instance of OrgsAlarmTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsAlarmTemplates.
func NewOrgsAlarmTemplates(baseController baseController) *OrgsAlarmTemplates {
    orgsAlarmTemplates := OrgsAlarmTemplates{baseController: baseController}
    return &orgsAlarmTemplates
}

// ListOrgAlarmTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Get List of Org Alarm Templates
func (o *OrgsAlarmTemplates) ListOrgAlarmTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.AlarmTemplate],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates", orgId),
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.AlarmTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.AlarmTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgAlarmTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Available rules can be found in Orgs>Consts>getAlarmDefs
// The delivery dict is only required if different from the template delivery settings.
func (o *OrgsAlarmTemplates) CreateOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AlarmTemplate) (
    models.ApiResponse[models.AlarmTemplate],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates", orgId),
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
    
    var result models.AlarmTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnsuppressOrgSuppressedAlarms takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Un-Suppress Suppressed Alarms
func (o *OrgsAlarmTemplates) UnsuppressOrgSuppressedAlarms(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/suppress", orgId),
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ListOrgSuppressedAlarms takes context, orgId, scope as parameters and
// returns an models.ApiResponse with models.ResponseOrgSuppressAlarm data and
// an error if there was an issue with the request or response.
// Get List of Org Alarms Currently Suppressed
func (o *OrgsAlarmTemplates) ListOrgSuppressedAlarms(
    ctx context.Context,
    orgId uuid.UUID,
    scope *models.SuppressedAlarmScopeEnum) (
    models.ApiResponse[models.ResponseOrgSuppressAlarm],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/suppress", orgId),
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
    if scope != nil {
        req.QueryParam("scope", *scope)
    }
    
    var result models.ResponseOrgSuppressAlarm
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOrgSuppressAlarm](decoder)
    return models.NewApiResponse(result, resp), err
}

// SuppressOrgAlarm takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// In certain situations, for example, scheduled maintenance, you may want to suspend alarms to be triggered against Sites for a period of time.
func (o *OrgsAlarmTemplates) SuppressOrgAlarm(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SuppressedAlarm) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/suppress", orgId),
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

// DeleteOrgAlarmTemplate takes context, orgId, alarmtemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Alarm Template
func (o *OrgsAlarmTemplates) DeleteOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/%v", orgId, alarmtemplateId),
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgAlarmTemplate takes context, orgId, alarmtemplateId as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Get Org Alarm Template Details
func (o *OrgsAlarmTemplates) GetOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID) (
    models.ApiResponse[models.AlarmTemplate],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/%v", orgId, alarmtemplateId),
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
    
    var result models.AlarmTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgAlarmTemplate takes context, orgId, alarmtemplateId, body as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Update Org Alarm Template
func (o *OrgsAlarmTemplates) UpdateOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID,
    body *models.AlarmTemplate) (
    models.ApiResponse[models.AlarmTemplate],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/alarmtemplates/%v", orgId, alarmtemplateId),
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
    
    var result models.AlarmTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
