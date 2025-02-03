package models

import (
    "encoding/json"
    "fmt"
)

// MapWallPath represents a MapWallPath struct.
// JSON blob for wall definition (same format as wayfinding_path)
type MapWallPath struct {
    Coordinate           *string                `json:"coordinate,omitempty"`
    Nodes                []MapNode              `json:"nodes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapWallPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapWallPath) String() string {
    return fmt.Sprintf(
    	"MapWallPath[Coordinate=%v, Nodes=%v, AdditionalProperties=%v]",
    	m.Coordinate, m.Nodes, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapWallPath.
// It customizes the JSON marshaling process for MapWallPath objects.
func (m MapWallPath) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "coordinate", "nodes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapWallPath object to a map representation for JSON marshaling.
func (m MapWallPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Coordinate != nil {
        structMap["coordinate"] = m.Coordinate
    }
    if m.Nodes != nil {
        structMap["nodes"] = m.Nodes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapWallPath.
// It customizes the JSON unmarshaling process for MapWallPath objects.
func (m *MapWallPath) UnmarshalJSON(input []byte) error {
    var temp tempMapWallPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coordinate", "nodes")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Coordinate = temp.Coordinate
    m.Nodes = temp.Nodes
    return nil
}

// tempMapWallPath is a temporary struct used for validating the fields of MapWallPath.
type tempMapWallPath  struct {
    Coordinate *string   `json:"coordinate,omitempty"`
    Nodes      []MapNode `json:"nodes,omitempty"`
}
