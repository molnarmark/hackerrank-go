#!/bin/bash

nums=0

read totalNumbers
for i in $(seq 0 $totalNumbers)
do
  read num
  nums=$((nums + num))
done
printf "%.3f\n" `echo "$nums/$totalNumbers" | bc -l`