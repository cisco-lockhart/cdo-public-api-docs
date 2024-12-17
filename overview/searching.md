# Searching

Several list endpoints in the CDO Public API allow you to search across multiple fields using the Lucene query syntax.

## How to search on a field

For an endpoint `/foo/bar` that returns a list of objects with a field `baz`, you can search for entries where the field `baz` has the value example by using the `q` query parameter as follows:

```http
GET /foo/bar?q=baz:example
```

### Wildcard searches
Some fields support wildcard searches. For example, with an endpoint `/foo/bar` that returns a list of objects with a field `baz`, you can search for entries where baz contains the string `example` anywhere in its value using the `q` query parameter:

```http
GET /foo/bar?q=baz:*example*
```

### Time Range Queries

Date fields may allow time range queries.

For example, with an endpoint `/foo/bar` that returns a list of objects with a field `loginTime` you can search for entries where `loginTime` falls between 10 AM and 1 PM UTC on December 17, 2024, using the following query:

```http
GET /foo/bar?q=loginTime:[2024-12-17Z10:00:00Z TO 2024-12-17Z13:00:00Z]
```

>**Note**:

- All dates must comply with [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339).
- To make an exclusive time range query (for the lower bound, upper bound, or both), replace the `[` with a `{` or the  `]` with a `}`. For example:
   - `/foo/bar?q=loginTime:{2024-12-17Z10:00:00Z TO 2024-12-17Z13:00:00Z]` makes the lower bound exclusive.
   - `/foo/bar?q=loginTime:[2024-12-17Z10:00:00Z TO 2024-12-17Z13:00:00Z}` makes the upper bound exclusive.
   - `/foo/bar?q=loginTime:{2024-12-17Z10:00:00Z TO 2024-12-17Z13:00:00Z}` makes both bounds exclusive.
 - Use the `*` wildcard to make open-ended searches. For example:
   - `/foo/bar?q=loginTime:{\* TO 2024-12-17Z13:00:00Z}` retrieves all records where the login time is earlier than 1 PM UTC on December 17, 2024.
 - When making these API calls using CuRL or Postman, you must URL-encode the query. For example:
   - `/foo/bar?q=loginTime:[2024-12-17Z10:00:00Z TO 2024-12-17Z13:00:00Z]` should be encoded as `/foo/bar?q=loginTime:%5B2024-12-17Z10:00:00Z%20TO%202024-12-17Z13:00:00Z%5D`.

## Which fields can I search on?

To determine whether a particular field in a list API endpoint is searchable, attempt a search (as described above). If the field is not searchable, the API will respond with an `HTTP 400` error, along with a list of searchable fields for that endpoint.

The same principle applies to determining whether a field supports wildcard searches.