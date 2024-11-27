package pg

func (pg *PgAdapter) UpdateContractCallApproved(messageID string, executeHash string) error {
	updateData := map[string]interface{}{
		"execute_hash": executeHash,
		"status":       APPROVED,
	}
	record := pg.PgClient.Model(&RelayData{}).Where("id = ?", messageID).Updates(updateData)
	if record.Error != nil {
		return record.Error
	}
	return nil
}

func (pg *PgAdapter) FindContractCallByCommnadId(commandId string) (*CallContract, error) {
	var contractCall CallContract
	result := pg.PgClient.Where("command_id = ?", commandId).First(&contractCall)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contractCall, nil
}
