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

func (pg *PgAdapter) GetDApps() ([]*DApp, error) {
	var dApps []*DApp
	if err := pg.PgClient.Preload("CustodialGroup").Preload("CustodialGroup.Custodials").Find(&dApps).Error; err != nil {
		return nil, err
	}
	return dApps, nil
}

func (pg *PgAdapter) SaveDApp(dApp *DApp) error {
	return pg.PgClient.Create(dApp).Error
}

func (pg *PgAdapter) UpdateDApp(dApp *DApp) error {
	// First find the existing DApp
	existingDApp := &DApp{}
	if err := pg.PgClient.First(existingDApp, dApp.ID).Error; err != nil {
		return err
	}

	// Update the DApp with all fields including associations
	result := pg.PgClient.Model(existingDApp).
		Updates(map[string]interface{}{
			"chain_name":             dApp.ChainName,
			"btc_address_hex":        dApp.BTCAddressHex,
			"public_key_hex":         dApp.PublicKeyHex,
			"smart_contract_address": dApp.SmartContractAddress,
			"chain_id":               dApp.ChainID,
			"chain_endpoint":         dApp.ChainEndpoint,
			"rpc_url":                dApp.RPCUrl,
			"access_token":           dApp.AccessToken,
			"token_contract_address": dApp.TokenContractAddress,
			"custodial_group_id":     dApp.CustodialGroupID,
		})

	if result.Error != nil {
		return result.Error
	}

	// Update the CustodialGroup association
	if err := pg.PgClient.Model(existingDApp).Association("CustodialGroup").Replace(dApp.CustodialGroup); err != nil {
		return err
	}

	return nil
}

func (pg *PgAdapter) ToggleDApp(ID uint) error {
	dApps := pg.PgClient
	var result DApp
	if err := dApps.Where("id = ?", ID).First(&result).Error; err != nil {
		return err
	}
	result.State = !result.State
	return dApps.Save(&result).Error
}

func (pg *PgAdapter) DeleteDApp(ID uint) error {
	dApps := pg.PgClient
	return dApps.Where("id = ?", ID).Delete(&DApp{}).Error
}
