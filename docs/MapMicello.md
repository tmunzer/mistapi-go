# MapMicello

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | the account key that has access to the map | 
**DefaultLevelId** | **int32** | micello floor/level id | 
**MapId** | **string** | micello map id | 
**VendorName** | [**MapMicelloVendorName**](MapMicelloVendorName.md) |  | [default to MAPMICELLOVENDORNAME_MICELLO]

## Methods

### NewMapMicello

`func NewMapMicello(accountKey string, defaultLevelId int32, mapId string, vendorName MapMicelloVendorName, ) *MapMicello`

NewMapMicello instantiates a new MapMicello object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapMicelloWithDefaults

`func NewMapMicelloWithDefaults() *MapMicello`

NewMapMicelloWithDefaults instantiates a new MapMicello object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *MapMicello) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *MapMicello) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *MapMicello) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetDefaultLevelId

`func (o *MapMicello) GetDefaultLevelId() int32`

GetDefaultLevelId returns the DefaultLevelId field if non-nil, zero value otherwise.

### GetDefaultLevelIdOk

`func (o *MapMicello) GetDefaultLevelIdOk() (*int32, bool)`

GetDefaultLevelIdOk returns a tuple with the DefaultLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLevelId

`func (o *MapMicello) SetDefaultLevelId(v int32)`

SetDefaultLevelId sets DefaultLevelId field to given value.


### GetMapId

`func (o *MapMicello) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *MapMicello) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *MapMicello) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetVendorName

`func (o *MapMicello) GetVendorName() MapMicelloVendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *MapMicello) GetVendorNameOk() (*MapMicelloVendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *MapMicello) SetVendorName(v MapMicelloVendorName)`

SetVendorName sets VendorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


