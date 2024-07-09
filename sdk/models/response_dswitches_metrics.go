package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDswitchesMetrics represents a ResponseDswitchesMetrics struct.
type ResponseDswitchesMetrics struct {
    InactiveWiredVlans   DswitchesMetricsInactiveWiredVlans `json:"inactive_wired_vlans"`
    PoeCompliance        DswitchesMetricsPoeCompliance      `json:"poe_compliance"`
    SwitchApAffinity     DswitchesMetricsSwitchApAffinity   `json:"switch_ap_affinity"`
    VersionCompliance    DswitchesMetricsVersionCompliance  `json:"version_compliance"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDswitchesMetrics.
// It customizes the JSON marshaling process for ResponseDswitchesMetrics objects.
func (r ResponseDswitchesMetrics) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDswitchesMetrics object to a map representation for JSON marshaling.
func (r ResponseDswitchesMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["inactive_wired_vlans"] = r.InactiveWiredVlans.toMap()
    structMap["poe_compliance"] = r.PoeCompliance.toMap()
    structMap["switch_ap_affinity"] = r.SwitchApAffinity.toMap()
    structMap["version_compliance"] = r.VersionCompliance.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDswitchesMetrics.
// It customizes the JSON unmarshaling process for ResponseDswitchesMetrics objects.
func (r *ResponseDswitchesMetrics) UnmarshalJSON(input []byte) error {
    var temp responseDswitchesMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "inactive_wired_vlans", "poe_compliance", "switch_ap_affinity", "version_compliance")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.InactiveWiredVlans = *temp.InactiveWiredVlans
    r.PoeCompliance = *temp.PoeCompliance
    r.SwitchApAffinity = *temp.SwitchApAffinity
    r.VersionCompliance = *temp.VersionCompliance
    return nil
}

// responseDswitchesMetrics is a temporary struct used for validating the fields of ResponseDswitchesMetrics.
type responseDswitchesMetrics  struct {
    InactiveWiredVlans *DswitchesMetricsInactiveWiredVlans `json:"inactive_wired_vlans"`
    PoeCompliance      *DswitchesMetricsPoeCompliance      `json:"poe_compliance"`
    SwitchApAffinity   *DswitchesMetricsSwitchApAffinity   `json:"switch_ap_affinity"`
    VersionCompliance  *DswitchesMetricsVersionCompliance  `json:"version_compliance"`
}

func (r *responseDswitchesMetrics) validate() error {
    var errs []string
    if r.InactiveWiredVlans == nil {
        errs = append(errs, "required field `inactive_wired_vlans` is missing for type `Response_Dswitches_Metrics`")
    }
    if r.PoeCompliance == nil {
        errs = append(errs, "required field `poe_compliance` is missing for type `Response_Dswitches_Metrics`")
    }
    if r.SwitchApAffinity == nil {
        errs = append(errs, "required field `switch_ap_affinity` is missing for type `Response_Dswitches_Metrics`")
    }
    if r.VersionCompliance == nil {
        errs = append(errs, "required field `version_compliance` is missing for type `Response_Dswitches_Metrics`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}