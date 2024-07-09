package mistapi

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// SitesDevicesStats represents a controller struct.
type SitesDevicesStats struct {
	baseController
}

// NewSitesDevicesStats creates a new instance of SitesDevicesStats.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesStats.
func NewSitesDevicesStats(baseController baseController) *SitesDevicesStats {
	sitesDevicesStats := SitesDevicesStats{baseController: baseController}
	return &sitesDevicesStats
}

// CountSiteBgpStats takes context, siteId, state, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count BGP Stats
func (s *SitesDevicesStats) CountSiteBgpStats(
	ctx context.Context,
	siteId uuid.UUID,
	state *string,
	distinct *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/bgp_peers/count", siteId),
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
	if state != nil {
		req.QueryParam("state", *state)
	}
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteBgpStats takes context, siteId as parameters and
// returns an models.ApiResponse with models.ResponseSearchBgps data and
// an error if there was an issue with the request or response.
// Search BGP Stats
func (s *SitesDevicesStats) SearchSiteBgpStats(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.ResponseSearchBgps],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/bgp_peers/search", siteId),
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

	var result models.ResponseSearchBgps
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSearchBgps](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteDevicesStats takes context, siteId, mType, status, page, limit as parameters and
// returns an models.ApiResponse with []models.ListSiteDevicesStatsResponse data and
// an error if there was an issue with the request or response.
// Get List of Site Devices Stats
func (s *SitesDevicesStats) ListSiteDevicesStats(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeWithAllEnum,
	status *models.StatDeviceStatusFilterEnum,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.ListSiteDevicesStatsResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/devices", siteId),
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
	if status != nil {
		req.QueryParam("status", *status)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.ListSiteDevicesStatsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ListSiteDevicesStatsResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceStats takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.StatsDevice2 data and
// an error if there was an issue with the request or response.
// Get Site Device Stats Details
func (s *SitesDevicesStats) GetSiteDeviceStats(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	models.ApiResponse[models.StatsDevice2],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/devices/%v", siteId, deviceId),
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

	var result models.StatsDevice2
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsDevice2](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteAllClientsStatsByDevice takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with [][]models.ClientWirelessStats data and
// an error if there was an issue with the request or response.
// Get wireless client stat by Device
func (s *SitesDevicesStats) GetSiteAllClientsStatsByDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	models.ApiResponse[[][]models.ClientWirelessStats],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/devices/%v/clients", siteId, deviceId),
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

	var result [][]models.ClientWirelessStats
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[][]models.ClientWirelessStats](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteGatewayMetrics takes context, siteId as parameters and
// returns an models.ApiResponse with models.GatewayMetrics data and
// an error if there was an issue with the request or response.
// Get Site Gateway Metrics
func (s *SitesDevicesStats) GetSiteGatewayMetrics(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.GatewayMetrics],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/gateways/metrics", siteId),
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

	var result models.GatewayMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GatewayMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteMxEdgesStats takes context, siteId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.MxedgeStats data and
// an error if there was an issue with the request or response.
// Get List of Site MxEdges Stats
func (s *SitesDevicesStats) ListSiteMxEdgesStats(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.MxedgeStats],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/mxedges", siteId),
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

	var result []models.MxedgeStats
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.MxedgeStats](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteMxEdgeStats takes context, siteId, mxedgeId, start, end, duration as parameters and
// returns an models.ApiResponse with models.MxedgeStats data and
// an error if there was an issue with the request or response.
// Get One Site MxEdge Stats
func (s *SitesDevicesStats) GetSiteMxEdgeStats(
	ctx context.Context,
	siteId uuid.UUID,
	mxedgeId uuid.UUID,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.MxedgeStats],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/mxedges/%v", siteId, mxedgeId),
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.MxedgeStats
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MxedgeStats](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteSwOrGwPorts takes context, siteId, distinct, fullDuplex, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, stpState, stpRole, authState, up, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Switch/Gateway Ports
func (s *SitesDevicesStats) CountSiteSwOrGwPorts(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SitePortsCountDistinctEnum,
	fullDuplex *bool,
	mac *string,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	portId *string,
	portMac *string,
	powerDraw *float64,
	txPkts *int,
	rxPkts *int,
	rxBytes *int,
	txBps *int,
	rxBps *int,
	txMcastPkts *int,
	txBcastPkts *int,
	rxMcastPkts *int,
	rxBcastPkts *int,
	speed *int,
	stpState *models.CountSiteSwOrGwPortsStpStateEnum,
	stpRole *models.CountSiteSwOrGwPortsStpRoleEnum,
	authState *models.CountSiteSwOrGwPortsAuthStateEnum,
	up *bool,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/ports/count", siteId),
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if powerDraw != nil {
		req.QueryParam("power_draw", *powerDraw)
	}
	if txPkts != nil {
		req.QueryParam("tx_pkts", *txPkts)
	}
	if rxPkts != nil {
		req.QueryParam("rx_pkts", *rxPkts)
	}
	if rxBytes != nil {
		req.QueryParam("rx_bytes", *rxBytes)
	}
	if txBps != nil {
		req.QueryParam("tx_bps", *txBps)
	}
	if rxBps != nil {
		req.QueryParam("rx_bps", *rxBps)
	}
	if txMcastPkts != nil {
		req.QueryParam("tx_mcast_pkts", *txMcastPkts)
	}
	if txBcastPkts != nil {
		req.QueryParam("tx_bcast_pkts", *txBcastPkts)
	}
	if rxMcastPkts != nil {
		req.QueryParam("rx_mcast_pkts", *rxMcastPkts)
	}
	if rxBcastPkts != nil {
		req.QueryParam("rx_bcast_pkts", *rxBcastPkts)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if up != nil {
		req.QueryParam("up", *up)
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

// SearchSiteSwOrGwPorts takes context, siteId, fullDuplex, mac, deviceType, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txErrors, rxErrors, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, macLimit, macCount, up, active, jitter, loss, latency, stpState, stpRole, xcvrPartNumber, authState, lteImsi, lteIccid, lteImei, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseSwitchPortSearch data and
// an error if there was an issue with the request or response.
// Search Switch / Gateway Ports
func (s *SitesDevicesStats) SearchSiteSwOrGwPorts(
	ctx context.Context,
	siteId uuid.UUID,
	fullDuplex *bool,
	mac *string,
	deviceType *models.SearchSiteSwOrGwPortsDeviceTypeEnum,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	portId *string,
	portMac *string,
	powerDraw *float64,
	txPkts *int,
	rxPkts *int,
	rxBytes *int,
	txBps *int,
	rxBps *int,
	txErrors *int,
	rxErrors *int,
	txMcastPkts *int,
	txBcastPkts *int,
	rxMcastPkts *int,
	rxBcastPkts *int,
	speed *int,
	macLimit *int,
	macCount *int,
	up *bool,
	active *bool,
	jitter *float64,
	loss *float64,
	latency *float64,
	stpState *models.SearchSiteSwOrGwPortsStpStateEnum,
	stpRole *models.SearchSiteSwOrGwPortsStpRoleEnum,
	xcvrPartNumber *string,
	authState *models.SearchSiteSwOrGwPortsAuthStateEnum,
	lteImsi *string,
	lteIccid *string,
	lteImei *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseSwitchPortSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/ports/search", siteId),
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
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if deviceType != nil {
		req.QueryParam("device_type", *deviceType)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if powerDraw != nil {
		req.QueryParam("power_draw", *powerDraw)
	}
	if txPkts != nil {
		req.QueryParam("tx_pkts", *txPkts)
	}
	if rxPkts != nil {
		req.QueryParam("rx_pkts", *rxPkts)
	}
	if rxBytes != nil {
		req.QueryParam("rx_bytes", *rxBytes)
	}
	if txBps != nil {
		req.QueryParam("tx_bps", *txBps)
	}
	if rxBps != nil {
		req.QueryParam("rx_bps", *rxBps)
	}
	if txErrors != nil {
		req.QueryParam("tx_errors", *txErrors)
	}
	if rxErrors != nil {
		req.QueryParam("rx_errors", *rxErrors)
	}
	if txMcastPkts != nil {
		req.QueryParam("tx_mcast_pkts", *txMcastPkts)
	}
	if txBcastPkts != nil {
		req.QueryParam("tx_bcast_pkts", *txBcastPkts)
	}
	if rxMcastPkts != nil {
		req.QueryParam("rx_mcast_pkts", *rxMcastPkts)
	}
	if rxBcastPkts != nil {
		req.QueryParam("rx_bcast_pkts", *rxBcastPkts)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if macLimit != nil {
		req.QueryParam("mac_limit", *macLimit)
	}
	if macCount != nil {
		req.QueryParam("mac_count", *macCount)
	}
	if up != nil {
		req.QueryParam("up", *up)
	}
	if active != nil {
		req.QueryParam("active", *active)
	}
	if jitter != nil {
		req.QueryParam("jitter", *jitter)
	}
	if loss != nil {
		req.QueryParam("loss", *loss)
	}
	if latency != nil {
		req.QueryParam("latency", *latency)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if xcvrPartNumber != nil {
		req.QueryParam("xcvr_part_number", *xcvrPartNumber)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if lteImsi != nil {
		req.QueryParam("lte_imsi", *lteImsi)
	}
	if lteIccid != nil {
		req.QueryParam("lte_iccid", *lteIccid)
	}
	if lteImei != nil {
		req.QueryParam("lte_imei", *lteImei)
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

	var result models.ResponseSwitchPortSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSwitchPortSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteSwitchPorts takes context, siteId, distinct, fullDuplex, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, stpState, stpRole, authState, up, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Switch/Gateway Ports
func (s *SitesDevicesStats) CountSiteSwitchPorts(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteSwitchPortsCountDistinctEnum,
	fullDuplex *bool,
	mac *string,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	portId *string,
	portMac *string,
	powerDraw *float64,
	txPkts *int,
	rxPkts *int,
	rxBytes *int,
	txBps *int,
	rxBps *int,
	txMcastPkts *int,
	txBcastPkts *int,
	rxMcastPkts *int,
	rxBcastPkts *int,
	speed *int,
	stpState *models.CountSiteSwitchPortsStpStateEnum,
	stpRole *models.CountSiteSwitchPortsStpRoleEnum,
	authState *models.CountSiteSwitchPortsAuthStateEnum,
	up *bool,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/switch_ports/count", siteId),
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if powerDraw != nil {
		req.QueryParam("power_draw", *powerDraw)
	}
	if txPkts != nil {
		req.QueryParam("tx_pkts", *txPkts)
	}
	if rxPkts != nil {
		req.QueryParam("rx_pkts", *rxPkts)
	}
	if rxBytes != nil {
		req.QueryParam("rx_bytes", *rxBytes)
	}
	if txBps != nil {
		req.QueryParam("tx_bps", *txBps)
	}
	if rxBps != nil {
		req.QueryParam("rx_bps", *rxBps)
	}
	if txMcastPkts != nil {
		req.QueryParam("tx_mcast_pkts", *txMcastPkts)
	}
	if txBcastPkts != nil {
		req.QueryParam("tx_bcast_pkts", *txBcastPkts)
	}
	if rxMcastPkts != nil {
		req.QueryParam("rx_mcast_pkts", *rxMcastPkts)
	}
	if rxBcastPkts != nil {
		req.QueryParam("rx_bcast_pkts", *rxBcastPkts)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if up != nil {
		req.QueryParam("up", *up)
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

// SearchSiteSwitchPorts takes context, siteId, fullDuplex, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, stpState, stpRole, authState, up, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseSwitchPortSearch data and
// an error if there was an issue with the request or response.
// Search Switch / Gateway Ports
func (s *SitesDevicesStats) SearchSiteSwitchPorts(
	ctx context.Context,
	siteId uuid.UUID,
	fullDuplex *bool,
	mac *string,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	portId *string,
	portMac *string,
	powerDraw *float64,
	txPkts *int,
	rxPkts *int,
	rxBytes *int,
	txBps *int,
	rxBps *int,
	txMcastPkts *int,
	txBcastPkts *int,
	rxMcastPkts *int,
	rxBcastPkts *int,
	speed *int,
	stpState *models.SearchSiteSwitchPortsStpStateEnum,
	stpRole *models.SearchSiteSwitchPortsStpRoleEnum,
	authState *models.SearchSiteSwitchPortsAuthStateEnum,
	up *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseSwitchPortSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/switch_ports/search", siteId),
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
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if powerDraw != nil {
		req.QueryParam("power_draw", *powerDraw)
	}
	if txPkts != nil {
		req.QueryParam("tx_pkts", *txPkts)
	}
	if rxPkts != nil {
		req.QueryParam("rx_pkts", *rxPkts)
	}
	if rxBytes != nil {
		req.QueryParam("rx_bytes", *rxBytes)
	}
	if txBps != nil {
		req.QueryParam("tx_bps", *txBps)
	}
	if rxBps != nil {
		req.QueryParam("rx_bps", *rxBps)
	}
	if txMcastPkts != nil {
		req.QueryParam("tx_mcast_pkts", *txMcastPkts)
	}
	if txBcastPkts != nil {
		req.QueryParam("tx_bcast_pkts", *txBcastPkts)
	}
	if rxMcastPkts != nil {
		req.QueryParam("rx_mcast_pkts", *rxMcastPkts)
	}
	if rxBcastPkts != nil {
		req.QueryParam("rx_bcast_pkts", *rxBcastPkts)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if up != nil {
		req.QueryParam("up", *up)
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

	var result models.ResponseSwitchPortSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSwitchPortSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
