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

// MSPsAdmins represents a controller struct.
type MSPsAdmins struct {
	baseController
}

// NewMSPsAdmins creates a new instance of MSPsAdmins.
// It takes a baseController as a parameter and returns a pointer to the MSPsAdmins.
func NewMSPsAdmins(baseController baseController) *MSPsAdmins {
	mSPsAdmins := MSPsAdmins{baseController: baseController}
	return &mSPsAdmins
}

// ListMspAdmins takes context, mspId as parameters and
// returns an models.ApiResponse with []models.Admin data and
// an error if there was an issue with the request or response.
// Get List of MSP Admins
func (m *MSPsAdmins) ListMspAdmins(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[[]models.Admin],
	error) {
	req := m.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/msps/%v/admins", mspId))
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

	var result []models.Admin
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Admin](decoder)
	return models.NewApiResponse(result, resp), err
}

// RevokeMspAdmin takes context, mspId, adminId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This removes all privileges this admin has against the MSP. This goes deep all the way to the sites
func (m *MSPsAdmins) RevokeMspAdmin(
	ctx context.Context,
	mspId uuid.UUID,
	adminId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/msps/%v/admins/%v", mspId, adminId),
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

// GetMspAdmin takes context, mspId, adminId as parameters and
// returns an models.ApiResponse with models.Admin data and
// an error if there was an issue with the request or response.
// Get MSP Admins
func (m *MSPsAdmins) GetMspAdmin(
	ctx context.Context,
	mspId uuid.UUID,
	adminId uuid.UUID) (
	models.ApiResponse[models.Admin],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/admins/%v", mspId, adminId),
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

	var result models.Admin
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Admin](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateMspAdmin takes context, mspId, adminId, body as parameters and
// returns an models.ApiResponse with models.Admin data and
// an error if there was an issue with the request or response.
// Update MSP Admin
func (m *MSPsAdmins) UpdateMspAdmin(
	ctx context.Context,
	mspId uuid.UUID,
	adminId uuid.UUID,
	body *models.Admin) (
	models.ApiResponse[models.Admin],
	error) {
	req := m.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/msps/%v/admins/%v", mspId, adminId),
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

	var result models.Admin
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Admin](decoder)
	return models.NewApiResponse(result, resp), err
}

// InviteMspAdmin takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Admin data and
// an error if there was an issue with the request or response.
// Invite MSP Admin
// **Note**: An email will also be sent to the user with a link to https://manage.mist.com/verify/invite?token=:token
func (m *MSPsAdmins) InviteMspAdmin(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.Admin) (
	models.ApiResponse[models.Admin],
	error) {
	req := m.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/msps/%v/invites", mspId),
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

	var result models.Admin
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Admin](decoder)
	return models.NewApiResponse(result, resp), err
}

// UninviteMspAdmin takes context, mspId, inviteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete admin invite
func (m *MSPsAdmins) UninviteMspAdmin(
	ctx context.Context,
	mspId uuid.UUID,
	inviteId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/msps/%v/invites/%v", mspId, inviteId),
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

// UpdateMspAdminInvite takes context, mspId, inviteId, body as parameters and
// returns an models.ApiResponse with models.Admin data and
// an error if there was an issue with the request or response.
// Update MSP admin invite
func (m *MSPsAdmins) UpdateMspAdminInvite(
	ctx context.Context,
	mspId uuid.UUID,
	inviteId uuid.UUID,
	body *models.Admin) (
	models.ApiResponse[models.Admin],
	error) {
	req := m.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/msps/%v/invites/%v", mspId, inviteId),
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

	var result models.Admin
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Admin](decoder)
	return models.NewApiResponse(result, resp), err
}
