package main

import (
	"context"
	"log"
	"time"

	sdk "github.com/fox-one/mixin-sdk"
)

func doReadNetwork(ctx context.Context) {
	//读snapshots,只读10条
	snapshots, err := sdk.ReadNetwork(ctx, "965e5c6e-434c-3fa9-b780-c50f43cd955c", time.Time{}, "", 10, "")
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read network", snapshots)
}

func doUserReadNetwork(ctx context.Context, user *sdk.User) {
	//读某个usersnapshots,只读10条
	//读取Mixin Network的公共快照snapshot_id。
	snapshots, err := user.ReadNetwork(ctx, "965e5c6e-434c-3fa9-b780-c50f43cd955c", time.Time{}, "", 10)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read network", snapshots)
}
//读多条
func doReadSnapshots(ctx context.Context, user *sdk.User) {
	snapshots, err := user.ReadSnapshots(ctx, "965e5c6e-434c-3fa9-b780-c50f43cd955c", time.Time{}, "", 10)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read snapshots", snapshots)
}
//通过snapshotid读一条
func doReadSnapshot(ctx context.Context, user *sdk.User, snapshotID string) {
	snapshot, err := user.ReadSnapshot(ctx, snapshotID)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read snapshot", snapshot)
}

func doReadTransfer(ctx context.Context, user *sdk.User, traceID string) {
	//ReadTransfer读取该traceID的快照
	snapshot, err := user.ReadTransfer(ctx, traceID)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read transfer", snapshot)
}

func doReadExternal(ctx context.Context) {
	//读取外部快照
	snapshots, err := sdk.ReadExternal(ctx, "", "", "", time.Time{}, 10)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read deposit snapshots", snapshots)
}
