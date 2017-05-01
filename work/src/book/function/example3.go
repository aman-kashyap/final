package main

import "fmt"

/*func half(x int) (int,bool){
	if(x%2==0){
		return 1,true
	}else{
		return 0,false
	}
	
}
func main(){
	fmt.Println(half(10))

}*/
/*func greatest(args ...int) int{
	x:=args[2]
	for _, v:= range args{
		if v>x{
			x=v
		}
	}
	return x
}
func main(){
	
	fmt.Println(greatest(1,2,3,4,5,6,45,90))
}*/
/*func makeOddGenerator() func() uint{
	i := uint(1)
	return func() (ret uint) {
		ret= i
		i+=2
		return
	}
}
func main(){
	nextOdd :=makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	
}*/
func fib(x uint) uint{
	if x==0{
		return 0
	}
	if x==1{
		return 1
	}
	if (fib(x)%2==0){
	return fib(x-1)+fib(x-2)
	}
}
func main(){
	fmt.Println(fib(7))
}