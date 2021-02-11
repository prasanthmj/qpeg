# QPeg
A simple query parser using PEG grammar. 

The syntax of the query is like this :
```
item.name=laptop item.spec.ram > 8gb item.spec.ssd=yes item.spec.ssd.capacity > 512gb sort_by:price/asc
```
or 
```
item.name=laptop (item.maker=asus | item.maker=coconics) item.spec.ssd=yes
```

### The PEG grammar
See the PEG grammar file here: /qp/query.peg

Command to generate the query.go file:
```
cd qp
pigeon -o query.go query.peg
```