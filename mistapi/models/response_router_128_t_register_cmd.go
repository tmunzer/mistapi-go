package models

import (
    "encoding/json"
    "fmt"
)

// ResponseRouter128tRegisterCmd represents a ResponseRouter128tRegisterCmd struct.
type ResponseRouter128tRegisterCmd struct {
    ConductorCmd         *string                `json:"conductor_cmd,omitempty"`
    RegistrationCode     *string                `json:"registration_code,omitempty"`
    RouterShellCmd       *string                `json:"router_shell_cmd,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRouter128tRegisterCmd,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRouter128tRegisterCmd) String() string {
    return fmt.Sprintf(
    	"ResponseRouter128tRegisterCmd[ConductorCmd=%v, RegistrationCode=%v, RouterShellCmd=%v, AdditionalProperties=%v]",
    	r.ConductorCmd, r.RegistrationCode, r.RouterShellCmd, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRouter128tRegisterCmd.
// It customizes the JSON marshaling process for ResponseRouter128tRegisterCmd objects.
func (r ResponseRouter128tRegisterCmd) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "conductor_cmd", "registration_code", "router_shell_cmd"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseRouter128tRegisterCmd object to a map representation for JSON marshaling.
func (r ResponseRouter128tRegisterCmd) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRouter128tRegisterCmd.
// It customizes the JSON unmarshaling process for ResponseRouter128tRegisterCmd objects.
func (r *ResponseRouter128tRegisterCmd) UnmarshalJSON(input []byte) error {
    var temp tempResponseRouter128tRegisterCmd
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

// tempResponseRouter128tRegisterCmd is a temporary struct used for validating the fields of ResponseRouter128tRegisterCmd.
type tempResponseRouter128tRegisterCmd  struct {
    ConductorCmd     *string `json:"conductor_cmd,omitempty"`
    RegistrationCode *string `json:"registration_code,omitempty"`
    RouterShellCmd   *string `json:"router_shell_cmd,omitempty"`
}
