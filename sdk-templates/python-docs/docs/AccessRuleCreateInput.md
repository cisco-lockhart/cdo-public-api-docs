# AccessRuleCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_group_uid** | **str** | The unique identifier, represented as a UUID, of the Access Group associated with the Access Rule. | 
**active_rule** | **bool** |  | [optional] 
**destination_dynamic_object** | [**DestinationDynamicObjectContent**](DestinationDynamicObjectContent.md) |  | [optional] 
**destination_network** | [**DestinationNetworkContent**](DestinationNetworkContent.md) |  | [optional] 
**destination_port** | [**DestinationPortContent**](DestinationPortContent.md) |  | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the device/manager associated with the Access Rule. Points to the shared Access Group in the case of a shared Access Rule being created. | 
**index** | **int** | The position of the Access Rule in the orded list of rules in an Access Group. | 
**log_settings** | [**LogSettings**](LogSettings.md) |  | [optional] 
**protocol** | [**ProtocolContent**](ProtocolContent.md) |  | [optional] 
**remark** | **str** | A human-readable remark. This is typically used to describe the intentions of the access rule. | [optional] 
**rule_action** | **str** | The rule&#39;s action. | [optional] 
**rule_time_range** | [**RuleTimeRangeContent**](RuleTimeRangeContent.md) |  | [optional] 
**source_dynamic_object** | [**SourceDynamicObjectContent**](SourceDynamicObjectContent.md) |  | [optional] 
**source_network** | [**SourceNetworkContent**](SourceNetworkContent.md) |  | [optional] 
**source_port** | [**SourcePortContent**](SourcePortContent.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.access_rule_create_input import AccessRuleCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRuleCreateInput from a JSON string
access_rule_create_input_instance = AccessRuleCreateInput.from_json(json)
# print the JSON string representation of the object
print(AccessRuleCreateInput.to_json())

# convert the object into a dict
access_rule_create_input_dict = access_rule_create_input_instance.to_dict()
# create an instance of AccessRuleCreateInput from a dict
access_rule_create_input_form_dict = access_rule_create_input.from_dict(access_rule_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


