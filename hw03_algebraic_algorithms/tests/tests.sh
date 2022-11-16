go build -o tests.exe .

for n in {0..12}
do
    inf=Fibo/test.$n.in
    outf=Fibo/test.$n.out
    task=fibonacci

    echo START N = $(head -1 $inf) | tee testresult.txt
    echo -n $task recursive: | tee testresult.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=recursive | tee testresult.txt
    echo -n fibonacci iterative: | tee testresult.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=iterative | tee testresult.txt
    echo -n fibonacci goldenratio: | tee testresult.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=goldenratio | tee testresult.txt
    echo -n fibonacci matrix: | tee testresult.txt
    ./tests.exe -in=$inf -out=$outf -task=$task -alg=matrix | tee  testresult.txt 
    echo -------------------------------------------------------- | tee testresult.txt
done

rm *.exe