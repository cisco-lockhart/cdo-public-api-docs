# CdoTransaction


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cancellable** | **bool** |  | [optional] 
**cdo_transaction_status** | **str** | Status of the transaction. | [optional] 
**entity_uid** | **str** | Unique identifier of the entity that the transaction is triggered on. This can be empty, for a transaction that is not tied to an entity, such as transactions which refresh RA VPN sessions. | [optional] 
**entity_url** | **str** | URL to access the entity that the transaction is triggered on. This can be empty, for a transaction that is not tied to an entity, such as transactions which refresh RA VPN sessions. | [optional] 
**error_details** | **Dict[str, str]** | Transaction error details, if any. | [optional] 
**error_message** | **str** | Transaction error message, if any. | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the transaction status was last updated. | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the transaction was triggered. | [optional] 
**tenant_uid** | **str** | Unique identifier of the tenant that the transaction was triggered on. | [optional] 
**transaction_details** | **Dict[str, str]** | Transaction details, if any. | [optional] 
**transaction_polling_url** | **str** | Polling URL to track the progress of the transaction. | [optional] 
**transaction_type** | **str** | Type of the transaction. | [optional] 
**transaction_uid** | **str** | Unique identifier of the transaction triggered. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction

# TODO update the JSON string below
json = "{}"
# create an instance of CdoTransaction from a JSON string
cdo_transaction_instance = CdoTransaction.from_json(json)
# print the JSON string representation of the object
print(CdoTransaction.to_json())

# convert the object into a dict
cdo_transaction_dict = cdo_transaction_instance.to_dict()
# create an instance of CdoTransaction from a dict
cdo_transaction_form_dict = cdo_transaction.from_dict(cdo_transaction_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


