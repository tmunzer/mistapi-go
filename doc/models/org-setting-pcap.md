
# Org Setting Pcap

## Structure

`OrgSettingPcap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `*string` | Optional | - |
| `MaxPktLen` | `*int` | Optional | Max_len of non-management packets to capture<br><br>**Default**: `128`<br><br>**Constraints**: `<= 128` |

## Example (as JSON)

```json
{
  "bucket": "myorg_pcap",
  "max_pkt_len": 128
}
```

