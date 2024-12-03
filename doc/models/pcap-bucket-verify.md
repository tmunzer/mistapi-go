
# Pcap Bucket Verify

*This model accepts additional fields of type interface{}.*

## Structure

`PcapBucketVerify`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `string` | Required | - |
| `VerifyToken` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bucket": "company-private-pcap",
  "verify_token": "eyJhbGciOiJIUzI1J9.eyJzdWIiOiIxMjM0joiMjgxOG5MDIyfQ.2rzcRvMA3Eg09NnjCAC-1EWMRtxAnFDM",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

