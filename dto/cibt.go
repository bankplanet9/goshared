package dto

import (
	"encoding/xml"
	"strconv"
)

type CIBTResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Header  struct {
		Text            string `xml:",chardata"`
		MessageResponse struct {
			Text  string `xml:",chardata"`
			GId   string `xml:"GId,attr"`
			MId   string `xml:"MId,attr"`
			MTs   string `xml:"MTs,attr"`
			METs  string `xml:"METs,attr"`
			Xmlns string `xml:"xmlns,attr"`
			P     struct {
				Text string `xml:",chardata"`
				SId  string `xml:"SId,attr"`
				PId  string `xml:"PId,attr"`
				PNs  string `xml:"PNs,attr"`
				R    struct {
					Text string `xml:",chardata"`
					L    string `xml:"L,attr"`
					C    string `xml:"C,attr"`
					D    string `xml:"D,attr"`
				} `xml:"R"`
			} `xml:"P"`
			Tx struct {
				Text string `xml:",chardata"`
				TxNs string `xml:"TxNs,attr"`
				R    struct {
					Text string `xml:",chardata"`
					L    string `xml:"L,attr"`
					C    string `xml:"C,attr"`
					D    string `xml:"D,attr"`
				} `xml:"R"`
			} `xml:"Tx"`
		} `xml:"MessageResponse"`
	} `xml:"Header"`
	Body struct {
		Text       string `xml:",chardata"`
		MGResponse struct {
			Text       string `xml:",chardata"`
			Xmlns      string `xml:"xmlns,attr"`
			MGResponse struct {
				Text        string `xml:",chardata"`
				Xmlns       string `xml:"xmlns,attr"`
				RIReqOutput struct {
					Text    string `xml:",chardata"`
					Xmlns   string `xml:"xmlns,attr"`
					Subject struct {
						Text     string `xml:",chardata"`
						Inquired struct {
							Text          string `xml:",chardata"`
							CBSubjectCode string `xml:"CBSubjectCode"`
							Person        struct {
								Text             string `xml:",chardata"`
								Gender           string `xml:"Gender"`
								FirstName        string `xml:"FirstName"`
								Patronymic       string `xml:"Patronymic"`
								Surname          string `xml:"Surname"`
								PrevFirstName    string `xml:"PrevFirstName"`
								PrevSurname      string `xml:"PrevSurname"`
								MotherMaidenName string `xml:"MotherMaidenName"`
								DateOfBirth      string `xml:"DateOfBirth"`
								PlaceOfBirth     string `xml:"PlaceOfBirth"`
								CountryOfBirth   string `xml:"CountryOfBirth"`
								INN              string `xml:"INN"`
								Documents        struct {
									Text string `xml:",chardata"`
									Main struct {
										Text               string `xml:",chardata"`
										DocType            string `xml:"DocType"`
										DocSeriesAndNumber string `xml:"DocSeriesAndNumber"`
										DocIssueDate       string `xml:"DocIssueDate"`
										DocIssueAuthPlace  string `xml:"DocIssueAuthPlace"`
									} `xml:"Main"`
								} `xml:"Documents"`
								Address struct {
									Text string `xml:",chardata"`
									Main struct {
										Text             string `xml:",chardata"`
										StreetNumAndName string `xml:"StreetNumAndName"`
										City             string `xml:"City"`
										District         string `xml:"District"`
										Region           string `xml:"Region"`
										PostCode         string `xml:"PostCode"`
										Country          string `xml:"Country"`
										LivedSinceDate   string `xml:"LivedSinceDate"`
									} `xml:"Main"`
								} `xml:"Address"`
								Contact struct {
									Text          string `xml:",chardata"`
									ContactType   string `xml:"ContactType"`
									ContactNumber string `xml:"ContactNumber"`
								} `xml:"Contact"`
								SoleTrade struct {
									Text              string `xml:",chardata"`
									SoleTradeName     string `xml:"SoleTradeName"`
									RegistrationNum   string `xml:"RegistrationNum"`
									EstablishDate     string `xml:"EstablishDate"`
									RegistrationPlace string `xml:"RegistrationPlace"`
									ContactNumber     string `xml:"ContactNumber"`
								} `xml:"SoleTrade"`
								Employement struct {
									Text            string `xml:",chardata"`
									EmployerName    string `xml:"EmployerName"`
									Status          string `xml:"Status"`
									ContactNumber   string `xml:"ContactNumber"`
									DateHired       string `xml:"DateHired"`
									DateTermination string `xml:"DateTermination"`
									OccupationType  string `xml:"OccupationType"`
									JobTitle        string `xml:"JobTitle"`
									AnnualIncome    string `xml:"AnnualIncome"`
								} `xml:"Employement"`
							} `xml:"Person"`
						} `xml:"Inquired"`
						Matched struct {
							Text          string `xml:",chardata"`
							CBSubjectCode string `xml:"CBSubjectCode"`
							FISubjectCode string `xml:"FISubjectCode"`
							FlgMatched    string `xml:"FlgMatched"`
							Person        struct {
								Text             string `xml:",chardata"`
								Gender           string `xml:"Gender"`
								FirstName        string `xml:"FirstName"`
								Patronymic       string `xml:"Patronymic"`
								Surname          string `xml:"Surname"`
								PrevFirstName    string `xml:"PrevFirstName"`
								PrevSurname      string `xml:"PrevSurname"`
								MotherMaidenName string `xml:"MotherMaidenName"`
								MaritalStatus    string `xml:"MaritalStatus"`
								NumDependents    string `xml:"NumDependents"`
								DateOfBirth      string `xml:"DateOfBirth"`
								PlaceOfBirth     string `xml:"PlaceOfBirth"`
								CountryOfBirth   string `xml:"CountryOfBirth"`
								INN              string `xml:"INN"`
								Address          struct {
									Text    string `xml:",chardata"`
									Current []struct {
										Text             string `xml:",chardata"`
										AddressType      string `xml:"AddressType"`
										StreetNumAndName string `xml:"StreetNumAndName"`
										City             string `xml:"City"`
										District         string `xml:"District"`
										Region           string `xml:"Region"`
										PostCode         string `xml:"PostCode"`
										Country          string `xml:"Country"`
										LivedSinceDate   string `xml:"LivedSinceDate"`
									} `xml:"Current"`
									Historical []struct {
										Text             string `xml:",chardata"`
										AddressType      string `xml:"AddressType"`
										StreetNumAndName string `xml:"StreetNumAndName"`
										City             string `xml:"City"`
										District         string `xml:"District"`
										Region           string `xml:"Region"`
										PostCode         string `xml:"PostCode"`
										Country          string `xml:"Country"`
										LivedSinceDate   string `xml:"LivedSinceDate"`
									} `xml:"Historical"`
								} `xml:"Address"`
								Documents struct {
									Text               string `xml:",chardata"`
									DocType            string `xml:"DocType"`
									DocSeriesAndNumber string `xml:"DocSeriesAndNumber"`
									DocIssueDate       string `xml:"DocIssueDate"`
									DocIssueAuthPlace  string `xml:"DocIssueAuthPlace"`
								} `xml:"Documents"`
								Contacts []struct {
									Text          string `xml:",chardata"`
									ContactType   string `xml:"ContactType"`
									ContactNumber string `xml:"ContactNumber"`
								} `xml:"Contacts"`
								Employements struct {
									Text            string `xml:",chardata"`
									EmpType         string `xml:"EmpType"`
									EmployerName    string `xml:"EmployerName"`
									Status          string `xml:"Status"`
									ContactNumber   string `xml:"ContactNumber"`
									DateHired       string `xml:"DateHired"`
									DateTermination string `xml:"DateTermination"`
									OccupationType  string `xml:"OccupationType"`
									JobTitle        string `xml:"JobTitle"`
									AnnualIncome    string `xml:"AnnualIncome"`
								} `xml:"Employements"`
							} `xml:"Person"`
						} `xml:"Matched"`
					} `xml:"Subject"`
					CreditHistory struct {
						Text              string `xml:",chardata"`
						InquiredOperation struct {
							Text           string `xml:",chardata"`
							CBContractCode string `xml:"CBContractCode"`
							ContractType   string `xml:"ContractType"`
							Phase          string `xml:"Phase"`
						} `xml:"InquiredOperation"`
						Score struct {
							Text         string `xml:",chardata"`
							Score        string `xml:"Score"`
							Range        string `xml:"Range"`
							ScoreMessage string `xml:"ScoreMessage"`
							ScoreFactor  []struct {
								Text        string `xml:",chardata"`
								Description string `xml:"Description"`
							} `xml:"ScoreFactor"`
						} `xml:"Score"`
						GeneralData struct {
							Text            string `xml:",chardata"`
							TotalContracts  string `xml:"TotalContracts"`
							TotalInstitutes string `xml:"TotalInstitutes"`
						} `xml:"GeneralData"`
						Contract struct {
							Text        string `xml:",chardata"`
							Instalments struct {
								Text    string `xml:",chardata"`
								Summary struct {
									Text           string `xml:",chardata"`
									TotalRequested string `xml:"TotalRequested"`
									TotalRefused   string `xml:"TotalRefused"`
									TotalRenounced string `xml:"TotalRenounced"`
									TotalActive    string `xml:"TotalActive"`
									TotalClosed    string `xml:"TotalClosed"`
								} `xml:"Summary"`
								ACSummary struct {
									Text                string `xml:",chardata"`
									TotInstallmentsAmn  string `xml:"TotInstallmentsAmn"`
									TotRemainInstallAmn string `xml:"TotRemainInstallAmn"`
									TotPastDueAmn       string `xml:"TotPastDueAmn"`
								} `xml:"ACSummary"`
								GSummary struct {
									Text                string `xml:",chardata"`
									TotInstallmentsAmn  string `xml:"TotInstallmentsAmn"`
									TotRemainInstallAmn string `xml:"TotRemainInstallAmn"`
									TotPastDueAmn       string `xml:"TotPastDueAmn"`
								} `xml:"GSummary"`
								GrantedContract []struct {
									Text       string `xml:",chardata"`
									CommonData struct {
										Text           string `xml:",chardata"`
										FICode         string `xml:"FICode"`
										CBContractCode string `xml:"CBContractCode"`
										ContractType   string `xml:"ContractType"`
										Phase          string `xml:"Phase"`
										StartDate      string `xml:"StartDate"`
										LastUpdateDate string `xml:"LastUpdateDate"`
										Currency       string `xml:"Currency"`
										CreditCurrency string `xml:"CreditCurrency"`
										Role           string `xml:"Role"`
										FIContractCode string `xml:"FIContractCode"`
									} `xml:"CommonData"`
									FlgReorg            string `xml:"FlgReorg"`
									DebtReason          string `xml:"DebtReason"`
									PaymentMethod       string `xml:"PaymentMethod"`
									Amount              string `xml:"Amount"`
									InstallmentNum      string `xml:"InstallmentNum"`
									Periodicity         string `xml:"Periodicity"`
									InstallmentAmn      string `xml:"InstallmentAmn"`
									EndDate             string `xml:"EndDate"`
									LastPaymentDate     string `xml:"LastPaymentDate"`
									NextInstallDate     string `xml:"NextInstallDate"`
									NextInstallAmn      string `xml:"NextInstallAmn"`
									PersonalGuarantees  string `xml:"PersonalGuarantees"`
									RealGuarantees      string `xml:"RealGuarantees"`
									Remarks             string `xml:"Remarks"`
									RemainInstallNum    string `xml:"RemainInstallNum"`
									RemainAmn           string `xml:"RemainAmn"`
									PastDueInstallNum   string `xml:"PastDueInstallNum"`
									PastDueInstallAmn   string `xml:"PastDueInstallAmn"`
									PastDueDaysNum      string `xml:"PastDueDaysNum"`
									MaxUnpaidAmn        string `xml:"MaxUnpaidAmn"`
									MaxDueInstallNum    string `xml:"MaxDueInstallNum"`
									MaxDueInstallDate   string `xml:"MaxDueInstallDate"`
									MaxMonthsDefaultNum string `xml:"MaxMonthsDefaultNum"`
									MaxPastDueDaysNum   string `xml:"MaxPastDueDaysNum"`
									MaxPastDueDaysDate  string `xml:"MaxPastDueDaysDate"`
									WorstStatus         string `xml:"WorstStatus"`
									WorstStatusDate     string `xml:"WorstStatusDate"`
									LeaseGoodType       string `xml:"LeaseGoodType"`
									LeaseGoodValue      string `xml:"LeaseGoodValue"`
									FlgNewUsed          string `xml:"FlgNewUsed"`
									LeaseGoodBrand      string `xml:"LeaseGoodBrand"`
									LeaseGoodNum        string `xml:"LeaseGoodNum"`
									LeaseGoodDate       string `xml:"LeaseGoodDate"`
									LinkedSubjects      []struct {
										Text              string `xml:",chardata"`
										CBSubjectCode     string `xml:"CBSubjectCode"`
										LinkedSubjectDesc string `xml:"LinkedSubjectDesc"`
										Role              string `xml:"Role"`
									} `xml:"LinkedSubjects"`
									History []struct {
										Text           string `xml:",chardata"`
										Year           string `xml:"Year"`
										Month          string `xml:"Month"`
										DueInstallNum  string `xml:"DueInstallNum"`
										PastDueDaysNum string `xml:"PastDueDaysNum"`
										Status         string `xml:"Status"`
									} `xml:"History"`
									Collaterals []struct {
										Text             string `xml:",chardata"`
										CBSubjectCode    string `xml:"CBSubjectCode"`
										GuarantorDesc    string `xml:"GuarantorDesc"`
										FICollateralCode string `xml:"FICollateralCode"`
										CountUsed        string `xml:"CountUsed"`
										CollateralType   string `xml:"CollateralType"`
										GuaranteedAmn    string `xml:"GuaranteedAmn"`
										Real             struct {
											Text             string `xml:",chardata"`
											AssetCode        string `xml:"AssetCode"`
											AssetDescription string `xml:"AssetDescription"`
											StartDate        string `xml:"StartDate"`
											EndDate          string `xml:"EndDate"`
										} `xml:"Real"`
										Personal struct {
											Text         string `xml:",chardata"`
											CustomerType string `xml:"CustomerType"`
											StartDate    string `xml:"StartDate"`
											EndDate      string `xml:"EndDate"`
										} `xml:"Personal"`
									} `xml:"Collaterals"`
								} `xml:"GrantedContract"`
							} `xml:"Instalments"`
							NonInstalments struct {
								Text    string `xml:",chardata"`
								Summary struct {
									Text           string `xml:",chardata"`
									TotalRequested string `xml:"TotalRequested"`
									TotalRefused   string `xml:"TotalRefused"`
									TotalRenounced string `xml:"TotalRenounced"`
									TotalActive    string `xml:"TotalActive"`
									TotalClosed    string `xml:"TotalClosed"`
								} `xml:"Summary"`
								ACSummary struct {
									Text           string `xml:",chardata"`
									TotCreditLimit string `xml:"TotCreditLimit"`
									TotUtilizAmn   string `xml:"TotUtilizAmn"`
									TotOverdrAmn   string `xml:"TotOverdrAmn"`
								} `xml:"ACSummary"`
								GSummary struct {
									Text           string `xml:",chardata"`
									TotCreditLimit string `xml:"TotCreditLimit"`
									TotUtilizAmn   string `xml:"TotUtilizAmn"`
									TotOverdrAmn   string `xml:"TotOverdrAmn"`
								} `xml:"GSummary"`
							} `xml:"NonInstalments"`
							CreditCards struct {
								Text    string `xml:",chardata"`
								Summary struct {
									Text           string `xml:",chardata"`
									TotalRequested string `xml:"TotalRequested"`
									TotalRefused   string `xml:"TotalRefused"`
									TotalRenounced string `xml:"TotalRenounced"`
									TotalActive    string `xml:"TotalActive"`
									TotalClosed    string `xml:"TotalClosed"`
								} `xml:"Summary"`
								ACSummary struct {
									Text           string `xml:",chardata"`
									TotCreditLimit string `xml:"TotCreditLimit"`
									TotRemainAmn   string `xml:"TotRemainAmn"`
									TotPastDueAmn  string `xml:"TotPastDueAmn"`
								} `xml:"ACSummary"`
								GSummary struct {
									Text           string `xml:",chardata"`
									TotCreditLimit string `xml:"TotCreditLimit"`
									TotRemainAmn   string `xml:"TotRemainAmn"`
									TotPastDueAmn  string `xml:"TotPastDueAmn"`
								} `xml:"GSummary"`
							} `xml:"CreditCards"`
						} `xml:"Contract"`
					} `xml:"CreditHistory"`
					ClaraResponse string `xml:"ClaraResponse"`
				} `xml:"RI_Req_Output"`
			} `xml:"MGResponse"`
		} `xml:"MGResponse"`
	} `xml:"Body"`
}

// Score returns the client's score in numeric representation (0 - 600)
func (c CIBTResponse) Score() int {
	score, _ := strconv.Atoi(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.Score.Score)
	return score
}

// Range returns the client's credit range of A,B,C...
func (c CIBTResponse) Range() string {
	return c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.Score.Range
}

func (c CIBTResponse) TotalActiveContracts() int {
	active, _ := strconv.Atoi(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.
		Contract.Instalments.Summary.TotalActive)
	return active
}

func (c CIBTResponse) TotalClosedContracts() int {
	closed, _ := strconv.Atoi(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.
		Contract.Instalments.Summary.TotalClosed)
	return closed
}

func (c CIBTResponse) TotalContracts() int {
	total, _ := strconv.Atoi(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.
		GeneralData.TotalContracts)
	return total
}

// TotalACRemainSum shows a total remain sum of installments for the client as A (borrower) and C (co-borrower)
func (c CIBTResponse) TotalACRemainSum() float64 {
	remain, _ := strconv.ParseFloat(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.Contract.Instalments.ACSummary.TotRemainInstallAmn, 64)
	return remain
}

// TotalGRemainSum shows a total remain sum of installments for the client as G (guarantor)
func (c CIBTResponse) TotalGRemainSum() float64 {
	remain, _ := strconv.ParseFloat(c.Body.MGResponse.MGResponse.RIReqOutput.CreditHistory.Contract.Instalments.GSummary.TotRemainInstallAmn, 64)
	return remain
}
