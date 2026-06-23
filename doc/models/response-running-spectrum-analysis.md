
# Response Running Spectrum Analysis

Running spectrum analysis session for a site

## Structure

`ResponseRunningSpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band currently being scanned by spectrum analysis, such as 24, 5, or 6 |
| `Channels` | `[]int` | Optional | List of channels being scanned in the spectrum analysis |
| `DeviceId` | `*uuid.UUID` | Optional | Device ID of the AP that is running spectrum analysis |
| `Duration` | `*int` | Optional | Length of the running spectrum analysis session, in seconds |
| `Format` | `*string` | Optional | Output format for the running spectrum analysis data, such as json or stream |
| `StartedTime` | `*int` | Optional | Timestamp when the spectrum analysis was started |
| `Width` | `*int` | Optional | Channel width used during the spectrum analysis scan, in MHz |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseRunningSpectrumAnalysis := models.ResponseRunningSpectrumAnalysis{
        Band:                 models.ToPointer("band4"),
        Channels:             []int{
            193,
            194,
            195,
        },
        DeviceId:             models.ToPointer(uuid.MustParse("00000530-0000-0000-0000-000000000000")),
        Duration:             models.ToPointer(250),
        Format:               models.ToPointer("format8"),
    }

}
```

