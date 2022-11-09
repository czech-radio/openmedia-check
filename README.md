# openmedia-files-checker

**The program finds wrongly organized OpenMedia rundown files within the given directory.**

## Features

- Reports wrongly placed files.
- Moves wrongly placed files in folders.
- Reports number of contacts per week.

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

## Usage

The "OpenMedia export folder must be mounted (accessible) `/xyz/cro.cz/Rundowns/2022/W01`.
You can change the path to the directory beginning with `WXX`.

```bash
./openmedia-files-checker -i "/path/to/mounted/Rundowns/2022/W01 /path/to/mounted/Rundowns/2022/W33" [optional -o log.txt] [optional -w write changes]
```
## Flags

- `-i` - input folder which ends with `WXX` or multiple input folders in doublequotes `"/path/to/2022/201 /path/to/2022/W02"`
- `-o` - output log file in json format ie `log.json`
- `-c` - check contact counts in files (can be slow)
- `-w` - do changes to filesystem

## Error messages
thy contains four fileds
- `date` - timestamp of when it occurs
- `type` - can be either `info`, `error` or `warning`
- `message` - human readable message what happened
- `status` - can be either `0` - info, `1` - error, `2` - warning (usually when file operations are done) this should be machine readable

When it runs well, you should see something like this on the output:

```json
{
    "date": "2022-11-09 10:44:27.1448431 +0100 CET m=+488.826201429",
    "type": "info",
    "message": "/mnt/cro.cz/Rundowns/2022/W01 test result: 498/498 SUCCESS!",
    "status": 0
}
```
