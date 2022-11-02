#!/bin/bash 

rm log.txt
for year in `seq 2020 2022`; do

for i in `ls /mnt/cro.cz/Rundowns/$year`;
do ./../openmedia_files_checker -i /mnt/cro.cz/Rundowns/$year/$i -o log.txt ;
done;
done
