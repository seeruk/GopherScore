Step 1
======

* Obligatory "Hello, World!".
* Notice package name does not match folder.
    * Any executable should be in a main package.
* `println()` is a thing, but don't use it.
    * Not part of the standard library, part of the runtime. Could disappear.
    * Prints to stderr. `fmt` package allows you to choose any `io.Writer`.
