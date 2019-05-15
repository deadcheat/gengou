package gengou

import (
	"bytes"
	"encoding/json"
)

// GengouCode is a type alias for using Codes of 元号
type GengouCode int

// GengouCodeFromString get code from string
func GengouCodeFromString(s string) GengouCode {
	return idMap[s]
}

// GengouCodeToString make code to string
func GengouCodeToString(c GengouCode) string {
	return strMap[c]
}

// MarshalJSON marshals the enum as a quoted json string
func (c GengouCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(GengouCodeToString(c))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (c *GengouCode) UnmarshalJSON(b []byte) error {
	var j string
	_ = json.Unmarshal(b, &j)
	*c = GengouCodeFromString(j)
	return nil
}

const (
	Taika GengouCode = iota
	Hakuchi
	Undecided0654
	Shuchou
	Undecided0686
	Taihou
	Keiun
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
	Tennou
	Enryaku
	Daidou
	Kounin
	Tenchou
	Jouwa834
	Kashou848
	Ninju
	Saikou
	Tennan
	Jougan859
	Gangyou
	Ninnna
	Kanpyou
	Shoutai
	Engi
	Enchou
	Jouhei
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
	Eishou1046
	Tenki
	Kouhei
	Jiryaku
	Enkyuu
	Jouhou
	Jouryaku
	Eihou
	Outoku
	Kanji
	Kahou
	Eichou
	Joutoku
	Kouwa1099
	Chouji
	Kashou1106
	Tennnin
	Tennei
	Eikyuu
	Gennei
	Houan
	Tenji
	Daiji
	Tenshou1131
	Choushou
	Houen
	Eiji
	Kouji1142
	Tennyou
	Kyuuan
	Ninpei
	Kyuuju
	Hougen
	Heiji
	Eiryaku
	Ouhou
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
	Ryakuou
	Kouei
	Jouwa1345
	Kannnou
	Bunnna
	Enbun
	Kouan1361
	Jouji
	Ouan
	Eiwa
	Kouryaku
	Eitoku
	Shitoku
	Kakyou
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
	Eishou1504
	Daiei
	Kyouroku
	Tenbun
	Kouji1555
	Eiroku
	Genki
	Tenshou1571
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
	Reiwa
)

var idMap = map[string]GengouCode{
	"Taika":          Taika,
	"Hakuchi":        Hakuchi,
	"Shuchou":        Shuchou,
	"Taihou":         Taihou,
	"Keiun":          Keiun,
	"Wadou":          Wadou,
	"Reiki":          Reiki,
	"Yourou":         Yourou,
	"Jinki":          Jinki,
	"Tenpyou":        Tenpyou,
	"Tenpyoukanpou":  Tenpyoukanpou,
	"Tenpyoushouhou": Tenpyoushouhou,
	"Tenpyouhouji":   Tenpyouhouji,
	"Tenpyoujingo":   Tenpyoujingo,
	"Jingokeiun":     Jingokeiun,
	"Houki":          Houki,
	"Tennou":         Tennou,
	"Enryaku":        Enryaku,
	"Daidou":         Daidou,
	"Kounin":         Kounin,
	"Tenchou":        Tenchou,
	"Jouwa834":       Jouwa834,
	"Kashou848":      Kashou848,
	"Ninju":          Ninju,
	"Saikou":         Saikou,
	"Tennan":         Tennan,
	"Jougan859":      Jougan859,
	"Gangyou":        Gangyou,
	"Ninnna":         Ninnna,
	"Kanpyou":        Kanpyou,
	"Shoutai":        Shoutai,
	"Engi":           Engi,
	"Enchou":         Enchou,
	"Jouhei":         Jouhei,
	"Tengyou":        Tengyou,
	"Tenryaku":       Tenryaku,
	"Tentoku":        Tentoku,
	"Ouwa":           Ouwa,
	"Kouhou":         Kouhou,
	"AnnnaAnwa":      AnnnaAnwa,
	"Tenroku":        Tenroku,
	"Tennen":         Tennen,
	"Jougen976":      Jougen976,
	"Tengen":         Tengen,
	"Eikan":          Eikan,
	"Kannna":         Kannna,
	"Eien":           Eien,
	"Eiso":           Eiso,
	"Shouryaku":      Shouryaku,
	"Choutoku":       Choutoku,
	"Chouhou":        Chouhou,
	"Kankou":         Kankou,
	"Chouwa":         Chouwa,
	"Kannnin":        Kannnin,
	"Jian":           Jian,
	"Manju":          Manju,
	"Chougen":        Chougen,
	"Chouryaku":      Chouryaku,
	"Choukyuu":       Choukyuu,
	"Kantoku":        Kantoku,
	"Eishou1046":     Eishou1046,
	"Tenki":          Tenki,
	"Kouhei":         Kouhei,
	"Jiryaku":        Jiryaku,
	"Enkyuu":         Enkyuu,
	"Jouhou":         Jouhou,
	"Jouryaku":       Jouryaku,
	"Eihou":          Eihou,
	"Outoku":         Outoku,
	"Kanji":          Kanji,
	"Kahou":          Kahou,
	"Eichou":         Eichou,
	"Joutoku":        Joutoku,
	"Kouwa1099":      Kouwa1099,
	"Chouji":         Chouji,
	"Kashou1106":     Kashou1106,
	"Tennnin":        Tennnin,
	"Tennei":         Tennei,
	"Eikyuu":         Eikyuu,
	"Gennei":         Gennei,
	"Houan":          Houan,
	"Tenji":          Tenji,
	"Daiji":          Daiji,
	"Tenshou1131":    Tenshou1131,
	"Choushou":       Choushou,
	"Houen":          Houen,
	"Eiji":           Eiji,
	"Kouji1142":      Kouji1142,
	"Tennyou":        Tennyou,
	"Kyuuan":         Kyuuan,
	"Ninpei":         Ninpei,
	"Kyuuju":         Kyuuju,
	"Hougen":         Hougen,
	"Heiji":          Heiji,
	"Eiryaku":        Eiryaku,
	"Ouhou":          Ouhou,
	"Choukan":        Choukan,
	"Eiman":          Eiman,
	"Ninnan":         Ninnan,
	"Kaou":           Kaou,
	"Shouan1171":     Shouan1171,
	"Angen":          Angen,
	"Jishou":         Jishou,
	"Youwa":          Youwa,
	"Juei":           Juei,
	"Genryaku":       Genryaku,
	"Bunji":          Bunji,
	"Kenkyuu":        Kenkyuu,
	"Shouji":         Shouji,
	"Kennnin":        Kennnin,
	"Genkyuu":        Genkyuu,
	"Kennei":         Kennei,
	"Jougen1207":     Jougen1207,
	"Kenryaku":       Kenryaku,
	"Kenpou":         Kenpou,
	"Joukyuu":        Joukyuu,
	"Jouou1222":      Jouou1222,
	"Gennnin":        Gennnin,
	"Karoku":         Karoku,
	"Antei":          Antei,
	"Kanki":          Kanki,
	"Jouei":          Jouei,
	"Tenpuku":        Tenpuku,
	"Bunryaku":       Bunryaku,
	"Katei":          Katei,
	"Ryakunin":       Ryakunin,
	"Ennou":          Ennou,
	"Ninji":          Ninji,
	"Kangen":         Kangen,
	"Houji":          Houji,
	"Kenchou":        Kenchou,
	"Kougen":         Kougen,
	"Shouka":         Shouka,
	"Shougen":        Shougen,
	"Bunnou":         Bunnou,
	"Kouchou":        Kouchou,
	"Bunnei":         Bunnei,
	"Kenji":          Kenji,
	"Kouan1278":      Kouan1278,
	"Shouou":         Shouou,
	"Einin":          Einin,
	"Shouan1299":     Shouan1299,
	"Kengen":         Kengen,
	"Kagen":          Kagen,
	"Tokuji":         Tokuji,
	"Enkyou1308":     Enkyou1308,
	"Ouchou":         Ouchou,
	"Shouwa1312":     Shouwa1312,
	"Bunpou":         Bunpou,
	"Gennou":         Gennou,
	"Genkou1321":     Genkou1321,
	"Shouchuu":       Shouchuu,
	"Karyaku":        Karyaku,
	"GentokuD":       GentokuD,
	"GentokuJ":       GentokuJ,
	"Genkou1331":     Genkou1331,
	"KenmuN":         KenmuN,
	"Engen":          Engen,
	"Koukoku":        Koukoku,
	"Shouhei":        Shouhei,
	"Kentoku":        Kentoku,
	"Bunchuu":        Bunchuu,
	"Tenju":          Tenju,
	"Kouwa1381":      Kouwa1381,
	"Genchuu":        Genchuu,
	"KenmuH":         KenmuH,
	"Ryakuou":        Ryakuou,
	"Kouei":          Kouei,
	"Jouwa1345":      Jouwa1345,
	"Kannnou":        Kannnou,
	"Bunnna":         Bunnna,
	"Enbun":          Enbun,
	"Kouan1361":      Kouan1361,
	"Jouji":          Jouji,
	"Ouan":           Ouan,
	"Eiwa":           Eiwa,
	"Kouryaku":       Kouryaku,
	"Eitoku":         Eitoku,
	"Shitoku":        Shitoku,
	"Kakyou":         Kakyou,
	"Kouou":          Kouou,
	"Meitoku":        Meitoku,
	"Ouei":           Ouei,
	"Shouchou":       Shouchou,
	"Eikyou":         Eikyou,
	"Kakitsu":        Kakitsu,
	"Bunnan":         Bunnan,
	"Houtoku":        Houtoku,
	"Kyoutoku":       Kyoutoku,
	"Koushou":        Koushou,
	"Chouroku":       Chouroku,
	"Kanshou":        Kanshou,
	"Bunshou":        Bunshou,
	"Ounin":          Ounin,
	"Bunmei":         Bunmei,
	"Choukyou":       Choukyou,
	"Entoku":         Entoku,
	"Meiou":          Meiou,
	"Bunki":          Bunki,
	"Eishou1504":     Eishou1504,
	"Daiei":          Daiei,
	"Kyouroku":       Kyouroku,
	"Tenbun":         Tenbun,
	"Kouji1555":      Kouji1555,
	"Eiroku":         Eiroku,
	"Genki":          Genki,
	"Tenshou1571":    Tenshou1571,
	"Bunroku":        Bunroku,
	"Keichou":        Keichou,
	"Gennna":         Gennna,
	"Kannei":         Kannei,
	"Shouhou":        Shouhou,
	"Keian":          Keian,
	"Jouou1652":      Jouou1652,
	"Meireki":        Meireki,
	"Manji":          Manji,
	"Kanbun":         Kanbun,
	"Enpou":          Enpou,
	"Tennna":         Tennna,
	"Joukyou":        Joukyou,
	"Genroku":        Genroku,
	"Houei":          Houei,
	"Shoutoku":       Shoutoku,
	"Kyouhou":        Kyouhou,
	"Genbun":         Genbun,
	"Kanpou":         Kanpou,
	"Enkyou1744":     Enkyou1744,
	"Kannen":         Kannen,
	"Houreki":        Houreki,
	"Meiwa":          Meiwa,
	"Annei":          Annei,
	"Tenmei":         Tenmei,
	"Kansei":         Kansei,
	"Kyouwa":         Kyouwa,
	"Bunka":          Bunka,
	"Bunsei":         Bunsei,
	"Tenpou":         Tenpou,
	"Kouka":          Kouka,
	"Kaei":           Kaei,
	"Ansei":          Ansei,
	"Mannen":         Mannen,
	"Bunkyuu":        Bunkyuu,
	"Genji":          Genji,
	"Keiou":          Keiou,
	"Meiji":          Meiji,
	"Taishou":        Taishou,
	"Shouwa1926":     Shouwa1926,
	"Heisei":         Heisei,
	"Reiwa":          Reiwa,
}

var strMap = map[GengouCode]string{
	Taika:          "Taika",
	Hakuchi:        "Hakuchi",
	Shuchou:        "Shuchou",
	Taihou:         "Taihou",
	Keiun:          "Keiun",
	Wadou:          "Wadou",
	Reiki:          "Reiki",
	Yourou:         "Yourou",
	Jinki:          "Jinki",
	Tenpyou:        "Tenpyou",
	Tenpyoukanpou:  "Tenpyoukanpou",
	Tenpyoushouhou: "Tenpyoushouhou",
	Tenpyouhouji:   "Tenpyouhouji",
	Tenpyoujingo:   "Tenpyoujingo",
	Jingokeiun:     "Jingokeiun",
	Houki:          "Houki",
	Tennou:         "Tennou",
	Enryaku:        "Enryaku",
	Daidou:         "Daidou",
	Kounin:         "Kounin",
	Tenchou:        "Tenchou",
	Jouwa834:       "Jouwa834",
	Kashou848:      "Kashou848",
	Ninju:          "Ninju",
	Saikou:         "Saikou",
	Tennan:         "Tennan",
	Jougan859:      "Jougan859",
	Gangyou:        "Gangyou",
	Ninnna:         "Ninnna",
	Kanpyou:        "Kanpyou",
	Shoutai:        "Shoutai",
	Engi:           "Engi",
	Enchou:         "Enchou",
	Jouhei:         "Jouhei",
	Tengyou:        "Tengyou",
	Tenryaku:       "Tenryaku",
	Tentoku:        "Tentoku",
	Ouwa:           "Ouwa",
	Kouhou:         "Kouhou",
	AnnnaAnwa:      "AnnnaAnwa",
	Tenroku:        "Tenroku",
	Tennen:         "Tennen",
	Jougen976:      "Jougen976",
	Tengen:         "Tengen",
	Eikan:          "Eikan",
	Kannna:         "Kannna",
	Eien:           "Eien",
	Eiso:           "Eiso",
	Shouryaku:      "Shouryaku",
	Choutoku:       "Choutoku",
	Chouhou:        "Chouhou",
	Kankou:         "Kankou",
	Chouwa:         "Chouwa",
	Kannnin:        "Kannnin",
	Jian:           "Jian",
	Manju:          "Manju",
	Chougen:        "Chougen",
	Chouryaku:      "Chouryaku",
	Choukyuu:       "Choukyuu",
	Kantoku:        "Kantoku",
	Eishou1046:     "Eishou1046",
	Tenki:          "Tenki",
	Kouhei:         "Kouhei",
	Jiryaku:        "Jiryaku",
	Enkyuu:         "Enkyuu",
	Jouhou:         "Jouhou",
	Jouryaku:       "Jouryaku",
	Eihou:          "Eihou",
	Outoku:         "Outoku",
	Kanji:          "Kanji",
	Kahou:          "Kahou",
	Eichou:         "Eichou",
	Joutoku:        "Joutoku",
	Kouwa1099:      "Kouwa1099",
	Chouji:         "Chouji",
	Kashou1106:     "Kashou1106",
	Tennnin:        "Tennnin",
	Tennei:         "Tennei",
	Eikyuu:         "Eikyuu",
	Gennei:         "Gennei",
	Houan:          "Houan",
	Tenji:          "Tenji",
	Daiji:          "Daiji",
	Tenshou1131:    "Tenshou1131",
	Choushou:       "Choushou",
	Houen:          "Houen",
	Eiji:           "Eiji",
	Kouji1142:      "Kouji1142",
	Tennyou:        "Tennyou",
	Kyuuan:         "Kyuuan",
	Ninpei:         "Ninpei",
	Kyuuju:         "Kyuuju",
	Hougen:         "Hougen",
	Heiji:          "Heiji",
	Eiryaku:        "Eiryaku",
	Ouhou:          "Ouhou",
	Choukan:        "Choukan",
	Eiman:          "Eiman",
	Ninnan:         "Ninnan",
	Kaou:           "Kaou",
	Shouan1171:     "Shouan1171",
	Angen:          "Angen",
	Jishou:         "Jishou",
	Youwa:          "Youwa",
	Juei:           "Juei",
	Genryaku:       "Genryaku",
	Bunji:          "Bunji",
	Kenkyuu:        "Kenkyuu",
	Shouji:         "Shouji",
	Kennnin:        "Kennnin",
	Genkyuu:        "Genkyuu",
	Kennei:         "Kennei",
	Jougen1207:     "Jougen1207",
	Kenryaku:       "Kenryaku",
	Kenpou:         "Kenpou",
	Joukyuu:        "Joukyuu",
	Jouou1222:      "Jouou1222",
	Gennnin:        "Gennnin",
	Karoku:         "Karoku",
	Antei:          "Antei",
	Kanki:          "Kanki",
	Jouei:          "Jouei",
	Tenpuku:        "Tenpuku",
	Bunryaku:       "Bunryaku",
	Katei:          "Katei",
	Ryakunin:       "Ryakunin",
	Ennou:          "Ennou",
	Ninji:          "Ninji",
	Kangen:         "Kangen",
	Houji:          "Houji",
	Kenchou:        "Kenchou",
	Kougen:         "Kougen",
	Shouka:         "Shouka",
	Shougen:        "Shougen",
	Bunnou:         "Bunnou",
	Kouchou:        "Kouchou",
	Bunnei:         "Bunnei",
	Kenji:          "Kenji",
	Kouan1278:      "Kouan1278",
	Shouou:         "Shouou",
	Einin:          "Einin",
	Shouan1299:     "Shouan1299",
	Kengen:         "Kengen",
	Kagen:          "Kagen",
	Tokuji:         "Tokuji",
	Enkyou1308:     "Enkyou1308",
	Ouchou:         "Ouchou",
	Shouwa1312:     "Shouwa1312",
	Bunpou:         "Bunpou",
	Gennou:         "Gennou",
	Genkou1321:     "Genkou1321",
	Shouchuu:       "Shouchuu",
	Karyaku:        "Karyaku",
	GentokuD:       "GentokuD",
	GentokuJ:       "GentokuJ",
	Genkou1331:     "Genkou1331",
	KenmuN:         "KenmuN",
	Engen:          "Engen",
	Koukoku:        "Koukoku",
	Shouhei:        "Shouhei",
	Kentoku:        "Kentoku",
	Bunchuu:        "Bunchuu",
	Tenju:          "Tenju",
	Kouwa1381:      "Kouwa1381",
	Genchuu:        "Genchuu",
	KenmuH:         "KenmuH",
	Ryakuou:        "Ryakuou",
	Kouei:          "Kouei",
	Jouwa1345:      "Jouwa1345",
	Kannnou:        "Kannnou",
	Bunnna:         "Bunnna",
	Enbun:          "Enbun",
	Kouan1361:      "Kouan1361",
	Jouji:          "Jouji",
	Ouan:           "Ouan",
	Eiwa:           "Eiwa",
	Kouryaku:       "Kouryaku",
	Eitoku:         "Eitoku",
	Shitoku:        "Shitoku",
	Kakyou:         "Kakyou",
	Kouou:          "Kouou",
	Meitoku:        "Meitoku",
	Ouei:           "Ouei",
	Shouchou:       "Shouchou",
	Eikyou:         "Eikyou",
	Kakitsu:        "Kakitsu",
	Bunnan:         "Bunnan",
	Houtoku:        "Houtoku",
	Kyoutoku:       "Kyoutoku",
	Koushou:        "Koushou",
	Chouroku:       "Chouroku",
	Kanshou:        "Kanshou",
	Bunshou:        "Bunshou",
	Ounin:          "Ounin",
	Bunmei:         "Bunmei",
	Choukyou:       "Choukyou",
	Entoku:         "Entoku",
	Meiou:          "Meiou",
	Bunki:          "Bunki",
	Eishou1504:     "Eishou1504",
	Daiei:          "Daiei",
	Kyouroku:       "Kyouroku",
	Tenbun:         "Tenbun",
	Kouji1555:      "Kouji1555",
	Eiroku:         "Eiroku",
	Genki:          "Genki",
	Tenshou1571:    "Tenshou1571",
	Bunroku:        "Bunroku",
	Keichou:        "Keichou",
	Gennna:         "Gennna",
	Kannei:         "Kannei",
	Shouhou:        "Shouhou",
	Keian:          "Keian",
	Jouou1652:      "Jouou1652",
	Meireki:        "Meireki",
	Manji:          "Manji",
	Kanbun:         "Kanbun",
	Enpou:          "Enpou",
	Tennna:         "Tennna",
	Joukyou:        "Joukyou",
	Genroku:        "Genroku",
	Houei:          "Houei",
	Shoutoku:       "Shoutoku",
	Kyouhou:        "Kyouhou",
	Genbun:         "Genbun",
	Kanpou:         "Kanpou",
	Enkyou1744:     "Enkyou1744",
	Kannen:         "Kannen",
	Houreki:        "Houreki",
	Meiwa:          "Meiwa",
	Annei:          "Annei",
	Tenmei:         "Tenmei",
	Kansei:         "Kansei",
	Kyouwa:         "Kyouwa",
	Bunka:          "Bunka",
	Bunsei:         "Bunsei",
	Tenpou:         "Tenpou",
	Kouka:          "Kouka",
	Kaei:           "Kaei",
	Ansei:          "Ansei",
	Mannen:         "Mannen",
	Bunkyuu:        "Bunkyuu",
	Genji:          "Genji",
	Keiou:          "Keiou",
	Meiji:          "Meiji",
	Taishou:        "Taishou",
	Shouwa1926:     "Shouwa1926",
	Heisei:         "Heisei",
	Reiwa:          "Reiwa",
}
