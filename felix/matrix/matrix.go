package matrix

import (

)

type Matrix struct{
	matrix [][]int
}

func NewMatrix(list [][]int) *Matrix{
	m:=new(Matrix)
	m.matrix=list
	return m
}

func (m *Matrix)GetList() [][]int{
	return m.matrix
}

//Nicht Nebenläufig!!!!!!!!!!!!!
func (m *Matrix)Mult(m2 *Matrix) *Matrix{
    var bMatrix [][]int = m2.GetList()
    var product [][]int 
    product=make([][]int,len(m.matrix))	
    for i:=0;i<len(m.matrix);i++{
    	product[i]=make([]int,len(bMatrix[0]))
    }
    
    for row := 0; row < len(m.matrix); row++{
        for col:= 0; col < len(bMatrix[0]); col++{
            // Multiply the row of A by the column of B to get the row, column of product.
            for inner:= 0; inner < len(m.matrix[0]); inner++ {
                	product[row][col] += m.matrix[row][inner] * bMatrix[inner][col];
            }
        }
    }   
    newMatrix:=new(Matrix)
    newMatrix.matrix=product
    return newMatrix
}

//Nebenläufig
func (m *Matrix)MultGo(m2 *Matrix) *Matrix{
    var bMatrix [][]int = m2.GetList()
	//channel:=make(chan int)
    //channel2:=make(chan int)
    
    var product [][]int 
    product=make([][]int,len(m.matrix))	
    for i:=0;i<len(m.matrix);i++{
    	product[i]=make([]int,len(bMatrix[0]))
    }
    
    for row := 0; row < len(m.matrix); row++{
        for col:= 0; col < len(bMatrix[0]); col++{
            // Multiply the row of A by the column of B to get the row, column of product.
            channel:=make(chan int,len(m.matrix[0]))
            var erg int = 0
            for i:=0;i<len(m.matrix[0]);i++{
            	go multHelp(m.matrix[row][i],bMatrix[i][col],channel)
            }
            for i:=0;i<len(m.matrix[0]);i++{
            	erg +=<-channel
            }
            product[row][col] += erg;            
        }
    }   
    newMatrix:=new(Matrix)
    newMatrix.matrix=product
    return newMatrix
}

func multHelp(a,b int,result chan int){
	var ergebnis int
	ergebnis = a*b
	result<-ergebnis
}