package models

import (
    "encoding/json"
)

// MxedgeStatsMemoryStat represents a MxedgeStatsMemoryStat struct.
// Memory usage
type MxedgeStatsMemoryStat struct {
    // The amount of memory, in kibibytes, that has been used more recently and is usually not reclaimed unless absolutely necessary.
    Active               *int           `json:"active,omitempty"`
    // An estimate of how much memory is available for starting new applications, without swapping.
    Available            *int64         `json:"available,omitempty"`
    // The amount, in kibibytes, of temporary storage for raw disk blocks.
    Buffers              *int           `json:"buffers,omitempty"`
    // The amount of physical RAM, in kibibytes, used as cache memory.
    Cached               *int           `json:"cached,omitempty"`
    // The amount of physical RAM, in kibibytes, left unused by the system
    Free                 *int64         `json:"free,omitempty"`
    // The amount of memory, in kibibytes, that has been used less recently and is more eligible to be reclaimed for other purposes.
    Inactive             *int           `json:"inactive,omitempty"`
    // The amount of memory, in kibibytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile.
    SwapCached           *int           `json:"swap_cached,omitempty"`
    // The total amount of swap free, in kibibytes.
    SwapFree             *int           `json:"swap_free,omitempty"`
    // The total amount of swap available, in kibibytes.
    SwapTotal            *int           `json:"swap_total,omitempty"`
    // Total amount of usable RAM, in kibibytes, which is physical RAM minus a number of reserved bits and the kernel binary code
    Total                *int64         `json:"total,omitempty"`
    Usage                *int           `json:"usage,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsMemoryStat.
// It customizes the JSON marshaling process for MxedgeStatsMemoryStat objects.
func (m MxedgeStatsMemoryStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsMemoryStat object to a map representation for JSON marshaling.
func (m MxedgeStatsMemoryStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Active != nil {
        structMap["active"] = m.Active
    }
    if m.Available != nil {
        structMap["available"] = m.Available
    }
    if m.Buffers != nil {
        structMap["buffers"] = m.Buffers
    }
    if m.Cached != nil {
        structMap["cached"] = m.Cached
    }
    if m.Free != nil {
        structMap["free"] = m.Free
    }
    if m.Inactive != nil {
        structMap["inactive"] = m.Inactive
    }
    if m.SwapCached != nil {
        structMap["swap_cached"] = m.SwapCached
    }
    if m.SwapFree != nil {
        structMap["swap_free"] = m.SwapFree
    }
    if m.SwapTotal != nil {
        structMap["swap_total"] = m.SwapTotal
    }
    if m.Total != nil {
        structMap["total"] = m.Total
    }
    if m.Usage != nil {
        structMap["usage"] = m.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsMemoryStat.
// It customizes the JSON unmarshaling process for MxedgeStatsMemoryStat objects.
func (m *MxedgeStatsMemoryStat) UnmarshalJSON(input []byte) error {
    var temp mxedgeStatsMemoryStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "active", "available", "buffers", "cached", "free", "inactive", "swap_cached", "swap_free", "swap_total", "total", "usage")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Active = temp.Active
    m.Available = temp.Available
    m.Buffers = temp.Buffers
    m.Cached = temp.Cached
    m.Free = temp.Free
    m.Inactive = temp.Inactive
    m.SwapCached = temp.SwapCached
    m.SwapFree = temp.SwapFree
    m.SwapTotal = temp.SwapTotal
    m.Total = temp.Total
    m.Usage = temp.Usage
    return nil
}

// mxedgeStatsMemoryStat is a temporary struct used for validating the fields of MxedgeStatsMemoryStat.
type mxedgeStatsMemoryStat  struct {
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
