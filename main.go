/*Panshin Roman, 05.03.2022*/

package main

import "fmt"
import "math/rand"
import "math"
import "time"

const kol int  = 8

type quenPosition struct{
	iPos, jPos int 
} 

func printDesk(quens [kol]quenPosition){
	var desk [8][8]string
	
	for i:= 0; i < 8; i++{
		for j:= 0; j < 8; j++{
			desk[i][j] = "0"
		}
	}
	for i:= 0; i < kol; i++{
		desk[quens[i].iPos][quens[i].jPos] = "👸"
	}
	for i:= 0; i < 8; i++{
		for j:= 0; j < 8; j++{
			fmt.Print(desk[i][j], " ")
		}
		fmt.Println()
	}
	
}

func checkPositionReplay(quen quenPosition, quens [kol]quenPosition) bool{
	for i:= 0; i < kol; i++{
		if(quen == quens[i]){
			return true
		}
	}
	return false
}

func G(quens [kol]quenPosition) [kol]quenPosition{
	var result [kol]quenPosition = quens
	point := rand.Intn(kol)
	var nQuen quenPosition
	
	for{
		nQuen = quenPosition{iPos: rand.Intn(8), jPos: rand.Intn(8)}
		if(!checkPositionReplay(nQuen, quens)){
			break
		}
	}

	result[point] = nQuen
	return result
}

func fNumberOfAtacks(quens [kol]quenPosition) int{
	var result int = 0
	for i:= 0; i < kol; i++{
		for j := 0; j < kol; j++{
			if(i != j){
				if(quens[i].iPos == quens[j].iPos){
					result++
				}else if(quens[i].jPos == quens[j].jPos){
					result++
				}else if(math.Abs(float64(quens[i].jPos - quens[j].jPos)) == math.Abs(float64(quens[i].iPos - quens[j].iPos))){
					result++
				}
			}
		}
	}
	return result 
}

func main(){
	rand.Seed(time.Now().UnixNano())
	var quens [kol]quenPosition
	for i:=0; i < kol; i++ {
		buffQuen := quenPosition{iPos: i, jPos: i}
		quens[i] = buffQuen 

	}
	var t float64 = 10000000
	var alpha float64 = 0.99999999
	var minF int = int(t)
	var minDesk [kol]quenPosition
	fmt.Println(quens)
	for{
		Xquens := G(quens)
		fx := fNumberOfAtacks(Xquens)
		fk1 := fNumberOfAtacks(quens)
		delta := fx - fk1
		if(delta < 0){
			quens = Xquens
			if(fx == 0){
				printDesk(quens)
				break
			}
			if(fx < minF){
				minF = fx
				minDesk = Xquens
			}
			
		}
		if(delta > 0){
			var p float64 = math.Exp(float64(delta) / t)
			var r float64 = rand.Float64()
			if(p > r){
				quens = Xquens
				t *= alpha
				if (t < 1.0){
					printDesk(minDesk)
					fmt.Println("min f = ", minF)
					break
				}
			} 
		}
	}
}
