
# Iotproxy Visionline

Visionline integration settings for IoT proxy

## Structure

`IotproxyVisionline`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessId` | `*string` | Optional | Access ID for the Visionline service |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Host` | `*string` | Optional | Hostname or IP of the Visionline collector |
| `Password` | `*string` | Optional | Password for the Visionline service |
| `Port` | `*int` | Optional | TCP port of the Visionline collector<br><br>**Default**: `443` |
| `Username` | `*string` | Optional | Username for the Visionline service |

## Example (as JSON)

```json
{
  "access_id": "790e6c1790e6c18541d",
  "enabled": false,
  "host": "visionline_collector1.local",
  "port": 443,
  "username": "card_administrator",
  "password": "password2"
}
```

