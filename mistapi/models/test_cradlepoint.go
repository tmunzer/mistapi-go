package models

import (
    "encoding/json"
    "fmt"
)

// TestCradlepoint represents a TestCradlepoint struct.
type TestCradlepoint struct {
    // if status is `inactive` this field returns the reason for it being inactive.
    Error                *string                        `json:"error,omitempty"`
    // status of integration detected during last sync. enum: `active`, `inactive`
    LastStatus           *TestCradlepointLastStatusEnum `json:"last_status,omitempty"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for TestCradlepoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TestCradlepoint) String() string {
    return fmt.Sprintf(
    	"TestCradlepoint[Error=%v, LastStatus=%v, AdditionalProperties=%v]",
    	t.Error, t.LastStatus, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TestCradlepoint.
// It customizes the JSON marshaling process for TestCradlepoint objects.
func (t TestCradlepoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "error", "last_status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TestCradlepoint object to a map representation for JSON marshaling.
func (t TestCradlepoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Error != nil {
        structMap["error"] = t.Error
    }
    if t.LastStatus != nil {
        structMap["last_status"] = t.LastStatus
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TestCradlepoint.
// It customizes the JSON unmarshaling process for TestCradlepoint objects.
func (t *TestCradlepoint) UnmarshalJSON(input []byte) error {
    var temp tempTestCradlepoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "error", "last_status")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Error = temp.Error
    t.LastStatus = temp.LastStatus
    return nil
}

// tempTestCradlepoint is a temporary struct used for validating the fields of TestCradlepoint.
type tempTestCradlepoint  struct {
    Error      *string                        `json:"error,omitempty"`
    LastStatus *TestCradlepointLastStatusEnum `json:"last_status,omitempty"`
}
