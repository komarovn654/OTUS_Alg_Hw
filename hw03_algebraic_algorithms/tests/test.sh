#!/bin/bash

function runTest {
    if !(./tests.exe $1 $2 $3 $4 $5 ) 
    then exit 1 
    fi
}

function StartFibonacciTest {
    resf=testfib.txt
    task=fibonacci
    
    touch $resf
    echo \|*\| recursive*\| iterative*\| goldenratio*\| matrix*\| >> $resf
    echo \| ------ *\| ------ *\| ------ *\| ------ *\| ------ *\| >> $resf
    for n in {0..12}
    do
        inf=Fibo/test.$n.in
        outf=Fibo/test.$n.out

        echo N = $(head -1 $inf)
        echo -n \|N = $(head -1 $inf)* >> $resf
        runTest $inf $outf $task recursive $resf
        runTest $inf $outf $task iterative $resf
        runTest $inf $outf $task goldenratio $resf
        runTest $inf $outf $task matrix $resf
        echo \| >> $resf
    done
}

function StartPowerTest {
    resf=testpwr.txt
    task=power
    
    touch $resf
    echo \|*\| iterative*\| sqrmultiply*\| binary*\| >> $resf
    echo \| ------ *\| ------ *\| ------ *\| ------ *\| >> $resf
    for n in {0..9}
    do
        inf=Power/test.$n.in
        outf=Power/test.$n.out

        echo N = $(head -1 $inf)
        echo -n \|N = $(head -1 $inf)* >> $resf
        runTest $inf $outf $task iterative $resf
        runTest $inf $outf $task sqrmultiply $resf
        runTest $inf $outf $task binary $resf
        echo \| >> $resf
    done
}

function StartPrimeTest {
    resf=testprime.txt
    task=prime
    
    touch $resf
    echo \|*\| brutforce*\| brutforceopt*\| erat*\| eratmem*\| eratopt*\| >> $resf
    echo \| ------ *\| ------ *\| ------ *\| ------ *\| ------ *\| ------ *\| >> $resf
    for n in {0..14}
    do
        inf=Primes/test.$n.in
        outf=Primes/test.$n.out

        echo N = $(head -1 $inf)
        echo -n \|N = $(head -1 $inf)* >> $resf
        runTest $inf $outf $task brutforce $resf
        runTest $inf $outf $task brutforceopt $resf
        runTest $inf $outf $task erat $resf
        runTest $inf $outf $task eratmem $resf
        runTest $inf $outf $task eratopt $resf
        echo \| >> $resf
    done
}

rm testfib.txt testprime.txt testpwr.txt README.md
go build -o tests.exe . 

echo start fibonacci test ...
StartFibonacciTest
echo start power test ...
StartPowerTest
echo start prime test ...
StartPrimeTest

bash makeRM.sh

rm testfib.txt testprime.txt testpwr.txt *.exe