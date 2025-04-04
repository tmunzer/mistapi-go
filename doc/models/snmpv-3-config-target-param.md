
# Snmpv 3 Config Target Param

*This model accepts additional fields of type interface{}.*

## Structure

`Snmpv3ConfigTargetParam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MessageProcessingModel` | [`*models.Snmpv3ConfigTargetParamMessProcessModelEnum`](../../doc/models/snmpv-3-config-target-param-mess-process-model-enum.md) | Optional | enum: `v1`, `v2c`, `v3` |
| `Name` | `*string` | Optional | - |
| `NotifyFilter` | `*string` | Optional | Refer to profile-name in notify_filter |
| `SecurityLevel` | [`*models.Snmpv3ConfigTargetParamSecurityLevelEnum`](../../doc/models/snmpv-3-config-target-param-security-level-enum.md) | Optional | enum: `authentication`, `none`, `privacy` |
| `SecurityModel` | [`*models.Snmpv3ConfigTargetParamSecurityModelEnum`](../../doc/models/snmpv-3-config-target-param-security-model-enum.md) | Optional | enum: `usm`, `v1`, `v2c` |
| `SecurityName` | `*string` | Optional | Refer to security_name in usm |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "security_name": "m01620",
  "message_processing_model": "v1",
  "name": "name0",
  "notify_filter": "notify_filter2",
  "security_level": "privacy",
  "security_model": "v2c",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

