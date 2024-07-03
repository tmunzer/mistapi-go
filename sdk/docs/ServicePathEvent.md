# ServicePathEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Policy** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VpnName** | Pointer to **string** |  | [optional] 
**VpnPath** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePathEvent

`func NewServicePathEvent() *ServicePathEvent`

NewServicePathEvent instantiates a new ServicePathEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePathEventWithDefaults

`func NewServicePathEventWithDefaults() *ServicePathEvent`

NewServicePathEventWithDefaults instantiates a new ServicePathEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ServicePathEvent) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ServicePathEvent) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ServicePathEvent) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ServicePathEvent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ServicePathEvent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServicePathEvent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServicePathEvent) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ServicePathEvent) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrgId

`func (o *ServicePathEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ServicePathEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ServicePathEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ServicePathEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPolicy

`func (o *ServicePathEvent) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ServicePathEvent) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ServicePathEvent) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ServicePathEvent) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPortId

`func (o *ServicePathEvent) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *ServicePathEvent) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *ServicePathEvent) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *ServicePathEvent) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSiteId

`func (o *ServicePathEvent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ServicePathEvent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ServicePathEvent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ServicePathEvent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetText

`func (o *ServicePathEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ServicePathEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ServicePathEvent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ServicePathEvent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServicePathEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServicePathEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServicePathEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServicePathEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *ServicePathEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePathEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePathEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServicePathEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ServicePathEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServicePathEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServicePathEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServicePathEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVpnName

`func (o *ServicePathEvent) GetVpnName() string`

GetVpnName returns the VpnName field if non-nil, zero value otherwise.

### GetVpnNameOk

`func (o *ServicePathEvent) GetVpnNameOk() (*string, bool)`

GetVpnNameOk returns a tuple with the VpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnName

`func (o *ServicePathEvent) SetVpnName(v string)`

SetVpnName sets VpnName field to given value.

### HasVpnName

`func (o *ServicePathEvent) HasVpnName() bool`

HasVpnName returns a boolean if a field has been set.

### GetVpnPath

`func (o *ServicePathEvent) GetVpnPath() string`

GetVpnPath returns the VpnPath field if non-nil, zero value otherwise.

### GetVpnPathOk

`func (o *ServicePathEvent) GetVpnPathOk() (*string, bool)`

GetVpnPathOk returns a tuple with the VpnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPath

`func (o *ServicePathEvent) SetVpnPath(v string)`

SetVpnPath sets VpnPath field to given value.

### HasVpnPath

`func (o *ServicePathEvent) HasVpnPath() bool`

HasVpnPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


