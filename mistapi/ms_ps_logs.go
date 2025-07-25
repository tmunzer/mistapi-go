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

// MSPsLogs represents a controller struct.
type MSPsLogs struct {
    baseController
}

// NewMSPsLogs creates a new instance of MSPsLogs.
// It takes a baseController as a parameter and returns a pointer to the MSPsLogs.
func NewMSPsLogs(baseController baseController) *MSPsLogs {
    mSPsLogs := MSPsLogs{baseController: baseController}
    return &mSPsLogs
}

// ListMspAuditLogs takes context, mspId, siteId, adminName, message, sort, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseLogSearch data and
// an error if there was an issue with the request or response.
// Get list of change logs for the current MSP
func (m *MSPsLogs) ListMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListMspLogsSortEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseLogSearch],
    error) {
    req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/logs")
    req.AppendTemplateParams(mspId)
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
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseLogSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseLogSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountMspAuditLogs takes context, mspId, distinct, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Audit Logs.
func (m *MSPsLogs) CountMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspLogsCountDistinctEnum,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/logs/count")
    req.AppendTemplateParams(mspId)
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}
