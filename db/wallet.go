package db

type Wallet struct {
	Model
	Wallet string `gorm:",not null" json:"wallet"`
}

// NewWallet returns a Wallet pointer.
func NewWallet(wallet string) *Wallet {
	return &Wallet{Wallet: wallet}
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
func GetWalletByAdresses(wallet string) (Wallet, error) {
	db := GetConn()
	defer db.Close()
	toreturn := Wallet{}
	err := db.Model(&Wallet{}).Where(
		"wallet = ?",
		wallet,
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
