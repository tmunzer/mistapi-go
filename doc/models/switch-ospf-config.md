
# Switch Ospf Config

OSPF routing configuration for a Junos switch

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchOspfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Areas` | [`map[string]models.SwitchOspfConfigArea`](../../doc/models/switch-ospf-config-area.md) | Optional | Property key is the area name. Defines the OSPF areas configured on the switch. |
| `Enabled` | `*bool` | Optional | Enable OSPF on the switch<br><br>**Default**: `false` |
| `ExportPolicy` | `*string` | Optional | optional, for basic scenario, `import_policy` can be specified and can be applied to all networks in all areas if not explicitly specified |
| `ImportPolicy` | `*string` | Optional | optional, for basic scenario, `import_policy` can be specified and can be applied to all networks in all areas if not explicitly specified |
| `ReferenceBandwidth` | [`*models.SwitchOspfConfigReferenceBandwidth`](../../doc/models/containers/switch-ospf-config-reference-bandwidth.md) | Optional | Reference bandwidth. Integer(100000) or String (10g) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchOspfConfig := models.SwitchOspfConfig{
        Areas:                map[string]models.SwitchOspfConfigArea{
            "key0": models.SwitchOspfConfigArea{
                NoSummary:            models.ToPointer(false),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            "key1": models.SwitchOspfConfigArea{
                NoSummary:            models.ToPointer(false),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            "key2": models.SwitchOspfConfigArea{
                NoSummary:            models.ToPointer(false),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        Enabled:              models.ToPointer(false),
        ExportPolicy:         models.ToPointer("export_policy0"),
        ImportPolicy:         models.ToPointer("import_policy4"),
        ReferenceBandwidth:   models.ToPointer(models.SwitchOspfConfigReferenceBandwidthContainer.FromNumber(100000)),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

