package storage

import "strconv"

type Wallet struct {
	Id      string
	Balance uint64
}

type Transaction struct {
	Hash      string
	TimeStamp int64 `sql:"time_stamp"`
	Sender    string
	Receiver  string
	Amount    int64
	External  bool
}

type TransactionsSlice []Transaction

type RequestConfig struct {
	TimeStampFrom int64
	TimeStampTo   int64
	Sender        string
	Receiver      string
}

type Param struct {
	Field string
	Value string
}

type Params []Param

func (ts TransactionsSlice) ToCSVStrings() (csvs [][]string) {
	csvs = append(csvs, []string{"hash", "time_stamp", "sender", "receiver", "amount", "external"})
	for _, r := range ts {
		csvs = append(csvs, []string{r.Hash, strconv.FormatInt(r.TimeStamp, 10), r.Sender, r.Receiver, strconv.FormatInt(r.Amount, 10), strconv.FormatBool(r.External)})
	}
	return
}

func (rq *RequestConfig) ToParams() Params {
	p := make(Params, 0)
	if len(rq.Receiver) != 0 {
		p = append(p, Param{" receiver = ", rq.Receiver})
	}
	if len(rq.Sender) != 0 {
		p = append(p, Param{" sender = ", rq.Sender})
	}
	if rq.TimeStampFrom != 0 {
		p = append(p, Param{" time_stamp > ", strconv.FormatInt(rq.TimeStampFrom, 10)})
	}
	if rq.TimeStampTo != 0 {
		p = append(p, Param{" time_stamp < ", strconv.FormatInt(rq.TimeStampTo, 10)})
	}

	return p
}
