
# Org Setting Pcap

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingPcap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `*string` | Optional | - |
| `MaxPktLen` | `*int` | Optional | Max_len of non-management packets to capture<br>**Default**: `128`<br>**Constraints**: `<= 128` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bucket": "myorg_pcap",
  "max_pkt_len": 128,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

