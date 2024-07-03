# ResponseClaimLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryAdded** | [**[]ResponseClaimLicenseInventoryItem**](ResponseClaimLicenseInventoryItem.md) |  | 
**InventoryDuplicated** | [**[]ResponseClaimLicenseInventoryItem**](ResponseClaimLicenseInventoryItem.md) |  | 
**LicenseAdded** | [**[]ResponseClaimLicenseLicenseItem**](ResponseClaimLicenseLicenseItem.md) |  | 
**LicenseDuplicated** | [**[]ResponseClaimLicenseLicenseItem**](ResponseClaimLicenseLicenseItem.md) |  | 
**LicenseError** | [**[]ResponseClaimLicenseLicenseErrorItem**](ResponseClaimLicenseLicenseErrorItem.md) |  | 

## Methods

### NewResponseClaimLicense

`func NewResponseClaimLicense(inventoryAdded []ResponseClaimLicenseInventoryItem, inventoryDuplicated []ResponseClaimLicenseInventoryItem, licenseAdded []ResponseClaimLicenseLicenseItem, licenseDuplicated []ResponseClaimLicenseLicenseItem, licenseError []ResponseClaimLicenseLicenseErrorItem, ) *ResponseClaimLicense`

NewResponseClaimLicense instantiates a new ResponseClaimLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseClaimLicenseWithDefaults

`func NewResponseClaimLicenseWithDefaults() *ResponseClaimLicense`

NewResponseClaimLicenseWithDefaults instantiates a new ResponseClaimLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryAdded

`func (o *ResponseClaimLicense) GetInventoryAdded() []ResponseClaimLicenseInventoryItem`

GetInventoryAdded returns the InventoryAdded field if non-nil, zero value otherwise.

### GetInventoryAddedOk

`func (o *ResponseClaimLicense) GetInventoryAddedOk() (*[]ResponseClaimLicenseInventoryItem, bool)`

GetInventoryAddedOk returns a tuple with the InventoryAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAdded

`func (o *ResponseClaimLicense) SetInventoryAdded(v []ResponseClaimLicenseInventoryItem)`

SetInventoryAdded sets InventoryAdded field to given value.


### GetInventoryDuplicated

`func (o *ResponseClaimLicense) GetInventoryDuplicated() []ResponseClaimLicenseInventoryItem`

GetInventoryDuplicated returns the InventoryDuplicated field if non-nil, zero value otherwise.

### GetInventoryDuplicatedOk

`func (o *ResponseClaimLicense) GetInventoryDuplicatedOk() (*[]ResponseClaimLicenseInventoryItem, bool)`

GetInventoryDuplicatedOk returns a tuple with the InventoryDuplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDuplicated

`func (o *ResponseClaimLicense) SetInventoryDuplicated(v []ResponseClaimLicenseInventoryItem)`

SetInventoryDuplicated sets InventoryDuplicated field to given value.


### GetLicenseAdded

`func (o *ResponseClaimLicense) GetLicenseAdded() []ResponseClaimLicenseLicenseItem`

GetLicenseAdded returns the LicenseAdded field if non-nil, zero value otherwise.

### GetLicenseAddedOk

`func (o *ResponseClaimLicense) GetLicenseAddedOk() (*[]ResponseClaimLicenseLicenseItem, bool)`

GetLicenseAddedOk returns a tuple with the LicenseAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAdded

`func (o *ResponseClaimLicense) SetLicenseAdded(v []ResponseClaimLicenseLicenseItem)`

SetLicenseAdded sets LicenseAdded field to given value.


### GetLicenseDuplicated

`func (o *ResponseClaimLicense) GetLicenseDuplicated() []ResponseClaimLicenseLicenseItem`

GetLicenseDuplicated returns the LicenseDuplicated field if non-nil, zero value otherwise.

### GetLicenseDuplicatedOk

`func (o *ResponseClaimLicense) GetLicenseDuplicatedOk() (*[]ResponseClaimLicenseLicenseItem, bool)`

GetLicenseDuplicatedOk returns a tuple with the LicenseDuplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseDuplicated

`func (o *ResponseClaimLicense) SetLicenseDuplicated(v []ResponseClaimLicenseLicenseItem)`

SetLicenseDuplicated sets LicenseDuplicated field to given value.


### GetLicenseError

`func (o *ResponseClaimLicense) GetLicenseError() []ResponseClaimLicenseLicenseErrorItem`

GetLicenseError returns the LicenseError field if non-nil, zero value otherwise.

### GetLicenseErrorOk

`func (o *ResponseClaimLicense) GetLicenseErrorOk() (*[]ResponseClaimLicenseLicenseErrorItem, bool)`

GetLicenseErrorOk returns a tuple with the LicenseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseError

`func (o *ResponseClaimLicense) SetLicenseError(v []ResponseClaimLicenseLicenseErrorItem)`

SetLicenseError sets LicenseError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


