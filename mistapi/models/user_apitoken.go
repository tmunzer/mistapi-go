package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// UserApitoken represents a UserApitoken struct.
// User API Token
type UserApitoken struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Key                  *string        `json:"key,omitempty"`
    LastUsed             Optional[int]  `json:"last_used"`
    // name of the token
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UserApitoken.
// It customizes the JSON marshaling process for UserApitoken objects.
func (u UserApitoken) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UserApitoken object to a map representation for JSON marshaling.
func (u UserApitoken) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CreatedTime != nil {
        structMap["created_time"] = u.CreatedTime
    }
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.Key != nil {
        structMap["key"] = u.Key
    }
    if u.LastUsed.IsValueSet() {
        if u.LastUsed.Value() != nil {
            structMap["last_used"] = u.LastUsed.Value()
        } else {
            structMap["last_used"] = nil
        }
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserApitoken.
// It customizes the JSON unmarshaling process for UserApitoken objects.
func (u *UserApitoken) UnmarshalJSON(input []byte) error {
    var temp userApitoken
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "id", "key", "last_used", "name")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.CreatedTime = temp.CreatedTime
    u.Id = temp.Id
    u.Key = temp.Key
    u.LastUsed = temp.LastUsed
    u.Name = temp.Name
    return nil
}

// userApitoken is a temporary struct used for validating the fields of UserApitoken.
type userApitoken  struct {
    CreatedTime *float64      `json:"created_time,omitempty"`
    Id          *uuid.UUID    `json:"id,omitempty"`
    Key         *string       `json:"key,omitempty"`
    LastUsed    Optional[int] `json:"last_used"`
    Name        *string       `json:"name,omitempty"`
}
