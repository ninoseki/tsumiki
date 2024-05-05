# Tsumiki

A collection of OPA functions.

## Semantic Versions

OPA has [built-in semver functions](https://www.openpolicyagent.org/docs/latest/policy-reference/#semver).

But a function for coercing is missing. `semver.coerce` is for that.

```rego
semver.coerce("1")
>>> 1.0.0
semver.coerce("1.0")
>>> 1.0.0
```

## Package URL

```rego
purl.parse("pkg:gem/jruby-launcher@1.1.2?platform=java")
>>> {
>>>     "name": "jruby-launcher",
>>>     "namespace": "",
>>>     "qualifiers": {
>>>         "key": "platform",
>>>         "value": "java"
>>>     },
>>>     "subpath": "",
>>>     "type": "gem",
>>>     "version": "1.1.2"
>>> }
```

## Validators

```rego
validators.is_url("http://example.com")
>>> true
validators.is_domain("example.com")
>>> true
validators.is_email("test@example.com")
>>> true
validators.is_ipv4("1.1.1.1")
>>> true
validators.is_ipv6("2001:4860:0:2001::68")
>>> true
```
