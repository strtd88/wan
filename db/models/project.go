package models

type Project struct {
	ID                    int
	Name                  string
	Description           string
	Dir                   string
	Status                string
	Category              string
	CreatedAt             string
	UpdatedAt             string
	BusinessValueShort    string
	BusinessValueLong     string
	BusinessValueAchieved string
	ObsidianNotePath      string
}
