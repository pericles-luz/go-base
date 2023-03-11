# go-base
Cinto de utilidades para não precisar replicar funções básicas em todas as aplicações
Várias funções básicas para facilitar o desenvolvimento de aplicações em Go

## Funções

### Utils

#### Conversores

```go
func MapInterfaceToBytes(data map[string]interface{}) []byte
```

Converte um map de interface para um slice de bytes

```go
func ByteToMapInterface(bytes []byte) map[string]interface{}
```

Converte um slice de bytes para um map de interface

```go
func MapStringToMapInterface(data map[string]string) map[string]interface{}
```

Converte um map de string para um map de interface

#### Arquivos

```go
func GetBaseDirectory(directory string) string
```

Retorna o diretório base da aplicação

```go
func FileExists(path string) bool
```

Verifica se um arquivo ou diretório existe

#### Geral

```go
func isTesting() bool
```

Verifica se a aplicação está em modo de teste

```go
func Hash256(s string) string
```

Gera um hash SHA256

```go
GetOnlyNumbers(str string) string
```

Retorna apenas os números de uma string

```go
func CompleteWithZeros(str string, length int) string
```

Completa uma string com zeros a esquerda

#### JWT

Estrutura usada para gerenciar o token JWT

```go
func (j *JwtServer) Valid(token string) bool
```

Verifica se o token é válido

```go
func (j *JwtServer) Create(payload map[string]interface{}) (string, error)
```

Cria um token JWT

```go
func (j *JwtServer) Parse(token string) (map[string]interface{}, error)
```

Decodifica um token JWT

```go
func ExtractValue(key string, jwt string) (interface{}, error)
```

Extrai um valor do token JWT

#### Validadores

```go
func ValidateCPF(cpf string) bool
```

Valida um CPF

```go
func ValidateUUID(subject string) bool
```

Valida um UUID

```go
func ValidateTimestamp(timestamp string) bool
```

Valida um timestamp

```go
func ValidateDDD(ddd string) bool
```

Valida um DDD

```go
func ValidatePhoneNumber(phoneNumber string) bool
```

Valida um número de telefone

```go
func HasOnlyNumbers(str string) bool
```

Verifica se uma string possui apenas números

```go
func ValidateCellphoneNumber(phoneNumber string) error
```

Valida um número de celular

```go
func ValidateLandlineNumber(phoneNumber string) error
```

Valida um número de telefone fixo

```go
func ValidateEmail(email string) bool
```

Valida um email

#### Configuração

O pacote `\conf` é responsável por carregar as configurações da aplicação

#### Database

O pacote `\database` é responsável por gerenciar a conexão com o banco de dados, além de disponibilizar funções para facilitar a manipulação de dados

#### Mensagens

O pacote `\messaging` é responsável por gerenciar as mensagens a serem enviadas ao usuário ou cliente