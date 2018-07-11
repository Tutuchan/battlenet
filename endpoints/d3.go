package endpoints

import (
	"strconv"
)

var (
	// D3 API
	endpointD3     = API() + "d3/"
	endpointD3Data = endpointD3 + "data/"

	// D3 ACT
	// 	GETACTINDEX /D3/DATA/ACT
	D3Acts = endpointD3Data + "act"
	// 	GETACT 		/D3/DATA/ACT/{ACTID}
	GetD3Act = func(actID int) string { return D3Acts + "/" + strconv.Itoa(actID) }

	// D3 ARTISAN AND RECIPE
	endpointD3Artisan = endpointD3Data + "artisan/"
	// 	GETARTISAN 	/D3/DATA/ARTISAN/{ARTISANSLUG}
	GetD3Artisan = func(artisanSlug string) string { return endpointD3Artisan + artisanSlug }
	// 	GETRECIPE 	/D3/DATA/ARTISAN/{ARTISANSLUG}/RECIPE/{RECIPESLUG}
	GetD3Recipe = func(artisanSlug, recipeSlug string) string { return GetD3Artisan(artisanSlug) + "/" + recipeSlug }

	// D3 FOLLOWER
	endpointD3Follower = endpointD3Data + "follower/"
	// 	GETFOLLOWER /D3/DATA/FOLLOWER/{FOLLOWERSLUG}
	GetD3Follower = func(followerSlug string) string { return endpointD3Follower + followerSlug }

	// D3 CHARACTER CLASS AND SKILL
	endpointD3CharacterClass = endpointD3Data + "hero/"
	// 	GETCHARACTERCLASS 	/D3/DATA/HERO/{CLASSSLUG}
	GetD3CharacterClass = func(classSlug string) string { return endpointD3CharacterClass + classSlug }
	// 	GETAPISKILL 		/D3/DATA/HERO/{CLASSSLUG}/SKILL/{SKILLSLUG}
	GetD3ApiSkill = func(classSlug, skillSlug string) string {
		return GetD3CharacterClass(classSlug) + "/skill/" + skillSlug
	}

	// D3 ITEM TYPE

	// 	GETITEMTYPEINDEX 	/D3/DATA/ITEM-TYPE

	// 	GETITEMTYPE 		/D3/DATA/ITEM-TYPE/{ITEMTYPESLUG}

	// D3 ITEM

	// 	GETITEM /D3/DATA/ITEM/{ITEMSLUGANDID}

	// D3 PROFILE

	// 	GETAPIACCOUNT 				/D3/PROFILE/{ACCOUNT}/

	// 	GETAPIHERO 					/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}

	// 	GETAPIDETAILEDHEROITEMS 	/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/ITEMS

	// 	GETAPIDETAILEDFOLLOWERITEMS /D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/FOLLOWER-ITEMS

	// D3 GAME DATA API
	// --------------------------------------------------------------------
	// D3
	endpointDataD3 = endpointData + "d3/"
	// 	SEASON INDEX 		/DATA/D3/SEASON/
	D3SeasonIndex = endpointDataD3 + "season/"
	// 	SEASON 				/DATA/D3/SEASON/:ID
	D3Season = func(id int) string { return D3SeasonIndex + strconv.Itoa(id) }
	// 	SEASON LEADERBOARD 	/DATA/D3/SEASON/:ID/LEADERBOARD/:LEADERBOARD
	D3SeasonLeaderboard = func(id int, leaderboard string) string { return D3Season(id) + "/leaderboard/" + leaderboard }
	// 	ERA INDEX 			/DATA/D3/ERA/
	D3EraIndex = endpointD3Data + "era/"
	// 	ERA 				/DATA/D3/ERA/:ID
	D3Era = func(id int) string { return D3EraIndex + strconv.Itoa(id) }
	// 	ERA LEADERBOARD 	/DATA/D3/ERA/:ID/LEADERBOARD/:LEADERBOARD
	D3EraLeaderboard = func(id int, leaderboard string) string { return D3Era(id) + "/leaderboard/" + leaderboard }
)
