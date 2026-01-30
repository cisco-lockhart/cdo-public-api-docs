# CliMacroCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | A detailed description of what the CLI Macro does. | [optional] 
**device_type** | **str** | The type of device the CLI Macro is intended for. | [optional] 
**name** | **str** | A human-readable name for the CLI Macro. | [optional] 
**script** | **str** | The script content of the CLI Macro that will be executed on the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cli_macro_create_input import CliMacroCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CliMacroCreateInput from a JSON string
cli_macro_create_input_instance = CliMacroCreateInput.from_json(json)
# print the JSON string representation of the object
print(CliMacroCreateInput.to_json())

# convert the object into a dict
cli_macro_create_input_dict = cli_macro_create_input_instance.to_dict()
# create an instance of CliMacroCreateInput from a dict
cli_macro_create_input_form_dict = cli_macro_create_input.from_dict(cli_macro_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


