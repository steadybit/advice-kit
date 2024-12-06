# Contributing Guidelines

## Installing Required Tools

```sh
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.0
```

## Executing the Generator

```sh
./build.sh
```

## Releasing

 1. Update `CHANGELOG.md`
 2. Set the tag: `git tag -a go/advice_kit_api/v0.0.1-beta.7 -m go/advice_kit_api/v0.0.1-beta.7`
 3. Push the tag: `git push origin go/advice_kit_api/v0.0.1-beta.7`
