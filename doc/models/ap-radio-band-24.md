
# Ap Radio Band 24

Radio Band AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApRadioBand24`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | **Default**: `false` |
| `AntGain` | `models.Optional[int]` | Optional | **Default**: `0`<br>**Constraints**: `>= 0`, `<= 10` |
| `AntennaMode` | [`*models.RadioBandAntennaModeEnum`](../../doc/models/radio-band-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br>**Default**: `"default"` |
| `Bandwidth` | [`*models.Dot11Bandwidth24Enum`](../../doc/models/dot-11-bandwidth-24-enum.md) | Optional | channel width for the 2.4GHz band. enum: `20`, `40`<br>**Default**: `20` |
| `Channel` | `models.Optional[int]` | Optional | For Device. (primary) channel for the band, 0 means using the Site Setting<br>**Constraints**: `>= 1`, `<= 13` |
| `Channels` | `models.Optional[[]int]` | Optional | For RFTemplates. List of channels, null or empty array means auto |
| `Disabled` | `*bool` | Optional | whether to disable the radio<br>**Default**: `false` |
| `Power` | `models.Optional[int]` | Optional | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …<br>**Constraints**: `>= 3`, `<= 25` |
| `PowerMax` | `models.Optional[int]` | Optional | when power=0, max tx power to use, HW-specific values will be used if not set<br>**Default**: `17`<br>**Constraints**: `>= 3`, `<= 18` |
| `PowerMin` | `models.Optional[int]` | Optional | when power=0, min tx power to use, HW-specific values will be used if not set<br>**Default**: `8`<br>**Constraints**: `>= 3`, `<= 18` |
| `Preamble` | [`*models.RadioBandPreambleEnum`](../../doc/models/radio-band-preamble-enum.md) | Optional | enum: `auto`, `long`, `short`<br>**Default**: `"short"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_rrm_disable": false,
  "ant_gain": 0,
  "antenna_mode": "default",
  "bandwidth": 20,
  "channel": 6,
  "channels": null,
  "disabled": false,
  "power": 3,
  "power_max": 17,
  "power_min": 8,
  "preamble": "short",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

