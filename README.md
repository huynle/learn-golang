# learn-golang
learning how to effectively write golang

This will be a complete package trying to touch every aspect of the language.

The motivation comes from the tour that golang 1.8 provided

to run this tour, get the source from golang and run

`go tool tour`


# Learning by tests



# Techniques to Concentrate
- simple
- readable
- maintainable

# Best Practices
- avoid nesting by handling errors first
- `switch v.(type)` to handle switches by the var type
- Put important codes first in the import list, and helpers last
- package name will always appear before the identifier you choose
- documentation
  - convention to write "method" name first in the sentence a nd about the code itself, godoc will pick this up
  - when a package gets complex, `doc.go` is convention to be placed in the package

# Test Codes
- test codes is only compiled at test time
-
