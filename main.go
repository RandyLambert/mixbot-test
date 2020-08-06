package main

import (
	"context"
	"github.com/fox-one/mixin-sdk"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"

	//"time"
)

//按照json格式输出
func printJSON(prefix string, item interface{}) {
	msg, _ := jsoniter.MarshalIndent(item, "", "    ")
	log.Println(prefix, string(msg))
}

const(
EOS	 = "6cfe566e-4aad-470b-8c9a-2fd35b49c68d"
CNB	 = "965e5c6e-434c-3fa9-b780-c50f43cd955c"
BTC	 = "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
ETC	 = "2204c1ee-0ea2-4add-bb9a-b3719cfff93a"
XRP	 = "23dfb5a5-5d7b-48b6-905f-3970e3176e27"
XEM	 = "27921032-f73e-434e-955f-43d55672ee31"
ETH	 = "43d61dcd-e413-450d-80b8-101d5e903357"
DASH = "6472e7e3-75fd-48b6-b1dc-28d294ee1476"
DOGE = "6770a1e5-6086-44d5-b60f-545f9d9e8ffd"
LTC	 = "76c802a2-7c88-447f-a93e-c29c9e5dd9c8"
SC	 = "990c4c29-57e9-48f6-9819-7d986ea44985"
ZEN	 = "a2c5d22b-62a2-4c13-b3f0-013290dbac60"
ZEC	 = "c996abc9-d94e-4494-b1cf-2a3fd3ac5714"
BCH	 = "fd11b6e3-0b87-41f1-a41f-f0e9b49e5bf0"
USDT = "815b0b1a-2764-3736-8faa-42d694fa620a"
//EOS的存币地址与其它的币有些不同，它由两部分组成： account_name and account tag, 如果你向Mixin Network存入EOS，你需要填两项数据： account name 是eoswithmixin,备注里输入你的account_tag,比如0aa2b00fad2c69059ca1b50de2b45569.

)

func InitConf(){
	workDir,err := os.Getwd()
	if err !=nil {
		panic(err.Error())
	}
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())

	}
}
func main() {

	InitConf() //初始化conf
	user,err := mixin.NewUser(viper.GetString("mixin.UserID"),
		viper.GetString("mixin.SessionId"),
		viper.GetString("mixin.PrivateKey"),
		viper.GetString("mixin.PinToken")) //主机器人生成user
		//                                             就是PrivateKey
		//user, err := sdk.NewUser(ClientID, SessionID, SessionKey, PINToken)
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.Background()
	publicKey := doAsset(ctx, user)//返回充值的地址

	p := viper.GetString("mixin.PinCode") //为新创建的子bot设置的密码,和原来的设置的一样
	u := doCreateUser(ctx, user, p) //创建子bot

	//doAssetFee(ctx, u) //通过asset_id读资产费用

	publicKey1 := doAsset(ctx, u)//返回子地址的充值地址

	//doAssets(ctx, u) //返回多个充值地址
	//fmt.Println("test",string(publicKey1))

	//assetID := "965e5c6e-434c-3fa9-b780-c50f43cd955c" //cnb
	assetID   := USDT
	doTransfer(ctx, user, assetID, u.UserID, "0.0001", "ping", viper.GetString("mixin.PinCode"))// user转账给u
	time.Sleep(time.Second * 5)
	snap := doTransfer(ctx, u, assetID, user.UserID, "0.0001", "pong", p) // u转账给user

	doWithdraw(ctx, user, assetID, publicKey1, "0.0001", "ping", viper.GetString("mixin.PinCode")) //从user中提现到publicKey1
	time.Sleep(time.Second * 5)
	doWithdraw(ctx, u, assetID, publicKey, "0.0001", "pong", p) //从u中提现到publicKey

	//doReadNetwork(ctx)//?? 读的网络公共Snapshots
	//{
	//"snapshot_id": "09bc17ed-b825-4bd1-8d2c-0e12aa9a023f",
	//"asset_id": "965e5c6e-434c-3fa9-b780-c50f43cd955c",
	//"created_at": "2020-06-26T05:31:49.927382Z",
	//"source": "TRANSFER_INITIALIZED",
	//"amount": "-0.0001",
	//"opening_balance": "0",
	//"closing_balance": "0",
	//"type": "snapshot",
	//"asset": {
	//"asset_id": "965e5c6e-434c-3fa9-b780-c50f43cd955c",
	//"chain_id": "43d61dcd-e413-450d-80b8-101d5e903357",
	//"asset_key": "0xec2a0550a2e4da2a027b3fc06f70ba15a94a6dac",
	//"symbol": "CNB",
	//"name": "Chui Niu Bi",
	//"icon_url": "https://mixin-images.zeromesh.net/0sQY63dDMkWTURkJVjowWY6Le4ICjAFuu3ANVyZA4uI3UdkbuOT5fjJUT82ArNYmZvVcxDXyNjxoOv0TAYbQTNKS=s128",
	//"price_usd": "0",
	//"change_usd": "0",
	//"balance": "0"
	//}
	doUserReadNetwork(ctx, user) //?? 读u的网络公共的Snapshots
	//doReadSnapshots(ctx, user) // 读user的snapshots,区别是信息量不同??

	//doReadSnapshot(ctx, u, snap.SnapshotID) //按照snapshotid读取snapshot信息
	//"snapshot_id": "89ccf495-7db2-4825-8617-bf5b4ec5da4d",
	//"trace_id": "c7b9ec04-8c88-40fc-87c4-bd2772b723b6",
	//"asset_id": "815b0b1a-2764-3736-8faa-42d694fa620a",
	//"opponent_id": "a39c3ffc-c308-4c12-b117-fb7410cdbb43",
	//"created_at": "2020-06-26T05:36:16.611716Z",
	//"source": "",
	//"amount": "-0.0001",
	//"opening_balance": "0.0001",
	//"closing_balance": "0",
	//"type": "transfer"

	doReadTransfer(ctx, u, snap.TraceID) //按照TraceID读取snapshot信息

	//doReadExternal(ctx)

	//doReadNetworkInfo(ctx) //读取网络信息???
	//做交易???
	//doTransaction(ctx, user, USDT/*"965e5c6e-434c-3fa9-b780-c50f43cd955c"*/, "XINT55hZYxzrtqJsWViUbyoxytJ6RoKUZfpnSCQTbgX8fjcdQ7GwjRySLxiPMWxAMhoN6KPa7SFkyv9FQXC3fGJuKHLf3est", "0.0001", "test", viper.GetString("mixin.PinCode"))

	// Messenger

	//conversation := doCreateConversation(ctx, user)//创建会话,测试发送信息类型的消息??
	//doMessage(ctx, user, &mixin.MessageRequest{
	//	ConversationID: conversation.ConversationID,
	//	MessageID:      uuid.Must(uuid.NewV4()).String(),
	//	Category:       "PLAIN_TEXT",
	//	Data:           base64.StdEncoding.EncodeToString([]byte("Just A Test")),
	//})
	//doReadConversation(ctx, user, conversation.ConversationID)

	Handler{}.Run(ctx, user)
}
