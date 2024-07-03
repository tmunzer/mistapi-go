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

// checks if the ZoneStatsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneStatsDetails{}

// ZoneStatsDetails Zone details statistics
type ZoneStatsDetails struct {
	// list of ble assets currently in the zone and when they entered
	Assets []string `json:"assets,omitempty"`
	ClientWaits ZoneStatsDetailsClientWaits `json:"client_waits"`
	// list of clients currently in the zone and when they entered
	Clients []string `json:"clients,omitempty"`
	// id of the zone
	Id string `json:"id"`
	// map_id of the zone
	MapId string `json:"map_id"`
	// name of the zone
	Name string `json:"name"`
	NumClients int32 `json:"num_clients"`
	// sdkclient wait time right now
	NumSdkclients int32 `json:"num_sdkclients"`
	// list of sdkclients currently in the zone and when they entered
	Sdkclients []string `json:"sdkclients,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZoneStatsDetails ZoneStatsDetails

// NewZoneStatsDetails instantiates a new ZoneStatsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneStatsDetails(clientWaits ZoneStatsDetailsClientWaits, id string, mapId string, name string, numClients int32, numSdkclients int32) *ZoneStatsDetails {
	this := ZoneStatsDetails{}
	this.ClientWaits = clientWaits
	this.Id = id
	this.MapId = mapId
	this.Name = name
	this.NumClients = numClients
	this.NumSdkclients = numSdkclients
	return &this
}

// NewZoneStatsDetailsWithDefaults instantiates a new ZoneStatsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneStatsDetailsWithDefaults() *ZoneStatsDetails {
	this := ZoneStatsDetails{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ZoneStatsDetails) GetAssets() []string {
	if o == nil || IsNil(o.Assets) {
		var ret []string
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetAssetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ZoneStatsDetails) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []string and assigns it to the Assets field.
func (o *ZoneStatsDetails) SetAssets(v []string) {
	o.Assets = v
}

// GetClientWaits returns the ClientWaits field value
func (o *ZoneStatsDetails) GetClientWaits() ZoneStatsDetailsClientWaits {
	if o == nil {
		var ret ZoneStatsDetailsClientWaits
		return ret
	}

	return o.ClientWaits
}

// GetClientWaitsOk returns a tuple with the ClientWaits field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetClientWaitsOk() (*ZoneStatsDetailsClientWaits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientWaits, true
}

// SetClientWaits sets field value
func (o *ZoneStatsDetails) SetClientWaits(v ZoneStatsDetailsClientWaits) {
	o.ClientWaits = v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ZoneStatsDetails) GetClients() []string {
	if o == nil || IsNil(o.Clients) {
		var ret []string
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetClientsOk() ([]string, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ZoneStatsDetails) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []string and assigns it to the Clients field.
func (o *ZoneStatsDetails) SetClients(v []string) {
	o.Clients = v
}

// GetId returns the Id field value
func (o *ZoneStatsDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ZoneStatsDetails) SetId(v string) {
	o.Id = v
}

// GetMapId returns the MapId field value
func (o *ZoneStatsDetails) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *ZoneStatsDetails) SetMapId(v string) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *ZoneStatsDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZoneStatsDetails) SetName(v string) {
	o.Name = v
}

// GetNumClients returns the NumClients field value
func (o *ZoneStatsDetails) GetNumClients() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetNumClientsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumClients, true
}

// SetNumClients sets field value
func (o *ZoneStatsDetails) SetNumClients(v int32) {
	o.NumClients = v
}

// GetNumSdkclients returns the NumSdkclients field value
func (o *ZoneStatsDetails) GetNumSdkclients() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumSdkclients
}

// GetNumSdkclientsOk returns a tuple with the NumSdkclients field value
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetNumSdkclientsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumSdkclients, true
}

// SetNumSdkclients sets field value
func (o *ZoneStatsDetails) SetNumSdkclients(v int32) {
	o.NumSdkclients = v
}

// GetSdkclients returns the Sdkclients field value if set, zero value otherwise.
func (o *ZoneStatsDetails) GetSdkclients() []string {
	if o == nil || IsNil(o.Sdkclients) {
		var ret []string
		return ret
	}
	return o.Sdkclients
}

// GetSdkclientsOk returns a tuple with the Sdkclients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStatsDetails) GetSdkclientsOk() ([]string, bool) {
	if o == nil || IsNil(o.Sdkclients) {
		return nil, false
	}
	return o.Sdkclients, true
}

// HasSdkclients returns a boolean if a field has been set.
func (o *ZoneStatsDetails) HasSdkclients() bool {
	if o != nil && !IsNil(o.Sdkclients) {
		return true
	}

	return false
}

// SetSdkclients gets a reference to the given []string and assigns it to the Sdkclients field.
func (o *ZoneStatsDetails) SetSdkclients(v []string) {
	o.Sdkclients = v
}

func (o ZoneStatsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneStatsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	toSerialize["client_waits"] = o.ClientWaits
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	toSerialize["id"] = o.Id
	toSerialize["map_id"] = o.MapId
	toSerialize["name"] = o.Name
	toSerialize["num_clients"] = o.NumClients
	toSerialize["num_sdkclients"] = o.NumSdkclients
	if !IsNil(o.Sdkclients) {
		toSerialize["sdkclients"] = o.Sdkclients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZoneStatsDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_waits",
		"id",
		"map_id",
		"name",
		"num_clients",
		"num_sdkclients",
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

	varZoneStatsDetails := _ZoneStatsDetails{}

	err = json.Unmarshal(data, &varZoneStatsDetails)

	if err != nil {
		return err
	}

	*o = ZoneStatsDetails(varZoneStatsDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "client_waits")
		delete(additionalProperties, "clients")
		delete(additionalProperties, "id")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "num_clients")
		delete(additionalProperties, "num_sdkclients")
		delete(additionalProperties, "sdkclients")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZoneStatsDetails struct {
	value *ZoneStatsDetails
	isSet bool
}

func (v NullableZoneStatsDetails) Get() *ZoneStatsDetails {
	return v.value
}

func (v *NullableZoneStatsDetails) Set(val *ZoneStatsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneStatsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneStatsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneStatsDetails(val *ZoneStatsDetails) *NullableZoneStatsDetails {
	return &NullableZoneStatsDetails{value: val, isSet: true}
}

func (v NullableZoneStatsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneStatsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


