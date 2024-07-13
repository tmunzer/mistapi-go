package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PcapBucket represents a PcapBucket struct.
type PcapBucket struct {
    Bucket               string         `json:"bucket"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PcapBucket.
// It customizes the JSON marshaling process for PcapBucket objects.
func (p PcapBucket) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PcapBucket object to a map representation for JSON marshaling.
func (p PcapBucket) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["bucket"] = p.Bucket
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PcapBucket.
// It customizes the JSON unmarshaling process for PcapBucket objects.
func (p *PcapBucket) UnmarshalJSON(input []byte) error {
    var temp pcapBucket
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bucket")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Bucket = *temp.Bucket
    return nil
}

// pcapBucket is a temporary struct used for validating the fields of PcapBucket.
type pcapBucket  struct {
    Bucket *string `json:"bucket"`
}

func (p *pcapBucket) validate() error {
    var errs []string
    if p.Bucket == nil {
        errs = append(errs, "required field `bucket` is missing for type `Pcap_Bucket`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}