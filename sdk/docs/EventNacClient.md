# EventNacClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** | authentication type, e.g. “eap-tls”, “eap-ttls”, “eap-teap”, “mab”, “psk”, “device-auth” | [optional] 
**Bssid** | Pointer to **string** |  | [optional] 
**DeviceMac** | Pointer to **string** |  | [optional] 
**DryrunNacruleId** | Pointer to **string** | NAC Policy Dry Run Rule ID, if present and matched | [optional] 
**DryrunNacruleMatched** | Pointer to **bool** |  | [optional] 
**IdpId** | Pointer to **string** |  | [optional] 
**IdpRole** | Pointer to **[]string** |  | [optional] 
**IdpUsername** | Pointer to **string** | Username presented to the Identity Provider | [optional] 
**Mac** | Pointer to **string** | Client MAC address | [optional] 
**MxedgeId** | Pointer to **string** | Mist Edge ID used to connect to cloud | [optional] 
**NacruleId** | Pointer to **string** | NAC Policy Rule ID, if matched | [optional] 
**NacruleMatched** | Pointer to **bool** |  | [optional] 
**NasVendor** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortId** | Pointer to **string** |  | [optional] 
**PortType** | Pointer to **string** | Type of client i.e wired, wireless | [optional] 
**RandomMac** | Pointer to **bool** |  | [optional] 
**RespAttrs** | Pointer to **[]string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssid** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** | event type, e.g. NAC_CLIENT_PERMIT | [optional] 
**UsermacLabels** | Pointer to **[]string** | labels derived from usermac entry | [optional] 
**Username** | Pointer to **string** | Username presented by the client | [optional] 
**Vlan** | Pointer to **string** | Vlan ID | [optional] 
**VlanSource** | Pointer to **string** | Vlan source, e.g. \&quot;nactag\&quot;, \&quot;usermac\&quot; | [optional] 

## Methods

### NewEventNacClient

`func NewEventNacClient() *EventNacClient`

NewEventNacClient instantiates a new EventNacClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNacClientWithDefaults

`func NewEventNacClientWithDefaults() *EventNacClient`

NewEventNacClientWithDefaults instantiates a new EventNacClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *EventNacClient) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *EventNacClient) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *EventNacClient) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *EventNacClient) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetAuthType

`func (o *EventNacClient) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EventNacClient) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EventNacClient) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *EventNacClient) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBssid

`func (o *EventNacClient) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *EventNacClient) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *EventNacClient) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *EventNacClient) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetDeviceMac

`func (o *EventNacClient) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *EventNacClient) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *EventNacClient) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *EventNacClient) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDryrunNacruleId

`func (o *EventNacClient) GetDryrunNacruleId() string`

GetDryrunNacruleId returns the DryrunNacruleId field if non-nil, zero value otherwise.

### GetDryrunNacruleIdOk

`func (o *EventNacClient) GetDryrunNacruleIdOk() (*string, bool)`

GetDryrunNacruleIdOk returns a tuple with the DryrunNacruleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunNacruleId

`func (o *EventNacClient) SetDryrunNacruleId(v string)`

SetDryrunNacruleId sets DryrunNacruleId field to given value.

### HasDryrunNacruleId

`func (o *EventNacClient) HasDryrunNacruleId() bool`

HasDryrunNacruleId returns a boolean if a field has been set.

### GetDryrunNacruleMatched

`func (o *EventNacClient) GetDryrunNacruleMatched() bool`

GetDryrunNacruleMatched returns the DryrunNacruleMatched field if non-nil, zero value otherwise.

### GetDryrunNacruleMatchedOk

`func (o *EventNacClient) GetDryrunNacruleMatchedOk() (*bool, bool)`

GetDryrunNacruleMatchedOk returns a tuple with the DryrunNacruleMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunNacruleMatched

`func (o *EventNacClient) SetDryrunNacruleMatched(v bool)`

SetDryrunNacruleMatched sets DryrunNacruleMatched field to given value.

### HasDryrunNacruleMatched

`func (o *EventNacClient) HasDryrunNacruleMatched() bool`

HasDryrunNacruleMatched returns a boolean if a field has been set.

### GetIdpId

`func (o *EventNacClient) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *EventNacClient) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *EventNacClient) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *EventNacClient) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetIdpRole

`func (o *EventNacClient) GetIdpRole() []string`

GetIdpRole returns the IdpRole field if non-nil, zero value otherwise.

### GetIdpRoleOk

`func (o *EventNacClient) GetIdpRoleOk() (*[]string, bool)`

GetIdpRoleOk returns a tuple with the IdpRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpRole

`func (o *EventNacClient) SetIdpRole(v []string)`

SetIdpRole sets IdpRole field to given value.

### HasIdpRole

`func (o *EventNacClient) HasIdpRole() bool`

HasIdpRole returns a boolean if a field has been set.

### GetIdpUsername

`func (o *EventNacClient) GetIdpUsername() string`

GetIdpUsername returns the IdpUsername field if non-nil, zero value otherwise.

### GetIdpUsernameOk

`func (o *EventNacClient) GetIdpUsernameOk() (*string, bool)`

GetIdpUsernameOk returns a tuple with the IdpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUsername

`func (o *EventNacClient) SetIdpUsername(v string)`

SetIdpUsername sets IdpUsername field to given value.

### HasIdpUsername

`func (o *EventNacClient) HasIdpUsername() bool`

HasIdpUsername returns a boolean if a field has been set.

### GetMac

`func (o *EventNacClient) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EventNacClient) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EventNacClient) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *EventNacClient) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMxedgeId

`func (o *EventNacClient) GetMxedgeId() string`

GetMxedgeId returns the MxedgeId field if non-nil, zero value otherwise.

### GetMxedgeIdOk

`func (o *EventNacClient) GetMxedgeIdOk() (*string, bool)`

GetMxedgeIdOk returns a tuple with the MxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeId

`func (o *EventNacClient) SetMxedgeId(v string)`

SetMxedgeId sets MxedgeId field to given value.

### HasMxedgeId

`func (o *EventNacClient) HasMxedgeId() bool`

HasMxedgeId returns a boolean if a field has been set.

### GetNacruleId

`func (o *EventNacClient) GetNacruleId() string`

GetNacruleId returns the NacruleId field if non-nil, zero value otherwise.

### GetNacruleIdOk

`func (o *EventNacClient) GetNacruleIdOk() (*string, bool)`

GetNacruleIdOk returns a tuple with the NacruleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacruleId

`func (o *EventNacClient) SetNacruleId(v string)`

SetNacruleId sets NacruleId field to given value.

### HasNacruleId

`func (o *EventNacClient) HasNacruleId() bool`

HasNacruleId returns a boolean if a field has been set.

### GetNacruleMatched

`func (o *EventNacClient) GetNacruleMatched() bool`

GetNacruleMatched returns the NacruleMatched field if non-nil, zero value otherwise.

### GetNacruleMatchedOk

`func (o *EventNacClient) GetNacruleMatchedOk() (*bool, bool)`

GetNacruleMatchedOk returns a tuple with the NacruleMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNacruleMatched

`func (o *EventNacClient) SetNacruleMatched(v bool)`

SetNacruleMatched sets NacruleMatched field to given value.

### HasNacruleMatched

`func (o *EventNacClient) HasNacruleMatched() bool`

HasNacruleMatched returns a boolean if a field has been set.

### GetNasVendor

`func (o *EventNacClient) GetNasVendor() string`

GetNasVendor returns the NasVendor field if non-nil, zero value otherwise.

### GetNasVendorOk

`func (o *EventNacClient) GetNasVendorOk() (*string, bool)`

GetNasVendorOk returns a tuple with the NasVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasVendor

`func (o *EventNacClient) SetNasVendor(v string)`

SetNasVendor sets NasVendor field to given value.

### HasNasVendor

`func (o *EventNacClient) HasNasVendor() bool`

HasNasVendor returns a boolean if a field has been set.

### GetOrgId

`func (o *EventNacClient) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventNacClient) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventNacClient) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventNacClient) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortId

`func (o *EventNacClient) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EventNacClient) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EventNacClient) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EventNacClient) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortType

`func (o *EventNacClient) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EventNacClient) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EventNacClient) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EventNacClient) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetRandomMac

`func (o *EventNacClient) GetRandomMac() bool`

GetRandomMac returns the RandomMac field if non-nil, zero value otherwise.

### GetRandomMacOk

`func (o *EventNacClient) GetRandomMacOk() (*bool, bool)`

GetRandomMacOk returns a tuple with the RandomMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomMac

`func (o *EventNacClient) SetRandomMac(v bool)`

SetRandomMac sets RandomMac field to given value.

### HasRandomMac

`func (o *EventNacClient) HasRandomMac() bool`

HasRandomMac returns a boolean if a field has been set.

### GetRespAttrs

`func (o *EventNacClient) GetRespAttrs() []string`

GetRespAttrs returns the RespAttrs field if non-nil, zero value otherwise.

### GetRespAttrsOk

`func (o *EventNacClient) GetRespAttrsOk() (*[]string, bool)`

GetRespAttrsOk returns a tuple with the RespAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespAttrs

`func (o *EventNacClient) SetRespAttrs(v []string)`

SetRespAttrs sets RespAttrs field to given value.

### HasRespAttrs

`func (o *EventNacClient) HasRespAttrs() bool`

HasRespAttrs returns a boolean if a field has been set.

### GetSiteId

`func (o *EventNacClient) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EventNacClient) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EventNacClient) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *EventNacClient) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsid

`func (o *EventNacClient) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *EventNacClient) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *EventNacClient) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *EventNacClient) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventNacClient) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventNacClient) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventNacClient) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventNacClient) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *EventNacClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventNacClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventNacClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventNacClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsermacLabels

`func (o *EventNacClient) GetUsermacLabels() []string`

GetUsermacLabels returns the UsermacLabels field if non-nil, zero value otherwise.

### GetUsermacLabelsOk

`func (o *EventNacClient) GetUsermacLabelsOk() (*[]string, bool)`

GetUsermacLabelsOk returns a tuple with the UsermacLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsermacLabels

`func (o *EventNacClient) SetUsermacLabels(v []string)`

SetUsermacLabels sets UsermacLabels field to given value.

### HasUsermacLabels

`func (o *EventNacClient) HasUsermacLabels() bool`

HasUsermacLabels returns a boolean if a field has been set.

### GetUsername

`func (o *EventNacClient) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EventNacClient) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EventNacClient) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EventNacClient) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVlan

`func (o *EventNacClient) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *EventNacClient) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *EventNacClient) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *EventNacClient) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanSource

`func (o *EventNacClient) GetVlanSource() string`

GetVlanSource returns the VlanSource field if non-nil, zero value otherwise.

### GetVlanSourceOk

`func (o *EventNacClient) GetVlanSourceOk() (*string, bool)`

GetVlanSourceOk returns a tuple with the VlanSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSource

`func (o *EventNacClient) SetVlanSource(v string)`

SetVlanSource sets VlanSource field to given value.

### HasVlanSource

`func (o *EventNacClient) HasVlanSource() bool`

HasVlanSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


