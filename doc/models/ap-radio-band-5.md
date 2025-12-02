
# Ap Radio Band 5

Radio Band AP settings

## Structure

`ApRadioBand5`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | **Default**: `false` |
| `AntGain` | `models.Optional[int]` | Optional | **Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `AntennaBeamPattern` | [`*models.RadioBandAntennaBeamPatternEnum`](../../doc/models/radio-band-antenna-beam-pattern-enum.md) | Optional | enum: `narrow`, `medium`, `wide` |
| `AntennaMode` | [`*models.RadioBandAntennaModeEnum`](../../doc/models/radio-band-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br><br>**Default**: `"default"` |
| `Bandwidth` | [`*models.Dot11Bandwidth5Enum`](../../doc/models/dot-11-bandwidth-5-enum.md) | Optional | channel width for the 5GHz band. enum: `0`(disabled, response only), `20`, `40`, `80`<br><br>**Default**: `40` |
| `Channel` | `models.Optional[int]` | Optional | For Device. (primary) channel for the band, 0 means using the Site Setting |
| `Channels` | `models.Optional[[]int]` | Optional | For RFTemplates. List of channels, null or empty array means auto |
| `Disabled` | `*bool` | Optional | Whether to disable the radio<br><br>**Default**: `false` |
| `Power` | `models.Optional[int]` | Optional | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …<br><br>**Constraints**: `>= 5`, `<= 25` |
| `PowerMax` | `models.Optional[int]` | Optional | When power=0, max tx power to use, HW-specific values will be used if not set<br><br>**Default**: `17`<br><br>**Constraints**: `>= 5`, `<= 17` |
| `PowerMin` | `models.Optional[int]` | Optional | When power=0, min tx power to use, HW-specific values will be used if not set<br><br>**Default**: `8`<br><br>**Constraints**: `>= 5`, `<= 17` |
| `Preamble` | [`*models.RadioBandPreambleEnum`](../../doc/models/radio-band-preamble-enum.md) | Optional | enum: `auto`, `long`, `short`<br><br>**Default**: `"short"` |

## Example (as JSON)

```json
{
  "allow_rrm_disable": false,
  "ant_gain": 0,
  "antenna_mode": "default",
  "bandwidth": 40,
  "channel": 100,
  "channels": null,
  "disabled": false,
  "power": 6,
  "power_max": 17,
  "power_min": 8,
  "preamble": "short",
  "antenna_beam_pattern": "medium"
}
```

