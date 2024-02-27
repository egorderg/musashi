# musashi

## Grammar

```
  <program>    --> <form>*
  <form>       --> (<datum>* | <symbol>*)
  <symbol>     --> <unicode>+ | <unicode>+(?=<string>) | <unicode>+(?=<number>) | <unicode>+(?=<comment>) 
  <comment>    --> #<any>\n

  <datum>      --> <boolean> | <int> | <float> | <string> | <nil>
  <nil>        --> nil
  <boolean>    --> true | false
  <int>        --> 0 | 1 | ... | 9 | 1.23
  <float>      --> <int>.<int>
  <string>     --> "<unicode>*"
```
