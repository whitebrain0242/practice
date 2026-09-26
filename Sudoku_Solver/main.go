package main

import "fmt"

func solve(board *[9][9]int)bool  {
	for row:=0;row<9;row++{
		for col:=0;col<9;col++{
			//找到一个空格
			if board[row][col]!=0{
				continue
			}
			for num:=1;num<=9;num++{
				//判断可行性
				if !isValid(board,row,col,num){
					continue
				}
				//填入
				board[row][col]=num
				//递归解决其他部分
				if solve(board){
					return true
				}
				board[row][col]=0
			}
			//1-9都不行
			return false
		}
	}
	//没有找到空格
	return true
}
func isValid(board *[9][9]int,row int,col int,num int)bool{
	for i:=0;i<9;i++{
		if board[row][i]==num{
			return  false
		}
	}
	for i:=0;i<9;i++{
		if board[i][col]==num{
			return false
		}
	}
	var startRow=row/3*3
	var startCol=col/3*3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if board[startRow+i][startCol+j]==num{
				return false
			}
		}
	}
	return true
}
func printBoard(board *[9][9]int)bool{
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			fmt.Print(board[i][j]," ")
		}
		fmt.Println()
	}
	return true
}
func main()  {
    var board [9][9]int
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			fmt.Scan(&board[i][j])
		}
	}
	if solve(&board){
		printBoard(&board)
	}else{
		fmt.Println("无解")
	}

}