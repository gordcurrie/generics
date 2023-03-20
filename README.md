# Go Generics Examples

Prior to generics being introduced in `v1.18` when you wanted to write code that could handle different types you had a few options

* You could re-implement your code for each type you needed to support, or
* You could have your code accept empty interfaces `interface{}` and leverage type assertion.

Examples of this can be found in the pre118 package. https://github.com/gordcurrie/generics/tree/main/pre118

1.18 introduced Generics to the langauge, allowing you to add type constraints to your code as well as adding the key word any, which is an alias to `interface{}`

Examples of this can be found in the post118 package. https://github.com/gordcurrie/generics/tree/main/post118
