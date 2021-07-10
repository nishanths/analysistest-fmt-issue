This repository demonstrates an issue with `analysistest`.

## Explanation

The file `analyzer.go` defines an `analysis.Analyzer` that produces.

When operating on the testdata file `normalpkg.go`, the analyzer (due to buggy
behavior) produces the following syntactically invalid Go output when fixes are
applied:
```
package normalpkg

<<<>>> // want "^removing$"
<<<>>> // want "^removing$"
```

But the output does not match the golden file contents, which is:
```
package normalpkg

// want "^removing$"
// want "^removing$"
```

But running:
```
go test
```
incorrectly results in a PASS, even though the output and the golden file contents do not match.

