# Go Modernize Example

This project demonstrates examples of Go code that can be automatically modernized using the `modernize` analyzer in [golang.org/x/tools/gopls/internal/analysis/modernize](https://pkg.go.dev/golang.org/x/tools/gopls/internal/analysis/modernize).

## Examples Included

The `main.go` file contains examples of the following code patterns that can be modernized:

1. **min/max functions** (Go 1.21): Replace if/else conditional assignment with built-in min/max
2. **slices.Sort** (Go 1.21): Replace sort.Slice with simpler slices.Sort
3. **any type** (Go 1.18): Replace interface{} with any
4. **slices.Clone** (Go 1.21): Replace append([]T(nil), s...) with slices.Clone
5. **fmt.Appendf** (Go 1.19): Replace []byte(fmt.Sprintf()) with fmt.Appendf(nil, ...)
6. **maps package functions** (Go 1.21): Replace map copying loops with maps.Clone
7. **slices.Delete** (Go 1.21): Replace append(s[:i], s[i+1:]...) with slices.Delete
8. **range-over-int** (Go 1.22): Replace traditional for loops with for i := range n {}
9. **strings.SplitSeq** (Go 1.24): Replace strings.Split in for-range with more efficient SplitSeq
10. **omitzero** (Go 1.24): Replace omitempty with omitzero on struct tags

Additionally, the `context_test.go` file demonstrates:
11. **t.Context** (Go 1.24): Replace context.WithCancel in tests with t.Context()

## Running Modernize

### Apply Changes

To apply all modernization fixes, run:

```bash
go tool modernize -test -fix ./...
```

You may need to run it multiple times to apply all fixes if there are conflicts.

### View Diff Without Applying Changes

To see what changes would be made without actually modifying the files:

```bash
go tool modernize -test -diff ./...
```

This shows a unified diff of the changes that would be applied.

## Continuous Integration

This project includes a GitHub Actions workflow (`.github/workflows/modernize-check.yml`) that automatically runs the modernize tool on each push and generates a summary of any potential improvements in the GitHub Actions output. This helps maintain awareness of code that could be modernized without immediately breaking the build.

The workflow will:
1. Run the modernize tool with the `-diff` option
2. If differences are found, output them to the GitHub Actions summary page
3. Show the command needed to apply the suggested changes

This helps team members stay aware of potential modernizations while making the decision to apply them at their convenience.

## Requirements

- Go 1.24 or later is recommended to leverage all modernization features.

