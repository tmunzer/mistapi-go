
# Snmpv 3 Config Target Param

SNMPv3 target parameter profile

## Structure

`Snmpv3ConfigTargetParam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MessageProcessingModel` | [`*models.Snmpv3ConfigTargetParamMessProcessModelEnum`](../../doc/models/snmpv-3-config-target-param-mess-process-model-enum.md) | Optional | enum: `v1`, `v2c`, `v3` |
| `Name` | `*string` | Optional | Target parameter profile name |
| `NotifyFilter` | `*string` | Optional | Notification filter profile referenced by this target parameter profile |
| `SecurityLevel` | [`*models.Snmpv3ConfigTargetParamSecurityLevelEnum`](../../doc/models/snmpv-3-config-target-param-security-level-enum.md) | Optional | enum: `authentication`, `none`, `privacy` |
| `SecurityModel` | [`*models.Snmpv3ConfigTargetParamSecurityModelEnum`](../../doc/models/snmpv-3-config-target-param-security-model-enum.md) | Optional | enum: `usm`, `v1`, `v2c` |
| `SecurityName` | `*string` | Optional | USM security name referenced by this target parameter profile |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigTargetParam := models.Snmpv3ConfigTargetParam{
        MessageProcessingModel: models.ToPointer(models.Snmpv3ConfigTargetParamMessProcessModelEnum_V2C),
        Name:                   models.ToPointer("name6"),
        NotifyFilter:           models.ToPointer("notify_filter2"),
        SecurityLevel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityLevelEnum_AUTHENTICATION),
        SecurityModel:          models.ToPointer(models.Snmpv3ConfigTargetParamSecurityModelEnum_V1),
        SecurityName:           models.ToPointer("m01620"),
    }

}
```

