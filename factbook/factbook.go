package factbook

import (
	"fmt"
	"math/big"
	"sort"
	"time"

	"crypto/sha256"
)

type Factbook struct {
	Countries map[string]Country `json:"countries"`
}

//
func (f Factbook) CountryOfTheWeek(when time.Time) Country {
	keys := []string{}
	for key, _ := range f.Countries {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	_, week := when.ISOWeek()

	hash := sha256.New()
	/* max 52 iso weeks */
	hash.Write([]byte{byte(week)})
	data := hash.Sum(nil)
	num := big.NewInt(0).Mod(
		big.NewInt(0).SetBytes(data),
		big.NewInt(int64(len(keys))),
	).Int64()

	return f.Countries[keys[num]]
}

type Country struct {
	Data Data `json:"data"`
}

type Data struct {
	Name         string       `json:"name"`
	Introduction Introduction `json:"introduction"`
	Geography    Geography    `json:"geography"`
	People       People       `json:"people"`
	Government   Government   `json:"government"`
	Economy      Economy      `json:"economy"`
	// Energy
	// Communications Communications `json:"communnications"`
	// Transportation
	MilitaryAndSecurity struct {
		// ...
		TerroirstGroups string `json:"terrorist_groups"`
	} `json:"military_and_security"`
	TransnationalIssues TransnationalIssues `json:"transnational_issues"`
}

type Introduction struct {
	Background string `json:"background"`
}

type TransnationalIssues struct {
	Disputes             []string             `json:"disputes"`
	TraffickingInPersons TraffickingInPersons `json:"trafficking_in_persons"`
	IllicitDrugs         struct {
		Note string `json:"note"`
	} `json:"illicit_drugs"`
}

type TraffickingInPersons struct {
	CurrentSituation string `json:"current_situation"`
	TierRating       string `json:"tier_rating"`
}

type Economy struct {
	Overview string `json:"overview"`
	GDP      struct {
		// PurchasingPowerParity
		// OfficialExchangeRate
		// RealGrowthRate
		// PerCapitaPurchasingPowerParity
		// Composition
		// BySectorOfOrigin
	} `json:"gdp"`

	// GrossNationalSaving
	// AgricultureProducts
	// Industries
	// IndustrialProductionGrowthRate
	// LaborForce
	// ByOccupation
	// UnemploymentRate
	// DistributionOfFamilyIncome
	// Budget
	// Expenditrues
	// TaxesAndOtherRevenues
	// BudgetSurplus
	// PublicDebt
	// FiscalYear
	// InflatioNRate
	// CentralBankDiscountRate
	// ComercialBankPrimeLendingRate
	// StockOfNarrowMoney
	// StockOfBroadMoney
	// StockOfDomesticCredit
	// MarketValueOfPubliclyTradedShares
	// CurrentAccountBalance
	// Exports
	// Imports
	// Commodities
	// Partners
	// ReservesOfForeignExchangeAndGold
	// ExternalDebt
	// StockOfDirectForeignInvestment
	// ExchangeRates
}

type Government struct {
	CountryName struct {
		ConventionalLongForm  string `json:"conventional_long_form"`
		ConventionalShortForm string `json:"conventional_short_form"`
		LocalLongForm         string `json:"local_long_form"`
		LocalShortForm        string `json:"local_short_form"`
		Etymology             string `json:"etymology"`
	} `json:"country_name"`
	Type string `json:"government_type"`
	// Capital
	// Independency
	// NationalHolidays
	// Consititution
	LegalSystem string `json:"legal_system"`
	// Citizenship
	// Suffrage
	ExecutiveBranch struct {
		HeadOfGovernment string `json:"head_of_government"`
	} `json:"executive_branch"`
	LegislativeBranch struct {
		Description string `json:"description"`
	} `json:"legislative_branch"`
	JudicialBranch struct {
		HighestCourts string `json:"highest_courts"`
	} `json:"judicial_branch"`
	// PoliticalPartiesAndLeaders
	// PoliticalPressureGroupsAndLeaders
	// InternationalOrgnizationParticipation
	// DiplomaticRepresentaion
	FlagDescription struct {
		Description string `json:"description"`
		Note        string `json:"note"`
	} `json:"flag_description"`
	// NationalSymbol
	// Colors
	// NationalAnthem
}

type People struct {
	// Population Population `json:"population"`
	Nationality Nationality `json:"nationality"`
	// EthnicGroups EthnicGroups `json:"ethnic_groups"`
	// Languages
	// Religions
	// AgeStructure
	// DependencyRatios
	// MedianAge
	// PopulationGrowthRate
	// BirthRate
	// DeathRate
	// NetMigrationRate
	PopulationDistribution string `json:"population_distribution"`
	// Urbanization
	MajorUrbanAreas UrbanAreas `json:"major_urban_areas"`
	// SexRatio
	// MaternalMortalityRate
	// InfantMortalityRate
	// LifeExpectancyAtBirth
	// TotalFertilityRate
	// ContraceptivePrevalenceRate
	// HealthExpenditures
	// PhysiciansDensity
	// HospitalBedDensity
	// DrinkingWaterSource
	// SanitiationFacilityAccess
	// HIVAids
	// AdultObesity
	// EducationExpenditures
	// Literacy
	// SchoolLifeExpentancy
	// YouthUnemployment
	MajorInfectiousDiseases struct {
		Note string `json:"note"`
	} `json:"major_infectious_diseases"`
}

type UrbanAreas struct {
	Places Places `json:"places"`
	Date   string `json:"date"`
}

type Places []Place

type Place struct {
	Place      string  `json:"place"`
	Population float64 `json:"population"`
	IsCapital  bool    `json:"is_capital"`
}

type Nationality struct {
	Noun      string `json:"noun"`
	Adjective string `json:"adjective"`
}

type Geography struct {
	Location      string      `json:"location"`
	Coordinates   Coordinates `json:"geographic_coordinates"`
	MapReferences string      `json:"map_references`
	// Area
	// LandBoundaries
	// Coastline
	// MaritimeClaims
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
	// Elevation
	// NaturalResources
	// LandUse
	// IrrigatedLand
	PopulationDistribution string `json:"population_distribution"`
	// NaturalHazards
	// Environment
}

type Coordinate struct {
	Degrees    int    `json:"degrees"`
	Minutes    int    `json:"minutes"`
	Hemisphere string `json:"hemisphere"`
}

func (c Coordinate) GeoString() string {
	prefix := ""
	switch c.Hemisphere {
	case "S", "W":
		prefix = "-"
	}
	return fmt.Sprintf("%s%d.%d", prefix, c.Degrees, c.Minutes)
}

type Coordinates struct {
	Latitude  Coordinate `json:"latitude"`
	Longitude Coordinate `json:"longitude"`
}

func (c Coordinates) GeoString() string {
	return fmt.Sprintf(
		"%s,%s,3,0,0",
		c.Longitude.GeoString(),
		c.Latitude.GeoString(),
	)
	// 9.0,51.0,3,0,0
}
