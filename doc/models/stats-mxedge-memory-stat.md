
# Stats Mxedge Memory Stat

Memory usage

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeMemoryStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Active` | `*int` | Optional | The amount of memory, in kibibytes, that has been used more recently and is usually not reclaimed unless absolutely necessary. |
| `Available` | `*int64` | Optional | An estimate of how much memory is available for starting new applications, without swapping. |
| `Buffers` | `*int` | Optional | The amount, in kibibytes, of temporary storage for raw disk blocks. |
| `Cached` | `*int` | Optional | The amount of physical RAM, in kibibytes, used as cache memory. |
| `Free` | `*int64` | Optional | The amount of physical RAM, in kibibytes, left unused by the system |
| `Inactive` | `*int` | Optional | The amount of memory, in kibibytes, that has been used less recently and is more eligible to be reclaimed for other purposes. |
| `SwapCached` | `*int` | Optional | The amount of memory, in kibibytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile. |
| `SwapFree` | `*int` | Optional | The total amount of swap free, in kibibytes. |
| `SwapTotal` | `*int` | Optional | The total amount of swap available, in kibibytes. |
| `Total` | `*int64` | Optional | Total amount of usable RAM, in kibibytes, which is physical RAM minus a number of reserved bits and the kernel binary code |
| `Usage` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "active": 394936320,
  "available": 4699291648,
  "buffers": 107646976,
  "cached": 478060544,
  "free": 4330659840,
  "inactive": 211980288,
  "swap_cached": 0,
  "swap_free": 1022357504,
  "swap_total": 1022357504,
  "total": 8365957120,
  "usage": 48,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

