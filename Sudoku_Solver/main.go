package main

import(
	"fmt"
	"runtime"
)
//核心填空函数:尝试填入-递归-失败-撤销
func solve(board *[9][9]int)bool{
	for row:=0;row<9;row++{
		for col:=0;col<9;col++{
			//找到一个空位置
			if board[row][col]!=0{
				continue
			}
			//尝试1-9
			for num:=1;num<10;num++{
				//判断是否符合
				if !isValid(board,row,col,num){
					continue
				}
				//填入
				board[row][col]=num
				//递归解决
				if solve(board){
					return true
				}
				//失败
				board[row][col]=0
			}
			return false
		}
	}
	return true
} 
//检查行 列 3*3
func isValid(board *[9][9]int,row int ,col int,num int)bool{
	for i:=0;i<9;i++{
		if board[row][i]==num{
			return false
		}
	}
	for i:=0;i<9;i++{
		if board[i][col]==num{
			return false
		}
	}
	startRow:= row/3*3
	startCol:= col/3*3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if board[startCol+i][startRow+j]==num{
				return false
			}
		}
	}
	return true
}

func printBoard(board *[9][9]int){
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			fmt.Println("%d",board[i][j])
		}
	}
}


func main(){
	var board=[9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if solve (&board){
		printBoard(&board)
	}else{
		fmt.Println("无解")
	}

}