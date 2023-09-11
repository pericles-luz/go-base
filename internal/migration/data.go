package migration

import "database/sql"

func GenerateTestData(db *sql.DB) error {
	// stmt, err := db.Prepare(`
	// 	INSERT INTO Token(TokenUD, CO_CPF, CO_Token, TP_MeioEnvio, TS_Criacao)
	// 	VALUES(?, ?, ?, ?, ?)`)
	// if err != nil {
	// 	return err
	// }
	// _, err = stmt.Exec(TEST_TOKEN_UD, TEST_CPF, TEST_TOKEN, SEND_MEDIA_SMS, time.Now().UTC().Add(-3*time.Hour).Add(time.Hour*-1).Format("2006-01-02 15:04:05"))
	// if err != nil {
	// 	return err
	// }
	// _, err = stmt.Exec(uuid.NewString(), TEST_FINANCE_DIRECTOR_CPF, TEST_TOKEN, SEND_MEDIA_SMS, time.Now().UTC().Add(-3*time.Hour).Add(time.Hour*-1).Format("2006-01-02 15:04:05"))
	// if err != nil {
	// 	return err
	// }
	// err = stmt.Close()
	// if err != nil {
	// 	return err
	// }
	return nil
}
