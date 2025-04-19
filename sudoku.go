package main

import(
	"fmt"
	"os"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

//which big box
func boxnumber(i int, j int) int{
	return int(i/3)*3 + int(j/3);
}

//number within big box
func boxcoord(i int, j int) int{
	return 3*(i%3) + (j%3);
}

func sum(numbers []int) int{
    var total int = 0;
    for _, num := range numbers {
        total += num
    }
    return total
}

//check if a vector (of 1s ans 0s) 
//sums to 1 and if yes where the 1 is
func checkOne(numbers []int) (bool, int){
	var dtm bool = false;
	var pos int = 10;
	for ind, num := range numbers {
		if num == 1 && dtm == true{
			return false, 10;
		}
		if num == 1{
			pos = ind;
			dtm = true;
		}
	}
	return dtm, pos;
}

func stepfunc(num int) int{
	if num > 0{
		return 1;
	}
	return 0;

}

/* box is 9x9x9 array that summarises 
   the information of board in a different way:
   
   for each 0<=i,j <9, box[i][j] is an 
   array of 1s and 0s. 

   box[i][j][k] is 1 if there is a 
   possibility that the number j + 1 
   in the ith big box can be placed 
   in position k there, given what is 
   known thus far. 
   box[i][j][k] is zero otherwise. 
   
   When all 81 arrays 
   box[i][j] have only one 1 and 8 0s, 
   then the board is solved.

   positions indices of big and small boxes:
	 _  _  _
   |0_|1_|2_|
   |3_|4_|5_|
   |6_|7_|8_|

*/

func parseBoard(board [9][9]int, counter int) ([9][9][9]int, int){
	var box [9][9][9]int;
	for i := 0; i < 9; i++{
		for j:= 0;j < 9; j++ {
			if board[i][j]!=0 {
				box[boxnumber(i,j)][board[i][j] - 1][boxcoord(i,j)] = 1;
				counter += 1;
			}
			// else{
			// 	for k:= 0; k < 9; k++{
			// 		box[boxnumber(i,j)][j][k] = 1;
			// 	}
			// }

		}
	}

	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if sum(box[i][j][:]) == 0{
				for k:= 0; k < 9; k++{
					box[i][j][k] = 1;
				}
			}

		}
	}
	return box, counter;
}

func parseBox(box [9][9][9]int) [9][9]int{
	var fin_board [9][9]int;
	var pos int;
	var dtm bool = false;
	
	for i := 0; i < 9; i++{
		for j:= 0;j < 9; j++ {
			dtm, pos = checkOne(box[i][j][:]);
			if dtm {
				fin_board[int(i/3)*3 + int(pos/3)][(i%3)*3 + pos%3] = j + 1;
			}
		}
	}

	return fin_board;
}

func checkcolumn(box *[9][9][9]int, i int, j int) {
	//index of other big boxes in the same column
	var x int = (i + 3)%9;
	var y int = (i + 6)%9;
	
	var count0 int = 0;
	var count1 int = 0;
	var count2 int = 0;

	var pos int;
	var dtm bool;

	//if j + 1 in box i is determined, do nothing
	if sum(box[i][j][:]) == 1{
		return;
	}

	for k:= 0; k < 3; k++{
		dtm, pos = checkOne(box[x][j][:]);
		if dtm {
			box[i][j][pos%3] = 0;
			box[i][j][pos%3 + 3] = 0;
			box[i][j][pos%3 + 6] = 0;
		}
		dtm, pos = checkOne(box[y][j][:]);
		if dtm {
			box[i][j][pos%3] = 0;
			box[i][j][pos%3 + 3] = 0;
			box[i][j][pos%3 + 6] = 0;
		}
		
	}

	for k:= 0; k < 3; k++{
		count0 += box[x][j][3*k];
		count1 += box[x][j][3*k + 1];
		count2 += box[x][j][3*k + 2];
		count0 += box[y][j][3*k];
		count1 += box[y][j][3*k + 1];
		count2 += box[y][j][3*k + 2];
	}

	count0 = stepfunc(count0);
	count1 = stepfunc(count1);
	count2 = stepfunc(count2);


	if count0 + count1 + count2 == 2{
		if count0 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][3*k + 1] = 0;
				box[i][j][3*k + 2] = 0;
			}
		}
		if count1 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][3*k] = 0;
				box[i][j][3*k + 2] = 0;
			}
		}
		if count2 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][3*k] = 0;
				box[i][j][3*k + 1] = 0;
			}
		}
	}
}

func checkrow(box *[9][9][9]int, i int, j int){
	//index of other big boxes in the same row
	var x int = int(i/3)*3 + (i + 1)%3;
	var y int = int(i/3)*3 + (i + 2)%3;

	var count0 int = 0;
	var count1 int = 0;
	var count2 int = 0;

	var pos int;
	var dtm bool;

	//if j + 1 in box i is determined, do nothing
	if sum(box[i][j][:]) == 1{
		return;
	}

	for k:= 0; k < 3; k++{
		dtm, pos = checkOne(box[x][j][:]);
		if dtm {
			box[i][j][int(pos/3)*3] = 0;
			box[i][j][int(pos/3)*3 + 1] = 0;
			box[i][j][int(pos/3)*3 + 2] = 0;
		}
		dtm, pos = checkOne(box[y][j][:]);
		if dtm {
			box[i][j][int(pos/3)*3] = 0;
			box[i][j][int(pos/3)*3 + 1] = 0;
			box[i][j][int(pos/3)*3 + 2] = 0;
		}
		
	}

	for k:= 0; k < 3; k++{
		count0 += box[x][j][k];
		count1 += box[x][j][3 + k];
		count2 += box[x][j][6 + k];
		count0 += box[y][j][k];
		count1 += box[y][j][3 + k];
		count2 += box[y][j][6 + k];
	}

	count0 = stepfunc(count0);
	count1 = stepfunc(count1);
	count2 = stepfunc(count2);

	if count0 + count1 + count2 == 2{
		if count0 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][3 + k] = 0;
				box[i][j][6 + k] = 0;
			}
		}
		if count1 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][k] = 0;
				box[i][j][6 + k] = 0;
			}
		}
		if count2 == 0{
			for k:= 0;k < 3;k++{
				box[i][j][k] = 0;
				box[i][j][3 + k] = 0;
			}
		}
	}
}

func checkBox(box *[9][9][9]int, i int, j int){
	var pos int;
	var dtm bool = false;

	for k := 0; k < 9 ; k++{
		if k == j {
			continue;
		}
		dtm, pos = checkOne(box[i][k][:]);
		if dtm {
			box[i][j][pos] = 0;
		}
	}

}

func loopOnce(box *[9][9][9]int, counter int){
	counter += 1;
	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if sum(box[i][j][:]) > 1{
				checkcolumn(box,i,j);
				checkrow(box,i,j);
				checkBox(box,i,j);
			}
			if sum(box[i][j][:]) == 1 {
				counter += 1;
			}
		}

	}

}

func main() {

	var board [9][9] int;	

	var in_board string;
	var buffer [9][9] int;
	var chessman byte;
	var arg1_len int;
	//var entrynum int = 0;
	var argc int = len(os.Args);

	if argc < 9 {
		fmt.Println("Warning: Fewer than nine inputs means final lines will be populated with 0s."); 
	}

	for i := 1; i < argc; i++{
		in_board = os.Args[i];
		arg1_len = len(in_board);
		
		for j := 0; j < arg1_len && j < 17; j++{
			
			if j%2==0 {
				chessman = in_board[j];
				if chessman < 0x30 || chessman > 0x39{
					fmt.Println("Input error: Entries must consist of integers between 0 and 9 inclusive:", chessman);
			 		return;
				}
				buffer[i - 1][j/2] = int(chessman - 0x30);

			}else if j%2==1 && in_board[j] != 0x2c {
				fmt.Println("Input error: Entries must be separated by commas without spaces.");
				return;
			}
		}
	}

	board = buffer;

	// board = [9][9]int{
	// 	{0,6,0,0,0,0,0,0,1},
	// 	{0,4,0,6,5,0,0,0,9},
	// 	{0,0,0,0,0,7,0,6,2},
	// 	{1,8,0,0,0,9,0,3,0},
	// 	{0,0,0,0,1,0,0,0,0},
	// 	{0,2,0,3,0,0,0,4,5},
	// 	{7,5,0,4,0,0,0,0,0},
	// 	{4,0,0,0,3,6,0,7,0},
	// 	{6,0,0,0,0,0,0,2,0},
	// };
	// board = [9][9]int{
	// 	{0,0,5,0,0,7,8,0,0},
	// 	{0,0,0,6,9,0,2,0,0},
	// 	{2,7,0,0,3,0,0,0,9},
	// 	{5,0,0,0,0,0,0,2,0},
	// 	{0,3,1,0,2,0,7,9,0},
	// 	{0,2,0,0,0,0,0,0,5},
	// 	{8,0,0,0,4,0,0,6,7},
	// 	{0,0,4,0,5,8,0,0,0},
	// 	{0,0,3,1,0,0,9,0,0},
	// };


	fmt.Println("*.*.*.*.*.*.*.*.*.*");
	fmt.Println("The input is:");
	fmt.Println("*.*.*.*.*.*.*.*.*.*");

	for i := 0; i < 9 ; i++{
		fmt.Println(board[i]);
	}	
	
	var fin_board [9][9]int;

	var counter int = 0;
	var box [9][9][9]int;

	box, counter = parseBoard(board,counter);

	for i := 0; i < 20; i++{
		loopOnce(&box, counter);
		fin_board = parseBox(box);
	}


	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if(sum(box[i][j][:])!=1){
				fmt.Println("Board has no solution/is undetermined.");
				return;
			}
		}
	} 


	fmt.Println("*.*.*.*.*.*.*.*.*.*");
	fmt.Println("The solution is:");
	fmt.Println("*.*.*.*.*.*.*.*.*.*");

	for i := 0; i < 9 ; i++{
		fmt.Println(fin_board[i]);
	}

}
