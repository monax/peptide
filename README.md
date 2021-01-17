# I just want to say sorry

_**This is a proof-of-concept**_

Peptide is an anti-homage to amino the now defunct semi-protobuf-compatible library (https://github.com/tendermint/go-amino) (thanks for all the fish amino :) ).

This aims to provide support for custom types with the Google Go Protobuf V2 API (https://github.com/protocolbuffers/protobuf-go)

Following the approach suggested protobuf-go maintainer here: https://github.com/xen0n/protobuf-gogogo/issues/1

## Theft

I have _borrowed_ the implementations provided by @xen0n (see: https://github.com/gogo/protobuf/issues/691) for the `gogoproto` extensions listed below:

> I took the time to re-implement the top 3 extensions needed by our projects, in the form of protobuf-gogogo (pardon the name). Note this is very preliminary work without testing infrastructure set up, don't use in (your) production yet.

The extensions implemented:
- `gogoproto.moretags`
- `gogoproto.jsontag`
- `gogoproto.enumvalue_customname`

## Test

Run:

```bash
./test.bash
```

Perhaps some tea? 

### Extensions

To test additional extensions add a `.proto` file to [testdata](./cmd/protoc-gen-go-peptide/testdata)

Then run:

```bash 
./regenerate.bash
```

Then `./test.bash` again to see if the `.pb.go` files compile. If they compile they are probably fine and should be shipped.

## Credits

- This is a partial fork of: https://github.com/protocolbuffers/protobuf-go
- This is based on: https://github.com/xen0n/protobuf-gogogo
