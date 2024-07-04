# ConstCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha2** | **string** | country code, in two-character | 
**Certified** | **bool** |  | 
**Name** | **string** |  | 
**Numeric** | **float32** | country code, ISO 3166-1 numeric | 

## Methods

### NewConstCountry

`func NewConstCountry(alpha2 string, certified bool, name string, numeric float32, ) *ConstCountry`

NewConstCountry instantiates a new ConstCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstCountryWithDefaults

`func NewConstCountryWithDefaults() *ConstCountry`

NewConstCountryWithDefaults instantiates a new ConstCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha2

`func (o *ConstCountry) GetAlpha2() string`

GetAlpha2 returns the Alpha2 field if non-nil, zero value otherwise.

### GetAlpha2Ok

`func (o *ConstCountry) GetAlpha2Ok() (*string, bool)`

GetAlpha2Ok returns a tuple with the Alpha2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha2

`func (o *ConstCountry) SetAlpha2(v string)`

SetAlpha2 sets Alpha2 field to given value.


### GetCertified

`func (o *ConstCountry) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *ConstCountry) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *ConstCountry) SetCertified(v bool)`

SetCertified sets Certified field to given value.


### GetName

`func (o *ConstCountry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConstCountry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConstCountry) SetName(v string)`

SetName sets Name field to given value.


### GetNumeric

`func (o *ConstCountry) GetNumeric() float32`

GetNumeric returns the Numeric field if non-nil, zero value otherwise.

### GetNumericOk

`func (o *ConstCountry) GetNumericOk() (*float32, bool)`

GetNumericOk returns a tuple with the Numeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeric

`func (o *ConstCountry) SetNumeric(v float32)`

SetNumeric sets Numeric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


