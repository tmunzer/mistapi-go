package models

import (
    "encoding/json"
)

// ReponseRouter128tRegisterCmd represents a ReponseRouter128tRegisterCmd struct.
type ReponseRouter128tRegisterCmd struct {
    ConductorCmd         *string                `json:"conductor_cmd,omitempty"`
    RegistrationCode     *string                `json:"registration_code,omitempty"`
    RouterShellCmd       *string                `json:"router_shell_cmd,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReponseRouter128tRegisterCmd.
// It customizes the JSON marshaling process for ReponseRouter128tRegisterCmd objects.
func (r ReponseRouter128tRegisterCmd) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "conductor_cmd", "registration_code", "router_shell_cmd"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReponseRouter128tRegisterCmd object to a map representation for JSON marshaling.
func (r ReponseRouter128tRegisterCmd) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ReponseRouter128tRegisterCmd.
// It customizes the JSON unmarshaling process for ReponseRouter128tRegisterCmd objects.
func (r *ReponseRouter128tRegisterCmd) UnmarshalJSON(input []byte) error {
    var temp tempReponseRouter128tRegisterCmd
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "conductor_cmd", "registration_code", "router_shell_cmd")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ConductorCmd = temp.ConductorCmd
    r.RegistrationCode = temp.RegistrationCode
    r.RouterShellCmd = temp.RouterShellCmd
    return nil
}

// tempReponseRouter128tRegisterCmd is a temporary struct used for validating the fields of ReponseRouter128tRegisterCmd.
type tempReponseRouter128tRegisterCmd  struct {
    ConductorCmd     *string `json:"conductor_cmd,omitempty"`
    RegistrationCode *string `json:"registration_code,omitempty"`
    RouterShellCmd   *string `json:"router_shell_cmd,omitempty"`
}
