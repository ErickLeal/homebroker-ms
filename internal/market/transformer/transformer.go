package transformer

import (
	"github.com/ErickLeal/homebroker-ms/internal/market/dto"
	"github.com/ErickLeal/homebroker-ms/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID, "name")
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, entity.OrderType(input.OrderType))
	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}
	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  string(order.Type),
		Status:     string(order.Status),
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID:       t.BuyingOrder.Investor.ID,
			SellerID:      t.SellingOrder.Investor.ID,
			AssetID:       t.SellingOrder.Asset.ID,
			Price:         t.Price,
			Shares:        t.SellingOrder.Shares - t.SellingOrder.PendingShares,
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}
	output.TransactionsOutput = transactionsOutput
	return output
}
