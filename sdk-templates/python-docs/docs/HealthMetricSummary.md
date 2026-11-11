# HealthMetricSummary


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**average** | **float** |  | [optional] 
**latest** | **float** |  | [optional] 
**max** | **float** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.health_metric_summary import HealthMetricSummary

# TODO update the JSON string below
json = "{}"
# create an instance of HealthMetricSummary from a JSON string
health_metric_summary_instance = HealthMetricSummary.from_json(json)
# print the JSON string representation of the object
print(HealthMetricSummary.to_json())

# convert the object into a dict
health_metric_summary_dict = health_metric_summary_instance.to_dict()
# create an instance of HealthMetricSummary from a dict
health_metric_summary_form_dict = health_metric_summary.from_dict(health_metric_summary_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


