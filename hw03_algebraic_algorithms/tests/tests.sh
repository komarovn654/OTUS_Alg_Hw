go build -o tests.exe .

for n in {0..13}
do
    ./tests.exe -in=test.$n.in -out=test.$n.out -task=fibonacci -alg=recursive   
done

rm *.exe