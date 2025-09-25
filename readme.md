# StructTags

`structtags` provides a way of parsing struct tag Go fields.

The goal is to provide some ways to parse struct tags:
- Some projects need a full parsing (key, values)
- Some others only need the key and the raw value.
- Other projects need to escape the comma.
- Etc.

## Usage

- `structtags.ParseToMap(tag)`:
    - Parses a struct tag to a `map[string]string`.
- `structtags.ParseToMapMultikeys(tag)`:
    - Parses a struct tag to a `map[string][]string`.
    - For non-conventional tags where the key is repeated.
- `structtags.ParseToMapValues(tag)`:
    - Parses a struct tag to a `map[string][]string`.
    - The value is split on a comma. (support comma escaped by backslash)
- `structtags.ParseToSlice(tag)`:
    - Parses a struct tag to a slice of `parser.Tag`.
- `structtags.ParseToSliceValues(tag)`:
    - Parses a struct tag to a slice of `slicevalues.Tag`.
    - The value is split on a comma. (support comma escaped by backslash)
- `structtags.ParseToFatih(tag)`:
    - Parses a struct tag to a `*structtag.Tags`.
    - The value is split on a comma.
- `structtags.ParseToFatihExtended(tag)`:
    - Parses a struct tag to a `*structtag.Tags`.
    - Extended: support comma escaped by backslash.

The `parser` package provides the tooling to parse a struct tag and its associated value.

To implement a custom parser, you can implement the `parser.Filler` interface.

## Plan

This is the first version of the module, and I want to extend it based on feedback so that the API can evolve and break.

## Notes

The struct tag specifications say that struct tags can be any string.

The key/value syntax, the comma separator, and the space separator are conventions based on `reflect.StructTag` and `json` implementation.

`reflect.StructTag` behaves like the struct tags are `map[string]string`, but with one difference:
The first key always wins if there are multiple keys with the same name.

We can say that the `reflect.StructTag` doesn't support multiple keys with the same name.
But some rare projects/libraries use multiple keys with the same name.

Also, the specification doesn't talk about comma escaping inside the value.

Maybe the specification should clarify those points.

## References

- https://go.dev/ref/spec#Struct_types
- https://go.dev/ref/spec#string_lit
- https://github.com/golang/go/blob/411c250d64304033181c46413a6e9381e8fe9b82/src/reflect/type.go#L1030-L1108
- https://github.com/golang/tools/blob/master/go/analysis/passes/structtag/structtag.go
- https://github.com/fatih/structtag/blob/master/tags.go
