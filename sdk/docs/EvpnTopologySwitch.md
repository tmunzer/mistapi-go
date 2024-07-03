# EvpnTopologySwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ModelSwitch**](ModelSwitch.md) |  | [optional] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**Downlinks** | Pointer to **[]string** |  | [optional] [readonly] 
**Esilaglinks** | Pointer to **[]string** |  | [optional] 
**EvpnId** | Pointer to **int32** |  | [optional] [readonly] 
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Pod** | Pointer to **int32** | optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.  * for CLOS, to group dist / access switches into pods * for ERB/CRB, to group dist / esilag-access into pods | [optional] [default to 1]
**Pods** | Pointer to **[]int32** | by default, core switches are assumed to be connecting all pods.  if you want to limit the pods, you can specify pods. | [optional] 
**Role** | Pointer to [**EvpnTopologySwitchRole**](EvpnTopologySwitchRole.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SuggestedDownlinks** | Pointer to **[]string** |  | [optional] [readonly] 
**SuggestedEsilaglinks** | Pointer to **[]string** |  | [optional] [readonly] 
**SuggestedUplinks** | Pointer to **[]string** |  | [optional] [readonly] 
**Uplinks** | Pointer to **[]string** | if not specified in the request, suggested ones will be used | [optional] [readonly] 

## Methods

### NewEvpnTopologySwitch

`func NewEvpnTopologySwitch() *EvpnTopologySwitch`

NewEvpnTopologySwitch instantiates a new EvpnTopologySwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvpnTopologySwitchWithDefaults

`func NewEvpnTopologySwitchWithDefaults() *EvpnTopologySwitch`

NewEvpnTopologySwitchWithDefaults instantiates a new EvpnTopologySwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *EvpnTopologySwitch) GetConfig() ModelSwitch`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EvpnTopologySwitch) GetConfigOk() (*ModelSwitch, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EvpnTopologySwitch) SetConfig(v ModelSwitch)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EvpnTopologySwitch) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *EvpnTopologySwitch) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *EvpnTopologySwitch) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *EvpnTopologySwitch) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *EvpnTopologySwitch) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDownlinks

`func (o *EvpnTopologySwitch) GetDownlinks() []string`

GetDownlinks returns the Downlinks field if non-nil, zero value otherwise.

### GetDownlinksOk

`func (o *EvpnTopologySwitch) GetDownlinksOk() (*[]string, bool)`

GetDownlinksOk returns a tuple with the Downlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinks

`func (o *EvpnTopologySwitch) SetDownlinks(v []string)`

SetDownlinks sets Downlinks field to given value.

### HasDownlinks

`func (o *EvpnTopologySwitch) HasDownlinks() bool`

HasDownlinks returns a boolean if a field has been set.

### GetEsilaglinks

`func (o *EvpnTopologySwitch) GetEsilaglinks() []string`

GetEsilaglinks returns the Esilaglinks field if non-nil, zero value otherwise.

### GetEsilaglinksOk

`func (o *EvpnTopologySwitch) GetEsilaglinksOk() (*[]string, bool)`

GetEsilaglinksOk returns a tuple with the Esilaglinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsilaglinks

`func (o *EvpnTopologySwitch) SetEsilaglinks(v []string)`

SetEsilaglinks sets Esilaglinks field to given value.

### HasEsilaglinks

`func (o *EvpnTopologySwitch) HasEsilaglinks() bool`

HasEsilaglinks returns a boolean if a field has been set.

### GetEvpnId

`func (o *EvpnTopologySwitch) GetEvpnId() int32`

GetEvpnId returns the EvpnId field if non-nil, zero value otherwise.

### GetEvpnIdOk

`func (o *EvpnTopologySwitch) GetEvpnIdOk() (*int32, bool)`

GetEvpnIdOk returns a tuple with the EvpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnId

`func (o *EvpnTopologySwitch) SetEvpnId(v int32)`

SetEvpnId sets EvpnId field to given value.

### HasEvpnId

`func (o *EvpnTopologySwitch) HasEvpnId() bool`

HasEvpnId returns a boolean if a field has been set.

### GetMac

`func (o *EvpnTopologySwitch) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EvpnTopologySwitch) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EvpnTopologySwitch) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *EvpnTopologySwitch) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *EvpnTopologySwitch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvpnTopologySwitch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvpnTopologySwitch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvpnTopologySwitch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPod

`func (o *EvpnTopologySwitch) GetPod() int32`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *EvpnTopologySwitch) GetPodOk() (*int32, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *EvpnTopologySwitch) SetPod(v int32)`

SetPod sets Pod field to given value.

### HasPod

`func (o *EvpnTopologySwitch) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPods

`func (o *EvpnTopologySwitch) GetPods() []int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *EvpnTopologySwitch) GetPodsOk() (*[]int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *EvpnTopologySwitch) SetPods(v []int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *EvpnTopologySwitch) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetRole

`func (o *EvpnTopologySwitch) GetRole() EvpnTopologySwitchRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EvpnTopologySwitch) GetRoleOk() (*EvpnTopologySwitchRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EvpnTopologySwitch) SetRole(v EvpnTopologySwitchRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *EvpnTopologySwitch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSiteId

`func (o *EvpnTopologySwitch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EvpnTopologySwitch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EvpnTopologySwitch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *EvpnTopologySwitch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSuggestedDownlinks

`func (o *EvpnTopologySwitch) GetSuggestedDownlinks() []string`

GetSuggestedDownlinks returns the SuggestedDownlinks field if non-nil, zero value otherwise.

### GetSuggestedDownlinksOk

`func (o *EvpnTopologySwitch) GetSuggestedDownlinksOk() (*[]string, bool)`

GetSuggestedDownlinksOk returns a tuple with the SuggestedDownlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDownlinks

`func (o *EvpnTopologySwitch) SetSuggestedDownlinks(v []string)`

SetSuggestedDownlinks sets SuggestedDownlinks field to given value.

### HasSuggestedDownlinks

`func (o *EvpnTopologySwitch) HasSuggestedDownlinks() bool`

HasSuggestedDownlinks returns a boolean if a field has been set.

### GetSuggestedEsilaglinks

`func (o *EvpnTopologySwitch) GetSuggestedEsilaglinks() []string`

GetSuggestedEsilaglinks returns the SuggestedEsilaglinks field if non-nil, zero value otherwise.

### GetSuggestedEsilaglinksOk

`func (o *EvpnTopologySwitch) GetSuggestedEsilaglinksOk() (*[]string, bool)`

GetSuggestedEsilaglinksOk returns a tuple with the SuggestedEsilaglinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedEsilaglinks

`func (o *EvpnTopologySwitch) SetSuggestedEsilaglinks(v []string)`

SetSuggestedEsilaglinks sets SuggestedEsilaglinks field to given value.

### HasSuggestedEsilaglinks

`func (o *EvpnTopologySwitch) HasSuggestedEsilaglinks() bool`

HasSuggestedEsilaglinks returns a boolean if a field has been set.

### GetSuggestedUplinks

`func (o *EvpnTopologySwitch) GetSuggestedUplinks() []string`

GetSuggestedUplinks returns the SuggestedUplinks field if non-nil, zero value otherwise.

### GetSuggestedUplinksOk

`func (o *EvpnTopologySwitch) GetSuggestedUplinksOk() (*[]string, bool)`

GetSuggestedUplinksOk returns a tuple with the SuggestedUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedUplinks

`func (o *EvpnTopologySwitch) SetSuggestedUplinks(v []string)`

SetSuggestedUplinks sets SuggestedUplinks field to given value.

### HasSuggestedUplinks

`func (o *EvpnTopologySwitch) HasSuggestedUplinks() bool`

HasSuggestedUplinks returns a boolean if a field has been set.

### GetUplinks

`func (o *EvpnTopologySwitch) GetUplinks() []string`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *EvpnTopologySwitch) GetUplinksOk() (*[]string, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *EvpnTopologySwitch) SetUplinks(v []string)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *EvpnTopologySwitch) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


