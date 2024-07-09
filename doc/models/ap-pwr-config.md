
# Ap Pwr Config

power related configs

## Structure

`ApPwrConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Base` | `*int` | Optional | additional power to request during negotiating with PSE over PoE, in mW<br>**Default**: `0` |
| `PreferUsbOverWifi` | `*bool` | Optional | whether to enable power out to peripheral, meanwhile will reduce power to wifi (only for AP45 at power mode)<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "base": 2000,
  "prefer_usb_over_wifi": false
}
```

