# ApPwrConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **float32** | additional power to request during negotiating with PSE over PoE, in mW | [optional] [default to 0]
**PreferUsbOverWifi** | Pointer to **bool** | whether to enable power out to peripheral, meanwhile will reduce power to wifi (only for AP45 at power mode) | [optional] [default to false]

## Methods

### NewApPwrConfig

`func NewApPwrConfig() *ApPwrConfig`

NewApPwrConfig instantiates a new ApPwrConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPwrConfigWithDefaults

`func NewApPwrConfigWithDefaults() *ApPwrConfig`

NewApPwrConfigWithDefaults instantiates a new ApPwrConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *ApPwrConfig) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ApPwrConfig) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ApPwrConfig) SetBase(v float32)`

SetBase sets Base field to given value.

### HasBase

`func (o *ApPwrConfig) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetPreferUsbOverWifi

`func (o *ApPwrConfig) GetPreferUsbOverWifi() bool`

GetPreferUsbOverWifi returns the PreferUsbOverWifi field if non-nil, zero value otherwise.

### GetPreferUsbOverWifiOk

`func (o *ApPwrConfig) GetPreferUsbOverWifiOk() (*bool, bool)`

GetPreferUsbOverWifiOk returns a tuple with the PreferUsbOverWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferUsbOverWifi

`func (o *ApPwrConfig) SetPreferUsbOverWifi(v bool)`

SetPreferUsbOverWifi sets PreferUsbOverWifi field to given value.

### HasPreferUsbOverWifi

`func (o *ApPwrConfig) HasPreferUsbOverWifi() bool`

HasPreferUsbOverWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


