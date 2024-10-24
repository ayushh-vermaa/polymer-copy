package rewards

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// CardDetail represents the detailed information of a credit card.
type CardDetail struct {
	CardKey                 string               `bson:"card_key" json:"cardKey"`                                    // Rewards Credit Card API unique card key
	CardIssuer              string               `bson:"card_issuer" json:"cardIssuer"`                              // Issuing bank
	CardName                string               `bson:"card_name" json:"cardName"`                                  // Name of credit card
	CardNetwork             string               `bson:"card_network" json:"cardNetwork"`                            // Network (e.g., Visa)
	CardType                string               `bson:"card_type" json:"cardType"`                                  // Type (e.g., Personal)
	CardUrl                 string               `bson:"card_url" json:"cardUrl"`                                    // Card bank URL
	AnnualFee               float64              `bson:"annual_fee" json:"annualFee"`                                // Annual fee in USD
	FxFee                   float64              `bson:"fx_fee" json:"fxFee"`                                        // Foreign transaction fee (%)
	IsFxFee                 int                  `bson:"is_fx_fee" json:"isFxFee"`                                   // Does card have foreign exchange fee? (false=no, true=yes)
	CreditRange             string               `bson:"credit_range" json:"creditRange"`                            // Credit required
	BaseSpendAmount         float64              `bson:"base_spend_amount" json:"baseSpendAmount"`                   // Points earned per dollar spend
	BaseSpendEarnType       string               `bson:"base_spend_earn_type" json:"baseSpendEarnType"`              // Redemption program (e.g., American Express Membership Rewards)
	BaseSpendEarnCategory   string               `bson:"base_spend_earn_category" json:"baseSpendEarnCategory"`      // Spend earning category (e.g., Travel)
	BaseSpendEarnCurrency   string               `bson:"base_spend_earn_currency" json:"baseSpendEarnCurrency"`      // Spend earning currency type (rewards, miles, cashback, crypto, points)
	BaseSpendEarnValuation  float64              `bson:"base_spend_earn_valuation" json:"baseSpendEarnValuation"`    // Subjective valuation of a point
	BaseSpendEarnIsCash     int                  `bson:"base_spend_earn_is_cash" json:"baseSpendEarnIsCash"`         // Can points be converted to a statement credit? (false=no, true=yes)
	BaseSpendEarnCashValue  float64              `bson:"base_spend_earn_cash_value" json:"baseSpendEarnCashValue"`   // If points can be cashed out, value per point
	IsSignupBonus           int                  `bson:"is_signup_bonus" json:"isSignupBonus"`                       // Does card have sign-up bonus? (false=no, true=yes)
	SignupBonusAmount       string               `bson:"signup_bonus_amount" json:"signupBonusAmount"`               // Amount of sign-up miles, points, cash
	SignupBonusType         string               `bson:"signup_bonus_type" json:"signupBonusType"`                   // Redemption program for sign-up bonus
	SignupBonusCategory     string               `bson:"signup_bonus_category" json:"signupBonusCategory"`           // Sign-up bonus category (e.g., Travel)
	SignUpBonusItem         string               `bson:"sign_up_bonus_item" json:"signUpBonusItem"`                  // Sign-up bonus item (e.g., Gift Card, miles, points, cash)
	SignupBonusSpend        float64              `bson:"signup_bonus_spend" json:"signupBonusSpend"`                 // Minimum spend required
	SignupBonusLength       float64              `bson:"signup_bonus_length" json:"signupBonusLength"`               // Length of time units
	SignupBonusLengthPeriod string               `bson:"signup_bonus_length_period" json:"signupBonusLengthPeriod"`  // Period for bonus length (day, month, or year)
	SignupAnnualFee         float64              `bson:"signup_annual_fee" json:"signupAnnualFee"`                   // First year annual fee
	IsSignupAnnualFeeWaived int                  `bson:"is_signup_annual_fee_waived" json:"isSignupAnnualFeeWaived"` // Is annual fee waived first year? (false=no, true=yes)
	SignupStatementCredit   float64              `bson:"signup_statement_credit" json:"signupStatementCredit"`       // Additional sign-up bonus if applicable
	SignupBonusDesc         string               `bson:"signup_bonus_desc" json:"signupBonusDesc"`                   // Sign-up bonus description
	TrustedTraveler         string               `bson:"trusted_traveler" json:"trustedTraveler"`                    // Trusted traveler credit description
	IsTrustedTraveler       int                  `bson:"is_trusted_traveler" json:"isTrustedTraveler"`               // Does card have trusted traveler credit? (false=no, true=yes)
	LoungeAccess            string               `bson:"lounge_access" json:"loungeAccess"`                          // Lounge access benefit description
	IsLoungeAccess          int                  `bson:"is_lounge_access" json:"isLoungeAccess"`                     // Does card include lounge access? (false=no, true=yes)
	FreeHotelNight          string               `bson:"free_hotel_night" json:"freeHotelNight"`                     // Free annual hotel night certificate description
	IsFreeHotelNight        int                  `bson:"is_free_hotel_night" json:"isFreeHotelNight"`                // Does card include a free annual hotel night certificate? (false=no, true=yes)
	FreeCheckedBag          string               `bson:"free_checked_bag" json:"freeCheckedBag"`                     // Free airline checked bag description
	IsFreeCheckedBag        int                  `bson:"is_free_checked_bag" json:"isFreeCheckedBag"`                // Does card include a free checked bag? (false=no, true=yes)
	IsActive                int                  `bson:"is_active" json:"isActive"`                                  // Is card currently open for applications? (false=no, true=yes)
	Benefit                 []Benefit            `bson:"benefit" json:"benefit"`                                     // List of card benefits
	SpendBonusCategory      []SpendBonusCategory `bson:"spend_bonus_category" json:"spendBonusCategory"`             // List of spend bonus categories
	AnnualSpend             []AnnualSpend        `bson:"annual_spend" json:"annualSpend"`                            // Annual spend bonuses
}

// Benefit represents a single benefit of a credit card.
type Benefit struct {
	BenefitTitle string `bson:"benefit_title" json:"benefitTitle"` // Card benefit title
	BenefitDesc  string `bson:"benefit_desc" json:"benefitDesc"`   // Card benefit description
}

// SpendBonusCategory represents a single spend bonus category of a credit card.
type SpendBonusCategory struct {
	SpendBonusCategoryType     string  `bson:"spend_bonus_category_type" json:"spendBonusCategoryType"`         // Spend bonus category type
	SpendBonusCategoryName     string  `bson:"spend_bonus_category_name" json:"spendBonusCategoryName"`         // Spend bonus category name (e.g., Dining)
	SpendBonusCategoryID       int     `bson:"spend_bonus_category_id" json:"spendBonusCategoryID"`             // Unique Spend Bonus Category ID
	SpendBonusCategoryGroup    string  `bson:"spend_bonus_category_group" json:"spendBonusCategoryGroup"`       // Spend bonus category group (e.g., Dining)
	SpendBonusSubcategoryGroup string  `bson:"spend_bonus_subcategory_group" json:"spendBonusSubcategoryGroup"` // Spend bonus subcategory group (e.g., All Dining)
	SpendBonusDesc             string  `bson:"spend_bonus_desc" json:"spendBonusDesc"`                          // Spend bonus description
	EarnMultiplier             float64 `bson:"earn_multiplier" json:"earnMultiplier"`                           // Points/miles per dollar
	IsDateLimit                int     `bson:"is_date_limit" json:"isDateLimit"`                                // Is category date limited? (false=no, true=yes)
	LimitBeginDate             string  `bson:"limit_begin_date,omitempty" json:"limitBeginDate,omitempty"`      // Date spend bonus begins
	LimitEndDate               string  `bson:"limit_end_date,omitempty" json:"limitEndDate,omitempty"`          // Date spend bonus ends
	IsSpendLimit               int     `bson:"is_spend_limit" json:"isSpendLimit"`                              // Is there a spend limit? (false=no, true=yes)
	SpendLimit                 float64 `bson:"spend_limit" json:"spendLimit"`                                   // Spend limit amount if applicable
	SpendLimitResetPeriod      string  `bson:"spend_limit_reset_period" json:"spendLimitResetPeriod"`           // Period when the spend limit resets (e.g., Year)
}

// AnnualSpend represents an annual spend bonus of a credit card.
type AnnualSpend struct {
	AnnualSpendDesc string `bson:"annual_spend_desc" json:"annualSpendDesc"` // Description if card includes bonus for reaching an annual spend
}

// FetchCardDetail fetches detail from the API in format of CardDetail for the
// specified cardKey string and returns with any error.
func FetchCardDetail(cardKey string) (*CardDetail, error) {
	params := []string{cardKey}
	resp, err := FetchEndpoint("card_detail", params)
	if err != nil {
		return &CardDetail{}, err
	}
	defer resp.Body.Close()

	var cardDetails []CardDetail
	if err := json.NewDecoder(resp.Body).Decode(&cardDetails); err != nil {
		return &CardDetail{}, err
	}

	if len(cardDetails) == 0 {
		err := fmt.Errorf("no details found for cardKey: %s", cardKey)
		return &CardDetail{}, err
	}

	return &cardDetails[0], nil
}

// IsApplicable determines if the bonus applies based on the given merchant
// category.
func (bonus *SpendBonusCategory) IsApplicable(categoryID int) bool {
	return bonus.SpendBonusCategoryID == categoryID
}

// RewardValue gets the value of a reward for a card in dollars per dollar after
// accounting for point conversions to cash.
func (card *CardDetail) RewardValue(rewardAmount float64) float64 {
	cent, _ := decimal.NewFromString("0.01")
	mult := decimal.NewFromFloat(rewardAmount).Mul(cent)
	if card.BaseSpendEarnIsCash == 1 {
		// If can be directly converted to cash, use that multiplier
		cashValue := decimal.NewFromFloat(card.BaseSpendEarnCashValue)
		totalValue := cashValue.Mul(mult)
		return totalValue.InexactFloat64()
	}
	// Otherwise just return the points assuming each is worth one cent
	return mult.InexactFloat64()
}
