// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// SearchMxedge represents a SearchMxedge struct.
type SearchMxedge struct {
	Distro   *string  `json:"distro,omitempty"`
	LastSeen *float64 `json:"last_seen,omitempty"`
	Model    *string  `json:"model,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	MxclusterId *uuid.UUID `json:"mxcluster_id,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	MxedgeId *uuid.UUID `json:"mxedge_id,omitempty"`
	// The name of the tunnel
	Name                 *string                `json:"name,omitempty"`
	OrgId                *uuid.UUID             `json:"org_id,omitempty"`
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	TuntermVersion       *string                `json:"tunterm_version,omitempty"`
	Uptime               *int                   `json:"uptime,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchMxedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchMxedge) String() string {
	return fmt.Sprintf(
		"SearchMxedge[Distro=%v, LastSeen=%v, Model=%v, MxclusterId=%v, MxedgeId=%v, Name=%v, OrgId=%v, SiteId=%v, TuntermVersion=%v, Uptime=%v, AdditionalProperties=%v]",
		s.Distro, s.LastSeen, s.Model, s.MxclusterId, s.MxedgeId, s.Name, s.OrgId, s.SiteId, s.TuntermVersion, s.Uptime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchMxedge.
// It customizes the JSON marshaling process for SearchMxedge objects.
func (s SearchMxedge) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"distro", "last_seen", "model", "mxcluster_id", "mxedge_id", "name", "org_id", "site_id", "tunterm_version", "uptime"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SearchMxedge object to a map representation for JSON marshaling.
func (s SearchMxedge) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Distro != nil {
		structMap["distro"] = s.Distro
	}
	if s.LastSeen != nil {
		structMap["last_seen"] = s.LastSeen
	}
	if s.Model != nil {
		structMap["model"] = s.Model
	}
	if s.MxclusterId != nil {
		structMap["mxcluster_id"] = s.MxclusterId
	}
	if s.MxedgeId != nil {
		structMap["mxedge_id"] = s.MxedgeId
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.TuntermVersion != nil {
		structMap["tunterm_version"] = s.TuntermVersion
	}
	if s.Uptime != nil {
		structMap["uptime"] = s.Uptime
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchMxedge.
// It customizes the JSON unmarshaling process for SearchMxedge objects.
func (s *SearchMxedge) UnmarshalJSON(input []byte) error {
	var temp tempSearchMxedge
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "distro", "last_seen", "model", "mxcluster_id", "mxedge_id", "name", "org_id", "site_id", "tunterm_version", "uptime")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Distro = temp.Distro
	s.LastSeen = temp.LastSeen
	s.Model = temp.Model
	s.MxclusterId = temp.MxclusterId
	s.MxedgeId = temp.MxedgeId
	s.Name = temp.Name
	s.OrgId = temp.OrgId
	s.SiteId = temp.SiteId
	s.TuntermVersion = temp.TuntermVersion
	s.Uptime = temp.Uptime
	return nil
}

// tempSearchMxedge is a temporary struct used for validating the fields of SearchMxedge.
type tempSearchMxedge struct {
	Distro         *string    `json:"distro,omitempty"`
	LastSeen       *float64   `json:"last_seen,omitempty"`
	Model          *string    `json:"model,omitempty"`
	MxclusterId    *uuid.UUID `json:"mxcluster_id,omitempty"`
	MxedgeId       *uuid.UUID `json:"mxedge_id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	OrgId          *uuid.UUID `json:"org_id,omitempty"`
	SiteId         *uuid.UUID `json:"site_id,omitempty"`
	TuntermVersion *string    `json:"tunterm_version,omitempty"`
	Uptime         *int       `json:"uptime,omitempty"`
}
