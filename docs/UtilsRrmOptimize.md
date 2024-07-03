# UtilsRrmOptimize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bands** | **[]string** | list of bands | 
**Macs** | Pointer to **[]string** | targeting AP (neighbor APs may get changed, too), default is empty for ALL APs | [optional] 
**TxpowerOnly** | Pointer to **bool** | only changng TX Power (will not disconnect clients) | [optional] [default to false]

## Methods

### NewUtilsRrmOptimize

`func NewUtilsRrmOptimize(bands []string, ) *UtilsRrmOptimize`

NewUtilsRrmOptimize instantiates a new UtilsRrmOptimize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsRrmOptimizeWithDefaults

`func NewUtilsRrmOptimizeWithDefaults() *UtilsRrmOptimize`

NewUtilsRrmOptimizeWithDefaults instantiates a new UtilsRrmOptimize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBands

`func (o *UtilsRrmOptimize) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *UtilsRrmOptimize) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *UtilsRrmOptimize) SetBands(v []string)`

SetBands sets Bands field to given value.


### GetMacs

`func (o *UtilsRrmOptimize) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *UtilsRrmOptimize) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *UtilsRrmOptimize) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *UtilsRrmOptimize) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetTxpowerOnly

`func (o *UtilsRrmOptimize) GetTxpowerOnly() bool`

GetTxpowerOnly returns the TxpowerOnly field if non-nil, zero value otherwise.

### GetTxpowerOnlyOk

`func (o *UtilsRrmOptimize) GetTxpowerOnlyOk() (*bool, bool)`

GetTxpowerOnlyOk returns a tuple with the TxpowerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxpowerOnly

`func (o *UtilsRrmOptimize) SetTxpowerOnly(v bool)`

SetTxpowerOnly sets TxpowerOnly field to given value.

### HasTxpowerOnly

`func (o *UtilsRrmOptimize) HasTxpowerOnly() bool`

HasTxpowerOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


