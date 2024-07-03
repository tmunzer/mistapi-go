# RoutingPolicyTermAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | Pointer to **bool** |  | [optional] 
**AddCommunity** | Pointer to **[]string** |  | [optional] 
**AddTargetVrfs** | Pointer to **[]string** | for SSR, hub decides how VRF routes are leaked on spoke | [optional] 
**Community** | Pointer to **[]string** | when used as export policy, optional | [optional] 
**ExcludeAsPath** | Pointer to **[]string** | when used as export policy, optional. To exclude certain AS | [optional] 
**ExcludeCommunity** | Pointer to **[]string** |  | [optional] 
**ExportCommunitites** | Pointer to **[]string** | when used as export policy, optional | [optional] 
**LocalPreference** | Pointer to **string** | optional, for an import policy, local_preference can be changed | [optional] 
**PrependAsPath** | Pointer to **[]string** | when used as export policy, optional. By default, the local AS will be prepended, to change it | [optional] 

## Methods

### NewRoutingPolicyTermAction

`func NewRoutingPolicyTermAction() *RoutingPolicyTermAction`

NewRoutingPolicyTermAction instantiates a new RoutingPolicyTermAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyTermActionWithDefaults

`func NewRoutingPolicyTermActionWithDefaults() *RoutingPolicyTermAction`

NewRoutingPolicyTermActionWithDefaults instantiates a new RoutingPolicyTermAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccept

`func (o *RoutingPolicyTermAction) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *RoutingPolicyTermAction) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *RoutingPolicyTermAction) SetAccept(v bool)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *RoutingPolicyTermAction) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetAddCommunity

`func (o *RoutingPolicyTermAction) GetAddCommunity() []string`

GetAddCommunity returns the AddCommunity field if non-nil, zero value otherwise.

### GetAddCommunityOk

`func (o *RoutingPolicyTermAction) GetAddCommunityOk() (*[]string, bool)`

GetAddCommunityOk returns a tuple with the AddCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCommunity

`func (o *RoutingPolicyTermAction) SetAddCommunity(v []string)`

SetAddCommunity sets AddCommunity field to given value.

### HasAddCommunity

`func (o *RoutingPolicyTermAction) HasAddCommunity() bool`

HasAddCommunity returns a boolean if a field has been set.

### GetAddTargetVrfs

`func (o *RoutingPolicyTermAction) GetAddTargetVrfs() []string`

GetAddTargetVrfs returns the AddTargetVrfs field if non-nil, zero value otherwise.

### GetAddTargetVrfsOk

`func (o *RoutingPolicyTermAction) GetAddTargetVrfsOk() (*[]string, bool)`

GetAddTargetVrfsOk returns a tuple with the AddTargetVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTargetVrfs

`func (o *RoutingPolicyTermAction) SetAddTargetVrfs(v []string)`

SetAddTargetVrfs sets AddTargetVrfs field to given value.

### HasAddTargetVrfs

`func (o *RoutingPolicyTermAction) HasAddTargetVrfs() bool`

HasAddTargetVrfs returns a boolean if a field has been set.

### GetCommunity

`func (o *RoutingPolicyTermAction) GetCommunity() []string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *RoutingPolicyTermAction) GetCommunityOk() (*[]string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *RoutingPolicyTermAction) SetCommunity(v []string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *RoutingPolicyTermAction) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetExcludeAsPath

`func (o *RoutingPolicyTermAction) GetExcludeAsPath() []string`

GetExcludeAsPath returns the ExcludeAsPath field if non-nil, zero value otherwise.

### GetExcludeAsPathOk

`func (o *RoutingPolicyTermAction) GetExcludeAsPathOk() (*[]string, bool)`

GetExcludeAsPathOk returns a tuple with the ExcludeAsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAsPath

`func (o *RoutingPolicyTermAction) SetExcludeAsPath(v []string)`

SetExcludeAsPath sets ExcludeAsPath field to given value.

### HasExcludeAsPath

`func (o *RoutingPolicyTermAction) HasExcludeAsPath() bool`

HasExcludeAsPath returns a boolean if a field has been set.

### GetExcludeCommunity

`func (o *RoutingPolicyTermAction) GetExcludeCommunity() []string`

GetExcludeCommunity returns the ExcludeCommunity field if non-nil, zero value otherwise.

### GetExcludeCommunityOk

`func (o *RoutingPolicyTermAction) GetExcludeCommunityOk() (*[]string, bool)`

GetExcludeCommunityOk returns a tuple with the ExcludeCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCommunity

`func (o *RoutingPolicyTermAction) SetExcludeCommunity(v []string)`

SetExcludeCommunity sets ExcludeCommunity field to given value.

### HasExcludeCommunity

`func (o *RoutingPolicyTermAction) HasExcludeCommunity() bool`

HasExcludeCommunity returns a boolean if a field has been set.

### GetExportCommunitites

`func (o *RoutingPolicyTermAction) GetExportCommunitites() []string`

GetExportCommunitites returns the ExportCommunitites field if non-nil, zero value otherwise.

### GetExportCommunititesOk

`func (o *RoutingPolicyTermAction) GetExportCommunititesOk() (*[]string, bool)`

GetExportCommunititesOk returns a tuple with the ExportCommunitites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportCommunitites

`func (o *RoutingPolicyTermAction) SetExportCommunitites(v []string)`

SetExportCommunitites sets ExportCommunitites field to given value.

### HasExportCommunitites

`func (o *RoutingPolicyTermAction) HasExportCommunitites() bool`

HasExportCommunitites returns a boolean if a field has been set.

### GetLocalPreference

`func (o *RoutingPolicyTermAction) GetLocalPreference() string`

GetLocalPreference returns the LocalPreference field if non-nil, zero value otherwise.

### GetLocalPreferenceOk

`func (o *RoutingPolicyTermAction) GetLocalPreferenceOk() (*string, bool)`

GetLocalPreferenceOk returns a tuple with the LocalPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPreference

`func (o *RoutingPolicyTermAction) SetLocalPreference(v string)`

SetLocalPreference sets LocalPreference field to given value.

### HasLocalPreference

`func (o *RoutingPolicyTermAction) HasLocalPreference() bool`

HasLocalPreference returns a boolean if a field has been set.

### GetPrependAsPath

`func (o *RoutingPolicyTermAction) GetPrependAsPath() []string`

GetPrependAsPath returns the PrependAsPath field if non-nil, zero value otherwise.

### GetPrependAsPathOk

`func (o *RoutingPolicyTermAction) GetPrependAsPathOk() (*[]string, bool)`

GetPrependAsPathOk returns a tuple with the PrependAsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependAsPath

`func (o *RoutingPolicyTermAction) SetPrependAsPath(v []string)`

SetPrependAsPath sets PrependAsPath field to given value.

### HasPrependAsPath

`func (o *RoutingPolicyTermAction) HasPrependAsPath() bool`

HasPrependAsPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


