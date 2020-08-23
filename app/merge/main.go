package main

import (
	backpackItems "backend/app/service/backpack/items/model"
	backpackUser "backend/app/service/backpack/user/model"
	account "backend/app/service/user/account/model"
	money "backend/app/service/user/money/model"
	vip "backend/app/service/user/vip/model"
	"backend/pkg/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

const (
	MaxProcess = 50
)

var db *gorm.DB
var createChan = make(chan interface{}, MaxProcess)
var sem = make(chan int, MaxProcess)
var wg *sync.WaitGroup

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)
	wg = &sync.WaitGroup{}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
		Logger:                 newLogger,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(MaxProcess)

	go Server()

	fmt.Println("Start to merge!")
	now := time.Now().Unix()
	InitDB()
	UserMerge()
	SkinMerge()
	wg.Wait()
	fmt.Printf("All merged! Used time: %ds\n", time.Now().Unix()-now)
}

func SkinMerge() {
	// skin
	fmt.Println("Start to merge skins data!")

	rows, err := db.Model(&Skin{}).Rows()
	if err != nil {
		panic(err)
	}

	skininfo := make(map[string]*Skin)
	skinowners := make(map[int64][]int32)
	count := int32(1)
	for rows.Next() {
		var skin Skin
		attr := backpackItems.Attributes{}
		err := db.ScanRows(rows, &skin)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// skin data
		skin.ID = count
		skininfo[skin.UID] = &skin
		attr["skin_path"] = skin.Model
		attr["arm_path"] = skin.Arm
		attr["uid"] = skin.UID

		// Owners
		if len(skin.Owners) > 0 {
			owners := make([]int64, 0)
			err = json.Unmarshal(skin.Owners, &owners)
			if err != nil {
				fmt.Println("Failed to merge Skin owners:", skin.UID, err)
			}
			for _, v := range owners {
				skinowners[v] = append(skinowners[v], skin.ID)
			}
		}

		pb := &backpackItems.Item{
			ID:          skin.ID,
			Name:        skin.Name,
			Description: "",
			Type:        1,
			Price:       0,
			Discount:    1,
			Available:   false,
			Attributes:  []byte("{}"),
		}
		err = pb.SetAttributes(&attr)
		if err != nil {
			fmt.Println("Failed to merge Skin:", skin.UID, err)
			continue
		}
		Create(pb)
		count++
		fmt.Fprintf(os.Stdout, "Merged %d skins\r", count)
	}

	rows.Close()

	// user
	fmt.Println("Start to merge skin users data!")

	rows, err = db.Model(&SkinStore{}).Where("skin IS NOT NULL").Rows()

	if err != nil {
		panic(err)
	}

	count = 0
	for rows.Next() {
		var user SkinStore
		err := db.ScanRows(rows, &user)
		if err != nil {
			fmt.Println(err)
			continue
		}

		skins, err := user.GetSkins()
		if err != nil {
			fmt.Println("Failed to merge UID:", user.UID, err)
			continue
		}

		// merge own skin
		var items backpackUser.Items
		now := time.Now().Unix()
		for _, v := range *skins {
			skin, ok := skininfo[v.UID]
			leftTime := v.Time - now
			// invalid skin or expired skin
			if !ok || (leftTime <= 0 && v.Time != 0) {
				continue
			}
			items = append(items, &backpackUser.Item{
				ID:       skin.ID,
				Amount:   0,
				ExprTime: leftTime,
			})
		}
		// 独占
		if _, ok := skinowners[int64(user.UID)]; ok {
			for _, v := range skinowners[int64(user.UID)] {
				items = append(items, &backpackUser.Item{
					ID:       v,
					Amount:   0,
					ExprTime: 0,
				})
			}
		}

		// jump when user do not have items
		if len(items) == 0 {
			continue
		}

		bp := &backpackUser.User{
			UID: int64(user.UID),
		}
		bp.AddItems(&items)
		bpModel, err := bp.GetModel()
		if err != nil {
			fmt.Println("Failed to merge UID:", user.UID, err)
			continue
		}
		Create(bpModel)
		count++
		fmt.Fprintf(os.Stdout, "Merged %d users\r", count)
	}

	fmt.Println("Merged", count, "user")

	rows.Close()
}

func UserMerge() {
	fmt.Println("Start to merge user data!")
	rows, err := db.Model(&User{}).Rows()
	if err != nil {
		panic(err)
	}

	count := 0
	for rows.Next() {
		var user User
		err := db.ScanRows(rows, &user)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// merge
		Create(&account.Info{
			UID:       int64(user.UID),
			SteamID:   user.Steamid,
			Username:  user.Username,
			FirstJoin: int64(user.Firstjoin),
		})
		if user.Money > 0 {
			Create(&money.Money{
				UID: int64(user.UID),
				RMB: user.Money,
			})
		}
		if user.Vippoint > 0 || user.Vipexpired > time.Now().Unix() {
			Create(&vip.VIP{
				UID:        int64(user.UID),
				Point:      int(user.Vippoint),
				ExpireDate: user.Vipexpired,
			})
		}
		count++
		fmt.Fprintf(os.Stdout, "Merged %d users\r", count)
	}
	fmt.Println("Merged", count, "user")
	rows.Close()
}

func Create(value interface{}) {
	createChan <- value
	wg.Add(1)
}

func Server() {
	for {
		sem <- 1
		go func() {
			db.Create(<-createChan)
			<-sem
			wg.Done()
		}()
	}
}

func InitDB() {
	if err := db.AutoMigrate(&account.Info{}); err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&money.Money{}); err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&vip.VIP{}); err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&backpackUser.UserModel{}); err != nil {
		fmt.Println(err)
	}
	if err := db.AutoMigrate(&backpackItems.Item{}); err != nil {
		fmt.Println(err)
	}
}
