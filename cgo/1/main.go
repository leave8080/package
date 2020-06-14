package main

/*
#include <stdio.h>
void Print(int i)
{
	printf("%d\n",i);
}
*/
import "C"

func main() {
	C.Print(1)
}
