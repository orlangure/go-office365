package schema

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// AuditRecord represents an event or action returned by Audit endpoint.
type AuditRecord struct {
	ID             *string             `json:"Id"`
	RecordType     *AuditLogRecordType `json:"RecordType"`
	CreationTime   *string             `json:"CreationTime"`
	Operation      *string             `json:"Operation"`
	OrganizationID *string             `json:"OrganizationId"`
	UserType       *UserType           `json:"UserType"`
	UserKey        *string             `json:"UserKey"`
	Workload       *string             `json:"Workload,omitempty"`
	ResultStatus   *string             `json:"ResultStatus,omitempty"`
	ObjectID       *string             `json:"ObjectId,omitempty"`
	UserID         *string             `json:"UserId"`
	ClientIP       *string             `json:"ClientIP"`
	Scope          *AuditLogScope      `json:"Scope,omitempty"`
}

// AuditLogRecordType identifies the type of AuditRecord.
// https://docs.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-schema#enum-auditlogrecordtype---type-edmint32
type AuditLogRecordType int

// UnmarshalJSON unmarshals either a string or a int into an AuditLogRecordType.
func (t *AuditLogRecordType) UnmarshalJSON(b []byte) error {
	var raw json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	var i int
	if err := json.Unmarshal(raw, &i); err == nil {
		tmp := AuditLogRecordType(i)
		*t = tmp
		return nil
	}

	// if not int, it has to be a string
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return err
	}

	// if the string appears in the lookup table
	if tmp, err := GetRecordType(s); err == nil {
		*t = *tmp
		return nil
	}

	// it's probably a number represented as string: try to parse it
	if i, err := strconv.Atoi(s); err == nil {
		tmp := AuditLogRecordType(i)
		*t = tmp
		return nil
	}

	return fmt.Errorf("unexpected audit log record type %s", string(b))
}

// AuditLogRecordType enum.
const (
	ExchangeAdminType                         AuditLogRecordType = 1
	ExchangeItemType                          AuditLogRecordType = 2
	ExchangeItemGroupType                     AuditLogRecordType = 3
	SharePointType                            AuditLogRecordType = 4
	SharePointFileOperationType               AuditLogRecordType = 6
	OneDriveType                              AuditLogRecordType = 7
	AzureActiveDirectoryType                  AuditLogRecordType = 8
	AzureActiveDirectoryAccountLogonType      AuditLogRecordType = 9
	DataCenterSecurityCmdletType              AuditLogRecordType = 10
	ComplianceDLPSharePointType               AuditLogRecordType = 11
	ComplianceDLPExchangeType                 AuditLogRecordType = 13
	SharePointSharingOperationType            AuditLogRecordType = 14
	AzureActiveDirectoryStsLogonType          AuditLogRecordType = 15
	SkypeForBusinessPSTNUsageType             AuditLogRecordType = 16
	SkypeForBusinessUsersBlockedType          AuditLogRecordType = 17
	SecurityComplianceCenterEOPCmdletType     AuditLogRecordType = 18
	ExchangeAggregatedOperationType           AuditLogRecordType = 19
	PowerBIAuditType                          AuditLogRecordType = 20
	CRMType                                   AuditLogRecordType = 21
	YammerType                                AuditLogRecordType = 22
	SkypeForBusinessCmdletsType               AuditLogRecordType = 23
	DiscoveryType                             AuditLogRecordType = 24
	MicrosoftTeamsType                        AuditLogRecordType = 25
	ThreatIntelligenceType                    AuditLogRecordType = 28
	MailSubmissionType                        AuditLogRecordType = 29
	MicrosoftFlowType                         AuditLogRecordType = 30
	AeDType                                   AuditLogRecordType = 31
	MicrosoftStreamType                       AuditLogRecordType = 32
	ComplianceDLPSharePointClassificationType AuditLogRecordType = 33
	ThreatFinderType                          AuditLogRecordType = 34
	ProjectType                               AuditLogRecordType = 35
	SharePointListOperationType               AuditLogRecordType = 36
	SharePointCommentOperationType            AuditLogRecordType = 37
	DataGovernanceType                        AuditLogRecordType = 38
	KaizalaType                               AuditLogRecordType = 39
	SecurityComplianceAlertsType              AuditLogRecordType = 40
	ThreatIntelligenceUrlType                 AuditLogRecordType = 41
	SecurityComplianceInsightsType            AuditLogRecordType = 42
	MIPLabelType                              AuditLogRecordType = 43
	WorkplaceAnalyticsType                    AuditLogRecordType = 44
	PowerAppsAppType                          AuditLogRecordType = 45
	PowerAppsPlanType                         AuditLogRecordType = 46
	ThreatIntelligenceAtpContentType          AuditLogRecordType = 47
	LabelContentExplorerType                  AuditLogRecordType = 48
	TeamsHealthcareType                       AuditLogRecordType = 49
	ExchangeItemAggregatedType                AuditLogRecordType = 50
	HygieneEventType                          AuditLogRecordType = 51
	DataInsightsRestApiAuditType              AuditLogRecordType = 52
	InformationBarrierPolicyApplicationType   AuditLogRecordType = 53
	SharePointListItemOperationType           AuditLogRecordType = 54
	SharePointContentTypeOperationType        AuditLogRecordType = 55
	SharePointFieldOperationType              AuditLogRecordType = 56
	MicrosoftTeamsAdminType                   AuditLogRecordType = 57
	HRSignalType                              AuditLogRecordType = 58
	MicrosoftTeamsDeviceType                  AuditLogRecordType = 59
	MicrosoftTeamsAnalyticsType               AuditLogRecordType = 60
	InformationWorkerProtectionType           AuditLogRecordType = 61
	CampaignType                              AuditLogRecordType = 62
	DLPEndpointType                           AuditLogRecordType = 63
	AirInvestigationType                      AuditLogRecordType = 64
	QuarantineType                            AuditLogRecordType = 65
	MicrosoftFormsType                        AuditLogRecordType = 66
	ApplicationAuditType                      AuditLogRecordType = 67
	ComplianceSupervisionExchangeType         AuditLogRecordType = 68
	CustomerKeyServiceEncryptionType          AuditLogRecordType = 69
	OfficeNativeType                          AuditLogRecordType = 70
	MipAutoLabelSharePointItemType            AuditLogRecordType = 71
	MipAutoLabelSharePointPolicyLocationType  AuditLogRecordType = 72
	MicrosoftTeamsShiftsType                  AuditLogRecordType = 73
	MipAutoLabelExchangeItemType              AuditLogRecordType = 75
	CortanaBriefingType                       AuditLogRecordType = 76
	WDATPAlertsType                           AuditLogRecordType = 78
	SensitivityLabelPolicyMatchType           AuditLogRecordType = 82
	SensitivityLabelActionType                AuditLogRecordType = 83
	SensitivityLabeledFileActionType          AuditLogRecordType = 84
	AttackSimType                             AuditLogRecordType = 85
	AirManualInvestigationType                AuditLogRecordType = 86
	SecurityComplianceRBACType                AuditLogRecordType = 87
	UserTrainingType                          AuditLogRecordType = 88
	AirAdminActionInvestigationType           AuditLogRecordType = 89
	MSTICType                                 AuditLogRecordType = 90
	PhysicalBadgingSignalType                 AuditLogRecordType = 91
	AipDiscoverType                           AuditLogRecordType = 93
	AipSensitivityLabelActionType             AuditLogRecordType = 94
	AipProtectionActionType                   AuditLogRecordType = 95
	AipFileDeletedType                        AuditLogRecordType = 96
	AipHeartBeatType                          AuditLogRecordType = 97
	MCASAlertsType                            AuditLogRecordType = 98
	OnPremisesFileShareScannerDlpType         AuditLogRecordType = 99
	OnPremisesSharePointScannerDlpType        AuditLogRecordType = 100
	ExchangeSearchType                        AuditLogRecordType = 101
	SharePointSearchType                      AuditLogRecordType = 102
	PrivacyInsightsType                       AuditLogRecordType = 103
	MyAnalyticsSettingsType                   AuditLogRecordType = 105
	SecurityComplianceUserChangeType          AuditLogRecordType = 106
	ComplianceDLPExchangeClassificationType   AuditLogRecordType = 107
	MipExactDataMatchType                     AuditLogRecordType = 109
	MS365DCustomDetectionType                 AuditLogRecordType = 113
	CoreReportingSettingsType                 AuditLogRecordType = 147
	ComplianceConnectorType                   AuditLogRecordType = 148
	OMEPortalType                             AuditLogRecordType = 154
	DataShareOperationType                    AuditLogRecordType = 174
	EduDataLakeDownloadOperationType          AuditLogRecordType = 181
	MicrosoftGraphDataConnectOperationType    AuditLogRecordType = 183
	PowerPagesSiteType                        AuditLogRecordType = 186
	PlannerPlanType                           AuditLogRecordType = 188
	PlannerCopyPlanType                       AuditLogRecordType = 189
	PlannerTaskType                           AuditLogRecordType = 190
	PlannerRosterType                         AuditLogRecordType = 191
	PlannerPlanListType                       AuditLogRecordType = 192
	PlannerTaskListType                       AuditLogRecordType = 193
	PlannerTenantSettingsType                 AuditLogRecordType = 194
	ProjectForTheWebProjectType               AuditLogRecordType = 195
	ProjectForTheWebTaskType                  AuditLogRecordType = 196
	ProjectForTheWebRoadmapType               AuditLogRecordType = 197
	ProjectForTheWebRoadmapItemType           AuditLogRecordType = 198
	ProjectForTheWebProjectSettingsType       AuditLogRecordType = 199
	ProjectForTheWebRoadmapSettingsType       AuditLogRecordType = 200
	MicrosoftTodoAuditType                    AuditLogRecordType = 202
	VivaGoalsType                             AuditLogRecordType = 216
	MicrosoftGraphDataConnectConsentType      AuditLogRecordType = 217
	AttackSimAdminType                        AuditLogRecordType = 218
	TeamsUpdatesType                          AuditLogRecordType = 230
	DefenderExpertsforXDRAdminType            AuditLogRecordType = 237
	PlannerRosterSensitivityLabelType         AuditLogRecordType = 231
	VfamCreatePolicyType                      AuditLogRecordType = 251
	VfamUpdatePolicyType                      AuditLogRecordType = 252
	VfamDeletePolicyType                      AuditLogRecordType = 253
	CopilotInteractionType                    AuditLogRecordType = 261
)

func (t AuditLogRecordType) String() string {
	literals := map[AuditLogRecordType]string{
		ExchangeAdminType:                         "ExchangeAdmin",
		ExchangeItemType:                          "ExchangeItem",
		ExchangeItemGroupType:                     "ExchangeItemGroup",
		SharePointType:                            "SharePoint",
		SharePointFileOperationType:               "SharePointFileOperation",
		OneDriveType:                              "OneDrive",
		AzureActiveDirectoryType:                  "AzureActiveDirectory",
		AzureActiveDirectoryAccountLogonType:      "AzureActiveDirectoryAccountLogon",
		DataCenterSecurityCmdletType:              "DataCenterSecurityCmdlet",
		ComplianceDLPSharePointType:               "ComplianceDLPSharePoint",
		ComplianceDLPExchangeType:                 "ComplianceDLPExchange",
		SharePointSharingOperationType:            "SharePointSharingOperation",
		AzureActiveDirectoryStsLogonType:          "AzureActiveDirectoryStsLogon",
		SkypeForBusinessPSTNUsageType:             "SkypeForBusinessPSTNUsage",
		SkypeForBusinessUsersBlockedType:          "SkypeForBusinessUsersBlocked",
		SecurityComplianceCenterEOPCmdletType:     "SecurityComplianceCenterEOPCmdlet",
		ExchangeAggregatedOperationType:           "ExchangeAggregatedOperation",
		PowerBIAuditType:                          "PowerBIAudit",
		CRMType:                                   "CRM",
		YammerType:                                "Yammer",
		SkypeForBusinessCmdletsType:               "SkypeForBusinessCmdlets",
		DiscoveryType:                             "Discovery",
		MicrosoftTeamsType:                        "MicrosoftTeams",
		ThreatIntelligenceType:                    "ThreatIntelligence",
		MailSubmissionType:                        "MailSubmission",
		MicrosoftFlowType:                         "MicrosoftFlow",
		AeDType:                                   "AeD",
		MicrosoftStreamType:                       "MicrosoftStream",
		ComplianceDLPSharePointClassificationType: "ComplianceDLPSharePointClassification",
		ThreatFinderType:                          "ThreatFinder",
		ProjectType:                               "Project",
		SharePointListOperationType:               "SharePointListOperation",
		SharePointCommentOperationType:            "SharePointCommentOperation",
		DataGovernanceType:                        "DataGovernance",
		KaizalaType:                               "Kaizala",
		SecurityComplianceAlertsType:              "SecurityComplianceAlerts",
		ThreatIntelligenceUrlType:                 "ThreatIntelligenceUrl",
		SecurityComplianceInsightsType:            "SecurityComplianceInsights",
		MIPLabelType:                              "MIPLabel",
		WorkplaceAnalyticsType:                    "WorkplaceAnalytics",
		PowerAppsAppType:                          "PowerAppsApp",
		PowerAppsPlanType:                         "PowerAppsPlan",
		ThreatIntelligenceAtpContentType:          "ThreatIntelligenceAtpContent",
		LabelContentExplorerType:                  "LabelContentExplorer",
		TeamsHealthcareType:                       "TeamsHealthcare",
		ExchangeItemAggregatedType:                "ExchangeItemAggregated",
		HygieneEventType:                          "HygieneEvent",
		DataInsightsRestApiAuditType:              "DataInsightsRestApiAudit",
		InformationBarrierPolicyApplicationType:   "InformationBarrierPolicyApplication",
		SharePointListItemOperationType:           "SharePointListItemOperation",
		SharePointContentTypeOperationType:        "SharePointContentTypeOperation",
		SharePointFieldOperationType:              "SharePointFieldOperation",
		MicrosoftTeamsAdminType:                   "MicrosoftTeamsAdmin",
		HRSignalType:                              "HRSignal",
		MicrosoftTeamsDeviceType:                  "MicrosoftTeamsDevice",
		MicrosoftTeamsAnalyticsType:               "MicrosoftTeamsAnalytics",
		InformationWorkerProtectionType:           "InformationWorkerProtection",
		CampaignType:                              "Campaign",
		DLPEndpointType:                           "DLPEndpoint",
		AirInvestigationType:                      "AirInvestigation",
		QuarantineType:                            "Quarantine",
		MicrosoftFormsType:                        "MicrosoftForms",
		ApplicationAuditType:                      "ApplicationAudit",
		ComplianceSupervisionExchangeType:         "ComplianceSupervisionExchange",
		CustomerKeyServiceEncryptionType:          "CustomerKeyServiceEncryption",
		OfficeNativeType:                          "OfficeNative",
		MipAutoLabelSharePointItemType:            "MipAutoLabelSharePointItem",
		MipAutoLabelSharePointPolicyLocationType:  "MipAutoLabelSharePointPolicyLocation",
		MicrosoftTeamsShiftsType:                  "MicrosoftTeamsShifts",
		MipAutoLabelExchangeItemType:              "MipAutoLabelExchangeItem",
		CortanaBriefingType:                       "CortanaBriefing",
		WDATPAlertsType:                           "WDATPAlerts",
		SensitivityLabelPolicyMatchType:           "SensitivityLabelPolicyMatch",
		SensitivityLabelActionType:                "SensitivityLabelAction",
		SensitivityLabeledFileActionType:          "SensitivityLabeledFileAction",
		AttackSimType:                             "AttackSim",
		AirManualInvestigationType:                "AirManualInvestigation",
		SecurityComplianceRBACType:                "SecurityComplianceRBAC",
		UserTrainingType:                          "UserTraining",
		AirAdminActionInvestigationType:           "AirAdminActionInvestigation",
		MSTICType:                                 "MSTIC",
		PhysicalBadgingSignalType:                 "PhysicalBadgingSignal",
		AipDiscoverType:                           "AipDiscover",
		AipSensitivityLabelActionType:             "AipSensitivityLabelAction",
		AipProtectionActionType:                   "AipProtectionAction",
		AipFileDeletedType:                        "AipFileDeleted",
		AipHeartBeatType:                          "AipHeartBeat",
		MCASAlertsType:                            "MCASAlerts",
		OnPremisesFileShareScannerDlpType:         "OnPremisesFileShareScannerDlp",
		OnPremisesSharePointScannerDlpType:        "OnPremisesSharePointScannerDlp",
		ExchangeSearchType:                        "ExchangeSearch",
		SharePointSearchType:                      "SharePointSearch",
		PrivacyInsightsType:                       "PrivacyInsights",
		MyAnalyticsSettingsType:                   "MyAnalyticsSettings",
		SecurityComplianceUserChangeType:          "SecurityComplianceUserChange",
		ComplianceDLPExchangeClassificationType:   "ComplianceDLPExchangeClassification",
		MipExactDataMatchType:                     "MipExactDataMatch",
		MS365DCustomDetectionType:                 "MS365DCustomDetection",
		CoreReportingSettingsType:                 "CoreReportingSettings",
		ComplianceConnectorType:                   "ComplianceConnector",
		OMEPortalType:                             "OMEPortal",
		DataShareOperationType:                    "DataShareOperation",
		EduDataLakeDownloadOperationType:          "EduDataLakeDownloadOperation",
		MicrosoftGraphDataConnectOperationType:    "MicrosoftGraphDataConnectOperation",
		PowerPagesSiteType:                        "PowerPagesSite",
		PlannerPlanType:                           "PlannerPlan",
		PlannerCopyPlanType:                       "PlannerCopyPlan",
		PlannerTaskType:                           "PlannerTask",
		PlannerRosterType:                         "PlannerRoster",
		PlannerPlanListType:                       "PlannerPlanList",
		PlannerTaskListType:                       "PlannerTaskList",
		PlannerTenantSettingsType:                 "PlannerTenantSettings",
		ProjectForTheWebProjectType:               "ProjectForTheWebProject",
		ProjectForTheWebTaskType:                  "ProjectForTheWebTask",
		ProjectForTheWebRoadmapType:               "ProjectForTheWebRoadmap",
		ProjectForTheWebRoadmapItemType:           "ProjectForTheWebRoadmapItem",
		ProjectForTheWebProjectSettingsType:       "ProjectForTheWebProjectSettings",
		ProjectForTheWebRoadmapSettingsType:       "ProjectForTheWebRoadmapSettings",
		MicrosoftTodoAuditType:                    "MicrosoftTodoAudit",
		VivaGoalsType:                             "VivaGoals",
		MicrosoftGraphDataConnectConsentType:      "MicrosoftGraphDataConnectConsent",
		AttackSimAdminType:                        "AttackSimAdmin",
		TeamsUpdatesType:                          "TeamsUpdates",
		DefenderExpertsforXDRAdminType:            "DefenderExpertsforXDRAdmin",
		PlannerRosterSensitivityLabelType:         "PlannerRosterSensitivityLabel",
		VfamCreatePolicyType:                      "VfamCreatePolicy",
		VfamUpdatePolicyType:                      "VfamUpdatePolicy",
		VfamDeletePolicyType:                      "VfamDeletePolicy",
		CopilotInteractionType:                    "CopilotInteraction",
	}
	return literals[t]
}

var literals = map[string]AuditLogRecordType{
	"ExchangeAdmin":                         ExchangeAdminType,
	"ExchangeItem":                          ExchangeItemType,
	"ExchangeItemGroup":                     ExchangeItemGroupType,
	"SharePoint":                            SharePointType,
	"SharePointFileOperation":               SharePointFileOperationType,
	"OneDrive":                              OneDriveType,
	"AzureActiveDirectory":                  AzureActiveDirectoryType,
	"AzureActiveDirectoryAccountLogon":      AzureActiveDirectoryAccountLogonType,
	"DataCenterSecurityCmdlet":              DataCenterSecurityCmdletType,
	"ComplianceDLPSharePoint":               ComplianceDLPSharePointType,
	"ComplianceDLPExchange":                 ComplianceDLPExchangeType,
	"SharePointSharingOperation":            SharePointSharingOperationType,
	"AzureActiveDirectoryStsLogon":          AzureActiveDirectoryStsLogonType,
	"SkypeForBusinessPSTNUsage":             SkypeForBusinessPSTNUsageType,
	"SkypeForBusinessUsersBlocked":          SkypeForBusinessUsersBlockedType,
	"SecurityComplianceCenterEOPCmdlet":     SecurityComplianceCenterEOPCmdletType,
	"ExchangeAggregatedOperation":           ExchangeAggregatedOperationType,
	"PowerBIAudit":                          PowerBIAuditType,
	"CRM":                                   CRMType,
	"Yammer":                                YammerType,
	"SkypeForBusinessCmdlets":               SkypeForBusinessCmdletsType,
	"Discovery":                             DiscoveryType,
	"MicrosoftTeams":                        MicrosoftTeamsType,
	"ThreatIntelligence":                    ThreatIntelligenceType,
	"MailSubmission":                        MailSubmissionType,
	"MicrosoftFlow":                         MicrosoftFlowType,
	"AeD":                                   AeDType,
	"MicrosoftStream":                       MicrosoftStreamType,
	"ComplianceDLPSharePointClassification": ComplianceDLPSharePointClassificationType,
	"ThreatFinder":                          ThreatFinderType,
	"Project":                               ProjectType,
	"SharePointListOperation":               SharePointListOperationType,
	"SharePointCommentOperation":            SharePointCommentOperationType,
	"DataGovernance":                        DataGovernanceType,
	"Kaizala":                               KaizalaType,
	"SecurityComplianceAlerts":              SecurityComplianceAlertsType,
	"ThreatIntelligenceUrl":                 ThreatIntelligenceUrlType,
	"SecurityComplianceInsights":            SecurityComplianceInsightsType,
	"MIPLabel":                              MIPLabelType,
	"WorkplaceAnalytics":                    WorkplaceAnalyticsType,
	"PowerAppsApp":                          PowerAppsAppType,
	"PowerAppsPlan":                         PowerAppsPlanType,
	"ThreatIntelligenceAtpContent":          ThreatIntelligenceAtpContentType,
	"LabelContentExplorer":                  LabelContentExplorerType,
	"TeamsHealthcare":                       TeamsHealthcareType,
	"ExchangeItemAggregated":                ExchangeItemAggregatedType,
	"HygieneEvent":                          HygieneEventType,
	"DataInsightsRestApiAudit":              DataInsightsRestApiAuditType,
	"InformationBarrierPolicyApplication":   InformationBarrierPolicyApplicationType,
	"SharePointListItemOperation":           SharePointListItemOperationType,
	"SharePointContentTypeOperation":        SharePointContentTypeOperationType,
	"SharePointFieldOperation":              SharePointFieldOperationType,
	"MicrosoftTeamsAdmin":                   MicrosoftTeamsAdminType,
	"HRSignal":                              HRSignalType,
	"MicrosoftTeamsDevice":                  MicrosoftTeamsDeviceType,
	"MicrosoftTeamsAnalytics":               MicrosoftTeamsAnalyticsType,
	"InformationWorkerProtection":           InformationWorkerProtectionType,
	"Campaign":                              CampaignType,
	"DLPEndpoint":                           DLPEndpointType,
	"AirInvestigation":                      AirInvestigationType,
	"Quarantine":                            QuarantineType,
	"MicrosoftForms":                        MicrosoftFormsType,
	"ApplicationAudit":                      ApplicationAuditType,
	"ComplianceSupervisionExchange":         ComplianceSupervisionExchangeType,
	"CustomerKeyServiceEncryption":          CustomerKeyServiceEncryptionType,
	"OfficeNative":                          OfficeNativeType,
	"MipAutoLabelSharePointItem":            MipAutoLabelSharePointItemType,
	"MipAutoLabelSharePointPolicyLocation":  MipAutoLabelSharePointPolicyLocationType,
	"MicrosoftTeamsShifts":                  MicrosoftTeamsShiftsType,
	"MipAutoLabelExchangeItem":              MipAutoLabelExchangeItemType,
	"CortanaBriefing":                       CortanaBriefingType,
	"WDATPAlerts":                           WDATPAlertsType,
	"SensitivityLabelPolicyMatch":           SensitivityLabelPolicyMatchType,
	"SensitivityLabelAction":                SensitivityLabelActionType,
	"SensitivityLabeledFileAction":          SensitivityLabeledFileActionType,
	"AttackSim":                             AttackSimType,
	"AirManualInvestigation":                AirManualInvestigationType,
	"SecurityComplianceRBAC":                SecurityComplianceRBACType,
	"UserTraining":                          UserTrainingType,
	"AirAdminActionInvestigation":           AirAdminActionInvestigationType,
	"MSTIC":                                 MSTICType,
	"PhysicalBadgingSignal":                 PhysicalBadgingSignalType,
	"AipDiscover":                           AipDiscoverType,
	"AipSensitivityLabelAction":             AipSensitivityLabelActionType,
	"AipProtectionAction":                   AipProtectionActionType,
	"AipFileDeleted":                        AipFileDeletedType,
	"AipHeartBeat":                          AipHeartBeatType,
	"MCASAlerts":                            MCASAlertsType,
	"OnPremisesFileShareScannerDlp":         OnPremisesFileShareScannerDlpType,
	"OnPremisesSharePointScannerDlp":        OnPremisesSharePointScannerDlpType,
	"ExchangeSearch":                        ExchangeSearchType,
	"SharePointSearch":                      SharePointSearchType,
	"PrivacyInsights":                       PrivacyInsightsType,
	"MyAnalyticsSettings":                   MyAnalyticsSettingsType,
	"SecurityComplianceUserChange":          SecurityComplianceUserChangeType,
	"ComplianceDLPExchangeClassification":   ComplianceDLPExchangeClassificationType,
	"MipExactDataMatch":                     MipExactDataMatchType,
	"MS365DCustomDetection":                 MS365DCustomDetectionType,
	"CoreReportingSettings":                 CoreReportingSettingsType,
	"ComplianceConnector":                   ComplianceConnectorType,
	"OMEPortal":                             OMEPortalType,
	"DataShareOperation":                    DataShareOperationType,
	"EduDataLakeDownloadOperation":          EduDataLakeDownloadOperationType,
	"MicrosoftGraphDataConnectOperation":    MicrosoftGraphDataConnectOperationType,
	"PowerPagesSite":                        PowerPagesSiteType,
	"PlannerPlan":                           PlannerPlanType,
	"PlannerCopyPlan":                       PlannerCopyPlanType,
	"PlannerTask":                           PlannerTaskType,
	"PlannerRoster":                         PlannerRosterType,
	"PlannerPlanList":                       PlannerPlanListType,
	"PlannerTaskList":                       PlannerTaskListType,
	"PlannerTenantSettings":                 PlannerTenantSettingsType,
	"ProjectForTheWebProject":               ProjectForTheWebProjectType,
	"ProjectForTheWebTask":                  ProjectForTheWebTaskType,
	"ProjectForTheWebRoadmap":               ProjectForTheWebRoadmapType,
	"ProjectForTheWebRoadmapItem":           ProjectForTheWebRoadmapItemType,
	"ProjectForTheWebProjectSettings":       ProjectForTheWebProjectSettingsType,
	"ProjectForTheWebRoadmapSettings":       ProjectForTheWebRoadmapSettingsType,
	"MicrosoftTodoAudit":                    MicrosoftTodoAuditType,
	"VivaGoals":                             VivaGoalsType,
	"MicrosoftGraphDataConnectConsent":      MicrosoftGraphDataConnectConsentType,
	"AttackSimAdmin":                        AttackSimAdminType,
	"TeamsUpdates":                          TeamsUpdatesType,
	"DefenderExpertsforXDRAdmin":            DefenderExpertsforXDRAdminType,
	"PlannerRosterSensitivityLabel":         PlannerRosterSensitivityLabelType,
	"VfamCreatePolicy":                      VfamCreatePolicyType,
	"VfamUpdatePolicy":                      VfamUpdatePolicyType,
	"VfamDeletePolicy":                      VfamDeletePolicyType,
	"CopilotInteraction":                    CopilotInteractionType,
}

// GetRecordType returns the RecordType for the provided string.
func GetRecordType(s string) (*AuditLogRecordType, error) {
	t, ok := literals[s]
	if !ok {
		return nil, fmt.Errorf("record type %s invalid", s)
	}
	return &t, nil
}

// UserType identifies the type of user in AuditRecord.
// https://docs.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-schema#enum-user-type---type-edmint32
type UserType int

// UserType enum.
const (
	Regular UserType = iota
	Reserved
	Admin
	DcAdmin
	System
	Application
	ServicePrincipal
	CustomPolicy
	SystemPolicy
)

func (t UserType) String() string {
	literals := map[UserType]string{
		Regular:          "Regular",
		Reserved:         "Reserved",
		Admin:            "Admin",
		DcAdmin:          "DcAdmin",
		System:           "System",
		Application:      "Application",
		ServicePrincipal: "ServicePrincipal",
		CustomPolicy:     "CustomPolicy",
		SystemPolicy:     "SystemPolicy",
	}
	return literals[t]
}

// AuditLogScope identifies the scope of an AuditRecord.
// https://docs.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-schema#auditlogscope
type AuditLogScope int

// AuditLogScope enum.
const (
	Online AuditLogScope = iota
	Onprem
)

func (s AuditLogScope) String() string {
	literals := map[AuditLogScope]string{
		Online: "Online",
		Onprem: "Onprem",
	}
	return literals[s]
}

// ContentType represents a type and source of aggregated actions and events
// generated by the Microsoft Office 365 Management Activity API.
type ContentType int

// ContentType enum.
const (
	AuditAzureActiveDirectory ContentType = iota
	AuditExchange
	AuditSharePoint
	AuditGeneral
	DLPAll
)

func (c ContentType) String() string {
	literals := map[ContentType]string{
		AuditAzureActiveDirectory: "Audit.AzureActiveDirectory",
		AuditExchange:             "Audit.Exchange",
		AuditSharePoint:           "Audit.SharePoint",
		AuditGeneral:              "Audit.General",
		DLPAll:                    "DLP.All",
	}
	return literals[c]
}

var contentTypes = map[string]ContentType{
	"Audit.AzureActiveDirectory": AuditAzureActiveDirectory,
	"Audit.Exchange":             AuditExchange,
	"Audit.SharePoint":           AuditSharePoint,
	"Audit.General":              AuditGeneral,
	"DLP.All":                    DLPAll,
}

// GetContentType returns the ContentType represented
// by the provided string literal.
func GetContentType(s string) (*ContentType, error) {
	if v, ok := contentTypes[s]; ok {
		return &v, nil
	}
	return nil, fmt.Errorf("ContentType invalid")
}

// GetContentTypes returns the list of ContentType.
func GetContentTypes() []ContentType {
	var result []ContentType
	for _, t := range contentTypes {
		result = append(result, t)
	}
	return result
}

// ContentTypeValid validates that a string is a valid ContentType.
func ContentTypeValid(s string) bool {
	if _, err := GetContentType(s); err != nil {
		return false
	}
	return true
}
