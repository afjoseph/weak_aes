# POC attack on AES with semi-known keys

I used this logic in many CTFs, but usually in hacky Python scripts. This is an attempt to make it more formal. Furthermore, this is a portable solution (well, as portable as Go can be) to this [CNIT141 project](https://samsclass.info/141/proj/C201.htm).

# Dependencies

Go 1.16+

# Usage

```
go run .
# If Go gives you troubles about missing packages, run `go mod tidy`
```

The `bruteforcer` package is quite portable and can be used in CTFs nicely.

# Testing

```
go test ./...
```

Not every function is tested, unfortunately.
