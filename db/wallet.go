package db

type Wallet struct {
	Model
	Address string `gorm:",not null" json:"address"`
}

// NewWallet returns a Wallet pointer.
func NewWallet(wallet string) *Wallet {
	return &Wallet{Address: wallet}
}

// Save the Wallet in database.
func (w *Wallet) Save() error {
	db := GetConn()
	defer db.Close()
	db = db.Save(w)
	return db.Error
}

// Delete Wallet in db
func (w *Wallet) Delete() error {
	db := GetConn()
	defer db.Close()
	db = db.Delete(w)
	return db.Error
}

// GetWalletByAdresses get wallet from adresses.
func GetWalletByAdresses(Address string) (Wallet, error) {
	db := GetConn()
	defer db.Close()
	toreturn := Wallet{}
	err := db.Model(&Wallet{}).Where(
		"address = ?",
		Address,
	).Find(&toreturn)
	return toreturn, err.Error
}

// GetAllWallets return all Wallet in db
func GetAllWallets() ([]Wallet, error) {
	db := GetConn()
	defer db.Close()
	var Wallets []Wallet
	res := db.Model(Wallet{}).Find(&Wallets)
	if res.Error != nil {
		return Wallets, res.Error
	}
	return Wallets, nil
}
