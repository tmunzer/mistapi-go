package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MapNode represents a MapNode struct.
// Nodes on maps
type MapNode struct {
    Edges                map[string]string `json:"edges,omitempty"`
    Name                 string            `json:"name"`
    Position             *MapNodePosition  `json:"position,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapNode.
// It customizes the JSON marshaling process for MapNode objects.
func (m MapNode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapNode object to a map representation for JSON marshaling.
func (m MapNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Edges != nil {
        structMap["edges"] = m.Edges
    }
    structMap["name"] = m.Name
    if m.Position != nil {
        structMap["position"] = m.Position.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapNode.
// It customizes the JSON unmarshaling process for MapNode objects.
func (m *MapNode) UnmarshalJSON(input []byte) error {
    var temp mapNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "edges", "name", "position")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Edges = temp.Edges
    m.Name = *temp.Name
    m.Position = temp.Position
    return nil
}

// mapNode is a temporary struct used for validating the fields of MapNode.
type mapNode  struct {
    Edges    map[string]string `json:"edges,omitempty"`
    Name     *string           `json:"name"`
    Position *MapNodePosition  `json:"position,omitempty"`
}

func (m *mapNode) validate() error {
    var errs []string
    if m.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Map_Node`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
