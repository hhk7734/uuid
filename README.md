# uuid

The uuid package generates and inspects UUIDs based on
[RFC 4122](http://tools.ietf.org/html/rfc4122)
and DCE 1.1: Authentication and Security Services.

This package is based on the `github.com/google/uuid` package.

The difference from `githubcom/google/uuid` is that when UUID is stored in SQL, it is stored as `[]byte`.

## Install

`go get -u github.com/hhk7734/uuid/v2`
