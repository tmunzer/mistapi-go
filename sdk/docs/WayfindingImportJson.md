# WayfindingImportJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | the client id | 
**ClientSecret** | **string** | the client secret | 
**CustomerId** | **int32** | the jibestream customer record id | 
**EndpointUrl** | **string** | the map contents endpoint host | 
**MapId** | **string** | micello map id | 
**Mmpp** | **int32** | millimeter per pixel | 
**Ppm** | **float32** | pixel per meter, same as the map JSON value. | 
**VendorName** | [**MapMicelloVendorName**](MapMicelloVendorName.md) |  | [default to MAPMICELLOVENDORNAME_MICELLO]
**VenueId** | **int32** | the venue or organization id | 
**AccountKey** | **string** | the account key that has access to the map | 
**DefaultLevelId** | **int32** | micello floor/level id | 

## Methods

### NewWayfindingImportJson

`func NewWayfindingImportJson(clientId string, clientSecret string, customerId int32, endpointUrl string, mapId string, mmpp int32, ppm float32, vendorName MapMicelloVendorName, venueId int32, accountKey string, defaultLevelId int32, ) *WayfindingImportJson`

NewWayfindingImportJson instantiates a new WayfindingImportJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayfindingImportJsonWithDefaults

`func NewWayfindingImportJsonWithDefaults() *WayfindingImportJson`

NewWayfindingImportJsonWithDefaults instantiates a new WayfindingImportJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *WayfindingImportJson) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WayfindingImportJson) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WayfindingImportJson) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *WayfindingImportJson) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *WayfindingImportJson) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *WayfindingImportJson) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCustomerId

`func (o *WayfindingImportJson) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WayfindingImportJson) GetCustomerIdOk() (*int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WayfindingImportJson) SetCustomerId(v int32)`

SetCustomerId sets CustomerId field to given value.


### GetEndpointUrl

`func (o *WayfindingImportJson) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *WayfindingImportJson) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *WayfindingImportJson) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.


### GetMapId

`func (o *WayfindingImportJson) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *WayfindingImportJson) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *WayfindingImportJson) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetMmpp

`func (o *WayfindingImportJson) GetMmpp() int32`

GetMmpp returns the Mmpp field if non-nil, zero value otherwise.

### GetMmppOk

`func (o *WayfindingImportJson) GetMmppOk() (*int32, bool)`

GetMmppOk returns a tuple with the Mmpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpp

`func (o *WayfindingImportJson) SetMmpp(v int32)`

SetMmpp sets Mmpp field to given value.


### GetPpm

`func (o *WayfindingImportJson) GetPpm() float32`

GetPpm returns the Ppm field if non-nil, zero value otherwise.

### GetPpmOk

`func (o *WayfindingImportJson) GetPpmOk() (*float32, bool)`

GetPpmOk returns a tuple with the Ppm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpm

`func (o *WayfindingImportJson) SetPpm(v float32)`

SetPpm sets Ppm field to given value.


### GetVendorName

`func (o *WayfindingImportJson) GetVendorName() MapMicelloVendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *WayfindingImportJson) GetVendorNameOk() (*MapMicelloVendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *WayfindingImportJson) SetVendorName(v MapMicelloVendorName)`

SetVendorName sets VendorName field to given value.


### GetVenueId

`func (o *WayfindingImportJson) GetVenueId() int32`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *WayfindingImportJson) GetVenueIdOk() (*int32, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *WayfindingImportJson) SetVenueId(v int32)`

SetVenueId sets VenueId field to given value.


### GetAccountKey

`func (o *WayfindingImportJson) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *WayfindingImportJson) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *WayfindingImportJson) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetDefaultLevelId

`func (o *WayfindingImportJson) GetDefaultLevelId() int32`

GetDefaultLevelId returns the DefaultLevelId field if non-nil, zero value otherwise.

### GetDefaultLevelIdOk

`func (o *WayfindingImportJson) GetDefaultLevelIdOk() (*int32, bool)`

GetDefaultLevelIdOk returns a tuple with the DefaultLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLevelId

`func (o *WayfindingImportJson) SetDefaultLevelId(v int32)`

SetDefaultLevelId sets DefaultLevelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


