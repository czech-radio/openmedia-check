# openmedia-check

**The program finds wrongly organized OpenMedia rundown files within the given directory.**

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
  go build -o openmedia-check[.exe]
  ```

## Usage

**Prerequisites**

The "OpenMedia export folder must be mounted (accessible) `/xyz/cro.cz/Rundowns/2022/W01`.
You can change the path to the directory beginning with `WXX`.

```bash
./openmedia-check -i "/path/to/rundowns/2022/W01 /path/to/Rundowns/2022/W32" [-o <output_name>] [-w] [-c]
```

### Flags

- `-i` - The input folder ending with `WXX` or multiple input folders in doublequotes e.g. `"/path/to/rundowns/2022/W01 /path/to/rundowns/2022/W02"`
- `-o` - The output log name (default `openmedia.log`)
- `-c` - Check contact counts in files.
- `-w` - Write changes to file system.

### Errors

JSON message contains the following fields:
- `date` - timestamp of when it occurs
- `type` - can be either `info`, `error` or `warning`
- `message` - human readable message what happened
- `status` - can be either `0` - info, `1` - error, `2` - warning (usually when file operations are done) this should be machine readable

When it runs well, you should see something like this on the output:

```json
{
    "date": "2022-11-09T10:44:27",
    "level": "info",
    "message": "/mnt/cro.cz/Rundowns/2022/W01 test result: 498/498 SUCCESS!",
    "status": 0
}
```
