# RemoteSyslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to [**RemoteSyslogArchive**](RemoteSyslogArchive.md) |  | [optional] 
**Console** | Pointer to [**RemoteSyslogConsole**](RemoteSyslogConsole.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Files** | Pointer to [**[]RemoteSyslogFileConfig**](RemoteSyslogFileConfig.md) |  | [optional] 
**Network** | Pointer to **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] 
**SendToAllServers** | Pointer to **bool** |  | [optional] [default to false]
**Servers** | Pointer to [**[]RemoteSyslogServer**](RemoteSyslogServer.md) |  | [optional] 
**TimeFormat** | Pointer to [**RemoteSyslogTimeFormat**](RemoteSyslogTimeFormat.md) |  | [optional] 
**Users** | Pointer to [**[]RemoteSyslogUser**](RemoteSyslogUser.md) |  | [optional] 

## Methods

### NewRemoteSyslog

`func NewRemoteSyslog() *RemoteSyslog`

NewRemoteSyslog instantiates a new RemoteSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogWithDefaults

`func NewRemoteSyslogWithDefaults() *RemoteSyslog`

NewRemoteSyslogWithDefaults instantiates a new RemoteSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *RemoteSyslog) GetArchive() RemoteSyslogArchive`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RemoteSyslog) GetArchiveOk() (*RemoteSyslogArchive, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RemoteSyslog) SetArchive(v RemoteSyslogArchive)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RemoteSyslog) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetConsole

`func (o *RemoteSyslog) GetConsole() RemoteSyslogConsole`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *RemoteSyslog) GetConsoleOk() (*RemoteSyslogConsole, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *RemoteSyslog) SetConsole(v RemoteSyslogConsole)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *RemoteSyslog) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetEnabled

`func (o *RemoteSyslog) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RemoteSyslog) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RemoteSyslog) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RemoteSyslog) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFiles

`func (o *RemoteSyslog) GetFiles() []RemoteSyslogFileConfig`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RemoteSyslog) GetFilesOk() (*[]RemoteSyslogFileConfig, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RemoteSyslog) SetFiles(v []RemoteSyslogFileConfig)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RemoteSyslog) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetNetwork

`func (o *RemoteSyslog) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RemoteSyslog) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RemoteSyslog) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RemoteSyslog) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSendToAllServers

`func (o *RemoteSyslog) GetSendToAllServers() bool`

GetSendToAllServers returns the SendToAllServers field if non-nil, zero value otherwise.

### GetSendToAllServersOk

`func (o *RemoteSyslog) GetSendToAllServersOk() (*bool, bool)`

GetSendToAllServersOk returns a tuple with the SendToAllServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToAllServers

`func (o *RemoteSyslog) SetSendToAllServers(v bool)`

SetSendToAllServers sets SendToAllServers field to given value.

### HasSendToAllServers

`func (o *RemoteSyslog) HasSendToAllServers() bool`

HasSendToAllServers returns a boolean if a field has been set.

### GetServers

`func (o *RemoteSyslog) GetServers() []RemoteSyslogServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RemoteSyslog) GetServersOk() (*[]RemoteSyslogServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RemoteSyslog) SetServers(v []RemoteSyslogServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RemoteSyslog) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTimeFormat

`func (o *RemoteSyslog) GetTimeFormat() RemoteSyslogTimeFormat`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *RemoteSyslog) GetTimeFormatOk() (*RemoteSyslogTimeFormat, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *RemoteSyslog) SetTimeFormat(v RemoteSyslogTimeFormat)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *RemoteSyslog) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetUsers

`func (o *RemoteSyslog) GetUsers() []RemoteSyslogUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RemoteSyslog) GetUsersOk() (*[]RemoteSyslogUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RemoteSyslog) SetUsers(v []RemoteSyslogUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RemoteSyslog) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


