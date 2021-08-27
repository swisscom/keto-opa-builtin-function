# OPA built-in function to check relations against ORY Keto

This project creates a built-in function to check relations against an ORY Keto instance.

## <span style="color:red">:: POC CODE - CRINGE POSSIBLE ::</span>

## Building

As this is a built-in function, when compiled you are also compiling OPA, so you need to use this
executable instead of the usual OPA one.

You can do this with Docker by running the usual:
```shell script
docker build -t opa-keto-builtin-function .
```
This multistage Docker build will compile the executable with the built-in function and move it to a distroless/cc container.

You can then run it as you would with a normal OPA container.
