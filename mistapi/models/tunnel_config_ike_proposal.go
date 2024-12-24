package models

import (
    "encoding/json"
    "fmt"
)

// TunnelConfigIkeProposal represents a TunnelConfigIkeProposal struct.
type TunnelConfigIkeProposal struct {
    // enum: `md5`, `sha1`, `sha2`
    AuthAlgo             *TunnelConfigAuthAlgoEnum         `json:"auth_algo,omitempty"`
    // enum:
    // * 1
    // * 2 (1024-bit)
    // * 5
    // * 14 (default, 2048-bit)
    // * 15 (3072-bit)
    // * 16 (4096-bit)
    // * 19 (256-bit ECP)
    // * 20 (384-bit ECP)
    // * 21 (521-bit ECP)
    // * 24 (2048-bit ECP)
    DhGroup              *TunnelConfigIkeDhGroupEnum       `json:"dh_group,omitempty"`
    // enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`
    EncAlgo              Optional[TunnelConfigEncAlgoEnum] `json:"enc_algo"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigIkeProposal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigIkeProposal) String() string {
    return fmt.Sprintf(
    	"TunnelConfigIkeProposal[AuthAlgo=%v, DhGroup=%v, EncAlgo=%v, AdditionalProperties=%v]",
    	t.AuthAlgo, t.DhGroup, t.EncAlgo, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigIkeProposal.
// It customizes the JSON marshaling process for TunnelConfigIkeProposal objects.
func (t TunnelConfigIkeProposal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "auth_algo", "dh_group", "enc_algo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigIkeProposal object to a map representation for JSON marshaling.
func (t TunnelConfigIkeProposal) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AuthAlgo != nil {
        structMap["auth_algo"] = t.AuthAlgo
    }
    if t.DhGroup != nil {
        structMap["dh_group"] = t.DhGroup
    }
    if t.EncAlgo.IsValueSet() {
        if t.EncAlgo.Value() != nil {
            structMap["enc_algo"] = t.EncAlgo.Value()
        } else {
            structMap["enc_algo"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigIkeProposal.
// It customizes the JSON unmarshaling process for TunnelConfigIkeProposal objects.
func (t *TunnelConfigIkeProposal) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigIkeProposal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "dh_group", "enc_algo")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.AuthAlgo = temp.AuthAlgo
    t.DhGroup = temp.DhGroup
    t.EncAlgo = temp.EncAlgo
    return nil
}

// tempTunnelConfigIkeProposal is a temporary struct used for validating the fields of TunnelConfigIkeProposal.
type tempTunnelConfigIkeProposal  struct {
    AuthAlgo *TunnelConfigAuthAlgoEnum         `json:"auth_algo,omitempty"`
    DhGroup  *TunnelConfigIkeDhGroupEnum       `json:"dh_group,omitempty"`
    EncAlgo  Optional[TunnelConfigEncAlgoEnum] `json:"enc_algo"`
}
