# MetricAggregationItem


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**aggregations** | [**List[AggregationThresholdItem]**](AggregationThresholdItem.md) |  | [optional] 
**metric** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.metric_aggregation_item import MetricAggregationItem

# TODO update the JSON string below
json = "{}"
# create an instance of MetricAggregationItem from a JSON string
metric_aggregation_item_instance = MetricAggregationItem.from_json(json)
# print the JSON string representation of the object
print(MetricAggregationItem.to_json())

# convert the object into a dict
metric_aggregation_item_dict = metric_aggregation_item_instance.to_dict()
# create an instance of MetricAggregationItem from a dict
metric_aggregation_item_form_dict = metric_aggregation_item.from_dict(metric_aggregation_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


