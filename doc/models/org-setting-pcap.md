
# Org Setting Pcap

## Structure

`OrgSettingPcap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `*string` | Optional | - |
| `MaxPktLen` | `*int` | Optional | max_len of non-management packets to capture<br>**Default**: `128`<br>**Constraints**: `<= 128` |

## Example (as JSON)

```json
{
  "bucket": "myorg_pcap",
  "max_pkt_len": 128
}
```

