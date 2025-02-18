
# Ap Uplink Port Config

AP Uplink port configuration

*This model accepts additional fields of type interface{}.*

## Structure

`ApUplinkPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dot1x` | `*bool` | Optional | Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch<br>**Default**: `false` |
| `KeepWlansUpIfDown` | `*bool` | Optional | By default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dot1x": false,
  "keep_wlans_up_if_down": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

