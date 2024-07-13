package models

import (
    "encoding/json"
)

// SiteSettingMxedge represents a SiteSettingMxedge struct.
// site mist edges form a cluster of radsecproxy servers
type SiteSettingMxedge struct {
    // configure cloud-assisted dynamic authorization service on this cluster of mist edges
    MistDas              *MxedgeDas       `json:"mist_das,omitempty"`
    // MxEdge Radsec Configuration
    Radsec               *MxclusterRadsec `json:"radsec,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingMxedge.
// It customizes the JSON marshaling process for SiteSettingMxedge objects.
func (s SiteSettingMxedge) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingMxedge object to a map representation for JSON marshaling.
func (s SiteSettingMxedge) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.MistDas != nil {
        structMap["mist_das"] = s.MistDas.toMap()
    }
    if s.Radsec != nil {
        structMap["radsec"] = s.Radsec.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingMxedge.
// It customizes the JSON unmarshaling process for SiteSettingMxedge objects.
func (s *SiteSettingMxedge) UnmarshalJSON(input []byte) error {
    var temp siteSettingMxedge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mist_das", "radsec")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.MistDas = temp.MistDas
    s.Radsec = temp.Radsec
    return nil
}

// siteSettingMxedge is a temporary struct used for validating the fields of SiteSettingMxedge.
type siteSettingMxedge  struct {
    MistDas *MxedgeDas       `json:"mist_das,omitempty"`
    Radsec  *MxclusterRadsec `json:"radsec,omitempty"`
}
