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

// OrgsMaps represents a controller struct.
type OrgsMaps struct {
    baseController
}

// NewOrgsMaps creates a new instance of OrgsMaps.
// It takes a baseController as a parameter and returns a pointer to the OrgsMaps.
func NewOrgsMaps(baseController baseController) *OrgsMaps {
    orgsMaps := OrgsMaps{baseController: baseController}
    return &orgsMaps
}

// ImportOrgMaps takes context, orgId, autoDeviceprofileAssignment, csv, file, json as parameters and
// returns an models.ApiResponse with models.ResponseMapImport data and
// an error if there was an issue with the request or response.
// Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches
// ### CSV File Format
// ```csv
// Vendor AP name,Mist AP Mac
// US Office AP-2 - 5c:5b:35:00:00:02,5c5b35000002
// ```
func (o *OrgsMaps) ImportOrgMaps(
    ctx context.Context,
    orgId uuid.UUID,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapOrgImportFileJson) (
    models.ApiResponse[models.ResponseMapImport],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/maps/import", orgId),
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
    formFields := []https.FormParam{}
    if autoDeviceprofileAssignment != nil {
        auto_deviceprofile_assignmentParam := https.FormParam{Key: "auto_deviceprofile_assignment", Value: *autoDeviceprofileAssignment, Headers: http.Header{}}
        formFields = append(formFields, auto_deviceprofile_assignmentParam)
    }
    if csv != nil {
        csvParam := https.FormParam{Key: "csv", Value: *csv, Headers: http.Header{}}
        formFields = append(formFields, csvParam)
    }
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    var result models.ResponseMapImport
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMapImport](decoder)
    return models.NewApiResponse(result, resp), err
}

// ImportOrgMapToSite takes context, orgId, siteName, autoDeviceprofileAssignment, csv, file, json as parameters and
// returns an models.ApiResponse with models.ResponseMapImport data and
// an error if there was an issue with the request or response.
// Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches
// #### Request
// ```
// "json": a JSON string describing your upload
// "file": a binary file
// ```
func (o *OrgsMaps) ImportOrgMapToSite(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    models.ApiResponse[models.ResponseMapImport],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/sites/%v/maps/import", orgId, siteName),
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
    formFields := []https.FormParam{}
    if autoDeviceprofileAssignment != nil {
        auto_deviceprofile_assignmentParam := https.FormParam{Key: "auto_deviceprofile_assignment", Value: *autoDeviceprofileAssignment, Headers: http.Header{}}
        formFields = append(formFields, auto_deviceprofile_assignmentParam)
    }
    if csv != nil {
        csvParam := https.FormParam{Key: "csv", Value: *csv, Headers: http.Header{}}
        formFields = append(formFields, csvParam)
    }
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    var result models.ResponseMapImport
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMapImport](decoder)
    return models.NewApiResponse(result, resp), err
}
