# CallTroubleshootData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApNumClients** | Pointer to **float32** |  | [optional] 
**ApRtt** | Pointer to **float32** |  | [optional] 
**ClientCpu** | Pointer to **float32** |  | [optional] 
**ClientNStreams** | Pointer to **float32** |  | [optional] 
**ClientRadioBand** | Pointer to **float32** |  | [optional] 
**ClientRssi** | Pointer to **float32** |  | [optional] 
**ClientRxBytes** | Pointer to **float32** |  | [optional] 
**ClientRxRates** | Pointer to **float32** |  | [optional] 
**ClientRxRetries** | Pointer to **float32** |  | [optional] 
**ClientTxBytes** | Pointer to **float32** |  | [optional] 
**ClientTxRates** | Pointer to **float32** |  | [optional] 
**ClientTxRetries** | Pointer to **float32** |  | [optional] 
**ClientVpnDistaince** | Pointer to **float32** |  | [optional] 
**ClientWifiVersion** | Pointer to **float32** |  | [optional] 
**Expected** | Pointer to **float32** |  | [optional] 
**RadioBandwidth** | Pointer to **float32** |  | [optional] 
**RadioChannel** | Pointer to **float32** |  | [optional] 
**RadioTxPower** | Pointer to **float32** |  | [optional] 
**RadioUtil** | Pointer to **float32** |  | [optional] 
**RadioUtilInterference** | Pointer to **float32** |  | [optional] 
**SiteNumClients** | Pointer to **float32** |  | [optional] 
**WanAvgDownloadMbps** | Pointer to **float32** |  | [optional] 
**WanAvgUploadMbps** | Pointer to **float32** |  | [optional] 
**WanJitter** | Pointer to **float32** |  | [optional] 
**WanMaxDownloadMbps** | Pointer to **float32** |  | [optional] 
**WanMaxUploadMbps** | Pointer to **float32** |  | [optional] 
**WanRtt** | Pointer to **float32** |  | [optional] 

## Methods

### NewCallTroubleshootData

`func NewCallTroubleshootData() *CallTroubleshootData`

NewCallTroubleshootData instantiates a new CallTroubleshootData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTroubleshootDataWithDefaults

`func NewCallTroubleshootDataWithDefaults() *CallTroubleshootData`

NewCallTroubleshootDataWithDefaults instantiates a new CallTroubleshootData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApNumClients

`func (o *CallTroubleshootData) GetApNumClients() float32`

GetApNumClients returns the ApNumClients field if non-nil, zero value otherwise.

### GetApNumClientsOk

`func (o *CallTroubleshootData) GetApNumClientsOk() (*float32, bool)`

GetApNumClientsOk returns a tuple with the ApNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNumClients

`func (o *CallTroubleshootData) SetApNumClients(v float32)`

SetApNumClients sets ApNumClients field to given value.

### HasApNumClients

`func (o *CallTroubleshootData) HasApNumClients() bool`

HasApNumClients returns a boolean if a field has been set.

### GetApRtt

`func (o *CallTroubleshootData) GetApRtt() float32`

GetApRtt returns the ApRtt field if non-nil, zero value otherwise.

### GetApRttOk

`func (o *CallTroubleshootData) GetApRttOk() (*float32, bool)`

GetApRttOk returns a tuple with the ApRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRtt

`func (o *CallTroubleshootData) SetApRtt(v float32)`

SetApRtt sets ApRtt field to given value.

### HasApRtt

`func (o *CallTroubleshootData) HasApRtt() bool`

HasApRtt returns a boolean if a field has been set.

### GetClientCpu

`func (o *CallTroubleshootData) GetClientCpu() float32`

GetClientCpu returns the ClientCpu field if non-nil, zero value otherwise.

### GetClientCpuOk

`func (o *CallTroubleshootData) GetClientCpuOk() (*float32, bool)`

GetClientCpuOk returns a tuple with the ClientCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCpu

`func (o *CallTroubleshootData) SetClientCpu(v float32)`

SetClientCpu sets ClientCpu field to given value.

### HasClientCpu

`func (o *CallTroubleshootData) HasClientCpu() bool`

HasClientCpu returns a boolean if a field has been set.

### GetClientNStreams

`func (o *CallTroubleshootData) GetClientNStreams() float32`

GetClientNStreams returns the ClientNStreams field if non-nil, zero value otherwise.

### GetClientNStreamsOk

`func (o *CallTroubleshootData) GetClientNStreamsOk() (*float32, bool)`

GetClientNStreamsOk returns a tuple with the ClientNStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNStreams

`func (o *CallTroubleshootData) SetClientNStreams(v float32)`

SetClientNStreams sets ClientNStreams field to given value.

### HasClientNStreams

`func (o *CallTroubleshootData) HasClientNStreams() bool`

HasClientNStreams returns a boolean if a field has been set.

### GetClientRadioBand

`func (o *CallTroubleshootData) GetClientRadioBand() float32`

GetClientRadioBand returns the ClientRadioBand field if non-nil, zero value otherwise.

### GetClientRadioBandOk

`func (o *CallTroubleshootData) GetClientRadioBandOk() (*float32, bool)`

GetClientRadioBandOk returns a tuple with the ClientRadioBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRadioBand

`func (o *CallTroubleshootData) SetClientRadioBand(v float32)`

SetClientRadioBand sets ClientRadioBand field to given value.

### HasClientRadioBand

`func (o *CallTroubleshootData) HasClientRadioBand() bool`

HasClientRadioBand returns a boolean if a field has been set.

### GetClientRssi

`func (o *CallTroubleshootData) GetClientRssi() float32`

GetClientRssi returns the ClientRssi field if non-nil, zero value otherwise.

### GetClientRssiOk

`func (o *CallTroubleshootData) GetClientRssiOk() (*float32, bool)`

GetClientRssiOk returns a tuple with the ClientRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRssi

`func (o *CallTroubleshootData) SetClientRssi(v float32)`

SetClientRssi sets ClientRssi field to given value.

### HasClientRssi

`func (o *CallTroubleshootData) HasClientRssi() bool`

HasClientRssi returns a boolean if a field has been set.

### GetClientRxBytes

`func (o *CallTroubleshootData) GetClientRxBytes() float32`

GetClientRxBytes returns the ClientRxBytes field if non-nil, zero value otherwise.

### GetClientRxBytesOk

`func (o *CallTroubleshootData) GetClientRxBytesOk() (*float32, bool)`

GetClientRxBytesOk returns a tuple with the ClientRxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRxBytes

`func (o *CallTroubleshootData) SetClientRxBytes(v float32)`

SetClientRxBytes sets ClientRxBytes field to given value.

### HasClientRxBytes

`func (o *CallTroubleshootData) HasClientRxBytes() bool`

HasClientRxBytes returns a boolean if a field has been set.

### GetClientRxRates

`func (o *CallTroubleshootData) GetClientRxRates() float32`

GetClientRxRates returns the ClientRxRates field if non-nil, zero value otherwise.

### GetClientRxRatesOk

`func (o *CallTroubleshootData) GetClientRxRatesOk() (*float32, bool)`

GetClientRxRatesOk returns a tuple with the ClientRxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRxRates

`func (o *CallTroubleshootData) SetClientRxRates(v float32)`

SetClientRxRates sets ClientRxRates field to given value.

### HasClientRxRates

`func (o *CallTroubleshootData) HasClientRxRates() bool`

HasClientRxRates returns a boolean if a field has been set.

### GetClientRxRetries

`func (o *CallTroubleshootData) GetClientRxRetries() float32`

GetClientRxRetries returns the ClientRxRetries field if non-nil, zero value otherwise.

### GetClientRxRetriesOk

`func (o *CallTroubleshootData) GetClientRxRetriesOk() (*float32, bool)`

GetClientRxRetriesOk returns a tuple with the ClientRxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRxRetries

`func (o *CallTroubleshootData) SetClientRxRetries(v float32)`

SetClientRxRetries sets ClientRxRetries field to given value.

### HasClientRxRetries

`func (o *CallTroubleshootData) HasClientRxRetries() bool`

HasClientRxRetries returns a boolean if a field has been set.

### GetClientTxBytes

`func (o *CallTroubleshootData) GetClientTxBytes() float32`

GetClientTxBytes returns the ClientTxBytes field if non-nil, zero value otherwise.

### GetClientTxBytesOk

`func (o *CallTroubleshootData) GetClientTxBytesOk() (*float32, bool)`

GetClientTxBytesOk returns a tuple with the ClientTxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTxBytes

`func (o *CallTroubleshootData) SetClientTxBytes(v float32)`

SetClientTxBytes sets ClientTxBytes field to given value.

### HasClientTxBytes

`func (o *CallTroubleshootData) HasClientTxBytes() bool`

HasClientTxBytes returns a boolean if a field has been set.

### GetClientTxRates

`func (o *CallTroubleshootData) GetClientTxRates() float32`

GetClientTxRates returns the ClientTxRates field if non-nil, zero value otherwise.

### GetClientTxRatesOk

`func (o *CallTroubleshootData) GetClientTxRatesOk() (*float32, bool)`

GetClientTxRatesOk returns a tuple with the ClientTxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTxRates

`func (o *CallTroubleshootData) SetClientTxRates(v float32)`

SetClientTxRates sets ClientTxRates field to given value.

### HasClientTxRates

`func (o *CallTroubleshootData) HasClientTxRates() bool`

HasClientTxRates returns a boolean if a field has been set.

### GetClientTxRetries

`func (o *CallTroubleshootData) GetClientTxRetries() float32`

GetClientTxRetries returns the ClientTxRetries field if non-nil, zero value otherwise.

### GetClientTxRetriesOk

`func (o *CallTroubleshootData) GetClientTxRetriesOk() (*float32, bool)`

GetClientTxRetriesOk returns a tuple with the ClientTxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTxRetries

`func (o *CallTroubleshootData) SetClientTxRetries(v float32)`

SetClientTxRetries sets ClientTxRetries field to given value.

### HasClientTxRetries

`func (o *CallTroubleshootData) HasClientTxRetries() bool`

HasClientTxRetries returns a boolean if a field has been set.

### GetClientVpnDistaince

`func (o *CallTroubleshootData) GetClientVpnDistaince() float32`

GetClientVpnDistaince returns the ClientVpnDistaince field if non-nil, zero value otherwise.

### GetClientVpnDistainceOk

`func (o *CallTroubleshootData) GetClientVpnDistainceOk() (*float32, bool)`

GetClientVpnDistainceOk returns a tuple with the ClientVpnDistaince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnDistaince

`func (o *CallTroubleshootData) SetClientVpnDistaince(v float32)`

SetClientVpnDistaince sets ClientVpnDistaince field to given value.

### HasClientVpnDistaince

`func (o *CallTroubleshootData) HasClientVpnDistaince() bool`

HasClientVpnDistaince returns a boolean if a field has been set.

### GetClientWifiVersion

`func (o *CallTroubleshootData) GetClientWifiVersion() float32`

GetClientWifiVersion returns the ClientWifiVersion field if non-nil, zero value otherwise.

### GetClientWifiVersionOk

`func (o *CallTroubleshootData) GetClientWifiVersionOk() (*float32, bool)`

GetClientWifiVersionOk returns a tuple with the ClientWifiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWifiVersion

`func (o *CallTroubleshootData) SetClientWifiVersion(v float32)`

SetClientWifiVersion sets ClientWifiVersion field to given value.

### HasClientWifiVersion

`func (o *CallTroubleshootData) HasClientWifiVersion() bool`

HasClientWifiVersion returns a boolean if a field has been set.

### GetExpected

`func (o *CallTroubleshootData) GetExpected() float32`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *CallTroubleshootData) GetExpectedOk() (*float32, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *CallTroubleshootData) SetExpected(v float32)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *CallTroubleshootData) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetRadioBandwidth

`func (o *CallTroubleshootData) GetRadioBandwidth() float32`

GetRadioBandwidth returns the RadioBandwidth field if non-nil, zero value otherwise.

### GetRadioBandwidthOk

`func (o *CallTroubleshootData) GetRadioBandwidthOk() (*float32, bool)`

GetRadioBandwidthOk returns a tuple with the RadioBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioBandwidth

`func (o *CallTroubleshootData) SetRadioBandwidth(v float32)`

SetRadioBandwidth sets RadioBandwidth field to given value.

### HasRadioBandwidth

`func (o *CallTroubleshootData) HasRadioBandwidth() bool`

HasRadioBandwidth returns a boolean if a field has been set.

### GetRadioChannel

`func (o *CallTroubleshootData) GetRadioChannel() float32`

GetRadioChannel returns the RadioChannel field if non-nil, zero value otherwise.

### GetRadioChannelOk

`func (o *CallTroubleshootData) GetRadioChannelOk() (*float32, bool)`

GetRadioChannelOk returns a tuple with the RadioChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioChannel

`func (o *CallTroubleshootData) SetRadioChannel(v float32)`

SetRadioChannel sets RadioChannel field to given value.

### HasRadioChannel

`func (o *CallTroubleshootData) HasRadioChannel() bool`

HasRadioChannel returns a boolean if a field has been set.

### GetRadioTxPower

`func (o *CallTroubleshootData) GetRadioTxPower() float32`

GetRadioTxPower returns the RadioTxPower field if non-nil, zero value otherwise.

### GetRadioTxPowerOk

`func (o *CallTroubleshootData) GetRadioTxPowerOk() (*float32, bool)`

GetRadioTxPowerOk returns a tuple with the RadioTxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioTxPower

`func (o *CallTroubleshootData) SetRadioTxPower(v float32)`

SetRadioTxPower sets RadioTxPower field to given value.

### HasRadioTxPower

`func (o *CallTroubleshootData) HasRadioTxPower() bool`

HasRadioTxPower returns a boolean if a field has been set.

### GetRadioUtil

`func (o *CallTroubleshootData) GetRadioUtil() float32`

GetRadioUtil returns the RadioUtil field if non-nil, zero value otherwise.

### GetRadioUtilOk

`func (o *CallTroubleshootData) GetRadioUtilOk() (*float32, bool)`

GetRadioUtilOk returns a tuple with the RadioUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioUtil

`func (o *CallTroubleshootData) SetRadioUtil(v float32)`

SetRadioUtil sets RadioUtil field to given value.

### HasRadioUtil

`func (o *CallTroubleshootData) HasRadioUtil() bool`

HasRadioUtil returns a boolean if a field has been set.

### GetRadioUtilInterference

`func (o *CallTroubleshootData) GetRadioUtilInterference() float32`

GetRadioUtilInterference returns the RadioUtilInterference field if non-nil, zero value otherwise.

### GetRadioUtilInterferenceOk

`func (o *CallTroubleshootData) GetRadioUtilInterferenceOk() (*float32, bool)`

GetRadioUtilInterferenceOk returns a tuple with the RadioUtilInterference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioUtilInterference

`func (o *CallTroubleshootData) SetRadioUtilInterference(v float32)`

SetRadioUtilInterference sets RadioUtilInterference field to given value.

### HasRadioUtilInterference

`func (o *CallTroubleshootData) HasRadioUtilInterference() bool`

HasRadioUtilInterference returns a boolean if a field has been set.

### GetSiteNumClients

`func (o *CallTroubleshootData) GetSiteNumClients() float32`

GetSiteNumClients returns the SiteNumClients field if non-nil, zero value otherwise.

### GetSiteNumClientsOk

`func (o *CallTroubleshootData) GetSiteNumClientsOk() (*float32, bool)`

GetSiteNumClientsOk returns a tuple with the SiteNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNumClients

`func (o *CallTroubleshootData) SetSiteNumClients(v float32)`

SetSiteNumClients sets SiteNumClients field to given value.

### HasSiteNumClients

`func (o *CallTroubleshootData) HasSiteNumClients() bool`

HasSiteNumClients returns a boolean if a field has been set.

### GetWanAvgDownloadMbps

`func (o *CallTroubleshootData) GetWanAvgDownloadMbps() float32`

GetWanAvgDownloadMbps returns the WanAvgDownloadMbps field if non-nil, zero value otherwise.

### GetWanAvgDownloadMbpsOk

`func (o *CallTroubleshootData) GetWanAvgDownloadMbpsOk() (*float32, bool)`

GetWanAvgDownloadMbpsOk returns a tuple with the WanAvgDownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAvgDownloadMbps

`func (o *CallTroubleshootData) SetWanAvgDownloadMbps(v float32)`

SetWanAvgDownloadMbps sets WanAvgDownloadMbps field to given value.

### HasWanAvgDownloadMbps

`func (o *CallTroubleshootData) HasWanAvgDownloadMbps() bool`

HasWanAvgDownloadMbps returns a boolean if a field has been set.

### GetWanAvgUploadMbps

`func (o *CallTroubleshootData) GetWanAvgUploadMbps() float32`

GetWanAvgUploadMbps returns the WanAvgUploadMbps field if non-nil, zero value otherwise.

### GetWanAvgUploadMbpsOk

`func (o *CallTroubleshootData) GetWanAvgUploadMbpsOk() (*float32, bool)`

GetWanAvgUploadMbpsOk returns a tuple with the WanAvgUploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAvgUploadMbps

`func (o *CallTroubleshootData) SetWanAvgUploadMbps(v float32)`

SetWanAvgUploadMbps sets WanAvgUploadMbps field to given value.

### HasWanAvgUploadMbps

`func (o *CallTroubleshootData) HasWanAvgUploadMbps() bool`

HasWanAvgUploadMbps returns a boolean if a field has been set.

### GetWanJitter

`func (o *CallTroubleshootData) GetWanJitter() float32`

GetWanJitter returns the WanJitter field if non-nil, zero value otherwise.

### GetWanJitterOk

`func (o *CallTroubleshootData) GetWanJitterOk() (*float32, bool)`

GetWanJitterOk returns a tuple with the WanJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanJitter

`func (o *CallTroubleshootData) SetWanJitter(v float32)`

SetWanJitter sets WanJitter field to given value.

### HasWanJitter

`func (o *CallTroubleshootData) HasWanJitter() bool`

HasWanJitter returns a boolean if a field has been set.

### GetWanMaxDownloadMbps

`func (o *CallTroubleshootData) GetWanMaxDownloadMbps() float32`

GetWanMaxDownloadMbps returns the WanMaxDownloadMbps field if non-nil, zero value otherwise.

### GetWanMaxDownloadMbpsOk

`func (o *CallTroubleshootData) GetWanMaxDownloadMbpsOk() (*float32, bool)`

GetWanMaxDownloadMbpsOk returns a tuple with the WanMaxDownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMaxDownloadMbps

`func (o *CallTroubleshootData) SetWanMaxDownloadMbps(v float32)`

SetWanMaxDownloadMbps sets WanMaxDownloadMbps field to given value.

### HasWanMaxDownloadMbps

`func (o *CallTroubleshootData) HasWanMaxDownloadMbps() bool`

HasWanMaxDownloadMbps returns a boolean if a field has been set.

### GetWanMaxUploadMbps

`func (o *CallTroubleshootData) GetWanMaxUploadMbps() float32`

GetWanMaxUploadMbps returns the WanMaxUploadMbps field if non-nil, zero value otherwise.

### GetWanMaxUploadMbpsOk

`func (o *CallTroubleshootData) GetWanMaxUploadMbpsOk() (*float32, bool)`

GetWanMaxUploadMbpsOk returns a tuple with the WanMaxUploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMaxUploadMbps

`func (o *CallTroubleshootData) SetWanMaxUploadMbps(v float32)`

SetWanMaxUploadMbps sets WanMaxUploadMbps field to given value.

### HasWanMaxUploadMbps

`func (o *CallTroubleshootData) HasWanMaxUploadMbps() bool`

HasWanMaxUploadMbps returns a boolean if a field has been set.

### GetWanRtt

`func (o *CallTroubleshootData) GetWanRtt() float32`

GetWanRtt returns the WanRtt field if non-nil, zero value otherwise.

### GetWanRttOk

`func (o *CallTroubleshootData) GetWanRttOk() (*float32, bool)`

GetWanRttOk returns a tuple with the WanRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanRtt

`func (o *CallTroubleshootData) SetWanRtt(v float32)`

SetWanRtt sets WanRtt field to given value.

### HasWanRtt

`func (o *CallTroubleshootData) HasWanRtt() bool`

HasWanRtt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


