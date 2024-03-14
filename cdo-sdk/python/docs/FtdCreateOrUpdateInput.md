# FtdCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Specify a human-readable name for the device. | 
**licenses** | **List[str]** | Specify a set of licenses to apply to the device. | 
**virtual** | **bool** | Indicate whether the FTD is a virtual or a physical device. | [optional] 
**fmc_access_policy_uid** | **str** | Specify the unique identifier of the FMC access policy to apply to this device. | 
**performance_tier** | **str** | Specify the performance tier of the FTDv (required only if isVirtual is set to true) | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**device_type** | **str** | Specify the type of the FTD. The only supported type of FTD is CDFMC_MANAGED_FTD | 

## Example

```python
from openapi_client.models.ftd_create_or_update_input import FtdCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdCreateOrUpdateInput from a JSON string
ftd_create_or_update_input_instance = FtdCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(FtdCreateOrUpdateInput.to_json())

# convert the object into a dict
ftd_create_or_update_input_dict = ftd_create_or_update_input_instance.to_dict()
# create an instance of FtdCreateOrUpdateInput from a dict
ftd_create_or_update_input_form_dict = ftd_create_or_update_input.from_dict(ftd_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


