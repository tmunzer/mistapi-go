
# Gateway Metrics

Gateway metric scores returned by the site gateway metrics API

## Structure

`GatewayMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigSuccess` | `*float64` | Optional | Gateway configuration success score |
| `VersionCompliance` | [`*models.GatewayComplianceVersion`](../../doc/models/gateway-compliance-version.md) | Optional | Version compliance metric for gateway devices |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMetrics := models.GatewayMetrics{
        ConfigSuccess:        models.ToPointer(float64(99.9)),
        VersionCompliance:    models.ToPointer(models.GatewayComplianceVersion{
            MajorVersion:         map[string]models.GatewayComplianceMajorVersionProperties{
                "key0": models.GatewayComplianceMajorVersionProperties{
                    MajorCount:           models.ToPointer(80),
                    MajorVersion:         models.ToPointer("major_version0"),
                },
                "key1": models.GatewayComplianceMajorVersionProperties{
                    MajorCount:           models.ToPointer(80),
                    MajorVersion:         models.ToPointer("major_version0"),
                },
                "key2": models.GatewayComplianceMajorVersionProperties{
                    MajorCount:           models.ToPointer(80),
                    MajorVersion:         models.ToPointer("major_version0"),
                },
            },
            Score:                models.ToPointer(float64(149.42)),
            Type:                 models.ToPointer("type2"),
        }),
    }

}
```

