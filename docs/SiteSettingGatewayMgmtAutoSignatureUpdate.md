# SiteSettingGatewayMgmtAutoSignatureUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**TimeOfDay** | Pointer to **string** | optional, Mist will decide the timing | [optional] 

## Methods

### NewSiteSettingGatewayMgmtAutoSignatureUpdate

`func NewSiteSettingGatewayMgmtAutoSignatureUpdate() *SiteSettingGatewayMgmtAutoSignatureUpdate`

NewSiteSettingGatewayMgmtAutoSignatureUpdate instantiates a new SiteSettingGatewayMgmtAutoSignatureUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingGatewayMgmtAutoSignatureUpdateWithDefaults

`func NewSiteSettingGatewayMgmtAutoSignatureUpdateWithDefaults() *SiteSettingGatewayMgmtAutoSignatureUpdate`

NewSiteSettingGatewayMgmtAutoSignatureUpdateWithDefaults instantiates a new SiteSettingGatewayMgmtAutoSignatureUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEnable

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *SiteSettingGatewayMgmtAutoSignatureUpdate) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


