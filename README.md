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

```go
func InterfaceToInt(in interface{}) int
```

Converte uma interface para um inteiro

```go
func StringToInt(str string) int
```

Converte uma string para um inteiro

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

### JWT

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

### Configuração

O pacote `\conf` é responsável por carregar as configurações da aplicação

```go
func (db *Database) NewDB(cfg *conf.DBConfiguration) error
```
Cria uma nova conexão com o banco de dados

```go
func (db *Database) GetDatabase() *sql.DB
```

Retorna a conexão com o banco de dados

```go
func (db *Database) IsConnected() bool
```

Verifica se a conexão com o banco de dados está ativa

```go
func (db *Database) Close() error
```

Fecha a conexão com o banco de dados

```go
func (db *Database) GetResult() (sql.Result, error)
```

Retorna o resultado da última query executada

```go
func (db *Database) GetLastInsertId() uint64
```

Retorna o ID da última inserção

```go
func (db *Database) GetStmt(sql string) (*sql.Stmt, error)
```

Retorna um statement para ser usado em queries

```go
func (db *Database) Exec(sql string, data ...interface{}) error
```

Executa uma query

```go
func (db *Database) GetOne(sql string, data ...interface{}) ([]byte, error)
```

Executa uma query e retorna apenas um resultado. Independente do tipo de dados, o resultado é retornado em um slice de bytes

```go
func (db *Database) GetRecord(sqlString string, data ...interface{}) (map[string]interface{}, error)
```

Executa uma query e retorna apenas um resultado. O resultado é retornado em um map de interface

```go
func (db *Database) GetRecords(sqlString string, data ...interface{}) ([]map[string]interface{}, error)
```

Executa uma query e retorna vários resultados. O resultado é retornado em um slice de map de interface

```go
func (db *Database) Insert(tableName string, data map[string]interface{}) error
```

Insere um registro no banco de dados

```go
func (db *Database) Update(tableName string, data map[string]interface{}) error
```

Atualiza um registro no banco de dados

```go
func (db *Database) IsMySQL() bool
```

Verifica se o banco de dados é MySQL

```go
func (db *Database) IsPostgreSQL() bool
```

Verifica se o banco de dados é PostgreSQL

```go
func (db *Database) IsSQLite() bool
```

Verifica se o banco de dados é SQLite


#### Database

O pacote `\database` é responsável por gerenciar a conexão com o banco de dados, além de disponibilizar funções para facilitar a manipulação de dados

```go
func (p *Pool) AddConnection(name string, db *sql.DB) error
```

Adiciona uma nova conexão ao pool

```go
func (p *Pool) GetConnection(name string) (*sql.DB, error)
```

Retorna uma conexão do pool

```go
func (p *Pool) CloseConnection(name string) error
```

Fecha uma conexão do pool

```go
func (p *Pool) IsConnected(name string) bool
```

Verifica se uma conexão está ativa


#### Mensagens

O pacote `\messaging` é responsável por gerenciar as mensagens a serem enviadas ao usuário ou cliente

#### Legal One