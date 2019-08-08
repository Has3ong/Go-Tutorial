# Go-Error_Processing_1

#### Error Processing

* Go is the built-in type and has an interface type called error.
* Go error is sent and received through this error interface, which has one method as follows.
* Developers can create a custom error type that implements this interface.

```
f, err := os.Open("unnamedfile")
if err != nil {
	log.Fatal(err.Error())
	//log.Fatal(err)
}
fmt.Println(f.Name())
```

```
func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("%d", n)
}

a, err := notZero(1)

if err != nil {
	log.Fatal(err)
}
fmt.Println(a)


b, err := notZero(0)

if err != nil {
	log.Fatal(err)
}
fmt.Println(b)
```

#### Error Definition

```
var err1 error = errors.New("Error occurred - 1")
err2 := errors.New("Error occurred - 2")

fmt.Println(err1)
fmt.Println(err1.Error())

fmt.Println(err2)
fmt.Println(err2.Error())
```