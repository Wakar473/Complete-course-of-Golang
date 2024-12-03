package main

import("fmt")
 func main(){
	for i := 1 ; i <= 10;i++ {

		if i %2 ==0{
			fmt.Println(i,"even")
		
		}else{
			fmt.Println(i,"odd")
		}
	}
 }

/*
Let's walk through this program:
• Create a variable i of type int and give it the value 1
• Is i less than or equal to 10 ? Yes: jump to the block
• Is the remainder of i ÷ 2 equal to 0 ? No: jump to the else block
• Print i followed by odd
• Increment i (the statement after the condition)
• Is i less than or equal to 10 ? Yes: jump to the block
• Is the remainder of i ÷ 2 equal to 0 ? Yes: jump to the if block
• Print i followed by even */