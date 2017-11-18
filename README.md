## PHCSV
### CSV

Parameters:
```
Copyright 2017 @IrekRomaniuk. All jdk-rights reversed.
Usage:
  -min int
        Min of lines in csv files (or will skip)
  -name string
        Name of custom list (default "Shields")
  -p string
        Phantom password
  -path string
        path to read csv file from (default "./")
  -u string
        Phantom username (default "admin")
  -url string
        Phantom REST endpoint (default "https://10.34.1.110/rest/decided_list/")
  -v    Prints current version
```

Example of use:

```
$ ./phcsv.sh 
List size: 1230
url: https://10.34.1.110/rest/decided_list/mylist
Response ID: 0
```