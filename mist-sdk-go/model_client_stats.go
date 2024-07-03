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
	"gopkg.in/validator.v2"
	"fmt"
)

// ClientStats - struct for ClientStats
type ClientStats struct {
	StatsWiredClient *StatsWiredClient
	ArrayOfClientWirelessStats *[]ClientWirelessStats
}

// StatsWiredClientAsClientStats is a convenience function that returns StatsWiredClient wrapped in ClientStats
func StatsWiredClientAsClientStats(v *StatsWiredClient) ClientStats {
	return ClientStats{
		StatsWiredClient: v,
	}
}

// []ClientWirelessStatsAsClientStats is a convenience function that returns []ClientWirelessStats wrapped in ClientStats
func ArrayOfClientWirelessStatsAsClientStats(v *[]ClientWirelessStats) ClientStats {
	return ClientStats{
		ArrayOfClientWirelessStats: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClientStats) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StatsWiredClient
	err = newStrictDecoder(data).Decode(&dst.StatsWiredClient)
	if err == nil {
		jsonStatsWiredClient, _ := json.Marshal(dst.StatsWiredClient)
		if string(jsonStatsWiredClient) == "{}" { // empty struct
			dst.StatsWiredClient = nil
		} else {
			if err = validator.Validate(dst.StatsWiredClient); err != nil {
				dst.StatsWiredClient = nil
			} else {
				match++
			}
		}
	} else {
		dst.StatsWiredClient = nil
	}

	// try to unmarshal data into ArrayOfClientWirelessStats
	err = newStrictDecoder(data).Decode(&dst.ArrayOfClientWirelessStats)
	if err == nil {
		jsonArrayOfClientWirelessStats, _ := json.Marshal(dst.ArrayOfClientWirelessStats)
		if string(jsonArrayOfClientWirelessStats) == "{}" { // empty struct
			dst.ArrayOfClientWirelessStats = nil
		} else {
			if err = validator.Validate(dst.ArrayOfClientWirelessStats); err != nil {
				dst.ArrayOfClientWirelessStats = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfClientWirelessStats = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.StatsWiredClient = nil
		dst.ArrayOfClientWirelessStats = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ClientStats)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ClientStats)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClientStats) MarshalJSON() ([]byte, error) {
	if src.StatsWiredClient != nil {
		return json.Marshal(&src.StatsWiredClient)
	}

	if src.ArrayOfClientWirelessStats != nil {
		return json.Marshal(&src.ArrayOfClientWirelessStats)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClientStats) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StatsWiredClient != nil {
		return obj.StatsWiredClient
	}

	if obj.ArrayOfClientWirelessStats != nil {
		return obj.ArrayOfClientWirelessStats
	}

	// all schemas are nil
	return nil
}

type NullableClientStats struct {
	value *ClientStats
	isSet bool
}

func (v NullableClientStats) Get() *ClientStats {
	return v.value
}

func (v *NullableClientStats) Set(val *ClientStats) {
	v.value = val
	v.isSet = true
}

func (v NullableClientStats) IsSet() bool {
	return v.isSet
}

func (v *NullableClientStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientStats(val *ClientStats) *NullableClientStats {
	return &NullableClientStats{value: val, isSet: true}
}

func (v NullableClientStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

