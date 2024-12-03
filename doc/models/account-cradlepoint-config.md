
# Account Cradlepoint Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountCradlepointConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpApiId` | `*string` | Optional | - |
| `CpApiKey` | `*string` | Optional | - |
| `EcmApiId` | `*string` | Optional | - |
| `EcmApiKey` | `*string` | Optional | - |
| `EnableLldp` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cp_api_id": "84446d61-2206-4ea5-855a-0043f980be54",
  "cp_api_key": "79c329da9893e34099c7d8ad5cb9c941",
  "ecm_api_id": "73446d61-2206-4ea5-855a-0043f980be62",
  "ecm_api_key": "68b329da9893e34099c7d8ad5cb9c9405",
  "enable_lldp": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

