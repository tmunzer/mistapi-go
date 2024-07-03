# \SitesClientsNACAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteNacClientEvents**](SitesClientsNACAPI.md#CountSiteNacClientEvents) | **Get** /api/v1/sites/{site_id}/nac_clients/events/count | countSiteNacClientEvents
[**CountSiteNacClients**](SitesClientsNACAPI.md#CountSiteNacClients) | **Get** /api/v1/sites/{site_id}/nac_clients/count | countSiteNacClients
[**SearchSiteNacClients**](SitesClientsNACAPI.md#SearchSiteNacClients) | **Get** /api/v1/sites/{site_id}/nac_clients/search | searchSiteNacClients
[**SearcsearcSiteNacClientEventsacClientEvents**](SitesClientsNACAPI.md#SearcsearcSiteNacClientEventsacClientEvents) | **Get** /api/v1/sites/{site_id}/nac_clients/events/search | searcsearcSiteNacClientEventsacClientEvents



## CountSiteNacClientEvents

> RepsonseCount CountSiteNacClientEvents(ctx, siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteNacClientEvents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.site_nac_client_events_count_distinct("") // SiteNacClientEventsCountDistinct |  (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsNACAPI.CountSiteNacClientEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsNACAPI.CountSiteNacClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteNacClientEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsNACAPI.CountSiteNacClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteNacClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteNacClientEventsCountDistinct**](SiteNacClientEventsCountDistinct.md) |  | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**RepsonseCount**](RepsonseCount.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountSiteNacClients

> RepsonseCount CountSiteNacClients(ctx, siteId).Distinct(distinct).LastNacruleId(lastNacruleId).NacruleMatched(nacruleMatched).AuthType(authType).LastVlanId(lastVlanId).LastNasVendor(lastNasVendor).IdpId(idpId).LastSsid(lastSsid).LastUsername(lastUsername).Timestamp(timestamp).LastAp(lastAp).Mac(mac).LastStatus(lastStatus).Type_(type_).MdmComplianceStatus(mdmComplianceStatus).MdmProvider(mdmProvider).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

countSiteNacClients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.site_nac_clients_count_distinct("") // SiteNacClientsCountDistinct | NAC Policy Rule ID, if matched (optional) (default to "type")
	lastNacruleId := "lastNacruleId_example" // string | NAC Policy Rule ID, if matched (optional)
	nacruleMatched := true // bool | NAC Policy Rule Matched (optional)
	authType := "authType_example" // string | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” (optional)
	lastVlanId := "lastVlanId_example" // string | Vlan ID (optional)
	lastNasVendor := "lastNasVendor_example" // string | vendor of NAS device (optional)
	idpId := "idpId_example" // string | SSO ID, if present and used (optional)
	lastSsid := "lastSsid_example" // string | SSID (optional)
	lastUsername := "lastUsername_example" // string | Username presented by the client (optional)
	timestamp := float32(8.14) // float32 | start time, in epoch (optional)
	lastAp := "lastAp_example" // string | AP MAC connected to by client (optional)
	mac := "mac_example" // string | MAC address (optional)
	lastStatus := "lastStatus_example" // string | Connection status of client i.e “permitted”, “denied, “session_ended” (optional)
	type_ := "type__example" // string | Client type i.e. “wireless”, “wired” etc. (optional)
	mdmComplianceStatus := "mdmComplianceStatus_example" // string | MDM compliancy of client i.e “compliant”, “not compliant” (optional)
	mdmProvider := "mdmProvider_example" // string | MDM provider of client’s organisation eg “intune”, “jamf” (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsNACAPI.CountSiteNacClients(context.Background(), siteId).Distinct(distinct).LastNacruleId(lastNacruleId).NacruleMatched(nacruleMatched).AuthType(authType).LastVlanId(lastVlanId).LastNasVendor(lastNasVendor).IdpId(idpId).LastSsid(lastSsid).LastUsername(lastUsername).Timestamp(timestamp).LastAp(lastAp).Mac(mac).LastStatus(lastStatus).Type_(type_).MdmComplianceStatus(mdmComplianceStatus).MdmProvider(mdmProvider).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsNACAPI.CountSiteNacClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteNacClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsNACAPI.CountSiteNacClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteNacClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteNacClientsCountDistinct**](SiteNacClientsCountDistinct.md) | NAC Policy Rule ID, if matched | [default to &quot;type&quot;]
 **lastNacruleId** | **string** | NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **bool** | NAC Policy Rule Matched | 
 **authType** | **string** | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” | 
 **lastVlanId** | **string** | Vlan ID | 
 **lastNasVendor** | **string** | vendor of NAS device | 
 **idpId** | **string** | SSO ID, if present and used | 
 **lastSsid** | **string** | SSID | 
 **lastUsername** | **string** | Username presented by the client | 
 **timestamp** | **float32** | start time, in epoch | 
 **lastAp** | **string** | AP MAC connected to by client | 
 **mac** | **string** | MAC address | 
 **lastStatus** | **string** | Connection status of client i.e “permitted”, “denied, “session_ended” | 
 **type_** | **string** | Client type i.e. “wireless”, “wired” etc. | 
 **mdmComplianceStatus** | **string** | MDM compliancy of client i.e “compliant”, “not compliant” | 
 **mdmProvider** | **string** | MDM provider of client’s organisation eg “intune”, “jamf” | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**RepsonseCount**](RepsonseCount.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteNacClients

> ResponseClientNacSearch SearchSiteNacClients(ctx, siteId).NacruleId(nacruleId).NacruleMatched(nacruleMatched).AuthType(authType).Vlan(vlan).NasVendor(nasVendor).IdpId(idpId).Ssid(ssid).Username(username).Timestamp(timestamp).Ap(ap).Mac(mac).MxedgeId(mxedgeId).NacruleName(nacruleName).Status(status).Type_(type_).MdmCompliance(mdmCompliance).MdmProvider(mdmProvider).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

searchSiteNacClients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nacruleId := "nacruleId_example" // string | NAC Policy Rule ID, if matched (optional)
	nacruleMatched := true // bool | NAC Policy Rule Matched (optional)
	authType := "authType_example" // string | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” (optional)
	vlan := "vlan_example" // string | Vlan name or ID assigned to the client (optional)
	nasVendor := "nasVendor_example" // string | vendor of NAS device (optional)
	idpId := "idpId_example" // string | SSO ID, if present and used (optional)
	ssid := "ssid_example" // string | SSID (optional)
	username := "username_example" // string | Username presented by the client (optional)
	timestamp := float32(8.14) // float32 | start time, in epoch (optional)
	ap := "ap_example" // string | AP MAC connected to by client (optional)
	mac := "mac_example" // string | MAC address (optional)
	mxedgeId := "mxedgeId_example" // string | ID of Mist Edge that the client is connected through (optional)
	nacruleName := "nacruleName_example" // string | NAC Policy Rule Name matched (optional)
	status := "status_example" // string | Connection status of client i.e “permitted”, “denied, “session_ended” (optional)
	type_ := "type__example" // string | Client type i.e. “wireless”, “wired” etc. (optional)
	mdmCompliance := "mdmCompliance_example" // string | MDM compliancy of client i.e “compliant”, “not compliant” (optional)
	mdmProvider := "mdmProvider_example" // string | MDM provider of client’s organisation eg “intune”, “jamf” (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsNACAPI.SearchSiteNacClients(context.Background(), siteId).NacruleId(nacruleId).NacruleMatched(nacruleMatched).AuthType(authType).Vlan(vlan).NasVendor(nasVendor).IdpId(idpId).Ssid(ssid).Username(username).Timestamp(timestamp).Ap(ap).Mac(mac).MxedgeId(mxedgeId).NacruleName(nacruleName).Status(status).Type_(type_).MdmCompliance(mdmCompliance).MdmProvider(mdmProvider).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsNACAPI.SearchSiteNacClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteNacClients`: ResponseClientNacSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsNACAPI.SearchSiteNacClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteNacClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nacruleId** | **string** | NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **bool** | NAC Policy Rule Matched | 
 **authType** | **string** | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” | 
 **vlan** | **string** | Vlan name or ID assigned to the client | 
 **nasVendor** | **string** | vendor of NAS device | 
 **idpId** | **string** | SSO ID, if present and used | 
 **ssid** | **string** | SSID | 
 **username** | **string** | Username presented by the client | 
 **timestamp** | **float32** | start time, in epoch | 
 **ap** | **string** | AP MAC connected to by client | 
 **mac** | **string** | MAC address | 
 **mxedgeId** | **string** | ID of Mist Edge that the client is connected through | 
 **nacruleName** | **string** | NAC Policy Rule Name matched | 
 **status** | **string** | Connection status of client i.e “permitted”, “denied, “session_ended” | 
 **type_** | **string** | Client type i.e. “wireless”, “wired” etc. | 
 **mdmCompliance** | **string** | MDM compliancy of client i.e “compliant”, “not compliant” | 
 **mdmProvider** | **string** | MDM provider of client’s organisation eg “intune”, “jamf” | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**ResponseClientNacSearch**](ResponseClientNacSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearcsearcSiteNacClientEventsacClientEvents

> ResponseEventsNacClientSearch SearcsearcSiteNacClientEventsacClientEvents(ctx, siteId).Type_(type_).NacruleId(nacruleId).NacruleMatched(nacruleMatched).DryrunNacruleId(dryrunNacruleId).DryrunNacruleMatched(dryrunNacruleMatched).AuthType(authType).Vlan(vlan).VlanSource(vlanSource).NasVendor(nasVendor).Bssid(bssid).IdpId(idpId).IdpRole(idpRole).IdpUsername(idpUsername).RespAttrs(respAttrs).Ssid(ssid).Username(username).UsermacLabels(usermacLabels).Ap(ap).RandomMac(randomMac).Mac(mac).LookupTimeTaken(lookupTimeTaken).Timestamp(timestamp).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searcsearcSiteNacClientEventsacClientEvents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) (optional)
	nacruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NAC Policy Rule ID, if matched (optional)
	nacruleMatched := true // bool | NAC Policy Rule Matched (optional)
	dryrunNacruleId := "dryrunNacruleId_example" // string | NAC Policy Dry Run Rule ID, if present and matched (optional)
	dryrunNacruleMatched := true // bool | True - if dryrun rule present and matched with priority, False - if not matched or not present (optional)
	authType := "authType_example" // string | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” (optional)
	vlan := int32(56) // int32 | Vlan ID (optional)
	vlanSource := "vlanSource_example" // string | Vlan source, e.g. \"nactag\", \"usermac\" (optional)
	nasVendor := "nasVendor_example" // string | vendor of NAS device (optional)
	bssid := "bssid_example" // string | SSID (optional)
	idpId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SSO ID, if present and used (optional)
	idpRole := "idpRole_example" // string | IDP returned roles/groups for the user (optional)
	idpUsername := "idpUsername_example" // string | Username presented to the Identity Provider (optional)
	respAttrs := []string{"Inner_example"} // []string | Radius attributes returned by NAC to NAS Devive (optional)
	ssid := "ssid_example" // string | SSID (optional)
	username := "username_example" // string | Username presented by the client (optional)
	usermacLabels := []string{"Inner_example"} // []string | labels derived from usermac entry (optional)
	ap := "ap_example" // string | AP MAC (optional)
	randomMac := true // bool | AP random macMAC (optional)
	mac := "mac_example" // string | MAC address (optional)
	lookupTimeTaken := float32(8.14) // float32 | Lookup(IDP etc.,) time taken in seconds (optional)
	timestamp := float32(8.14) // float32 | time, in epoch (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsNACAPI.SearcsearcSiteNacClientEventsacClientEvents(context.Background(), siteId).Type_(type_).NacruleId(nacruleId).NacruleMatched(nacruleMatched).DryrunNacruleId(dryrunNacruleId).DryrunNacruleMatched(dryrunNacruleMatched).AuthType(authType).Vlan(vlan).VlanSource(vlanSource).NasVendor(nasVendor).Bssid(bssid).IdpId(idpId).IdpRole(idpRole).IdpUsername(idpUsername).RespAttrs(respAttrs).Ssid(ssid).Username(username).UsermacLabels(usermacLabels).Ap(ap).RandomMac(randomMac).Mac(mac).LookupTimeTaken(lookupTimeTaken).Timestamp(timestamp).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsNACAPI.SearcsearcSiteNacClientEventsacClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearcsearcSiteNacClientEventsacClientEvents`: ResponseEventsNacClientSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsNACAPI.SearcsearcSiteNacClientEventsacClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearcsearcSiteNacClientEventsacClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listNacEventsDefinitions) | 
 **nacruleId** | **string** | NAC Policy Rule ID, if matched | 
 **nacruleMatched** | **bool** | NAC Policy Rule Matched | 
 **dryrunNacruleId** | **string** | NAC Policy Dry Run Rule ID, if present and matched | 
 **dryrunNacruleMatched** | **bool** | True - if dryrun rule present and matched with priority, False - if not matched or not present | 
 **authType** | **string** | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “device-auth” | 
 **vlan** | **int32** | Vlan ID | 
 **vlanSource** | **string** | Vlan source, e.g. \&quot;nactag\&quot;, \&quot;usermac\&quot; | 
 **nasVendor** | **string** | vendor of NAS device | 
 **bssid** | **string** | SSID | 
 **idpId** | **string** | SSO ID, if present and used | 
 **idpRole** | **string** | IDP returned roles/groups for the user | 
 **idpUsername** | **string** | Username presented to the Identity Provider | 
 **respAttrs** | **[]string** | Radius attributes returned by NAC to NAS Devive | 
 **ssid** | **string** | SSID | 
 **username** | **string** | Username presented by the client | 
 **usermacLabels** | **[]string** | labels derived from usermac entry | 
 **ap** | **string** | AP MAC | 
 **randomMac** | **bool** | AP random macMAC | 
 **mac** | **string** | MAC address | 
 **lookupTimeTaken** | **float32** | Lookup(IDP etc.,) time taken in seconds | 
 **timestamp** | **float32** | time, in epoch | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseEventsNacClientSearch**](ResponseEventsNacClientSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

