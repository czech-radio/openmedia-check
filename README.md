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


when it runs well, you should see something like this on the output:
```
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
