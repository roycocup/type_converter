package main

import "strconv"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

const (
	MinUint8 = 0
	MaxUint8 = 255

	MinUint16 = 0
	MaxUint16 = 65535

	MinInt32  = 0
	MaxUint32 = 4294967295

	MinUint64 = 0
	MaxUint64 = 18446744073709551615

	MinInt8 = -128
	MaxInt8 = 127

	MinInt16 = -32768
	MaxInt16 = 32767

	minInt32 = -2147483648
	MaxInt32 = 2147483647

	MinInt64 = -9223372036854775808
	MaxInt64 = 9223372036854775807
)

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}

func StringToInt64(str string) int64 {
	id, _ := strconv.ParseInt(str, 10, 64)
	return id
}

func Uint64ToString(u uint64) string {

	return strconv.FormatUint(u, 10)
}

func Int64ToInt(i64 int64) (i int) {
	i = 0
	if i64 > int64(MaxInt) || i64 < 0 {
		i = 0
	} else {
		i = int(i)
	}

	return
}

func Int64ToUint64(i64 int64) (u uint64) {
	if i64 < 0 {
		u = 0
	} else {
		u = uint64(i64)
	}

	return
}
