
# Wlan App Qos

app qos wlan settings

## Structure

`WlanAppQos`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | [`map[string]models.WlanAppQosAppsProperties`](../../doc/models/wlan-app-qos-apps-properties.md) | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Others` | [`[]models.WlanAppQosOthersItem`](../../doc/models/wlan-app-qos-others-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "apps": {
    "skype-business-video": {
      "dscp": 32,
      "dst_subnet": "10.2.0.0/16",
      "src_subnet": "10.2.0.0/16"
    }
  },
  "enabled": false,
  "others": [
    {
      "dscp": 44,
      "dst_subnet": "dst_subnet2",
      "port_ranges": "port_ranges6",
      "protocol": "protocol2",
      "src_subnet": "src_subnet0"
    },
    {
      "dscp": 44,
      "dst_subnet": "dst_subnet2",
      "port_ranges": "port_ranges6",
      "protocol": "protocol2",
      "src_subnet": "src_subnet0"
    }
  ]
}
```

