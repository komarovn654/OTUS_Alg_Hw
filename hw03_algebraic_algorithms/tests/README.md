# Test result
## Comparison of performance methods
### Fibonacci num calculated
|              | recursive       | iterative          | goldenratio         | matrix              |
| ------       | ------          | ------             | ------              | ------              |
|N = 0         | PASS_117.612µs  | PASS_95.962µs      | PASS_85.111µs       | PASS_115.749µs      |
|N = 1         | PASS_80.041µs   | PASS_94.93µs       | PASS_63.169µs       | PASS_60.294µs       |
|N = 2         | PASS_59.232µs   | PASS_51.939µs      | PASS_119.806µs      | PASS_87.605µs       |
|N = 3         | PASS_131.719µs  | PASS_92.656µs      | PASS_94.209µs       | PASS_218.403µs      |
|N = 4         | PASS_81.805µs   | PASS_59.934µs      | PASS_58.44µs        | PASS_68.399µs       |
|N = 5         | PASS_65.905µs   | PASS_72.246µs      | PASS_131.548µs      | PASS_125.447µs      |
|N = 10        | PASS_114.276µs  | PASS_125.838µs     | PASS_67.618µs       | PASS_124.696µs      |
|N = 100       | TIMEOUT         | PASS_65.273µs      | FAIL 41.698µs       | PASS_134.554µs      |
|N = 1000      | TIMEOUT         | PASS_209.265µs     | FAIL 131.148µs      | PASS_185.481µs      |
|N = 10000     | TIMEOUT         | PASS_1.4591ms      | FAIL 232.84µs       | PASS_242.188µs      |
|N = 100000    | TIMEOUT         | PASS_53.677085ms   | FAIL 2.745723ms     | PASS_7.589403ms     |
|N = 1000000   | TIMEOUT         | PASS_4.405196943s  | FAIL 117.956467ms   | PASS_198.9759ms     |
|N = 10000000  | TIMEOUT         | TIMEOUT            | FAIL 11.415597004s  | PASS_13.383815502s  |
### Exponentiation
|                  | iterative          | sqrmultiply        | binary          |
| ------           | ------             | ------             | ------          |
|N = 2             | PASS_117.222µs     | PASS_58.521µs      | PASS_122.141µs  |
|N = 123456789     | PASS_58.831µs      | PASS_48.772µs      | PASS_100.59µs   |
|N = 1.001         | PASS_93.106µs      | PASS_147.84µs      | PASS_141.087µs  |
|N = 1.0001        | PASS_136.068µs     | PASS_119.145µs     | PASS_102.283µs  |
|N = 1.00001       | PASS_464.929µs     | PASS_297.813µs     | PASS_56.877µs   |
|N = 1.000001      | PASS_1.510737ms    | PASS_2.318855ms    | PASS_52.489µs   |
|N = 1.0000001     | PASS_17.025097ms   | PASS_12.922099ms   | PASS_108.695µs  |
|N = 1.00000001    | PASS_124.733172ms  | PASS_78.59152ms    | PASS_46.859µs   |
|N = 1.000000001   | PASS_1.039841477s  | PASS_753.777561ms  | PASS_50.214µs   |
|N = 1.0000000001  | PASS_10.3528082s   | PASS_5.920624706s  | PASS_127.542µs  |
### Prime numbers count calculated
|                | brutforce          | brutforceopt        | erat                | eratmem             | eratopt            |
| ------         | ------             | ------              | ------              | ------              | ------             |
|N = 10          | PASS_105.81µs      | PASS_75.513µs       | PASS_122.391µs      | PASS_95.441µs       | PASS_69.982µs      |
|N = 1           | PASS_61.527µs      | PASS_139.183µs      | PASS_48.061µs       | PASS_55.706µs       | PASS_128.292µs     |
|N = 2           | PASS_57.348µs      | PASS_110.859µs      | PASS_106.461µs      | PASS_111.681µs      | PASS_151.797µs     |
|N = 3           | PASS_117.392µs     | PASS_46.237µs       | PASS_42.29µs        | PASS_47.45µs        | PASS_105.359µs     |
|N = 4           | PASS_63.12µs       | PASS_54.212µs       | PASS_54.222µs       | PASS_90.15µs        | PASS_49.413µs      |
|N = 5           | PASS_113.104µs     | PASS_61.115µs       | PASS_120.618µs      | PASS_101.763µs      | PASS_126.99µs      |
|N = 100         | PASS_72.868µs      | PASS_104.268µs      | PASS_158.9µs        | PASS_53.111µs       | PASS_54.102µs      |
|N = 1000        | PASS_460.1µs       | PASS_82.466µs       | PASS_106.331µs      | PASS_63.15µs        | PASS_131.108µs     |
|N = 10000       | PASS_26.198386ms   | PASS_293.064µs      | PASS_71.195µs       | PASS_220.988µs      | PASS_276.854µs     |
|N = 100000      | PASS_1.712108484s  | PASS_4.056491ms     | PASS_282.635µs      | PASS_596.858µs      | PASS_1.371514ms    |
|N = 1000000     | TIMEOUT            | PASS_74.892845ms    | PASS_6.025575ms     | PASS_5.791553ms     | PASS_10.941874ms   |
|N = 10000000    | TIMEOUT            | PASS_1.604447717s   | PASS_67.400655ms    | PASS_66.08686ms     | PASS_76.822664ms   |
|N = 100000000   | TIMEOUT            | PASS_35.774089778s  | PASS_1.491373443s   | PASS_1.517258255s   | PASS_779.795406ms  |
|N = 1000000000  | TIMEOUT            | TIMEOUT             | PASS_14.413052044s  | PASS_19.107394859s  | PASS_7.507332884s  |
|N = 123456789   | TIMEOUT            | PASS_47.003690142s  | PASS_1.732093739s   | PASS_1.946994423s   | PASS_953.576684ms  |
