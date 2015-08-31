package fee

type Fee struct {
	Value int
}

func (f Fee) Add(addFee Fee) Fee {
	return Fee{f.Value + addFee.Value}
}

func (f Fee) Multiple(rate Rate) Fee {
	return Fee{int(float64(f.Value) * rate.Value)}
}
