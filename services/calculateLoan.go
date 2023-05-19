package services

import (
	"fmt"
	"math"

	"github.com/dustin/go-humanize"
	"jgnovak.net/loanshark-api/models"
)

func CalculateLoan(loanRequest models.LoanRequest) models.LoanResponse {
	// Convert a flaot value like 12500.75 to a currency string like $12,500.75
	loanPaymentFormatter := func(val float64) string {
		return fmt.Sprintf("$%v", humanize.FormatFloat("#,###.##", val))
	}

	// Calculate loan details
	correctedRate := loanRequest.Rate / 100.0
	monthlyPayment := calculateMonthlyPayment(loanRequest.Amount, correctedRate, loanRequest.Term)
	totalCost := monthlyPayment * float64(loanRequest.Term)
	totalInterest := totalCost - loanRequest.Amount
	remainingBalance := loanRequest.Amount

	var loanResponse models.LoanResponse

	// Calculate loan payments
	var payments []models.Payment
	for i := 1; i <= loanRequest.Term; i++ {
		interestPayment := remainingBalance * (correctedRate / 12.0)
		principalPayment := monthlyPayment - interestPayment
		remainingBalance -= principalPayment

		// Correct for possible negative interest value in final payments
		if i == loanRequest.Term {
			principalPayment += remainingBalance
			monthlyPayment = principalPayment + interestPayment
			remainingBalance = 0.0
		}

		loanPayment := models.Payment{
			Month:     i,
			Payment:   loanPaymentFormatter(monthlyPayment),
			Interest:  loanPaymentFormatter(interestPayment),
			Principal: loanPaymentFormatter(principalPayment),
			Balance:   loanPaymentFormatter(remainingBalance),
		}

		payments = append(payments, loanPayment)
	}

	loanResponse.Amount = loanPaymentFormatter(loanRequest.Amount)
	loanResponse.Rate = fmt.Sprintf("%.2f", loanRequest.Rate)
	loanResponse.MonthlyPayment = loanPaymentFormatter(monthlyPayment)
	loanResponse.TotalInterest = loanPaymentFormatter(totalInterest)
	loanResponse.TotalCost = loanPaymentFormatter(totalCost)
	loanResponse.Payments = payments

	return loanResponse
}

func calculateMonthlyPayment(loanAmount float64, annualInterestRate float64, termInMonths int) float64 {
	monthlyInterestRate := annualInterestRate / 12.0
	onePlusMonthlyRate := monthlyInterestRate + 1.0
	mathPower := math.Pow(onePlusMonthlyRate, float64(termInMonths))

	numerator := loanAmount * monthlyInterestRate * mathPower
	denominator := mathPower - 1.0

	return numerator / denominator
}
