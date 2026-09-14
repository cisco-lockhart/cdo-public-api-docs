# RemoteReadinessResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**check_identifier** | **str** |  | [optional] 
**device_address** | **str** |  | [optional] 
**error_msg** | **str** |  | [optional] 
**progress** | **int** |  | [optional] 
**stage** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.remote_readiness_result import RemoteReadinessResult

# TODO update the JSON string below
json = "{}"
# create an instance of RemoteReadinessResult from a JSON string
remote_readiness_result_instance = RemoteReadinessResult.from_json(json)
# print the JSON string representation of the object
print(RemoteReadinessResult.to_json())

# convert the object into a dict
remote_readiness_result_dict = remote_readiness_result_instance.to_dict()
# create an instance of RemoteReadinessResult from a dict
remote_readiness_result_form_dict = remote_readiness_result.from_dict(remote_readiness_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


