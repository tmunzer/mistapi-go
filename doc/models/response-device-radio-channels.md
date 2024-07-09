
# Response Device Radio Channels

## Structure

`ResponseDeviceRadioChannels`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band2440mhzAllowed` | `bool` | Required | - |
| `Band24Channels` | `map[string]interface{}` | Required | Property key is the channel width |
| `Band24Enabled` | `bool` | Required | - |
| `Band5Channels` | `map[string]interface{}` | Required | Property key is the channel width |
| `Band5Enabled` | `bool` | Required | - |
| `Band6Channels` | `map[string]interface{}` | Optional | Property key is the channel width |
| `Band6Enabled` | `*bool` | Optional | - |
| `Certified` | `bool` | Required | - |
| `Code` | `int` | Required | - |
| `DfsOk` | `bool` | Required | - |
| `Key` | `string` | Required | - |
| `Name` | `string` | Required | - |
| `Uses` | `string` | Required | - |

## Example (as JSON)

```json
{
  "band24_40mhz_allowed": false,
  "band24_channels": {
    "20": {
      "key1": "val1",
      "key2": "val2"
    },
    "40": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band24_enabled": false,
  "band5_channels": {
    "20": {
      "key1": "val1",
      "key2": "val2"
    },
    "40": {
      "key1": "val1",
      "key2": "val2"
    },
    "80": {
      "key1": "val1",
      "key2": "val2"
    },
    "dfs": {
      "key1": "val1",
      "key2": "val2"
    },
    "outdoor": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band5_enabled": false,
  "band6_channels": {
    "160": {
      "key1": "val1",
      "key2": "val2"
    },
    "20": {
      "key1": "val1",
      "key2": "val2"
    },
    "40": {
      "key1": "val1",
      "key2": "val2"
    },
    "80": {
      "key1": "val1",
      "key2": "val2"
    },
    "psc": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "certified": false,
  "code": 60,
  "dfs_ok": false,
  "key": "key8",
  "name": "name8",
  "uses": "uses0",
  "band6_enabled": false
}
```

