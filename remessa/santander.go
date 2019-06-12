package remessa

import "github.com/padmoney/cobranca"

type Santander struct {
	conta      cobranca.Conta
	params     Params
	sequential int64
}

func NewSantander(conta cobranca.Conta, params Params, sequential int64) Santander {
	return Santander{
		conta:      conta,
		params:     params,
		sequential: sequential,
	}
}
