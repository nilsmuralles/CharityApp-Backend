package models

import (
	"time"
)

type CampaignDonationsResponse struct {
	Campaign                string  `json:"campaign"`
	MonetaryTotal           float64 `json:"monetary_total"`
	CantNoMonetaryDonations int     `json:"no_monetary_donations"`
	TotalArticulosDonados   int     `json:"articulos_donados"`
}

type TrandingDonationsResponse struct {
	Month          string  `json:"month"`
	TotalDonations float64 `json:"total_donations"`
	UniqueDonors   int     `json:"unique_donors"`
	PayMethod      string  `json:"pay_method"`
}

type VolunteerParticipationResponse struct {
	Campaign       string  `json:"campaign"`
	Volunteer      string  `json:"volunteer"`
	Participations int     `json:"participations"`
	TotalHours     float64 `json:"total_hours"`
}

type DonorsDistributionResponse struct {
	Category         string  `json:"category"`
	TotalDonors      int     `json:"total_donors"`
	MonetaryDonors   int     `json:"monetary_donors"`
	NoMonetaryDonors int     `json:"no_monetary_donors"`
	TotalAmount      float64 `json:"total_amount"`
}

type CampaignEfficiencyResponse struct {
	Campaign            string    `json:"campaign"`
	InitDate            time.Time `json:"init_date"`
	EndDate             time.Time `json:"end_date"`
	TotalDonations      float64   `json:"total_donations"`
	UniqueVolunteers    int       `json:"unique_volunteers"`
	Participations      int       `json:"participations"`
	RecognitionsAwarded int       `json:"recognitions_awarded"`
}
