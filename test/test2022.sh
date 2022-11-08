#!/bin/bash 

rm log.json
for year in `seq 2020 2022`; do

for i in `ls /mnt/cro.cz/Rundowns/$year`;
do ./../openmedia_files_checker -i /mnt/cro.cz/Rundowns/$year/$i -o log.json -w ;
done;
done

sed -e 's/\]\[/\,/g' log.json
