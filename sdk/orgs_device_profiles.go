package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// OrgsDeviceProfiles represents a controller struct.
type OrgsDeviceProfiles struct {
	baseController
}

// NewOrgsDeviceProfiles creates a new instance of OrgsDeviceProfiles.
// It takes a baseController as a parameter and returns a pointer to the OrgsDeviceProfiles.
func NewOrgsDeviceProfiles(baseController baseController) *OrgsDeviceProfiles {
	orgsDeviceProfiles := OrgsDeviceProfiles{baseController: baseController}
	return &orgsDeviceProfiles
}

// ListOrgDeviceProfiles takes context, orgId, mType, limit, page as parameters and
// returns an models.ApiResponse with []models.ListOrgDeviceProfilesResponse data and
// an error if there was an issue with the request or response.
// Get List of Org Device Profiles
func (o *OrgsDeviceProfiles) ListOrgDeviceProfiles(
	ctx context.Context,
	orgId uuid.UUID,
	mType *models.DeviceTypeEnum,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.ListOrgDeviceProfilesResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.ListOrgDeviceProfilesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ListOrgDeviceProfilesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgDeviceProfiles takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.CreateOrgDeviceProfilesResponse data and
// an error if there was an issue with the request or response.
// Create Device Profile
func (o *OrgsDeviceProfiles) CreateOrgDeviceProfiles(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Deviceprofile2) (
	models.ApiResponse[models.CreateOrgDeviceProfilesResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.CreateOrgDeviceProfilesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CreateOrgDeviceProfilesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgDeviceProfile takes context, orgId, deviceprofileId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Device Profile
func (o *OrgsDeviceProfiles) DeleteOrgDeviceProfile(
	ctx context.Context,
	orgId uuid.UUID,
	deviceprofileId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles/%v", orgId, deviceprofileId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// GetOrgDeviceProfile takes context, orgId, deviceprofileId as parameters and
// returns an models.ApiResponse with models.GetOrgDeviceProfileResponse data and
// an error if there was an issue with the request or response.
// Get Org device Profile Details
func (o *OrgsDeviceProfiles) GetOrgDeviceProfile(
	ctx context.Context,
	orgId uuid.UUID,
	deviceprofileId uuid.UUID) (
	models.ApiResponse[models.GetOrgDeviceProfileResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles/%v", orgId, deviceprofileId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.GetOrgDeviceProfileResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrgDeviceProfileResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgDeviceProfile takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.UpdateOrgDeviceProfileResponse data and
// an error if there was an issue with the request or response.
// Update Org Device Profile
func (o *OrgsDeviceProfiles) UpdateOrgDeviceProfile(
	ctx context.Context,
	orgId uuid.UUID,
	deviceprofileId uuid.UUID,
	body *models.Deviceprofile2) (
	models.ApiResponse[models.UpdateOrgDeviceProfileResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles/%v", orgId, deviceprofileId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.UpdateOrgDeviceProfileResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UpdateOrgDeviceProfileResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// AssignOrgDeviceProfile takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Assign Org Device Profile to Devices
func (o *OrgsDeviceProfiles) AssignOrgDeviceProfile(
	ctx context.Context,
	orgId uuid.UUID,
	deviceprofileId uuid.UUID,
	body *models.MacAddresses) (
	models.ApiResponse[models.ResponseAssignSuccess],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles/%v/assign", orgId, deviceprofileId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseAssignSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}

// UnassignOrgDeviceProfiles takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Unassign Org Device Profile from Devices
func (o *OrgsDeviceProfiles) UnassignOrgDeviceProfiles(
	ctx context.Context,
	orgId uuid.UUID,
	deviceprofileId uuid.UUID,
	body *models.MacAddresses) (
	models.ApiResponse[models.ResponseAssignSuccess],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/deviceprofiles/%v/unassign", orgId, deviceprofileId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseAssignSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}
