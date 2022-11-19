# Test result
## Comparison of performance methods
### Fibonacci num calculated
|              | recursive      | iterative          | goldenratio   | matrix             |
| ------       | ------         | ------             | ------        | ------             |
|N = 0         | PASS_981ns     | PASS_1.192µs       | PASS_661ns    | PASS_331ns         |
|N = 1         | PASS_1.212µs   | PASS_632ns         | PASS_941ns    | PASS_3.186µs       |
|N = 2         | PASS_912ns     | PASS_851ns         | PASS_1.743µs  | PASS_12.283µs      |
|N = 3         | PASS_1.873µs   | PASS_2.054µs       | PASS_1.754µs  | PASS_15.399µs      |
|N = 4         | PASS_1.643µs   | PASS_892ns         | PASS_801ns    | PASS_6.914µs       |
|N = 5         | PASS_2.545µs   | PASS_2.866µs       | PASS_1.953µs  | PASS_16.01µs       |
|N = 10        | PASS_11.962µs  | PASS_2.855µs       | PASS_1.773µs  | PASS_10.86µs       |
|N = 100       | TIMEOUT        | PASS_7.724µs       | FAIL 1.533µs  | PASS_33.934µs      |
|N = 1000      | TIMEOUT        | PASS_94.629µs      | FAIL 2.024µs  | PASS_82.706µs      |
|N = 10000     | TIMEOUT        | PASS_1.651011ms    | FAIL 1.824µs  | PASS_145.756µs     |
|N = 100000    | TIMEOUT        | PASS_48.85116ms    | FAIL 2.565µs  | PASS_4.843346ms    |
|N = 1000000   | TIMEOUT        | PASS_4.347169518s  | FAIL 2.334µs  | PASS_90.09826ms    |
|N = 10000000  | TIMEOUT        | TIMEOUT            | FAIL 1.453µs  | PASS_2.689663942s  |
### Exponentiation
|                  | iterative           | sqrmultiply        | binary      |
| ------           | ------              | ------             | ------      |
|N = 2             | PASS_241ns          | PASS_371ns         | PASS_140ns  |
|N = 123456789     | PASS_271ns          | PASS_260ns         | PASS_110ns  |
|N = 1.001         | PASS_2.915µs        | PASS_2.746µs       | PASS_150ns  |
|N = 1.0001        | PASS_14.518µs       | PASS_18.906µs      | PASS_170ns  |
|N = 1.00001       | PASS_143.481µs      | PASS_212.141µs     | PASS_421ns  |
|N = 1.000001      | PASS_3.162769ms     | PASS_2.302854ms    | PASS_531ns  |
|N = 1.0000001     | PASS_14.040124ms    | PASS_6.856433ms    | PASS_220ns  |
|N = 1.00000001    | PASS_108.521489ms   | PASS_76.646015ms   | PASS_280ns  |
|N = 1.000000001   | PASS_1.044088092s   | PASS_771.391798ms  | PASS_471ns  |
|N = 1.0000000001  | PASS_10.337684479s  | PASS_5.880348324s  | PASS_711ns  |
### Prime numbers count calculated
|                | brutforce         | brutforceopt        | erat                | eratmem             | eratopt           |
| ------         | ------            | ------              | ------              | ------              | ------            |
|N = 10          | PASS_932ns        | PASS_4.909µs        | PASS_1.012µs        | PASS_752ns          | PASS_1.693µs      |
|N = 1           | PASS_250ns        | PASS_120ns          | PASS_271ns          | PASS_241ns          | PASS_421ns        |
|N = 2           | PASS_791ns        | PASS_440ns          | PASS_80ns           | PASS_441ns          | PASS_351ns        |
|N = 3           | PASS_381ns        | PASS_5.35µs         | PASS_632ns          | PASS_1.062µs        | PASS_1.132µs      |
|N = 4           | PASS_410ns        | PASS_431ns          | PASS_671ns          | PASS_772ns          | PASS_380ns        |
|N = 5           | PASS_471ns        | PASS_5.971µs        | PASS_531ns          | PASS_572ns          | PASS_1.162µs      |
|N = 100         | PASS_13.556µs     | PASS_9.708µs        | PASS_1.082µs        | PASS_1.533µs        | PASS_6.182µs      |
|N = 1000        | PASS_408.762µs    | PASS_64.672µs       | PASS_14.929µs       | PASS_26.971µs       | PASS_15.339µs     |
|N = 10000       | PASS_28.817292ms  | PASS_688.702µs      | PASS_67.227µs       | PASS_154.232µs      | PASS_91.092µs     |
|N = 100000      | PASS_1.70289677s  | PASS_5.372236ms     | PASS_341.265µs      | PASS_563.275µs      | PASS_1.336226ms   |
|N = 1000000     | TIMEOUT           | PASS_78.782641ms    | PASS_6.346719ms     | PASS_5.73288ms      | PASS_7.623172ms   |
|N = 10000000    | TIMEOUT           | PASS_1.560180989s   | PASS_77.795698ms    | PASS_70.098064ms    | PASS_73.945879ms  |
|N = 100000000   | TIMEOUT           | PASS_35.366058253s  | PASS_1.369013744s   | PASS_1.387902429s   | PASS_739.86973ms  |
|N = 1000000000  | TIMEOUT           | TIMEOUT             | PASS_14.814265515s  | PASS_19.697557496s  | PASS_7.73532167s  |
|N = 123456789   | TIMEOUT           | PASS_46.819019378s  | PASS_1.660091147s   | PASS_1.859342747s   | PASS_887.80268ms  |
