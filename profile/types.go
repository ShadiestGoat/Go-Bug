package profile

/*
╭──────────────────────────────╮
│            Common            │
╰──────────────────────────────╯
*/

type PoliticalCompass struct {
	Liberal float64
	RightWing float64
}

type Infp struct {
	SI float64
	IE float64
	FT float64
	JP float64
}

type Annoying struct {
	Edge float64
	Ego float64
	Temper float64
	Mature float64
	Selfish float64
	Talkative float64
}

type Personality struct {
	Creepy float64
	Cute float64
	Funny float64
	Kind float64
	OpenMind float64
	Chaotic float64
	Evil float64
}

type Politics struct {
	Engagement float64
}

type PersonProfileBase struct {
	Annoying Annoying
	Personality Personality
	Infp Infp
}

type WeightsPersonality struct {
	Personality
	Annoying float64
	INFP float64
	Politics float64
}

type Weights struct {
	Annoying Annoying
	Personality WeightsPersonality
	Politics IdealPolitics
	Infp Infp
}

type Voted struct {
	PersonProfileBase
	Politics Politics
}

type IdealPolitics struct {
	Politics
	PoliticalCompass
}

type Ideal struct {
	PersonProfileBase
	Politics IdealPolitics
}
