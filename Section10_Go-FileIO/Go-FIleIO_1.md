## FileIO_1

#### Write

```
file, err := os.Create("TEST.txt")
errCheck1(err)
string := "Hello Golang\n Hello VisualStudioCode"
n2, err := file.WriteString(string)
errCheck2(err)

fmt.Println(n2)
defer file.Close()
```

#### Read
```
file, err := os.Open("TEST.txt")
errCheck1(err)

fileInfo, err := file.Stat()

fd1 := make([]byte, fileInfo.Size())
ct1, err := file.Read(fd1)

fmt.Println(fileInfo)
fmt.Println(fileInfo.Name(), "\n")
fmt.Println(fileInfo.Size(), "\n")
fmt.Println(fileInfo.ModTime(), "\n")
fmt.Println(ct1)
fmt.Println(string(fd1))
```