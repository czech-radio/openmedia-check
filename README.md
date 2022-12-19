# openmedia-check

[![build](https://github.com/czech-radio/openmedia-check/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-check/actions/workflows/main.yml)

**The program reports and possibly relocates incorrectly organized OpenMedia rundowns and contacts.**

**Disclaimer:** Although we developed this package as an open-source, we use it internally to work with a specific XML file exported from the OpenMedia broadcast system. Feel free to read the source code.

## Features

- Reports wrongly placed files.
- Moves wrongly placed files in folders.
- Reports number of contacts per week.

## Installation

- Step 1
  ```bash
  git clone https://github.com/czech-radio/openmedia-check.git
  ```
- Step 2
  ```bash
  cd openmedia-check
  ```
- Step 3
  ```bash
  go mod tidy
  ```
- Step 4
  ```bash
  go build
  ```

## Usage

This program can operate only on *week* directores i.e. `W01`-`W52(3)`!
The OpenMedia export folder must accessible (should mounted if you are on Linux mahine).

Program needs `$ANNOVA` system variable and/or `.env` file which contains such variable to be filled in before run.
`$ANNOVA` should be the path of data repository root folder, the one holding sub-folders named `Contacts`, `Rundowns`.
```bash
export ANNOVA=/mnt/cro.cz/

```

If it is all set well, you can scan any system directory containing rundown files, for its stats disk operations.


The basic usage is as follows:

```bash
./openmedia-check -i "path/to/rundowns/2022/W01 path/to/contacts/2022/W02"
```

### Flags

- `-i` - The input folder or multiple input folders e.g. `"/path/to/rundowns/2022/W01 /path/to/contacts/2022/W02"`
- `-o` - The output log name (default `openmedia.log`)
- `-w` - Write changes to file system.

### Output

When it runs well, you should see something like this on the output:

```json
{"index":498,"status":"FAILURE","action":"mv","data":{"date":"2022-01-15","week":"W02","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Sobota_15_01_2022_2_9905892_20220115234504.xml","dest":"data/Rundowns/2022/W02"}}
{"index":499,"status":"SUCCESS","action":"none","data":{"date":"2022-01-06","week":"W01","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Thu__06_01_2022_2_9802737_20220106234505.xml","dest":"data/Rundowns/2022/W01"}}
{"index":500,"status":"FAILURE","action":"mv","data":{"date":"2022-01-13","week":"W02","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Thu__13_01_2022_2_9882830_20220113234503.xml","dest":"data/Rundowns/2022/W02"}}
{"index":501,"status":"SUCCESS","action":"none","data":{"date":"2022-01-04","week":"W01","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Tue__04_01_2022_2_9774691_20220104234504.xml","dest":"data/Rundowns/2022/W01"}}
{"index":502,"status":"FAILURE","action":"mv","data":{"date":"2022-01-11","week":"W02","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Tue__11_01_2022_2_9857136_20220111234504.xml","dest":"data/Rundowns/2022/W02"}}
{"index":503,"status":"SUCCESS","action":"none","data":{"date":"2022-01-05","week":"W01","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Wed__05_01_2022_2_9788319_20220105234504.xml","dest":"data/Rundowns/2022/W01"}}
{"index":504,"status":"FAILURE","action":"mv","data":{"date":"2022-01-12","week":"W02","file":"data/Rundowns/2022/W01/RD_20-24_RŽ_Sport_-_Wed__12_01_2022_2_9870175_20220112234504.xml","dest":"data/Rundowns/2022/W02"}}
```

The each line is a valid JSON object and contains the following fields:

- `index` - Sequential batch item index.
- `status` - either "SUCCESS" or "FAILURE"
- `action` - either "none", "mv", "rm"
- `data`
  - `date` - Rundown file schedule date.  
  - `week` - String representing detected week-number i.e.: `W23`.
  - `file` - Path to a source file
  - `dest` - Path to destination folder (where the file should be)

## Contribution

Propose new feature, enhance existing feature or fix a bug.


Some usefull commands:

```bash
go fmt
go vet
go test
```
