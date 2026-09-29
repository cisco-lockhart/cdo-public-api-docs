# LocalReadinessResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_address** | **str** |  | [optional] 
**error_msg** | **str** |  | [optional] 
**log_path** | **str** |  | [optional] 
**progress** | **int** |  | [optional] 
**stage** | **str** |  | [optional] 
**version** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.local_readiness_result import LocalReadinessResult

# TODO update the JSON string below
json = "{}"
# create an instance of LocalReadinessResult from a JSON string
local_readiness_result_instance = LocalReadinessResult.from_json(json)
# print the JSON string representation of the object
print(LocalReadinessResult.to_json())

# convert the object into a dict
local_readiness_result_dict = local_readiness_result_instance.to_dict()
# create an instance of LocalReadinessResult from a dict
local_readiness_result_form_dict = local_readiness_result.from_dict(local_readiness_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


