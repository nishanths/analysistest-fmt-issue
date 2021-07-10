This repository demonstrates an issue with `analysistest`. See https://github.com/golang/go/issues/47118.

## Explanation

The file `analyzer.go` defines an `analysis.Analyzer`.

The test in `analyzer_test.go` uses `analysistest.RunWithSuggestedFixes()`.

When operating on the testdata file `normal.go`, the analyzer (due to buggy
behavior) produces the following syntactically invalid Go output when fixes are
applied:
```go
package normalpkg

<<<>>> // want "^removing$"
<<<>>> // want "^removing$"
```

This output does not match the golden file content:
```go
package normalpkg

// want "^removing$"
// want "^removing$"
```

But running:
```
go test
```
results in a PASS, though the output and the golden file do not match.
