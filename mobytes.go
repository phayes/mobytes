/*
Mo'Bytes - more common []byte functions for golang

Functions include bitwise operatoins (AND, OR, XOR, NOT, and ANDNOT) and other useful functions for working with []byte's.
*/
package mobytes

// Bitwise AND
func AND(a, b []byte) []byte {
	opedlen := min(len(a), len(b))
	result := make([]byte, opedlen)
	for i := 0; i < opedlen; i++ {
		result[i] = a[i] & b[i]
	}
	return result
}

// Bitwise OR
func OR(a, b []byte) []byte {
	opedlen := min(len(a), len(b))
	result := make([]byte, opedlen)
	for i := 0; i < opedlen; i++ {
		result[i] = a[i] | b[i]
	}
	return result
}

// Bitwise XOR
func XOR(a, b []byte) []byte {
	opedlen := min(len(a), len(b))
	result := make([]byte, opedlen)
	for i := 0; i < opedlen; i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

// Bitwise AND NOT (clear)
func ANDNOT(a, b []byte) []byte {
	opedlen := min(len(a), len(b))
	result := make([]byte, opedlen)
	for i := 0; i < opedlen; i++ {
		result[i] = a[i] &^ b[i]
	}
	return result
}

// Bitwise NOT. Replaces all 0 bits with 1 bits and all 1 bits with 0 bits.
func NOT(a []byte) []byte {
	opedlen := len(a)
	result := make([]byte, opedlen)
	for i := 0; i < opedlen; i++ {
		result[i] = 255 - a[i]
	}
	return result
}

// Split the []byte every n bytes
func SplitEvery(a []byte, n int) [][]byte {
	length := len(a)
	result := make([][]byte, length/n)
	i := 0
	for j := 0; j < length; j += n {
		if j+n < length {
			result[i] = a[j : j+n]
		} else {
			result[i] = a[j:]
		}
		i++
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
