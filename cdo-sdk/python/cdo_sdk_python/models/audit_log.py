# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.1.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class AuditLog(BaseModel):
    """
    The list of items retrieved.
    """ # noqa: E501
    uid: StrictStr = Field(description="The unique identifier of the Audit Log.")
    event_type: Optional[StrictStr] = Field(default=None, description="The type of the Audit Log event.", alias="eventType")
    username: Optional[StrictStr] = Field(default=None, description="The name/email of the of user the Audit Log refers to.")
    event_description: Optional[StrictStr] = Field(default=None, description="The description of the Audit Log event.", alias="eventDescription")
    event_time: Optional[datetime] = Field(default=None, description="The time (UTC; represented using the RFC-3339 standard) at which the Audit Log event was created.", alias="eventTime")
    roles: Optional[List[StrictStr]] = Field(default=None, description="The roles of the user who did the Audit log operation")
    __properties: ClassVar[List[str]] = ["uid", "eventType", "username", "eventDescription", "eventTime", "roles"]

    @field_validator('event_type')
    def event_type_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['USER_LOGGED_IN', 'USER_ADDED_TO_TENANT', 'USER_REMOVED_FROM_TENANT', 'USER_ROLE_CHANGED', 'AD_GROUP_ADDED', 'AD_GROUP_DELETED', 'AD_GROUP_ROLE_CHANGED', 'USER_LOGIN_DATA']):
            raise ValueError("must be one of enum values ('USER_LOGGED_IN', 'USER_ADDED_TO_TENANT', 'USER_REMOVED_FROM_TENANT', 'USER_ROLE_CHANGED', 'AD_GROUP_ADDED', 'AD_GROUP_DELETED', 'AD_GROUP_ROLE_CHANGED', 'USER_LOGIN_DATA')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of AuditLog from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of AuditLog from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "uid": obj.get("uid"),
            "eventType": obj.get("eventType"),
            "username": obj.get("username"),
            "eventDescription": obj.get("eventDescription"),
            "eventTime": obj.get("eventTime"),
            "roles": obj.get("roles")
        })
        return _obj

