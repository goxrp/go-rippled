package gorippled

// TransferFeeDecimal calculates a transfer fee using the
// transfer rate from the Account Root object.
func TransferFeeDecimal(transferRate int64) float64 {
	return (float64(transferRate) / 1000000000.0) - 1.0
}
