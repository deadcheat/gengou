package gengou

import "time"

var (
	jst                 = time.FixedZone("Asia/Tokyo", 9*60*60)
	nanbokueraFrom      = time.Date(1329, time.September, 22, 0, 0, 0, 0, jst)
	nanbokueraTo        = time.Date(1394, time.August, 2, 0, 0, 0, 0, jst)
	gengoDataNanbokuEra = []Gengou{

		// from
		Gengou{
			C:    153,
			Name: "元徳（大覚寺統）",
			Kana: []string{"げんとく"},
			From: time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1331, time.September, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    154,
			Name: "元徳（持明院統）",
			Kana: []string{"げんとく"},
			From: time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1332, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    155,
			Name: "元弘",
			Kana: []string{"げんこう"},
			From: time.Date(1331, time.September, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    156,
			Name: "建武（南朝）",
			Kana: []string{"けんむ"},
			From: time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1336, time.April, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    157,
			Name: "延元（南朝）",
			Kana: []string{"えんげん"},
			From: time.Date(1336, time.April, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1340, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    158,
			Name: "興国（南朝）",
			Kana: []string{"こうこく"},
			From: time.Date(1340, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1347, time.January, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    159,
			Name: "正平（南朝）",
			Kana: []string{"しょうへい"},
			From: time.Date(1347, time.January, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1370, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    160,
			Name: "建徳（南朝）",
			Kana: []string{"けんとく"},
			From: time.Date(1370, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1372, time.April, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    161,
			Name: "文中（南朝）",
			Kana: []string{"ぶんちゅう"},
			From: time.Date(1372, time.April, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1375, time.June, 26, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    162,
			Name: "天授（南朝）",
			Kana: []string{"てんじゅ"},
			From: time.Date(1375, time.June, 26, 0, 0, 0, 0, jst),
			To:   time.Date(1381, time.March, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    163,
			Name: "弘和（南朝）",
			Kana: []string{"こうわ"},
			From: time.Date(1381, time.March, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1384, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    164,
			Name: "元中（南朝）",
			Kana: []string{"げんちゅう"},
			From: time.Date(1384, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1392, time.November, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    165,
			Name: "建武（北朝）",
			Kana: []string{"けんむ"},
			From: time.Date(1334, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1338, time.October, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    166,
			Name: "暦応（北朝）",
			Kana: []string{"りゃくおう", "れきおう"},
			From: time.Date(1338, time.October, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1342, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    167,
			Name: "康永（北朝）",
			Kana: []string{"こうえい"},
			From: time.Date(1342, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1345, time.November, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    168,
			Name: "貞和（北朝）",
			Kana: []string{"じょうわ", "ていわ"},
			From: time.Date(1345, time.November, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1350, time.April, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    169,
			Name: "観応（北朝）",
			Kana: []string{"かんのう", "かんおう"},
			From: time.Date(1350, time.April, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1352, time.November, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    170,
			Name: "文和（北朝）",
			Kana: []string{"ぶんな", "ぶんわ"},
			From: time.Date(1352, time.November, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1356, time.April, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    171,
			Name: "延文（北朝）",
			Kana: []string{"えんぶん"},
			From: time.Date(1356, time.April, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1361, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    172,
			Name: "康安（北朝）",
			Kana: []string{"こうあん"},
			From: time.Date(1361, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1362, time.October, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    173,
			Name: "貞治（北朝）",
			Kana: []string{"じょうじ", "ていじ"},
			From: time.Date(1362, time.October, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1368, time.March, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    174,
			Name: "応安（北朝）",
			Kana: []string{"おうあん"},
			From: time.Date(1368, time.March, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1375, time.March, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    175,
			Name: "永和（北朝）",
			Kana: []string{"えいわ"},
			From: time.Date(1375, time.March, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1379, time.April, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    176,
			Name: "康暦（北朝）",
			Kana: []string{"こうりゃく"},
			From: time.Date(1379, time.April, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1381, time.March, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    177,
			Name: "永徳（北朝）",
			Kana: []string{"えいとく"},
			From: time.Date(1381, time.March, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1384, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    178,
			Name: "至徳（北朝）",
			Kana: []string{"しとく"},
			From: time.Date(1384, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1387, time.October, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    179,
			Name: "嘉慶（北朝）",
			Kana: []string{"かきょう", "かけい"},
			From: time.Date(1387, time.October, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1389, time.March, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    180,
			Name: "康応（北朝）",
			Kana: []string{"こうおう"},
			From: time.Date(1389, time.March, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1390, time.April, 12, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    181,
			Name: "明徳",
			Kana: []string{"めいとく"},
			From: time.Date(1390, time.April, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1394, time.August, 2, 0, 0, 0, 0, jst),
		},
	}
	gengoData = []Gengou{

		Gengou{
			C:    0,
			Name: "大化",
			Kana: []string{"たいか"},
			From: time.Date(645, time.July, 17, 0, 0, 0, 0, jst),
			To:   time.Date(650, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    1,
			Name: "白雉",
			Kana: []string{"はくち"},
			From: time.Date(650, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(654, time.November, 24, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    2,
			Name: "朱鳥",
			Kana: []string{"しゅちょう", "すちょう", "あかみどり"},
			From: time.Date(686, time.August, 14, 0, 0, 0, 0, jst),
			To:   time.Date(686, time.October, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    3,
			Name: "大宝",
			Kana: []string{"たいほう", "だいほう"},
			From: time.Date(701, time.May, 3, 0, 0, 0, 0, jst),
			To:   time.Date(704, time.June, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    4,
			Name: "慶雲",
			Kana: []string{"けいうん", "きょううん"},
			From: time.Date(704, time.June, 16, 0, 0, 0, 0, jst),
			To:   time.Date(708, time.February, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    5,
			Name: "和銅",
			Kana: []string{"わどう"},
			From: time.Date(708, time.February, 7, 0, 0, 0, 0, jst),
			To:   time.Date(715, time.October, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    6,
			Name: "霊亀",
			Kana: []string{"れいき"},
			From: time.Date(715, time.October, 3, 0, 0, 0, 0, jst),
			To:   time.Date(717, time.December, 24, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    7,
			Name: "養老",
			Kana: []string{"ようろう"},
			From: time.Date(717, time.December, 24, 0, 0, 0, 0, jst),
			To:   time.Date(724, time.March, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    8,
			Name: "神亀",
			Kana: []string{"じんき"},
			From: time.Date(724, time.March, 3, 0, 0, 0, 0, jst),
			To:   time.Date(729, time.September, 2, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    9,
			Name: "天平",
			Kana: []string{"てんぴょう"},
			From: time.Date(729, time.September, 2, 0, 0, 0, 0, jst),
			To:   time.Date(749, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    10,
			Name: "天平感宝",
			Kana: []string{"てんぴょうかんぽう"},
			From: time.Date(749, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(749, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    11,
			Name: "天平勝宝",
			Kana: []string{"てんぴょうしょうほう"},
			From: time.Date(749, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(757, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    12,
			Name: "天平宝字",
			Kana: []string{"てんぴょうほうじ"},
			From: time.Date(757, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(765, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    13,
			Name: "天平神護",
			Kana: []string{"てんぴょうじんご"},
			From: time.Date(765, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(767, time.September, 13, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    14,
			Name: "神護景雲",
			Kana: []string{"じんごけいうん"},
			From: time.Date(767, time.September, 13, 0, 0, 0, 0, jst),
			To:   time.Date(770, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    15,
			Name: "宝亀",
			Kana: []string{"ほうき"},
			From: time.Date(770, time.October, 23, 0, 0, 0, 0, jst),
			To:   time.Date(781, time.January, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    16,
			Name: "天応",
			Kana: []string{"てんおう", "てんのう"},
			From: time.Date(781, time.January, 30, 0, 0, 0, 0, jst),
			To:   time.Date(782, time.September, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    17,
			Name: "延暦",
			Kana: []string{"えんりゃく"},
			From: time.Date(782, time.September, 30, 0, 0, 0, 0, jst),
			To:   time.Date(806, time.June, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    18,
			Name: "大同",
			Kana: []string{"だいどう"},
			From: time.Date(806, time.June, 8, 0, 0, 0, 0, jst),
			To:   time.Date(810, time.October, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    19,
			Name: "弘仁",
			Kana: []string{"こうにん"},
			From: time.Date(810, time.October, 20, 0, 0, 0, 0, jst),
			To:   time.Date(824, time.February, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    20,
			Name: "天長",
			Kana: []string{"てんちょう"},
			From: time.Date(824, time.February, 8, 0, 0, 0, 0, jst),
			To:   time.Date(834, time.February, 14, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    21,
			Name: "承和",
			Kana: []string{"じょうわ"},
			From: time.Date(834, time.February, 14, 0, 0, 0, 0, jst),
			To:   time.Date(848, time.July, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    22,
			Name: "嘉祥",
			Kana: []string{"かしょう", "かじょう"},
			From: time.Date(848, time.July, 16, 0, 0, 0, 0, jst),
			To:   time.Date(851, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    23,
			Name: "仁寿",
			Kana: []string{"にんじゅ"},
			From: time.Date(851, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(854, time.December, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    24,
			Name: "斉衡",
			Kana: []string{"さいこう"},
			From: time.Date(854, time.December, 23, 0, 0, 0, 0, jst),
			To:   time.Date(857, time.March, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    25,
			Name: "天安",
			Kana: []string{"てんあん", "てんなん"},
			From: time.Date(857, time.March, 20, 0, 0, 0, 0, jst),
			To:   time.Date(859, time.May, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    26,
			Name: "貞観",
			Kana: []string{"じょうがん"},
			From: time.Date(859, time.May, 20, 0, 0, 0, 0, jst),
			To:   time.Date(877, time.June, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    27,
			Name: "元慶",
			Kana: []string{"がんぎょう"},
			From: time.Date(877, time.June, 1, 0, 0, 0, 0, jst),
			To:   time.Date(885, time.March, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    28,
			Name: "仁和",
			Kana: []string{"にんな"},
			From: time.Date(885, time.March, 11, 0, 0, 0, 0, jst),
			To:   time.Date(889, time.May, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    29,
			Name: "寛平",
			Kana: []string{"かんぴょう", "かんぺい", "かんへい"},
			From: time.Date(889, time.May, 30, 0, 0, 0, 0, jst),
			To:   time.Date(898, time.May, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    30,
			Name: "昌泰",
			Kana: []string{"しょうたい"},
			From: time.Date(898, time.May, 20, 0, 0, 0, 0, jst),
			To:   time.Date(901, time.August, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    31,
			Name: "延喜",
			Kana: []string{"えんぎ"},
			From: time.Date(901, time.August, 31, 0, 0, 0, 0, jst),
			To:   time.Date(923, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    32,
			Name: "延長",
			Kana: []string{"えんちょう"},
			From: time.Date(923, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(931, time.May, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    33,
			Name: "承平",
			Kana: []string{"じょうへい", "しょうへい"},
			From: time.Date(931, time.May, 16, 0, 0, 0, 0, jst),
			To:   time.Date(938, time.June, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    34,
			Name: "天慶",
			Kana: []string{"てんぎょう"},
			From: time.Date(938, time.June, 22, 0, 0, 0, 0, jst),
			To:   time.Date(947, time.May, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    35,
			Name: "天暦",
			Kana: []string{"てんりゃく"},
			From: time.Date(947, time.May, 15, 0, 0, 0, 0, jst),
			To:   time.Date(957, time.November, 21, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    36,
			Name: "天徳",
			Kana: []string{"てんとく"},
			From: time.Date(957, time.November, 21, 0, 0, 0, 0, jst),
			To:   time.Date(961, time.March, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    37,
			Name: "応和",
			Kana: []string{"おうわ"},
			From: time.Date(961, time.March, 5, 0, 0, 0, 0, jst),
			To:   time.Date(964, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    38,
			Name: "康保",
			Kana: []string{"こうほう"},
			From: time.Date(964, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(968, time.September, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    39,
			Name: "安和",
			Kana: []string{"あんな", "あんわ"},
			From: time.Date(968, time.September, 8, 0, 0, 0, 0, jst),
			To:   time.Date(970, time.May, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    40,
			Name: "天禄",
			Kana: []string{"てんろく"},
			From: time.Date(970, time.May, 3, 0, 0, 0, 0, jst),
			To:   time.Date(974, time.January, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    41,
			Name: "天延",
			Kana: []string{"てんえん"},
			From: time.Date(974, time.January, 16, 0, 0, 0, 0, jst),
			To:   time.Date(976, time.August, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    42,
			Name: "貞元",
			Kana: []string{"じょうげん"},
			From: time.Date(976, time.August, 11, 0, 0, 0, 0, jst),
			To:   time.Date(978, time.December, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    43,
			Name: "天元",
			Kana: []string{"てんげん"},
			From: time.Date(978, time.December, 31, 0, 0, 0, 0, jst),
			To:   time.Date(983, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    44,
			Name: "永観",
			Kana: []string{"えいかん"},
			From: time.Date(983, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(985, time.May, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    45,
			Name: "寛和",
			Kana: []string{"かんな"},
			From: time.Date(985, time.May, 19, 0, 0, 0, 0, jst),
			To:   time.Date(987, time.May, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    46,
			Name: "永延",
			Kana: []string{"えいえん"},
			From: time.Date(987, time.May, 5, 0, 0, 0, 0, jst),
			To:   time.Date(989, time.September, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    47,
			Name: "永祚",
			Kana: []string{"えいそ"},
			From: time.Date(989, time.September, 10, 0, 0, 0, 0, jst),
			To:   time.Date(990, time.November, 26, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    48,
			Name: "正暦",
			Kana: []string{"しょうりゃく"},
			From: time.Date(990, time.November, 26, 0, 0, 0, 0, jst),
			To:   time.Date(995, time.March, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    49,
			Name: "長徳",
			Kana: []string{"ちょうとく"},
			From: time.Date(995, time.March, 25, 0, 0, 0, 0, jst),
			To:   time.Date(999, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    50,
			Name: "長保",
			Kana: []string{"ちょうほう"},
			From: time.Date(999, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1004, time.August, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    51,
			Name: "寛弘",
			Kana: []string{"かんこう"},
			From: time.Date(1004, time.August, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1013, time.February, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    52,
			Name: "長和",
			Kana: []string{"ちょうわ"},
			From: time.Date(1013, time.February, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1017, time.May, 21, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    53,
			Name: "寛仁",
			Kana: []string{"かんにん"},
			From: time.Date(1017, time.May, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1021, time.March, 17, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    54,
			Name: "治安",
			Kana: []string{"じあん"},
			From: time.Date(1021, time.March, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1024, time.August, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    55,
			Name: "万寿",
			Kana: []string{"まんじゅ"},
			From: time.Date(1024, time.August, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1028, time.August, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    56,
			Name: "長元",
			Kana: []string{"ちょうげん"},
			From: time.Date(1028, time.August, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1037, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    57,
			Name: "長暦",
			Kana: []string{"ちょうりゃく"},
			From: time.Date(1037, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1040, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    58,
			Name: "長久",
			Kana: []string{"ちょうきゅう"},
			From: time.Date(1040, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1044, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    59,
			Name: "寛徳",
			Kana: []string{"かんとく"},
			From: time.Date(1044, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1046, time.May, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    60,
			Name: "永承",
			Kana: []string{"えいしょう", "えいじょう"},
			From: time.Date(1046, time.May, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1053, time.February, 2, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    61,
			Name: "天喜",
			Kana: []string{"てんき", "てんぎ"},
			From: time.Date(1053, time.February, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1058, time.September, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    62,
			Name: "康平",
			Kana: []string{"こうへい"},
			From: time.Date(1058, time.September, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1065, time.September, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    63,
			Name: "治暦",
			Kana: []string{"じりゃく"},
			From: time.Date(1065, time.September, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1069, time.May, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    64,
			Name: "延久",
			Kana: []string{"えんきゅう"},
			From: time.Date(1069, time.May, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1074, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    65,
			Name: "承保",
			Kana: []string{"じょうほう", "しょうほう"},
			From: time.Date(1074, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1077, time.December, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    66,
			Name: "承暦",
			Kana: []string{"じょうりゃく", "しょうりゃく"},
			From: time.Date(1077, time.December, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1081, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    67,
			Name: "永保",
			Kana: []string{"えいほう"},
			From: time.Date(1081, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1084, time.March, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    68,
			Name: "応徳",
			Kana: []string{"おうとく"},
			From: time.Date(1084, time.March, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1087, time.May, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    69,
			Name: "寛治",
			Kana: []string{"かんじ"},
			From: time.Date(1087, time.May, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1095, time.January, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    70,
			Name: "嘉保",
			Kana: []string{"かほう"},
			From: time.Date(1095, time.January, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1097, time.January, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    71,
			Name: "永長",
			Kana: []string{"えいちょう"},
			From: time.Date(1097, time.January, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1097, time.December, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    72,
			Name: "承徳",
			Kana: []string{"じょうとく", "しょうとく"},
			From: time.Date(1097, time.December, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1099, time.September, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    73,
			Name: "康和",
			Kana: []string{"こうわ"},
			From: time.Date(1099, time.September, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1104, time.March, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    74,
			Name: "長治",
			Kana: []string{"ちょうじ"},
			From: time.Date(1104, time.March, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1106, time.May, 13, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    75,
			Name: "嘉承",
			Kana: []string{"かしょう", "かじょう"},
			From: time.Date(1106, time.May, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1108, time.September, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    76,
			Name: "天仁",
			Kana: []string{"てんにん"},
			From: time.Date(1108, time.September, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1110, time.July, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    77,
			Name: "天永",
			Kana: []string{"てんえい"},
			From: time.Date(1110, time.July, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1113, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    78,
			Name: "永久",
			Kana: []string{"えいきゅう"},
			From: time.Date(1113, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1118, time.April, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    79,
			Name: "元永",
			Kana: []string{"げんえい"},
			From: time.Date(1118, time.April, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1120, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    80,
			Name: "保安",
			Kana: []string{"ほうあん"},
			From: time.Date(1120, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1124, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    81,
			Name: "天治",
			Kana: []string{"てんじ"},
			From: time.Date(1124, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1126, time.February, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    82,
			Name: "大治",
			Kana: []string{"だいじ"},
			From: time.Date(1126, time.February, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1131, time.February, 28, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    83,
			Name: "天承",
			Kana: []string{"てんしょう", "てんじょう"},
			From: time.Date(1131, time.February, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1132, time.September, 21, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    84,
			Name: "長承",
			Kana: []string{"ちょうしょう"},
			From: time.Date(1132, time.September, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1135, time.June, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    85,
			Name: "保延",
			Kana: []string{"ほうえん"},
			From: time.Date(1135, time.June, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1141, time.August, 13, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    86,
			Name: "永治",
			Kana: []string{"えいじ"},
			From: time.Date(1141, time.August, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1142, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    87,
			Name: "康治",
			Kana: []string{"こうじ"},
			From: time.Date(1142, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1144, time.March, 28, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    88,
			Name: "天養",
			Kana: []string{"てんよう"},
			From: time.Date(1144, time.March, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1145, time.August, 12, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    89,
			Name: "久安",
			Kana: []string{"きゅうあん"},
			From: time.Date(1145, time.August, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1151, time.February, 14, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    90,
			Name: "仁平",
			Kana: []string{"にんぺい", "にんぴょう"},
			From: time.Date(1151, time.February, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1154, time.December, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    91,
			Name: "久寿",
			Kana: []string{"きゅうじゅ"},
			From: time.Date(1154, time.December, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1156, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    92,
			Name: "保元",
			Kana: []string{"ほうげん"},
			From: time.Date(1156, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1159, time.May, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    93,
			Name: "平治",
			Kana: []string{"へいじ"},
			From: time.Date(1159, time.May, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1160, time.February, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    94,
			Name: "永暦",
			Kana: []string{"えいりゃく"},
			From: time.Date(1160, time.February, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1161, time.September, 24, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    95,
			Name: "応保",
			Kana: []string{"おうほう", "おうほ"},
			From: time.Date(1161, time.September, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1163, time.May, 4, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    96,
			Name: "長寛",
			Kana: []string{"ちょうかん"},
			From: time.Date(1163, time.May, 4, 0, 0, 0, 0, jst),
			To:   time.Date(1165, time.July, 14, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    97,
			Name: "永万",
			Kana: []string{"えいまん"},
			From: time.Date(1165, time.July, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1166, time.September, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    98,
			Name: "仁安",
			Kana: []string{"にんあん"},
			From: time.Date(1166, time.September, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1169, time.May, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    99,
			Name: "嘉応",
			Kana: []string{"かおう"},
			From: time.Date(1169, time.May, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1171, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    100,
			Name: "承安",
			Kana: []string{"しょうあん"},
			From: time.Date(1171, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1175, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    101,
			Name: "安元",
			Kana: []string{"あんげん"},
			From: time.Date(1175, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1177, time.August, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    102,
			Name: "治承",
			Kana: []string{"じしょう"},
			From: time.Date(1177, time.August, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1181, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    103,
			Name: "養和",
			Kana: []string{"ようわ"},
			From: time.Date(1181, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1182, time.June, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    104,
			Name: "寿永",
			Kana: []string{"じゅえい"},
			From: time.Date(1182, time.June, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1184, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    105,
			Name: "元暦",
			Kana: []string{"げんりゃく"},
			From: time.Date(1184, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1185, time.September, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    106,
			Name: "文治",
			Kana: []string{"ぶんじ"},
			From: time.Date(1185, time.September, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1190, time.May, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    107,
			Name: "建久",
			Kana: []string{"けんきゅう"},
			From: time.Date(1190, time.May, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1199, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    108,
			Name: "正治",
			Kana: []string{"しょうじ"},
			From: time.Date(1199, time.May, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1201, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    109,
			Name: "建仁",
			Kana: []string{"けんにん"},
			From: time.Date(1201, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1204, time.March, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    110,
			Name: "元久",
			Kana: []string{"げんきゅう"},
			From: time.Date(1204, time.March, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1206, time.June, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    111,
			Name: "建永",
			Kana: []string{"けんえい"},
			From: time.Date(1206, time.June, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1207, time.November, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    112,
			Name: "承元",
			Kana: []string{"じょうげん"},
			From: time.Date(1207, time.November, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1211, time.April, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    113,
			Name: "建暦",
			Kana: []string{"けんりゃく"},
			From: time.Date(1211, time.April, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1214, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    114,
			Name: "建保",
			Kana: []string{"けんぽう"},
			From: time.Date(1214, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1219, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    115,
			Name: "承久",
			Kana: []string{"じょうきゅう"},
			From: time.Date(1219, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1222, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    116,
			Name: "貞応",
			Kana: []string{"じょうおう"},
			From: time.Date(1222, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1224, time.December, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    117,
			Name: "元仁",
			Kana: []string{"げんにん"},
			From: time.Date(1224, time.December, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1225, time.May, 28, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    118,
			Name: "嘉禄",
			Kana: []string{"かろく"},
			From: time.Date(1225, time.May, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1228, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    119,
			Name: "安貞",
			Kana: []string{"あんてい"},
			From: time.Date(1228, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1229, time.March, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    120,
			Name: "寛喜",
			Kana: []string{"かんき"},
			From: time.Date(1229, time.March, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1232, time.April, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    121,
			Name: "貞永",
			Kana: []string{"じょうえい"},
			From: time.Date(1232, time.April, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1233, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    122,
			Name: "天福",
			Kana: []string{"てんぷく"},
			From: time.Date(1233, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1234, time.November, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    123,
			Name: "文暦",
			Kana: []string{"ぶんりゃく"},
			From: time.Date(1234, time.November, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1235, time.November, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    124,
			Name: "嘉禎",
			Kana: []string{"かてい"},
			From: time.Date(1235, time.November, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1238, time.December, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    125,
			Name: "暦仁",
			Kana: []string{"りゃくにん"},
			From: time.Date(1238, time.December, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1239, time.March, 13, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    126,
			Name: "延応",
			Kana: []string{"えんおう"},
			From: time.Date(1239, time.March, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1240, time.August, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    127,
			Name: "仁治",
			Kana: []string{"にんじ"},
			From: time.Date(1240, time.August, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1243, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    128,
			Name: "寛元",
			Kana: []string{"かんげん"},
			From: time.Date(1243, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1247, time.April, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    129,
			Name: "宝治",
			Kana: []string{"ほうじ"},
			From: time.Date(1247, time.April, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1249, time.May, 2, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    130,
			Name: "建長",
			Kana: []string{"けんちょう"},
			From: time.Date(1249, time.May, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1256, time.October, 24, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    131,
			Name: "康元",
			Kana: []string{"こうげん"},
			From: time.Date(1256, time.October, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1257, time.March, 31, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    132,
			Name: "正嘉",
			Kana: []string{"しょうか"},
			From: time.Date(1257, time.March, 31, 0, 0, 0, 0, jst),
			To:   time.Date(1259, time.April, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    133,
			Name: "正元",
			Kana: []string{"しょうげん"},
			From: time.Date(1259, time.April, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1260, time.May, 24, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    134,
			Name: "文応",
			Kana: []string{"ぶんおう"},
			From: time.Date(1260, time.May, 24, 0, 0, 0, 0, jst),
			To:   time.Date(1261, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    135,
			Name: "弘長",
			Kana: []string{"こうちょう"},
			From: time.Date(1261, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1264, time.March, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    136,
			Name: "文永",
			Kana: []string{"ぶんえい"},
			From: time.Date(1264, time.March, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1275, time.May, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    137,
			Name: "建治",
			Kana: []string{"けんじ"},
			From: time.Date(1275, time.May, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1278, time.March, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    138,
			Name: "弘安",
			Kana: []string{"こうあん"},
			From: time.Date(1278, time.March, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1288, time.May, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    139,
			Name: "正応",
			Kana: []string{"しょうおう"},
			From: time.Date(1288, time.May, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1293, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    140,
			Name: "永仁",
			Kana: []string{"えいにん"},
			From: time.Date(1293, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1299, time.May, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    141,
			Name: "正安",
			Kana: []string{"しょうあん"},
			From: time.Date(1299, time.May, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1302, time.December, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    142,
			Name: "乾元",
			Kana: []string{"けんげん"},
			From: time.Date(1302, time.December, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1303, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    143,
			Name: "嘉元",
			Kana: []string{"かげん"},
			From: time.Date(1303, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1307, time.January, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    144,
			Name: "徳治",
			Kana: []string{"とくじ"},
			From: time.Date(1307, time.January, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1308, time.November, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    145,
			Name: "延慶",
			Kana: []string{"えんきょう"},
			From: time.Date(1308, time.November, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1311, time.May, 17, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    146,
			Name: "応長",
			Kana: []string{"おうちょう"},
			From: time.Date(1311, time.May, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1312, time.April, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    147,
			Name: "正和",
			Kana: []string{"しょうわ"},
			From: time.Date(1312, time.April, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1317, time.March, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    148,
			Name: "文保",
			Kana: []string{"ぶんぽう"},
			From: time.Date(1317, time.March, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1319, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    149,
			Name: "元応",
			Kana: []string{"げんおう"},
			From: time.Date(1319, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1321, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    150,
			Name: "元亨",
			Kana: []string{"げんこう"},
			From: time.Date(1321, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1324, time.December, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    151,
			Name: "正中",
			Kana: []string{"しょうちゅう"},
			From: time.Date(1324, time.December, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1326, time.May, 28, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    152,
			Name: "嘉暦",
			Kana: []string{"かりゃく"},
			From: time.Date(1326, time.May, 28, 0, 0, 0, 0, jst),
			To:   time.Date(1329, time.September, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    182,
			Name: "応永",
			Kana: []string{"おうえい"},
			From: time.Date(1394, time.August, 2, 0, 0, 0, 0, jst),
			To:   time.Date(1428, time.June, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    183,
			Name: "正長",
			Kana: []string{"しょうちょう"},
			From: time.Date(1428, time.June, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1429, time.October, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    184,
			Name: "永享",
			Kana: []string{"えいきょう"},
			From: time.Date(1429, time.October, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1441, time.March, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    185,
			Name: "嘉吉",
			Kana: []string{"かきつ"},
			From: time.Date(1441, time.March, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1444, time.February, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    186,
			Name: "文安",
			Kana: []string{"ぶんあん"},
			From: time.Date(1444, time.February, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1449, time.August, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    187,
			Name: "宝徳",
			Kana: []string{"ほうとく"},
			From: time.Date(1449, time.August, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1452, time.August, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    188,
			Name: "享徳",
			Kana: []string{"きょうとく"},
			From: time.Date(1452, time.August, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1455, time.September, 6, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    189,
			Name: "康正",
			Kana: []string{"こうしょう"},
			From: time.Date(1455, time.September, 6, 0, 0, 0, 0, jst),
			To:   time.Date(1457, time.October, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    190,
			Name: "長禄",
			Kana: []string{"ちょうろく"},
			From: time.Date(1457, time.October, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1461, time.February, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    191,
			Name: "寛正",
			Kana: []string{"かんしょう"},
			From: time.Date(1461, time.February, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1466, time.March, 14, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    192,
			Name: "文正",
			Kana: []string{"ぶんしょう"},
			From: time.Date(1466, time.March, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1467, time.April, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    193,
			Name: "応仁",
			Kana: []string{"おうにん"},
			From: time.Date(1467, time.April, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1469, time.June, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    194,
			Name: "文明",
			Kana: []string{"ぶんめい"},
			From: time.Date(1469, time.June, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1487, time.August, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    195,
			Name: "長享",
			Kana: []string{"ちょうきょう"},
			From: time.Date(1487, time.August, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1489, time.September, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    196,
			Name: "延徳",
			Kana: []string{"えんとく"},
			From: time.Date(1489, time.September, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1492, time.August, 12, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    197,
			Name: "明応",
			Kana: []string{"めいおう"},
			From: time.Date(1492, time.August, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1501, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    198,
			Name: "文亀",
			Kana: []string{"ぶんき"},
			From: time.Date(1501, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1504, time.March, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    199,
			Name: "永正",
			Kana: []string{"えいしょう"},
			From: time.Date(1504, time.March, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1521, time.September, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    200,
			Name: "大永",
			Kana: []string{"だいえい"},
			From: time.Date(1521, time.September, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1528, time.September, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    201,
			Name: "享禄",
			Kana: []string{"きょうろく"},
			From: time.Date(1528, time.September, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1532, time.August, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    202,
			Name: "天文",
			Kana: []string{"てんぶん"},
			From: time.Date(1532, time.August, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1555, time.November, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    203,
			Name: "弘治",
			Kana: []string{"こうじ"},
			From: time.Date(1555, time.November, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1558, time.March, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    204,
			Name: "永禄",
			Kana: []string{"えいろく"},
			From: time.Date(1558, time.March, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1570, time.May, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    205,
			Name: "元亀",
			Kana: []string{"げんき"},
			From: time.Date(1570, time.May, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1573, time.August, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    206,
			Name: "天正",
			Kana: []string{"てんしょう"},
			From: time.Date(1573, time.August, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1593, time.January, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    207,
			Name: "文禄",
			Kana: []string{"ぶんろく"},
			From: time.Date(1593, time.January, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1596, time.December, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    208,
			Name: "慶長",
			Kana: []string{"けいちょう"},
			From: time.Date(1596, time.December, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1615, time.September, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    209,
			Name: "元和",
			Kana: []string{"げんな"},
			From: time.Date(1615, time.September, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1624, time.April, 17, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    210,
			Name: "寛永",
			Kana: []string{"かんえい"},
			From: time.Date(1624, time.April, 17, 0, 0, 0, 0, jst),
			To:   time.Date(1645, time.January, 13, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    211,
			Name: "正保",
			Kana: []string{"しょうほう"},
			From: time.Date(1645, time.January, 13, 0, 0, 0, 0, jst),
			To:   time.Date(1648, time.April, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    212,
			Name: "慶安",
			Kana: []string{"けいあん"},
			From: time.Date(1648, time.April, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1652, time.October, 20, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    213,
			Name: "承応",
			Kana: []string{"じょうおう"},
			From: time.Date(1652, time.October, 20, 0, 0, 0, 0, jst),
			To:   time.Date(1655, time.May, 18, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    214,
			Name: "明暦",
			Kana: []string{"めいれき"},
			From: time.Date(1655, time.May, 18, 0, 0, 0, 0, jst),
			To:   time.Date(1658, time.August, 21, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    215,
			Name: "万治",
			Kana: []string{"まんじ"},
			From: time.Date(1658, time.August, 21, 0, 0, 0, 0, jst),
			To:   time.Date(1661, time.May, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    216,
			Name: "寛文",
			Kana: []string{"かんぶん"},
			From: time.Date(1661, time.May, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1673, time.October, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    217,
			Name: "延宝",
			Kana: []string{"えんぽう"},
			From: time.Date(1673, time.October, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1681, time.November, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    218,
			Name: "天和",
			Kana: []string{"てんな"},
			From: time.Date(1681, time.November, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1684, time.April, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    219,
			Name: "貞享",
			Kana: []string{"じょうきょう"},
			From: time.Date(1684, time.April, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1688, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    220,
			Name: "元禄",
			Kana: []string{"げんろく"},
			From: time.Date(1688, time.October, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1704, time.April, 16, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    221,
			Name: "宝永",
			Kana: []string{"ほうえい"},
			From: time.Date(1704, time.April, 16, 0, 0, 0, 0, jst),
			To:   time.Date(1711, time.June, 11, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    222,
			Name: "正徳",
			Kana: []string{"しょうとく"},
			From: time.Date(1711, time.June, 11, 0, 0, 0, 0, jst),
			To:   time.Date(1716, time.August, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    223,
			Name: "享保",
			Kana: []string{"きょうほう"},
			From: time.Date(1716, time.August, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1736, time.June, 7, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    224,
			Name: "元文",
			Kana: []string{"げんぶん"},
			From: time.Date(1736, time.June, 7, 0, 0, 0, 0, jst),
			To:   time.Date(1741, time.April, 12, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    225,
			Name: "寛保",
			Kana: []string{"かんぽう"},
			From: time.Date(1741, time.April, 12, 0, 0, 0, 0, jst),
			To:   time.Date(1744, time.April, 3, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    226,
			Name: "延享",
			Kana: []string{"えんきょう"},
			From: time.Date(1744, time.April, 3, 0, 0, 0, 0, jst),
			To:   time.Date(1748, time.August, 5, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    227,
			Name: "寛延",
			Kana: []string{"かんえん"},
			From: time.Date(1748, time.August, 5, 0, 0, 0, 0, jst),
			To:   time.Date(1751, time.December, 14, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    228,
			Name: "宝暦",
			Kana: []string{"ほうれき"},
			From: time.Date(1751, time.December, 14, 0, 0, 0, 0, jst),
			To:   time.Date(1764, time.June, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    229,
			Name: "明和",
			Kana: []string{"めいわ"},
			From: time.Date(1764, time.June, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1772, time.December, 10, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    230,
			Name: "安永",
			Kana: []string{"あんえい"},
			From: time.Date(1772, time.December, 10, 0, 0, 0, 0, jst),
			To:   time.Date(1781, time.April, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    231,
			Name: "天明",
			Kana: []string{"てんめい"},
			From: time.Date(1781, time.April, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1789, time.February, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    232,
			Name: "寛政",
			Kana: []string{"かんせい"},
			From: time.Date(1789, time.February, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1801, time.March, 19, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    233,
			Name: "享和",
			Kana: []string{"きょうわ"},
			From: time.Date(1801, time.March, 19, 0, 0, 0, 0, jst),
			To:   time.Date(1804, time.March, 22, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    234,
			Name: "文化",
			Kana: []string{"ぶんか"},
			From: time.Date(1804, time.March, 22, 0, 0, 0, 0, jst),
			To:   time.Date(1818, time.May, 26, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    235,
			Name: "文政",
			Kana: []string{"ぶんせい"},
			From: time.Date(1818, time.May, 26, 0, 0, 0, 0, jst),
			To:   time.Date(1831, time.January, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    236,
			Name: "天保",
			Kana: []string{"てんぽう"},
			From: time.Date(1831, time.January, 23, 0, 0, 0, 0, jst),
			To:   time.Date(1845, time.January, 9, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    237,
			Name: "弘化",
			Kana: []string{"こうか"},
			From: time.Date(1845, time.January, 9, 0, 0, 0, 0, jst),
			To:   time.Date(1848, time.April, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    238,
			Name: "嘉永",
			Kana: []string{"かえい"},
			From: time.Date(1848, time.April, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1855, time.January, 15, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    239,
			Name: "安政",
			Kana: []string{"あんせい"},
			From: time.Date(1855, time.January, 15, 0, 0, 0, 0, jst),
			To:   time.Date(1860, time.April, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    240,
			Name: "万延",
			Kana: []string{"まんえん"},
			From: time.Date(1860, time.April, 8, 0, 0, 0, 0, jst),
			To:   time.Date(1861, time.March, 29, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    241,
			Name: "文久",
			Kana: []string{"ぶんきゅう"},
			From: time.Date(1861, time.March, 29, 0, 0, 0, 0, jst),
			To:   time.Date(1864, time.March, 27, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    242,
			Name: "元治",
			Kana: []string{"げんじ"},
			From: time.Date(1864, time.March, 27, 0, 0, 0, 0, jst),
			To:   time.Date(1865, time.May, 1, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    243,
			Name: "慶応",
			Kana: []string{"けいおう"},
			From: time.Date(1865, time.May, 1, 0, 0, 0, 0, jst),
			To:   time.Date(1868, time.October, 23, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    244,
			Name: "明治",
			Kana: []string{"めいじ"},
			From: time.Date(1868, time.January, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1912, time.July, 30, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    245,
			Name: "大正",
			Kana: []string{"たいしょう"},
			From: time.Date(1912, time.July, 30, 0, 0, 0, 0, jst),
			To:   time.Date(1926, time.December, 25, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    246,
			Name: "昭和",
			Kana: []string{"しょうわ"},
			From: time.Date(1926, time.December, 25, 0, 0, 0, 0, jst),
			To:   time.Date(1989, time.January, 8, 0, 0, 0, 0, jst),
		},

		Gengou{
			C:    247,
			Name: "平成",
			Kana: []string{"へいせい"},
			From: time.Date(1989, time.January, 8, 0, 0, 0, 0, jst),
			To:   time.Date(2019, time.April, 30, 0, 0, 0, 0, jst),
		},
	}
)
