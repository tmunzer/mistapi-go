
# Iotendpoint Stats

IoT endpoint statistics

## Structure

`IotendpointStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | MAC address of the AP the endpoint was seen on |
| `Id` | `*string` | Optional | Unique identifier for the IoT endpoint |
| `Lqi` | `*int` | Optional | Link Quality Indicator (0–255)<br><br>**Constraints**: `>= 0`, `<= 255` |
| `Mac` | `*string` | Optional | MAC address of the IoT endpoint |
| `Mfg` | `*string` | Optional | Manufacturer name |
| `Model` | `*string` | Optional | Model name |
| `Timestamp` | `*float64` | Optional | Epoch timestamp of the last observation, in seconds |
| `Type` | `*string` | Optional | IoT endpoint type. enum: `zigbee` |

## Example (as JSON)

```json
{
  "ap_mac": "5c5b350e0001",
  "id": "63f9e299182b63f9",
  "mac": "63f9e299182b63f9",
  "mfg": "Assa Abloy",
  "model": "Assa Abloy",
  "type": "zigbee",
  "lqi": 2
}
```

