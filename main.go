package main

import (
	"fmt"
	"math"
)


const inflationRate = 6.5
func main() {
	
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scanln(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	
	futureValue, futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,years)
	fmt.Printf("Future value: %0.2f\nFuture real value: %0.2f\n",  futureValue,futureRealValue)
	


}

func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues (investmentAmount, expectedReturnRate,years float64) (futureValue float64, futureRealValue float64){
	fv:= investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv := futureValue / math.Pow((1+inflationRate/100), years)

	return fv,rfv
}
