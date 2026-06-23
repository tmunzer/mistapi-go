
# Insight Metrics Results Item

Insight metric result item, returned either as a number or an object depending on the requested metric

## Class Name

`InsightMetricsResultsItem`

## Cases

| Type | Factory Method |
|  --- | --- |
| `float64` | models.InsightMetricsResultsItemContainer.FromPrecision(float64 precision) |
| `interface{}` | models.InsightMetricsResultsItemContainer.FromObject(interface{} object) |

## float64

### Initialization Code

#### Example

```go
value := models.InsightMetricsResultsItemContainer.FromPrecision(float64(0))
```

## interface{}

### Initialization Code

#### Example

```go
value := models.InsightMetricsResultsItemContainer.FromObject(interface{}("[key1, val1][key2, val2]"))
```

