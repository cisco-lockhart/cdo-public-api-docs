# AsaConfigCommandsDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_commands** | **List[str]** | The list of ASA configuration commands that will be applied to the device upon deployment. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the ASA device the config commands will be applied to upon deployment. | [optional] 
**status** | **str** | The status of the configuration command generation process. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_config_commands_dto import AsaConfigCommandsDto

# TODO update the JSON string below
json = "{}"
# create an instance of AsaConfigCommandsDto from a JSON string
asa_config_commands_dto_instance = AsaConfigCommandsDto.from_json(json)
# print the JSON string representation of the object
print(AsaConfigCommandsDto.to_json())

# convert the object into a dict
asa_config_commands_dto_dict = asa_config_commands_dto_instance.to_dict()
# create an instance of AsaConfigCommandsDto from a dict
asa_config_commands_dto_form_dict = asa_config_commands_dto.from_dict(asa_config_commands_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


