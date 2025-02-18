package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Radsec represents a Radsec struct.
// Radsec settings
type Radsec struct {
    CoaEnabled           *bool                  `json:"coa_enabled,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    IdleTimeout          *int                   `json:"idle_timeout,omitempty"`
    // To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids
    MxclusterIds         []uuid.UUID            `json:"mxcluster_ids,omitempty"`
    // Default is site.mxedge.radsec.proxy_hosts which must be a superset of all `wlans[*].radsec.proxy_hosts`. When `radsec.proxy_hosts` are not used, tunnel peers (org or site mxedges) are used irrespective of `use_site_mxedge`
    ProxyHosts           []string               `json:"proxy_hosts,omitempty"`
    // Name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge.
    ServerName           *string                `json:"server_name,omitempty"`
    // List of Radsec Servers. Only if not Mist Edge.
    Servers              []RadsecServer         `json:"servers,omitempty"`
    // use mxedge(s) as radsecproxy
    UseMxedge            *bool                  `json:"use_mxedge,omitempty"`
    // To use Site mxedges when this WLAN does not use mxtunnel
    UseSiteMxedge        *bool                  `json:"use_site_mxedge,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Radsec,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r Radsec) String() string {
    return fmt.Sprintf(
    	"Radsec[CoaEnabled=%v, Enabled=%v, IdleTimeout=%v, MxclusterIds=%v, ProxyHosts=%v, ServerName=%v, Servers=%v, UseMxedge=%v, UseSiteMxedge=%v, AdditionalProperties=%v]",
    	r.CoaEnabled, r.Enabled, r.IdleTimeout, r.MxclusterIds, r.ProxyHosts, r.ServerName, r.Servers, r.UseMxedge, r.UseSiteMxedge, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Radsec.
// It customizes the JSON marshaling process for Radsec objects.
func (r Radsec) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "coa_enabled", "enabled", "idle_timeout", "mxcluster_ids", "proxy_hosts", "server_name", "servers", "use_mxedge", "use_site_mxedge"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the Radsec object to a map representation for JSON marshaling.
func (r Radsec) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.CoaEnabled != nil {
        structMap["coa_enabled"] = r.CoaEnabled
    }
    if r.Enabled != nil {
        structMap["enabled"] = r.Enabled
    }
    if r.IdleTimeout != nil {
        structMap["idle_timeout"] = r.IdleTimeout
    }
    if r.MxclusterIds != nil {
        structMap["mxcluster_ids"] = r.MxclusterIds
    }
    if r.ProxyHosts != nil {
        structMap["proxy_hosts"] = r.ProxyHosts
    }
    if r.ServerName != nil {
        structMap["server_name"] = r.ServerName
    }
    if r.Servers != nil {
        structMap["servers"] = r.Servers
    }
    if r.UseMxedge != nil {
        structMap["use_mxedge"] = r.UseMxedge
    }
    if r.UseSiteMxedge != nil {
        structMap["use_site_mxedge"] = r.UseSiteMxedge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Radsec.
// It customizes the JSON unmarshaling process for Radsec objects.
func (r *Radsec) UnmarshalJSON(input []byte) error {
    var temp tempRadsec
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coa_enabled", "enabled", "idle_timeout", "mxcluster_ids", "proxy_hosts", "server_name", "servers", "use_mxedge", "use_site_mxedge")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.CoaEnabled = temp.CoaEnabled
    r.Enabled = temp.Enabled
    r.IdleTimeout = temp.IdleTimeout
    r.MxclusterIds = temp.MxclusterIds
    r.ProxyHosts = temp.ProxyHosts
    r.ServerName = temp.ServerName
    r.Servers = temp.Servers
    r.UseMxedge = temp.UseMxedge
    r.UseSiteMxedge = temp.UseSiteMxedge
    return nil
}

// tempRadsec is a temporary struct used for validating the fields of Radsec.
type tempRadsec  struct {
    CoaEnabled    *bool          `json:"coa_enabled,omitempty"`
    Enabled       *bool          `json:"enabled,omitempty"`
    IdleTimeout   *int           `json:"idle_timeout,omitempty"`
    MxclusterIds  []uuid.UUID    `json:"mxcluster_ids,omitempty"`
    ProxyHosts    []string       `json:"proxy_hosts,omitempty"`
    ServerName    *string        `json:"server_name,omitempty"`
    Servers       []RadsecServer `json:"servers,omitempty"`
    UseMxedge     *bool          `json:"use_mxedge,omitempty"`
    UseSiteMxedge *bool          `json:"use_site_mxedge,omitempty"`
}
