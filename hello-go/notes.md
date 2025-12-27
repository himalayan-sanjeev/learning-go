Key idea:
- Go executes packages (directories), not files.
- Every .go file starts with package
- One directory = one package
- package main = executable
- Execution is top-to-bottom.
- There is exactly one main()
- One package per directory
- Curly braces are mandatory
- No semicolons (mostly)
- No exceptions
- No implicit conversions
- Explicit is better than clever
- Go does not have classes in the Ruby / Java sense.[No classes, no inheritance]
- Go replaces object inheritance with behavior composition.

Program Structure (Typical Go App):
hello-go/
├── go.mod
├── main.go
├── handlers/
├── services/
└── repositories/


Variables:
- Types are explicit or inferred e.g,: var name string = "Hello"
- := is shorthand inside functions only

Constants:
- Constants are compile-time values.
- const maxCount = 10

Functions:
- Types on parameters and return values
- Explicit return

Conditionals:
- Parentheses optional
- Curly braces mandatory
- No end

Loops:
-
-


Structs + Explicit Functions:
- no classes, no inheritance

Why Go Removed Inheritance
- Because inheritance causes:
- Tight coupling
- Fragile base classes
- Hidden behavior
- Hard refactors

Go enforces:
- Small interfaces
- Explicit dependencies
- Flat hierarchies
- Predictable behavior
- This is why Go scales well in large teams.

Mental Translation Table [Rails > Go]:
Rails >	Go
Class	> Struct
Inheritance	> Composition
Concern	> Embedded struct
Callback	> Explicit function
Duck typing	> Interfaces
Magic	> Readability

