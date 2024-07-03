# JseInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractEndTime** | Pointer to **int32** |  | [optional] 
**ContractId** | Pointer to **string** |  | [optional] 
**ContractSku** | Pointer to **string** |  | [optional] 
**ContractStartTime** | Pointer to **int32** |  | [optional] 
**ContractType** | Pointer to **string** | Contract type (Maintenance / Subscription / Premium / Gov AdvCare / Gov TAC / High Sec / AdvCare / Gov Premium) | [optional] 
**CustomerPo** | Pointer to **string** | po number associated with this SKU | [optional] 
**Distributor** | Pointer to **string** | distributor name | [optional] 
**EolTime** | Pointer to **int32** | end of life time | [optional] 
**EosTime** | Pointer to **int32** | end of support time | [optional] 
**InstalledAddress** | Pointer to **string** | address where the device is installed. It is a combination of address , region , country , zip | [optional] 
**Model** | Pointer to **string** | model of the install base inventory | [optional] 
**OrderId** | Pointer to **string** | order ID associated with this SKU | [optional] 
**Reseller** | Pointer to **string** | reseller name | [optional] 
**Serial** | Pointer to **string** | serial Number of the inventory | [optional] 
**ShippedTime** | Pointer to **float32** | Shipped date | [optional] 
**Sku** | Pointer to **string** | serviceable device stock | [optional] 
**Type** | Pointer to [**JseInventoryItemType**](JseInventoryItemType.md) |  | [optional] 
**WarrantyEndTime** | Pointer to **int32** |  | [optional] 
**WarrantyStartTime** | Pointer to **int32** |  | [optional] 
**WarrantyType** | Pointer to **string** |  | [optional] 

## Methods

### NewJseInventoryItem

`func NewJseInventoryItem() *JseInventoryItem`

NewJseInventoryItem instantiates a new JseInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJseInventoryItemWithDefaults

`func NewJseInventoryItemWithDefaults() *JseInventoryItem`

NewJseInventoryItemWithDefaults instantiates a new JseInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractEndTime

`func (o *JseInventoryItem) GetContractEndTime() int32`

GetContractEndTime returns the ContractEndTime field if non-nil, zero value otherwise.

### GetContractEndTimeOk

`func (o *JseInventoryItem) GetContractEndTimeOk() (*int32, bool)`

GetContractEndTimeOk returns a tuple with the ContractEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEndTime

`func (o *JseInventoryItem) SetContractEndTime(v int32)`

SetContractEndTime sets ContractEndTime field to given value.

### HasContractEndTime

`func (o *JseInventoryItem) HasContractEndTime() bool`

HasContractEndTime returns a boolean if a field has been set.

### GetContractId

`func (o *JseInventoryItem) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *JseInventoryItem) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *JseInventoryItem) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *JseInventoryItem) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetContractSku

`func (o *JseInventoryItem) GetContractSku() string`

GetContractSku returns the ContractSku field if non-nil, zero value otherwise.

### GetContractSkuOk

`func (o *JseInventoryItem) GetContractSkuOk() (*string, bool)`

GetContractSkuOk returns a tuple with the ContractSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSku

`func (o *JseInventoryItem) SetContractSku(v string)`

SetContractSku sets ContractSku field to given value.

### HasContractSku

`func (o *JseInventoryItem) HasContractSku() bool`

HasContractSku returns a boolean if a field has been set.

### GetContractStartTime

`func (o *JseInventoryItem) GetContractStartTime() int32`

GetContractStartTime returns the ContractStartTime field if non-nil, zero value otherwise.

### GetContractStartTimeOk

`func (o *JseInventoryItem) GetContractStartTimeOk() (*int32, bool)`

GetContractStartTimeOk returns a tuple with the ContractStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStartTime

`func (o *JseInventoryItem) SetContractStartTime(v int32)`

SetContractStartTime sets ContractStartTime field to given value.

### HasContractStartTime

`func (o *JseInventoryItem) HasContractStartTime() bool`

HasContractStartTime returns a boolean if a field has been set.

### GetContractType

`func (o *JseInventoryItem) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *JseInventoryItem) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *JseInventoryItem) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *JseInventoryItem) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCustomerPo

`func (o *JseInventoryItem) GetCustomerPo() string`

GetCustomerPo returns the CustomerPo field if non-nil, zero value otherwise.

### GetCustomerPoOk

`func (o *JseInventoryItem) GetCustomerPoOk() (*string, bool)`

GetCustomerPoOk returns a tuple with the CustomerPo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPo

`func (o *JseInventoryItem) SetCustomerPo(v string)`

SetCustomerPo sets CustomerPo field to given value.

### HasCustomerPo

`func (o *JseInventoryItem) HasCustomerPo() bool`

HasCustomerPo returns a boolean if a field has been set.

### GetDistributor

`func (o *JseInventoryItem) GetDistributor() string`

GetDistributor returns the Distributor field if non-nil, zero value otherwise.

### GetDistributorOk

`func (o *JseInventoryItem) GetDistributorOk() (*string, bool)`

GetDistributorOk returns a tuple with the Distributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributor

`func (o *JseInventoryItem) SetDistributor(v string)`

SetDistributor sets Distributor field to given value.

### HasDistributor

`func (o *JseInventoryItem) HasDistributor() bool`

HasDistributor returns a boolean if a field has been set.

### GetEolTime

`func (o *JseInventoryItem) GetEolTime() int32`

GetEolTime returns the EolTime field if non-nil, zero value otherwise.

### GetEolTimeOk

`func (o *JseInventoryItem) GetEolTimeOk() (*int32, bool)`

GetEolTimeOk returns a tuple with the EolTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolTime

`func (o *JseInventoryItem) SetEolTime(v int32)`

SetEolTime sets EolTime field to given value.

### HasEolTime

`func (o *JseInventoryItem) HasEolTime() bool`

HasEolTime returns a boolean if a field has been set.

### GetEosTime

`func (o *JseInventoryItem) GetEosTime() int32`

GetEosTime returns the EosTime field if non-nil, zero value otherwise.

### GetEosTimeOk

`func (o *JseInventoryItem) GetEosTimeOk() (*int32, bool)`

GetEosTimeOk returns a tuple with the EosTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEosTime

`func (o *JseInventoryItem) SetEosTime(v int32)`

SetEosTime sets EosTime field to given value.

### HasEosTime

`func (o *JseInventoryItem) HasEosTime() bool`

HasEosTime returns a boolean if a field has been set.

### GetInstalledAddress

`func (o *JseInventoryItem) GetInstalledAddress() string`

GetInstalledAddress returns the InstalledAddress field if non-nil, zero value otherwise.

### GetInstalledAddressOk

`func (o *JseInventoryItem) GetInstalledAddressOk() (*string, bool)`

GetInstalledAddressOk returns a tuple with the InstalledAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAddress

`func (o *JseInventoryItem) SetInstalledAddress(v string)`

SetInstalledAddress sets InstalledAddress field to given value.

### HasInstalledAddress

`func (o *JseInventoryItem) HasInstalledAddress() bool`

HasInstalledAddress returns a boolean if a field has been set.

### GetModel

`func (o *JseInventoryItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JseInventoryItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JseInventoryItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *JseInventoryItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrderId

`func (o *JseInventoryItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *JseInventoryItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *JseInventoryItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *JseInventoryItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetReseller

`func (o *JseInventoryItem) GetReseller() string`

GetReseller returns the Reseller field if non-nil, zero value otherwise.

### GetResellerOk

`func (o *JseInventoryItem) GetResellerOk() (*string, bool)`

GetResellerOk returns a tuple with the Reseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReseller

`func (o *JseInventoryItem) SetReseller(v string)`

SetReseller sets Reseller field to given value.

### HasReseller

`func (o *JseInventoryItem) HasReseller() bool`

HasReseller returns a boolean if a field has been set.

### GetSerial

`func (o *JseInventoryItem) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *JseInventoryItem) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *JseInventoryItem) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *JseInventoryItem) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShippedTime

`func (o *JseInventoryItem) GetShippedTime() float32`

GetShippedTime returns the ShippedTime field if non-nil, zero value otherwise.

### GetShippedTimeOk

`func (o *JseInventoryItem) GetShippedTimeOk() (*float32, bool)`

GetShippedTimeOk returns a tuple with the ShippedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedTime

`func (o *JseInventoryItem) SetShippedTime(v float32)`

SetShippedTime sets ShippedTime field to given value.

### HasShippedTime

`func (o *JseInventoryItem) HasShippedTime() bool`

HasShippedTime returns a boolean if a field has been set.

### GetSku

`func (o *JseInventoryItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *JseInventoryItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *JseInventoryItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *JseInventoryItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetType

`func (o *JseInventoryItem) GetType() JseInventoryItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JseInventoryItem) GetTypeOk() (*JseInventoryItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JseInventoryItem) SetType(v JseInventoryItemType)`

SetType sets Type field to given value.

### HasType

`func (o *JseInventoryItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWarrantyEndTime

`func (o *JseInventoryItem) GetWarrantyEndTime() int32`

GetWarrantyEndTime returns the WarrantyEndTime field if non-nil, zero value otherwise.

### GetWarrantyEndTimeOk

`func (o *JseInventoryItem) GetWarrantyEndTimeOk() (*int32, bool)`

GetWarrantyEndTimeOk returns a tuple with the WarrantyEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEndTime

`func (o *JseInventoryItem) SetWarrantyEndTime(v int32)`

SetWarrantyEndTime sets WarrantyEndTime field to given value.

### HasWarrantyEndTime

`func (o *JseInventoryItem) HasWarrantyEndTime() bool`

HasWarrantyEndTime returns a boolean if a field has been set.

### GetWarrantyStartTime

`func (o *JseInventoryItem) GetWarrantyStartTime() int32`

GetWarrantyStartTime returns the WarrantyStartTime field if non-nil, zero value otherwise.

### GetWarrantyStartTimeOk

`func (o *JseInventoryItem) GetWarrantyStartTimeOk() (*int32, bool)`

GetWarrantyStartTimeOk returns a tuple with the WarrantyStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyStartTime

`func (o *JseInventoryItem) SetWarrantyStartTime(v int32)`

SetWarrantyStartTime sets WarrantyStartTime field to given value.

### HasWarrantyStartTime

`func (o *JseInventoryItem) HasWarrantyStartTime() bool`

HasWarrantyStartTime returns a boolean if a field has been set.

### GetWarrantyType

`func (o *JseInventoryItem) GetWarrantyType() string`

GetWarrantyType returns the WarrantyType field if non-nil, zero value otherwise.

### GetWarrantyTypeOk

`func (o *JseInventoryItem) GetWarrantyTypeOk() (*string, bool)`

GetWarrantyTypeOk returns a tuple with the WarrantyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyType

`func (o *JseInventoryItem) SetWarrantyType(v string)`

SetWarrantyType sets WarrantyType field to given value.

### HasWarrantyType

`func (o *JseInventoryItem) HasWarrantyType() bool`

HasWarrantyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


