# ThousandEyesOutageResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affected_domains** | **List[str]** |  | [optional] 
**alert_id** | **str** |  | [optional] 
**application_name** | **str** |  | [optional] 
**duration_seconds** | **int** |  | [optional] 
**end_date** | **datetime** |  | [optional] 
**errors** | **List[str]** |  | [optional] 
**outage_type** | **str** |  | [optional] 
**regions** | **List[str]** |  | [optional] 
**severity** | **str** |  | [optional] 
**start_date** | **datetime** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_outage_response import ThousandEyesOutageResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesOutageResponse from a JSON string
thousand_eyes_outage_response_instance = ThousandEyesOutageResponse.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesOutageResponse.to_json())

# convert the object into a dict
thousand_eyes_outage_response_dict = thousand_eyes_outage_response_instance.to_dict()
# create an instance of ThousandEyesOutageResponse from a dict
thousand_eyes_outage_response_form_dict = thousand_eyes_outage_response.from_dict(thousand_eyes_outage_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


