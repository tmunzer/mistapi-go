
# Skyatp List

*This model accepts additional fields of type interface{}.*

## Structure

`SkyatpList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Domains` | [`[]models.SkyatpListDomain`](../../doc/models/skyatp-list-domain.md) | Optional | - |
| `Ips` | [`[]models.SkyatpListIp`](../../doc/models/skyatp-list-ip.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "domains": [
    {
      "comment": "comment6",
      "domain": "domain6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "ips": [
    {
      "comment": "comment0",
      "ip": "ip0",
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

