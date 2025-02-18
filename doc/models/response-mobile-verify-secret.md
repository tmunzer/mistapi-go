
# Response Mobile Verify Secret

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMobileVerifySecret`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `Secret` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "secret": "secret2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

