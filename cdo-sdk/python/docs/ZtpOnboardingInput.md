# ZtpOnboardingInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Specify a human-readable name for the device. | 
**serial_number** | **str** | Specify the serial number of the FTD device to be onboarded. When a device with this serial number is plugged in and connected to the Internet, it will automatically be registered to this SCC tenant. | 
**admin_password** | **str** | Specify the initial provisioning password. This is required for setting up the FTD, and can be ignored if password is already set on the device. | [optional] 
**fmc_access_policy_uid** | **str** | Specify the unique identifier, represented as a UUID, of the FMC access policy to apply to this device. | 
**licenses** | **List[str]** | Specify a set of licenses to apply to the device. | 
**device_group_uid** | **str** | Specify the unique identifier, represented as a UUID, of the device group which the device will be a part of after it finishes registering with SCC. | [optional] 

## Example

```python
from cdo_sdk_python.models.ztp_onboarding_input import ZtpOnboardingInput

# TODO update the JSON string below
json = "{}"
# create an instance of ZtpOnboardingInput from a JSON string
ztp_onboarding_input_instance = ZtpOnboardingInput.from_json(json)
# print the JSON string representation of the object
print(ZtpOnboardingInput.to_json())

# convert the object into a dict
ztp_onboarding_input_dict = ztp_onboarding_input_instance.to_dict()
# create an instance of ZtpOnboardingInput from a dict
ztp_onboarding_input_form_dict = ztp_onboarding_input.from_dict(ztp_onboarding_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


