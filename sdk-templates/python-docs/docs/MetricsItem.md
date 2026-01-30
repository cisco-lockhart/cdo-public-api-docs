# MetricsItem


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** |  | [optional] 
**metrics** | [**Dict[str, Metric]**](Metric.md) |  | [optional] 
**name** | **str** |  | [optional] 
**uid** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.metrics_item import MetricsItem

# TODO update the JSON string below
json = "{}"
# create an instance of MetricsItem from a JSON string
metrics_item_instance = MetricsItem.from_json(json)
# print the JSON string representation of the object
print(MetricsItem.to_json())

# convert the object into a dict
metrics_item_dict = metrics_item_instance.to_dict()
# create an instance of MetricsItem from a dict
metrics_item_form_dict = metrics_item.from_dict(metrics_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


