package gengou

import (
	"bytes"
	"encoding/json"
)

// GengoCode is a type alias for using Codes of 元号
type GengoCode int

// GengoCodeFromString get code from string
func GengoCodeFromString(s string) GengoCode {
	return idMap[s]
}

// GengoCodeToString make code to string
func GengoCodeToString(c GengoCode) string {
	return strMap[c]
}

// MarshalJSON marshals the enum as a quoted json string
func (c GengoCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(GengoCodeToString(c))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (c *GengoCode) UnmarshalJSON(b []byte) error {
	var j string
	_ = json.Unmarshal(b, &j)
	*c = GengoCodeFromString(j)
	return nil
}

const (
	Taika GengoCode = iota
	Hakuchi
	ShuchouSuchouAkamidori
	TaihouDaihou
	KeiunKyouun
	Wadou
	Reiki
	Yourou
	Jinki
	Tenpyou
	Tenpyoukanpou
	Tenpyoushouhou
	Tenpyouhouji
	Tenpyoujingo
	Jingokeiun
	Houki
	TennouTennnou
	Enryaku
	Daidou
	Kounin
	Tenchou
	Jouwa
	KashouKajou848
	Ninju
	Saikou
	TennanTennnan
	Jougan859
	Gangyou
	Ninnna
	KanpyouKanpeiKanhei
	Shoutai
	Engi
	Enchou
	JouheiShouhei
	Tengyou
	Tenryaku
	Tentoku
	Ouwa
	Kouhou
	AnnnaAnwa
	Tenroku
	Tennen
	Jougen976
	Tengen
	Eikan
	Kannna
	Eien
	Eiso
	Shouryaku
	Choutoku
	Chouhou
	Kankou
	Chouwa
	Kannnin
	Jian
	Manju
	Chougen
	Chouryaku
	Choukyuu
	Kantoku
	EishouEijou
	TenkiTengi
	Kouhei
	Jiryaku
	Enkyuu
	JouhouShouhou
	JouryakuShouryaku
	Eihou
	Outoku
	Kanji
	Kahou
	Eichou
	JoutokuShoutoku
	Kouwa1099
	Chouji
	KashouKajou1106
	Tennnin
	Tennei
	Eikyuu
	Gennei
	Houan
	Tenji
	Daiji
	TenshouTenjou
	Choushou
	Houen
	Eiji
	Kouji1142
	Tennyou
	Kyuuan
	NinpeiNinpyou
	Kyuuju
	Hougen
	Heiji
	Eiryaku
	OuhouOuho
	Choukan
	Eiman
	Ninnan
	Kaou
	Shouan1171
	Angen
	Jishou
	Youwa
	Juei
	Genryaku
	Bunji
	Kenkyuu
	Shouji
	Kennnin
	Genkyuu
	Kennei
	Jougen1207
	Kenryaku
	Kenpou
	Joukyuu
	Jouou1222
	Gennnin
	Karoku
	Antei
	Kanki
	Jouei
	Tenpuku
	Bunryaku
	Katei
	Ryakunin
	Ennou
	Ninji
	Kangen
	Houji
	Kenchou
	Kougen
	Shouka
	Shougen
	Bunnou
	Kouchou
	Bunnei
	Kenji
	Kouan1278
	Shouou
	Einin
	Shouan1299
	Kengen
	Kagen
	Tokuji
	Enkyou1308
	Ouchou
	Shouwa1312
	Bunpou
	Gennou
	Genkou1321
	Shouchuu
	Karyaku
	GentokuD
	GentokuJ
	Genkou1331
	KenmuN
	Engen
	Koukoku
	Shouhei
	Kentoku
	Bunchuu
	Tenju
	Kouwa1381
	Genchuu
	KenmuH
	RyakuouRekiou
	Kouei
	JouwaTeiwa
	KannnouKannou
	BunnnaBunwa
	Enbun
	Kouan1361
	JoujiTeiji
	Ouan
	Eiwa
	Kouryaku
	Eitoku
	Shitoku
	KakyouKakei
	Kouou
	Meitoku
	Ouei
	Shouchou
	Eikyou
	Kakitsu
	Bunnan
	Houtoku
	Kyoutoku
	Koushou
	Chouroku
	Kanshou
	Bunshou
	Ounin
	Bunmei
	Choukyou
	Entoku
	Meiou
	Bunki
	Eishou
	Daiei
	Kyouroku
	Tenbun
	Kouji1555
	Eiroku
	Genki
	Tenshou
	Bunroku
	Keichou
	Gennna
	Kannei
	Shouhou
	Keian
	Jouou1652
	Meireki
	Manji
	Kanbun
	Enpou
	Tennna
	Joukyou
	Genroku
	Houei
	Shoutoku
	Kyouhou
	Genbun
	Kanpou
	Enkyou1744
	Kannen
	Houreki
	Meiwa
	Annei
	Tenmei
	Kansei
	Kyouwa
	Bunka
	Bunsei
	Tenpou
	Kouka
	Kaei
	Ansei
	Mannen
	Bunkyuu
	Genji
	Keiou
	Meiji
	Taishou
	Shouwa1926
	Heisei
)

var idMap = map[string]GengoCode{
	"Taika":                  Taika,
	"Hakuchi":                Hakuchi,
	"ShuchouSuchouAkamidori": ShuchouSuchouAkamidori,
	"TaihouDaihou":           TaihouDaihou,
	"KeiunKyouun":            KeiunKyouun,
	"Wadou":                  Wadou,
	"Reiki":                  Reiki,
	"Yourou":                 Yourou,
	"Jinki":                  Jinki,
	"Tenpyou":                Tenpyou,
	"Tenpyoukanpou":          Tenpyoukanpou,
	"Tenpyoushouhou":         Tenpyoushouhou,
	"Tenpyouhouji":           Tenpyouhouji,
	"Tenpyoujingo":           Tenpyoujingo,
	"Jingokeiun":             Jingokeiun,
	"Houki":                  Houki,
	"TennouTennnou":          TennouTennnou,
	"Enryaku":                Enryaku,
	"Daidou":                 Daidou,
	"Kounin":                 Kounin,
	"Tenchou":                Tenchou,
	"Jouwa":                  Jouwa,
	"KashouKajou848":         KashouKajou848,
	"Ninju":                  Ninju,
	"Saikou":                 Saikou,
	"TennanTennnan":          TennanTennnan,
	"Jougan859":              Jougan859,
	"Gangyou":                Gangyou,
	"Ninnna":                 Ninnna,
	"KanpyouKanpeiKanhei":    KanpyouKanpeiKanhei,
	"Shoutai":                Shoutai,
	"Engi":                   Engi,
	"Enchou":                 Enchou,
	"JouheiShouhei":          JouheiShouhei,
	"Tengyou":                Tengyou,
	"Tenryaku":               Tenryaku,
	"Tentoku":                Tentoku,
	"Ouwa":                   Ouwa,
	"Kouhou":                 Kouhou,
	"AnnnaAnwa":              AnnnaAnwa,
	"Tenroku":                Tenroku,
	"Tennen":                 Tennen,
	"Jougen976":              Jougen976,
	"Tengen":                 Tengen,
	"Eikan":                  Eikan,
	"Kannna":                 Kannna,
	"Eien":                   Eien,
	"Eiso":                   Eiso,
	"Shouryaku":              Shouryaku,
	"Choutoku":               Choutoku,
	"Chouhou":                Chouhou,
	"Kankou":                 Kankou,
	"Chouwa":                 Chouwa,
	"Kannnin":                Kannnin,
	"Jian":                   Jian,
	"Manju":                  Manju,
	"Chougen":                Chougen,
	"Chouryaku":              Chouryaku,
	"Choukyuu":               Choukyuu,
	"Kantoku":                Kantoku,
	"EishouEijou":            EishouEijou,
	"TenkiTengi":             TenkiTengi,
	"Kouhei":                 Kouhei,
	"Jiryaku":                Jiryaku,
	"Enkyuu":                 Enkyuu,
	"JouhouShouhou":          JouhouShouhou,
	"JouryakuShouryaku":      JouryakuShouryaku,
	"Eihou":                  Eihou,
	"Outoku":                 Outoku,
	"Kanji":                  Kanji,
	"Kahou":                  Kahou,
	"Eichou":                 Eichou,
	"JoutokuShoutoku":        JoutokuShoutoku,
	"Kouwa1099":              Kouwa1099,
	"Chouji":                 Chouji,
	"KashouKajou1106":        KashouKajou1106,
	"Tennnin":                Tennnin,
	"Tennei":                 Tennei,
	"Eikyuu":                 Eikyuu,
	"Gennei":                 Gennei,
	"Houan":                  Houan,
	"Tenji":                  Tenji,
	"Daiji":                  Daiji,
	"TenshouTenjou":          TenshouTenjou,
	"Choushou":               Choushou,
	"Houen":                  Houen,
	"Eiji":                   Eiji,
	"Kouji1142":              Kouji1142,
	"Tennyou":                Tennyou,
	"Kyuuan":                 Kyuuan,
	"NinpeiNinpyou":          NinpeiNinpyou,
	"Kyuuju":                 Kyuuju,
	"Hougen":                 Hougen,
	"Heiji":                  Heiji,
	"Eiryaku":                Eiryaku,
	"OuhouOuho":              OuhouOuho,
	"Choukan":                Choukan,
	"Eiman":                  Eiman,
	"Ninnan":                 Ninnan,
	"Kaou":                   Kaou,
	"Shouan1171":             Shouan1171,
	"Angen":                  Angen,
	"Jishou":                 Jishou,
	"Youwa":                  Youwa,
	"Juei":                   Juei,
	"Genryaku":               Genryaku,
	"Bunji":                  Bunji,
	"Kenkyuu":                Kenkyuu,
	"Shouji":                 Shouji,
	"Kennnin":                Kennnin,
	"Genkyuu":                Genkyuu,
	"Kennei":                 Kennei,
	"Jougen1207":             Jougen1207,
	"Kenryaku":               Kenryaku,
	"Kenpou":                 Kenpou,
	"Joukyuu":                Joukyuu,
	"Jouou1222":              Jouou1222,
	"Gennnin":                Gennnin,
	"Karoku":                 Karoku,
	"Antei":                  Antei,
	"Kanki":                  Kanki,
	"Jouei":                  Jouei,
	"Tenpuku":                Tenpuku,
	"Bunryaku":               Bunryaku,
	"Katei":                  Katei,
	"Ryakunin":               Ryakunin,
	"Ennou":                  Ennou,
	"Ninji":                  Ninji,
	"Kangen":                 Kangen,
	"Houji":                  Houji,
	"Kenchou":                Kenchou,
	"Kougen":                 Kougen,
	"Shouka":                 Shouka,
	"Shougen":                Shougen,
	"Bunnou":                 Bunnou,
	"Kouchou":                Kouchou,
	"Bunnei":                 Bunnei,
	"Kenji":                  Kenji,
	"Kouan1278":              Kouan1278,
	"Shouou":                 Shouou,
	"Einin":                  Einin,
	"Shouan1299":             Shouan1299,
	"Kengen":                 Kengen,
	"Kagen":                  Kagen,
	"Tokuji":                 Tokuji,
	"Enkyou1308":             Enkyou1308,
	"Ouchou":                 Ouchou,
	"Shouwa1312":             Shouwa1312,
	"Bunpou":                 Bunpou,
	"Gennou":                 Gennou,
	"Genkou1321":             Genkou1321,
	"Shouchuu":               Shouchuu,
	"Karyaku":                Karyaku,
	"GentokuD":               GentokuD,
	"GentokuJ":               GentokuJ,
	"Genkou1331":             Genkou1331,
	"KenmuN":                 KenmuN,
	"Engen":                  Engen,
	"Koukoku":                Koukoku,
	"Shouhei":                Shouhei,
	"Kentoku":                Kentoku,
	"Bunchuu":                Bunchuu,
	"Tenju":                  Tenju,
	"Kouwa1381":              Kouwa1381,
	"Genchuu":                Genchuu,
	"KenmuH":                 KenmuH,
	"RyakuouRekiou":          RyakuouRekiou,
	"Kouei":                  Kouei,
	"JouwaTeiwa":             JouwaTeiwa,
	"KannnouKannou":          KannnouKannou,
	"BunnnaBunwa":            BunnnaBunwa,
	"Enbun":                  Enbun,
	"Kouan1361":              Kouan1361,
	"JoujiTeiji":             JoujiTeiji,
	"Ouan":                   Ouan,
	"Eiwa":                   Eiwa,
	"Kouryaku":               Kouryaku,
	"Eitoku":                 Eitoku,
	"Shitoku":                Shitoku,
	"KakyouKakei":            KakyouKakei,
	"Kouou":                  Kouou,
	"Meitoku":                Meitoku,
	"Ouei":                   Ouei,
	"Shouchou":               Shouchou,
	"Eikyou":                 Eikyou,
	"Kakitsu":                Kakitsu,
	"Bunnan":                 Bunnan,
	"Houtoku":                Houtoku,
	"Kyoutoku":               Kyoutoku,
	"Koushou":                Koushou,
	"Chouroku":               Chouroku,
	"Kanshou":                Kanshou,
	"Bunshou":                Bunshou,
	"Ounin":                  Ounin,
	"Bunmei":                 Bunmei,
	"Choukyou":               Choukyou,
	"Entoku":                 Entoku,
	"Meiou":                  Meiou,
	"Bunki":                  Bunki,
	"Eishou":                 Eishou,
	"Daiei":                  Daiei,
	"Kyouroku":               Kyouroku,
	"Tenbun":                 Tenbun,
	"Kouji1555":              Kouji1555,
	"Eiroku":                 Eiroku,
	"Genki":                  Genki,
	"Tenshou":                Tenshou,
	"Bunroku":                Bunroku,
	"Keichou":                Keichou,
	"Gennna":                 Gennna,
	"Kannei":                 Kannei,
	"Shouhou":                Shouhou,
	"Keian":                  Keian,
	"Jouou1652":              Jouou1652,
	"Meireki":                Meireki,
	"Manji":                  Manji,
	"Kanbun":                 Kanbun,
	"Enpou":                  Enpou,
	"Tennna":                 Tennna,
	"Joukyou":                Joukyou,
	"Genroku":                Genroku,
	"Houei":                  Houei,
	"Shoutoku":               Shoutoku,
	"Kyouhou":                Kyouhou,
	"Genbun":                 Genbun,
	"Kanpou":                 Kanpou,
	"Enkyou1744":             Enkyou1744,
	"Kannen":                 Kannen,
	"Houreki":                Houreki,
	"Meiwa":                  Meiwa,
	"Annei":                  Annei,
	"Tenmei":                 Tenmei,
	"Kansei":                 Kansei,
	"Kyouwa":                 Kyouwa,
	"Bunka":                  Bunka,
	"Bunsei":                 Bunsei,
	"Tenpou":                 Tenpou,
	"Kouka":                  Kouka,
	"Kaei":                   Kaei,
	"Ansei":                  Ansei,
	"Mannen":                 Mannen,
	"Bunkyuu":                Bunkyuu,
	"Genji":                  Genji,
	"Keiou":                  Keiou,
	"Meiji":                  Meiji,
	"Taishou":                Taishou,
	"Shouwa1926":             Shouwa1926,
	"Heisei":                 Heisei,
}

var strMap = map[GengoCode]string{
	Taika:                  "Taika",
	Hakuchi:                "Hakuchi",
	ShuchouSuchouAkamidori: "ShuchouSuchouAkamidori",
	TaihouDaihou:           "TaihouDaihou",
	KeiunKyouun:            "KeiunKyouun",
	Wadou:                  "Wadou",
	Reiki:                  "Reiki",
	Yourou:                 "Yourou",
	Jinki:                  "Jinki",
	Tenpyou:                "Tenpyou",
	Tenpyoukanpou:          "Tenpyoukanpou",
	Tenpyoushouhou:         "Tenpyoushouhou",
	Tenpyouhouji:           "Tenpyouhouji",
	Tenpyoujingo:           "Tenpyoujingo",
	Jingokeiun:             "Jingokeiun",
	Houki:                  "Houki",
	TennouTennnou:          "TennouTennnou",
	Enryaku:                "Enryaku",
	Daidou:                 "Daidou",
	Kounin:                 "Kounin",
	Tenchou:                "Tenchou",
	Jouwa:                  "Jouwa",
	KashouKajou848:         "KashouKajou848",
	Ninju:                  "Ninju",
	Saikou:                 "Saikou",
	TennanTennnan:          "TennanTennnan",
	Jougan859:              "Jougan859",
	Gangyou:                "Gangyou",
	Ninnna:                 "Ninnna",
	KanpyouKanpeiKanhei:    "KanpyouKanpeiKanhei",
	Shoutai:                "Shoutai",
	Engi:                   "Engi",
	Enchou:                 "Enchou",
	JouheiShouhei:          "JouheiShouhei",
	Tengyou:                "Tengyou",
	Tenryaku:               "Tenryaku",
	Tentoku:                "Tentoku",
	Ouwa:                   "Ouwa",
	Kouhou:                 "Kouhou",
	AnnnaAnwa:              "AnnnaAnwa",
	Tenroku:                "Tenroku",
	Tennen:                 "Tennen",
	Jougen976:              "Jougen976",
	Tengen:                 "Tengen",
	Eikan:                  "Eikan",
	Kannna:                 "Kannna",
	Eien:                   "Eien",
	Eiso:                   "Eiso",
	Shouryaku:              "Shouryaku",
	Choutoku:               "Choutoku",
	Chouhou:                "Chouhou",
	Kankou:                 "Kankou",
	Chouwa:                 "Chouwa",
	Kannnin:                "Kannnin",
	Jian:                   "Jian",
	Manju:                  "Manju",
	Chougen:                "Chougen",
	Chouryaku:              "Chouryaku",
	Choukyuu:               "Choukyuu",
	Kantoku:                "Kantoku",
	EishouEijou:            "EishouEijou",
	TenkiTengi:             "TenkiTengi",
	Kouhei:                 "Kouhei",
	Jiryaku:                "Jiryaku",
	Enkyuu:                 "Enkyuu",
	JouhouShouhou:          "JouhouShouhou",
	JouryakuShouryaku:      "JouryakuShouryaku",
	Eihou:                  "Eihou",
	Outoku:                 "Outoku",
	Kanji:                  "Kanji",
	Kahou:                  "Kahou",
	Eichou:                 "Eichou",
	JoutokuShoutoku:        "JoutokuShoutoku",
	Kouwa1099:              "Kouwa1099",
	Chouji:                 "Chouji",
	KashouKajou1106:        "KashouKajou1106",
	Tennnin:                "Tennnin",
	Tennei:                 "Tennei",
	Eikyuu:                 "Eikyuu",
	Gennei:                 "Gennei",
	Houan:                  "Houan",
	Tenji:                  "Tenji",
	Daiji:                  "Daiji",
	TenshouTenjou:          "TenshouTenjou",
	Choushou:               "Choushou",
	Houen:                  "Houen",
	Eiji:                   "Eiji",
	Kouji1142:              "Kouji1142",
	Tennyou:                "Tennyou",
	Kyuuan:                 "Kyuuan",
	NinpeiNinpyou:          "NinpeiNinpyou",
	Kyuuju:                 "Kyuuju",
	Hougen:                 "Hougen",
	Heiji:                  "Heiji",
	Eiryaku:                "Eiryaku",
	OuhouOuho:              "OuhouOuho",
	Choukan:                "Choukan",
	Eiman:                  "Eiman",
	Ninnan:                 "Ninnan",
	Kaou:                   "Kaou",
	Shouan1171:             "Shouan1171",
	Angen:                  "Angen",
	Jishou:                 "Jishou",
	Youwa:                  "Youwa",
	Juei:                   "Juei",
	Genryaku:               "Genryaku",
	Bunji:                  "Bunji",
	Kenkyuu:                "Kenkyuu",
	Shouji:                 "Shouji",
	Kennnin:                "Kennnin",
	Genkyuu:                "Genkyuu",
	Kennei:                 "Kennei",
	Jougen1207:             "Jougen1207",
	Kenryaku:               "Kenryaku",
	Kenpou:                 "Kenpou",
	Joukyuu:                "Joukyuu",
	Jouou1222:              "Jouou1222",
	Gennnin:                "Gennnin",
	Karoku:                 "Karoku",
	Antei:                  "Antei",
	Kanki:                  "Kanki",
	Jouei:                  "Jouei",
	Tenpuku:                "Tenpuku",
	Bunryaku:               "Bunryaku",
	Katei:                  "Katei",
	Ryakunin:               "Ryakunin",
	Ennou:                  "Ennou",
	Ninji:                  "Ninji",
	Kangen:                 "Kangen",
	Houji:                  "Houji",
	Kenchou:                "Kenchou",
	Kougen:                 "Kougen",
	Shouka:                 "Shouka",
	Shougen:                "Shougen",
	Bunnou:                 "Bunnou",
	Kouchou:                "Kouchou",
	Bunnei:                 "Bunnei",
	Kenji:                  "Kenji",
	Kouan1278:              "Kouan1278",
	Shouou:                 "Shouou",
	Einin:                  "Einin",
	Shouan1299:             "Shouan1299",
	Kengen:                 "Kengen",
	Kagen:                  "Kagen",
	Tokuji:                 "Tokuji",
	Enkyou1308:             "Enkyou1308",
	Ouchou:                 "Ouchou",
	Shouwa1312:             "Shouwa1312",
	Bunpou:                 "Bunpou",
	Gennou:                 "Gennou",
	Genkou1321:             "Genkou1321",
	Shouchuu:               "Shouchuu",
	Karyaku:                "Karyaku",
	GentokuD:               "GentokuD",
	GentokuJ:               "GentokuJ",
	Genkou1331:             "Genkou1331",
	KenmuN:                 "KenmuN",
	Engen:                  "Engen",
	Koukoku:                "Koukoku",
	Shouhei:                "Shouhei",
	Kentoku:                "Kentoku",
	Bunchuu:                "Bunchuu",
	Tenju:                  "Tenju",
	Kouwa1381:              "Kouwa1381",
	Genchuu:                "Genchuu",
	KenmuH:                 "KenmuH",
	RyakuouRekiou:          "RyakuouRekiou",
	Kouei:                  "Kouei",
	JouwaTeiwa:             "JouwaTeiwa",
	KannnouKannou:          "KannnouKannou",
	BunnnaBunwa:            "BunnnaBunwa",
	Enbun:                  "Enbun",
	Kouan1361:              "Kouan1361",
	JoujiTeiji:             "JoujiTeiji",
	Ouan:                   "Ouan",
	Eiwa:                   "Eiwa",
	Kouryaku:               "Kouryaku",
	Eitoku:                 "Eitoku",
	Shitoku:                "Shitoku",
	KakyouKakei:            "KakyouKakei",
	Kouou:                  "Kouou",
	Meitoku:                "Meitoku",
	Ouei:                   "Ouei",
	Shouchou:               "Shouchou",
	Eikyou:                 "Eikyou",
	Kakitsu:                "Kakitsu",
	Bunnan:                 "Bunnan",
	Houtoku:                "Houtoku",
	Kyoutoku:               "Kyoutoku",
	Koushou:                "Koushou",
	Chouroku:               "Chouroku",
	Kanshou:                "Kanshou",
	Bunshou:                "Bunshou",
	Ounin:                  "Ounin",
	Bunmei:                 "Bunmei",
	Choukyou:               "Choukyou",
	Entoku:                 "Entoku",
	Meiou:                  "Meiou",
	Bunki:                  "Bunki",
	Eishou:                 "Eishou",
	Daiei:                  "Daiei",
	Kyouroku:               "Kyouroku",
	Tenbun:                 "Tenbun",
	Kouji1555:              "Kouji1555",
	Eiroku:                 "Eiroku",
	Genki:                  "Genki",
	Tenshou:                "Tenshou",
	Bunroku:                "Bunroku",
	Keichou:                "Keichou",
	Gennna:                 "Gennna",
	Kannei:                 "Kannei",
	Shouhou:                "Shouhou",
	Keian:                  "Keian",
	Jouou1652:              "Jouou1652",
	Meireki:                "Meireki",
	Manji:                  "Manji",
	Kanbun:                 "Kanbun",
	Enpou:                  "Enpou",
	Tennna:                 "Tennna",
	Joukyou:                "Joukyou",
	Genroku:                "Genroku",
	Houei:                  "Houei",
	Shoutoku:               "Shoutoku",
	Kyouhou:                "Kyouhou",
	Genbun:                 "Genbun",
	Kanpou:                 "Kanpou",
	Enkyou1744:             "Enkyou1744",
	Kannen:                 "Kannen",
	Houreki:                "Houreki",
	Meiwa:                  "Meiwa",
	Annei:                  "Annei",
	Tenmei:                 "Tenmei",
	Kansei:                 "Kansei",
	Kyouwa:                 "Kyouwa",
	Bunka:                  "Bunka",
	Bunsei:                 "Bunsei",
	Tenpou:                 "Tenpou",
	Kouka:                  "Kouka",
	Kaei:                   "Kaei",
	Ansei:                  "Ansei",
	Mannen:                 "Mannen",
	Bunkyuu:                "Bunkyuu",
	Genji:                  "Genji",
	Keiou:                  "Keiou",
	Meiji:                  "Meiji",
	Taishou:                "Taishou",
	Shouwa1926:             "Shouwa1926",
	Heisei:                 "Heisei",
}
