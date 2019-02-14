package gengou

import "time"

var (
	jst                 = time.FixedZone("Asia/Tokyo", 9*60*60)
	nanbokueraFrom      = time.Date(1329, time.September, 22, 0, 0, 0, 0, jst)
	nanbokueraTo        = time.Date(1394, time.August, 2, 0, 0, 0, 0, jst)
	gengoDataNanbokuEra = []Gengo{

		// from
		Gengo{
			C:    153,
			Name: "元徳（大覚寺統）",
			Kana: "げんとく",
			From: time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1331, time.September, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    154,
			Name: "元徳（持明院統）",
			Kana: "げんとく",
			From: time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1332, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    155,
			Name: "元弘",
			Kana: "げんこう",
			From: time.Date(1331, time.September, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    156,
			Name: "建武（南朝）",
			Kana: "けんむ",
			From: time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1336, time.April, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    157,
			Name: "延元（南朝）",
			Kana: "えんげん",
			From: time.Date(1336, time.April, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1340, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    158,
			Name: "興国（南朝）",
			Kana: "こうこく",
			From: time.Date(1340, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1347, time.January, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    159,
			Name: "正平（南朝）",
			Kana: "しょうへい",
			From: time.Date(1347, time.January, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1370, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    160,
			Name: "建徳（南朝）",
			Kana: "けんとく",
			From: time.Date(1370, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1372, time.April, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    161,
			Name: "文中（南朝）",
			Kana: "ぶんちゅう",
			From: time.Date(1372, time.April, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1375, time.June, 26, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    162,
			Name: "天授（南朝）",
			Kana: "てんじゅ",
			From: time.Date(1375, time.June, 26, 0, 0, 0, 0, jst),
			To:   time.Date(1381, time.March, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    163,
			Name: "弘和（南朝）",
			Kana: "こうわ",
			From: time.Date(1381, time.March, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1384, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    164,
			Name: "元中（南朝）",
			Kana: "げんちゅう",
			From: time.Date(1384, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1392, time.November, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    165,
			Name: "建武（北朝）",
			Kana: "けんむ",
			From: time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1338, time.October, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    166,
			Name: "暦応（北朝）",
			Kana: "りゃくおう/れきおう",
			From: time.Date(1338, time.October, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1342, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    167,
			Name: "康永（北朝）",
			Kana: "こうえい",
			From: time.Date(1342, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1345, time.November, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    168,
			Name: "貞和（北朝）",
			Kana: "じょうわ/ていわ",
			From: time.Date(1345, time.November, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1350, time.April, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    169,
			Name: "観応（北朝）",
			Kana: "かんのう/かんおう",
			From: time.Date(1350, time.April, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1352, time.November, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    170,
			Name: "文和（北朝）",
			Kana: "ぶんな/ぶんわ",
			From: time.Date(1352, time.November, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1356, time.April, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    171,
			Name: "延文（北朝）",
			Kana: "えんぶん",
			From: time.Date(1356, time.April, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1361, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    172,
			Name: "康安（北朝）",
			Kana: "こうあん",
			From: time.Date(1361, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1362, time.October, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    173,
			Name: "貞治（北朝）",
			Kana: "じょうじ/ていじ",
			From: time.Date(1362, time.October, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1368, time.March, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    174,
			Name: "応安（北朝）",
			Kana: "おうあん",
			From: time.Date(1368, time.March, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1375, time.March, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    175,
			Name: "永和（北朝）",
			Kana: "えいわ",
			From: time.Date(1375, time.March, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1379, time.April, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    176,
			Name: "康暦（北朝）",
			Kana: "こうりゃく",
			From: time.Date(1379, time.April, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1381, time.March, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    177,
			Name: "永徳（北朝）",
			Kana: "えいとく",
			From: time.Date(1381, time.March, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1384, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    178,
			Name: "至徳（北朝）",
			Kana: "しとく",
			From: time.Date(1384, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1387, time.October, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    179,
			Name: "嘉慶（北朝）",
			Kana: "かきょう/かけい",
			From: time.Date(1387, time.October, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1389, time.March, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    180,
			Name: "康応（北朝）",
			Kana: "こうおう",
			From: time.Date(1389, time.March, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1390, time.April, 12, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    181,
			Name: "明徳",
			Kana: "めいとく",
			From: time.Date(1390, time.April, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1394, time.August, 2, 0, 0, 0, 0, jst),
		},
	}
	gengoData = []Gengo{

		Gengo{
			C:    0,
			Name: "大化",
			Kana: "たいか",
			From: time.Date(645, time.July, 17, 0, 0, 0, 0, jst),
			To:   time.Date(650, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    1,
			Name: "白雉",
			Kana: "はくち",
			From: time.Date(650, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(654, time.November, 24, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    2,
			Name: "朱鳥",
			Kana: "しゅちょう/すちょう/あかみどり",
			From: time.Date(686, time.August, 14, 0, 0, 0, 0, jst),
			To:   time.Date(686, time.October, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    3,
			Name: "大宝",
			Kana: "たいほう/だいほう",
			From: time.Date(701, time.May, 3, 0, 0, 0, 0, jst),
			To:   time.Date(704, time.June, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    4,
			Name: "慶雲",
			Kana: "けいうん/きょううん",
			From: time.Date(704, time.June, 16, 0, 0, 0, 0, jst),
			To:   time.Date(708, time.February, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    5,
			Name: "和銅",
			Kana: "わどう",
			From: time.Date(708, time.February, 7, 0, 0, 0, 0, jst),
			To:   time.Date(715, time.October, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    6,
			Name: "霊亀",
			Kana: "れいき",
			From: time.Date(715, time.October, 3, 0, 0, 0, 0, jst),
			To:   time.Date(717, time.December, 24, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    7,
			Name: "養老",
			Kana: "ようろう",
			From: time.Date(717, time.December, 24, 0, 0, 0, 0, jst),
			To:   time.Date(724, time.March, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    8,
			Name: "神亀",
			Kana: "じんき",
			From: time.Date(724, time.March, 3, 0, 0, 0, 0, jst),
			To:   time.Date(729, time.September, 2, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    9,
			Name: "天平",
			Kana: "てんぴょう",
			From: time.Date(729, time.September, 2, 0, 0, 0, 0, jst),
			To:   time.Date(749, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    10,
			Name: "天平感宝",
			Kana: "てんぴょうかんぽう",
			From: time.Date(749, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(749, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    11,
			Name: "天平勝宝",
			Kana: "てんぴょうしょうほう",
			From: time.Date(749, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(757, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    12,
			Name: "天平宝字",
			Kana: "てんぴょうほうじ",
			From: time.Date(757, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(765, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    13,
			Name: "天平神護",
			Kana: "てんぴょうじんご",
			From: time.Date(765, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(767, time.September, 13, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    14,
			Name: "神護景雲",
			Kana: "じんごけいうん",
			From: time.Date(767, time.September, 13, 0, 0, 0, 0, jst),
			To:   time.Date(770, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    15,
			Name: "宝亀",
			Kana: "ほうき",
			From: time.Date(770, time.October, 23, 0, 0, 0, 0, jst),
			To:   time.Date(781, time.January, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    16,
			Name: "天応",
			Kana: "てんおう/てんのう",
			From: time.Date(781, time.January, 30, 0, 0, 0, 0, jst),
			To:   time.Date(782, time.September, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    17,
			Name: "延暦",
			Kana: "えんりゃく",
			From: time.Date(782, time.September, 30, 0, 0, 0, 0, jst),
			To:   time.Date(806, time.June, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    18,
			Name: "大同",
			Kana: "だいどう",
			From: time.Date(806, time.June, 8, 0, 0, 0, 0, jst),
			To:   time.Date(810, time.October, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    19,
			Name: "弘仁",
			Kana: "こうにん",
			From: time.Date(810, time.October, 20, 0, 0, 0, 0, jst),
			To:   time.Date(824, time.February, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    20,
			Name: "天長",
			Kana: "てんちょう",
			From: time.Date(824, time.February, 8, 0, 0, 0, 0, jst),
			To:   time.Date(834, time.February, 14, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    21,
			Name: "承和",
			Kana: "じょうわ",
			From: time.Date(834, time.February, 14, 0, 0, 0, 0, jst),
			To:   time.Date(848, time.July, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    22,
			Name: "嘉祥",
			Kana: "かしょう/かじょう",
			From: time.Date(848, time.July, 16, 0, 0, 0, 0, jst),
			To:   time.Date(851, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    23,
			Name: "仁寿",
			Kana: "にんじゅ",
			From: time.Date(851, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(854, time.December, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    24,
			Name: "斉衡",
			Kana: "さいこう",
			From: time.Date(854, time.December, 23, 0, 0, 0, 0, jst),
			To:   time.Date(857, time.March, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    25,
			Name: "天安",
			Kana: "てんあん/てんなん",
			From: time.Date(857, time.March, 20, 0, 0, 0, 0, jst),
			To:   time.Date(859, time.May, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    26,
			Name: "貞観",
			Kana: "じょうがん",
			From: time.Date(859, time.May, 20, 0, 0, 0, 0, jst),
			To:   time.Date(877, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    27,
			Name: "元慶",
			Kana: "がんぎょう",
			From: time.Date(877, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(885, time.March, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    28,
			Name: "仁和",
			Kana: "にんな",
			From: time.Date(885, time.March, 11, 0, 0, 0, 0, jst),
			To:   time.Date(889, time.May, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    29,
			Name: "寛平",
			Kana: "かんぴょう/かんぺい/かんへい",
			From: time.Date(889, time.May, 30, 0, 0, 0, 0, jst),
			To:   time.Date(898, time.May, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    30,
			Name: "昌泰",
			Kana: "しょうたい",
			From: time.Date(898, time.May, 20, 0, 0, 0, 0, jst),
			To:   time.Date(901, time.August, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    31,
			Name: "延喜",
			Kana: "えんぎ",
			From: time.Date(901, time.August, 31, 0, 0, 0, 0, jst),
			To:   time.Date(923, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    32,
			Name: "延長",
			Kana: "えんちょう",
			From: time.Date(923, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(931, time.May, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    33,
			Name: "承平",
			Kana: "じょうへい/しょうへい",
			From: time.Date(931, time.May, 16, 0, 0, 0, 0, jst),
			To:   time.Date(938, time.June, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    34,
			Name: "天慶",
			Kana: "てんぎょう",
			From: time.Date(938, time.June, 22, 0, 0, 0, 0, jst),
			To:   time.Date(947, time.May, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    35,
			Name: "天暦",
			Kana: "てんりゃく",
			From: time.Date(947, time.May, 15, 0, 0, 0, 0, jst),
			To:   time.Date(957, time.November, 21, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    36,
			Name: "天徳",
			Kana: "てんとく",
			From: time.Date(957, time.November, 21, 0, 0, 0, 0, jst),
			To:   time.Date(961, time.March, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    37,
			Name: "応和",
			Kana: "おうわ",
			From: time.Date(961, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(964, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    38,
			Name: "康保",
			Kana: "こうほう",
			From: time.Date(964, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(968, time.September, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    39,
			Name: "安和",
			Kana: "あんな/あんわ",
			From: time.Date(968, time.September, 8, 0, 0, 0, 0, jst),
			To:   time.Date(970, time.May, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    40,
			Name: "天禄",
			Kana: "てんろく",
			From: time.Date(970, time.May, 3, 0, 0, 0, 0, jst),
			To:   time.Date(974, time.January, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    41,
			Name: "天延",
			Kana: "てんえん",
			From: time.Date(974, time.January, 16, 0, 0, 0, 0, jst),
			To:   time.Date(976, time.August, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    42,
			Name: "貞元",
			Kana: "じょうげん",
			From: time.Date(976, time.August, 11, 0, 0, 0, 0, jst),
			To:   time.Date(978, time.December, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    43,
			Name: "天元",
			Kana: "てんげん",
			From: time.Date(978, time.December, 31, 0, 0, 0, 0, jst),
			To:   time.Date(983, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    44,
			Name: "永観",
			Kana: "えいかん",
			From: time.Date(983, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(985, time.May, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    45,
			Name: "寛和",
			Kana: "かんな",
			From: time.Date(985, time.May, 19, 0, 0, 0, 0, jst),
			To:   time.Date(987, time.May, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    46,
			Name: "永延",
			Kana: "えいえん",
			From: time.Date(987, time.May, 5, 0, 0, 0, 0, jst),
			To:   time.Date(989, time.September, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    47,
			Name: "永祚",
			Kana: "えいそ",
			From: time.Date(989, time.September, 10, 0, 0, 0, 0, jst),
			To:   time.Date(990, time.November, 26, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    48,
			Name: "正暦",
			Kana: "しょうりゃく",
			From: time.Date(990, time.November, 26, 0, 0, 0, 0, jst),
			To:   time.Date(995, time.March, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    49,
			Name: "長徳",
			Kana: "ちょうとく",
			From: time.Date(995, time.March, 25, 0, 0, 0, 0, jst),
			To:   time.Date(999, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    50,
			Name: "長保",
			Kana: "ちょうほう",
			From: time.Date(999, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1004, time.August, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    51,
			Name: "寛弘",
			Kana: "かんこう",
			From: time.Date(1004, time.August, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1013, time.February, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    52,
			Name: "長和",
			Kana: "ちょうわ",
			From: time.Date(1013, time.February, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1017, time.May, 21, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    53,
			Name: "寛仁",
			Kana: "かんにん",
			From: time.Date(1017, time.May, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1021, time.March, 17, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    54,
			Name: "治安",
			Kana: "じあん",
			From: time.Date(1021, time.March, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1024, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    55,
			Name: "万寿",
			Kana: "まんじゅ",
			From: time.Date(1024, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1028, time.August, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    56,
			Name: "長元",
			Kana: "ちょうげん",
			From: time.Date(1028, time.August, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1037, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    57,
			Name: "長暦",
			Kana: "ちょうりゃく",
			From: time.Date(1037, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1040, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    58,
			Name: "長久",
			Kana: "ちょうきゅう",
			From: time.Date(1040, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1044, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    59,
			Name: "寛徳",
			Kana: "かんとく",
			From: time.Date(1044, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1046, time.May, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    60,
			Name: "永承",
			Kana: "えいしょう/えいじょう",
			From: time.Date(1046, time.May, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1053, time.February, 2, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    61,
			Name: "天喜",
			Kana: "てんき/てんぎ",
			From: time.Date(1053, time.February, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1058, time.September, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    62,
			Name: "康平",
			Kana: "こうへい",
			From: time.Date(1058, time.September, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1065, time.September, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    63,
			Name: "治暦",
			Kana: "じりゃく",
			From: time.Date(1065, time.September, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1069, time.May, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    64,
			Name: "延久",
			Kana: "えんきゅう",
			From: time.Date(1069, time.May, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1074, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    65,
			Name: "承保",
			Kana: "じょうほう/しょうほう",
			From: time.Date(1074, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1077, time.December, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    66,
			Name: "承暦",
			Kana: "じょうりゃく/しょうりゃく",
			From: time.Date(1077, time.December, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1081, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    67,
			Name: "永保",
			Kana: "えいほう",
			From: time.Date(1081, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1084, time.March, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    68,
			Name: "応徳",
			Kana: "おうとく",
			From: time.Date(1084, time.March, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1087, time.May, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    69,
			Name: "寛治",
			Kana: "かんじ",
			From: time.Date(1087, time.May, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1095, time.January, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    70,
			Name: "嘉保",
			Kana: "かほう",
			From: time.Date(1095, time.January, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1097, time.January, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    71,
			Name: "永長",
			Kana: "えいちょう",
			From: time.Date(1097, time.January, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1097, time.December, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    72,
			Name: "承徳",
			Kana: "じょうとく/しょうとく",
			From: time.Date(1097, time.December, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1099, time.September, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    73,
			Name: "康和",
			Kana: "こうわ",
			From: time.Date(1099, time.September, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1104, time.March, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    74,
			Name: "長治",
			Kana: "ちょうじ",
			From: time.Date(1104, time.March, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1106, time.May, 13, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    75,
			Name: "嘉承",
			Kana: "かしょう/かじょう",
			From: time.Date(1106, time.May, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1108, time.September, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    76,
			Name: "天仁",
			Kana: "てんにん",
			From: time.Date(1108, time.September, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1110, time.July, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    77,
			Name: "天永",
			Kana: "てんえい",
			From: time.Date(1110, time.July, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1113, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    78,
			Name: "永久",
			Kana: "えいきゅう",
			From: time.Date(1113, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1118, time.April, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    79,
			Name: "元永",
			Kana: "げんえい",
			From: time.Date(1118, time.April, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1120, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    80,
			Name: "保安",
			Kana: "ほうあん",
			From: time.Date(1120, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1124, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    81,
			Name: "天治",
			Kana: "てんじ",
			From: time.Date(1124, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1126, time.February, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    82,
			Name: "大治",
			Kana: "だいじ",
			From: time.Date(1126, time.February, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1131, time.February, 28, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    83,
			Name: "天承",
			Kana: "てんしょう/てんじょう",
			From: time.Date(1131, time.February, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1132, time.September, 21, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    84,
			Name: "長承",
			Kana: "ちょうしょう",
			From: time.Date(1132, time.September, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1135, time.June, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    85,
			Name: "保延",
			Kana: "ほうえん",
			From: time.Date(1135, time.June, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1141, time.August, 13, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    86,
			Name: "永治",
			Kana: "えいじ",
			From: time.Date(1141, time.August, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1142, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    87,
			Name: "康治",
			Kana: "こうじ",
			From: time.Date(1142, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1144, time.March, 28, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    88,
			Name: "天養",
			Kana: "てんよう",
			From: time.Date(1144, time.March, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1145, time.August, 12, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    89,
			Name: "久安",
			Kana: "きゅうあん",
			From: time.Date(1145, time.August, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1151, time.February, 14, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    90,
			Name: "仁平",
			Kana: "にんぺい/にんぴょう",
			From: time.Date(1151, time.February, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1154, time.December, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    91,
			Name: "久寿",
			Kana: "きゅうじゅ",
			From: time.Date(1154, time.December, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1156, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    92,
			Name: "保元",
			Kana: "ほうげん",
			From: time.Date(1156, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1159, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    93,
			Name: "平治",
			Kana: "へいじ",
			From: time.Date(1159, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1160, time.February, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    94,
			Name: "永暦",
			Kana: "えいりゃく",
			From: time.Date(1160, time.February, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1161, time.September, 24, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    95,
			Name: "応保",
			Kana: "おうほう/おうほ",
			From: time.Date(1161, time.September, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1163, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    96,
			Name: "長寛",
			Kana: "ちょうかん",
			From: time.Date(1163, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1165, time.July, 14, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    97,
			Name: "永万",
			Kana: "えいまん",
			From: time.Date(1165, time.July, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1166, time.September, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    98,
			Name: "仁安",
			Kana: "にんあん",
			From: time.Date(1166, time.September, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1169, time.May, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    99,
			Name: "嘉応",
			Kana: "かおう",
			From: time.Date(1169, time.May, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1171, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    100,
			Name: "承安",
			Kana: "しょうあん",
			From: time.Date(1171, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1175, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    101,
			Name: "安元",
			Kana: "あんげん",
			From: time.Date(1175, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1177, time.August, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    102,
			Name: "治承",
			Kana: "じしょう",
			From: time.Date(1177, time.August, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1181, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    103,
			Name: "養和",
			Kana: "ようわ",
			From: time.Date(1181, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1182, time.June, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    104,
			Name: "寿永",
			Kana: "じゅえい",
			From: time.Date(1182, time.June, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1184, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    105,
			Name: "元暦",
			Kana: "げんりゃく",
			From: time.Date(1184, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1185, time.September, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    106,
			Name: "文治",
			Kana: "ぶんじ",
			From: time.Date(1185, time.September, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1190, time.May, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    107,
			Name: "建久",
			Kana: "けんきゅう",
			From: time.Date(1190, time.May, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1199, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    108,
			Name: "正治",
			Kana: "しょうじ",
			From: time.Date(1199, time.May, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1201, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    109,
			Name: "建仁",
			Kana: "けんにん",
			From: time.Date(1201, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1204, time.March, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    110,
			Name: "元久",
			Kana: "げんきゅう",
			From: time.Date(1204, time.March, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1206, time.June, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    111,
			Name: "建永",
			Kana: "けんえい",
			From: time.Date(1206, time.June, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1207, time.November, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    112,
			Name: "承元",
			Kana: "じょうげん",
			From: time.Date(1207, time.November, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1211, time.April, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    113,
			Name: "建暦",
			Kana: "けんりゃく",
			From: time.Date(1211, time.April, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1214, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    114,
			Name: "建保",
			Kana: "けんぽう",
			From: time.Date(1214, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1219, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    115,
			Name: "承久",
			Kana: "じょうきゅう",
			From: time.Date(1219, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1222, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    116,
			Name: "貞応",
			Kana: "じょうおう",
			From: time.Date(1222, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1224, time.December, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    117,
			Name: "元仁",
			Kana: "げんにん",
			From: time.Date(1224, time.December, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1225, time.May, 28, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    118,
			Name: "嘉禄",
			Kana: "かろく",
			From: time.Date(1225, time.May, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1228, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    119,
			Name: "安貞",
			Kana: "あんてい",
			From: time.Date(1228, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1229, time.March, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    120,
			Name: "寛喜",
			Kana: "かんき",
			From: time.Date(1229, time.March, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1232, time.April, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    121,
			Name: "貞永",
			Kana: "じょうえい",
			From: time.Date(1232, time.April, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1233, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    122,
			Name: "天福",
			Kana: "てんぷく",
			From: time.Date(1233, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1234, time.November, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    123,
			Name: "文暦",
			Kana: "ぶんりゃく",
			From: time.Date(1234, time.November, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1235, time.November, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    124,
			Name: "嘉禎",
			Kana: "かてい",
			From: time.Date(1235, time.November, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1238, time.December, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    125,
			Name: "暦仁",
			Kana: "りゃくにん",
			From: time.Date(1238, time.December, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1239, time.March, 13, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    126,
			Name: "延応",
			Kana: "えんおう",
			From: time.Date(1239, time.March, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1240, time.August, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    127,
			Name: "仁治",
			Kana: "にんじ",
			From: time.Date(1240, time.August, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1243, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    128,
			Name: "寛元",
			Kana: "かんげん",
			From: time.Date(1243, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1247, time.April, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    129,
			Name: "宝治",
			Kana: "ほうじ",
			From: time.Date(1247, time.April, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1249, time.May, 2, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    130,
			Name: "建長",
			Kana: "けんちょう",
			From: time.Date(1249, time.May, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1256, time.October, 24, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    131,
			Name: "康元",
			Kana: "こうげん",
			From: time.Date(1256, time.October, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1257, time.March, 31, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    132,
			Name: "正嘉",
			Kana: "しょうか",
			From: time.Date(1257, time.March, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1259, time.April, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    133,
			Name: "正元",
			Kana: "しょうげん",
			From: time.Date(1259, time.April, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1260, time.May, 24, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    134,
			Name: "文応",
			Kana: "ぶんおう",
			From: time.Date(1260, time.May, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1261, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    135,
			Name: "弘長",
			Kana: "こうちょう",
			From: time.Date(1261, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1264, time.March, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    136,
			Name: "文永",
			Kana: "ぶんえい",
			From: time.Date(1264, time.March, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1275, time.May, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    137,
			Name: "建治",
			Kana: "けんじ",
			From: time.Date(1275, time.May, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1278, time.March, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    138,
			Name: "弘安",
			Kana: "こうあん",
			From: time.Date(1278, time.March, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1288, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    139,
			Name: "正応",
			Kana: "しょうおう",
			From: time.Date(1288, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1293, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    140,
			Name: "永仁",
			Kana: "えいにん",
			From: time.Date(1293, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1299, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    141,
			Name: "正安",
			Kana: "しょうあん",
			From: time.Date(1299, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1302, time.December, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    142,
			Name: "乾元",
			Kana: "けんげん",
			From: time.Date(1302, time.December, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1303, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    143,
			Name: "嘉元",
			Kana: "かげん",
			From: time.Date(1303, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1307, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    144,
			Name: "徳治",
			Kana: "とくじ",
			From: time.Date(1307, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1308, time.November, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    145,
			Name: "延慶",
			Kana: "えんきょう",
			From: time.Date(1308, time.November, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1311, time.May, 17, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    146,
			Name: "応長",
			Kana: "おうちょう",
			From: time.Date(1311, time.May, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1312, time.April, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    147,
			Name: "正和",
			Kana: "しょうわ",
			From: time.Date(1312, time.April, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1317, time.March, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    148,
			Name: "文保",
			Kana: "ぶんぽう",
			From: time.Date(1317, time.March, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1319, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    149,
			Name: "元応",
			Kana: "げんおう",
			From: time.Date(1319, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1321, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    150,
			Name: "元亨",
			Kana: "げんこう",
			From: time.Date(1321, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1324, time.December, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    151,
			Name: "正中",
			Kana: "しょうちゅう",
			From: time.Date(1324, time.December, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1326, time.May, 28, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    152,
			Name: "嘉暦",
			Kana: "かりゃく",
			From: time.Date(1326, time.May, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    182,
			Name: "応永",
			Kana: "おうえい",
			From: time.Date(1394, time.August, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1428, time.June, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    183,
			Name: "正長",
			Kana: "しょうちょう",
			From: time.Date(1428, time.June, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1429, time.October, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    184,
			Name: "永享",
			Kana: "えいきょう",
			From: time.Date(1429, time.October, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1441, time.March, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    185,
			Name: "嘉吉",
			Kana: "かきつ",
			From: time.Date(1441, time.March, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1444, time.February, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    186,
			Name: "文安",
			Kana: "ぶんあん",
			From: time.Date(1444, time.February, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1449, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    187,
			Name: "宝徳",
			Kana: "ほうとく",
			From: time.Date(1449, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1452, time.August, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    188,
			Name: "享徳",
			Kana: "きょうとく",
			From: time.Date(1452, time.August, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1455, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    189,
			Name: "康正",
			Kana: "こうしょう",
			From: time.Date(1455, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1457, time.October, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    190,
			Name: "長禄",
			Kana: "ちょうろく",
			From: time.Date(1457, time.October, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1461, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    191,
			Name: "寛正",
			Kana: "かんしょう",
			From: time.Date(1461, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1466, time.March, 14, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    192,
			Name: "文正",
			Kana: "ぶんしょう",
			From: time.Date(1466, time.March, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1467, time.April, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    193,
			Name: "応仁",
			Kana: "おうにん",
			From: time.Date(1467, time.April, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1469, time.June, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    194,
			Name: "文明",
			Kana: "ぶんめい",
			From: time.Date(1469, time.June, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1487, time.August, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    195,
			Name: "長享",
			Kana: "ちょうきょう",
			From: time.Date(1487, time.August, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1489, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    196,
			Name: "延徳",
			Kana: "えんとく",
			From: time.Date(1489, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1492, time.August, 12, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    197,
			Name: "明応",
			Kana: "めいおう",
			From: time.Date(1492, time.August, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1501, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    198,
			Name: "文亀",
			Kana: "ぶんき",
			From: time.Date(1501, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1504, time.March, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    199,
			Name: "永正",
			Kana: "えいしょう",
			From: time.Date(1504, time.March, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1521, time.September, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    200,
			Name: "大永",
			Kana: "だいえい",
			From: time.Date(1521, time.September, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1528, time.September, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    201,
			Name: "享禄",
			Kana: "きょうろく",
			From: time.Date(1528, time.September, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1532, time.August, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    202,
			Name: "天文",
			Kana: "てんぶん",
			From: time.Date(1532, time.August, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1555, time.November, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    203,
			Name: "弘治",
			Kana: "こうじ",
			From: time.Date(1555, time.November, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1558, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    204,
			Name: "永禄",
			Kana: "えいろく",
			From: time.Date(1558, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1570, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    205,
			Name: "元亀",
			Kana: "げんき",
			From: time.Date(1570, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1573, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    206,
			Name: "天正",
			Kana: "てんしょう",
			From: time.Date(1573, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1593, time.January, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    207,
			Name: "文禄",
			Kana: "ぶんろく",
			From: time.Date(1593, time.January, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1596, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    208,
			Name: "慶長",
			Kana: "けいちょう",
			From: time.Date(1596, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1615, time.September, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    209,
			Name: "元和",
			Kana: "げんな",
			From: time.Date(1615, time.September, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1624, time.April, 17, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    210,
			Name: "寛永",
			Kana: "かんえい",
			From: time.Date(1624, time.April, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1645, time.January, 13, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    211,
			Name: "正保",
			Kana: "しょうほう",
			From: time.Date(1645, time.January, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1648, time.April, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    212,
			Name: "慶安",
			Kana: "けいあん",
			From: time.Date(1648, time.April, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1652, time.October, 20, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    213,
			Name: "承応",
			Kana: "じょうおう",
			From: time.Date(1652, time.October, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1655, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    214,
			Name: "明暦",
			Kana: "めいれき",
			From: time.Date(1655, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1658, time.August, 21, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    215,
			Name: "万治",
			Kana: "まんじ",
			From: time.Date(1658, time.August, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1661, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    216,
			Name: "寛文",
			Kana: "かんぶん",
			From: time.Date(1661, time.May, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1673, time.October, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    217,
			Name: "延宝",
			Kana: "えんぽう",
			From: time.Date(1673, time.October, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1681, time.November, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    218,
			Name: "天和",
			Kana: "てんな",
			From: time.Date(1681, time.November, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1684, time.April, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    219,
			Name: "貞享",
			Kana: "じょうきょう",
			From: time.Date(1684, time.April, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1688, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    220,
			Name: "元禄",
			Kana: "げんろく",
			From: time.Date(1688, time.October, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1704, time.April, 16, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    221,
			Name: "宝永",
			Kana: "ほうえい",
			From: time.Date(1704, time.April, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1711, time.June, 11, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    222,
			Name: "正徳",
			Kana: "しょうとく",
			From: time.Date(1711, time.June, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1716, time.August, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    223,
			Name: "享保",
			Kana: "きょうほう",
			From: time.Date(1716, time.August, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1736, time.June, 7, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    224,
			Name: "元文",
			Kana: "げんぶん",
			From: time.Date(1736, time.June, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1741, time.April, 12, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    225,
			Name: "寛保",
			Kana: "かんぽう",
			From: time.Date(1741, time.April, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1744, time.April, 3, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    226,
			Name: "延享",
			Kana: "えんきょう",
			From: time.Date(1744, time.April, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1748, time.August, 5, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    227,
			Name: "寛延",
			Kana: "かんえん",
			From: time.Date(1748, time.August, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1751, time.December, 14, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    228,
			Name: "宝暦",
			Kana: "ほうれき",
			From: time.Date(1751, time.December, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1764, time.June, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    229,
			Name: "明和",
			Kana: "めいわ",
			From: time.Date(1764, time.June, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1772, time.December, 10, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    230,
			Name: "安永",
			Kana: "あんえい",
			From: time.Date(1772, time.December, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1781, time.April, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    231,
			Name: "天明",
			Kana: "てんめい",
			From: time.Date(1781, time.April, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1789, time.February, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    232,
			Name: "寛政",
			Kana: "かんせい",
			From: time.Date(1789, time.February, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1801, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    233,
			Name: "享和",
			Kana: "きょうわ",
			From: time.Date(1801, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1804, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    234,
			Name: "文化",
			Kana: "ぶんか",
			From: time.Date(1804, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1818, time.May, 26, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    235,
			Name: "文政",
			Kana: "ぶんせい",
			From: time.Date(1818, time.May, 26, 0, 0, 0, 0, jst),
			To:   time.Date(1831, time.January, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    236,
			Name: "天保",
			Kana: "てんぽう",
			From: time.Date(1831, time.January, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1845, time.January, 9, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    237,
			Name: "弘化",
			Kana: "こうか",
			From: time.Date(1845, time.January, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1848, time.April, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    238,
			Name: "嘉永",
			Kana: "かえい",
			From: time.Date(1848, time.April, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1855, time.January, 15, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    239,
			Name: "安政",
			Kana: "あんせい",
			From: time.Date(1855, time.January, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1860, time.April, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    240,
			Name: "万延",
			Kana: "まんえん",
			From: time.Date(1860, time.April, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1861, time.March, 29, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    241,
			Name: "文久",
			Kana: "ぶんきゅう",
			From: time.Date(1861, time.March, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1864, time.March, 27, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    242,
			Name: "元治",
			Kana: "げんじ",
			From: time.Date(1864, time.March, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1865, time.May, 1, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    243,
			Name: "慶応",
			Kana: "けいおう",
			From: time.Date(1865, time.May, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1868, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    244,
			Name: "明治",
			Kana: "めいじ",
			From: time.Date(1868, time.January, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1912, time.July, 30, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    245,
			Name: "大正",
			Kana: "たいしょう",
			From: time.Date(1912, time.July, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1926, time.December, 25, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    246,
			Name: "昭和",
			Kana: "しょうわ",
			From: time.Date(1926, time.December, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1989, time.January, 8, 0, 0, 0, 0, jst),
		},

		Gengo{
			C:    247,
			Name: "平成",
			Kana: "へいせい",
			From: time.Date(1989, time.January, 8, 0, 0, 0, 0, jst),
			To:   time.Date(2019, time.April, 30, 0, 0, 0, 0, jst),
		},
	}
)
