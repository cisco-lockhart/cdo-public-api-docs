# TargetsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**targets_uuids** | **List[str]** | The list of UIDs of the targets to be deleted. | [optional] 

## Example

```python
from cdo_python_sdk.models.targets_request import TargetsRequest

# TODO update the JSON string below
json = "{}"
# create an instance of TargetsRequest from a JSON string
targets_request_instance = TargetsRequest.from_json(json)
# print the JSON string representation of the object
print(TargetsRequest.to_json())

# convert the object into a dict
targets_request_dict = targets_request_instance.to_dict()
# create an instance of TargetsRequest from a dict
targets_request_form_dict = targets_request.from_dict(targets_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


