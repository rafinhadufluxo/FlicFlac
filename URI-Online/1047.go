package main

import (
	"fmt"
)

func main() {

	//inicializando as variaveis

	var a int 
	var b int
	var c int 
	var d int
	var x int 
	var y int
	
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%d\n", &b )
	fmt.Scanf("%d\n", &y)
	
	if (a == y && b == y && y == x){
		c=x-y
		d=24+a-b
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c ,"MINUTO(S)\n")
	}
	if ( a == b && y > x) {
		c=y-x
		d=a-b
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if ( a == b && x > y){
		c=60-x+y
		d=24-a+b-1
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (x == y && a < b){
		c=0
		d=b-a
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (x == y && a > b){
		c=0
		d=24-a+b
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (b > a && y > x){
		c=y-x
		d=b-a
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (a < b && x > y){
		c=60-x+y
		d=b-a-1
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (a > b && x < y){
		c = y - x
		d = 24 - a -1 + b
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
	if (a > b && x > y){
		c= 60 + y - x
		d= 24 + b - a - 1
		fmt.Println("O JOGO DUROU", d ,"HORA(S) E ", c , "MINUTO(S)\n")
	}
}