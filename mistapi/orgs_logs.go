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

// OrgsLogs represents a controller struct.
type OrgsLogs struct {
    baseController
}

// NewOrgsLogs creates a new instance of OrgsLogs.
// It takes a baseController as a parameter and returns a pointer to the OrgsLogs.
func NewOrgsLogs(baseController baseController) *OrgsLogs {
    orgsLogs := OrgsLogs{baseController: baseController}
    return &orgsLogs
}

// ListOrgLogs takes context, orgId, siteId, adminName, message, sort, start, end, limit, page, duration as parameters and
// returns an models.ApiResponse with models.ResponseLogSearch data and
// an error if there was an issue with the request or response.
// Get List of change logs for the current Org
func (o *OrgsLogs) ListOrgLogs(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListOrgLogsSortEnum,
    start *int,
    end *int,
    limit *int,
    page *int,
    duration *string) (
    models.ApiResponse[models.ResponseLogSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v/logs", orgId))
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if adminName != nil {
        req.QueryParam("admin_name", *adminName)
    }
    if message != nil {
        req.QueryParam("message", *message)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.ResponseLogSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseLogSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgLogs takes context, orgId, distinct, adminId, adminName, siteId, message, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Audit Logs
func (o *OrgsLogs) CountOrgLogs(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgLogsCountDistinctEnum,
    adminId *string,
    adminName *string,
    siteId *string,
    message *string,
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
      fmt.Sprintf("/api/v1/orgs/%v/logs/count", orgId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if adminId != nil {
        req.QueryParam("admin_id", *adminId)
    }
    if adminName != nil {
        req.QueryParam("admin_name", *adminName)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if message != nil {
        req.QueryParam("message", *message)
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
