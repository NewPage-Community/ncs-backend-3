package main

import (
	itemsSrv "backend/app/service/backpack/items/api/grpc/v1"
	userSrv "backend/app/service/backpack/user/api/grpc/v1"
	accountSrv "backend/app/service/user/account/api/grpc/v1"
	moneySrv "backend/app/service/user/money/api/grpc/v1"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	MaxProcess = 10
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
	account = accountSrv.InitClient(accountSrv.ServiceAddr)
	defer func() {
		itemsSrv.Close()
		userSrv.Close()
		moneySrv.Close()
		accountSrv.Close()
	}()

	fmt.Printf("Start to refund items\n%s\n", refundStr)
	refund()
	fmt.Printf("Refund done\n")
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
		// Routine queue
		sem <- 1
		wg.Add(1)

		// 1 user 1 routine
		uid := uid
		go func() {
			// Routine done
			defer func() {
				<-sem
				wg.Done()
			}()

			isRefund := false

			for _, item := range refundItemsInfo {
				ctx := context.Background()

				// Check user have item
				res, err := user.HaveItem(ctx, &userSrv.HaveItemReq{
					Uid: uid,
					Id:  item.Id,
				})
				if err != nil {
					fmt.Printf("[ERROR] check %d item (%d) faild\n", uid, item.Id)
					fmt.Println(err)
					continue
				}
				// User do not have item
				if !res.Have {
					continue
				}

				// Remove it
				_, err = user.RemoveItem(ctx, &userSrv.RemoveItemReq{
					Uid:  uid,
					Item: &userSrv.Item{Id: item.Id},
					// Force remove unlimited item
					All: true,
				})
				if err != nil {
					fmt.Printf("[ERROR] remove %d item (%d) faild\n", uid, item.Id)
					fmt.Println(err)
					continue
				}

				isRefund = true

				// No refund if do not have price
				if item.Price == 0 {
					continue
				}

				// Refund!
				_, err = money.Give(ctx, &moneySrv.GiveReq{
					Uid:    uid,
					Money:  item.Price,
					Reason: fmt.Sprintf("退款 - %s", item.Name),
				})
				if err != nil {
					fmt.Printf("[ERROR] refund %d item (%d - %d) faild\n", uid, item.Id, item.Price)
					fmt.Println(err)
				} else {
					fmt.Printf("[SUCCESS] refund %d item (%d - %d)\n", uid, item.Id, item.Price)
				}
			}

			if isRefund {
				_, err = money.Give(context.Background(), &moneySrv.GiveReq{
					Uid:    uid,
					Money:  300,
					Reason: "回收补偿",
				})
				if err != nil {
					fmt.Printf("[ERROR] compensation %d faild\n", uid)
					fmt.Println(err)
				}
			}
		}()
	}

	wg.Wait()
}
