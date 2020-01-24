package big_o

/*
Master in Coding 
Section 3
All Big o Type
1) o(1)
2) o(lon n)
3) o(n)
4) o(n*log(n))
5) o(n^2)
6) o(2^n)
7) o(n!)   factorial times 	
https://www.bigocheatsheet.com/  here you get all notation with type

Which code is best? or pillar of Programming
	- Readable
	- Memory (Space Complexity)
	- Speed (Time Complexity)

What is cause time in function
	- Operation (+ - * /)
	- Comparisons (< > =)
	- Looping (for While)
	- Outside function call

What cause of space complexity? 
	- Variable
	- Data Structure
	- Function Call
	- Allocations	
	

	*/







//Big O o(n) notation example
var numberArray =[]int {1,2,3,4,5,6,7,8,9,10}
func notationO_n(numberArray){
	for _,v:=range numberArray{  //O(n) n is lenght of numberArray
		fmt.Println("value",v)
	}
}


//ans: O(n)

//Big O o(1) notation example
func notationO_n(numberArray){
		fmt.Println("value",numberArray[1]) //o(1)  this operation not dependent on array lenght because its execute only one time
	
}

//ans: O(1)


//exercise 1:
//calculate Big O notation

func getData(numberArray){
	a:=5     
	a:=5+10
	for _,v:=range numberArray{ 
		notationO_n(numberArray)
		fmt.Println("value",v)
		a++
	}
	fmt.Println("string")
}


/*Answer: 

func getData(numberArray){
	a:=5      o(1)
	a:=5+10 o(1)
	for _,v:=range numberArray{ 
		notationO_n(numberArray)  o(n)
		fmt.Println("value",v)  o(n)
		a++  o(n)
	} 
	fmt.Println("string")  o(1)
}

Big o(3+3n)
but in interview answer only o(n)

*/

func notationO_n^2(){
	numArray:=[]int{1,2,3,4,5}
	//int that code nexted array is present 
	for i:=0;i<len(numArray);i++{  // o(n)
		for j:=0;j<len(numArray);j++{ //o(n)
			fmt.Println(numArray[i],numArray[j])
		}
	}
}

//Answer o(n*n)  its called as o(n^2)

func notationO_n(){ //example second
	numArray:=[]int{1,2,3,4,5}
	//int that code nexted array is present 
	for i:=0;i<len(numArray);i++{  // o(n)
			fmt.Println(numArray[i])
		
	}
	numArray1:=[]int{1,2,3}
	for j:=0;j<len(numArray1);j++{ //o(n)
		fmt.Println(numArray1[j])
	}
}

//Answer o(n+n)  //its called as o(n)