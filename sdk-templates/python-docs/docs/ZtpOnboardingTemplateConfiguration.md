# ZtpOnboardingTemplateConfiguration


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**object_overrides** | [**List[FmcObjectOverride]**](FmcObjectOverride.md) | Specify a list of objects to override. All objects that have been marked as overridable in the template must be defined here. Note: Any overrides defined here that are not present in the template configuration will be ignored. | [optional] 
**template_variables** | [**List[FmcTemplateVariable]**](FmcTemplateVariable.md) | Specify the list of template variables configured on the template. All required template variables must be defined here, and should have values valid for the type of the variable. Any invalid input will result in the template not being applied to the device upon onboarding. Note: Any variables defined here that are not present in the template will be ignored. | [optional] 
**uid** | **str** | Specify the unique identifier, represented as a UUID, of the template to apply to the onboarded device. | 

## Example

```python
from scc_firewall_manager_sdk.models.ztp_onboarding_template_configuration import ZtpOnboardingTemplateConfiguration

# TODO update the JSON string below
json = "{}"
# create an instance of ZtpOnboardingTemplateConfiguration from a JSON string
ztp_onboarding_template_configuration_instance = ZtpOnboardingTemplateConfiguration.from_json(json)
# print the JSON string representation of the object
print(ZtpOnboardingTemplateConfiguration.to_json())

# convert the object into a dict
ztp_onboarding_template_configuration_dict = ztp_onboarding_template_configuration_instance.to_dict()
# create an instance of ZtpOnboardingTemplateConfiguration from a dict
ztp_onboarding_template_configuration_form_dict = ztp_onboarding_template_configuration.from_dict(ztp_onboarding_template_configuration_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


