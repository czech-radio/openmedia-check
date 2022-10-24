# rundown_files_checker
Program which founds bad modtimes in files whitin rundown repository (disk cro.cz)


# running program

```
git clone https://github.com/K0F/rundown_files_checker.git
cd rundown_files_checker
go mod tidy
go run .
```

You should have mounted "Annova" `cro.cz` smb in `/mnt/cro.cz/` or change the path in source file
