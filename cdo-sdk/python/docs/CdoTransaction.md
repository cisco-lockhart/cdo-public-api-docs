# CdoTransaction


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tenant_uid** | **str** | The unique identifier of the tenant that the transaction triggered on. | [optional] 
**transaction_uid** | **str** | The unique identifier of the transaction triggered. | [optional] 
**entity_uid** | **str** | The unique identifier of the entity that the transaction is triggered on. This can be empty, for a transaction that is not tied to an entity, such as transactions which refresh RA VPN sessions. | [optional] 
**entity_url** | **str** | A URL to access the entity that the transaction is triggered on. This can also be empty | [optional] 
**transaction_polling_url** | **str** | The URL to poll to track the progress of the transaction. | [optional] 
**submission_time** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the transaction was triggered | [optional] 
**last_updated_time** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the transaction status was last updated | [optional] 
**transaction_details** | **Dict[str, str]** | Transaction details, if any | [optional] 
**error_message** | **str** | Transaction error message, if any | [optional] 
**error_details** | **Dict[str, str]** | Transaction error details, if any | [optional] 
**transaction_type** | **str** | the type of the transaction | [optional] 
**cdo_transaction_status** | **str** | The status of the CDO transaction | [optional] 

## Example

```python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction

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


