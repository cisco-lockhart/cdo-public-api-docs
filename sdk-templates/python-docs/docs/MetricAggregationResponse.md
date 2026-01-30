# MetricAggregationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metric_aggregations** | [**List[MetricAggregationItem]**](MetricAggregationItem.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.metric_aggregation_response import MetricAggregationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MetricAggregationResponse from a JSON string
metric_aggregation_response_instance = MetricAggregationResponse.from_json(json)
# print the JSON string representation of the object
print(MetricAggregationResponse.to_json())

# convert the object into a dict
metric_aggregation_response_dict = metric_aggregation_response_instance.to_dict()
# create an instance of MetricAggregationResponse from a dict
metric_aggregation_response_form_dict = metric_aggregation_response.from_dict(metric_aggregation_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


