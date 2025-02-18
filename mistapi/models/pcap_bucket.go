package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// PcapBucket represents a PcapBucket struct.
type PcapBucket struct {
    Bucket               string                 `json:"bucket"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PcapBucket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PcapBucket) String() string {
    return fmt.Sprintf(
    	"PcapBucket[Bucket=%v, AdditionalProperties=%v]",
    	p.Bucket, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PcapBucket.
// It customizes the JSON marshaling process for PcapBucket objects.
func (p PcapBucket) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "bucket"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PcapBucket object to a map representation for JSON marshaling.
func (p PcapBucket) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["bucket"] = p.Bucket
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PcapBucket.
// It customizes the JSON unmarshaling process for PcapBucket objects.
func (p *PcapBucket) UnmarshalJSON(input []byte) error {
    var temp tempPcapBucket
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bucket")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Bucket = *temp.Bucket
    return nil
}

// tempPcapBucket is a temporary struct used for validating the fields of PcapBucket.
type tempPcapBucket  struct {
    Bucket *string `json:"bucket"`
}

func (p *tempPcapBucket) validate() error {
    var errs []string
    if p.Bucket == nil {
        errs = append(errs, "required field `bucket` is missing for type `pcap_bucket`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
