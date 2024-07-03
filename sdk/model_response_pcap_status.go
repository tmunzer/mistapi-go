/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the ResponsePcapStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePcapStatus{}

// ResponsePcapStatus struct for ResponsePcapStatus
type ResponsePcapStatus struct {
	ApMac NullableString `json:"ap_mac,omitempty"`
	// List of target APs to capture packets
	Aps []string `json:"aps,omitempty"`
	ClientMac NullableString `json:"client_mac,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	// List of APs where configuration attempt failed
	Failed []string `json:"failed,omitempty"`
	// List of target Gateways to capture packets if a gateway capture type is specified
	Gateways []string `json:"gateways,omitempty"`
	// unique id for the capture
	Id string `json:"id"`
	IncludesMcast *bool `json:"includes_mcast,omitempty"`
	// max number of packets configured by user
	MaxNumPackets *int32 `json:"max_num_packets,omitempty"`
	MaxPktLen *int32 `json:"max_pkt_len,omitempty"`
	// total number of packets captured by all AP, not applicable for type [client, new_assoc]
	NumPackets *int32 `json:"num_packets,omitempty"`
	// List of target APs successfully configured to capture packets
	Ok []string `json:"ok,omitempty"`
	PcapAps *map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
	// when `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user
	RadiotapTcpdumpExpression *string `json:"radiotap_tcpdump_expression,omitempty"`
	// when `type`==`scan`, scan_tcpdump_expression provided by the user
	ScanTcpdumpExpression *string `json:"scan_tcpdump_expression,omitempty"`
	Ssid NullableString `json:"ssid,omitempty"`
	StartedTime *int32 `json:"started_time,omitempty"`
	// List of target Switches to capture packets if a switch capture type is specified
	Switches []string `json:"switches,omitempty"`
	// tcpdump expression provided by the user (common)
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
	Type PcapType `json:"type"`
	// when `type`==`wired`, wired_tcpdump_expression provided by the user
	WiredTcpdumpExpression *string `json:"wired_tcpdump_expression,omitempty"`
	// when `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user
	WirelessTcpdumpExpression *string `json:"wireless_tcpdump_expression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponsePcapStatus ResponsePcapStatus

// NewResponsePcapStatus instantiates a new ResponsePcapStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePcapStatus(id string, type_ PcapType) *ResponsePcapStatus {
	this := ResponsePcapStatus{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewResponsePcapStatusWithDefaults instantiates a new ResponsePcapStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePcapStatusWithDefaults() *ResponsePcapStatus {
	this := ResponsePcapStatus{}
	return &this
}

// GetApMac returns the ApMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponsePcapStatus) GetApMac() string {
	if o == nil || IsNil(o.ApMac.Get()) {
		var ret string
		return ret
	}
	return *o.ApMac.Get()
}

// GetApMacOk returns a tuple with the ApMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsePcapStatus) GetApMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApMac.Get(), o.ApMac.IsSet()
}

// HasApMac returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasApMac() bool {
	if o != nil && o.ApMac.IsSet() {
		return true
	}

	return false
}

// SetApMac gets a reference to the given NullableString and assigns it to the ApMac field.
func (o *ResponsePcapStatus) SetApMac(v string) {
	o.ApMac.Set(&v)
}
// SetApMacNil sets the value for ApMac to be an explicit nil
func (o *ResponsePcapStatus) SetApMacNil() {
	o.ApMac.Set(nil)
}

// UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
func (o *ResponsePcapStatus) UnsetApMac() {
	o.ApMac.Unset()
}

// GetAps returns the Aps field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetAps() []string {
	if o == nil || IsNil(o.Aps) {
		var ret []string
		return ret
	}
	return o.Aps
}

// GetApsOk returns a tuple with the Aps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetApsOk() ([]string, bool) {
	if o == nil || IsNil(o.Aps) {
		return nil, false
	}
	return o.Aps, true
}

// HasAps returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasAps() bool {
	if o != nil && !IsNil(o.Aps) {
		return true
	}

	return false
}

// SetAps gets a reference to the given []string and assigns it to the Aps field.
func (o *ResponsePcapStatus) SetAps(v []string) {
	o.Aps = v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponsePcapStatus) GetClientMac() string {
	if o == nil || IsNil(o.ClientMac.Get()) {
		var ret string
		return ret
	}
	return *o.ClientMac.Get()
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsePcapStatus) GetClientMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientMac.Get(), o.ClientMac.IsSet()
}

// HasClientMac returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasClientMac() bool {
	if o != nil && o.ClientMac.IsSet() {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given NullableString and assigns it to the ClientMac field.
func (o *ResponsePcapStatus) SetClientMac(v string) {
	o.ClientMac.Set(&v)
}
// SetClientMacNil sets the value for ClientMac to be an explicit nil
func (o *ResponsePcapStatus) SetClientMacNil() {
	o.ClientMac.Set(nil)
}

// UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
func (o *ResponsePcapStatus) UnsetClientMac() {
	o.ClientMac.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *ResponsePcapStatus) SetDuration(v int32) {
	o.Duration = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetFailed() []string {
	if o == nil || IsNil(o.Failed) {
		var ret []string
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetFailedOk() ([]string, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []string and assigns it to the Failed field.
func (o *ResponsePcapStatus) SetFailed(v []string) {
	o.Failed = v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetGateways() []string {
	if o == nil || IsNil(o.Gateways) {
		var ret []string
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetGatewaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []string and assigns it to the Gateways field.
func (o *ResponsePcapStatus) SetGateways(v []string) {
	o.Gateways = v
}

// GetId returns the Id field value
func (o *ResponsePcapStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponsePcapStatus) SetId(v string) {
	o.Id = v
}

// GetIncludesMcast returns the IncludesMcast field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetIncludesMcast() bool {
	if o == nil || IsNil(o.IncludesMcast) {
		var ret bool
		return ret
	}
	return *o.IncludesMcast
}

// GetIncludesMcastOk returns a tuple with the IncludesMcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetIncludesMcastOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesMcast) {
		return nil, false
	}
	return o.IncludesMcast, true
}

// HasIncludesMcast returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasIncludesMcast() bool {
	if o != nil && !IsNil(o.IncludesMcast) {
		return true
	}

	return false
}

// SetIncludesMcast gets a reference to the given bool and assigns it to the IncludesMcast field.
func (o *ResponsePcapStatus) SetIncludesMcast(v bool) {
	o.IncludesMcast = &v
}

// GetMaxNumPackets returns the MaxNumPackets field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetMaxNumPackets() int32 {
	if o == nil || IsNil(o.MaxNumPackets) {
		var ret int32
		return ret
	}
	return *o.MaxNumPackets
}

// GetMaxNumPacketsOk returns a tuple with the MaxNumPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetMaxNumPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumPackets) {
		return nil, false
	}
	return o.MaxNumPackets, true
}

// HasMaxNumPackets returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasMaxNumPackets() bool {
	if o != nil && !IsNil(o.MaxNumPackets) {
		return true
	}

	return false
}

// SetMaxNumPackets gets a reference to the given int32 and assigns it to the MaxNumPackets field.
func (o *ResponsePcapStatus) SetMaxNumPackets(v int32) {
	o.MaxNumPackets = &v
}

// GetMaxPktLen returns the MaxPktLen field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetMaxPktLen() int32 {
	if o == nil || IsNil(o.MaxPktLen) {
		var ret int32
		return ret
	}
	return *o.MaxPktLen
}

// GetMaxPktLenOk returns a tuple with the MaxPktLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetMaxPktLenOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPktLen) {
		return nil, false
	}
	return o.MaxPktLen, true
}

// HasMaxPktLen returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasMaxPktLen() bool {
	if o != nil && !IsNil(o.MaxPktLen) {
		return true
	}

	return false
}

// SetMaxPktLen gets a reference to the given int32 and assigns it to the MaxPktLen field.
func (o *ResponsePcapStatus) SetMaxPktLen(v int32) {
	o.MaxPktLen = &v
}

// GetNumPackets returns the NumPackets field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetNumPackets() int32 {
	if o == nil || IsNil(o.NumPackets) {
		var ret int32
		return ret
	}
	return *o.NumPackets
}

// GetNumPacketsOk returns a tuple with the NumPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetNumPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumPackets) {
		return nil, false
	}
	return o.NumPackets, true
}

// HasNumPackets returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasNumPackets() bool {
	if o != nil && !IsNil(o.NumPackets) {
		return true
	}

	return false
}

// SetNumPackets gets a reference to the given int32 and assigns it to the NumPackets field.
func (o *ResponsePcapStatus) SetNumPackets(v int32) {
	o.NumPackets = &v
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetOk() []string {
	if o == nil || IsNil(o.Ok) {
		var ret []string
		return ret
	}
	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetOkOk() ([]string, bool) {
	if o == nil || IsNil(o.Ok) {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasOk() bool {
	if o != nil && !IsNil(o.Ok) {
		return true
	}

	return false
}

// SetOk gets a reference to the given []string and assigns it to the Ok field.
func (o *ResponsePcapStatus) SetOk(v []string) {
	o.Ok = v
}

// GetPcapAps returns the PcapAps field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetPcapAps() map[string]ResponsePcapAp {
	if o == nil || IsNil(o.PcapAps) {
		var ret map[string]ResponsePcapAp
		return ret
	}
	return *o.PcapAps
}

// GetPcapApsOk returns a tuple with the PcapAps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetPcapApsOk() (*map[string]ResponsePcapAp, bool) {
	if o == nil || IsNil(o.PcapAps) {
		return nil, false
	}
	return o.PcapAps, true
}

// HasPcapAps returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasPcapAps() bool {
	if o != nil && !IsNil(o.PcapAps) {
		return true
	}

	return false
}

// SetPcapAps gets a reference to the given map[string]ResponsePcapAp and assigns it to the PcapAps field.
func (o *ResponsePcapStatus) SetPcapAps(v map[string]ResponsePcapAp) {
	o.PcapAps = &v
}

// GetRadiotapTcpdumpExpression returns the RadiotapTcpdumpExpression field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetRadiotapTcpdumpExpression() string {
	if o == nil || IsNil(o.RadiotapTcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.RadiotapTcpdumpExpression
}

// GetRadiotapTcpdumpExpressionOk returns a tuple with the RadiotapTcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetRadiotapTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.RadiotapTcpdumpExpression) {
		return nil, false
	}
	return o.RadiotapTcpdumpExpression, true
}

// HasRadiotapTcpdumpExpression returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasRadiotapTcpdumpExpression() bool {
	if o != nil && !IsNil(o.RadiotapTcpdumpExpression) {
		return true
	}

	return false
}

// SetRadiotapTcpdumpExpression gets a reference to the given string and assigns it to the RadiotapTcpdumpExpression field.
func (o *ResponsePcapStatus) SetRadiotapTcpdumpExpression(v string) {
	o.RadiotapTcpdumpExpression = &v
}

// GetScanTcpdumpExpression returns the ScanTcpdumpExpression field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetScanTcpdumpExpression() string {
	if o == nil || IsNil(o.ScanTcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.ScanTcpdumpExpression
}

// GetScanTcpdumpExpressionOk returns a tuple with the ScanTcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetScanTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ScanTcpdumpExpression) {
		return nil, false
	}
	return o.ScanTcpdumpExpression, true
}

// HasScanTcpdumpExpression returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasScanTcpdumpExpression() bool {
	if o != nil && !IsNil(o.ScanTcpdumpExpression) {
		return true
	}

	return false
}

// SetScanTcpdumpExpression gets a reference to the given string and assigns it to the ScanTcpdumpExpression field.
func (o *ResponsePcapStatus) SetScanTcpdumpExpression(v string) {
	o.ScanTcpdumpExpression = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponsePcapStatus) GetSsid() string {
	if o == nil || IsNil(o.Ssid.Get()) {
		var ret string
		return ret
	}
	return *o.Ssid.Get()
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsePcapStatus) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ssid.Get(), o.Ssid.IsSet()
}

// HasSsid returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasSsid() bool {
	if o != nil && o.Ssid.IsSet() {
		return true
	}

	return false
}

// SetSsid gets a reference to the given NullableString and assigns it to the Ssid field.
func (o *ResponsePcapStatus) SetSsid(v string) {
	o.Ssid.Set(&v)
}
// SetSsidNil sets the value for Ssid to be an explicit nil
func (o *ResponsePcapStatus) SetSsidNil() {
	o.Ssid.Set(nil)
}

// UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
func (o *ResponsePcapStatus) UnsetSsid() {
	o.Ssid.Unset()
}

// GetStartedTime returns the StartedTime field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetStartedTime() int32 {
	if o == nil || IsNil(o.StartedTime) {
		var ret int32
		return ret
	}
	return *o.StartedTime
}

// GetStartedTimeOk returns a tuple with the StartedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetStartedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedTime) {
		return nil, false
	}
	return o.StartedTime, true
}

// HasStartedTime returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasStartedTime() bool {
	if o != nil && !IsNil(o.StartedTime) {
		return true
	}

	return false
}

// SetStartedTime gets a reference to the given int32 and assigns it to the StartedTime field.
func (o *ResponsePcapStatus) SetStartedTime(v int32) {
	o.StartedTime = &v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetSwitches() []string {
	if o == nil || IsNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetSwitchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *ResponsePcapStatus) SetSwitches(v []string) {
	o.Switches = v
}

// GetTcpdumpExpression returns the TcpdumpExpression field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetTcpdumpExpression() string {
	if o == nil || IsNil(o.TcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.TcpdumpExpression
}

// GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.TcpdumpExpression) {
		return nil, false
	}
	return o.TcpdumpExpression, true
}

// HasTcpdumpExpression returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasTcpdumpExpression() bool {
	if o != nil && !IsNil(o.TcpdumpExpression) {
		return true
	}

	return false
}

// SetTcpdumpExpression gets a reference to the given string and assigns it to the TcpdumpExpression field.
func (o *ResponsePcapStatus) SetTcpdumpExpression(v string) {
	o.TcpdumpExpression = &v
}

// GetType returns the Type field value
func (o *ResponsePcapStatus) GetType() PcapType {
	if o == nil {
		var ret PcapType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetTypeOk() (*PcapType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResponsePcapStatus) SetType(v PcapType) {
	o.Type = v
}

// GetWiredTcpdumpExpression returns the WiredTcpdumpExpression field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetWiredTcpdumpExpression() string {
	if o == nil || IsNil(o.WiredTcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.WiredTcpdumpExpression
}

// GetWiredTcpdumpExpressionOk returns a tuple with the WiredTcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetWiredTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.WiredTcpdumpExpression) {
		return nil, false
	}
	return o.WiredTcpdumpExpression, true
}

// HasWiredTcpdumpExpression returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasWiredTcpdumpExpression() bool {
	if o != nil && !IsNil(o.WiredTcpdumpExpression) {
		return true
	}

	return false
}

// SetWiredTcpdumpExpression gets a reference to the given string and assigns it to the WiredTcpdumpExpression field.
func (o *ResponsePcapStatus) SetWiredTcpdumpExpression(v string) {
	o.WiredTcpdumpExpression = &v
}

// GetWirelessTcpdumpExpression returns the WirelessTcpdumpExpression field value if set, zero value otherwise.
func (o *ResponsePcapStatus) GetWirelessTcpdumpExpression() string {
	if o == nil || IsNil(o.WirelessTcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.WirelessTcpdumpExpression
}

// GetWirelessTcpdumpExpressionOk returns a tuple with the WirelessTcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePcapStatus) GetWirelessTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.WirelessTcpdumpExpression) {
		return nil, false
	}
	return o.WirelessTcpdumpExpression, true
}

// HasWirelessTcpdumpExpression returns a boolean if a field has been set.
func (o *ResponsePcapStatus) HasWirelessTcpdumpExpression() bool {
	if o != nil && !IsNil(o.WirelessTcpdumpExpression) {
		return true
	}

	return false
}

// SetWirelessTcpdumpExpression gets a reference to the given string and assigns it to the WirelessTcpdumpExpression field.
func (o *ResponsePcapStatus) SetWirelessTcpdumpExpression(v string) {
	o.WirelessTcpdumpExpression = &v
}

func (o ResponsePcapStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePcapStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApMac.IsSet() {
		toSerialize["ap_mac"] = o.ApMac.Get()
	}
	if !IsNil(o.Aps) {
		toSerialize["aps"] = o.Aps
	}
	if o.ClientMac.IsSet() {
		toSerialize["client_mac"] = o.ClientMac.Get()
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IncludesMcast) {
		toSerialize["includes_mcast"] = o.IncludesMcast
	}
	if !IsNil(o.MaxNumPackets) {
		toSerialize["max_num_packets"] = o.MaxNumPackets
	}
	if !IsNil(o.MaxPktLen) {
		toSerialize["max_pkt_len"] = o.MaxPktLen
	}
	if !IsNil(o.NumPackets) {
		toSerialize["num_packets"] = o.NumPackets
	}
	if !IsNil(o.Ok) {
		toSerialize["ok"] = o.Ok
	}
	if !IsNil(o.PcapAps) {
		toSerialize["pcap_aps"] = o.PcapAps
	}
	if !IsNil(o.RadiotapTcpdumpExpression) {
		toSerialize["radiotap_tcpdump_expression"] = o.RadiotapTcpdumpExpression
	}
	if !IsNil(o.ScanTcpdumpExpression) {
		toSerialize["scan_tcpdump_expression"] = o.ScanTcpdumpExpression
	}
	if o.Ssid.IsSet() {
		toSerialize["ssid"] = o.Ssid.Get()
	}
	if !IsNil(o.StartedTime) {
		toSerialize["started_time"] = o.StartedTime
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !IsNil(o.TcpdumpExpression) {
		toSerialize["tcpdump_expression"] = o.TcpdumpExpression
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.WiredTcpdumpExpression) {
		toSerialize["wired_tcpdump_expression"] = o.WiredTcpdumpExpression
	}
	if !IsNil(o.WirelessTcpdumpExpression) {
		toSerialize["wireless_tcpdump_expression"] = o.WirelessTcpdumpExpression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponsePcapStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResponsePcapStatus := _ResponsePcapStatus{}

	err = json.Unmarshal(data, &varResponsePcapStatus)

	if err != nil {
		return err
	}

	*o = ResponsePcapStatus(varResponsePcapStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "aps")
		delete(additionalProperties, "client_mac")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "gateways")
		delete(additionalProperties, "id")
		delete(additionalProperties, "includes_mcast")
		delete(additionalProperties, "max_num_packets")
		delete(additionalProperties, "max_pkt_len")
		delete(additionalProperties, "num_packets")
		delete(additionalProperties, "ok")
		delete(additionalProperties, "pcap_aps")
		delete(additionalProperties, "radiotap_tcpdump_expression")
		delete(additionalProperties, "scan_tcpdump_expression")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "started_time")
		delete(additionalProperties, "switches")
		delete(additionalProperties, "tcpdump_expression")
		delete(additionalProperties, "type")
		delete(additionalProperties, "wired_tcpdump_expression")
		delete(additionalProperties, "wireless_tcpdump_expression")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponsePcapStatus struct {
	value *ResponsePcapStatus
	isSet bool
}

func (v NullableResponsePcapStatus) Get() *ResponsePcapStatus {
	return v.value
}

func (v *NullableResponsePcapStatus) Set(val *ResponsePcapStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePcapStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePcapStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePcapStatus(val *ResponsePcapStatus) *NullableResponsePcapStatus {
	return &NullableResponsePcapStatus{value: val, isSet: true}
}

func (v NullableResponsePcapStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePcapStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


