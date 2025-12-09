package dsa_brushup



func Matrix_set_zero(matrix [][]int) [][]int {
	
	rows:=len(matrix)
	cols:=len(matrix[0])

    rows_with_zero:=make([]bool , rows)
	cols_with_zero:=make([]bool,cols)


	for i:=0;i<len(matrix);i++ {
       for j:=0;j<len(matrix[i]);j++ {
          if matrix[i][j]== 0 {
            rows_with_zero[i]=true
			cols_with_zero[j]=true
		  }
	   }
	}

	//set rows zero

	for i:=0;i<rows;i++ {
      if rows_with_zero[i] {
		for j:=0;j<cols;j++ {
			matrix[i][j]=0
		}
	  }
	}

	//set cols zero
	for j:=0;j<cols;j++ {
		if cols_with_zero[j]  {
		    for i:=0;i<rows;i++ {
				matrix[i][j]=0
			}
		}
	}
	return matrix
}