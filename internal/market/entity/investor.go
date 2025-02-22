package entity

type Investor struct {
	ID           string
	Name         string
	AssetPostion []*InvestorAssetPosition
}

func NewInvestor(id, name string) *Investor {
	return &Investor{
		ID:           id,
		Name:         name,
		AssetPostion: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPostion = append(i.AssetPostion, assetPosition)
}

func (i *Investor) AdjustAssetPosition(assetID string, qtdShares int) {
	assetPostions := i.GetAssetPosition(assetID)
	if assetPostions == nil {
		i.AssetPostion = append(i.AssetPostion, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPostions.AddShares(qtdShares)
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPostion {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
func (iap *InvestorAssetPosition) AddShares(qtdShares int) {
	iap.Shares += qtdShares
}
