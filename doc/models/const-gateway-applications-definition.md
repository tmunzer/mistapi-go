
# Const Gateway Applications Definition

*This model accepts additional fields of type interface{}.*

## Structure

`ConstGatewayApplicationsDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppId` | `*bool` | Optional | - |
| `Key` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `SsrAppId` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "app_id": true,
  "key": "4shared",
  "name": "4shared",
  "ssr_app_id": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

