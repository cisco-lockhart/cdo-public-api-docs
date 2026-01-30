# AggregationThresholdItem


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**value** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.aggregation_threshold_item import AggregationThresholdItem

# TODO update the JSON string below
json = "{}"
# create an instance of AggregationThresholdItem from a JSON string
aggregation_threshold_item_instance = AggregationThresholdItem.from_json(json)
# print the JSON string representation of the object
print(AggregationThresholdItem.to_json())

# convert the object into a dict
aggregation_threshold_item_dict = aggregation_threshold_item_instance.to_dict()
# create an instance of AggregationThresholdItem from a dict
aggregation_threshold_item_form_dict = aggregation_threshold_item.from_dict(aggregation_threshold_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


