# ThousandEyesOutageQueryRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application_name** | **str** |  | [optional] 
**application_specific** | **bool** |  | [optional] 
**effective_page_size** | **int** |  | [optional] 
**end_date** | **datetime** |  | 
**last_evaluated_key** | **str** |  | [optional] 
**page_size** | **int** |  | [optional] 
**start_date** | **datetime** |  | 
**valid_date_range** | **bool** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_outage_query_request import ThousandEyesOutageQueryRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesOutageQueryRequest from a JSON string
thousand_eyes_outage_query_request_instance = ThousandEyesOutageQueryRequest.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesOutageQueryRequest.to_json())

# convert the object into a dict
thousand_eyes_outage_query_request_dict = thousand_eyes_outage_query_request_instance.to_dict()
# create an instance of ThousandEyesOutageQueryRequest from a dict
thousand_eyes_outage_query_request_form_dict = thousand_eyes_outage_query_request.from_dict(thousand_eyes_outage_query_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


