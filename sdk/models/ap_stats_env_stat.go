package models

import (
    "encoding/json"
)

// ApStatsEnvStat represents a ApStatsEnvStat struct.
// device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
type ApStatsEnvStat struct {
    AccelX               *float64       `json:"accel_x,omitempty"`
    AccelY               *float64       `json:"accel_y,omitempty"`
    AccelZ               *float64       `json:"accel_z,omitempty"`
    AmbientTemp          *int           `json:"ambient_temp,omitempty"`
    Attitude             *int           `json:"attitude,omitempty"`
    CpuTemp              *int           `json:"cpu_temp,omitempty"`
    Humidity             *int           `json:"humidity,omitempty"`
    MagneX               *float64       `json:"magne_x,omitempty"`
    MagneY               *float64       `json:"magne_y,omitempty"`
    MagneZ               *float64       `json:"magne_z,omitempty"`
    Pressure             *int           `json:"pressure,omitempty"`
    VcoreVoltage         *float64       `json:"vcore_voltage,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
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
    if a.AccelX != nil {
        structMap["accel_x"] = a.AccelX
    }
    if a.AccelY != nil {
        structMap["accel_y"] = a.AccelY
    }
    if a.AccelZ != nil {
        structMap["accel_z"] = a.AccelZ
    }
    if a.AmbientTemp != nil {
        structMap["ambient_temp"] = a.AmbientTemp
    }
    if a.Attitude != nil {
        structMap["attitude"] = a.Attitude
    }
    if a.CpuTemp != nil {
        structMap["cpu_temp"] = a.CpuTemp
    }
    if a.Humidity != nil {
        structMap["humidity"] = a.Humidity
    }
    if a.MagneX != nil {
        structMap["magne_x"] = a.MagneX
    }
    if a.MagneY != nil {
        structMap["magne_y"] = a.MagneY
    }
    if a.MagneZ != nil {
        structMap["magne_z"] = a.MagneZ
    }
    if a.Pressure != nil {
        structMap["pressure"] = a.Pressure
    }
    if a.VcoreVoltage != nil {
        structMap["vcore_voltage"] = a.VcoreVoltage
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
    AccelX       *float64 `json:"accel_x,omitempty"`
    AccelY       *float64 `json:"accel_y,omitempty"`
    AccelZ       *float64 `json:"accel_z,omitempty"`
    AmbientTemp  *int     `json:"ambient_temp,omitempty"`
    Attitude     *int     `json:"attitude,omitempty"`
    CpuTemp      *int     `json:"cpu_temp,omitempty"`
    Humidity     *int     `json:"humidity,omitempty"`
    MagneX       *float64 `json:"magne_x,omitempty"`
    MagneY       *float64 `json:"magne_y,omitempty"`
    MagneZ       *float64 `json:"magne_z,omitempty"`
    Pressure     *int     `json:"pressure,omitempty"`
    VcoreVoltage *float64 `json:"vcore_voltage,omitempty"`
}
