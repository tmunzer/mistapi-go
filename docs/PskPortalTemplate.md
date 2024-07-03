# PskPortalTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**PskPortalTemplateAlignment**](PskPortalTemplateAlignment.md) |  | [optional] [default to PSKPORTALTEMPLATEALIGNMENT_CENTER]
**Color** | Pointer to **string** |  | [optional] [default to "#1074bc"]
**Logo** | Pointer to **NullableString** | custom logo.  default null, uses Juniper Mist Logo | [optional] 
**PoweredBy** | Pointer to **bool** | whether to hide \&quot;Powered by Juniper Mist\&quot; and email footers | [optional] [default to false]

## Methods

### NewPskPortalTemplate

`func NewPskPortalTemplate() *PskPortalTemplate`

NewPskPortalTemplate instantiates a new PskPortalTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskPortalTemplateWithDefaults

`func NewPskPortalTemplateWithDefaults() *PskPortalTemplate`

NewPskPortalTemplateWithDefaults instantiates a new PskPortalTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *PskPortalTemplate) GetAlignment() PskPortalTemplateAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *PskPortalTemplate) GetAlignmentOk() (*PskPortalTemplateAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *PskPortalTemplate) SetAlignment(v PskPortalTemplateAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *PskPortalTemplate) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetColor

`func (o *PskPortalTemplate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PskPortalTemplate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PskPortalTemplate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PskPortalTemplate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLogo

`func (o *PskPortalTemplate) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PskPortalTemplate) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PskPortalTemplate) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PskPortalTemplate) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *PskPortalTemplate) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *PskPortalTemplate) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetPoweredBy

`func (o *PskPortalTemplate) GetPoweredBy() bool`

GetPoweredBy returns the PoweredBy field if non-nil, zero value otherwise.

### GetPoweredByOk

`func (o *PskPortalTemplate) GetPoweredByOk() (*bool, bool)`

GetPoweredByOk returns a tuple with the PoweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredBy

`func (o *PskPortalTemplate) SetPoweredBy(v bool)`

SetPoweredBy sets PoweredBy field to given value.

### HasPoweredBy

`func (o *PskPortalTemplate) HasPoweredBy() bool`

HasPoweredBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


