# coder
CLI for encoding strings according to 
**substitute rules:**

- `a -> 1`
- `e -> 2`
- `i -> 3`
- `o -> 4`
- `u -> 5`


---
How to install:  
---
`$ go build`

---
How to use :
---

`$ ./coder.go --h`  
will show help-menu

`$ ./coder -encode "hello"`
<br/> will return `"h2ll4"`

`$ ./coder -decode "h3 th2r2"`
<br/> - `"hi there"`

*letter case is ignored in a way :  
'A' and 'a' both encoding into 1*  
*while 1 is only decoded into 'a'*