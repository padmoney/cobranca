package remessa

import (
	"errors"

	"github.com/padmoney/cobranca"
)

type RemessaGenerator interface {
}

type Remessa struct {
	Generator RemessaGenerator
}

func New(conta cobranca.Conta, params Params, sequential int64) (Remessa, error) {
	remessa := Remessa{}

	switch conta.Banco {
	case cobranca.CodigoSantander:
		remessa.Generator = NewSantander(conta, params, sequential)
	default:
		return remessa, errors.New(cobranca.BancoNaoSuportado)
	}
	return remessa, nil
}

/*
   public function __construct($conta, array $params = [])
   {
       $this->conta = $conta;
       $this->params = $params;

       $agencia = explode('-', $conta->agencia());
       $this->agencia = isset($agencia[0]) ? $agencia[0] : $conta->agencia();
       $this->agencia_dv = $this->digitoAgencia();

       $conta_numero = explode('-', $conta->contaCorrente());
       $this->conta_numero = isset($conta_numero[0]) ? $conta_numero[0] : $conta->contaCorrente();
       $this->conta_dv = isset($conta_numero[1]) ? $conta_numero[1] : '';

       $this->nome_cedente = $conta->beneficiario()->nome();
       $this->carteira = $conta->carteira();
       $this->convenio = $conta->convenio();
       $this->variacao =  $conta->variacao();
       $this->sequencial_remessa = $conta->sequencialRemessa();
   }
*/
