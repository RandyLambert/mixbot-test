package main

import (
	"context"
	"log"

	sdk "github.com/fox-one/mixin-sdk"
)

//Read asset fee by asset_id.
//按asset_id读取资产费用。
func doAssetFee(ctx context.Context, user *sdk.User) {
	assetID := "43d61dcd-e413-450d-80b8-101d5e903357"
	fee, err := user.ReadAssetFee(ctx, assetID)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("asset fee", fee)

	if fee.AssetID != assetID {
		log.Panicf("fee asset should be %s but get %s", assetID, fee.AssetID)
	}

	if len(fee.Amount) == 0 {
		log.Panicln("empty fee amount") //没有剩余金额
	}
}
//判断地址是否有效
func validateAsset(asset *sdk.Asset) {
	if len(asset.Destination) == 0 {
		log.Panicln("empty destination", asset)
	}

	if asset.Balance.IsNegative() {
		log.Panicln("invalid balance")
	}
}

func doAsset(ctx context.Context, user *sdk.User) string {
	assetID := "965e5c6e-434c-3fa9-b780-c50f43cd955c"
	//ReadAsset get asset info, including balance, address info, etc.
	//ReadAsset 获取资产信息，包括余额、地址信息等。
	asset, err := user.ReadAsset(ctx, assetID)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("asset", asset)

	if asset.AssetID != assetID { //判断是否获取正确
		log.Panicf("asset should be %s but get %s\n", assetID, asset.AssetID)
	}

	validateAsset(asset)
	return asset.Destination //返回只资产的充值地址
}

func doAssets(ctx context.Context, user *sdk.User) {
	assets, err := user.ReadAssets(ctx)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("assets", assets)
}
