/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : binomial_expression_test.go
#   Last Modified : 2018-12-17 11:43
#   Describe      : 测试二项式公式的计算
#
# ====================================================*/

package basic

import (
	//"bufio"
	"log"
	//"os"
	//"strconv"
	"testing"
	//"time"
)

// CalculateBinomialRecursionTest 测试二项式公式的计算 使用递归
func TestCalculateBinomialRecursion(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialRecursion(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialRecursion=%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}

//// TestCalculateBinomialRecursionWithTime 测试二项式公式的计算 使用递归
//func TestCalculateBinomialRecursionWithTime(t *testing.T) {
	//outputFile, outputError := os.OpenFile("recursion.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//if outputError != nil {
		//log.Printf("An error occurred with file opening or creation\n")
		//return
	//}
	//defer outputFile.Close()
	//outputWriter := bufio.NewWriter(outputFile)
	//outputWriter.WriteString("n\tk\ttime\tresult\n")
	//for n := 34; n < 35; n++ {
		//for k := 1; k < n/2; k++ {
			//timeBegin := time.Now().UnixNano()
			//result := CalculateBinomialRecursion(k, n)
			//timeEnd := time.Now().UnixNano()
			//outputString := strconv.Itoa(n) + "\t" + strconv.Itoa(k) + "\t" + strconv.FormatInt(int64(result/1000), 10) + "\t" + strconv.FormatInt(timeEnd-timeBegin, 10) + "\n"
			//outputWriter.WriteString(outputString)

		//}
	//}
	//outputWriter.Flush()
//}

// TestCalculateBinomialMemo测试二项式公式的计算 使用备忘录方法
func TestCalculateBinomialMemo(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialMemo(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialMemo=%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}

// TestCalculateBinomialMemoWithTime 测试二项式公式的计算 使用备忘录方法
//func TestCalculateBinomialMemoWithTime(t *testing.T) {
	//outputFile, outputError := os.OpenFile("memo.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//if outputError != nil {
		//log.Printf("An error occurred with file opening or creation\n")
		//return
	//}
	//defer outputFile.Close()
	//outputWriter := bufio.NewWriter(outputFile)
	//outputWriter.WriteString("n\tk\ttime\tresult\n")
	//for n := 34; n < 35; n++ {
		//for k := 1; k <= n/2; k++ {
			//timeBegin := time.Now().UnixNano()
			//result := CalculateBinomialMemo(k, n)
			//timeEnd := time.Now().UnixNano()
			//outputString := strconv.Itoa(n) + "\t" + strconv.Itoa(k) + "\t" + strconv.FormatInt((timeEnd-timeBegin)/1000, 10) + "\t" + strconv.FormatInt(int64(result), 10) + "\n"
			//outputWriter.WriteString(outputString)

		//}
	//}
	//outputWriter.Flush()
//}

// TestCalculateBinomialIteration 测试二项式公式的计算 使用迭代方法
func TestCalculateBinomialIteration(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialIteration(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialIteration =%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}

//// TestCalculateBinomialIterationWithTime 测试二项式公式的计算 使用迭代方法
//func TestCalculateBinomialIterationWithTime(t *testing.T) {
	//outputFile, outputError := os.OpenFile("iteration.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//if outputError != nil {
		//log.Printf("An error occurred with file opening or creation\n")
		//return
	//}
	//defer outputFile.Close()
	//outputWriter := bufio.NewWriter(outputFile)
	//outputWriter.WriteString("n\tk\ttime\tresult\n")
	//for n := 34; n < 35; n++ {
		//for k := 1; k <= n/2; k++ {
			//timeBegin := time.Now().UnixNano()
			//result := CalculateBinomialIteration(k, n)
			//timeEnd := time.Now().UnixNano()
			//outputString := strconv.Itoa(n) + "\t" + strconv.Itoa(k) + "\t" + strconv.FormatInt((timeEnd-timeBegin)/1000, 10) + "\t" + strconv.FormatInt(int64(result), 10) + "\n"
			//outputWriter.WriteString(outputString)

		//}
	//}
	//outputWriter.Flush()
//}
