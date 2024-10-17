package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// EvpnTopology represents a EvpnTopology struct.
type EvpnTopology struct {
	// Switch settings
	Config    *SwitchMgmt `json:"config,omitempty"`
	Id        *uuid.UUID  `json:"id,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Overwrite *bool       `json:"overwrite,omitempty"`
	// Property key is the pod number
	PodNames             map[string]string    `json:"pod_names,omitempty"`
	Switches             []EvpnTopologySwitch `json:"switches"`
	AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopology.
// It customizes the JSON marshaling process for EvpnTopology objects.
func (e EvpnTopology) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopology object to a map representation for JSON marshaling.
func (e EvpnTopology) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, e.AdditionalProperties)
	if e.Config != nil {
		structMap["config"] = e.Config.toMap()
	}
	if e.Id != nil {
		structMap["id"] = e.Id
	}
	if e.Name != nil {
		structMap["name"] = e.Name
	}
	if e.Overwrite != nil {
		structMap["overwrite"] = e.Overwrite
	}
	if e.PodNames != nil {
		structMap["pod_names"] = e.PodNames
	}
	structMap["switches"] = e.Switches
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopology.
// It customizes the JSON unmarshaling process for EvpnTopology objects.
func (e *EvpnTopology) UnmarshalJSON(input []byte) error {
	var temp tempEvpnTopology
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "config", "id", "name", "overwrite", "pod_names", "switches")
	if err != nil {
		return err
	}

	e.AdditionalProperties = additionalProperties
	e.Config = temp.Config
	e.Id = temp.Id
	e.Name = temp.Name
	e.Overwrite = temp.Overwrite
	e.PodNames = temp.PodNames
	e.Switches = *temp.Switches
	return nil
}

// tempEvpnTopology is a temporary struct used for validating the fields of EvpnTopology.
type tempEvpnTopology struct {
	Config    *SwitchMgmt           `json:"config,omitempty"`
	Id        *uuid.UUID            `json:"id,omitempty"`
	Name      *string               `json:"name,omitempty"`
	Overwrite *bool                 `json:"overwrite,omitempty"`
	PodNames  map[string]string     `json:"pod_names,omitempty"`
	Switches  *[]EvpnTopologySwitch `json:"switches"`
}

func (e *tempEvpnTopology) validate() error {
	var errs []string
	if e.Switches == nil {
		errs = append(errs, "required field `switches` is missing for type `evpn_topology`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
