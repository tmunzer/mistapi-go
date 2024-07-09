package models

import (
    "encoding/json"
)

// AssetOfInterest represents a AssetOfInterest struct.
type AssetOfInterest struct {
    CheckpointPrep       *float64       `json:"_checkpoint_prep,omitempty"`
    CheckpointPreparer   *float64       `json:"_checkpoint_preparer,omitempty"`
    CheckpointScan       *float64       `json:"_checkpoint_scan,omitempty"`
    Timestamp            *float64       `json:"_timestamp,omitempty"`
    Ttl                  *float64       `json:"_ttl,omitempty"`
    ApMac                *string        `json:"ap_mac,omitempty"`
    Beam                 *float64       `json:"beam,omitempty"`
    By                   *string        `json:"by,omitempty"`
    CurrSite             *string        `json:"curr_site,omitempty"`
    DeviceName           *string        `json:"device_name,omitempty"`
    Id                   *string        `json:"id,omitempty"`
    LastSeen             *float64       `json:"last_seen,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    Manufacture          *string        `json:"manufacture,omitempty"`
    MapId                *string        `json:"map_id,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Rssi                 *float64       `json:"rssi,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AssetOfInterest.
// It customizes the JSON marshaling process for AssetOfInterest objects.
func (a AssetOfInterest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AssetOfInterest object to a map representation for JSON marshaling.
func (a AssetOfInterest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CheckpointPrep != nil {
        structMap["_checkpoint_prep"] = a.CheckpointPrep
    }
    if a.CheckpointPreparer != nil {
        structMap["_checkpoint_preparer"] = a.CheckpointPreparer
    }
    if a.CheckpointScan != nil {
        structMap["_checkpoint_scan"] = a.CheckpointScan
    }
    if a.Timestamp != nil {
        structMap["_timestamp"] = a.Timestamp
    }
    if a.Ttl != nil {
        structMap["_ttl"] = a.Ttl
    }
    if a.ApMac != nil {
        structMap["ap_mac"] = a.ApMac
    }
    if a.Beam != nil {
        structMap["beam"] = a.Beam
    }
    if a.By != nil {
        structMap["by"] = a.By
    }
    if a.CurrSite != nil {
        structMap["curr_site"] = a.CurrSite
    }
    if a.DeviceName != nil {
        structMap["device_name"] = a.DeviceName
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.LastSeen != nil {
        structMap["last_seen"] = a.LastSeen
    }
    if a.Mac != nil {
        structMap["mac"] = a.Mac
    }
    if a.Manufacture != nil {
        structMap["manufacture"] = a.Manufacture
    }
    if a.MapId != nil {
        structMap["map_id"] = a.MapId
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Rssi != nil {
        structMap["rssi"] = a.Rssi
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetOfInterest.
// It customizes the JSON unmarshaling process for AssetOfInterest objects.
func (a *AssetOfInterest) UnmarshalJSON(input []byte) error {
    var temp assetOfInterest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "_checkpoint_prep", "_checkpoint_preparer", "_checkpoint_scan", "_timestamp", "_ttl", "ap_mac", "beam", "by", "curr_site", "device_name", "id", "last_seen", "mac", "manufacture", "map_id", "name", "rssi")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.CheckpointPrep = temp.CheckpointPrep
    a.CheckpointPreparer = temp.CheckpointPreparer
    a.CheckpointScan = temp.CheckpointScan
    a.Timestamp = temp.Timestamp
    a.Ttl = temp.Ttl
    a.ApMac = temp.ApMac
    a.Beam = temp.Beam
    a.By = temp.By
    a.CurrSite = temp.CurrSite
    a.DeviceName = temp.DeviceName
    a.Id = temp.Id
    a.LastSeen = temp.LastSeen
    a.Mac = temp.Mac
    a.Manufacture = temp.Manufacture
    a.MapId = temp.MapId
    a.Name = temp.Name
    a.Rssi = temp.Rssi
    return nil
}

// assetOfInterest is a temporary struct used for validating the fields of AssetOfInterest.
type assetOfInterest  struct {
    CheckpointPrep     *float64 `json:"_checkpoint_prep,omitempty"`
    CheckpointPreparer *float64 `json:"_checkpoint_preparer,omitempty"`
    CheckpointScan     *float64 `json:"_checkpoint_scan,omitempty"`
    Timestamp          *float64 `json:"_timestamp,omitempty"`
    Ttl                *float64 `json:"_ttl,omitempty"`
    ApMac              *string  `json:"ap_mac,omitempty"`
    Beam               *float64 `json:"beam,omitempty"`
    By                 *string  `json:"by,omitempty"`
    CurrSite           *string  `json:"curr_site,omitempty"`
    DeviceName         *string  `json:"device_name,omitempty"`
    Id                 *string  `json:"id,omitempty"`
    LastSeen           *float64 `json:"last_seen,omitempty"`
    Mac                *string  `json:"mac,omitempty"`
    Manufacture        *string  `json:"manufacture,omitempty"`
    MapId              *string  `json:"map_id,omitempty"`
    Name               *string  `json:"name,omitempty"`
    Rssi               *float64 `json:"rssi,omitempty"`
}
