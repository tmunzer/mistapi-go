# OrgSiteSleWifiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAvailability** | Pointer to **float32** |  | [optional] 
**ApHealth** | Pointer to **float32** |  | [optional] 
**Capacity** | Pointer to **float32** |  | [optional] 
**Coverage** | Pointer to **float32** |  | [optional] 
**NumAps** | Pointer to **float32** |  | [optional] 
**NumClients** | Pointer to **float32** |  | [optional] 
**Roaming** | Pointer to **float32** |  | [optional] 
**SiteId** | **string** |  | [readonly] 
**SuccessfulConnect** | Pointer to **float32** |  | [optional] 
**Throughput** | Pointer to **float32** |  | [optional] 
**TimeToConnect** | Pointer to **float32** |  | [optional] 

## Methods

### NewOrgSiteSleWifiResult

`func NewOrgSiteSleWifiResult(siteId string, ) *OrgSiteSleWifiResult`

NewOrgSiteSleWifiResult instantiates a new OrgSiteSleWifiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSiteSleWifiResultWithDefaults

`func NewOrgSiteSleWifiResultWithDefaults() *OrgSiteSleWifiResult`

NewOrgSiteSleWifiResultWithDefaults instantiates a new OrgSiteSleWifiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApAvailability

`func (o *OrgSiteSleWifiResult) GetApAvailability() float32`

GetApAvailability returns the ApAvailability field if non-nil, zero value otherwise.

### GetApAvailabilityOk

`func (o *OrgSiteSleWifiResult) GetApAvailabilityOk() (*float32, bool)`

GetApAvailabilityOk returns a tuple with the ApAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAvailability

`func (o *OrgSiteSleWifiResult) SetApAvailability(v float32)`

SetApAvailability sets ApAvailability field to given value.

### HasApAvailability

`func (o *OrgSiteSleWifiResult) HasApAvailability() bool`

HasApAvailability returns a boolean if a field has been set.

### GetApHealth

`func (o *OrgSiteSleWifiResult) GetApHealth() float32`

GetApHealth returns the ApHealth field if non-nil, zero value otherwise.

### GetApHealthOk

`func (o *OrgSiteSleWifiResult) GetApHealthOk() (*float32, bool)`

GetApHealthOk returns a tuple with the ApHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApHealth

`func (o *OrgSiteSleWifiResult) SetApHealth(v float32)`

SetApHealth sets ApHealth field to given value.

### HasApHealth

`func (o *OrgSiteSleWifiResult) HasApHealth() bool`

HasApHealth returns a boolean if a field has been set.

### GetCapacity

`func (o *OrgSiteSleWifiResult) GetCapacity() float32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *OrgSiteSleWifiResult) GetCapacityOk() (*float32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *OrgSiteSleWifiResult) SetCapacity(v float32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *OrgSiteSleWifiResult) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCoverage

`func (o *OrgSiteSleWifiResult) GetCoverage() float32`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *OrgSiteSleWifiResult) GetCoverageOk() (*float32, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *OrgSiteSleWifiResult) SetCoverage(v float32)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *OrgSiteSleWifiResult) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetNumAps

`func (o *OrgSiteSleWifiResult) GetNumAps() float32`

GetNumAps returns the NumAps field if non-nil, zero value otherwise.

### GetNumApsOk

`func (o *OrgSiteSleWifiResult) GetNumApsOk() (*float32, bool)`

GetNumApsOk returns a tuple with the NumAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAps

`func (o *OrgSiteSleWifiResult) SetNumAps(v float32)`

SetNumAps sets NumAps field to given value.

### HasNumAps

`func (o *OrgSiteSleWifiResult) HasNumAps() bool`

HasNumAps returns a boolean if a field has been set.

### GetNumClients

`func (o *OrgSiteSleWifiResult) GetNumClients() float32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *OrgSiteSleWifiResult) GetNumClientsOk() (*float32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *OrgSiteSleWifiResult) SetNumClients(v float32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *OrgSiteSleWifiResult) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetRoaming

`func (o *OrgSiteSleWifiResult) GetRoaming() float32`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *OrgSiteSleWifiResult) GetRoamingOk() (*float32, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *OrgSiteSleWifiResult) SetRoaming(v float32)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *OrgSiteSleWifiResult) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetSiteId

`func (o *OrgSiteSleWifiResult) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OrgSiteSleWifiResult) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OrgSiteSleWifiResult) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSuccessfulConnect

`func (o *OrgSiteSleWifiResult) GetSuccessfulConnect() float32`

GetSuccessfulConnect returns the SuccessfulConnect field if non-nil, zero value otherwise.

### GetSuccessfulConnectOk

`func (o *OrgSiteSleWifiResult) GetSuccessfulConnectOk() (*float32, bool)`

GetSuccessfulConnectOk returns a tuple with the SuccessfulConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulConnect

`func (o *OrgSiteSleWifiResult) SetSuccessfulConnect(v float32)`

SetSuccessfulConnect sets SuccessfulConnect field to given value.

### HasSuccessfulConnect

`func (o *OrgSiteSleWifiResult) HasSuccessfulConnect() bool`

HasSuccessfulConnect returns a boolean if a field has been set.

### GetThroughput

`func (o *OrgSiteSleWifiResult) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *OrgSiteSleWifiResult) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *OrgSiteSleWifiResult) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *OrgSiteSleWifiResult) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetTimeToConnect

`func (o *OrgSiteSleWifiResult) GetTimeToConnect() float32`

GetTimeToConnect returns the TimeToConnect field if non-nil, zero value otherwise.

### GetTimeToConnectOk

`func (o *OrgSiteSleWifiResult) GetTimeToConnectOk() (*float32, bool)`

GetTimeToConnectOk returns a tuple with the TimeToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToConnect

`func (o *OrgSiteSleWifiResult) SetTimeToConnect(v float32)`

SetTimeToConnect sets TimeToConnect field to given value.

### HasTimeToConnect

`func (o *OrgSiteSleWifiResult) HasTimeToConnect() bool`

HasTimeToConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


