package main

import(
	 "fmt"
	 "strconv" 
	)

/// 									section 1	
//package level declerations
var name ="john"
var last = "smith"
var age =33

//group variables
var(
	name2 = "mary"
	last2 = "jane"
	age2 =28

)


func main(){
//var declarations
var a int 
a=10

var b int =20

c:= 30

// d :=40    declared but not used

fmt.Println("Hello guys")
fmt.Println(a)
fmt.Printf("%v , %T \n", b, b)
fmt.Println(c)

var s string
s = strconv.Itoa(b)//converts b int to string
fmt.Printf("%v,%T\n",s,s)


//										section 2
// boolean data

//var isTrue bool = false
is := 1 == 1
nt := 1 == 2
fmt.Printf("%v , %T \n", is, is)
fmt.Printf("%v , %T \n", nt, nt)


//integer types

var anIntStartingAtZero uint =10
fmt.Println(anIntStartingAtZero +10 )
fmt.Println(anIntStartingAtZero - 10)
fmt.Println(anIntStartingAtZero / 3)
fmt.Println(anIntStartingAtZero % 3)

// bit operators

j := 10  // 1010
i := 3	//0011
fmt.Println("_____\n")
fmt.Println(j & i)	// and
fmt.Println(j | i)	// or
fmt.Println(j ^ i)	// exclusive
fmt.Println(j &^ i)	//and not

//complex nums
var n complex64 = 1+ 2i
fmt.Printf("%v , %T \n", n, n)
fmt.Printf("%v , %T \n", real(n), real(n))
fmt.Printf("%v , %T \n", imag(n), imag(n))

// string any utf8 char
var aString string = "This a string"

// this is a bite in pos 5 will return number (ascii num)
fmt.Printf("%v , %T \n", aString[5], aString[5]) 

aBytes := []byte(aString)
fmt.Printf("%v , %T \n", aBytes,aBytes) // displays ascii values


// roon any utf32 char

var r rune = 'a'
fmt.Printf("%v , %T \n", r,r) 
					
										// section 3

 // constants

  const myConstInt int = 45
  const myConstString = "a const string"
fmt.Printf("%v , %T \n", myConstInt, myConstInt) 
fmt.Printf("%v , %T \n", myConstString,myConstString) 

fmt.Printf("%v , %T \n", myConstInt + 5, myConstInt + 5) 
 

 const (
 	 aa = iota // 0
 	 bb		   // 1
 	 cc		   // 2
 )


 const (

 	 _ = iota // throw away variable 

 	 // bits to __
 	 KB = 1 << (10 * iota) //1 << 5 is "1 times 2, 5 times" or 32
 	 					   // (1×2)^10 = 1024
 	 MG
 	 GB
 	 TB
 	 PB
 	 EB
 	 ZB
 	 YB		   // 2
 )

fmt.Println( aa,bb,cc)
fmt.Println( KB,MG,GB,TB,PB)

fileSize := 4000000.
fmt.Printf("%.2fMB\n",fileSize/MG)

//  example

								// section 3

//arrays
 
 fmt.Println("\n--- Arrays ----")


 grade := [...]int{70,80,90} //size of what i pass in

 fmt.Println(grade)

 grade[0]=100
 grade[1]=90
 grade[2]=80
 fmt.Println(grade[0])
 fmt.Println(len(grade))

// slice  
 fmt.Println("--- Slices ----")


nums := []int{10,20,30,40,50,60,70}
fmt.Printf("%v %T\n",nums,nums)

a2 := nums[:]  //slice of all num elements
b2 := nums[2:]  //slice of  3rd to end
c2 := nums[:4]  //slice of 0 to 3

 fmt.Println(a2,b2,c2)

 a3 := make([]int,3) //initilized to 0 values
 //a3 := make([]int,3,20) // slices dont need to be fixed size (20)
 fmt.Println(a3)
 fmt.Println(len(a3))
 fmt.Println(cap(a3))

// add elements
a3=append(a3,1)// when appending go will increase size
fmt.Println(a3)
fmt.Println(len(a3))
fmt.Println(cap(a3))

// slice of integers inside 1 index (spread operator ...)

 a4 := []int{} //initilized to 0 values
 //a3 := make([]int,3,20) // slices dont need to be fixed size (20)
 a4=append(a4,[]int {1,2,3,4,5,6}...)// when appending go will increase size
 fmt.Println(a4)
 fmt.Println(len(a4))
 fmt.Println(cap(a4))

 //removing from middle of slice
b4 := append(a4[:2],a4[4:]...) //a4 will change make sure we dont need a4 anymore
fmt.Println(b4)
fmt.Println(a4) //a4 array change bcuz of line 188


												// section 4

fmt.Println("\n--- Maps ----\n")


//declaring maps

 		     //key=string 
statePop := map[string]int{  // value=int
 
 		"Ca" : 392000,
 		"Az" : 292000,
 		"Dv" : 39200,
 		"La" : 112000,
 }

fmt.Println(statePop)

fmt.Println(statePop["Ca"])

statePop["La"] = 9900 // replacing value of'La'
delete(statePop,"Dv") //removing key

fmt.Println(statePop)

//if key not found will return 0 , 0 is not good result
pop ,ok := statePop["Ka"] //if not found will return false
fmt.Println(pop,ok)
//

}