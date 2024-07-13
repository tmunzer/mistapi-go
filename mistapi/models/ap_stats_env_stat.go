package models

import (
    "encoding/json"
)

// ApStatsEnvStat represents a ApStatsEnvStat struct.
// device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
type ApStatsEnvStat struct {
    AccelX               Optional[float64] `json:"accel_x"`
    AccelY               Optional[float64] `json:"accel_y"`
    AccelZ               Optional[float64] `json:"accel_z"`
    AmbientTemp          Optional[int]     `json:"ambient_temp"`
    Attitude             Optional[int]     `json:"attitude"`
    CpuTemp              Optional[int]     `json:"cpu_temp"`
    Humidity             Optional[int]     `json:"humidity"`
    MagneX               Optional[float64] `json:"magne_x"`
    MagneY               Optional[float64] `json:"magne_y"`
    MagneZ               Optional[float64] `json:"magne_z"`
    Pressure             Optional[float64] `json:"pressure"`
    VcoreVoltage         Optional[int]     `json:"vcore_voltage"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsEnvStat.
// It customizes the JSON marshaling process for ApStatsEnvStat objects.
func (a ApStatsEnvStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsEnvStat object to a map representation for JSON marshaling.
func (a ApStatsEnvStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AccelX.IsValueSet() {
        if a.AccelX.Value() != nil {
            structMap["accel_x"] = a.AccelX.Value()
        } else {
            structMap["accel_x"] = nil
        }
    }
    if a.AccelY.IsValueSet() {
        if a.AccelY.Value() != nil {
            structMap["accel_y"] = a.AccelY.Value()
        } else {
            structMap["accel_y"] = nil
        }
    }
    if a.AccelZ.IsValueSet() {
        if a.AccelZ.Value() != nil {
            structMap["accel_z"] = a.AccelZ.Value()
        } else {
            structMap["accel_z"] = nil
        }
    }
    if a.AmbientTemp.IsValueSet() {
        if a.AmbientTemp.Value() != nil {
            structMap["ambient_temp"] = a.AmbientTemp.Value()
        } else {
            structMap["ambient_temp"] = nil
        }
    }
    if a.Attitude.IsValueSet() {
        if a.Attitude.Value() != nil {
            structMap["attitude"] = a.Attitude.Value()
        } else {
            structMap["attitude"] = nil
        }
    }
    if a.CpuTemp.IsValueSet() {
        if a.CpuTemp.Value() != nil {
            structMap["cpu_temp"] = a.CpuTemp.Value()
        } else {
            structMap["cpu_temp"] = nil
        }
    }
    if a.Humidity.IsValueSet() {
        if a.Humidity.Value() != nil {
            structMap["humidity"] = a.Humidity.Value()
        } else {
            structMap["humidity"] = nil
        }
    }
    if a.MagneX.IsValueSet() {
        if a.MagneX.Value() != nil {
            structMap["magne_x"] = a.MagneX.Value()
        } else {
            structMap["magne_x"] = nil
        }
    }
    if a.MagneY.IsValueSet() {
        if a.MagneY.Value() != nil {
            structMap["magne_y"] = a.MagneY.Value()
        } else {
            structMap["magne_y"] = nil
        }
    }
    if a.MagneZ.IsValueSet() {
        if a.MagneZ.Value() != nil {
            structMap["magne_z"] = a.MagneZ.Value()
        } else {
            structMap["magne_z"] = nil
        }
    }
    if a.Pressure.IsValueSet() {
        if a.Pressure.Value() != nil {
            structMap["pressure"] = a.Pressure.Value()
        } else {
            structMap["pressure"] = nil
        }
    }
    if a.VcoreVoltage.IsValueSet() {
        if a.VcoreVoltage.Value() != nil {
            structMap["vcore_voltage"] = a.VcoreVoltage.Value()
        } else {
            structMap["vcore_voltage"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsEnvStat.
// It customizes the JSON unmarshaling process for ApStatsEnvStat objects.
func (a *ApStatsEnvStat) UnmarshalJSON(input []byte) error {
    var temp apStatsEnvStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accel_x", "accel_y", "accel_z", "ambient_temp", "attitude", "cpu_temp", "humidity", "magne_x", "magne_y", "magne_z", "pressure", "vcore_voltage")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AccelX = temp.AccelX
    a.AccelY = temp.AccelY
    a.AccelZ = temp.AccelZ
    a.AmbientTemp = temp.AmbientTemp
    a.Attitude = temp.Attitude
    a.CpuTemp = temp.CpuTemp
    a.Humidity = temp.Humidity
    a.MagneX = temp.MagneX
    a.MagneY = temp.MagneY
    a.MagneZ = temp.MagneZ
    a.Pressure = temp.Pressure
    a.VcoreVoltage = temp.VcoreVoltage
    return nil
}

// apStatsEnvStat is a temporary struct used for validating the fields of ApStatsEnvStat.
type apStatsEnvStat  struct {
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
