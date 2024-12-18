# FmcDeviceRecord

(FMC-managed FTDs only) The device record in FMC. A FMC-managed device on SCC can also be accessed directly using the FMC APIs; this field provides details.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the device on the FMC. | [optional] 
**link** | **str** | The endpoint to access this resource from on the FMC. | [optional] 

## Example

```python
from cdo_sdk_python.models.fmc_device_record import FmcDeviceRecord

# TODO update the JSON string below
json = "{}"
# create an instance of FmcDeviceRecord from a JSON string
fmc_device_record_instance = FmcDeviceRecord.from_json(json)
# print the JSON string representation of the object
print(FmcDeviceRecord.to_json())

# convert the object into a dict
fmc_device_record_dict = fmc_device_record_instance.to_dict()
# create an instance of FmcDeviceRecord from a dict
fmc_device_record_form_dict = fmc_device_record.from_dict(fmc_device_record_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


