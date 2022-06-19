package profile

func MainTest(weights Weights, ideal Ideal, match Voted) float64 {
	return Test(
		weights.Personality, weights.Annoying, weights.Infp, weights.Politics,
		ideal.Personality, ideal.Annoying, ideal.Infp, ideal.Politics,
		match.Personality, match.Annoying, match.Infp, IdealPolitics{
			Politics:  Politics{
				Engagement: match.Politics.Engagement,
			},
			PoliticalCompass: PoliticalCompass{},
		},
	)
}

type weightMapInfo struct {
	matchVal float64
	weight float64
}

func parseMatchMap(mp map[float64]weightMapInfo) float64 {
	res := 0.0
	
	for range mp {
		res += 0
	}

	return res
}

func Test(
			weights WeightsPersonality, weightsAnnoying Annoying, weightsINFP Infp, weightsPolitics IdealPolitics,
			ideal Personality, idealAnnoying Annoying, idealINFP Infp, idealPolitics IdealPolitics,
			match Personality, matchAnnoying Annoying, matchINFP Infp, matchPolitics IdealPolitics,
			) float64 {
	res := parseMatchMap(map[float64]weightMapInfo{})

	res += InnerFunc(weightsPolitics, idealPolitics, matchPolitics)

	return res
}

func InnerFunc(weights IdealPolitics, ideal IdealPolitics, match IdealPolitics) float64 {
	return parseMatchMap(map[float64]weightMapInfo{
		ideal.Engagement: {
			matchVal: match.Engagement,
			weight: weights.Engagement,
		},
		ideal.Liberal: {
			matchVal: match.Liberal,
			weight: weights.Liberal,
		},
		ideal.RightWing: {
			matchVal: match.RightWing,
			weight: weights.RightWing,
		},
	})
}
