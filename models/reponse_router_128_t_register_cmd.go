package models

import (
    "encoding/json"
)

// ReponseRouter128TRegisterCmd represents a ReponseRouter128TRegisterCmd struct.
type ReponseRouter128TRegisterCmd struct {
    ConductorCmd         *string        `json:"conductor_cmd,omitempty"`
    RegistrationCode     *string        `json:"registration_code,omitempty"`
    RouterShellCmd       *string        `json:"router_shell_cmd,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReponseRouter128TRegisterCmd.
// It customizes the JSON marshaling process for ReponseRouter128TRegisterCmd objects.
func (r ReponseRouter128TRegisterCmd) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReponseRouter128TRegisterCmd object to a map representation for JSON marshaling.
func (r ReponseRouter128TRegisterCmd) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ConductorCmd != nil {
        structMap["conductor_cmd"] = r.ConductorCmd
    }
    if r.RegistrationCode != nil {
        structMap["registration_code"] = r.RegistrationCode
    }
    if r.RouterShellCmd != nil {
        structMap["router_shell_cmd"] = r.RouterShellCmd
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReponseRouter128TRegisterCmd.
// It customizes the JSON unmarshaling process for ReponseRouter128TRegisterCmd objects.
func (r *ReponseRouter128TRegisterCmd) UnmarshalJSON(input []byte) error {
    var temp reponseRouter128TRegisterCmd
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "conductor_cmd", "registration_code", "router_shell_cmd")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ConductorCmd = temp.ConductorCmd
    r.RegistrationCode = temp.RegistrationCode
    r.RouterShellCmd = temp.RouterShellCmd
    return nil
}

// reponseRouter128TRegisterCmd is a temporary struct used for validating the fields of ReponseRouter128TRegisterCmd.
type reponseRouter128TRegisterCmd  struct {
    ConductorCmd     *string `json:"conductor_cmd,omitempty"`
    RegistrationCode *string `json:"registration_code,omitempty"`
    RouterShellCmd   *string `json:"router_shell_cmd,omitempty"`
}
