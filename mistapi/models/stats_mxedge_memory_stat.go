// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeMemoryStat represents a StatsMxedgeMemoryStat struct.
// Memory usage
type StatsMxedgeMemoryStat struct {
    // The amount of memory, in kilobytes, that has been used more recently and is usually not reclaimed unless absolutely necessary.
    Active               *int                   `json:"active,omitempty"`
    // An estimate of how much memory is available for starting new applications, without swapping.
    Available            *int64                 `json:"available,omitempty"`
    // The amount, in kilobytes, of temporary storage for raw disk blocks.
    Buffers              *int                   `json:"buffers,omitempty"`
    // The amount of physical RAM, in kilobytes, used as cache memory.
    Cached               *int                   `json:"cached,omitempty"`
    // The amount of physical RAM, in kilobytes, left unused by the system
    Free                 *int64                 `json:"free,omitempty"`
    // The amount of memory, in kilobytes, that has been used less recently and is more eligible to be reclaimed for other purposes.
    Inactive             *int                   `json:"inactive,omitempty"`
    // The amount of memory, in kilobytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile.
    SwapCached           *int                   `json:"swap_cached,omitempty"`
    // The total amount of swap free, in kilobytes.
    SwapFree             *int                   `json:"swap_free,omitempty"`
    // The total amount of swap available, in kilobytes.
    SwapTotal            *int                   `json:"swap_total,omitempty"`
    // Total amount of usable RAM, in kilobytes, which is physical RAM minus a number of reserved bits and the kernel binary code
    Total                *int64                 `json:"total,omitempty"`
    Usage                *int                   `json:"usage,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeMemoryStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeMemoryStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeMemoryStat[Active=%v, Available=%v, Buffers=%v, Cached=%v, Free=%v, Inactive=%v, SwapCached=%v, SwapFree=%v, SwapTotal=%v, Total=%v, Usage=%v, AdditionalProperties=%v]",
    	s.Active, s.Available, s.Buffers, s.Cached, s.Free, s.Inactive, s.SwapCached, s.SwapFree, s.SwapTotal, s.Total, s.Usage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeMemoryStat.
// It customizes the JSON marshaling process for StatsMxedgeMemoryStat objects.
func (s StatsMxedgeMemoryStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "active", "available", "buffers", "cached", "free", "inactive", "swap_cached", "swap_free", "swap_total", "total", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeMemoryStat object to a map representation for JSON marshaling.
func (s StatsMxedgeMemoryStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Active != nil {
        structMap["active"] = s.Active
    }
    if s.Available != nil {
        structMap["available"] = s.Available
    }
    if s.Buffers != nil {
        structMap["buffers"] = s.Buffers
    }
    if s.Cached != nil {
        structMap["cached"] = s.Cached
    }
    if s.Free != nil {
        structMap["free"] = s.Free
    }
    if s.Inactive != nil {
        structMap["inactive"] = s.Inactive
    }
    if s.SwapCached != nil {
        structMap["swap_cached"] = s.SwapCached
    }
    if s.SwapFree != nil {
        structMap["swap_free"] = s.SwapFree
    }
    if s.SwapTotal != nil {
        structMap["swap_total"] = s.SwapTotal
    }
    if s.Total != nil {
        structMap["total"] = s.Total
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeMemoryStat.
// It customizes the JSON unmarshaling process for StatsMxedgeMemoryStat objects.
func (s *StatsMxedgeMemoryStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeMemoryStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active", "available", "buffers", "cached", "free", "inactive", "swap_cached", "swap_free", "swap_total", "total", "usage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Active = temp.Active
    s.Available = temp.Available
    s.Buffers = temp.Buffers
    s.Cached = temp.Cached
    s.Free = temp.Free
    s.Inactive = temp.Inactive
    s.SwapCached = temp.SwapCached
    s.SwapFree = temp.SwapFree
    s.SwapTotal = temp.SwapTotal
    s.Total = temp.Total
    s.Usage = temp.Usage
    return nil
}

// tempStatsMxedgeMemoryStat is a temporary struct used for validating the fields of StatsMxedgeMemoryStat.
type tempStatsMxedgeMemoryStat  struct {
    Active     *int   `json:"active,omitempty"`
    Available  *int64 `json:"available,omitempty"`
    Buffers    *int   `json:"buffers,omitempty"`
    Cached     *int   `json:"cached,omitempty"`
    Free       *int64 `json:"free,omitempty"`
    Inactive   *int   `json:"inactive,omitempty"`
    SwapCached *int   `json:"swap_cached,omitempty"`
    SwapFree   *int   `json:"swap_free,omitempty"`
    SwapTotal  *int   `json:"swap_total,omitempty"`
    Total      *int64 `json:"total,omitempty"`
    Usage      *int   `json:"usage,omitempty"`
}
