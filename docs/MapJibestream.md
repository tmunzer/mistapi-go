# MapJibestream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | the client id | 
**ClientSecret** | **string** | the client secret | 
**CustomerId** | **int32** | the jibestream customer record id | 
**EndpointUrl** | **string** | the map contents endpoint host | 
**MapId** | **string** | the jibestream map id | 
**Mmpp** | **int32** | millimeter per pixel | 
**Ppm** | **float32** | pixel per meter, same as the map JSON value. | 
**VendorName** | [**MapJibestreamVendorName**](MapJibestreamVendorName.md) |  | [default to MAPJIBESTREAMVENDORNAME_JIBESTREAM]
**VenueId** | **int32** | the venue or organization id | 

## Methods

### NewMapJibestream

`func NewMapJibestream(clientId string, clientSecret string, customerId int32, endpointUrl string, mapId string, mmpp int32, ppm float32, vendorName MapJibestreamVendorName, venueId int32, ) *MapJibestream`

NewMapJibestream instantiates a new MapJibestream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapJibestreamWithDefaults

`func NewMapJibestreamWithDefaults() *MapJibestream`

NewMapJibestreamWithDefaults instantiates a new MapJibestream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MapJibestream) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MapJibestream) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MapJibestream) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *MapJibestream) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MapJibestream) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MapJibestream) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCustomerId

`func (o *MapJibestream) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MapJibestream) GetCustomerIdOk() (*int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MapJibestream) SetCustomerId(v int32)`

SetCustomerId sets CustomerId field to given value.


### GetEndpointUrl

`func (o *MapJibestream) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *MapJibestream) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *MapJibestream) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.


### GetMapId

`func (o *MapJibestream) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *MapJibestream) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *MapJibestream) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetMmpp

`func (o *MapJibestream) GetMmpp() int32`

GetMmpp returns the Mmpp field if non-nil, zero value otherwise.

### GetMmppOk

`func (o *MapJibestream) GetMmppOk() (*int32, bool)`

GetMmppOk returns a tuple with the Mmpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpp

`func (o *MapJibestream) SetMmpp(v int32)`

SetMmpp sets Mmpp field to given value.


### GetPpm

`func (o *MapJibestream) GetPpm() float32`

GetPpm returns the Ppm field if non-nil, zero value otherwise.

### GetPpmOk

`func (o *MapJibestream) GetPpmOk() (*float32, bool)`

GetPpmOk returns a tuple with the Ppm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpm

`func (o *MapJibestream) SetPpm(v float32)`

SetPpm sets Ppm field to given value.


### GetVendorName

`func (o *MapJibestream) GetVendorName() MapJibestreamVendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *MapJibestream) GetVendorNameOk() (*MapJibestreamVendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *MapJibestream) SetVendorName(v MapJibestreamVendorName)`

SetVendorName sets VendorName field to given value.


### GetVenueId

`func (o *MapJibestream) GetVenueId() int32`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *MapJibestream) GetVenueIdOk() (*int32, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *MapJibestream) SetVenueId(v int32)`

SetVenueId sets VenueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


