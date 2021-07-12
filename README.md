This repository demonstrates an issue with `analysistest`.

https://github.com/golang/go/issues/47118.

## Explanation

* `analyzer.go` defines an `analysis.Analyzer`.
* `analyzer_test.go` runs a test on the Analyzer using `analysistest.RunWithSuggestedFixes()`.

When operating on the testdata file `normal.go`, the Analyzer (due to buggy
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
