package models

import (
    "encoding/json"
    "fmt"
)

// FwupdateStat represents a FwupdateStat struct.
type FwupdateStat struct {
    Progress             Optional[int]                    `json:"progress"`
    // enum: `inprogress`, `failed`, `upgraded`
    Status               Optional[FwupdateStatStatusEnum] `json:"status"`
    StatusId             Optional[int]                    `json:"status_id"`
    // Epoch (seconds)
    Timestamp            *float64                         `json:"timestamp,omitempty"`
    WillRetry            Optional[bool]                   `json:"will_retry"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for FwupdateStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FwupdateStat) String() string {
    return fmt.Sprintf(
    	"FwupdateStat[Progress=%v, Status=%v, StatusId=%v, Timestamp=%v, WillRetry=%v, AdditionalProperties=%v]",
    	f.Progress, f.Status, f.StatusId, f.Timestamp, f.WillRetry, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FwupdateStat.
// It customizes the JSON marshaling process for FwupdateStat objects.
func (f FwupdateStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "progress", "status", "status_id", "timestamp", "will_retry"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FwupdateStat object to a map representation for JSON marshaling.
func (f FwupdateStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Progress.IsValueSet() {
        if f.Progress.Value() != nil {
            structMap["progress"] = f.Progress.Value()
        } else {
            structMap["progress"] = nil
        }
    }
    if f.Status.IsValueSet() {
        if f.Status.Value() != nil {
            structMap["status"] = f.Status.Value()
        } else {
            structMap["status"] = nil
        }
    }
    if f.StatusId.IsValueSet() {
        if f.StatusId.Value() != nil {
            structMap["status_id"] = f.StatusId.Value()
        } else {
            structMap["status_id"] = nil
        }
    }
    if f.Timestamp != nil {
        structMap["timestamp"] = f.Timestamp
    }
    if f.WillRetry.IsValueSet() {
        if f.WillRetry.Value() != nil {
            structMap["will_retry"] = f.WillRetry.Value()
        } else {
            structMap["will_retry"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FwupdateStat.
// It customizes the JSON unmarshaling process for FwupdateStat objects.
func (f *FwupdateStat) UnmarshalJSON(input []byte) error {
    var temp tempFwupdateStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "progress", "status", "status_id", "timestamp", "will_retry")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Progress = temp.Progress
    f.Status = temp.Status
    f.StatusId = temp.StatusId
    f.Timestamp = temp.Timestamp
    f.WillRetry = temp.WillRetry
    return nil
}

// tempFwupdateStat is a temporary struct used for validating the fields of FwupdateStat.
type tempFwupdateStat  struct {
    Progress  Optional[int]                    `json:"progress"`
    Status    Optional[FwupdateStatStatusEnum] `json:"status"`
    StatusId  Optional[int]                    `json:"status_id"`
    Timestamp *float64                         `json:"timestamp,omitempty"`
    WillRetry Optional[bool]                   `json:"will_retry"`
}
