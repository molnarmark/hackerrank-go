#!/bin/bash

read answer

if [ $answer == "y" ] || [ $answer == "Y" ]
then
  echo "YES"
else
  echo "NO"
fi