# FmcRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**commands** | [**List[Command]**](Command.md) | A list of commands to be executed on the device. | 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_request import FmcRequest

# TODO update the JSON string below
json = "{}"
# create an instance of FmcRequest from a JSON string
fmc_request_instance = FmcRequest.from_json(json)
# print the JSON string representation of the object
print(FmcRequest.to_json())

# convert the object into a dict
fmc_request_dict = fmc_request_instance.to_dict()
# create an instance of FmcRequest from a dict
fmc_request_form_dict = fmc_request.from_dict(fmc_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


