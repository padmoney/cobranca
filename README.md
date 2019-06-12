# cobranca

## Boleto

## Registro Online

### Santander

```
import "github.com/padmoney/cobranca/boleto"

// ...

// Crie uma conta
c := NewConta("033", "1234", "123456-7", "101", "9876543210")

// Informe os dados boleto

b := BoletoOnline{
Numero: "1",
NossoNumero: "0000000001-1",
}

r := NewOnlineRegister(c)
r.Register(b)
```
