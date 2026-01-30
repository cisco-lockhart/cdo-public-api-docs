# SdcPatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.sdc_patch_input import SdcPatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of SdcPatchInput from a JSON string
sdc_patch_input_instance = SdcPatchInput.from_json(json)
# print the JSON string representation of the object
print(SdcPatchInput.to_json())

# convert the object into a dict
sdc_patch_input_dict = sdc_patch_input_instance.to_dict()
# create an instance of SdcPatchInput from a dict
sdc_patch_input_form_dict = sdc_patch_input.from_dict(sdc_patch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


