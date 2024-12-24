package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// SsoMxedgeProxy represents a SsoMxedgeProxy struct.
// if `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster
type SsoMxedgeProxy struct {
    AcctServers          []SsoMxedgeProxyAcctServer `json:"acct_servers,omitempty"`
    AuthServers          []SsoMxedgeProxyAuthServer `json:"auth_servers,omitempty"`
    MxclusterId          *uuid.UUID                 `json:"mxcluster_id,omitempty"`
    // Operator name as Radius attribute while proxying
    OperatorName         *string                    `json:"operator_name,omitempty"`
    // public hostname/IPs
    ProxyHosts           []string                   `json:"proxy_hosts,omitempty"`
    // SSIDs that support eduroam
    Ssids                []string                   `json:"ssids,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for SsoMxedgeProxy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoMxedgeProxy) String() string {
    return fmt.Sprintf(
    	"SsoMxedgeProxy[AcctServers=%v, AuthServers=%v, MxclusterId=%v, OperatorName=%v, ProxyHosts=%v, Ssids=%v, AdditionalProperties=%v]",
    	s.AcctServers, s.AuthServers, s.MxclusterId, s.OperatorName, s.ProxyHosts, s.Ssids, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoMxedgeProxy.
// It customizes the JSON marshaling process for SsoMxedgeProxy objects.
func (s SsoMxedgeProxy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "acct_servers", "auth_servers", "mxcluster_id", "operator_name", "proxy_hosts", "ssids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SsoMxedgeProxy object to a map representation for JSON marshaling.
func (s SsoMxedgeProxy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AcctServers != nil {
        structMap["acct_servers"] = s.AcctServers
    }
    if s.AuthServers != nil {
        structMap["auth_servers"] = s.AuthServers
    }
    if s.MxclusterId != nil {
        structMap["mxcluster_id"] = s.MxclusterId
    }
    if s.OperatorName != nil {
        structMap["operator_name"] = s.OperatorName
    }
    if s.ProxyHosts != nil {
        structMap["proxy_hosts"] = s.ProxyHosts
    }
    if s.Ssids != nil {
        structMap["ssids"] = s.Ssids
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoMxedgeProxy.
// It customizes the JSON unmarshaling process for SsoMxedgeProxy objects.
func (s *SsoMxedgeProxy) UnmarshalJSON(input []byte) error {
    var temp tempSsoMxedgeProxy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_servers", "auth_servers", "mxcluster_id", "operator_name", "proxy_hosts", "ssids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AcctServers = temp.AcctServers
    s.AuthServers = temp.AuthServers
    s.MxclusterId = temp.MxclusterId
    s.OperatorName = temp.OperatorName
    s.ProxyHosts = temp.ProxyHosts
    s.Ssids = temp.Ssids
    return nil
}

// tempSsoMxedgeProxy is a temporary struct used for validating the fields of SsoMxedgeProxy.
type tempSsoMxedgeProxy  struct {
    AcctServers  []SsoMxedgeProxyAcctServer `json:"acct_servers,omitempty"`
    AuthServers  []SsoMxedgeProxyAuthServer `json:"auth_servers,omitempty"`
    MxclusterId  *uuid.UUID                 `json:"mxcluster_id,omitempty"`
    OperatorName *string                    `json:"operator_name,omitempty"`
    ProxyHosts   []string                   `json:"proxy_hosts,omitempty"`
    Ssids        []string                   `json:"ssids,omitempty"`
}
