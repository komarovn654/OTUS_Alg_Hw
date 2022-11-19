#!/bin/bash

touch README.md

echo \# Test result >> README.md
echo \#\# Comparison of performance methods >> README.md

echo \#\#\# Fibonacci num calculated >> README.md
cat testfib.txt | column -t  -s '*' >> README.md

echo \#\#\# Exponentiation >> README.md
cat testpwr.txt | column -t  -s '*' >> README.md

echo \#\#\# Prime numbers count calculated >> README.md
cat testprime.txt | column -t  -s '*' >> README.md