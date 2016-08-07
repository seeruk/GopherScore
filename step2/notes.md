Step 2
======

`main.go`:

* `goimports` is a thing.
* `init()` function is handy.
* Dependencies!
    * No version info.
        * Vendoring.
        * gopkg.in.
    * Always absolute paths.
        * "What about forks?" `gorename`.
        * "Why not just use local paths?" It breaks tooling, like `$ go build` (think of it like unit testing, code should work as units on their own).
* Short assignment syntax.
    * Type inference!
    * Zero-values, and `var` syntax.

`cmd/search.go`:

* Packages != namespaces.
    * Always 1 level deep.
    * Exported things, and non-exported things.
* Return types. ✓
* Anonymous functions. ✓
* High order functions. ✓
* Old `main.go` in here.
    * Execute it.
    * Show contextual help.
