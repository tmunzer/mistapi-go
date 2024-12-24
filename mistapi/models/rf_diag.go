package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// RfDiag represents a RfDiag struct.
// RF Diag
type RfDiag struct {
    // recording length in seconds, max is 180. Default value is also 180.
    Duration             *int                   `json:"duration,omitempty"`
    // if `type`==`client` or `asset`, mac of the device
    Mac                  *string                `json:"mac,omitempty"`
    // name of the recording, the name of the sdk client would be a good default choice
    Name                 string                 `json:"name"`
    // if `type`==`sdkclient`, sdkclient_id of this recording
    SdkclientId          *uuid.UUID             `json:"sdkclient_id,omitempty"`
    // enum: `asset`, `client`, `sdkclient`
    Type                 RfClientTypeEnum       `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RfDiag,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RfDiag) String() string {
    return fmt.Sprintf(
    	"RfDiag[Duration=%v, Mac=%v, Name=%v, SdkclientId=%v, Type=%v, AdditionalProperties=%v]",
    	r.Duration, r.Mac, r.Name, r.SdkclientId, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RfDiag.
// It customizes the JSON marshaling process for RfDiag objects.
func (r RfDiag) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "duration", "mac", "name", "sdkclient_id", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RfDiag object to a map representation for JSON marshaling.
func (r RfDiag) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Duration != nil {
        structMap["duration"] = r.Duration
    }
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    structMap["name"] = r.Name
    if r.SdkclientId != nil {
        structMap["sdkclient_id"] = r.SdkclientId
    }
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RfDiag.
// It customizes the JSON unmarshaling process for RfDiag objects.
func (r *RfDiag) UnmarshalJSON(input []byte) error {
    var temp tempRfDiag
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "mac", "name", "sdkclient_id", "type")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Duration = temp.Duration
    r.Mac = temp.Mac
    r.Name = *temp.Name
    r.SdkclientId = temp.SdkclientId
    r.Type = *temp.Type
    return nil
}

// tempRfDiag is a temporary struct used for validating the fields of RfDiag.
type tempRfDiag  struct {
    Duration    *int              `json:"duration,omitempty"`
    Mac         *string           `json:"mac,omitempty"`
    Name        *string           `json:"name"`
    SdkclientId *uuid.UUID        `json:"sdkclient_id,omitempty"`
    Type        *RfClientTypeEnum `json:"type"`
}

func (r *tempRfDiag) validate() error {
    var errs []string
    if r.Name == nil {
        errs = append(errs, "required field `name` is missing for type `rf_diag`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `rf_diag`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
