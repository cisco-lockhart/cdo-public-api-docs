# MetricAggregationListResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[MetricAggregationListItem]**](MetricAggregationListItem.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.metric_aggregation_list_response import MetricAggregationListResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MetricAggregationListResponse from a JSON string
metric_aggregation_list_response_instance = MetricAggregationListResponse.from_json(json)
# print the JSON string representation of the object
print(MetricAggregationListResponse.to_json())

# convert the object into a dict
metric_aggregation_list_response_dict = metric_aggregation_list_response_instance.to_dict()
# create an instance of MetricAggregationListResponse from a dict
metric_aggregation_list_response_form_dict = metric_aggregation_list_response.from_dict(metric_aggregation_list_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


