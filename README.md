# openmedia-files-checker

**The program finds wrongly organized OpenMedia rundown files within the given directory.**

## Features

- Report wrongly placed files.
- Report number of contacts per week.
- Match contacts to the given stations.

## Installation

- Step 1
  ```bash
  git clone https://github.com/czech-radio/openmedia-files-checker.git
  ```
- Step 2
  ```bash
  cd openmedia-files-checker
  ```
- Step 3
  ```bash
  go mod tidy
  ```
- Step 4
  ```bash
  go build
  ```

You should have mounted "Openmedia folder" `/xyz/cro.cz/Rundowns/2022/W01` or change the path to a directory begginning with `WXX`.

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
