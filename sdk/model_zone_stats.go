/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the ZoneStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneStats{}

// ZoneStats Zone statistics
type ZoneStats struct {
	AssetsWaits *ZoneStatsAssetsWaits `json:"assets_waits,omitempty"`
	ClientsWaits *ZoneStatsClientsWaits `json:"clients_waits,omitempty"`
	CreatedTime *int32 `json:"created_time,omitempty"`
	// id of the zone
	Id string `json:"id"`
	// map_id of the zone
	MapId string `json:"map_id"`
	ModifiedTime *int32 `json:"modified_time,omitempty"`
	// name of the zone
	Name string `json:"name"`
	// number of assets
	NumAssets *int32 `json:"num_assets,omitempty"`
	// number of wifi clients (unconnected + connected)
	NumClients *int32 `json:"num_clients,omitempty"`
	// number of sdk clients
	NumSdkclients *int32 `json:"num_sdkclients,omitempty"`
	OccupancyLimit *int32 `json:"occupancy_limit,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SdkclientsWaits *ZoneStatsSdkclientsWaits `json:"sdkclients_waits,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area
	Vertices []ZoneVertex `json:"vertices,omitempty"`
	VerticesM []ZoneVertexM `json:"vertices_m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZoneStats ZoneStats

// NewZoneStats instantiates a new ZoneStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneStats(id string, mapId string, name string) *ZoneStats {
	this := ZoneStats{}
	this.Id = id
	this.MapId = mapId
	this.Name = name
	return &this
}

// NewZoneStatsWithDefaults instantiates a new ZoneStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneStatsWithDefaults() *ZoneStats {
	this := ZoneStats{}
	return &this
}

// GetAssetsWaits returns the AssetsWaits field value if set, zero value otherwise.
func (o *ZoneStats) GetAssetsWaits() ZoneStatsAssetsWaits {
	if o == nil || IsNil(o.AssetsWaits) {
		var ret ZoneStatsAssetsWaits
		return ret
	}
	return *o.AssetsWaits
}

// GetAssetsWaitsOk returns a tuple with the AssetsWaits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetAssetsWaitsOk() (*ZoneStatsAssetsWaits, bool) {
	if o == nil || IsNil(o.AssetsWaits) {
		return nil, false
	}
	return o.AssetsWaits, true
}

// HasAssetsWaits returns a boolean if a field has been set.
func (o *ZoneStats) HasAssetsWaits() bool {
	if o != nil && !IsNil(o.AssetsWaits) {
		return true
	}

	return false
}

// SetAssetsWaits gets a reference to the given ZoneStatsAssetsWaits and assigns it to the AssetsWaits field.
func (o *ZoneStats) SetAssetsWaits(v ZoneStatsAssetsWaits) {
	o.AssetsWaits = &v
}

// GetClientsWaits returns the ClientsWaits field value if set, zero value otherwise.
func (o *ZoneStats) GetClientsWaits() ZoneStatsClientsWaits {
	if o == nil || IsNil(o.ClientsWaits) {
		var ret ZoneStatsClientsWaits
		return ret
	}
	return *o.ClientsWaits
}

// GetClientsWaitsOk returns a tuple with the ClientsWaits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetClientsWaitsOk() (*ZoneStatsClientsWaits, bool) {
	if o == nil || IsNil(o.ClientsWaits) {
		return nil, false
	}
	return o.ClientsWaits, true
}

// HasClientsWaits returns a boolean if a field has been set.
func (o *ZoneStats) HasClientsWaits() bool {
	if o != nil && !IsNil(o.ClientsWaits) {
		return true
	}

	return false
}

// SetClientsWaits gets a reference to the given ZoneStatsClientsWaits and assigns it to the ClientsWaits field.
func (o *ZoneStats) SetClientsWaits(v ZoneStatsClientsWaits) {
	o.ClientsWaits = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ZoneStats) GetCreatedTime() int32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret int32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetCreatedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ZoneStats) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int32 and assigns it to the CreatedTime field.
func (o *ZoneStats) SetCreatedTime(v int32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value
func (o *ZoneStats) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ZoneStats) SetId(v string) {
	o.Id = v
}

// GetMapId returns the MapId field value
func (o *ZoneStats) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *ZoneStats) SetMapId(v string) {
	o.MapId = v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *ZoneStats) GetModifiedTime() int32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret int32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetModifiedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *ZoneStats) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given int32 and assigns it to the ModifiedTime field.
func (o *ZoneStats) SetModifiedTime(v int32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *ZoneStats) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZoneStats) SetName(v string) {
	o.Name = v
}

// GetNumAssets returns the NumAssets field value if set, zero value otherwise.
func (o *ZoneStats) GetNumAssets() int32 {
	if o == nil || IsNil(o.NumAssets) {
		var ret int32
		return ret
	}
	return *o.NumAssets
}

// GetNumAssetsOk returns a tuple with the NumAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetNumAssetsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAssets) {
		return nil, false
	}
	return o.NumAssets, true
}

// HasNumAssets returns a boolean if a field has been set.
func (o *ZoneStats) HasNumAssets() bool {
	if o != nil && !IsNil(o.NumAssets) {
		return true
	}

	return false
}

// SetNumAssets gets a reference to the given int32 and assigns it to the NumAssets field.
func (o *ZoneStats) SetNumAssets(v int32) {
	o.NumAssets = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *ZoneStats) GetNumClients() int32 {
	if o == nil || IsNil(o.NumClients) {
		var ret int32
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetNumClientsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumClients) {
		return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *ZoneStats) HasNumClients() bool {
	if o != nil && !IsNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given int32 and assigns it to the NumClients field.
func (o *ZoneStats) SetNumClients(v int32) {
	o.NumClients = &v
}

// GetNumSdkclients returns the NumSdkclients field value if set, zero value otherwise.
func (o *ZoneStats) GetNumSdkclients() int32 {
	if o == nil || IsNil(o.NumSdkclients) {
		var ret int32
		return ret
	}
	return *o.NumSdkclients
}

// GetNumSdkclientsOk returns a tuple with the NumSdkclients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetNumSdkclientsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumSdkclients) {
		return nil, false
	}
	return o.NumSdkclients, true
}

// HasNumSdkclients returns a boolean if a field has been set.
func (o *ZoneStats) HasNumSdkclients() bool {
	if o != nil && !IsNil(o.NumSdkclients) {
		return true
	}

	return false
}

// SetNumSdkclients gets a reference to the given int32 and assigns it to the NumSdkclients field.
func (o *ZoneStats) SetNumSdkclients(v int32) {
	o.NumSdkclients = &v
}

// GetOccupancyLimit returns the OccupancyLimit field value if set, zero value otherwise.
func (o *ZoneStats) GetOccupancyLimit() int32 {
	if o == nil || IsNil(o.OccupancyLimit) {
		var ret int32
		return ret
	}
	return *o.OccupancyLimit
}

// GetOccupancyLimitOk returns a tuple with the OccupancyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetOccupancyLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.OccupancyLimit) {
		return nil, false
	}
	return o.OccupancyLimit, true
}

// HasOccupancyLimit returns a boolean if a field has been set.
func (o *ZoneStats) HasOccupancyLimit() bool {
	if o != nil && !IsNil(o.OccupancyLimit) {
		return true
	}

	return false
}

// SetOccupancyLimit gets a reference to the given int32 and assigns it to the OccupancyLimit field.
func (o *ZoneStats) SetOccupancyLimit(v int32) {
	o.OccupancyLimit = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ZoneStats) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ZoneStats) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ZoneStats) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSdkclientsWaits returns the SdkclientsWaits field value if set, zero value otherwise.
func (o *ZoneStats) GetSdkclientsWaits() ZoneStatsSdkclientsWaits {
	if o == nil || IsNil(o.SdkclientsWaits) {
		var ret ZoneStatsSdkclientsWaits
		return ret
	}
	return *o.SdkclientsWaits
}

// GetSdkclientsWaitsOk returns a tuple with the SdkclientsWaits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetSdkclientsWaitsOk() (*ZoneStatsSdkclientsWaits, bool) {
	if o == nil || IsNil(o.SdkclientsWaits) {
		return nil, false
	}
	return o.SdkclientsWaits, true
}

// HasSdkclientsWaits returns a boolean if a field has been set.
func (o *ZoneStats) HasSdkclientsWaits() bool {
	if o != nil && !IsNil(o.SdkclientsWaits) {
		return true
	}

	return false
}

// SetSdkclientsWaits gets a reference to the given ZoneStatsSdkclientsWaits and assigns it to the SdkclientsWaits field.
func (o *ZoneStats) SetSdkclientsWaits(v ZoneStatsSdkclientsWaits) {
	o.SdkclientsWaits = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ZoneStats) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ZoneStats) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ZoneStats) SetSiteId(v string) {
	o.SiteId = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *ZoneStats) GetVertices() []ZoneVertex {
	if o == nil || IsNil(o.Vertices) {
		var ret []ZoneVertex
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetVerticesOk() ([]ZoneVertex, bool) {
	if o == nil || IsNil(o.Vertices) {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *ZoneStats) HasVertices() bool {
	if o != nil && !IsNil(o.Vertices) {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []ZoneVertex and assigns it to the Vertices field.
func (o *ZoneStats) SetVertices(v []ZoneVertex) {
	o.Vertices = v
}

// GetVerticesM returns the VerticesM field value if set, zero value otherwise.
func (o *ZoneStats) GetVerticesM() []ZoneVertexM {
	if o == nil || IsNil(o.VerticesM) {
		var ret []ZoneVertexM
		return ret
	}
	return o.VerticesM
}

// GetVerticesMOk returns a tuple with the VerticesM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneStats) GetVerticesMOk() ([]ZoneVertexM, bool) {
	if o == nil || IsNil(o.VerticesM) {
		return nil, false
	}
	return o.VerticesM, true
}

// HasVerticesM returns a boolean if a field has been set.
func (o *ZoneStats) HasVerticesM() bool {
	if o != nil && !IsNil(o.VerticesM) {
		return true
	}

	return false
}

// SetVerticesM gets a reference to the given []ZoneVertexM and assigns it to the VerticesM field.
func (o *ZoneStats) SetVerticesM(v []ZoneVertexM) {
	o.VerticesM = v
}

func (o ZoneStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetsWaits) {
		toSerialize["assets_waits"] = o.AssetsWaits
	}
	if !IsNil(o.ClientsWaits) {
		toSerialize["clients_waits"] = o.ClientsWaits
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["id"] = o.Id
	toSerialize["map_id"] = o.MapId
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NumAssets) {
		toSerialize["num_assets"] = o.NumAssets
	}
	if !IsNil(o.NumClients) {
		toSerialize["num_clients"] = o.NumClients
	}
	if !IsNil(o.NumSdkclients) {
		toSerialize["num_sdkclients"] = o.NumSdkclients
	}
	if !IsNil(o.OccupancyLimit) {
		toSerialize["occupancy_limit"] = o.OccupancyLimit
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SdkclientsWaits) {
		toSerialize["sdkclients_waits"] = o.SdkclientsWaits
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Vertices) {
		toSerialize["vertices"] = o.Vertices
	}
	if !IsNil(o.VerticesM) {
		toSerialize["vertices_m"] = o.VerticesM
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZoneStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"map_id",
		"name",
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

	varZoneStats := _ZoneStats{}

	err = json.Unmarshal(data, &varZoneStats)

	if err != nil {
		return err
	}

	*o = ZoneStats(varZoneStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets_waits")
		delete(additionalProperties, "clients_waits")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "num_assets")
		delete(additionalProperties, "num_clients")
		delete(additionalProperties, "num_sdkclients")
		delete(additionalProperties, "occupancy_limit")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "sdkclients_waits")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "vertices")
		delete(additionalProperties, "vertices_m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZoneStats struct {
	value *ZoneStats
	isSet bool
}

func (v NullableZoneStats) Get() *ZoneStats {
	return v.value
}

func (v *NullableZoneStats) Set(val *ZoneStats) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneStats) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneStats(val *ZoneStats) *NullableZoneStats {
	return &NullableZoneStats{value: val, isSet: true}
}

func (v NullableZoneStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


