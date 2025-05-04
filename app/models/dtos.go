package models

import (
	"time"
)

type CampaignDonationsResponse struct {
	Campaign                string    `json:"campaign" gorm:"column:campaign"`
	EndDate                 time.Time `json:"end_date" gorm:"column:end_date"`
	MonetaryTotal           float64   `json:"monetary_total" gorm:"column:monetary_total"`
	CantNoMonetaryDonations int       `json:"no_monetary_donations" gorm:"column:no_monetary_donations"`
	TotalDonatedArticles    int       `json:"total_donated_articles" gorm:"column:total_donated_articles"`
}

type TrendingDonationsResponse struct {
	Month          string  `json:"month" gorm:"column:month"`
	TotalDonations float64 `json:"total_donations" gorm:"column:total_donations"`
	UniqueDonors   int     `json:"unique_donors" gorm:"column:unique_donors"`
	PayMethod      string  `json:"pay_method" gorm:"column:pay_method"`
}

type VolunteerParticipationResponse struct {
	Campaign       string  `json:"campaign" gorm:"column:campaign"`
	Volunteer      string  `json:"volunteer" gorm:"column:volunteer"`
	Participations int     `json:"participations" gorm:"column:participations"`
	TotalHours     float64 `json:"total_hours" gorm:"column:total_hours"`
}

type DonorsDistributionResponse struct {
	Category         string  `json:"category" gorm:"column:category"`
	TotalDonors      int     `json:"total_donors" gorm:"column:total_donors"`
	MonetaryDonors   int     `json:"monetary_donors" gorm:"column:monetary_donors"`
	NoMonetaryDonors int     `json:"no_monetary_donors" gorm:"column:no_monetary_donors"`
	TotalAmount      float64 `json:"total_amount" gorm:"column:total_amount"`
}

type CampaignEfficiencyResponse struct {
	Campaign            string    `json:"campaign" gorm:"column:campaign"`
	InitDate            time.Time `json:"init_date" gorm:"column:init_date"`
	EndDate             time.Time `json:"end_date" gorm:"column:end_date"`
	TotalDonations      float64   `json:"total_donations" gorm:"column:total_donations"`
	UniqueVolunteers    int       `json:"unique_volunteers" gorm:"column:unique_volunteers"`
	Participations      int       `json:"participations" gorm:"column:participations"`
	RecognitionsAwarded int       `json:"recognitions_awarded" gorm:"column:recognitions_awarded"`
}
