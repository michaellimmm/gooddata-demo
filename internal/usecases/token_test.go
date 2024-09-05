package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	type arg struct {
		privateKey string
		payload    TokenKey
	}

	const src = `
		{
			"alg":"RS256",
			"d":"eFDa_o25k18IUqgGCfXLXZRrci3Tbodtdd2pKOP9o_dlqSMopY_QrQgf5Ba2qQ1DHZ9xPrt8dHSvHZdwlgSWH9Yg6THXkz6dXkuchNUe6hCPgCd77XuPnErQVVsHQi7mrVSA7PGym0pL1OA0PkY-jRf-kjMwPiKBtBumdMG7dfe-LKwsbi2p3wejbcJW0KToL32FQqhjaVt8pN-lINC7LP8yXJtarl5xmdeu3rOtDjIRrCBGxA2o-3rMZU17403CMdV9q6KY-akpRD3m3Foq0Hlo4u3KJj0ZTDJ7-dgQRkp4yUj--RpZu5lmfnoRFLOxDzQ_uPlro217Nlgn8px1QQ",
			"dp":"ISp4j1i8bqV5I5Iop9mzrbDKmSxZ3107rZDe52boyeVYrQdUIzzplir2rjv2FyJEBeETb4x6y0FHKDNEx9PHEkYJSsDrsywaCVOv2xqI8ajbCMbnSkIRf_6heAl-SZS1i81VlJ65Jrm7o4XmPS5TYp7abZtEpdYWMpBMx8c8qoE",
			"dq":"QXiLePMew13MoLFni-HUmQuOdid2CMMu0pPBjSroTQXoqGyHrCNj6k71ZASIIe8LcgePFwkl9HC2gNZZokw842cFvP4MlBRAkj74o-cDNd3H-Tky6xls87Degi10sGULb6uxzyvLvhmW94IR1ZXm-5gfhiFSYh2ZpOm3uPK3PWk",
			"e":"AQAB",
			"kid":"zeals_gooddata_kid",
			"kty":"RSA",
			"n":"zp4NBeolXwz6Yh1RpM81K7AuTZgTdK3fgFpdU4sU-dC-8SkZUHrTibOuieUlWwxPeD7w393WSfGnwCEnB9YP5qUryt0d0xxizXKnLnNxs_ecq_YSlnxBETjy8MZtw_ou5BUNN8gh15VPGkoGOnGnVAyOJTAbRT0u5nVCxKQ95mrHdGULUlQMOlCKhP4v2dP4tT4PGKMVV0pZyI6D58XTRN6wx3eomnjfL-9ZeEPcudnfBcfixuLV55jF6TppUUhl93eh-dDyaUuTWt2Elo7dPPWAwh8mcOlK3wd89ykgDJOftEEC1FXaKt5qze4aooYD3s8BWw4gDYb2q5xAvJZ_ew",
			"p":"40StSPK3tmNagCRNW2Rq1nhaNDidD7cPmYzqDu5WvyyAi5307wEXvpiYSHLkkPrWkYj5sx9vBciabo8tKJgB4AQp_er5Jm6iPxv5-MiKOu-HsvZ5BKPqNpu3X0NNEHYuDZFUGc_CeX4VIETI8T1dfu-KIVf2m4lDJi7g2GKfJmE",
			"q":"6L0HEySZY-ABwaciO4_PDPGbfXGlBdlWAXrdL-KG29D6XfDITJTbedjIOFkJz1UcnqyInhSCYfXSJIEBnqiB94AnmkwBXNMhP-yZj-iMXuHPVH95bBepiMX80h7dz5cTcz2eSIGt0mrTMQ8umKFALGwhhiY4l0zeblsn3Jjbu1s",
			"qi":"M5fcIqBEsLjreq2_28Zxzaqfve6V4zykMwUjrZ7xjUjhOGcelhVqso_JliqJV8X0fAMJ2EoXNuXSEF5xtWcA26GnykOaGHFLcCb6NJ9h6vbekjo10Y6E9YWc9dEqVWkjxI1GFqkGkJ12myM4vGxhk9n0XaQPQisDy5a6ltZ8l4k"
		}`

	tests := []struct {
		name string
		arg  arg
	}{
		{
			name: "Given valid request Then return accesstoken",
			arg: arg{
				privateKey: src,
				payload: TokenKey{
					Kid: "zeals_gooddata_kid",
					Sub: "testuser",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arg := test.arg
			actual, err := GenerateToken(arg.privateKey, arg.payload)
			assert.Nil(t, err)
			assert.NotNil(t, actual)
		})
	}
}
