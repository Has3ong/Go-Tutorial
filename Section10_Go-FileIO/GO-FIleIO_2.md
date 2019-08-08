## FileIO_2

#### ioutils

* The Go standard package, the ioutil package, is a package that provides convenient I/O-related Utilities.
* If the input file is not very large, the ReadFile and WriteFile functions in this package can be used to read and write files conveniently.
* The example below is a code that copies files as they are using the ioutil.

```
s := "Hello Golang!\n File Write Test!\n"

err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
errCheck(err)

data, err := ioutil.ReadFile("sample.txt")

fmt.Println(string(data))
```

```
file, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

wt := bufio.NewWriter(file)
wt.WriteString("Hello Golang1!\nFile Write Test1!\n")
wt.Write([]byte("Hello Golang2!\nFile Write Test2!\n"))
```