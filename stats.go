package main

type MainStats struct {
	STR int `json:"Strength,omitempty"`
	DEX int `json:"Dexterity,omitempty"`
	VIT int `json:"Vitality,omitempty"`
	INT int `json:"Intelligence,omitempty"`
	MND int `json:"Mind,omitempty"`
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
	CRIT int `json:"Critical Hit,omitempty"`
	DET  int `json:"Determination,omitempty"`
	DH   int `json:"Direct Hit Rate,omitempty"`
	SKS  int `json:"Skill Speed,omitempty"`
	SPS  int `json:"Spell Speed,omitempty"`
	TNC  int `json:"Tenacity,omitempty"`
	PT   int `json:"Piety,omitempty"`
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
