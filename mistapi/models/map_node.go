package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MapNode represents a MapNode struct.
// Nodes on maps
type MapNode struct {
    Edges                map[string]string      `json:"edges,omitempty"`
    Name                 string                 `json:"name"`
    Position             *MapNodePosition       `json:"position,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapNode) String() string {
    return fmt.Sprintf(
    	"MapNode[Edges=%v, Name=%v, Position=%v, AdditionalProperties=%v]",
    	m.Edges, m.Name, m.Position, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapNode.
// It customizes the JSON marshaling process for MapNode objects.
func (m MapNode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "edges", "name", "position"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapNode object to a map representation for JSON marshaling.
func (m MapNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMapNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "edges", "name", "position")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Edges = temp.Edges
    m.Name = *temp.Name
    m.Position = temp.Position
    return nil
}

// tempMapNode is a temporary struct used for validating the fields of MapNode.
type tempMapNode  struct {
    Edges    map[string]string `json:"edges,omitempty"`
    Name     *string           `json:"name"`
    Position *MapNodePosition  `json:"position,omitempty"`
}

func (m *tempMapNode) validate() error {
    var errs []string
    if m.Name == nil {
        errs = append(errs, "required field `name` is missing for type `map_node`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
