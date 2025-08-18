
# Issued Client Certificates Results

*This model accepts additional fields of type interface{}.*

## Structure

`IssuedClientCertificatesResults`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.IssuedClientCertificate`](../../doc/models/issued-client-certificate.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "common_name": "common_name4",
      "created_time": 73.76,
      "device_id": "00001510-0000-0000-0000-000000000000",
      "modified_time": 5.2,
      "serial_number": "serial_number0",
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

