#!/bin/bash

read x
read y
read z

sides=($x $y $z)

if [ $y == $x ] && [ $z == $x ]
then
  echo "EQUILATERAL"
elif [ $x == $y ] || [ $x == $z ] || [ $y == $z ]
  then
    echo "ISOSCELES"
else
    echo "SCALENE"
fi