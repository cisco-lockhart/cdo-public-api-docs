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
from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class AccessGroup(BaseModel):
    """
    AccessGroup
    """ # noqa: E501
    uid: StrictStr = Field(description="The unique identifier of Access Group in CDO.")
    name: StrictStr = Field(description="The name of Access Group in CDO. Access Group names are unique in CDO.")
    entity_uid: StrictStr = Field(description="The unique identifier of the device/manager associated with the Access Group.", alias="entityUid")
    shared_access_group_uid: StrictStr = Field(description="The unique identifier of the shared access group manager associated with the Access Group.", alias="sharedAccessGroupUid")
    is_shared: Optional[StrictBool] = Field(default=None, description="The flag that identifies if access group is shared.", alias="isShared")
    applied_to: Optional[List[StrictStr]] = Field(default=None, description="The set of device unique identifiers to which this Access Group was applied. Only valid for shared access group.", alias="appliedTo")
    resources: Optional[List[Dict[str, Dict[str, Any]]]] = Field(default=None, description="The set of of interface and direction pairs or global resource.")
    created_date: Optional[datetime] = Field(default=None, description="The time (in UTC) at which Access Group was created, represented using the RFC-3339 standard.", alias="createdDate")
    updated_date: Optional[datetime] = Field(default=None, description="The time (in UTC) at which Access Group was updated, represented using the RFC-3339 standard.", alias="updatedDate")
    __properties: ClassVar[List[str]] = ["uid", "name", "entityUid", "sharedAccessGroupUid", "isShared", "appliedTo", "resources", "createdDate", "updatedDate"]

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
        """Create an instance of AccessGroup from a JSON string"""
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
        """Create an instance of AccessGroup from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "uid": obj.get("uid"),
            "name": obj.get("name"),
            "entityUid": obj.get("entityUid"),
            "sharedAccessGroupUid": obj.get("sharedAccessGroupUid"),
            "isShared": obj.get("isShared"),
            "appliedTo": obj.get("appliedTo"),
            "resources": obj.get("resources"),
            "createdDate": obj.get("createdDate"),
            "updatedDate": obj.get("updatedDate")
        })
        return _obj

