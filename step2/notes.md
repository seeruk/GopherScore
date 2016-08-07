Step 2
======

`main.go`:

* No longer single-file application.
* Dependencies
  * Absolute paths.
    * Single units.
  * Local vs. remote.
  * Standard library.
* Variables.
  * `var` vs. `:=` (and type inference).
* Pointers.
  * Can be identified by their type, e.g. `var foo *string`.
  * `&` generates a pointer to it's operand.
  * `*` dereferences a pointer (denotes a pointer's underlying value).

`score-calculator.go`:

* Structs.
  * Fields, sort of like properties of a class.
    * Can't set values though.
  * Methods.
    * Method receiver.
      * Appears in it's own argument list.
      * Pointer receivers.
        * When do you use them?
* Function / Method return values.
  * You can have multiple, more on that later.
* Interfaces.
  * Duck typing.
* Looping.
  * One looping construct to rule them all (`for`).

`models.go`:

* More than one struct in the same file.
* Struct field tags.

`api-client.go`:

* You can do all of this, with just the standard library?!
* Constructors.
* `defer`.
* `range`.
* `interface{}`.
* Maps.

`score-handler.go`:

* Struct used for dependencies.
* Example of pointer dereferencing.
* Error handling.
  * What could be improved?
* Encoding (marshalling) structs into JSON.

`score-calculator_test.go`:

* Brief overview.
* Sub-tests.
* Table-driven tests.
* Coverage.
  * Unfortunately, can be tricky to generate coverage for multiple packages.
