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

// OrgsJSI represents a controller struct.
type OrgsJSI struct {
	baseController
}

// NewOrgsJSI creates a new instance of OrgsJSI.
// It takes a baseController as a parameter and returns a pointer to the OrgsJSI.
func NewOrgsJSI(baseController baseController) *OrgsJSI {
	orgsJSI := OrgsJSI{baseController: baseController}
	return &orgsJSI
}

// ListOrgJsiDevices takes context, orgId, limit, page, model, serial, mac as parameters and
// returns an models.ApiResponse with []models.JseDevice data and
// an error if there was an issue with the request or response.
// Get List of Org devices that connected to JSI
func (o *OrgsJSI) ListOrgJsiDevices(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int,
	model *string,
	serial *string,
	mac *string) (
	models.ApiResponse[[]models.JseDevice],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/devices")
	req.AppendTemplateParams(orgId)
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
	if model != nil {
		req.QueryParam("model", *model)
	}
	if serial != nil {
		req.QueryParam("serial", *serial)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}

	var result []models.JseDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.JseDevice](decoder)
	return models.NewApiResponse(result, resp), err
}

// AdoptOrgJsiDevice takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseDeviceConfigCmd data and
// an error if there was an issue with the request or response.
// Adopt JSI devices
func (o *OrgsJSI) AdoptOrgJsiDevice(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.ResponseDeviceConfigCmd],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		"/api/v1/orgs/%v/jsi/devices/outbound_ssh_cmd",
	)
	req.AppendTemplateParams(orgId)
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

	var result models.ResponseDeviceConfigCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceConfigCmd](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgJsiDeviceShellSession takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Create Shell Session
func (o *OrgsJSI) CreateOrgJsiDeviceShellSession(
	ctx context.Context,
	orgId uuid.UUID,
	deviceMac string) (
	models.ApiResponse[models.WebsocketSessionWithUrl],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/jsi/devices/%v/shell")
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

	var result models.WebsocketSessionWithUrl
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOrgJsiPastPurchases takes context, orgId, limit, page, model, serial as parameters and
// returns an models.ApiResponse with []models.JsInventoryItem data and
// an error if there was an issue with the request or response.
// This gets all devices purchased from the accounts associated with the Org
// * Fetch Install base devices for all linked accounts and associated account of the linked accounts.
// * The primary and the associated account ids will be queries from SFDC by passing the linked account
// * Returns only the device centric details of the Install base device. No customer specific information will be returned.
func (o *OrgsJSI) ListOrgJsiPastPurchases(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int,
	model *string,
	serial *string) (
	models.ApiResponse[[]models.JsInventoryItem],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/inventory")
	req.AppendTemplateParams(orgId)
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
		"400": {Message: "Bad Request - no Juniper Account Linked", Unmarshaller: errors.NewResponseDetailString},
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
	if model != nil {
		req.QueryParam("model", *model)
	}
	if serial != nil {
		req.QueryParam("serial", *serial)
	}

	var result []models.JsInventoryItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.JsInventoryItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgJsiAssetsAndContracts takes context, orgId, distinct, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count devices purchased from the accounts associated with the Org
func (o *OrgsJSI) CountOrgJsiAssetsAndContracts(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.JsiInventoryCountDistinctEnum,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/inventory/count")
	req.AppendTemplateParams(orgId)
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
		"400": {Message: "Bad Request - no Juniper Account Linked", Unmarshaller: errors.NewResponseDetailString},
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

// SearchOrgJsiAssetsAndContracts takes context, orgId, claimed, model, serial, sku, status, warrantyType, eolAfter, eolBefore, eosAfter, eosBefore, versionEosAfter, versionEosBefore, hasSupport, sirtId, pbnId, text, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.JsInventorySearch data and
// an error if there was an issue with the request or response.
// This gets all devices purchased from the accounts associated with the Org
// * Fetch Install base devices for all linked accounts and associated account of the linked accounts.
// * The primary and the associated account ids will be queries from SFDC by passing the linked account
// * Returns only the device centric details of the Install base device. No customer specific information will be returned.
func (o *OrgsJSI) SearchOrgJsiAssetsAndContracts(
	ctx context.Context,
	orgId uuid.UUID,
	claimed *bool,
	model *string,
	serial *string,
	sku *string,
	status *models.DeviceStatusEnum,
	warrantyType *models.JsiWarrantyTypeEnum,
	eolAfter *string,
	eolBefore *string,
	eosAfter *string,
	eosBefore *string,
	versionEosAfter *string,
	versionEosBefore *string,
	hasSupport *bool,
	sirtId *string,
	pbnId *string,
	text *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.JsInventorySearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/inventory/search")
	req.AppendTemplateParams(orgId)
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
		"400": {Message: "Bad Request - no Juniper Account Linked", Unmarshaller: errors.NewResponseDetailString},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	if claimed != nil {
		req.QueryParam("claimed", *claimed)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if serial != nil {
		req.QueryParam("serial", *serial)
	}
	if sku != nil {
		req.QueryParam("sku", *sku)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if warrantyType != nil {
		req.QueryParam("warranty_type", *warrantyType)
	}
	if eolAfter != nil {
		req.QueryParam("eol_after", *eolAfter)
	}
	if eolBefore != nil {
		req.QueryParam("eol_before", *eolBefore)
	}
	if eosAfter != nil {
		req.QueryParam("eos_after", *eosAfter)
	}
	if eosBefore != nil {
		req.QueryParam("eos_before", *eosBefore)
	}
	if versionEosAfter != nil {
		req.QueryParam("version_eos_after", *versionEosAfter)
	}
	if versionEosBefore != nil {
		req.QueryParam("version_eos_before", *versionEosBefore)
	}
	if hasSupport != nil {
		req.QueryParam("has_support", *hasSupport)
	}
	if sirtId != nil {
		req.QueryParam("sirt_id", *sirtId)
	}
	if pbnId != nil {
		req.QueryParam("pbn_id", *pbnId)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.JsInventorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.JsInventorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgJsiPbn takes context, orgId, distinct, limit, start, end as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Get count of PBN advisories grouped by specified field
func (o *OrgsJSI) CountOrgJsiPbn(
	ctx context.Context,
	orgId uuid.UUID,
	distinct models.CountPbnDistinctEnum,
	limit *int,
	start *string,
	end *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/pbn/count")
	req.AppendTemplateParams(orgId)
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
	req.QueryParam("distinct", distinct)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgJsiPbn takes context, orgId, versions, mModels, customerRisk, id, bugType, limit, page, searchAfter, start, end as parameters and
// returns an models.ApiResponse with models.JsiPbnSearch data and
// an error if there was an issue with the request or response.
// Text search for PBN (Problem Bug Notification) advisories. Search can be done on versions, models, customer_risk, id, and bug_type fields.
func (o *OrgsJSI) SearchOrgJsiPbn(
	ctx context.Context,
	orgId uuid.UUID,
	versions *string,
	mModels *string,
	customerRisk *string,
	id *string,
	bugType *string,
	limit *int,
	page *int,
	searchAfter *string,
	start *string,
	end *string) (
	models.ApiResponse[models.JsiPbnSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/pbn/search")
	req.AppendTemplateParams(orgId)
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
	if versions != nil {
		req.QueryParam("versions", *versions)
	}
	if mModels != nil {
		req.QueryParam("models", *mModels)
	}
	if customerRisk != nil {
		req.QueryParam("customer_risk", *customerRisk)
	}
	if id != nil {
		req.QueryParam("id", *id)
	}
	if bugType != nil {
		req.QueryParam("bug_type", *bugType)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.JsiPbnSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.JsiPbnSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgJsiSirt takes context, orgId, distinct, limit, start, end as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Get count of SIRT advisories grouped by specified field
func (o *OrgsJSI) CountOrgJsiSirt(
	ctx context.Context,
	orgId uuid.UUID,
	distinct models.CountSirtDistinctEnum,
	limit *int,
	start *string,
	end *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/sirt/count")
	req.AppendTemplateParams(orgId)
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
	req.QueryParam("distinct", distinct)
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgJsiSirt takes context, orgId, versions, mModels, severity, id, limit, page, searchAfter, start, end as parameters and
// returns an models.ApiResponse with models.JsiSirtSearch data and
// an error if there was an issue with the request or response.
// Text search for SIRT (Security Incident Response Team) advisories. Search can be done on versions, models, severity, and id fields.
func (o *OrgsJSI) SearchOrgJsiSirt(
	ctx context.Context,
	orgId uuid.UUID,
	versions *string,
	mModels *string,
	severity *string,
	id *string,
	limit *int,
	page *int,
	searchAfter *string,
	start *string,
	end *string) (
	models.ApiResponse[models.JsiSirtSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/sirt/search")
	req.AppendTemplateParams(orgId)
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
	if versions != nil {
		req.QueryParam("versions", *versions)
	}
	if mModels != nil {
		req.QueryParam("models", *mModels)
	}
	if severity != nil {
		req.QueryParam("severity", *severity)
	}
	if id != nil {
		req.QueryParam("id", *id)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.JsiSirtSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.JsiSirtSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
