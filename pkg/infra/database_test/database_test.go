package database_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/infra/database"

	"github.com/stretchr/testify/require"
)

func TestDatabase(t *testing.T) {
	if os.Getenv("GITHUB") != "no" {
		t.Skip("Skip when running on github")
	}
	db, err := databaseConnection(t)
	require.NoError(t, err)
	defer db.Close()
	require.True(t, db.IsConnected())
	databaseDropTable(t, db)
	databaseCreateTable(t, db)
	databaseInsertRecord(t, db)
	databaseGetOne(t, db)
	databaseGetRecord(t, db)
	databaseGetRecords(t, db)
	databaseUpdateRecord(t, db)
}

func TestAgnuDatabase(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	db, err := databaseAgnuConnection(t)
	require.NoError(t, err)
	defer db.Close()
	require.True(t, db.IsConnected())
	databaseDropTable(t, db)
	databaseCreateTable(t, db)
	databaseInsertRecord(t, db)
	databaseGetOne(t, db)
	databaseGetRecord(t, db)
	databaseGetRecords(t, db)
}

func databaseConnection(t *testing.T) (*database.Database, error) {
	configuration, err := conf.NewInitialConfig("initial.test")
	var db *database.Database
	require.NoError(t, err)
	db, err = database.NewDatabase(configuration.DBConfiguration)
	require.NoError(t, err)
	require.True(t, db.IsConnected())
	return db, nil
}

func databaseAgnuConnection(t *testing.T) (*database.Database, error) {
	configuration, err := conf.NewInitialConfig("initial.test")
	var db *database.Database
	require.NoError(t, err)
	require.NotNil(t, configuration.AgnuDBConfiguration)
	db, err = database.NewDatabase(configuration.AgnuDBConfiguration)
	require.NoError(t, err)
	require.True(t, db.IsConnected())
	return db, nil
}

func databaseCreateTable(t *testing.T, db *database.Database) {
	err := db.Exec(`CREATE TABLE IF NOT EXISTS Pessoa (
		PessoaID int(11) NOT NULL AUTO_INCREMENT,
		CO_CPF char(11) DEFAULT NULL,
		DE_Pessoa varchar(128) DEFAULT NULL,
		DT_Nascimento DATE,
		SG_Sexo char(1) DEFAULT NULL,
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		PRIMARY KEY (PessoaID),
		KEY CO_CPF (CO_CPF) USING BTREE
	) ENGINE=MyISAM DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Identidade (
		IdentidadeID int(11) NOT NULL AUTO_INCREMENT,
		ID_Pessoa int(11) NOT NULL DEFAULT 0,
		DE_Numero char(12) DEFAULT NULL,
		SG_Orgao char(10) DEFAULT NULL,
		SG_Estado char(2) DEFAULT NULL,
		DT_Emissao DATE,
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		PRIMARY KEY (IdentidadeID),
		KEY ID_Pessoa (ID_Pessoa) USING BTREE
	) ENGINE=MyISAM DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Endereco (
		EnderecoID int(11) NOT NULL AUTO_INCREMENT,
		ID_Pessoa int(11) NOT NULL DEFAULT '0',
		CO_Cep char(8) NOT NULL DEFAULT '',
		DE_Endereco varchar(150) NOT NULL DEFAULT '',
		DE_Numero char(10) DEFAULT NULL,
		DE_Complemento varchar(30) DEFAULT NULL,
		DE_Bairro varchar(150) DEFAULT NULL,
		DE_Cidade varchar(150) NOT NULL DEFAULT '',
		SG_Estado char(2) NOT NULL DEFAULT '',
		TS_Operacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		SN_Publicidade tinyint(4) NOT NULL DEFAULT '0',
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		PRIMARY KEY (EnderecoID),
		KEY clienteCep (CO_Cep) USING BTREE,
		KEY enderecoPessoa (ID_Pessoa)
	) ENGINE=MyISAM DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Telefone (
		TelefoneID int(11) NOT NULL AUTO_INCREMENT,
		ID_Pessoa int(11) NOT NULL DEFAULT '0',
		CO_DDD char(2) NOT NULL DEFAULT '',
		CO_Telefone char(9) NOT NULL DEFAULT '',
		DE_Observacao varchar(128) DEFAULT NULL,
		ID_Status tinyint(4) NOT NULL DEFAULT '0',
		SN_WhatsApp tinyint(4) NOT NULL DEFAULT '0',
		SN_Publicidade tinyint(4) NOT NULL DEFAULT '1',
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		TS_Operacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (TelefoneID),
		KEY clienteTelefone (CO_DDD,CO_Telefone) USING BTREE,
		KEY telefonePessoa (ID_Pessoa)
	) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS ContaCorrente (
		ContaCorrenteID int(11) NOT NULL AUTO_INCREMENT,
		ID_Pessoa int(11) NOT NULL DEFAULT '0',
		CO_Banco char(3) NOT NULL DEFAULT '',
		CO_Agencia char(4) NOT NULL DEFAULT '',
		CO_AgenciaDv char(1) NOT NULL DEFAULT '',
		CO_ContaCorrente char(12) NOT NULL DEFAULT '',
		CO_ContaCorrenteDv char(1) NOT NULL DEFAULT '',
		CO_Operacao char(3) NOT NULL DEFAULT '',
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		PRIMARY KEY (ContaCorrenteID),
		KEY telefonePessoa (ID_Pessoa)
		) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Email (
			EmailID int(11) NOT NULL AUTO_INCREMENT,
			ID_Pessoa int(11) NOT NULL DEFAULT '0',
			EM_Email varchar(150) DEFAULT NULL,
			DE_Observacao varchar(128) DEFAULT NULL,
			ID_Status tinyint(4) NOT NULL DEFAULT '0',
			SN_Publicidade tinyint(4) NOT NULL DEFAULT '1',
			SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
			TS_Operacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (EmailID),
			KEY clienteEmail (EM_Email) USING BTREE,
			KEY emailPessoa (ID_Pessoa)
		) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Agenda (
		AgendaID varchar(36) NOT NULL DEFAULT '',
		DE_Agenda varchar(80) DEFAULT NULL,
		TP_Agenda tinyint(4) NOT NULL DEFAULT '0',
		ID_Atendente varchar(36) NOT NULL DEFAULT '0',
		ID_Alternativo varchar(36) NOT NULL DEFAULT '0',
		DE_Observacao varchar(250) DEFAULT NULL,
		ID_Cliente varchar(36) NOT NULL DEFAULT '0',
		ID_Status int(11) NOT NULL DEFAULT '0',
		TS_Agenda timestamp NULL DEFAULT NULL,
		TS_Criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		SN_Ativo tinyint(4) NOT NULL DEFAULT '1',
		PRIMARY KEY (AgendaID)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8`)
	require.NoError(t, err)
}

func databaseInsertRecord(t *testing.T, db *database.Database) {
	err := db.Exec(`REPLACE INTO Pessoa(PessoaID, CO_CPF, DE_Pessoa, DT_Nascimento, SG_Sexo) VALUES(0, ?, ?, ?, ?)`, "91134803672", "Testando Record", "2021-08-01", "M")
	require.NoError(t, err)
	require.Greater(t, db.GetLastInsertId(), uint64(0))
}

func databaseGetOne(t *testing.T, db *database.Database) {
	value, err := db.GetOne(`SELECT CO_CPF FROM Pessoa WHERE CO_CPF=(?)`, "91134803672")
	require.NoError(t, err)
	require.Equal(t, "91134803672", string(value))
}

func databaseGetRecord(t *testing.T, db *database.Database) {
	value, err := db.GetRecord(`SELECT PessoaID, CO_CPF FROM Pessoa WHERE CO_CPF=(?)`, "91134803672")
	require.NoError(t, err)
	require.Equal(t, "91134803672", value["CO_CPF"])
}

func databaseGetRecords(t *testing.T, db *database.Database) {
	value, err := db.GetRecords(`SELECT PessoaID, CO_CPF FROM Pessoa WHERE CO_CPF=(?)`, "91134803672")
	require.NoError(t, err)
	require.Greater(t, len(value), 0)
}

func databaseUpdateRecord(t *testing.T, db *database.Database) {
	command := map[string]interface{}{
		"DE_Tabela": "Pessoa",
		"registro": map[string]interface{}{
			"PessoaID":      1,
			"CO_CPF":        "00000000191",
			"DE_Pessoa":     "Testando Update",
			"DT_Nascimento": "2021-08-01",
			"SG_Sexo":       "M",
		},
	}
	err := db.Update(command["DE_Tabela"].(string), command["registro"].(map[string]interface{}))
	require.NoError(t, err)
	value, err := db.GetRecord(`SELECT PessoaID, CO_CPF, DE_Pessoa, DT_Nascimento, SG_Sexo FROM Pessoa WHERE CO_CPF=(?)`, "00000000191")
	require.NoError(t, err)
	require.Equal(t, "Testando Update", value["DE_Pessoa"])
	require.Equal(t, "2021-08-01", value["DT_Nascimento"])
	require.Equal(t, "M", value["SG_Sexo"])
}

func databaseDropTable(t *testing.T, db *database.Database) {
	err := db.Exec(`DROP TABLE IF EXISTS Pessoa`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Identidade`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Endereco`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Telefone`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS ContaCorrente`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Email`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Cartao`)
	require.NoError(t, err)
	err = db.Exec(`DROP TABLE IF EXISTS Agenda`)
	require.NoError(t, err)
}
