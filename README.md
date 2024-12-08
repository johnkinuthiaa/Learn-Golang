# Learning Golang using java and javascript concepts

loops -used fizzbuzz for for loops -
loops are almost the same with java but lil different in the initialization
### in javascript
```javascript
for(let i=0; i<=100 ;i++){
    if(i%5 === 0 && i%3 === 0){
        console.log("fizzbuzz")
    }else if (i%5 === 0){
        console.log("fizz")
    }else if(i%3===0){
        console.log("Buzz")
    }else{
        console.log(i)
    }
}
```
### in java
```javascript
for (int i=0;i<=100;i++){
    if(i%5==0 && i%3==0){
        System.out.println("fizzBuzz");
    }else if(i %5==0){
        System.out.println("fizz");
    }else if(i %3 ==0){
        System.out.println("Buzz");
    }else{
        System.out.println(i);
    }
}
```

### in Go,the only thing that changes is how we initialize the i or loop starting point
```go
//use i :=0 or var i=0
for i :=0; i<=100; i++ {
	
    if i%5==0 && i%3==0{
        fmt.println("fizzBuzz");
		
    }else if i %5==0{
        fmt.println("fizz");
		
    }else if i %3 ==0{
        fmt.println("Buzz");
		
    }else{
		fmt.println(i);}
}
```