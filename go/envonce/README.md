* Mimic test for net/http/transport.go

~~~
$ foo=1 bar=2 go run envOnce.go 
-- foo -- 
value =  1 
value =  1 
value =  1 
value =  1
~~~
