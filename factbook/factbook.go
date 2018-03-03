package factbook

type Factbook struct {
	Countries map[string]Country `json:"countries"`
}

type Country struct {
	Data Data `json:"data"`
}

type Data struct {
	Name         string       `json:"name"`
	Introduction Introduction `json:"introduction"`
}

type Introduction struct {
	Background string     `json:"background"`
	Geography  Geography  `json:"geography"`
	People     People     `json:"people"`
	Goverment  Government `json:"goverment"`
	Economy    Economy    `json:"economy"`
	// Energy
	// Communications Communications `json:"communnications"`
	// Transportation
	MilitaryAndSecurity struct {
		// ...
		TerroirstGroups string `json:"terrorist_groups"`
	} `json:"military_and_security"`
	// TransnationalIssues
}

type TransnationalIssues struct {
	Disputes []string `json:"disputes"`
	// TraffickingInPersons
	// IllicitDrugs
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
	// ExecutiveBranch
	// LegislativeBranch
	// JudicialBranch
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
}

type UrbanAreas struct {
	Places Places `json:"places"`
	Date   string `json:"date"`
}

type Places []Place

type Place struct {
	Place      string `json:"place"`
	Population int    `json:"population"`
	IsCapital  bool   `json:"is_capital"`
}

type Nationality struct {
	Noun      string `json:"noun"`
	Adjective string `json:"adjective"`
}

type Geography struct {
	Location      string      `json:"location"`
	Coordinates   Coordinates `json:geographic_coordinates`
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

type Coordinates struct {
	Latitude  Coordinate `json:"latitude"`
	Longitude Coordinate `json:"longitude"`
}
