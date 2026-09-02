# InsightsReportGenerationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**insights_uid** | **str** |  | [optional] 
**source** | [**DataSource**](DataSource.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.insights_report_generation_request import InsightsReportGenerationRequest

# TODO update the JSON string below
json = "{}"
# create an instance of InsightsReportGenerationRequest from a JSON string
insights_report_generation_request_instance = InsightsReportGenerationRequest.from_json(json)
# print the JSON string representation of the object
print(InsightsReportGenerationRequest.to_json())

# convert the object into a dict
insights_report_generation_request_dict = insights_report_generation_request_instance.to_dict()
# create an instance of InsightsReportGenerationRequest from a dict
insights_report_generation_request_form_dict = insights_report_generation_request.from_dict(insights_report_generation_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


