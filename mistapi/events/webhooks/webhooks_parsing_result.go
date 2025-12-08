// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package webhooks

import (
	"encoding/json"
	"errors"
	"fmt"
	"mistapi/models"
)

// WebhooksParsingResult represents a WebhooksParsingResult struct.
type WebhooksParsingResult struct {
	value                 any
	isAlarms              bool
	isAssetRawRssi        bool
	isAudits              bool
	isClientInfo          bool
	isClientJoin          bool
	isClientLatency       bool
	isClientSessions      bool
	isDeviceEvents        bool
	isDeviceUpdowns       bool
	isDiscoveredRawRssi   bool
	isGuestAuthorizations bool
	isLocation            bool
	isLocationAsset       bool
	isLocationCentrak     bool
	isLocationClient      bool
	isLocationSdk         bool
	isLocationUnclient    bool
	isMxedgeEvents        bool
	isNacAccounting       bool
	isNacEvents           bool
	isOccupancyAlerts     bool
	isPing                bool
	isSdkclientScanData   bool
	isSiteSle             bool
	isWifiConnRaw         bool
	isWifiUnconnRaw       bool
	isZone                bool
	isUnknown             bool
}

// String implements the fmt.Stringer interface for WebhooksParsingResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhooksParsingResult) String() string {
	return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WebhooksParsingResult.
// It customizes the JSON marshaling process for WebhooksParsingResult objects.
func (w WebhooksParsingResult) MarshalJSON() (
	[]byte,
	error) {
	if w.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.WebhooksContainer.From*` functions to initialize the WebhooksParsingResult object.")
	}
	return json.Marshal(w.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhooksParsingResult.
// It customizes the JSON unmarshaling process for WebhooksParsingResult objects.
func (w *WebhooksParsingResult) UnmarshalJSON(input []byte) error {
	result, err := models.UnmarshallOneOf(input,
		models.NewTypeHolder(&models.WebhookAlarms{}, true, &w.isAlarms),
		models.NewTypeHolder(&models.WebhookAssetRawRssi{}, true, &w.isAssetRawRssi),
		models.NewTypeHolder(&models.WebhookAudits{}, true, &w.isAudits),
		models.NewTypeHolder(&models.WebhookClientInfo{}, true, &w.isClientInfo),
		models.NewTypeHolder(&models.WebhookClientJoin{}, true, &w.isClientJoin),
		models.NewTypeHolder(&models.WebhookClientLatency{}, true, &w.isClientLatency),
		models.NewTypeHolder(&models.WebhookClientSessions{}, true, &w.isClientSessions),
		models.NewTypeHolder(&models.WebhookDeviceEvents{}, true, &w.isDeviceEvents),
		models.NewTypeHolder(&models.WebhookDeviceUpdowns{}, true, &w.isDeviceUpdowns),
		models.NewTypeHolder(&models.WebhookDiscoveredRawRssi{}, true, &w.isDiscoveredRawRssi),
		models.NewTypeHolder(&models.WebhookGuestAuthorizations{}, true, &w.isGuestAuthorizations),
		models.NewTypeHolder(&models.WebhookLocation{}, true, &w.isLocation),
		models.NewTypeHolder(&models.WebhookLocationAsset{}, true, &w.isLocationAsset),
		models.NewTypeHolder(new(interface{}), true, &w.isLocationCentrak),
		models.NewTypeHolder(&models.WebhookLocationClient{}, true, &w.isLocationClient),
		models.NewTypeHolder(&models.WebhookLocationSdk{}, true, &w.isLocationSdk),
		models.NewTypeHolder(&models.WebhookLocationUnclient{}, true, &w.isLocationUnclient),
		models.NewTypeHolder(&models.WebhookMxedgeEvents{}, true, &w.isMxedgeEvents),
		models.NewTypeHolder(&models.WebhookNacAccounting{}, true, &w.isNacAccounting),
		models.NewTypeHolder(&models.WebhookNacEvents{}, true, &w.isNacEvents),
		models.NewTypeHolder(&models.WebhookOccupancyAlerts{}, true, &w.isOccupancyAlerts),
		models.NewTypeHolder(&models.WebhookPing{}, true, &w.isPing),
		models.NewTypeHolder(&models.WebhookSdkclientScanData{}, true, &w.isSdkclientScanData),
		models.NewTypeHolder(&models.WebhookSiteSle{}, true, &w.isSiteSle),
		models.NewTypeHolder(&models.WebhookWifiConnRaw{}, true, &w.isWifiConnRaw),
		models.NewTypeHolder(&models.WebhookWifiUnconnRaw{}, true, &w.isWifiUnconnRaw),
		models.NewTypeHolder(&models.WebhookZone{}, true, &w.isZone),
	)

	w.value = result
	return err
}

func (w *WebhooksParsingResult) AsAlarms() (
	*models.WebhookAlarms,
	bool) {
	if !w.isAlarms {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookAlarms), true
}

func (w *WebhooksParsingResult) AsAssetRawRssi() (
	*models.WebhookAssetRawRssi,
	bool) {
	if !w.isAssetRawRssi {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookAssetRawRssi), true
}

func (w *WebhooksParsingResult) AsAudits() (
	*models.WebhookAudits,
	bool) {
	if !w.isAudits {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookAudits), true
}

func (w *WebhooksParsingResult) AsClientInfo() (
	*models.WebhookClientInfo,
	bool) {
	if !w.isClientInfo {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookClientInfo), true
}

func (w *WebhooksParsingResult) AsClientJoin() (
	*models.WebhookClientJoin,
	bool) {
	if !w.isClientJoin {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookClientJoin), true
}

func (w *WebhooksParsingResult) AsClientLatency() (
	*models.WebhookClientLatency,
	bool) {
	if !w.isClientLatency {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookClientLatency), true
}

func (w *WebhooksParsingResult) AsClientSessions() (
	*models.WebhookClientSessions,
	bool) {
	if !w.isClientSessions {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookClientSessions), true
}

func (w *WebhooksParsingResult) AsDeviceEvents() (
	*models.WebhookDeviceEvents,
	bool) {
	if !w.isDeviceEvents {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookDeviceEvents), true
}

func (w *WebhooksParsingResult) AsDeviceUpdowns() (
	*models.WebhookDeviceUpdowns,
	bool) {
	if !w.isDeviceUpdowns {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookDeviceUpdowns), true
}

func (w *WebhooksParsingResult) AsDiscoveredRawRssi() (
	*models.WebhookDiscoveredRawRssi,
	bool) {
	if !w.isDiscoveredRawRssi {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookDiscoveredRawRssi), true
}

func (w *WebhooksParsingResult) AsGuestAuthorizations() (
	*models.WebhookGuestAuthorizations,
	bool) {
	if !w.isGuestAuthorizations {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookGuestAuthorizations), true
}

func (w *WebhooksParsingResult) AsLocation() (
	*models.WebhookLocation,
	bool) {
	if !w.isLocation {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookLocation), true
}

func (w *WebhooksParsingResult) AsLocationAsset() (
	*models.WebhookLocationAsset,
	bool) {
	if !w.isLocationAsset {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookLocationAsset), true
}

func (w *WebhooksParsingResult) AsLocationCentrak() (
	*interface{},
	bool) {
	if !w.isLocationCentrak {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*interface{}), true
}

func (w *WebhooksParsingResult) AsLocationClient() (
	*models.WebhookLocationClient,
	bool) {
	if !w.isLocationClient {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookLocationClient), true
}

func (w *WebhooksParsingResult) AsLocationSdk() (
	*models.WebhookLocationSdk,
	bool) {
	if !w.isLocationSdk {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookLocationSdk), true
}

func (w *WebhooksParsingResult) AsLocationUnclient() (
	*models.WebhookLocationUnclient,
	bool) {
	if !w.isLocationUnclient {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookLocationUnclient), true
}

func (w *WebhooksParsingResult) AsMxedgeEvents() (
	*models.WebhookMxedgeEvents,
	bool) {
	if !w.isMxedgeEvents {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookMxedgeEvents), true
}

func (w *WebhooksParsingResult) AsNacAccounting() (
	*models.WebhookNacAccounting,
	bool) {
	if !w.isNacAccounting {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookNacAccounting), true
}

func (w *WebhooksParsingResult) AsNacEvents() (
	*models.WebhookNacEvents,
	bool) {
	if !w.isNacEvents {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookNacEvents), true
}

func (w *WebhooksParsingResult) AsOccupancyAlerts() (
	*models.WebhookOccupancyAlerts,
	bool) {
	if !w.isOccupancyAlerts {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookOccupancyAlerts), true
}

func (w *WebhooksParsingResult) AsPing() (
	*models.WebhookPing,
	bool) {
	if !w.isPing {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookPing), true
}

func (w *WebhooksParsingResult) AsSdkclientScanData() (
	*models.WebhookSdkclientScanData,
	bool) {
	if !w.isSdkclientScanData {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookSdkclientScanData), true
}

func (w *WebhooksParsingResult) AsSiteSle() (
	*models.WebhookSiteSle,
	bool) {
	if !w.isSiteSle {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookSiteSle), true
}

func (w *WebhooksParsingResult) AsWifiConnRaw() (
	*models.WebhookWifiConnRaw,
	bool) {
	if !w.isWifiConnRaw {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookWifiConnRaw), true
}

func (w *WebhooksParsingResult) AsWifiUnconnRaw() (
	*models.WebhookWifiUnconnRaw,
	bool) {
	if !w.isWifiUnconnRaw {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookWifiUnconnRaw), true
}

func (w *WebhooksParsingResult) AsZone() (
	*models.WebhookZone,
	bool) {
	if !w.isZone {
		return nil, false
	}
	if w.value == nil {
		return nil, true
	}
	return w.value.(*models.WebhookZone), true
}

func (w *WebhooksParsingResult) AsUnknown() bool {
	return w.isUnknown
}

// internalWebhooksParsingResult represents a webhooksParsingResult struct.
type internalWebhooksParsingResult struct{}

var WebhooksParsingResultContainer internalWebhooksParsingResult

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookAlarms value.
func (w *internalWebhooksParsingResult) FromAlarms(val *models.WebhookAlarms) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookAssetRawRssi value.
func (w *internalWebhooksParsingResult) FromAssetRawRssi(val *models.WebhookAssetRawRssi) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookAudits value.
func (w *internalWebhooksParsingResult) FromAudits(val *models.WebhookAudits) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookClientInfo value.
func (w *internalWebhooksParsingResult) FromClientInfo(val *models.WebhookClientInfo) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookClientJoin value.
func (w *internalWebhooksParsingResult) FromClientJoin(val *models.WebhookClientJoin) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookClientLatency value.
func (w *internalWebhooksParsingResult) FromClientLatency(val *models.WebhookClientLatency) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookClientSessions value.
func (w *internalWebhooksParsingResult) FromClientSessions(val *models.WebhookClientSessions) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookDeviceEvents value.
func (w *internalWebhooksParsingResult) FromDeviceEvents(val *models.WebhookDeviceEvents) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookDeviceUpdowns value.
func (w *internalWebhooksParsingResult) FromDeviceUpdowns(val *models.WebhookDeviceUpdowns) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookDiscoveredRawRssi value.
func (w *internalWebhooksParsingResult) FromDiscoveredRawRssi(val *models.WebhookDiscoveredRawRssi) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookGuestAuthorizations value.
func (w *internalWebhooksParsingResult) FromGuestAuthorizations(val *models.WebhookGuestAuthorizations) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookLocation value.
func (w *internalWebhooksParsingResult) FromLocation(val *models.WebhookLocation) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookLocationAsset value.
func (w *internalWebhooksParsingResult) FromLocationAsset(val *models.WebhookLocationAsset) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided interface{} value.
func (w *internalWebhooksParsingResult) FromLocationCentrak(val *interface{}) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookLocationClient value.
func (w *internalWebhooksParsingResult) FromLocationClient(val *models.WebhookLocationClient) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookLocationSdk value.
func (w *internalWebhooksParsingResult) FromLocationSdk(val *models.WebhookLocationSdk) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookLocationUnclient value.
func (w *internalWebhooksParsingResult) FromLocationUnclient(val *models.WebhookLocationUnclient) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookMxedgeEvents value.
func (w *internalWebhooksParsingResult) FromMxedgeEvents(val *models.WebhookMxedgeEvents) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookNacAccounting value.
func (w *internalWebhooksParsingResult) FromNacAccounting(val *models.WebhookNacAccounting) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookNacEvents value.
func (w *internalWebhooksParsingResult) FromNacEvents(val *models.WebhookNacEvents) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookOccupancyAlerts value.
func (w *internalWebhooksParsingResult) FromOccupancyAlerts(val *models.WebhookOccupancyAlerts) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookPing value.
func (w *internalWebhooksParsingResult) FromPing(val *models.WebhookPing) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookSdkclientScanData value.
func (w *internalWebhooksParsingResult) FromSdkclientScanData(val *models.WebhookSdkclientScanData) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookSiteSle value.
func (w *internalWebhooksParsingResult) FromSiteSle(val *models.WebhookSiteSle) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookWifiConnRaw value.
func (w *internalWebhooksParsingResult) FromWifiConnRaw(val *models.WebhookWifiConnRaw) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookWifiUnconnRaw value.
func (w *internalWebhooksParsingResult) FromWifiUnconnRaw(val *models.WebhookWifiUnconnRaw) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, wrapping the provided models.WebhookZone value.
func (w *internalWebhooksParsingResult) FromZone(val *models.WebhookZone) WebhooksParsingResult {
	return WebhooksParsingResult{value: &val}
}

// The internalWebhooksParsingResult instance, Triggered when no event identified.
func (w *internalWebhooksParsingResult) MarkUnknown() WebhooksParsingResult {
	return WebhooksParsingResult{
		value:     nil,
		isUnknown: true,
	}
}
