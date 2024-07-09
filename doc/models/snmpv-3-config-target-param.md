
# Snmpv 3 Config Target Param

## Structure

`Snmpv3ConfigTargetParam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MessageProcessingModel` | [`*models.Snmpv3ConfigTargetParamMessProcessModelEnum`](../../doc/models/snmpv-3-config-target-param-mess-process-model-enum.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `NotifyFilter` | `*string` | Optional | refer to profile-name in notify_filter |
| `SecurityLevel` | [`*models.Snmpv3ConfigTargetParamSecurityLevelEnum`](../../doc/models/snmpv-3-config-target-param-security-level-enum.md) | Optional | - |
| `SecurityModel` | [`*models.Snmpv3ConfigTargetParamSecurityModelEnum`](../../doc/models/snmpv-3-config-target-param-security-model-enum.md) | Optional | - |
| `SecurityName` | `*string` | Optional | refer to security_name in usm |

## Example (as JSON)

```json
{
  "security_name": "m01620",
  "message_processing_model": "v2c",
  "name": "name8",
  "notify_filter": "notify_filter0",
  "security_level": "privacy",
  "security_model": "usm"
}
```

