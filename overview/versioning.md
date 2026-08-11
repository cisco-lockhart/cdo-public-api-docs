# Versioning & Deprecation

## Versioning

API versioning enables APIs to evolve while providing consumers with a predictable mechanism for managing changes. When a change cannot be introduced without affecting existing integrations, a new API version may be provided.

## Deprecation

Deprecation is a standard stage in the API lifecycle and may apply to an entire API version or to individual API operations or endpoints.

Deprecation indicates that an API version, operation, or endpoint is no longer recommended for use and may be removed in the future. Where available, consumers should migrate to the recommended replacement, such as a newer API version or alternative operation.

Deprecation itself is not a breaking change. Deprecated APIs remain supported for **a minimum deprecation period of 12 months from the date of deprecation.**

Consumers should begin migration planning as soon as an API is marked as deprecated. Continued use of deprecated APIs is not recommended, as compatibility and availability are not guaranteed beyond the deprecation period.

## Sunsetting

Sunsetting is the permanent removal of support for an API version, operation, or endpoint and is considered a breaking change.

A deprecated API version, operation, or endpoint will remain available for **a minimum of 12 months from its deprecation date**. Once this 12-month deprecation period has elapsed, **it may be sunset at any time without further notice**.

Consumers are responsible for migrating to a supported replacement before the deprecation period expires. Integrations that continue to rely on deprecated APIs after this period may stop functioning when the API is sunset.