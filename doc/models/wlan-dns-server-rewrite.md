
# Wlan Dns Server Rewrite

for radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns)

## Structure

`WlanDnsServerRewrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `RadiusGroups` | `map[string]string` | Optional | map between radius_group and the desired DNS server (IPv4 only)<br>Property key is the RADIUS group, property value is the desired DNS Server |

## Example (as JSON)

```json
{
  "enabled": false,
  "radius_groups": {
    "contractor": "172.1.1.1",
    "guest": "8.8.8.8"
  }
}
```

