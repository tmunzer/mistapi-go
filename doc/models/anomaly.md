
# Anomaly

SLE anomaly data point with baseline, deviation, and contributing events

## Structure

`Anomaly`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | `[]string` | Required, Read-only | Event names associated with an anomaly record |
| `Since` | `*float64` | Optional, Read-only | Timestamp when the anomalous period began |
| `SleBaseline` | `float64` | Required, Read-only | Expected SLE value used as the anomaly baseline |
| `SleDeviation` | `float64` | Required, Read-only | Difference between the observed SLE value and the baseline |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    anomaly := models.Anomaly{
        Events:               "",
        SleBaseline:          0.0,
        SleDeviation:         0.0,
        Timestamp:            0.0,
    }

}
```

