# BulkOperationAsaDeviceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | A collection of up to 50 unique identifiers (UUIDs) for the ASA devices in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.bulk_operation_asa_device_request import BulkOperationAsaDeviceRequest

# TODO update the JSON string below
json = "{}"
# create an instance of BulkOperationAsaDeviceRequest from a JSON string
bulk_operation_asa_device_request_instance = BulkOperationAsaDeviceRequest.from_json(json)
# print the JSON string representation of the object
print(BulkOperationAsaDeviceRequest.to_json())

# convert the object into a dict
bulk_operation_asa_device_request_dict = bulk_operation_asa_device_request_instance.to_dict()
# create an instance of BulkOperationAsaDeviceRequest from a dict
bulk_operation_asa_device_request_form_dict = bulk_operation_asa_device_request.from_dict(bulk_operation_asa_device_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


