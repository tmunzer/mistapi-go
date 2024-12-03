
# Pcap Bucket

*This model accepts additional fields of type interface{}.*

## Structure

`PcapBucket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bucket": "company-private-pcap",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

