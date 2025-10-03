
# Org Setting Gateway Mgmt App Probing

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtAppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | APp-keys from [List Applications](/#operations/listApplications) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "apps": [
    "facebook"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

