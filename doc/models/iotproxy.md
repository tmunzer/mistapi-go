
# Iotproxy

IoT proxy configuration for the site

## Structure

`Iotproxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Visionline` | [`*models.IotproxyVisionline`](../../doc/models/iotproxy-visionline.md) | Optional | Visionline integration settings for IoT proxy |

## Example (as JSON)

```json
{
  "enabled": false,
  "visionline": {
    "access_id": "access_id8",
    "enabled": false,
    "host": "host4",
    "password": "password8",
    "port": 8
  }
}
```

