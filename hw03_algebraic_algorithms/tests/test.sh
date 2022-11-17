#!/bin/bash

go build -o tests.exe .
rm testfib.txt testprime.txt testpwr.txt

for n in {0..12}
do
    inf=Fibo/test.$n.in
    outf=Fibo/test.$n.out
    task=fibonacci

    echo START N = $(head -1 $inf) | tee -a testfib.txt
    echo -n $task recursive: | tee -a testfib.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=recursive | tee -a testfib.txt
    echo -n fibonacci iterative: | tee -a testfib.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=iterative | tee -a testfib.txt
    echo -n fibonacci goldenratio: | tee -a testfib.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=goldenratio | tee -a testfib.txt
    echo -n fibonacci matrix: | tee -a testfib.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=matrix | tee -a  testfib.txt 
    echo -------------------------------------------------------- | tee -a testfib.txt
done

for n in {0..9}
do
    inf=Power/test.$n.in
    outf=Power/test.$n.out
    task=power

    echo START N = $(head -1 $inf) | tee -a testpwr.txt
    echo -n $task iterative: | tee -a testpwr.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=iterative | tee -a testpwr.txt
    echo -n $task sqrmultiply: | tee -a testpwr.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=sqrmultiply | tee -a testpwr.txt
    echo -n $task binary: | tee -a testpwr.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=binary | tee -a testpwr.txt
    echo -------------------------------------------------------- | tee -a testpwr.txt
done


for n in {0..14}
do
    inf=Primes/test.$n.in
    outf=Primes/test.$n.out
    task=prime

    echo START N = $(head -1 $inf) | tee -a testprime.txt
    echo -n $task brutforce: | tee -a testprime.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=brutforce | tee -a testprime.txt
    echo -n $task brutforceopt: | tee -a testprime.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=brutforceopt | tee -a testprime.txt
    echo -n $task erat: | tee -a testprime.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=erat | tee -a testprime.txt
    echo -n $task eratmem: | tee -a testprime.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=eratmem | tee -a testprime.txt
    echo -n $task eratopt: | tee -a testprime.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=eratopt | tee -a testprime.txt
    echo -------------------------------------------------------- | tee -a testprime.txt
done

rm *.exe