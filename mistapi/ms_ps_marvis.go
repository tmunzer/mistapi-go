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

// MSPsMarvis represents a controller struct.
type MSPsMarvis struct {
    baseController
}

// NewMSPsMarvis creates a new instance of MSPsMarvis.
// It takes a baseController as a parameter and returns a pointer to the MSPsMarvis.
func NewMSPsMarvis(baseController baseController) *MSPsMarvis {
    mSPsMarvis := MSPsMarvis{baseController: baseController}
    return &mSPsMarvis
}

// CountMspsMarvisActions takes context, mspId, distinct, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCountMarvisActions data and
// an error if there was an issue with the request or response.
// Count Marvis actions
func (m *MSPsMarvis) CountMspsMarvisActions(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspMarvisSuggestionsCountDistinctEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCountMarvisActions],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/suggestion/count", mspId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseCountMarvisActions
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCountMarvisActions](decoder)
    return models.NewApiResponse(result, resp), err
}
