# coder
CLI for encoding strings according to 
**substitute rules:**

- `a -> 1`
- `e -> 2`
- `i -> 3`
- `o -> 4`
- `u -> 5`

---
How to use :
---

`$ go run cmd/coder.go --h`  
will show help-menu

`$ go run cmd/coder.go -encode "hello"`
<br/> will return `"h2ll4"`

`$ go run cmd/coder.go -decode "h3 th2r2"`
<br/> - `"hi there"`

*letter case is ignored in a way :  
'A' and 'a' both encoding into 1*  
*while 1 is only decoded into 'a'*