# RemoteSyslogServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]RemoteSyslogContent**](RemoteSyslogContent.md) |  | [optional] 
**ExplicitPriority** | Pointer to **bool** |  | [optional] 
**Facility** | Pointer to [**RemoteSyslogFacility**](RemoteSyslogFacility.md) |  | [optional] [default to REMOTESYSLOGFACILITY_ANY]
**Host** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 514]
**Protocol** | Pointer to [**RemoteSyslogServerProtocol**](RemoteSyslogServerProtocol.md) |  | [optional] [default to REMOTESYSLOGSERVERPROTOCOL_UDP]
**RoutingInstance** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to [**RemoteSyslogSeverity**](RemoteSyslogSeverity.md) |  | [optional] [default to REMOTESYSLOGSEVERITY_ANY]
**SourceAddress** | Pointer to **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] 
**StructuredData** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteSyslogServer

`func NewRemoteSyslogServer() *RemoteSyslogServer`

NewRemoteSyslogServer instantiates a new RemoteSyslogServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerWithDefaults

`func NewRemoteSyslogServerWithDefaults() *RemoteSyslogServer`

NewRemoteSyslogServerWithDefaults instantiates a new RemoteSyslogServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *RemoteSyslogServer) GetContents() []RemoteSyslogContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RemoteSyslogServer) GetContentsOk() (*[]RemoteSyslogContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RemoteSyslogServer) SetContents(v []RemoteSyslogContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *RemoteSyslogServer) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExplicitPriority

`func (o *RemoteSyslogServer) GetExplicitPriority() bool`

GetExplicitPriority returns the ExplicitPriority field if non-nil, zero value otherwise.

### GetExplicitPriorityOk

`func (o *RemoteSyslogServer) GetExplicitPriorityOk() (*bool, bool)`

GetExplicitPriorityOk returns a tuple with the ExplicitPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPriority

`func (o *RemoteSyslogServer) SetExplicitPriority(v bool)`

SetExplicitPriority sets ExplicitPriority field to given value.

### HasExplicitPriority

`func (o *RemoteSyslogServer) HasExplicitPriority() bool`

HasExplicitPriority returns a boolean if a field has been set.

### GetFacility

`func (o *RemoteSyslogServer) GetFacility() RemoteSyslogFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RemoteSyslogServer) GetFacilityOk() (*RemoteSyslogFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RemoteSyslogServer) SetFacility(v RemoteSyslogFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RemoteSyslogServer) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHost

`func (o *RemoteSyslogServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemoteSyslogServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemoteSyslogServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemoteSyslogServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMatch

`func (o *RemoteSyslogServer) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RemoteSyslogServer) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RemoteSyslogServer) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RemoteSyslogServer) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetPort

`func (o *RemoteSyslogServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteSyslogServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteSyslogServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemoteSyslogServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *RemoteSyslogServer) GetProtocol() RemoteSyslogServerProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RemoteSyslogServer) GetProtocolOk() (*RemoteSyslogServerProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RemoteSyslogServer) SetProtocol(v RemoteSyslogServerProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RemoteSyslogServer) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRoutingInstance

`func (o *RemoteSyslogServer) GetRoutingInstance() string`

GetRoutingInstance returns the RoutingInstance field if non-nil, zero value otherwise.

### GetRoutingInstanceOk

`func (o *RemoteSyslogServer) GetRoutingInstanceOk() (*string, bool)`

GetRoutingInstanceOk returns a tuple with the RoutingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingInstance

`func (o *RemoteSyslogServer) SetRoutingInstance(v string)`

SetRoutingInstance sets RoutingInstance field to given value.

### HasRoutingInstance

`func (o *RemoteSyslogServer) HasRoutingInstance() bool`

HasRoutingInstance returns a boolean if a field has been set.

### GetSeverity

`func (o *RemoteSyslogServer) GetSeverity() RemoteSyslogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RemoteSyslogServer) GetSeverityOk() (*RemoteSyslogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RemoteSyslogServer) SetSeverity(v RemoteSyslogSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RemoteSyslogServer) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSourceAddress

`func (o *RemoteSyslogServer) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *RemoteSyslogServer) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *RemoteSyslogServer) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *RemoteSyslogServer) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetStructuredData

`func (o *RemoteSyslogServer) GetStructuredData() bool`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *RemoteSyslogServer) GetStructuredDataOk() (*bool, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *RemoteSyslogServer) SetStructuredData(v bool)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *RemoteSyslogServer) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetTag

`func (o *RemoteSyslogServer) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RemoteSyslogServer) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RemoteSyslogServer) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RemoteSyslogServer) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


