# RfTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntGain24** | Pointer to **int32** |  | [optional] 
**AntGain5** | Pointer to **int32** |  | [optional] 
**AntGain6** | Pointer to **int32** |  | [optional] 
**Band24** | Pointer to [**RftemplateRadioBand24**](RftemplateRadioBand24.md) |  | [optional] 
**Band24Usage** | Pointer to [**RadioBand24Usage**](RadioBand24Usage.md) |  | [optional] 
**Band5** | Pointer to [**RftemplateRadioBand5**](RftemplateRadioBand5.md) |  | [optional] 
**Band5On24Radio** | Pointer to [**RftemplateRadioBand5**](RftemplateRadioBand5.md) |  | [optional] 
**Band6** | Pointer to [**RftemplateRadioBand6**](RftemplateRadioBand6.md) |  | [optional] 
**CountryCode** | Pointer to **string** | optional, country code to use. If specified, this gets applied to all sites using the RF Template | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModelSpecific** | Pointer to [**map[string]RfTemplateModelSpecificProperty**](RfTemplateModelSpecificProperty.md) | overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. \&quot;AP63\&quot;) | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** | The name of the RF template | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**ScanningEnabled** | Pointer to **bool** | whether scanning radio is enabled | [optional] 

## Methods

### NewRfTemplate

`func NewRfTemplate(name string, ) *RfTemplate`

NewRfTemplate instantiates a new RfTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfTemplateWithDefaults

`func NewRfTemplateWithDefaults() *RfTemplate`

NewRfTemplateWithDefaults instantiates a new RfTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntGain24

`func (o *RfTemplate) GetAntGain24() int32`

GetAntGain24 returns the AntGain24 field if non-nil, zero value otherwise.

### GetAntGain24Ok

`func (o *RfTemplate) GetAntGain24Ok() (*int32, bool)`

GetAntGain24Ok returns a tuple with the AntGain24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain24

`func (o *RfTemplate) SetAntGain24(v int32)`

SetAntGain24 sets AntGain24 field to given value.

### HasAntGain24

`func (o *RfTemplate) HasAntGain24() bool`

HasAntGain24 returns a boolean if a field has been set.

### GetAntGain5

`func (o *RfTemplate) GetAntGain5() int32`

GetAntGain5 returns the AntGain5 field if non-nil, zero value otherwise.

### GetAntGain5Ok

`func (o *RfTemplate) GetAntGain5Ok() (*int32, bool)`

GetAntGain5Ok returns a tuple with the AntGain5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain5

`func (o *RfTemplate) SetAntGain5(v int32)`

SetAntGain5 sets AntGain5 field to given value.

### HasAntGain5

`func (o *RfTemplate) HasAntGain5() bool`

HasAntGain5 returns a boolean if a field has been set.

### GetAntGain6

`func (o *RfTemplate) GetAntGain6() int32`

GetAntGain6 returns the AntGain6 field if non-nil, zero value otherwise.

### GetAntGain6Ok

`func (o *RfTemplate) GetAntGain6Ok() (*int32, bool)`

GetAntGain6Ok returns a tuple with the AntGain6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain6

`func (o *RfTemplate) SetAntGain6(v int32)`

SetAntGain6 sets AntGain6 field to given value.

### HasAntGain6

`func (o *RfTemplate) HasAntGain6() bool`

HasAntGain6 returns a boolean if a field has been set.

### GetBand24

`func (o *RfTemplate) GetBand24() RftemplateRadioBand24`

GetBand24 returns the Band24 field if non-nil, zero value otherwise.

### GetBand24Ok

`func (o *RfTemplate) GetBand24Ok() (*RftemplateRadioBand24, bool)`

GetBand24Ok returns a tuple with the Band24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24

`func (o *RfTemplate) SetBand24(v RftemplateRadioBand24)`

SetBand24 sets Band24 field to given value.

### HasBand24

`func (o *RfTemplate) HasBand24() bool`

HasBand24 returns a boolean if a field has been set.

### GetBand24Usage

`func (o *RfTemplate) GetBand24Usage() RadioBand24Usage`

GetBand24Usage returns the Band24Usage field if non-nil, zero value otherwise.

### GetBand24UsageOk

`func (o *RfTemplate) GetBand24UsageOk() (*RadioBand24Usage, bool)`

GetBand24UsageOk returns a tuple with the Band24Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Usage

`func (o *RfTemplate) SetBand24Usage(v RadioBand24Usage)`

SetBand24Usage sets Band24Usage field to given value.

### HasBand24Usage

`func (o *RfTemplate) HasBand24Usage() bool`

HasBand24Usage returns a boolean if a field has been set.

### GetBand5

`func (o *RfTemplate) GetBand5() RftemplateRadioBand5`

GetBand5 returns the Band5 field if non-nil, zero value otherwise.

### GetBand5Ok

`func (o *RfTemplate) GetBand5Ok() (*RftemplateRadioBand5, bool)`

GetBand5Ok returns a tuple with the Band5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5

`func (o *RfTemplate) SetBand5(v RftemplateRadioBand5)`

SetBand5 sets Band5 field to given value.

### HasBand5

`func (o *RfTemplate) HasBand5() bool`

HasBand5 returns a boolean if a field has been set.

### GetBand5On24Radio

`func (o *RfTemplate) GetBand5On24Radio() RftemplateRadioBand5`

GetBand5On24Radio returns the Band5On24Radio field if non-nil, zero value otherwise.

### GetBand5On24RadioOk

`func (o *RfTemplate) GetBand5On24RadioOk() (*RftemplateRadioBand5, bool)`

GetBand5On24RadioOk returns a tuple with the Band5On24Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5On24Radio

`func (o *RfTemplate) SetBand5On24Radio(v RftemplateRadioBand5)`

SetBand5On24Radio sets Band5On24Radio field to given value.

### HasBand5On24Radio

`func (o *RfTemplate) HasBand5On24Radio() bool`

HasBand5On24Radio returns a boolean if a field has been set.

### GetBand6

`func (o *RfTemplate) GetBand6() RftemplateRadioBand6`

GetBand6 returns the Band6 field if non-nil, zero value otherwise.

### GetBand6Ok

`func (o *RfTemplate) GetBand6Ok() (*RftemplateRadioBand6, bool)`

GetBand6Ok returns a tuple with the Band6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6

`func (o *RfTemplate) SetBand6(v RftemplateRadioBand6)`

SetBand6 sets Band6 field to given value.

### HasBand6

`func (o *RfTemplate) HasBand6() bool`

HasBand6 returns a boolean if a field has been set.

### GetCountryCode

`func (o *RfTemplate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RfTemplate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RfTemplate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *RfTemplate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *RfTemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RfTemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RfTemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *RfTemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *RfTemplate) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *RfTemplate) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *RfTemplate) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *RfTemplate) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *RfTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RfTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RfTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RfTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelSpecific

`func (o *RfTemplate) GetModelSpecific() map[string]RfTemplateModelSpecificProperty`

GetModelSpecific returns the ModelSpecific field if non-nil, zero value otherwise.

### GetModelSpecificOk

`func (o *RfTemplate) GetModelSpecificOk() (*map[string]RfTemplateModelSpecificProperty, bool)`

GetModelSpecificOk returns a tuple with the ModelSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSpecific

`func (o *RfTemplate) SetModelSpecific(v map[string]RfTemplateModelSpecificProperty)`

SetModelSpecific sets ModelSpecific field to given value.

### HasModelSpecific

`func (o *RfTemplate) HasModelSpecific() bool`

HasModelSpecific returns a boolean if a field has been set.

### GetModifiedTime

`func (o *RfTemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *RfTemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *RfTemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *RfTemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *RfTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RfTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RfTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *RfTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RfTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RfTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RfTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetScanningEnabled

`func (o *RfTemplate) GetScanningEnabled() bool`

GetScanningEnabled returns the ScanningEnabled field if non-nil, zero value otherwise.

### GetScanningEnabledOk

`func (o *RfTemplate) GetScanningEnabledOk() (*bool, bool)`

GetScanningEnabledOk returns a tuple with the ScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningEnabled

`func (o *RfTemplate) SetScanningEnabled(v bool)`

SetScanningEnabled sets ScanningEnabled field to given value.

### HasScanningEnabled

`func (o *RfTemplate) HasScanningEnabled() bool`

HasScanningEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


