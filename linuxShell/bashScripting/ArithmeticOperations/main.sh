#!/bin/bash
read expression

printf "%0.3f\n" `echo "$expression" | bc -l`