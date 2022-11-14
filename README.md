# openmedia-check

[![build](https://github.com/czech-radio/openmedia-check/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-check/actions/workflows/main.yml)

**The program reports and possibly correct incorrectly organized OpenMedia rundonws and contacts.**

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
We often refer to this path via environment variable `$env:ANNOVA` e.g `$env:ANNOA\Rundowns\2022\W01`

The basic usage is as follows:

```bash
./openmedia-check -i "$ANNOVA/path/to/rundowns/2022/W01 $ANNOVA/path/to/contacts/2022/W02"
```

### Flags

- `-i` - The input folder or multiple input folders e.g. `"/path/to/rundowns/2022/W01 /path/to/contacts/2022/W02"`
- `-o` - The output log name (default `openmedia.log`)
- `-c` - Check contact counts in files.
- `-w` - Write changes to file system.

### Output

When it runs well, you should see something like this on the output:

```json
{"index":0,"status":"SUCCESS","data":{"date":"2022-11-4","week":"44","file":"RD_00-05_Radiožurnál_-_Fri__04_11_2022_2_13519620_20221105001439.xml"}}
{"index":1,"status":"SUCCESS","data":{"date":"2022-10-31","week":"44","file":"RD_00-05_Radiožurnál_-_Mon__31_10_2022_2_13467409_20221101001437.xml"}}
{"index":2,"status":"SUCCESS","data":{"date":"2022-11-6","week":"44","file":"RD_00-05_Radiožurnál_-_Neděle_06_11_2022_2_13547024_20221107001352.xml"}}
{"index":3,"status":"SUCCESS","data":{"date":"2022-11-5","week":"44","file":"RD_00-05_Radiožurnál_-_Sobota_05_11_2022_2_13537307_20221106001425.xml"}}
{"index":4,"status":"SUCCESS","data":{"date":"2022-11-1","week":"44","file":"RD_00-05_Radiožurnál_-_Tue__01_11_2022_2_13478904_20221102001422.xml"}}
{"index":5,"status":"SUCCESS","data":{"date":"2022-11-2","week":"44","file":"RD_00-05_Radiožurnál_-_Wed__02_11_2022_2_13493128_20221103001430.xml"}}
{"index":6,"status":"SUCCESS","data":{"date":"2022-11-3","week":"44","file":"RD_00-05_Radiožurnál_-__Čt_03_11_2022_2_13506313_20221104001434.xml"}}
{"index":7,"status":"SUCCESS","data":{"date":"2022-11-6","week":"44","file":"RD_00-05_ČRo_Region_SC_-_Neděle_06_11_2022_2_13546661_20221107001347.xml"}}
{"index":8,"status":"SUCCESS","data":{"date":"2022-10-31","week":"44","file":"RD_00-05_ČRo_Region_SC_-_Pondělí_31_10_2022_2_13467101_20221101001433.xml"}}
{"index":9,"status":"SUCCESS","data":{"date":"2022-11-4","week":"44","file":"RD_00-05_ČRo_Region_SC_-_Pátek_04_11_2022_2_13519355_20221105001432.xml"}}
```

Each line is a JSON object and contains the following fields:

- `#` - Sequential batch item index.
- `date` - Timestamp of when it occurs.
- `level` - Can be either `info`, `error` or `warning`.
- `status` - Can be either `0` - info, `1` - error, `2` - warning.
- `message` - Human readable status message.


## Contribution

Propose new feature, enhance existing feature or fix a bug.


Some usefull commands:

```bash
go fmt
go vet
go test
```
