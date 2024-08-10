package main

type MainStats struct {
	STR int `json:"Strength"`
	DEX int `json:"Dexterity"`
	VIT int `json:"Vitality"`
	INT int `json:"Intelligence"`
	MND int `json:"Mind"`
}

func SumMainStats(stats ...MainStats) MainStats {
	var sum MainStats
	for _, stat := range stats {
		sum.STR += stat.STR
		sum.DEX += stat.DEX
		sum.VIT += stat.VIT
		sum.INT += stat.INT
		sum.MND += stat.MND
	}

	return sum
}

type SecondaryStats struct {
	CRIT int `json:"Critical Hit"`
	DET  int `json:"Determination"`
	DH   int `json:"Direct Hit Rate"`
	SKS  int `json:"Skill Speed"`
	SPS  int `json:"Spell Speed"`
	TNC  int `json:"Tenacity"`
	PT   int `json:"Piety"`
}

// Cap returnes capped secondary stats
func (st SecondaryStats) Cap(cap int) SecondaryStats {
	return SecondaryStats{
		CRIT: min(cap, st.CRIT),
		DET:  min(cap, st.DET),
		DH:   min(cap, st.DH),
		SKS:  min(cap, st.SKS),
		SPS:  min(cap, st.SPS),
		TNC:  min(cap, st.TNC),
		PT:   min(cap, st.PT),
	}
}

func SumSecondaryStats(stats ...SecondaryStats) SecondaryStats {
	var sum SecondaryStats
	for _, stat := range stats {
		sum.CRIT += stat.CRIT
		sum.DET += stat.DET
		sum.DH += stat.DH
		sum.SKS += stat.SKS
		sum.SPS += stat.SPS
		sum.TNC += stat.TNC
		sum.PT += stat.PT
	}

	return sum
}

type Stats struct {
	MainStats
	SecondaryStats
}

func SumStats(stats ...Stats) Stats {
	var sum Stats
	for _, stat := range stats {
		sum.MainStats = SumMainStats(sum.MainStats, stat.MainStats)
		sum.SecondaryStats = SumSecondaryStats(sum.SecondaryStats, stat.SecondaryStats)
	}

	return sum
}
