# openmedia-files-checker

**The program finds wrongly organized OpenMedia rundown files within the given directory.**

## Features

- bad modtimes in files whitin rundown repository (disk cro.cz)
- (todo) number of contacts per week
- (todo) match contacts to given stations

## Installation

```bash
git clone https://github.com/czech-radio/openmedia-files-checker.git
cd openmedia-files-checker
go mod tidy
go build
./openmedia-files-checker -i "/path/to/mounted/Rundowns/2022/W01 /path/to/mounted/Rundowns/2022/W33" [optional -o log.txt] [optional -w write changes]
```

You should have mounted "Openmedia folder" `/xyz/cro.cz/Rundowns/2022/W01` or change the path to a directory begginning with `WXX`

When it runs well, you should see something like this on the output:

```bash
2022/10/24 15:33:07 Checking Rundown count Year 2020
2020/W12: comparing file modtime to foldername: 324/324 PASSED!
2020/W13: comparing file modtime to foldername: 1632/1632 PASSED!
2020/W14: comparing file modtime to foldername: 456/456 PASSED!
2020/W15: comparing file modtime to foldername: 456/456 PASSED!
2020/W16: comparing file modtime to foldername: 456/456 PASSED!
2020/W17: comparing file modtime to foldername: 456/456 PASSED!
2020/W18: comparing file modtime to foldername: 456/456 PASSED!
...
```
