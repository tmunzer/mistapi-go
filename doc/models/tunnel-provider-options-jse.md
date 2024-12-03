
# Tunnel Provider Options Jse

for jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsJse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `NumUsers` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "JSE_ORG1",
  "num_users": 5,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

