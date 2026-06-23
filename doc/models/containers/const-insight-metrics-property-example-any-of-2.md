
# Const Insight Metrics Property Example Any of 2

Example values for an insight metric property, as an array or keyed object

## Class Name

`ConstInsightMetricsPropertyExampleAnyOf2`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`[]models.ConstInsightMetricsPropertyExample`](../../../doc/models/containers/const-insight-metrics-property-example.md) | models.ConstInsightMetricsPropertyExampleAnyOf2Container.FromArrayOfConstInsightMetricsPropertyExample([]models.ConstInsightMetricsPropertyExample arrayOfConstInsightMetricsPropertyExample) |
| [`map[string][]models.ConstInsightMetricsPropertyExample`](../../../doc/models/containers/const-insight-metrics-property-example.md) | models.ConstInsightMetricsPropertyExampleAnyOf2Container.FromMapOfArrayOfConstInsightMetricsPropertyExample2(map[string][]models.ConstInsightMetricsPropertyExample mapOfArrayOfConstInsightMetricsPropertyExample2) |

## []models.ConstInsightMetricsPropertyExample

### Initialization Code

#### Example

```go
value := models.ConstInsightMetricsPropertyExampleAnyOf2Container.FromArrayOfConstInsightMetricsPropertyExample([]models.ConstInsightMetricsPropertyExample{
    models.ConstInsightMetricsPropertyExampleContainer.FromNumber(73),
})
```

## map[string][]models.ConstInsightMetricsPropertyExample

### Initialization Code

#### Example

```go
value := models.ConstInsightMetricsPropertyExampleAnyOf2Container.FromMapOfArrayOfConstInsightMetricsPropertyExample2(map[string][]models.ConstInsightMetricsPropertyExample{
    "key0": map[string][]models.ConstInsightMetricsPropertyExample{
        "": nil,
        "": nil,
        "": nil,
    },
})
```

