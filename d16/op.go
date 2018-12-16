package main

var AllOps = map[string]OpFunc{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqri": eqri,
	"eqir": eqir,
	"eqrr": eqrr,
}

type OpFunc func(r Register, a, b, c int) Register

func addr(r Register, a, b, c int) Register {
	r[c] = r[a] + r[b]
	return r
}

func addi(r Register, a, b, c int) Register {
	r[c] = r[a] + b
	return r
}

func mulr(r Register, a, b, c int) Register {
	r[c] = r[a] * r[b]
	return r
}

func muli(r Register, a, b, c int) Register {
	r[c] = r[a] * b
	return r
}

func banr(r Register, a, b, c int) Register {
	r[c] = r[a] & r[b]
	return r
}

func bani(r Register, a, b, c int) Register {
	r[c] = r[a] & b
	return r
}

func borr(r Register, a, b, c int) Register {
	r[c] = r[a] | r[b]
	return r
}

func bori(r Register, a, b, c int) Register {
	r[c] = r[a] | b
	return r
}

func setr(r Register, a, b, c int) Register {
	r[c] = r[a]
	return r
}

func seti(r Register, a, b, c int) Register {
	r[c] = a
	return r
}

func gtir(r Register, a, b, c int) Register {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtri(r Register, a, b, c int) Register {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtrr(r Register, a, b, c int) Register {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqir(r Register, a, b, c int) Register {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqri(r Register, a, b, c int) Register {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqrr(r Register, a, b, c int) Register {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}
