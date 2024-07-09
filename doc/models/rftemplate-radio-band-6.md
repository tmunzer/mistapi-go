
# Rftemplate Radio Band 6

Radio Band AP settings

## Structure

`RftemplateRadioBand6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | **Default**: `false` |
| `AntGain` | `models.Optional[int]` | Optional | **Default**: `0`<br>**Constraints**: `>= 0`, `<= 10` |
| `AntennaMode` | [`*models.RadioBandAntennaModeEnum`](../../doc/models/radio-band-antenna-mode-enum.md) | Optional | **Default**: `"default"` |
| `Bandwidth` | [`*models.Dot11Bandwidth6Enum`](../../doc/models/dot-11-bandwidth-6-enum.md) | Optional | channel width for the 6GHz band<br>**Default**: `80` |
| `Channels` | `models.Optional[[]int]` | Optional | For RFTemplates. List of channels, null or empty array means auto |
| `Disabled` | `*bool` | Optional | whether to disable the radio<br>**Default**: `false` |
| `Power` | `models.Optional[int]` | Optional | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …<br>**Constraints**: `>= 5`, `<= 25` |
| `PowerMax` | `models.Optional[int]` | Optional | when power=0, max tx power to use, HW-specific values will be used if not set<br>**Default**: `18`<br>**Constraints**: `>= 5`, `<= 18` |
| `PowerMin` | `models.Optional[int]` | Optional | when power=0, min tx power to use, HW-specific values will be used if not set<br>**Default**: `8`<br>**Constraints**: `>= 5`, `<= 18` |
| `Preamble` | [`*models.RadioBandPreambleEnum`](../../doc/models/radio-band-preamble-enum.md) | Optional | **Default**: `"short"` |
| `StandardPower` | `*bool` | Optional | for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we'll fallback to Low Power Indoor if AFC failed<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "allow_rrm_disable": false,
  "ant_gain": 0,
  "antenna_mode": "default",
  "bandwidth": 20,
  "channels": [
    11,
    10,
    9
  ],
  "disabled": false,
  "power": 7,
  "power_max": 18,
  "power_min": 8,
  "preamble": "short",
  "standard_power": false
}
```
