package initialize

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// InitializeData 初始化测试数据
func InitializeData() {
	// 初始化游戏数据
	initializeGames()

	// 初始化活动数据
	initializeActivities()

	// 初始化分类数据
	initializeCategories()

	// 初始化游戏分类数据
	initializeGameCategories()

	// 初始化陪玩专家数据
	initializePlaymates()

	// 初始化社区帖子数据
	initializeCommunityPosts()

	// 初始化奖励订单数据
	initializeRewardOrders()
}

// 初始化游戏数据
func initializeGames() {
	var count int64
	global.GVA_DB.Model(&model.Game{}).Count(&count)
	if count > 0 {
		return
	}

	games := []model.Game{
		{Name: "王者荣耀", Icon: "shield", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCO6Fv1fyGzYvj3dZ-H2AFzfu7X6Y6_yA1taG4ZuIR2PMnd96dzALMWRYtklMtial0VLZJtSVsAP5y9m9b_xwTQoq2zVnAq5fBxhcejRsjtv0IFfLF3aOhYMHtKueVSCXz0ow78iQTjbi5HbOPPPjqeZ2tC6L9LrFZ7WoKghDYAwKmcr9P1XBCv0W5pbF7hZ5hbHCBu8G7UM2eH-_Buh3zy9Zts0inQl1gewpndP5Cqfi6i-yux-QSrRAebo6d1Z5tXpcefUZCMnE0"},
		{Name: "LOL", Icon: "videogame_asset", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuDwJMK-vPtlCamAbtTRphsnIFcd1mWDzm1_T7KpzrxayEqLvesNtlHnzdM4rYKM0uo9g2_54Jj3Y0EPW5FB1sDjdErswcmdN1SmvaBqTh5bJD9-J5ucaAwF12MByzHgfaBX5CqrBM321UsnOg5GpZieNcmb7SaK-iUOOlTN_jvoYnW1BDpgm3eWB-0IPh1akuOUcitl05xf5FC8iPrAgKeciCRH1Odd9DYxPNu8sE8cfOFkv7uk_HsPwBjE-FUF43lV3AZE7uLVp0I"},
		{Name: "Valorant", Icon: "target", Category: "shooting", Image: "https://neeko-copilot.bytedance.net/api/text2image?prompt=valorant%20game%20icon&size=400x400"},
		{Name: "原神", Icon: "auto_awesome", Category: "entertainment", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCdUkaUZGDBcCAcjxT3f9Z7kvGkPdT3w8bqDluc3iqSErfCGRjzL_0ZOcLQsR9kJ-B48g4oD0-LYMpVyMizWjSaGuOrv4IJuxr7Q9sM46dgvZLmFBazQNCInUgQSWWf7Ei_hTLzJ9zbHSg6XdbfMoeKpo2Kx7RdCyKulFDPeds0dbsHSSJqMoL89znPuERLPN2vXgzEH1c-UjXwpi61RYvSGU9GsNswjgAbWeq6oGsbLw6bfZC3vTxlIu_wBCAay7PeYyYLCQf_moM"},
		{Name: "绝地求生", Icon: "sports_esports", Category: "shooting", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuBvdKZ4vBDmJiGGoIdY5uN8vk6NP2skn5qTAfHBUjn_XhnvxhinOjbExV7IZTrLvHYx_cM7mnDxDdHV7liBkAmlFCHP1FjfQXAT9JFOFL-bW7VvCpDV8ALyyZQxWzGYxF_tBHLcywbN52GPzD8b_i9EG_geNM22_Ry3-u7yFzQx5X7gOB9M-HNmDCLsJJRyWYYYTwWyDSzv_0TPVZ5LBNP67K2Q2jfbzRMqsISc8gaxeDYUD1WoAt8P6AIc0EmqW3TKxYpUBjpmciE"},
		{Name: "CS:GO", Icon: "fitness_center", Category: "shooting", Image: "https://neeko-copilot.bytedance.net/api/text2image?prompt=csgo%20game%20icon&size=400x400"},
		{Name: "英雄联盟", Icon: "videogame_asset", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuDwJMK-vPtlCamAbtTRphsnIFcd1mWDzm1_T7KpzrxayEqLvesNtlHnzdM4rYKM0uo9g2_54Jj3Y0EPW5FB1sDjdErswcmdN1SmvaBqTh5bJD9-J5ucaAwF12MByzHgfaBX5CqrBM321UsnOg5GpZieNcmb7SaK-iUOOlTN_jvoYnW1BDpgm3eWB-0IPh1akuOUcitl05xf5FC8iPrAgKeciCRH1Odd9DYxPNu8sE8cfOFkv7uk_HsPwBjE-FUF43lV3AZE7uLVp0I"},
		{Name: "和平精英", Icon: "sports_esports", Category: "shooting", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuBvdKZ4vBDmJiGGoIdY5uN8vk6NP2skn5qTAfHBUjn_XhnvxhinOjbExV7IZTrLvHYx_cM7mnDxDdHV7liBkAmlFCHP1FjfQXAT9JFOFL-bW7VvCpDV8ALyyZQxWzGYxF_tBHLcywbN52GPzD8b_i9EG_geNM22_Ry3-u7yFzQx5X7gOB9M-HNmDCLsJJRyWYYYTwWyDSzv_0TPVZ5LBNP67K2Q2jfbzRMqsISc8gaxeDYUD1WoAt8P6AIc0EmqW3TKxYpUBjpmciE"},
		{Name: "Dota 2", Icon: "videogame_asset", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCig4vCjf8ynERRlJw6j7E-3ZfKnw3ksy3SLjVpnw99A5bjjtVmmZlDcRGgdDBVojT-NlTqiZlBWwqSKxmKeF1aPZzRuW4w9Z8HO0THinpJiLI2TpCy_uMtiVy2l0wIIqGj2yus3EC6mUevTVsHyo2bqqvBCKW-4KtNnENw1eGB07txvjLJTGb3zpzgmAV5am0uyzCz9IJbe7zOo9HRzgvkOS5mmL2g2mZJUTenpWh05M_usWd43ahAPDLisvmnDmUP7F4FgyP5730"},
		{Name: "永劫无间", Icon: "sports_esports", Category: "competitive", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuDQJJK1QQu5hytrrQz6-9PVcDQ_U5A5aFb5aKsG3nc4hta6PDVhkEuZqbQ9-9RE4huLdPSZb7w9-d65-ormxf6PLyrtY197b3IP8--uW7vVkwsPX9FnWtiQmusWN24Rxo7GI1vm8PS0q2vE6E9wI0j_J4HhWxALc7cV3dh4g6uw7PsYJ8_JxKqOvV1NUn1P-4arJV6psl5XStZc01itaEqm7KKkKpZOaxNqQhDLMlxli0poz7fcvoEPgDynsti4DVYLwiXzLV4OEyk"},
		{Name: "决战平安京", Icon: "videogame_asset", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuAAb-x90DK6MePE4PFHS7KgUv8Z0fWx3AAV_aSUBD9bgnPtp_DDOV0QcorEuoF9L8EEauuGUnerUPtt7UuFYcat1rXsFWECBpum8Sv-PopUh1cZh-XeVcAw6hEKnxTVcUfK3Z7KYuhWRupQw1Al_4QPLI1Izlb6ADoHV7_ZsTgLXUwExllC9KZji9o6fd9MuOBtSdM0KWdbCR473FQvnjNUZdXHCINeS8ndfvdFZZe6hvgx16OurtkhZkqYC3lTuH-fDBBYSBX7UV0"},
		{Name: "第五人格", Icon: "sports_esports", Category: "entertainment", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCLCwcS754hCh6ggV3wdNDwshCz9TOfWbVrRUuppWZhYjJ16ufic_fBkx9wwSCqYDXf_xyrEcyAnAv7CnBZ1ahLPHFRhU7W8UW_eWswyvg518Te0HFoJPrHQOVeb-Rf6EWLCZh5RvMqiqKY9VsqH_R7xy0A1uq0EZJEOU-M0CwRIPkQbDj2ZCHz1XgcZLwmQ3Mjpt58W8JiMZwGpPvksGwPNYZatlHvuDluE1FxB6Mg30HfpxovYJhQlCVua5iJMPF732KijA2z-YA"},
		{Name: "英雄联盟手游", Icon: "videogame_asset", Category: "moba", Image: "https://lh3.googleusercontent.com/aida-public/AB6AXuAoF3F1MQOLKt8g5tJmSwhNEBH72gw-ndT2ii2kf5h-Rx8x5--0z3McICzfZFJIy3Da3VXizRgM73JCWLAW4CFK_bWIWKQkH-I1iuMgmo19VLLkBYvCBi-a4fGhxUrWsQMC5HD2aF_48366vg4V5ccJy60tbg_b0KnU-S3BOf-TwKeEDxULpzZg8297caghuE3J3LRvIf9THvqDF59qkxwANu5YDuX5aRCnvcBa26nlCt_-0eZMSngG-FlsuGSvicmwFScS0sAzI7A"},
	}

	global.GVA_DB.Create(&games)
}

// 初始化活动数据
func initializeActivities() {
	var count int64
	global.GVA_DB.Model(&model.Activity{}).Count(&count)
	if count > 0 {
		return
	}

	activities := []model.Activity{
		{
			Title:       "新人专享首单优惠",
			Description: "首次下单立减10元，限时优惠",
			Discount:    10,
			Type:        "discount",
			ValidUntil:  time.Date(2026, 4, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			Title:       "周末狂欢活动",
			Description: "周末下单享8折优惠",
			Discount:    20,
			Type:        "weekend",
			ValidUntil:  time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	global.GVA_DB.Create(&activities)
}

// 初始化分类数据
func initializeCategories() {
	var count int64
	global.GVA_DB.Model(&model.Category{}).Count(&count)
	if count > 0 {
		return
	}

	categories := []model.Category{
		{Name: "热门游戏", Icon: "🎮"},
		{Name: "新手教学", Icon: "📚"},
		{Name: "陪练上分", Icon: "🏆"},
		{Name: "娱乐开黑", Icon: "🎉"},
		{Name: "赛事解说", Icon: "📺"},
	}

	global.GVA_DB.Create(&categories)
}

// 初始化游戏分类数据
func initializeGameCategories() {
	var count int64
	global.GVA_DB.Model(&model.GameCategory{}).Count(&count)
	if count > 0 {
		return
	}

	gameCategories := []model.GameCategory{
		{Name: "moba", Label: "MOBA"},
		{Name: "shooting", Label: "射击"},
		{Name: "competitive", Label: "竞技"},
		{Name: "entertainment", Label: "娱乐"},
		{Name: "online", Label: "网游"},
		{Name: "console", Label: "主机"},
	}

	global.GVA_DB.Create(&gameCategories)
}

// 初始化陪玩专家数据
func initializePlaymates() {
	var count int64
	global.GVA_DB.Model(&model.Playmate{}).Count(&count)
	if count > 0 {
		return
	}

	playmates := []model.Playmate{
		{
			UserID:      1,
			Nickname:    "林间小鹿",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I",
			Rating:      5.0,
			Price:       45,
			Likes:       128,
			Tags:        "王者,萌妹",
			IsOnline:    true,
			Game:        "honor_of_kings",
			Rank:        "challenger",
			Gender:      "female",
			Description: "",
			Level:       45,
			Title:       "王者荣耀 · 荣耀之巅",
		},
		{
			UserID:      2,
			Nickname:    "影子猎手",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o",
			Rating:      4.9,
			Price:       60,
			Likes:       95,
			Tags:        "国服第一,技术流",
			IsOnline:    true,
			Game:        "lol",
			Rank:        "challenger",
			Gender:      "male",
			Description: "",
			Level:       40,
			Title:       "LOL · 峡谷之巅",
		},
		{
			UserID:      3,
			Nickname:    "安妮喵呜",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg",
			Rating:      4.8,
			Price:       55,
			Likes:       156,
			Tags:        "温柔语聊,绝地求生",
			IsOnline:    false,
			Game:        "pubg",
			Rank:        "master",
			Gender:      "female",
			Description: "累计陪玩 1,200+ 小时",
			Level:       35,
			Title:       "绝地求生 · 吃鸡高手",
		},
		{
			UserID:      4,
			Nickname:    "星野 Kyo",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuA3Yx21l3XQH58JEjdPvp2NeoI5LIs_51ynV3rForFFjwT3Hd5AqMSy-sxYD_dlyN682W91abmSgg8KAw9tpHslBqThqBE3aE1ZVVOsHMHttJaF7wdtEpDJ2OL28yGHnfz11wPG1Jw2fXoB8C6dHlBSmulomn9y3CFDd8uDRgc2wm8DxSw67mMZ3pZTXdnJa4MLoz4Dl06hB9dHtby_V56tqiQv_vAAw9oI9xP2_AoxO74HKUfN1fVEaKcK9T4LY_KB9LY0itiXfHU",
			Rating:      4.9,
			Price:       50,
			Likes:       203,
			Tags:        "国服第一野王,性格超好",
			IsOnline:    true,
			Game:        "honor_of_kings",
			Rank:        "challenger",
			Gender:      "male",
			Description: "累计陪玩 1,200+ 小时",
			Level:       50,
			Title:       "王者荣耀 · 野王",
		},
		{
			UserID:      5,
			Nickname:    "皮卡丘爱吃肉",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuA255ssriqx239upv8-12k7zCZRVehwBMDwNbxBDVWFHTXeU9kDiEO9QDelxUR7qMzQ75J-4ro7_ZMZwik9PTEdzzd455C4MqYkMBHBgZUIIE3UUxXuZUzd9hBOzXfu3LsQIb_K1cS4c8zbYJaZXj4FIJciSMjb3OZ7mNFzGgGCjvimrN_9WQg4FONTf5MEWUNC7cIxE0HJ_2v3fuqE_KzIwg_BDneDCMgvQT--ApEg5QvmzBze7lomDKVKJDidgKcF1y0voJA1l",
			Rating:      5.0,
			Price:       40,
			Likes:       87,
			Tags:        "萌新友好,LOL",
			IsOnline:    true,
			Game:        "lol",
			Rank:        "diamond",
			Gender:      "female",
			Description: "",
			Level:       30,
			Title:       "LOL · 钻石选手",
		},
		{
			UserID:      6,
			Nickname:    "游戏小王子",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuC0dK_MZqKKAmyA4yP_BkKpGsS_W87iCSLeuAAmwt_w2J0Fv2TJ3nFinGP1J3GSQx67aimew8FhrXUj4yr1rs9GdeQQ0HoLz7PYrcmT-menQ1ggX6R3fg-gRg1AOiAechpMCNklo7nlzYm6oTif3LxCbdYzd3J-GAlZr6SEyxpjOmVPDUFE4dwfMogxpTTQg-RVyvcN33jdsEkwf8vjjX4NIE1tfsILkkmLj6dL8Mu-pfvMf__Z-Ik58KyaejQd1RAt3s-oj5nsyo8",
			Rating:      4.7,
			Price:       35,
			Likes:       64,
			Tags:        "Valorant,技术流",
			IsOnline:    false,
			Game:        "valorant",
			Rank:        "master",
			Gender:      "male",
			Description: "",
			Level:       25,
			Title:       "Valorant · 大师",
		},
		{
			UserID:      7,
			Nickname:    "小甜甜",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuAo3hiZMbuWCmrkmqDp0MNg-nzMlqn4AgGNgFofJ-IJJ59AeEGJZox-r1xg8drTN14GtBID78Yfm0NN6Wl-MX_ZHWaF8-LdFumF5to5bPc46yM5zIaucC3a4KP-Y2-i2tCwQd-duGHsg9F8AzLemDvG9s_6W3hw-pjjJN-PKkQI7oILcDHGkalxMkbZWOmaqIaik3owuV2Ghbrqr3MraHwibw8jkhxJWFQhPSpHM7em0RfJhdOUnqGVa3qBY2EA9DWUXnYNrTFFiM4",
			Rating:      4.9,
			Price:       48,
			Likes:       178,
			Tags:        "原神,声音甜美",
			IsOnline:    true,
			Game:        "genshin",
			Rank:        "adventure_rank_55",
			Gender:      "female",
			Description: "专注原神陪玩，熟悉所有角色和副本",
			Level:       40,
			Title:       "原神 · 冒险等级55",
		},
		{
			UserID:      8,
			Nickname:    "电竞女神",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuBU8XrMhk3VFwj3Ceer0yCVb-HelGKX5ryLoHAYthnekinYCfvHoJ83xSPEZdrL2tht2CTf_d1atj0kQiKVMY41s8kFOBgY2l5a9dPvoP6yXh3HyA9pdom7W1PkI1l7drYVsVSEeg-BnjsOK2tD_lvHVqzF3VZCAhg6pbcyZj11rhzX6V52RT4jlbNYqEKBRxP818vVewrnT3E6phVAdGXO9zQIWGaIvWWk6pXEzjpNwOh3xDn8FNjv-sGQOhsOt1srbfvB0MFCo-A",
			Rating:      5.0,
			Price:       65,
			Likes:       234,
			Tags:        "CS:GO,女战神",
			IsOnline:    true,
			Game:        "csgo",
			Rank:        "global_elite",
			Gender:      "female",
			Description: "职业选手退役，擅长各种地图",
			Level:       50,
			Title:       "CS:GO · 全球精英",
		},
		{
			UserID:      9,
			Nickname:    "绝地枪神",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuDwJXWYv1t2M1S8jlZ3nJwEK4xXnZfwXDgO_ryYTFAAPD3NJ6pmTUhpeJsra6qrl4sekbm3zZ7874XLIqf6W__ReW-xU87woTXwpcGzN5Q3b-1FAHT-7Q9ZJvEhBpGxJwr8JPIoII8MLGmKtsx2rIr1BNvzlerT6UrU1DTTnRE1vevhiuXzR1drRXXUNZdGjDRSQmxQUYJwpOCdchivrmLpaAcpR5c8eCBholl_o0fI_vrf4Lu_oTTxHGoqzJuNdhfVVbxvwguDqZQ",
			Rating:      4.8,
			Price:       52,
			Likes:       112,
			Tags:        "绝地求生,吃鸡高手",
			IsOnline:    false,
			Game:        "pubg",
			Rank:        "conqueror",
			Gender:      "male",
			Description: "K/D比3.0+，擅长各种地形作战",
			Level:       45,
			Title:       "绝地求生 · 征服者",
		},
		{
			UserID:      10,
			Nickname:    "王者小能手",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuA3Yx21l3XQH58JEjdPvp2NeoI5LIs_51ynV3rForFFjwT3Hd5AqMSy-sxYD_dlyN682W91abmSgg8KAw9tpHslBqThqBE3aE1ZVVOsHMHttJaF7wdtEpDJ2OL28yGHnfz11wPG1Jw2fXoB8C6dHlBSmulomn9y3CFDd8uDRgc2wm8DxSw67mMZ3pZTXdnJa4MLoz4Dl06hB9dHtby_V56tqiQv_vAAw9oI9xP2_AoxO74HKUfN1fVEaKcK9T4LY_KB9LY0itiXfHU",
			Rating:      4.7,
			Price:       38,
			Likes:       76,
			Tags:        "王者荣耀,全能选手",
			IsOnline:    true,
			Game:        "honor_of_kings",
			Rank:        "mythic",
			Gender:      "male",
			Description: "所有位置都能打，擅长团队配合",
			Level:       35,
			Title:       "王者荣耀 · 全能选手",
		},
		{
			UserID:      11,
			Nickname:    "LOL大神",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o",
			Rating:      4.9,
			Price:       68,
			Likes:       189,
			Tags:        "LOL,峡谷之巅",
			IsOnline:    true,
			Game:        "lol",
			Rank:        "master",
			Gender:      "male",
			Description: "峡谷之巅大师，擅长中单位置",
			Level:       45,
			Title:       "LOL · 峡谷之巅大师",
		},
		{
			UserID:      12,
			Nickname:    "萌妹陪玩",
			Avatar:      "https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I",
			Rating:      4.8,
			Price:       42,
			Likes:       145,
			Tags:        "声音好听,娱乐陪玩",
			IsOnline:    false,
			Game:        "honor_of_kings",
			Rank:        "diamond",
			Gender:      "female",
			Description: "性格开朗，善于活跃气氛",
			Level:       30,
			Title:       "王者荣耀 · 钻石选手",
		},
	}

	global.GVA_DB.Create(&playmates)

	// 为每个陪玩专家创建技能
	for _, playmate := range playmates {
		skills := []model.PlaymateSkill{
			{
				PlaymateID:  playmate.ID,
				Name:        playmate.Game,
				Price:       playmate.Price,
				Level:       playmate.Rank,
				Description: "排位赛/巅峰赛",
			},
			{
				PlaymateID:  playmate.ID,
				Name:        "语音连麦",
				Price:       playmate.Price * 0.8,
				Level:       "专业",
				Description: "聊天/指导/开黑",
			},
		}
		global.GVA_DB.Create(&skills)

		// 创建语音介绍
		voiceIntro := model.PlaymateVoiceIntroduction{
			PlaymateID: playmate.ID,
			URL:        "",
			Duration:   "00:15",
		}
		global.GVA_DB.Create(&voiceIntro)
	}
}

// 初始化社区帖子数据
func initializeCommunityPosts() {
	var count int64
	global.GVA_DB.Model(&model.CommunityPost{}).Count(&count)
	if count > 0 {
		return
	}

	communityPosts := []model.CommunityPost{
		{
			UserID:    1,
			Content:   "今天在峡谷遇到一个超级温柔的辅助，操作意识拉满！有人想一起组队排位吗？坐标艾欧尼亚，主玩AD。 ✨",
			Images:    "https://lh3.googleusercontent.com/aida-public/AB6AXuBU8XrMhk3VFwj3Ceer0yCVb-HelGKX5ryLoHAYthnekinYCfvHoJ83xSPEZdrL2tht2CTf_d1atj0kQiKVMY41s8kFOBgY2l5a9dPvoP6yXh3HyA9pdom7W1PkI1l7drYVsVSEeg-BnjsOK2tD_lvHVqzF3VZCAhg6pbcyZj11rhzX6V52RT4jlbNYqEKBRxP818vVewrnT3E6phVAdGXO9zQIWGaIvWWk6pXEzjpNwOh3xDn8FNjv-sGQOhsOt1srbfvB0MFCo-A,https://lh3.googleusercontent.com/aida-public/AB6AXuBq4AH6Y7D9L3js3AeN_vpUSza4-PdEYsZbo6sLRP3jUY5UwegYPx3xB9NLMbxzijv13O-siVI5RljmMLqN_Gr5WevHBvslpJbeCO97PCdtgKwNIXBrfQfGanIkq-uktmkkEyoBN9MXdNEwYcjzTObAUzdyii37QkpNmj9bZP7Bcl5uaq012JA-ku0hKny8SxUd9oCA63jFafkb_YzLfo3nDoPbHw27pT7fpKdYQd2sZiXMQL7sMljpCn5eqqS3c5aJtgaLu6pHQpU",
			Likes:     1200,
			Comments:  348,
			Game:      "英雄联盟",
		},
		{
			UserID:    2,
			Content:   "终于抽到雷电将军了！大家快来吸吸欧气，顺便问下配队方案，目前有行秋和班尼特。🌸",
			Images:    "https://lh3.googleusercontent.com/aida-public/AB6AXuDpMESPnsbi4t03wXNISsduDaneNSUQYy8CFJQ6WGwywyqmkUeGY562sFOkLCu6pFjjqqs0sbxMrg9JwqOHdCcO-aFFA3Uh2XcpKXedBl_JMAAdIZ19iCVx6GQyYCwO0zuk9xRpZOHHqzfIgD7XyQBa6MYkyr_AyukGeEUUYDV7giFJThyUVuogxAZZYyU6SgGNT62RNNxPFRg1-urZjonnT0y8Zh-OZyb_7b7dlBnbONQqjw2jCO-qEHaONKJ-0zPtkTmZXZzvCrc",
			Likes:     856,
			Comments:  120,
			Game:      "原神",
		},
		{
			UserID:    3,
			Content:   "\"今晚吃鸡，大吉大利！\" 🍗\n好久没这么酣畅淋漓地打过药包了，感谢各位队友的灵性拉枪。",
			Images:    "",
			Likes:     233,
			Comments:  42,
			Game:      "绝地求生",
		},
	}

	global.GVA_DB.Create(&communityPosts)
}

// 初始化奖励订单数据
func initializeRewardOrders() {
	var count int64
	global.GVA_DB.Model(&model.RewardOrder{}).Count(&count)
	if count > 0 {
		return
	}

	rewardOrders := []model.RewardOrder{
		{
			UserID:         1001,
			Game:           "英雄联盟",
			Content:        "寻找钻石以上段位的陪玩，一起冲大师！",
			Reward:         100,
			PaymentMethod:  "prepay",
			Status:         "available",
			Tags:           "钻石,排位,上分",
		},
		{
			UserID:         1002,
			Game:           "王者荣耀",
			Content:        "寻找荣耀王者陪玩，教我玩貂蝉",
			Reward:         80,
			PaymentMethod:  "postpay",
			Status:         "available",
			Tags:           "荣耀王者,貂蝉,教学",
		},
		{
			UserID:         1003,
			Game:           "绝地求生",
			Content:        "寻找KD3.0以上的陪玩，一起吃鸡",
			Reward:         120,
			PaymentMethod:  "prepay",
			Status:         "available",
			Tags:           "KD3.0,吃鸡,技术流",
		},
		{
			UserID:         1004,
			Game:           "原神",
			Content:        "寻找原神陪玩，一起打深渊12层",
			Reward:         60,
			PaymentMethod:  "postpay",
			Status:         "available",
			Tags:           "原神,深渊,攻略",
		},
		{
			UserID:         1005,
			Game:           "CS:GO",
			Content:        "寻找Global Elite陪玩，练习枪法",
			Reward:         150,
			PaymentMethod:  "prepay",
			Status:         "available",
			Tags:           "Global Elite,枪法,练习",
		},
	}

	global.GVA_DB.Create(&rewardOrders)
}
