# CliMacroPatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | A detailed description of what the CLI Macro does. | [optional] 
**name** | **str** | A human-readable name for the CLI Macro. | [optional] 
**script** | **str** | The script content of the CLI Macro that will be executed on the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cli_macro_patch_input import CliMacroPatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of CliMacroPatchInput from a JSON string
cli_macro_patch_input_instance = CliMacroPatchInput.from_json(json)
# print the JSON string representation of the object
print(CliMacroPatchInput.to_json())

# convert the object into a dict
cli_macro_patch_input_dict = cli_macro_patch_input_instance.to_dict()
# create an instance of CliMacroPatchInput from a dict
cli_macro_patch_input_form_dict = cli_macro_patch_input.from_dict(cli_macro_patch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


