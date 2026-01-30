# InterfaceMetricErrorAggregationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**interface_errors** | [**List[InterfaceMetricError]**](InterfaceMetricError.md) |  | [optional] 
**summary** | [**InterfaceErrorSummary**](InterfaceErrorSummary.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_metric_error_aggregation_response import InterfaceMetricErrorAggregationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceMetricErrorAggregationResponse from a JSON string
interface_metric_error_aggregation_response_instance = InterfaceMetricErrorAggregationResponse.from_json(json)
# print the JSON string representation of the object
print(InterfaceMetricErrorAggregationResponse.to_json())

# convert the object into a dict
interface_metric_error_aggregation_response_dict = interface_metric_error_aggregation_response_instance.to_dict()
# create an instance of InterfaceMetricErrorAggregationResponse from a dict
interface_metric_error_aggregation_response_form_dict = interface_metric_error_aggregation_response.from_dict(interface_metric_error_aggregation_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


