# Go-Basic_2

#### Enumeration

##### Declaring iota

* The iota keyword represents successive integer constants 0, 1, 2,…
* It resets to 0 whenever the word const appears in the source code,
* and increments after each const specification.
```
const (
    A = iota + 1
    B   // 2
    C   // 3
    _   // 4
    D   // 5
    E   // 6
    F   // 7
)
```

```
const(
    _ = iota + 5 * 2
    A   // 16
    B   // 17
    C   // 18
    D   // 19
    E   // 20
    F   // 21
)

```
