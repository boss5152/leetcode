package problems

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Fish() {
	// 設定亂數種子
	rand.Seed(time.Now().UnixNano())

	// 初始化遊戲參數
	var fishNum = 10
	var fishList = make([]int, fishNum)
	var score = 0

	// 生成魚的初始位置和分數
	for i := 0; i < fishNum; i++ {
		fishList[i] = rand.Intn(10) + 1
	}

	// 遊戲開始
	fmt.Println("歡迎來到捕魚機遊戲！")
	fmt.Println("遊戲規則：輸入1~10之間的整數，代表射擊魚的位置，擊中獲得相應分數。")
	fmt.Println("遊戲結束條件：擊中所有魚或者分數達到100。")

	// 遊戲循環
	for {
		fmt.Println("魚的位置：", fishList)
		fmt.Println("當前分數：", score)

		// 等待玩家輸入
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("請輸入射擊位置：")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		shotPos, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("輸入錯誤，請輸入1~10之間的整數！")
			continue
		}
		if shotPos < 1 || shotPos > 10 {
			fmt.Println("輸入錯誤，請輸入1~10之間的整數！")
			continue
		}

		// 判斷是否擊中魚
		hitFish := false
		for i := 0; i < fishNum; i++ {
			if fishList[i] == shotPos {
				hitFish = true
				score += fishList[i]
				fishList = append(fishList[:i], fishList[i+1:]...)
				fishNum--
				break
			}
		}
		if hitFish {
			fmt.Println("擊中魚，得分", score)
		} else {
			fmt.Println("沒有擊中魚")
		}

		// 判斷遊戲是否結束
		if fishNum == 0 || score >= 100 {
			fmt.Println("遊戲結束，得分", score)
			break
		}
	}
}
