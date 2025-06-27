
# Skyatp List

*This model accepts additional fields of type interface{}.*

## Structure

`SkyatpList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Domains` | [`[]models.SkyatpListDomain`](../../doc/models/skyatp-list-domain.md) | Optional | - |
| `Ip` | [`[]models.SkyatpListIp`](../../doc/models/skyatp-list-ip.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "domains": [
    {
      "comment": "comment6",
      "value": "value2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "ip": [
    {
      "comment": "comment8",
      "value": "value6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

