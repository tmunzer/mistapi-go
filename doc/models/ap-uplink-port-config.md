
# Ap Uplink Port Config

AP Uplink port configuration

## Structure

`ApUplinkPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dot1x` | `*bool` | Optional | Whether to do 802.1x against uplink switch. When enabled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch<br><br>**Default**: `false` |
| `KeepWlansUpIfDown` | `*bool` | Optional | By default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "dot1x": false,
  "keep_wlans_up_if_down": false
}
```

