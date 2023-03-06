package dpfm_api_output_formatter

type DeliverySDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	APIType             string             `json:"api_type"`
	DataConcatenation   *DataConcatenation `json:"DataConcatenation"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	Deleted             bool               `json:"deleted"`
	APIProcessingResult bool               `json:"api_processing_result"`
	APIProcessingError  *string            `json:"api_processing_error"`
}

type DataConcatenation struct {
	Header DeliveryHeader `json:"DeliveryHeader"`
	Item   []DeliveryItem `json:"DeliveryItem"`
}

type DeliveryHeader struct {
	ExchangedDeliveryNoticeDocumentIdentifier                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier" csv:"1"`
	DeliveryNoticeDocument                                         *string  `json:"DeliveryNoticeDocument" csv:"2"`
	ExchangedDocumentContextSpecifiedTransactionIdentifier         *string  `json:"ExchangedDocumentContextSpecifiedTransactionIdentifier" csv:"3"`
	ExchangedDeliveryNoticeDocumentName                            *string  `json:"ExchangedDeliveryNoticeDocumentName" csv:"4"`
	ExchangeDeliveryNoticeDocumentTypeCode                         *string  `json:"ExchangeDeliveryNoticeDocumentTypeCode" csv:"5"`
	ExchangedDeliveryNoticeDocumentIssueDate                       *string  `json:"ExchangedDeliveryNoticeDocumentIssueDate" csv:"6"`
	ExchangedDeliveryNoticeDocumentPurposeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentPurposeCode" csv:"7"`
	ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode       *string  `json:"ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode" csv:"8"`
	ExchangedDeliveryNoticeDocumentVersionIdentifier               *string  `json:"ExchangedDeliveryNoticeDocumentVersionIdentifier" csv:"9"`
	ExchangedDeliveryNoticeDocumentCategoryCode                    *string  `json:"ExchangedDeliveryNoticeDocumentCategoryCode" csv:"10"`
	ExchangedDeliveryNoticeDocumentSubtypeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentSubtypeCode" csv:"11"`
	NoteDeliveryNoticeSubjectText                                  *string  `json:"NoteDeliveryNoticeSubjectText" csv:"12"`
	NoteDeliveryNoticeContentText                                  *string  `json:"NoteDeliveryNoticeContentText" csv:"13"`
	NoteDeliveryNoticeIdentifier                                   *string  `json:"NoteDeliveryNoticeIdentifier" csv:"14"`
	SpecifiedBinaryFileIdentifier                                  *string  `json:"SpecifiedBinaryFileIdentifier" csv:"15"`
	SpecifiedBinaryFileVersionIdentifier                           *string  `json:"SpecifiedBinaryFileVersionIdentifier" csv:"16"`
	SpecifiedBinaryFileNameText                                    *string  `json:"SpecifiedBinaryFileNameText" csv:"17"`
	SpecifiedBineryFileURIIdentifier                               *string  `json:"SpecifiedBineryFileURIIdentifier" csv:"18"`
	SpecifiedBineryFileMIMECode                                    *string  `json:"SpecifiedBineryFileMIMECode" csv:"19"`
	SpecifiedBineryFileEncodingCode                                *string  `json:"SpecifiedBineryFileEncodingCode" csv:"20"`
	SpecifiedBineryFileCode                                        *string  `json:"SpecifiedBineryFileCode" csv:"21"`
	SpecifiedBinaryFileDescriptionText                             *string  `json:"SpecifiedBinaryFileDescriptionText" csv:"22"`
	TradeSellerIdentifier                                          *string  `json:"TradeSellerIdentifier" csv:"23"`
	TradeSellerGlobalIdentifier                                    *string  `json:"TradeSellerGlobalIdentifier" csv:"24"`
	TradeSellerName                                                *string  `json:"TradeSellerName" csv:"25"`
	TradeBillFromPartyRegisteredIdentifier                         *string  `json:"TradeBillFromPartyRegisteredIdentifier" csv:"26"`
	TradeContactSellerIdentifier                                   *string  `json:"TradeContactSellerIdentifier" csv:"27"`
	TradeContactSellerPersonName                                   *string  `json:"TradeContactSellerPersonName" csv:"28"`
	TradeContactSellerDepartmentName                               *string  `json:"TradeContactSellerDepartmentName" csv:"29"`
	SellerTelephoneNumber                                          *string  `json:"SellerTelephoneNumber" csv:"30"`
	SellerFaxNumber                                                *string  `json:"SellerFaxNumber" csv:"31"`
	SellerEmailAddress                                             *string  `json:"SellerEmailAddress" csv:"32"`
	SellerAddressPostalCode                                        *string  `json:"SellerAddressPostalCode" csv:"33"`
	SellerAddress                                                  *string  `json:"SellerAddress" csv:"34"`
	TradeBuyerIdentifier                                           *string  `json:"TradeBuyerIdentifier" csv:"35"`
	TradeBuyerGlobalIdentifier                                     *string  `json:"TradeBuyerGlobalIdentifier" csv:"36"`
	TradeBuyerName                                                 *string  `json:"TradeBuyerName" csv:"37"`
	TradeContactBuyerIdentifier                                    *string  `json:"TradeContactBuyerIdentifier" csv:"38"`
	TradeContactBuyerPersonName                                    *string  `json:"TradeContactBuyerPersonName" csv:"39"`
	TradeContactBuyerDepartmentName                                *string  `json:"TradeContactBuyerDepartmentName" csv:"40"`
	BuyerTelephoneNumber                                           *string  `json:"BuyerTelephoneNumber" csv:"41"`
	BuyerFaxNumber                                                 *string  `json:"BuyerFaxNumber" csv:"42"`
	BuyerEmailAddress                                              *string  `json:"BuyerEmailAddress" csv:"43"`
	BuyerAddressPostalCode                                         *string  `json:"BuyerAddressPostalCode" csv:"44"`
	BuyerAddress                                                   *string  `json:"BuyerAddress" csv:"45"`
	ReferencedOrdersDocumentIssureAssignedIdentifier               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier" csv:"46"`
	ReferencedOrdersDocumentIssueDate                              *string  `json:"ReferencedOrdersDocumentIssueDate" csv:"47"`
	ReferencedOrdersDocumentRevisionIdentifier                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier" csv:"48"`
	ReferencedOrdersDocumentName                                   *string  `json:"ReferencedOrdersDocumentName" csv:"49"`
	ReferencedOrdersDocumentInformationText                        *string  `json:"ReferencedOrdersDocumentInformationText" csv:"50"`
	ReferencedOrdersDocumentInformationPurposeCode                 *string  `json:"ReferencedOrdersDocumentInformationPurposeCode" csv:"51"`
	ReferencedOrdersDocumentInformationSubtypeCode                 *string  `json:"ReferencedOrdersDocumentInformationSubtypeCode" csv:"52"`
	ProjectIdentifier                                              *string  `json:"ProjectIdentifier" csv:"53"`
	ProjectName                                                    *string  `json:"ProjectName" csv:"54"`
	ReferencedSalesOrderDocumentIssureAddignedIdentifier           *string  `json:"ReferencedSalesOrderDocumentIssureAddignedIdentifier" csv:"55"`
	ReferencedSalesOrderDocumentIssueDate                          *string  `json:"ReferencedSalesOrderDocumentIssueDate" csv:"56"`
	ReferencedSalesOrderDocumentRevisionIdentifier                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier" csv:"57"`
	ReferencedSalesOrderDocumentName                               *string  `json:"ReferencedSalesOrderDocumentName" csv:"58"`
	ReferencedSalesOrderDocumentInformationText                    *string  `json:"ReferencedSalesOrderDocumentInformationText" csv:"59"`
	ReferencedSalesOrderDocumentSubtypeCode                        *string  `json:"ReferencedSalesOrderDocumentSubtypeCode" csv:"60"`
	TradeShipToPartyIdentifier                                     *string  `json:"TradeShipToPartyIdentifier" csv:"61"`
	TradeShipToPartyGlobalIdentifier                               *string  `json:"TradeShipToPartyGlobalIdentifier" csv:"62"`
	TradeShipToPartyName                                           *string  `json:"TradeShipToPartyName" csv:"63"`
	TradeShipToPartyContactIdentifier                              *string  `json:"TradeShipToPartyContactIdentifier" csv:"64"`
	TradeShipToPartyContactPersonName                              *string  `json:"TradeShipToPartyContactPersonName" csv:"65"`
	TradeShipToPartyContactDepartmentName                          *string  `json:"TradeShipToPartyContactDepartmentName" csv:"66"`
	TradeShipToPartyContactPersonIdentifier                        *string  `json:"TradeShipToPartyContactPersonIdentifier" csv:"67"`
	ShipToPartyTelephoneNumber                                     *string  `json:"ShipToPartyTelephoneNumber" csv:"68"`
	ShipToPartyFaxNumber                                           *string  `json:"ShipToPartyFaxNumber" csv:"69"`
	ShipToPartyEmailAddress                                        *string  `json:"ShipToPartyEmailAddress" csv:"70"`
	ShipToPartyAddressPostalCode                                   *string  `json:"ShipToPartyAddressPostalCode" csv:"71"`
	ShipToPartyAddress                                             *string  `json:"ShipToPartyAddress" csv:"72"`
	TradeShipFromPartyIdentifier                                   *string  `json:"TradeShipFromPartyIdentifier" csv:"73"`
	TradeShipFromPartyName                                         *string  `json:"TradeShipFromPartyName" csv:"74"`
	TradeLogiName                                                  *string  `json:"TradeLogiName" csv:"75"`
	TradeLogiContactIdentifier                                     *string  `json:"TradeLogiContactIdentifier" csv:"76"`
	TradeLogiContactPersonName                                     *string  `json:"TradeLogiContactPersonName" csv:"77"`
	TradeLogiContactDepartmentName                                 *string  `json:"TradeLogiContactDepartmentName" csv:"78"`
	TradeLogiContactPersonIdentifier                               *string  `json:"TradeLogiContactPersonIdentifier" csv:"79"`
	LogiTelephoneNumber                                            *string  `json:"LogiTelephoneNumber" csv:"80"`
	SupplyChainEventIdentifier                                     *string  `json:"SupplyChainEventIdentifier" csv:"81"`
	InstructedTemperatureControlCode                               *string  `json:"InstructedTemperatureControlCode" csv:"82"`
	TradeTaxCalculatedAmount                                       *float32 `json:"TradeTaxCalculatedAmount" csv:"83"`
	TradeTaxTypeCode                                               *string  `json:"TradeTaxTypeCode" csv:"84"`
	TradeTaxBasisAmount                                            *float32 `json:"TradeTaxBasisAmount" csv:"85"`
	TradeTaxCategoryCode                                           *string  `json:"TradeTaxCategoryCode" csv:"86"`
	TradeTaxCategoryName                                           *string  `json:"TradeTaxCategoryName" csv:"87"`
	TradeTaxRatePercent                                            *float32 `json:"TradeTaxRatePercent" csv:"88"`
	TradeTaxGrandTotalAmount                                       *float32 `json:"TradeTaxGrandTotalAmount" csv:"89"`
	TradeTaxCalculationMethod                                      *string  `json:"TradeTaxCalculationMethod" csv:"90"`
	TradeSettlementMonetarySummationTotalTaxAmount                 *float32 `json:"TradeSettlementMonetarySummationTotalTaxAmount" csv:"91"`
	TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount" csv:"92"`
	TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount   *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount" csv:"93"`
	TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount  *float32 `json:"TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount" csv:"94"`
}

type DeliveryItem struct {
	ExchangedDeliveryNoticeDocumentIdentifier                                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier" csv:"1"`
	DeliveryNoticeDocumentItemlineIdentifier                                       string   `json:"DeliveryNoticeDocumentItemlineIdentifier" csv:"95"`
	DeliveryNoticeDocumentItemlineStatusCode                                       *string  `json:"DeliveryNoticeDocumentItemlineStatusCode" csv:"96"`
	DeliveryNoticeDocumentItemlineStatusReasonCode                                 *string  `json:"DeliveryNoticeDocumentItemlineStatusReasonCode" csv:"97"`
	NoteDeliveryNoticeItemSubjectText                                              *string  `json:"NoteDeliveryNoticeItemSubjectText" csv:"98"`
	NoteDeliveryNoticeItemContentText                                              *string  `json:"NoteDeliveryNoticeItemContentText" csv:"99"`
	NoteDeliveryNoticeItemIdentifier                                               *string  `json:"NoteDeliveryNoticeItemIdentifier" csv:"100"`
	ReferencedSalesOrderDocumentIssuerAssignedIdentifier                           *string  `json:"ReferencedSalesOrderDocumentIssuerAssignedIdentifier" csv:"101"`
	ReferencedSalesOrderDocumentItemLineIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentItemLineIdentifier" csv:"102"`
	ReferencedSalesOrderDocumentRevisionIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier" csv:"103"`
	ReferencedOrdersDocumentIssureAssignedIdentifier                               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier" csv:"104"`
	ReferencedOrdersDocumentItemLineIdentifier                                     *string  `json:"ReferencedOrdersDocumentItemLineIdentifier" csv:"105"`
	ReferencedOrdersDocumentRevisionIdentifier                                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier" csv:"106"`
	TradePriceTypeCode                                                             *string  `json:"TradePriceTypeCode" csv:"107"`
	TradeDeliveryPriceChargeAmount                                                 *float32 `json:"TradeDeliveryPriceChargeAmount" csv:"108"`
	TradePriceBasisQuantity                                                        *float32 `json:"TradePriceBasisQuantity" csv:"109"`
	TradePriceBasisUnitCode                                                        *string  `json:"TradePriceBasisUnitCode" csv:"110"`
	SupplyChainTradeDeliveryPackageQuantity                                        *float32 `json:"SupplyChainTradeDeliveryPackageQuantity" csv:"111"`
	SupplyChainTradeDeliveryProductUnitQuantity                                    *float32 `json:"SupplyChainTradeDeliveryProductUnitQuantity" csv:"112"`
	SupplyChainTradeDeliveryPerPackageUnitQuantity                                 *float32 `json:"SupplyChainTradeDeliveryPerPackageUnitQuantity" csv:"113"`
	SupplyChainTradeDeliveryDespatchedQuantity                                     *float32 `json:"SupplyChainTradeDeliveryDespatchedQuantity" csv:"114"`
	SupplyChainTradeDeliveryRequestedQuantity                                      *float32 `json:"SupplyChainTradeDeliveryRequestedQuantity" csv:"115"`
	SupplyChainTradeDeliveryRemainingRequestedQuantity                             *float32 `json:"SupplyChainTradeDeliveryRemainingRequestedQuantity" csv:"116"`
	ItemTradeDeliverToPartyIdentifier                                              *string  `json:"ItemTradeDeliverToPartyIdentifier" csv:"117"`
	ItemTradeDeliverToPartyGlobalIdentifier                                        *string  `json:"ItemTradeDeliverToPartyGlobalIdentifier" csv:"118"`
	ItemTradeDeliverToPartyName                                                    *string  `json:"ItemTradeDeliverToPartyName" csv:"119"`
	ItemTradeDeliverToPartyContactPersonName                                       *string  `json:"ItemTradeDeliverToPartyContactPersonName" csv:"120"`
	ItemTradeDeliverToPartyContactDepartmentName                                   *string  `json:"ItemTradeDeliverToPartyContactDepartmentName" csv:"121"`
	ItemDeliverToPartyPhoneNumber                                                  *string  `json:"ItemDeliverToPartyPhoneNumber" csv:"122"`
	ItemDeliverToPartyFaxNumber                                                    *string  `json:"ItemDeliverToPartyFaxNumber" csv:"123"`
	ItemDeliverToPartyEMailAddress                                                 *string  `json:"ItemDeliverToPartyEMailAddress" csv:"124"`
	ItemDeliverToPartyAddressPostalCode                                            *string  `json:"ItemDeliverToPartyAddressPostalCode" csv:"125"`
	ItemDeliverToPartyAddress                                                      *string  `json:"ItemDeliverToPartyAddress" csv:"126"`
	SupplyChainDeliveryEventIdentifier                                             *string  `json:"SupplyChainDeliveryEventIdentifier" csv:"127"`
	SupplyChainDeliveryEventOccurrenceDate                                         *string  `json:"SupplyChainDeliveryEventOccurrenceDate" csv:"128"`
	SupplyChainEventTypeCode                                                       *string  `json:"SupplyChainEventTypeCode" csv:"129"`
	SupplyChainEventRequirementOccurrenceDate                                      *string  `json:"SupplyChainEventRequirementOccurrenceDate" csv:"130"`
	LogisticsLocationIdentification                                                *string  `json:"LogisticsLocationIdentification" csv:"131"`
	LogisticsLocationName                                                          *string  `json:"LogisticsLocationName" csv:"132"`
	SupplyChainEventPlannedOccurrenceDate                                          *string  `json:"SupplyChainEventPlannedOccurrenceDate" csv:"133"`
	TradeTaxTypeCode                                                               *string  `json:"TradeTaxTypeCode" csv:"134"`
	ItemTradeTaxBasisAmount                                                        *float32 `json:"ItemTradeTaxBasisAmount" csv:"135"`
	ItemTradeTaxCategoryCode                                                       *string  `json:"ItemTradeTaxCategoryCode" csv:"136"`
	TradeTaxCategoryName                                                           *string  `json:"TradeTaxCategoryName" csv:"137"`
	ItemTradeTaxRateApplicablePercent                                              *float32 `json:"ItemTradeTaxRateApplicablePercent" csv:"138"`
	ItemTradeTaxGrandTotalAmount                                                   *float32 `json:"ItemTradeTaxGrandTotalAmount" csv:"139"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount               *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount" csv:"140"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount" csv:"141"`
	TradeProductIdentifier                                                         *string  `json:"TradeProductIdentifier" csv:"142"`
	TradeProductGlobalIdentifier                                                   *string  `json:"TradeProductGlobalIdentifier" csv:"143"`
	TradeProductSellerAssignedIdentifier                                           *string  `json:"TradeProductSellerAssignedIdentifier" csv:"144"`
	TradeProductBuyerAssignedIdentifier                                            *string  `json:"TradeProductBuyerAssignedIdentifier" csv:"145"`
	TradeProductManufacturerAssignedIdentifier                                     *string  `json:"TradeProductManufacturerAssignedIdentifier" csv:"146"`
	TradeProductName                                                               *string  `json:"TradeProductName" csv:"147"`
	TradeProductDescription                                                        *string  `json:"TradeProductDescription" csv:"148"`
	TradeProductTypeCode                                                           *string  `json:"TradeProductTypeCode" csv:"149"`
	TradeProductEndItemTypeCode                                                    *string  `json:"TradeProductEndItemTypeCode" csv:"150"`
	TradeProductSizeCode                                                           *string  `json:"TradeProductSizeCode" csv:"151"`
	TradeProductSizeDescriptionText                                                *string  `json:"TradeProductSizeDescriptionText" csv:"152"`
	TradeManufacturerIdentifier                                                    *string  `json:"TradeManufacturerIdentifier" csv:"153"`
	TradeManufacturerName                                                          *string  `json:"TradeManufacturerName" csv:"154"`
	ReferencedLogisticsPackageUnitQuantity                                         *float32 `json:"ReferencedLogisticsPackageUnitQuantity" csv:"155"`
	ReferencedLogisticsPackageQuantityUnitCode                                     *string  `json:"ReferencedLogisticsPackageQuantityUnitCode" csv:"156"`
	ReferencedLogisticsPackageTypeCode                                             *string  `json:"ReferencedLogisticsPackageTypeCode" csv:"157"`
	ReferencedLogisticsPackageIdentifier                                           *string  `json:"ReferencedLogisticsPackageIdentifier" csv:"158"`
}
