# musashi

## Grammar

```
  <program>    --> <list>*
  <comment>    --> #<any>\n

  <datum>      --> <boolean> | <number> | <string> | <symbol> | <list>
  <list>       --> (<datum>*)
  <boolean>    --> true | false
  <nil>        --> nil
  <number>     --> 0 | 1 | ... | 9 | 1.23
  <string>     --> "<unicode>*"
  <symbol>     --> <unicode>+ | <unicode>+(?=<string>) | <unicode>+(?=<number>) | <unicode>+(?=<comment>) 
```
