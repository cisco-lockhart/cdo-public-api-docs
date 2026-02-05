# ThousandEyesOutageQueryResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**has_more_results** | **bool** |  | [optional] 
**last_evaluated_key** | **str** |  | [optional] 
**outages** | [**List[ThousandEyesOutageResponse]**](ThousandEyesOutageResponse.md) |  | [optional] 
**page_size** | **int** |  | [optional] 
**total_count** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.thousand_eyes_outage_query_response import ThousandEyesOutageQueryResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ThousandEyesOutageQueryResponse from a JSON string
thousand_eyes_outage_query_response_instance = ThousandEyesOutageQueryResponse.from_json(json)
# print the JSON string representation of the object
print(ThousandEyesOutageQueryResponse.to_json())

# convert the object into a dict
thousand_eyes_outage_query_response_dict = thousand_eyes_outage_query_response_instance.to_dict()
# create an instance of ThousandEyesOutageQueryResponse from a dict
thousand_eyes_outage_query_response_form_dict = thousand_eyes_outage_query_response.from_dict(thousand_eyes_outage_query_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


