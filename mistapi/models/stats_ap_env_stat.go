// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsApEnvStat represents a StatsApEnvStat struct.
// Device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
type StatsApEnvStat struct {
    AccelX               Optional[float64]      `json:"accel_x"`
    AccelY               Optional[float64]      `json:"accel_y"`
    AccelZ               Optional[float64]      `json:"accel_z"`
    AmbientTemp          Optional[int]          `json:"ambient_temp"`
    Attitude             Optional[int]          `json:"attitude"`
    CpuTemp              Optional[int]          `json:"cpu_temp"`
    Humidity             Optional[int]          `json:"humidity"`
    MagneX               Optional[float64]      `json:"magne_x"`
    MagneY               Optional[float64]      `json:"magne_y"`
    MagneZ               Optional[float64]      `json:"magne_z"`
    Pressure             Optional[float64]      `json:"pressure"`
    VcoreVoltage         Optional[int]          `json:"vcore_voltage"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApEnvStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApEnvStat) String() string {
    return fmt.Sprintf(
    	"StatsApEnvStat[AccelX=%v, AccelY=%v, AccelZ=%v, AmbientTemp=%v, Attitude=%v, CpuTemp=%v, Humidity=%v, MagneX=%v, MagneY=%v, MagneZ=%v, Pressure=%v, VcoreVoltage=%v, AdditionalProperties=%v]",
    	s.AccelX, s.AccelY, s.AccelZ, s.AmbientTemp, s.Attitude, s.CpuTemp, s.Humidity, s.MagneX, s.MagneY, s.MagneZ, s.Pressure, s.VcoreVoltage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApEnvStat.
// It customizes the JSON marshaling process for StatsApEnvStat objects.
func (s StatsApEnvStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "accel_x", "accel_y", "accel_z", "ambient_temp", "attitude", "cpu_temp", "humidity", "magne_x", "magne_y", "magne_z", "pressure", "vcore_voltage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApEnvStat object to a map representation for JSON marshaling.
func (s StatsApEnvStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AccelX.IsValueSet() {
        if s.AccelX.Value() != nil {
            structMap["accel_x"] = s.AccelX.Value()
        } else {
            structMap["accel_x"] = nil
        }
    }
    if s.AccelY.IsValueSet() {
        if s.AccelY.Value() != nil {
            structMap["accel_y"] = s.AccelY.Value()
        } else {
            structMap["accel_y"] = nil
        }
    }
    if s.AccelZ.IsValueSet() {
        if s.AccelZ.Value() != nil {
            structMap["accel_z"] = s.AccelZ.Value()
        } else {
            structMap["accel_z"] = nil
        }
    }
    if s.AmbientTemp.IsValueSet() {
        if s.AmbientTemp.Value() != nil {
            structMap["ambient_temp"] = s.AmbientTemp.Value()
        } else {
            structMap["ambient_temp"] = nil
        }
    }
    if s.Attitude.IsValueSet() {
        if s.Attitude.Value() != nil {
            structMap["attitude"] = s.Attitude.Value()
        } else {
            structMap["attitude"] = nil
        }
    }
    if s.CpuTemp.IsValueSet() {
        if s.CpuTemp.Value() != nil {
            structMap["cpu_temp"] = s.CpuTemp.Value()
        } else {
            structMap["cpu_temp"] = nil
        }
    }
    if s.Humidity.IsValueSet() {
        if s.Humidity.Value() != nil {
            structMap["humidity"] = s.Humidity.Value()
        } else {
            structMap["humidity"] = nil
        }
    }
    if s.MagneX.IsValueSet() {
        if s.MagneX.Value() != nil {
            structMap["magne_x"] = s.MagneX.Value()
        } else {
            structMap["magne_x"] = nil
        }
    }
    if s.MagneY.IsValueSet() {
        if s.MagneY.Value() != nil {
            structMap["magne_y"] = s.MagneY.Value()
        } else {
            structMap["magne_y"] = nil
        }
    }
    if s.MagneZ.IsValueSet() {
        if s.MagneZ.Value() != nil {
            structMap["magne_z"] = s.MagneZ.Value()
        } else {
            structMap["magne_z"] = nil
        }
    }
    if s.Pressure.IsValueSet() {
        if s.Pressure.Value() != nil {
            structMap["pressure"] = s.Pressure.Value()
        } else {
            structMap["pressure"] = nil
        }
    }
    if s.VcoreVoltage.IsValueSet() {
        if s.VcoreVoltage.Value() != nil {
            structMap["vcore_voltage"] = s.VcoreVoltage.Value()
        } else {
            structMap["vcore_voltage"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApEnvStat.
// It customizes the JSON unmarshaling process for StatsApEnvStat objects.
func (s *StatsApEnvStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsApEnvStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accel_x", "accel_y", "accel_z", "ambient_temp", "attitude", "cpu_temp", "humidity", "magne_x", "magne_y", "magne_z", "pressure", "vcore_voltage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AccelX = temp.AccelX
    s.AccelY = temp.AccelY
    s.AccelZ = temp.AccelZ
    s.AmbientTemp = temp.AmbientTemp
    s.Attitude = temp.Attitude
    s.CpuTemp = temp.CpuTemp
    s.Humidity = temp.Humidity
    s.MagneX = temp.MagneX
    s.MagneY = temp.MagneY
    s.MagneZ = temp.MagneZ
    s.Pressure = temp.Pressure
    s.VcoreVoltage = temp.VcoreVoltage
    return nil
}

// tempStatsApEnvStat is a temporary struct used for validating the fields of StatsApEnvStat.
type tempStatsApEnvStat  struct {
    AccelX       Optional[float64] `json:"accel_x"`
    AccelY       Optional[float64] `json:"accel_y"`
    AccelZ       Optional[float64] `json:"accel_z"`
    AmbientTemp  Optional[int]     `json:"ambient_temp"`
    Attitude     Optional[int]     `json:"attitude"`
    CpuTemp      Optional[int]     `json:"cpu_temp"`
    Humidity     Optional[int]     `json:"humidity"`
    MagneX       Optional[float64] `json:"magne_x"`
    MagneY       Optional[float64] `json:"magne_y"`
    MagneZ       Optional[float64] `json:"magne_z"`
    Pressure     Optional[float64] `json:"pressure"`
    VcoreVoltage Optional[int]     `json:"vcore_voltage"`
}
