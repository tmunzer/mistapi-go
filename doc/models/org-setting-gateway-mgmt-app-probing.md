
# Org Setting Gateway Mgmt App Probing

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtAppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | APp-keys from /api/v1/const/applications |
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

