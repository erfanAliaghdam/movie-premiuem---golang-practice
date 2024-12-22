# How to run tests?
> go test ./... -v

# How to check coverage?
- Step1:
    > go test ./... -coverprofile=cover.out
- Step2:
    >  go tool cover -html=cover.out
