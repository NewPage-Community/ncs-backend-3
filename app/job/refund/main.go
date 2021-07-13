package main

import (
	itemsSrv "backend/app/service/backpack/items/api/grpc"
	userSrv "backend/app/service/backpack/user/api/grpc"
	accountSrv "backend/app/service/user/account/api/grpc"
	moneySrv "backend/app/service/user/money/api/grpc"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	MaxProcess = 50
)

var items itemsSrv.ItemsClient
var user userSrv.UserClient
var money moneySrv.MoneyClient
var account accountSrv.AccountClient

var refundItems []int

func main() {
	// Get refund items
	refundStr := os.Getenv("REFUND_ITEMS")
	itemsStr := strings.Split(refundStr, ",")
	for _, v := range itemsStr {
		item, err := strconv.Atoi(v)
		if err == nil {
			refundItems = append(refundItems, item)
		}
	}

	// Init grpc
	items = itemsSrv.InitClient(itemsSrv.ServiceAddr)
	user = userSrv.InitClient(userSrv.ServiceAddr)
	money = moneySrv.InitClient(moneySrv.ServiceAddr)
	account = accountSrv.InitClient(moneySrv.ServiceAddr)
	defer func() {
		itemsSrv.Close()
		userSrv.Close()
		moneySrv.Close()
		accountSrv.Close()
	}()

	fmt.Printf("Start to refund items\n%s", refundStr)
	refund()
}

func refund() {
	// Get items info
	itemsLib, err := items.GetItems(context.Background(), &itemsSrv.GetItemsReq{
		Type: 0,
	})
	if err != nil {
		panic(err)
	}
	// Load items info
	refundItemsInfo := make(map[int]*itemsSrv.Item)
	for _, id := range refundItems {
		for _, item := range itemsLib.Items {
			if int32(id) == item.Id {
				refundItemsInfo[id] = item
			}
		}
	}

	// Get all UID
	accounts, err := account.GetAllUID(context.Background(), &accountSrv.GetAllUIDReq{})
	if err != nil {
		panic(err)
	}

	// Routine
	wg := &sync.WaitGroup{}
	sem := make(chan int, MaxProcess)

	// Refund all
	for _, uid := range accounts.Uid {
		for _, item := range refundItemsInfo {
			// Routine queue
			sem <- 1
			wg.Add(1)

			item := item
			go func() {
				// Routine done
				defer func() {
					<-sem
					wg.Done()
				}()

				ctx := context.Background()

				_, err = user.RemoveItem(ctx, &userSrv.RemoveItemReq{
					Uid:  uid,
					Item: &userSrv.Item{Id: item.Id},
				})
				if err != nil {
					fmt.Printf("[ERROR] remove %d item (%d) faild\n", uid, item.Id)
					return
				}
				if item.Price == 0 {
					return
				}
				_, err = money.Give(ctx, &moneySrv.GiveReq{
					Uid:    uid,
					Money:  item.Price,
					Reason: fmt.Sprintf("皮肤 %s 退款", item.Name),
				})
				if err != nil {
					fmt.Printf("[ERROR] refund %d item (%d - %d) faild\n", uid, item.Id, item.Price)
				} else {
					fmt.Printf("[SUCCESS] refund %d item (%d - %d)", uid, item.Id, item.Price)
				}
			}()
		}
	}

	wg.Wait()
}
