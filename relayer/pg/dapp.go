package pg

func (pg *PgAdapter) FindProtocolInfo(chainName string, contractAddress string) (*DApp, error) {
	var protocolInfo DApp

	query := pg.PgClient.Where("chain_name = ? AND smart_contract_address = ?", chainName, contractAddress)
	result := query.First(&protocolInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &protocolInfo, nil
}
