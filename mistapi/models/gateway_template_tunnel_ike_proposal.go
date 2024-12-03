package models

import (
    "encoding/json"
)

// GatewayTemplateTunnelIkeProposal represents a GatewayTemplateTunnelIkeProposal struct.
type GatewayTemplateTunnelIkeProposal struct {
    // enum: `md5`, `sha1`, `sha2`
    AuthAlgo             *TunnelConfigsAuthAlgoEnum           `json:"auth_algo,omitempty"`
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
    DhGroup              *GatewayTemplateTunnelIkeDhGroupEnum `json:"dh_group,omitempty"`
    // enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`
    EncAlgo              Optional[TunnelConfigsEncAlgoEnum]   `json:"enc_algo"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayTemplateTunnelIkeProposal.
// It customizes the JSON marshaling process for GatewayTemplateTunnelIkeProposal objects.
func (g GatewayTemplateTunnelIkeProposal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "auth_algo", "dh_group", "enc_algo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayTemplateTunnelIkeProposal object to a map representation for JSON marshaling.
func (g GatewayTemplateTunnelIkeProposal) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AuthAlgo != nil {
        structMap["auth_algo"] = g.AuthAlgo
    }
    if g.DhGroup != nil {
        structMap["dh_group"] = g.DhGroup
    }
    if g.EncAlgo.IsValueSet() {
        if g.EncAlgo.Value() != nil {
            structMap["enc_algo"] = g.EncAlgo.Value()
        } else {
            structMap["enc_algo"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayTemplateTunnelIkeProposal.
// It customizes the JSON unmarshaling process for GatewayTemplateTunnelIkeProposal objects.
func (g *GatewayTemplateTunnelIkeProposal) UnmarshalJSON(input []byte) error {
    var temp tempGatewayTemplateTunnelIkeProposal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "dh_group", "enc_algo")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.AuthAlgo = temp.AuthAlgo
    g.DhGroup = temp.DhGroup
    g.EncAlgo = temp.EncAlgo
    return nil
}

// tempGatewayTemplateTunnelIkeProposal is a temporary struct used for validating the fields of GatewayTemplateTunnelIkeProposal.
type tempGatewayTemplateTunnelIkeProposal  struct {
    AuthAlgo *TunnelConfigsAuthAlgoEnum           `json:"auth_algo,omitempty"`
    DhGroup  *GatewayTemplateTunnelIkeDhGroupEnum `json:"dh_group,omitempty"`
    EncAlgo  Optional[TunnelConfigsEncAlgoEnum]   `json:"enc_algo"`
}
