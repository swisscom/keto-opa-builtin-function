# opa-keto

## Introduction

opa-keto is an extension of Open Policy Agent with two additional functions to interact with 
[Ory Keto](https://www.ory.sh/keto/).

## Building

As this is a built-in function, when compiled you are also compiling OPA, so you need to use this
executable instead of the usual OPA one.

You can do this with Docker by running the usual:
```shell script
docker build -t opa-keto .
```
This multistage Docker build will compile the executable with the built-in function and move it to a distroless/cc container.

You can then run it as you would with a normal OPA container.

## Usage

Within your policy, you'll be able to use the following new functions:

- `ketoCheck(subject, relation, namespace, object)`
- `ketoExpand(relation, namespace, object)`

```rego
# TODO: Complete rego policy example
```

## Attribution

This project is based on the original work fork of [Daniel Gozalo - @dgozalo](https://github.com/dgozalo):
[keto-opa-builtin-function](https://github.com/dgozalo/keto-opa-builtin-function).
