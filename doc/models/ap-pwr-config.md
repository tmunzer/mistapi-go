
# Ap Pwr Config

Power related configs

*This model accepts additional fields of type interface{}.*

## Structure

`ApPwrConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Base` | `*int` | Optional | Additional power to request during negotiating with PSE over PoE, in mW<br><br>**Default**: `0` |
| `PreferUsbOverWifi` | `*bool` | Optional | Whether to enable power out to peripheral, meanwhile will reduce power to Wi-Fi (only for AP45 at power mode)<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "base": 2000,
  "prefer_usb_over_wifi": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

