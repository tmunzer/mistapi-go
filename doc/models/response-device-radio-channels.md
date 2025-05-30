
# Response Device Radio Channels

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceRadioChannels`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band2440mhzAllowed` | `bool` | Required | - |
| `Band24Channels` | `map[string][]int` | Required | Property key is the channel width |
| `Band24Enabled` | `bool` | Required | - |
| `Band5Channels` | `map[string][]int` | Required | Property key is the channel width |
| `Band5Enabled` | `bool` | Required | - |
| `Band6Channels` | `map[string][]int` | Optional | Property key is the channel width |
| `Band6Enabled` | `*bool` | Optional | - |
| `Certified` | `bool` | Required | - |
| `Code` | `int` | Required | - |
| `DfsOk` | `bool` | Required | - |
| `Key` | `string` | Required | - |
| `Name` | `string` | Required | - |
| `Uses` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band24_40mhz_allowed": false,
  "band24_channels": {
    "20": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ],
    "40": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ]
  },
  "band24_enabled": false,
  "band5_channels": {
    "20": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ],
    "40": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "80": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "dfs": [
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144
    ],
    "outdoor": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ]
  },
  "band5_enabled": false,
  "band6_channels": {
    "160": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221
    ],
    "20": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221,
      225,
      229,
      233
    ],
    "40": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221,
      225,
      229
    ],
    "80": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221
    ],
    "psc": [
      5,
      21,
      37,
      53,
      69,
      85,
      101,
      117,
      133,
      149,
      165,
      181,
      197,
      213,
      229
    ]
  },
  "certified": false,
  "code": 40,
  "dfs_ok": false,
  "key": "key6",
  "name": "name6",
  "uses": "uses8",
  "band6_enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

