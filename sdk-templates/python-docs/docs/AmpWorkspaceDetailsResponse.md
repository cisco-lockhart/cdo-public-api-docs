# AmpWorkspaceDetailsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**url** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.amp_workspace_details_response import AmpWorkspaceDetailsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AmpWorkspaceDetailsResponse from a JSON string
amp_workspace_details_response_instance = AmpWorkspaceDetailsResponse.from_json(json)
# print the JSON string representation of the object
print(AmpWorkspaceDetailsResponse.to_json())

# convert the object into a dict
amp_workspace_details_response_dict = amp_workspace_details_response_instance.to_dict()
# create an instance of AmpWorkspaceDetailsResponse from a dict
amp_workspace_details_response_form_dict = amp_workspace_details_response.from_dict(amp_workspace_details_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


