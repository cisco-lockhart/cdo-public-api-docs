# coding: utf-8

"""
    Cisco Security Cloud Control API

    Use the documentation to explore the endpoints Security Cloud Control has to offer

    The version of the OpenAPI document: 1.5.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner(BaseModel):
    """
    ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner
    """ # noqa: E501
    name: Optional[StrictStr] = None
    annotations: Optional[List[Dict[str, Any]]] = None
    declared_annotations: Optional[List[Dict[str, Any]]] = Field(default=None, alias="declaredAnnotations")
    sealed: Optional[StrictBool] = None
    specification_title: Optional[StrictStr] = Field(default=None, alias="specificationTitle")
    specification_version: Optional[StrictStr] = Field(default=None, alias="specificationVersion")
    specification_vendor: Optional[StrictStr] = Field(default=None, alias="specificationVendor")
    implementation_title: Optional[StrictStr] = Field(default=None, alias="implementationTitle")
    implementation_version: Optional[StrictStr] = Field(default=None, alias="implementationVersion")
    implementation_vendor: Optional[StrictStr] = Field(default=None, alias="implementationVendor")
    __properties: ClassVar[List[str]] = ["name", "annotations", "declaredAnnotations", "sealed", "specificationTitle", "specificationVersion", "specificationVendor", "implementationTitle", "implementationVersion", "implementationVendor"]

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
        """Create an instance of ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner from a JSON string"""
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
        """Create an instance of ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "name": obj.get("name"),
            "annotations": obj.get("annotations"),
            "declaredAnnotations": obj.get("declaredAnnotations"),
            "sealed": obj.get("sealed"),
            "specificationTitle": obj.get("specificationTitle"),
            "specificationVersion": obj.get("specificationVersion"),
            "specificationVendor": obj.get("specificationVendor"),
            "implementationTitle": obj.get("implementationTitle"),
            "implementationVersion": obj.get("implementationVersion"),
            "implementationVendor": obj.get("implementationVendor")
        })
        return _obj


