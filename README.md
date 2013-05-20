gofilter
========

gofilter creates filter function that remove elements from a slice where a provided predicate returns false. The difference to common filters is that this one can handle arbitrary typed slices by using reflect.MakeFunc from go1.1.

Requires go1.1.

Copyright: MIT license
