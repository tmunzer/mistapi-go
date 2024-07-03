# Sdktemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgImage** | Pointer to **string** |  | [optional] 
**BtnFlrBgcolor** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | whether this is the default template when there are multiple templates | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**HeaderTxt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** | name for identification purpose | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SearchTxtcolor** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**WelcomeMsg** | Pointer to **string** |  | [optional] 

## Methods

### NewSdktemplate

`func NewSdktemplate(name string, ) *Sdktemplate`

NewSdktemplate instantiates a new Sdktemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdktemplateWithDefaults

`func NewSdktemplateWithDefaults() *Sdktemplate`

NewSdktemplateWithDefaults instantiates a new Sdktemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgImage

`func (o *Sdktemplate) GetBgImage() string`

GetBgImage returns the BgImage field if non-nil, zero value otherwise.

### GetBgImageOk

`func (o *Sdktemplate) GetBgImageOk() (*string, bool)`

GetBgImageOk returns a tuple with the BgImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImage

`func (o *Sdktemplate) SetBgImage(v string)`

SetBgImage sets BgImage field to given value.

### HasBgImage

`func (o *Sdktemplate) HasBgImage() bool`

HasBgImage returns a boolean if a field has been set.

### GetBtnFlrBgcolor

`func (o *Sdktemplate) GetBtnFlrBgcolor() string`

GetBtnFlrBgcolor returns the BtnFlrBgcolor field if non-nil, zero value otherwise.

### GetBtnFlrBgcolorOk

`func (o *Sdktemplate) GetBtnFlrBgcolorOk() (*string, bool)`

GetBtnFlrBgcolorOk returns a tuple with the BtnFlrBgcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtnFlrBgcolor

`func (o *Sdktemplate) SetBtnFlrBgcolor(v string)`

SetBtnFlrBgcolor sets BtnFlrBgcolor field to given value.

### HasBtnFlrBgcolor

`func (o *Sdktemplate) HasBtnFlrBgcolor() bool`

HasBtnFlrBgcolor returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Sdktemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Sdktemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Sdktemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Sdktemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefault

`func (o *Sdktemplate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Sdktemplate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Sdktemplate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Sdktemplate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetForSite

`func (o *Sdktemplate) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Sdktemplate) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Sdktemplate) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Sdktemplate) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeaderTxt

`func (o *Sdktemplate) GetHeaderTxt() string`

GetHeaderTxt returns the HeaderTxt field if non-nil, zero value otherwise.

### GetHeaderTxtOk

`func (o *Sdktemplate) GetHeaderTxtOk() (*string, bool)`

GetHeaderTxtOk returns a tuple with the HeaderTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTxt

`func (o *Sdktemplate) SetHeaderTxt(v string)`

SetHeaderTxt sets HeaderTxt field to given value.

### HasHeaderTxt

`func (o *Sdktemplate) HasHeaderTxt() bool`

HasHeaderTxt returns a boolean if a field has been set.

### GetId

`func (o *Sdktemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sdktemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sdktemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sdktemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Sdktemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Sdktemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Sdktemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Sdktemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Sdktemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sdktemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sdktemplate) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *Sdktemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Sdktemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Sdktemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Sdktemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSearchTxtcolor

`func (o *Sdktemplate) GetSearchTxtcolor() string`

GetSearchTxtcolor returns the SearchTxtcolor field if non-nil, zero value otherwise.

### GetSearchTxtcolorOk

`func (o *Sdktemplate) GetSearchTxtcolorOk() (*string, bool)`

GetSearchTxtcolorOk returns a tuple with the SearchTxtcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTxtcolor

`func (o *Sdktemplate) SetSearchTxtcolor(v string)`

SetSearchTxtcolor sets SearchTxtcolor field to given value.

### HasSearchTxtcolor

`func (o *Sdktemplate) HasSearchTxtcolor() bool`

HasSearchTxtcolor returns a boolean if a field has been set.

### GetSiteId

`func (o *Sdktemplate) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Sdktemplate) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Sdktemplate) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Sdktemplate) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetWelcomeMsg

`func (o *Sdktemplate) GetWelcomeMsg() string`

GetWelcomeMsg returns the WelcomeMsg field if non-nil, zero value otherwise.

### GetWelcomeMsgOk

`func (o *Sdktemplate) GetWelcomeMsgOk() (*string, bool)`

GetWelcomeMsgOk returns a tuple with the WelcomeMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMsg

`func (o *Sdktemplate) SetWelcomeMsg(v string)`

SetWelcomeMsg sets WelcomeMsg field to given value.

### HasWelcomeMsg

`func (o *Sdktemplate) HasWelcomeMsg() bool`

HasWelcomeMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


