# gosum

>   see https://stackoverflow.com/questions/65304131/is-it-possible-to-load-a-go-dll-in-c-dll-on-windows/65327806#65327806

```bash
go build -ldflags "-s -w" -buildmode=c-shared -o gosum.dll
g++ -c test_gosum.cpp -o test_gosum.o -g -std=c++11
g++ test_gosum.o gosum.dll -o test_gosum.exe -g -std=c++11
.\test_gosum.exe
```