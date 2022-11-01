#!/bin/bash 

rm log.txt
for i in `ls /mnt/cro.cz/Rundowns/2022`;
do ./../openmedia-files-checker -i /mnt/cro.cz/Rundowns/2022/$i -o log.txt ;
done
