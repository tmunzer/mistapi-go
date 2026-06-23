
# Gateway Compliance Version

Version compliance metric for gateway devices

## Structure

`GatewayComplianceVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersion` | [`map[string]models.GatewayComplianceMajorVersionProperties`](../../doc/models/gateway-compliance-major-version-properties.md) | Optional | Gateway version-compliance entries keyed by gateway model |
| `Score` | `*float64` | Optional | Gateway software version compliance score |
| `Type` | `*string` | Optional | Device type represented by the gateway compliance metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayComplianceVersion := models.GatewayComplianceVersion{
        MajorVersion:         map[string]models.GatewayComplianceMajorVersionProperties{
            "key0": models.GatewayComplianceMajorVersionProperties{
                MajorCount:           models.ToPointer(80),
                MajorVersion:         models.ToPointer("major_version0"),
            },
        },
        Score:                models.ToPointer(float64(99.9)),
        Type:                 models.ToPointer("gateway"),
    }

}
```

