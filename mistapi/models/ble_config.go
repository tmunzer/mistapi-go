package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// BleConfig represents a BleConfig struct.
// BLE AP settings
type BleConfig struct {
    // whether Mist beacons is enabled
    BeaconEnabled           *bool                        `json:"beacon_enabled,omitempty"`
    // required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second
    BeaconRate              *int                         `json:"beacon_rate,omitempty"`
    // enum: `custom`, `default`
    BeaconRateMode          *BleConfigBeaconRateModeEnum `json:"beacon_rate_mode,omitempty"`
    // list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam)
    BeamDisabled            []int                        `json:"beam_disabled,omitempty"`
    // can be enabled if `beacon_enabled`==`true`, whether to send custom packet
    CustomBlePacketEnabled  *bool                        `json:"custom_ble_packet_enabled,omitempty"`
    // The custom frame to be sent out in this beacon. The frame must be a hexstring
    CustomBlePacketFrame    *string                      `json:"custom_ble_packet_frame,omitempty"`
    // Frequency (msec) of data emitted by custom ble beacon
    CustomBlePacketFreqMsec *int                         `json:"custom_ble_packet_freq_msec,omitempty"`
    // advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
    EddystoneUidAdvPower    *int                         `json:"eddystone_uid_adv_power,omitempty"`
    EddystoneUidBeams       *string                      `json:"eddystone_uid_beams,omitempty"`
    // only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled
    EddystoneUidEnabled     *bool                        `json:"eddystone_uid_enabled,omitempty"`
    // Frequency (msec) of data emmit by Eddystone-UID beacon
    EddystoneUidFreqMsec    *int                         `json:"eddystone_uid_freq_msec,omitempty"`
    // Eddystone-UID instance for the device
    EddystoneUidInstance    *string                      `json:"eddystone_uid_instance,omitempty"`
    // Eddystone-UID namespace
    EddystoneUidNamespace   *string                      `json:"eddystone_uid_namespace,omitempty"`
    // advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
    EddystoneUrlAdvPower    *int                         `json:"eddystone_url_adv_power,omitempty"`
    EddystoneUrlBeams       *string                      `json:"eddystone_url_beams,omitempty"`
    // only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled
    EddystoneUrlEnabled     *bool                        `json:"eddystone_url_enabled,omitempty"`
    // Frequency (msec) of data emit by Eddystone-UID beacon
    EddystoneUrlFreqMsec    *int                         `json:"eddystone_url_freq_msec,omitempty"`
    // URL pointed by Eddystone-URL beacon
    EddystoneUrlUrl         *string                      `json:"eddystone_url_url,omitempty"`
    // advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
    IbeaconAdvPower         *int                         `json:"ibeacon_adv_power,omitempty"`
    IbeaconBeams            *string                      `json:"ibeacon_beams,omitempty"`
    // can be enabled if `beacon_enabled`==`true`, whether to send iBeacon
    IbeaconEnabled          *bool                        `json:"ibeacon_enabled,omitempty"`
    // Frequency (msec) of data emmit for iBeacon
    IbeaconFreqMsec         *int                         `json:"ibeacon_freq_msec,omitempty"`
    // Major number for iBeacon
    IbeaconMajor            *int                         `json:"ibeacon_major,omitempty"`
    // Minor number for iBeacon
    IbeaconMinor            *int                         `json:"ibeacon_minor,omitempty"`
    // optional, if not specified, the same UUID as the beacon will be used
    IbeaconUuid             *uuid.UUID                   `json:"ibeacon_uuid,omitempty"`
    // required if `power_mode`==`custom`; else use `power_mode` as default
    Power                   *int                         `json:"power,omitempty"`
    // enum: `custom`, `default`
    PowerMode               *BleConfigPowerModeEnum      `json:"power_mode,omitempty"`
    AdditionalProperties    map[string]interface{}       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BleConfig.
// It customizes the JSON marshaling process for BleConfig objects.
func (b BleConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "beacon_enabled", "beacon_rate", "beacon_rate_mode", "beam_disabled", "custom_ble_packet_enabled", "custom_ble_packet_frame", "custom_ble_packet_freq_msec", "eddystone_uid_adv_power", "eddystone_uid_beams", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_adv_power", "eddystone_url_beams", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_adv_power", "ibeacon_beams", "ibeacon_enabled", "ibeacon_freq_msec", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "power", "power_mode"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BleConfig object to a map representation for JSON marshaling.
func (b BleConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.BeaconEnabled != nil {
        structMap["beacon_enabled"] = b.BeaconEnabled
    }
    if b.BeaconRate != nil {
        structMap["beacon_rate"] = b.BeaconRate
    }
    if b.BeaconRateMode != nil {
        structMap["beacon_rate_mode"] = b.BeaconRateMode
    }
    if b.BeamDisabled != nil {
        structMap["beam_disabled"] = b.BeamDisabled
    }
    if b.CustomBlePacketEnabled != nil {
        structMap["custom_ble_packet_enabled"] = b.CustomBlePacketEnabled
    }
    if b.CustomBlePacketFrame != nil {
        structMap["custom_ble_packet_frame"] = b.CustomBlePacketFrame
    }
    if b.CustomBlePacketFreqMsec != nil {
        structMap["custom_ble_packet_freq_msec"] = b.CustomBlePacketFreqMsec
    }
    if b.EddystoneUidAdvPower != nil {
        structMap["eddystone_uid_adv_power"] = b.EddystoneUidAdvPower
    }
    if b.EddystoneUidBeams != nil {
        structMap["eddystone_uid_beams"] = b.EddystoneUidBeams
    }
    if b.EddystoneUidEnabled != nil {
        structMap["eddystone_uid_enabled"] = b.EddystoneUidEnabled
    }
    if b.EddystoneUidFreqMsec != nil {
        structMap["eddystone_uid_freq_msec"] = b.EddystoneUidFreqMsec
    }
    if b.EddystoneUidInstance != nil {
        structMap["eddystone_uid_instance"] = b.EddystoneUidInstance
    }
    if b.EddystoneUidNamespace != nil {
        structMap["eddystone_uid_namespace"] = b.EddystoneUidNamespace
    }
    if b.EddystoneUrlAdvPower != nil {
        structMap["eddystone_url_adv_power"] = b.EddystoneUrlAdvPower
    }
    if b.EddystoneUrlBeams != nil {
        structMap["eddystone_url_beams"] = b.EddystoneUrlBeams
    }
    if b.EddystoneUrlEnabled != nil {
        structMap["eddystone_url_enabled"] = b.EddystoneUrlEnabled
    }
    if b.EddystoneUrlFreqMsec != nil {
        structMap["eddystone_url_freq_msec"] = b.EddystoneUrlFreqMsec
    }
    if b.EddystoneUrlUrl != nil {
        structMap["eddystone_url_url"] = b.EddystoneUrlUrl
    }
    if b.IbeaconAdvPower != nil {
        structMap["ibeacon_adv_power"] = b.IbeaconAdvPower
    }
    if b.IbeaconBeams != nil {
        structMap["ibeacon_beams"] = b.IbeaconBeams
    }
    if b.IbeaconEnabled != nil {
        structMap["ibeacon_enabled"] = b.IbeaconEnabled
    }
    if b.IbeaconFreqMsec != nil {
        structMap["ibeacon_freq_msec"] = b.IbeaconFreqMsec
    }
    if b.IbeaconMajor != nil {
        structMap["ibeacon_major"] = b.IbeaconMajor
    }
    if b.IbeaconMinor != nil {
        structMap["ibeacon_minor"] = b.IbeaconMinor
    }
    if b.IbeaconUuid != nil {
        structMap["ibeacon_uuid"] = b.IbeaconUuid
    }
    if b.Power != nil {
        structMap["power"] = b.Power
    }
    if b.PowerMode != nil {
        structMap["power_mode"] = b.PowerMode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BleConfig.
// It customizes the JSON unmarshaling process for BleConfig objects.
func (b *BleConfig) UnmarshalJSON(input []byte) error {
    var temp tempBleConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "beacon_enabled", "beacon_rate", "beacon_rate_mode", "beam_disabled", "custom_ble_packet_enabled", "custom_ble_packet_frame", "custom_ble_packet_freq_msec", "eddystone_uid_adv_power", "eddystone_uid_beams", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_adv_power", "eddystone_url_beams", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_adv_power", "ibeacon_beams", "ibeacon_enabled", "ibeacon_freq_msec", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "power", "power_mode")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.BeaconEnabled = temp.BeaconEnabled
    b.BeaconRate = temp.BeaconRate
    b.BeaconRateMode = temp.BeaconRateMode
    b.BeamDisabled = temp.BeamDisabled
    b.CustomBlePacketEnabled = temp.CustomBlePacketEnabled
    b.CustomBlePacketFrame = temp.CustomBlePacketFrame
    b.CustomBlePacketFreqMsec = temp.CustomBlePacketFreqMsec
    b.EddystoneUidAdvPower = temp.EddystoneUidAdvPower
    b.EddystoneUidBeams = temp.EddystoneUidBeams
    b.EddystoneUidEnabled = temp.EddystoneUidEnabled
    b.EddystoneUidFreqMsec = temp.EddystoneUidFreqMsec
    b.EddystoneUidInstance = temp.EddystoneUidInstance
    b.EddystoneUidNamespace = temp.EddystoneUidNamespace
    b.EddystoneUrlAdvPower = temp.EddystoneUrlAdvPower
    b.EddystoneUrlBeams = temp.EddystoneUrlBeams
    b.EddystoneUrlEnabled = temp.EddystoneUrlEnabled
    b.EddystoneUrlFreqMsec = temp.EddystoneUrlFreqMsec
    b.EddystoneUrlUrl = temp.EddystoneUrlUrl
    b.IbeaconAdvPower = temp.IbeaconAdvPower
    b.IbeaconBeams = temp.IbeaconBeams
    b.IbeaconEnabled = temp.IbeaconEnabled
    b.IbeaconFreqMsec = temp.IbeaconFreqMsec
    b.IbeaconMajor = temp.IbeaconMajor
    b.IbeaconMinor = temp.IbeaconMinor
    b.IbeaconUuid = temp.IbeaconUuid
    b.Power = temp.Power
    b.PowerMode = temp.PowerMode
    return nil
}

// tempBleConfig is a temporary struct used for validating the fields of BleConfig.
type tempBleConfig  struct {
    BeaconEnabled           *bool                        `json:"beacon_enabled,omitempty"`
    BeaconRate              *int                         `json:"beacon_rate,omitempty"`
    BeaconRateMode          *BleConfigBeaconRateModeEnum `json:"beacon_rate_mode,omitempty"`
    BeamDisabled            []int                        `json:"beam_disabled,omitempty"`
    CustomBlePacketEnabled  *bool                        `json:"custom_ble_packet_enabled,omitempty"`
    CustomBlePacketFrame    *string                      `json:"custom_ble_packet_frame,omitempty"`
    CustomBlePacketFreqMsec *int                         `json:"custom_ble_packet_freq_msec,omitempty"`
    EddystoneUidAdvPower    *int                         `json:"eddystone_uid_adv_power,omitempty"`
    EddystoneUidBeams       *string                      `json:"eddystone_uid_beams,omitempty"`
    EddystoneUidEnabled     *bool                        `json:"eddystone_uid_enabled,omitempty"`
    EddystoneUidFreqMsec    *int                         `json:"eddystone_uid_freq_msec,omitempty"`
    EddystoneUidInstance    *string                      `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace   *string                      `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlAdvPower    *int                         `json:"eddystone_url_adv_power,omitempty"`
    EddystoneUrlBeams       *string                      `json:"eddystone_url_beams,omitempty"`
    EddystoneUrlEnabled     *bool                        `json:"eddystone_url_enabled,omitempty"`
    EddystoneUrlFreqMsec    *int                         `json:"eddystone_url_freq_msec,omitempty"`
    EddystoneUrlUrl         *string                      `json:"eddystone_url_url,omitempty"`
    IbeaconAdvPower         *int                         `json:"ibeacon_adv_power,omitempty"`
    IbeaconBeams            *string                      `json:"ibeacon_beams,omitempty"`
    IbeaconEnabled          *bool                        `json:"ibeacon_enabled,omitempty"`
    IbeaconFreqMsec         *int                         `json:"ibeacon_freq_msec,omitempty"`
    IbeaconMajor            *int                         `json:"ibeacon_major,omitempty"`
    IbeaconMinor            *int                         `json:"ibeacon_minor,omitempty"`
    IbeaconUuid             *uuid.UUID                   `json:"ibeacon_uuid,omitempty"`
    Power                   *int                         `json:"power,omitempty"`
    PowerMode               *BleConfigPowerModeEnum      `json:"power_mode,omitempty"`
}
