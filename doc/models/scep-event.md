
# Scep Event

Mist SCEP PKI operation event reported for an organization

## Structure

`ScepEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProvider` | `*string` | Optional | MDM or certificate provider that triggered the SCEP operation |
| `CommonName` | `*string` | Optional | Common name presented in the SCEP certificate request |
| `DeviceId` | `*string` | Optional | Device identifier associated with the SCEP operation. Empty when the SCEP request did not include a device ID |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Text` | `*string` | Optional | Reason text describing the outcome of the SCEP operation |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | [`*models.ScepEventTypeEnum`](../../doc/models/scep-event-type-enum.md) | Optional | enum: `SCEP_PKI_OPERATION_FAILURE`, `SCEP_PKI_OPERATION_SUCCESS` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    scepEvent := models.ScepEvent{
        CertProvider:         models.ToPointer("jamf"),
        CommonName:           models.ToPointer("name@company.net bb08e3c5-a1d9-5f21-a3b7-cd0821eab8f6"),
        DeviceId:             models.ToPointer("bb08e3c5-a1d9-5f21-a3b7-cd0821eab8f6"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Text:                 models.ToPointer("invalid challenge/expired"),
        Type:                 models.ToPointer(models.ScepEventTypeEnum_SCEPPKIOPERATIONFAILURE),
    }

}
```

