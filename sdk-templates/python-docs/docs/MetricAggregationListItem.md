# MetricAggregationListItem


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_name** | **str** |  | [optional] 
**device_uid** | **str** |  | [optional] 
**interface_name** | **str** |  | [optional] 
**interface_uid** | **str** |  | [optional] 
**metric** | **str** |  | [optional] 
**tenant_name** | **str** |  | [optional] 
**tenant_uid** | **str** |  | [optional] 
**value** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.metric_aggregation_list_item import MetricAggregationListItem

# TODO update the JSON string below
json = "{}"
# create an instance of MetricAggregationListItem from a JSON string
metric_aggregation_list_item_instance = MetricAggregationListItem.from_json(json)
# print the JSON string representation of the object
print(MetricAggregationListItem.to_json())

# convert the object into a dict
metric_aggregation_list_item_dict = metric_aggregation_list_item_instance.to_dict()
# create an instance of MetricAggregationListItem from a dict
metric_aggregation_list_item_form_dict = metric_aggregation_list_item.from_dict(metric_aggregation_list_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


