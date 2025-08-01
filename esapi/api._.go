// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 9.1.0 (9732617): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
type API struct {
	Cat         *Cat
	Cluster     *Cluster
	Indices     *Indices
	Ingest      *Ingest
	Nodes       *Nodes
	Remote      *Remote
	Snapshot    *Snapshot
	Tasks       *Tasks
	AsyncSearch *AsyncSearch
	CCR         *CCR
	ILM         *ILM
	License     *License
	Migration   *Migration
	ML          *ML
	Monitoring  *Monitoring
	Rollup      *Rollup
	Security    *Security
	SQL         *SQL
	SSL         *SSL
	Watcher     *Watcher
	XPack       *XPack

	AutoscalingDeleteAutoscalingPolicy            AutoscalingDeleteAutoscalingPolicy
	AutoscalingGetAutoscalingCapacity             AutoscalingGetAutoscalingCapacity
	AutoscalingGetAutoscalingPolicy               AutoscalingGetAutoscalingPolicy
	AutoscalingPutAutoscalingPolicy               AutoscalingPutAutoscalingPolicy
	Bulk                                          Bulk
	Capabilities                                  Capabilities
	ClearScroll                                   ClearScroll
	ClosePointInTime                              ClosePointInTime
	ConnectorCheckIn                              ConnectorCheckIn
	ConnectorDelete                               ConnectorDelete
	ConnectorGet                                  ConnectorGet
	ConnectorLastSync                             ConnectorLastSync
	ConnectorList                                 ConnectorList
	ConnectorPost                                 ConnectorPost
	ConnectorPut                                  ConnectorPut
	ConnectorSecretDelete                         ConnectorSecretDelete
	ConnectorSecretGet                            ConnectorSecretGet
	ConnectorSecretPost                           ConnectorSecretPost
	ConnectorSecretPut                            ConnectorSecretPut
	ConnectorSyncJobCancel                        ConnectorSyncJobCancel
	ConnectorSyncJobCheckIn                       ConnectorSyncJobCheckIn
	ConnectorSyncJobClaim                         ConnectorSyncJobClaim
	ConnectorSyncJobDelete                        ConnectorSyncJobDelete
	ConnectorSyncJobError                         ConnectorSyncJobError
	ConnectorSyncJobGet                           ConnectorSyncJobGet
	ConnectorSyncJobList                          ConnectorSyncJobList
	ConnectorSyncJobPost                          ConnectorSyncJobPost
	ConnectorSyncJobUpdateStats                   ConnectorSyncJobUpdateStats
	ConnectorUpdateAPIKeyDocumentID               ConnectorUpdateAPIKeyDocumentID
	ConnectorUpdateActiveFiltering                ConnectorUpdateActiveFiltering
	ConnectorUpdateConfiguration                  ConnectorUpdateConfiguration
	ConnectorUpdateError                          ConnectorUpdateError
	ConnectorUpdateFeatures                       ConnectorUpdateFeatures
	ConnectorUpdateFiltering                      ConnectorUpdateFiltering
	ConnectorUpdateFilteringValidation            ConnectorUpdateFilteringValidation
	ConnectorUpdateIndexName                      ConnectorUpdateIndexName
	ConnectorUpdateName                           ConnectorUpdateName
	ConnectorUpdateNative                         ConnectorUpdateNative
	ConnectorUpdatePipeline                       ConnectorUpdatePipeline
	ConnectorUpdateScheduling                     ConnectorUpdateScheduling
	ConnectorUpdateServiceDocumentType            ConnectorUpdateServiceDocumentType
	ConnectorUpdateStatus                         ConnectorUpdateStatus
	Count                                         Count
	Create                                        Create
	DanglingIndicesDeleteDanglingIndex            DanglingIndicesDeleteDanglingIndex
	DanglingIndicesImportDanglingIndex            DanglingIndicesImportDanglingIndex
	DanglingIndicesListDanglingIndices            DanglingIndicesListDanglingIndices
	DeleteByQuery                                 DeleteByQuery
	DeleteByQueryRethrottle                       DeleteByQueryRethrottle
	Delete                                        Delete
	DeleteScript                                  DeleteScript
	EnrichDeletePolicy                            EnrichDeletePolicy
	EnrichExecutePolicy                           EnrichExecutePolicy
	EnrichGetPolicy                               EnrichGetPolicy
	EnrichPutPolicy                               EnrichPutPolicy
	EnrichStats                                   EnrichStats
	EqlDelete                                     EqlDelete
	EqlGet                                        EqlGet
	EqlGetStatus                                  EqlGetStatus
	EqlSearch                                     EqlSearch
	EsqlAsyncQueryDelete                          EsqlAsyncQueryDelete
	EsqlAsyncQueryGet                             EsqlAsyncQueryGet
	EsqlAsyncQuery                                EsqlAsyncQuery
	EsqlAsyncQueryStop                            EsqlAsyncQueryStop
	EsqlGetQuery                                  EsqlGetQuery
	EsqlListQueries                               EsqlListQueries
	EsqlQuery                                     EsqlQuery
	Exists                                        Exists
	ExistsSource                                  ExistsSource
	Explain                                       Explain
	FeaturesGetFeatures                           FeaturesGetFeatures
	FeaturesResetFeatures                         FeaturesResetFeatures
	FieldCaps                                     FieldCaps
	FleetDeleteSecret                             FleetDeleteSecret
	FleetGetSecret                                FleetGetSecret
	FleetGlobalCheckpoints                        FleetGlobalCheckpoints
	FleetMsearch                                  FleetMsearch
	FleetPostSecret                               FleetPostSecret
	FleetSearch                                   FleetSearch
	Get                                           Get
	GetScriptContext                              GetScriptContext
	GetScriptLanguages                            GetScriptLanguages
	GetScript                                     GetScript
	GetSource                                     GetSource
	GraphExplore                                  GraphExplore
	HealthReport                                  HealthReport
	Index                                         Index
	InferenceChatCompletionUnified                InferenceChatCompletionUnified
	InferenceCompletion                           InferenceCompletion
	InferenceDelete                               InferenceDelete
	InferenceGet                                  InferenceGet
	InferenceInference                            InferenceInference
	InferencePutAlibabacloud                      InferencePutAlibabacloud
	InferencePutAmazonbedrock                     InferencePutAmazonbedrock
	InferencePutAmazonsagemaker                   InferencePutAmazonsagemaker
	InferencePutAnthropic                         InferencePutAnthropic
	InferencePutAzureaistudio                     InferencePutAzureaistudio
	InferencePutAzureopenai                       InferencePutAzureopenai
	InferencePutCohere                            InferencePutCohere
	InferencePutCustom                            InferencePutCustom
	InferencePutDeepseek                          InferencePutDeepseek
	InferencePutElasticsearch                     InferencePutElasticsearch
	InferencePutElser                             InferencePutElser
	InferencePutGoogleaistudio                    InferencePutGoogleaistudio
	InferencePutGooglevertexai                    InferencePutGooglevertexai
	InferencePutHuggingFace                       InferencePutHuggingFace
	InferencePutJinaai                            InferencePutJinaai
	InferencePutMistral                           InferencePutMistral
	InferencePutOpenai                            InferencePutOpenai
	InferencePut                                  InferencePut
	InferencePutVoyageai                          InferencePutVoyageai
	InferencePutWatsonx                           InferencePutWatsonx
	InferenceRerank                               InferenceRerank
	InferenceSparseEmbedding                      InferenceSparseEmbedding
	InferenceStreamCompletion                     InferenceStreamCompletion
	InferenceTextEmbedding                        InferenceTextEmbedding
	InferenceUpdate                               InferenceUpdate
	Info                                          Info
	LogstashDeletePipeline                        LogstashDeletePipeline
	LogstashGetPipeline                           LogstashGetPipeline
	LogstashPutPipeline                           LogstashPutPipeline
	Mget                                          Mget
	Msearch                                       Msearch
	MsearchTemplate                               MsearchTemplate
	Mtermvectors                                  Mtermvectors
	OpenPointInTime                               OpenPointInTime
	Ping                                          Ping
	ProfilingFlamegraph                           ProfilingFlamegraph
	ProfilingStacktraces                          ProfilingStacktraces
	ProfilingStatus                               ProfilingStatus
	ProfilingTopnFunctions                        ProfilingTopnFunctions
	PutScript                                     PutScript
	QueryRulesDeleteRule                          QueryRulesDeleteRule
	QueryRulesDeleteRuleset                       QueryRulesDeleteRuleset
	QueryRulesGetRule                             QueryRulesGetRule
	QueryRulesGetRuleset                          QueryRulesGetRuleset
	QueryRulesListRulesets                        QueryRulesListRulesets
	QueryRulesPutRule                             QueryRulesPutRule
	QueryRulesPutRuleset                          QueryRulesPutRuleset
	QueryRulesTest                                QueryRulesTest
	RankEval                                      RankEval
	Reindex                                       Reindex
	ReindexRethrottle                             ReindexRethrottle
	RenderSearchTemplate                          RenderSearchTemplate
	ScriptsPainlessExecute                        ScriptsPainlessExecute
	Scroll                                        Scroll
	SearchApplicationDeleteBehavioralAnalytics    SearchApplicationDeleteBehavioralAnalytics
	SearchApplicationDelete                       SearchApplicationDelete
	SearchApplicationGetBehavioralAnalytics       SearchApplicationGetBehavioralAnalytics
	SearchApplicationGet                          SearchApplicationGet
	SearchApplicationList                         SearchApplicationList
	SearchApplicationPostBehavioralAnalyticsEvent SearchApplicationPostBehavioralAnalyticsEvent
	SearchApplicationPutBehavioralAnalytics       SearchApplicationPutBehavioralAnalytics
	SearchApplicationPut                          SearchApplicationPut
	SearchApplicationRenderQuery                  SearchApplicationRenderQuery
	SearchApplicationSearch                       SearchApplicationSearch
	SearchMvt                                     SearchMvt
	Search                                        Search
	SearchShards                                  SearchShards
	SearchTemplate                                SearchTemplate
	SearchableSnapshotsCacheStats                 SearchableSnapshotsCacheStats
	SearchableSnapshotsClearCache                 SearchableSnapshotsClearCache
	SearchableSnapshotsMount                      SearchableSnapshotsMount
	SearchableSnapshotsStats                      SearchableSnapshotsStats
	ShutdownDeleteNode                            ShutdownDeleteNode
	ShutdownGetNode                               ShutdownGetNode
	ShutdownPutNode                               ShutdownPutNode
	SimulateIngest                                SimulateIngest
	SlmDeleteLifecycle                            SlmDeleteLifecycle
	SlmExecuteLifecycle                           SlmExecuteLifecycle
	SlmExecuteRetention                           SlmExecuteRetention
	SlmGetLifecycle                               SlmGetLifecycle
	SlmGetStats                                   SlmGetStats
	SlmGetStatus                                  SlmGetStatus
	SlmPutLifecycle                               SlmPutLifecycle
	SlmStart                                      SlmStart
	SlmStop                                       SlmStop
	StreamsLogsDisable                            StreamsLogsDisable
	StreamsLogsEnable                             StreamsLogsEnable
	StreamsStatus                                 StreamsStatus
	SynonymsDeleteSynonym                         SynonymsDeleteSynonym
	SynonymsDeleteSynonymRule                     SynonymsDeleteSynonymRule
	SynonymsGetSynonym                            SynonymsGetSynonym
	SynonymsGetSynonymRule                        SynonymsGetSynonymRule
	SynonymsGetSynonymsSets                       SynonymsGetSynonymsSets
	SynonymsPutSynonym                            SynonymsPutSynonym
	SynonymsPutSynonymRule                        SynonymsPutSynonymRule
	TermsEnum                                     TermsEnum
	Termvectors                                   Termvectors
	TextStructureFindFieldStructure               TextStructureFindFieldStructure
	TextStructureFindMessageStructure             TextStructureFindMessageStructure
	TextStructureFindStructure                    TextStructureFindStructure
	TextStructureTestGrokPattern                  TextStructureTestGrokPattern
	TransformDeleteTransform                      TransformDeleteTransform
	TransformGetNodeStats                         TransformGetNodeStats
	TransformGetTransform                         TransformGetTransform
	TransformGetTransformStats                    TransformGetTransformStats
	TransformPreviewTransform                     TransformPreviewTransform
	TransformPutTransform                         TransformPutTransform
	TransformResetTransform                       TransformResetTransform
	TransformScheduleNowTransform                 TransformScheduleNowTransform
	TransformStartTransform                       TransformStartTransform
	TransformStopTransform                        TransformStopTransform
	TransformUpdateTransform                      TransformUpdateTransform
	TransformUpgradeTransforms                    TransformUpgradeTransforms
	UpdateByQuery                                 UpdateByQuery
	UpdateByQueryRethrottle                       UpdateByQueryRethrottle
	Update                                        Update
}

// Cat contains the Cat APIs
type Cat struct {
	Aliases              CatAliases
	Allocation           CatAllocation
	ComponentTemplates   CatComponentTemplates
	Count                CatCount
	Fielddata            CatFielddata
	Health               CatHealth
	Help                 CatHelp
	Indices              CatIndices
	MLDataFrameAnalytics CatMLDataFrameAnalytics
	MLDatafeeds          CatMLDatafeeds
	MLJobs               CatMLJobs
	MLTrainedModels      CatMLTrainedModels
	Master               CatMaster
	Nodeattrs            CatNodeattrs
	Nodes                CatNodes
	PendingTasks         CatPendingTasks
	Plugins              CatPlugins
	Recovery             CatRecovery
	Repositories         CatRepositories
	Segments             CatSegments
	Shards               CatShards
	Snapshots            CatSnapshots
	Tasks                CatTasks
	Templates            CatTemplates
	ThreadPool           CatThreadPool
	Transforms           CatTransforms
}

// Cluster contains the Cluster APIs
type Cluster struct {
	AllocationExplain            ClusterAllocationExplain
	DeleteComponentTemplate      ClusterDeleteComponentTemplate
	DeleteVotingConfigExclusions ClusterDeleteVotingConfigExclusions
	ExistsComponentTemplate      ClusterExistsComponentTemplate
	GetComponentTemplate         ClusterGetComponentTemplate
	GetSettings                  ClusterGetSettings
	Health                       ClusterHealth
	Info                         ClusterInfo
	PendingTasks                 ClusterPendingTasks
	PostVotingConfigExclusions   ClusterPostVotingConfigExclusions
	PutComponentTemplate         ClusterPutComponentTemplate
	PutSettings                  ClusterPutSettings
	RemoteInfo                   ClusterRemoteInfo
	Reroute                      ClusterReroute
	State                        ClusterState
	Stats                        ClusterStats
}

// Indices contains the Indices APIs
type Indices struct {
	AddBlock                IndicesAddBlock
	Analyze                 IndicesAnalyze
	CancelMigrateReindex    IndicesCancelMigrateReindex
	ClearCache              IndicesClearCache
	Clone                   IndicesClone
	Close                   IndicesClose
	CreateDataStream        IndicesCreateDataStream
	CreateFrom              IndicesCreateFrom
	Create                  IndicesCreate
	DataStreamsStats        IndicesDataStreamsStats
	DeleteAlias             IndicesDeleteAlias
	DeleteDataLifecycle     IndicesDeleteDataLifecycle
	DeleteDataStreamOptions IndicesDeleteDataStreamOptions
	DeleteDataStream        IndicesDeleteDataStream
	DeleteIndexTemplate     IndicesDeleteIndexTemplate
	Delete                  IndicesDelete
	DeleteTemplate          IndicesDeleteTemplate
	DiskUsage               IndicesDiskUsage
	Downsample              IndicesDownsample
	ExistsAlias             IndicesExistsAlias
	ExistsIndexTemplate     IndicesExistsIndexTemplate
	Exists                  IndicesExists
	ExistsTemplate          IndicesExistsTemplate
	ExplainDataLifecycle    IndicesExplainDataLifecycle
	FieldUsageStats         IndicesFieldUsageStats
	Flush                   IndicesFlush
	Forcemerge              IndicesForcemerge
	GetAlias                IndicesGetAlias
	GetDataLifecycle        IndicesGetDataLifecycle
	GetDataLifecycleStats   IndicesGetDataLifecycleStats
	GetDataStreamOptions    IndicesGetDataStreamOptions
	GetDataStream           IndicesGetDataStream
	GetDataStreamSettings   IndicesGetDataStreamSettings
	GetFieldMapping         IndicesGetFieldMapping
	GetIndexTemplate        IndicesGetIndexTemplate
	GetMapping              IndicesGetMapping
	GetMigrateReindexStatus IndicesGetMigrateReindexStatus
	Get                     IndicesGet
	GetSettings             IndicesGetSettings
	GetTemplate             IndicesGetTemplate
	MigrateReindex          IndicesMigrateReindex
	MigrateToDataStream     IndicesMigrateToDataStream
	ModifyDataStream        IndicesModifyDataStream
	Open                    IndicesOpen
	PromoteDataStream       IndicesPromoteDataStream
	PutAlias                IndicesPutAlias
	PutDataLifecycle        IndicesPutDataLifecycle
	PutDataStreamOptions    IndicesPutDataStreamOptions
	PutDataStreamSettings   IndicesPutDataStreamSettings
	PutIndexTemplate        IndicesPutIndexTemplate
	PutMapping              IndicesPutMapping
	PutSettings             IndicesPutSettings
	PutTemplate             IndicesPutTemplate
	Recovery                IndicesRecovery
	Refresh                 IndicesRefresh
	ReloadSearchAnalyzers   IndicesReloadSearchAnalyzers
	RemoveBlock             IndicesRemoveBlock
	ResolveCluster          IndicesResolveCluster
	ResolveIndex            IndicesResolveIndex
	Rollover                IndicesRollover
	Segments                IndicesSegments
	ShardStores             IndicesShardStores
	Shrink                  IndicesShrink
	SimulateIndexTemplate   IndicesSimulateIndexTemplate
	SimulateTemplate        IndicesSimulateTemplate
	Split                   IndicesSplit
	Stats                   IndicesStats
	UpdateAliases           IndicesUpdateAliases
	ValidateQuery           IndicesValidateQuery
}

// Ingest contains the Ingest APIs
type Ingest struct {
	DeleteGeoipDatabase      IngestDeleteGeoipDatabase
	DeleteIPLocationDatabase IngestDeleteIPLocationDatabase
	DeletePipeline           IngestDeletePipeline
	GeoIPStats               IngestGeoIPStats
	GetGeoipDatabase         IngestGetGeoipDatabase
	GetIPLocationDatabase    IngestGetIPLocationDatabase
	GetPipeline              IngestGetPipeline
	ProcessorGrok            IngestProcessorGrok
	PutGeoipDatabase         IngestPutGeoipDatabase
	PutIPLocationDatabase    IngestPutIPLocationDatabase
	PutPipeline              IngestPutPipeline
	Simulate                 IngestSimulate
}

// Nodes contains the Nodes APIs
type Nodes struct {
	ClearRepositoriesMeteringArchive NodesClearRepositoriesMeteringArchive
	GetRepositoriesMeteringInfo      NodesGetRepositoriesMeteringInfo
	HotThreads                       NodesHotThreads
	Info                             NodesInfo
	ReloadSecureSettings             NodesReloadSecureSettings
	Stats                            NodesStats
	Usage                            NodesUsage
}

// Remote contains the Remote APIs
type Remote struct {
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
	CleanupRepository         SnapshotCleanupRepository
	Clone                     SnapshotClone
	CreateRepository          SnapshotCreateRepository
	Create                    SnapshotCreate
	DeleteRepository          SnapshotDeleteRepository
	Delete                    SnapshotDelete
	GetRepository             SnapshotGetRepository
	Get                       SnapshotGet
	RepositoryAnalyze         SnapshotRepositoryAnalyze
	RepositoryVerifyIntegrity SnapshotRepositoryVerifyIntegrity
	Restore                   SnapshotRestore
	Status                    SnapshotStatus
	VerifyRepository          SnapshotVerifyRepository
}

// Tasks contains the Tasks APIs
type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

// AsyncSearch contains the AsyncSearch APIs
type AsyncSearch struct {
	Delete AsyncSearchDelete
	Get    AsyncSearchGet
	Status AsyncSearchStatus
	Submit AsyncSearchSubmit
}

// CCR contains the CCR APIs
type CCR struct {
	DeleteAutoFollowPattern CCRDeleteAutoFollowPattern
	FollowInfo              CCRFollowInfo
	Follow                  CCRFollow
	FollowStats             CCRFollowStats
	ForgetFollower          CCRForgetFollower
	GetAutoFollowPattern    CCRGetAutoFollowPattern
	PauseAutoFollowPattern  CCRPauseAutoFollowPattern
	PauseFollow             CCRPauseFollow
	PutAutoFollowPattern    CCRPutAutoFollowPattern
	ResumeAutoFollowPattern CCRResumeAutoFollowPattern
	ResumeFollow            CCRResumeFollow
	Stats                   CCRStats
	Unfollow                CCRUnfollow
}

// ILM contains the ILM APIs
type ILM struct {
	DeleteLifecycle    ILMDeleteLifecycle
	ExplainLifecycle   ILMExplainLifecycle
	GetLifecycle       ILMGetLifecycle
	GetStatus          ILMGetStatus
	MigrateToDataTiers ILMMigrateToDataTiers
	MoveToStep         ILMMoveToStep
	PutLifecycle       ILMPutLifecycle
	RemovePolicy       ILMRemovePolicy
	Retry              ILMRetry
	Start              ILMStart
	Stop               ILMStop
}

// License contains the License APIs
type License struct {
	Delete         LicenseDelete
	GetBasicStatus LicenseGetBasicStatus
	Get            LicenseGet
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations            MigrationDeprecations
	GetFeatureUpgradeStatus MigrationGetFeatureUpgradeStatus
	PostFeatureUpgrade      MigrationPostFeatureUpgrade
}

// ML contains the ML APIs
type ML struct {
	ClearTrainedModelDeploymentCache MLClearTrainedModelDeploymentCache
	CloseJob                         MLCloseJob
	DeleteCalendarEvent              MLDeleteCalendarEvent
	DeleteCalendarJob                MLDeleteCalendarJob
	DeleteCalendar                   MLDeleteCalendar
	DeleteDataFrameAnalytics         MLDeleteDataFrameAnalytics
	DeleteDatafeed                   MLDeleteDatafeed
	DeleteExpiredData                MLDeleteExpiredData
	DeleteFilter                     MLDeleteFilter
	DeleteForecast                   MLDeleteForecast
	DeleteJob                        MLDeleteJob
	DeleteModelSnapshot              MLDeleteModelSnapshot
	DeleteTrainedModelAlias          MLDeleteTrainedModelAlias
	DeleteTrainedModel               MLDeleteTrainedModel
	EstimateModelMemory              MLEstimateModelMemory
	EvaluateDataFrame                MLEvaluateDataFrame
	ExplainDataFrameAnalytics        MLExplainDataFrameAnalytics
	FlushJob                         MLFlushJob
	Forecast                         MLForecast
	GetBuckets                       MLGetBuckets
	GetCalendarEvents                MLGetCalendarEvents
	GetCalendars                     MLGetCalendars
	GetCategories                    MLGetCategories
	GetDataFrameAnalytics            MLGetDataFrameAnalytics
	GetDataFrameAnalyticsStats       MLGetDataFrameAnalyticsStats
	GetDatafeedStats                 MLGetDatafeedStats
	GetDatafeeds                     MLGetDatafeeds
	GetFilters                       MLGetFilters
	GetInfluencers                   MLGetInfluencers
	GetJobStats                      MLGetJobStats
	GetJobs                          MLGetJobs
	GetMemoryStats                   MLGetMemoryStats
	GetModelSnapshotUpgradeStats     MLGetModelSnapshotUpgradeStats
	GetModelSnapshots                MLGetModelSnapshots
	GetOverallBuckets                MLGetOverallBuckets
	GetRecords                       MLGetRecords
	GetTrainedModels                 MLGetTrainedModels
	GetTrainedModelsStats            MLGetTrainedModelsStats
	InferTrainedModel                MLInferTrainedModel
	Info                             MLInfo
	OpenJob                          MLOpenJob
	PostCalendarEvents               MLPostCalendarEvents
	PostData                         MLPostData
	PreviewDataFrameAnalytics        MLPreviewDataFrameAnalytics
	PreviewDatafeed                  MLPreviewDatafeed
	PutCalendarJob                   MLPutCalendarJob
	PutCalendar                      MLPutCalendar
	PutDataFrameAnalytics            MLPutDataFrameAnalytics
	PutDatafeed                      MLPutDatafeed
	PutFilter                        MLPutFilter
	PutJob                           MLPutJob
	PutTrainedModelAlias             MLPutTrainedModelAlias
	PutTrainedModelDefinitionPart    MLPutTrainedModelDefinitionPart
	PutTrainedModel                  MLPutTrainedModel
	PutTrainedModelVocabulary        MLPutTrainedModelVocabulary
	ResetJob                         MLResetJob
	RevertModelSnapshot              MLRevertModelSnapshot
	SetUpgradeMode                   MLSetUpgradeMode
	StartDataFrameAnalytics          MLStartDataFrameAnalytics
	StartDatafeed                    MLStartDatafeed
	StartTrainedModelDeployment      MLStartTrainedModelDeployment
	StopDataFrameAnalytics           MLStopDataFrameAnalytics
	StopDatafeed                     MLStopDatafeed
	StopTrainedModelDeployment       MLStopTrainedModelDeployment
	UpdateDataFrameAnalytics         MLUpdateDataFrameAnalytics
	UpdateDatafeed                   MLUpdateDatafeed
	UpdateFilter                     MLUpdateFilter
	UpdateJob                        MLUpdateJob
	UpdateModelSnapshot              MLUpdateModelSnapshot
	UpdateTrainedModelDeployment     MLUpdateTrainedModelDeployment
	UpgradeJobSnapshot               MLUpgradeJobSnapshot
	ValidateDetector                 MLValidateDetector
	Validate                         MLValidate
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk MonitoringBulk
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob    RollupDeleteJob
	GetJobs      RollupGetJobs
	GetCaps      RollupGetRollupCaps
	GetIndexCaps RollupGetRollupIndexCaps
	PutJob       RollupPutJob
	Search       RollupRollupSearch
	StartJob     RollupStartJob
	StopJob      RollupStopJob
}

// Security contains the Security APIs
type Security struct {
	ActivateUserProfile         SecurityActivateUserProfile
	Authenticate                SecurityAuthenticate
	BulkDeleteRole              SecurityBulkDeleteRole
	BulkPutRole                 SecurityBulkPutRole
	BulkUpdateAPIKeys           SecurityBulkUpdateAPIKeys
	ChangePassword              SecurityChangePassword
	ClearAPIKeyCache            SecurityClearAPIKeyCache
	ClearCachedPrivileges       SecurityClearCachedPrivileges
	ClearCachedRealms           SecurityClearCachedRealms
	ClearCachedRoles            SecurityClearCachedRoles
	ClearCachedServiceTokens    SecurityClearCachedServiceTokens
	CreateAPIKey                SecurityCreateAPIKey
	CreateCrossClusterAPIKey    SecurityCreateCrossClusterAPIKey
	CreateServiceToken          SecurityCreateServiceToken
	DelegatePki                 SecurityDelegatePki
	DeletePrivileges            SecurityDeletePrivileges
	DeleteRoleMapping           SecurityDeleteRoleMapping
	DeleteRole                  SecurityDeleteRole
	DeleteServiceToken          SecurityDeleteServiceToken
	DeleteUser                  SecurityDeleteUser
	DisableUserProfile          SecurityDisableUserProfile
	DisableUser                 SecurityDisableUser
	EnableUserProfile           SecurityEnableUserProfile
	EnableUser                  SecurityEnableUser
	EnrollKibana                SecurityEnrollKibana
	EnrollNode                  SecurityEnrollNode
	GetAPIKey                   SecurityGetAPIKey
	GetBuiltinPrivileges        SecurityGetBuiltinPrivileges
	GetPrivileges               SecurityGetPrivileges
	GetRoleMapping              SecurityGetRoleMapping
	GetRole                     SecurityGetRole
	GetServiceAccounts          SecurityGetServiceAccounts
	GetServiceCredentials       SecurityGetServiceCredentials
	GetSettings                 SecurityGetSettings
	GetToken                    SecurityGetToken
	GetUserPrivileges           SecurityGetUserPrivileges
	GetUserProfile              SecurityGetUserProfile
	GetUser                     SecurityGetUser
	GrantAPIKey                 SecurityGrantAPIKey
	HasPrivileges               SecurityHasPrivileges
	HasPrivilegesUserProfile    SecurityHasPrivilegesUserProfile
	InvalidateAPIKey            SecurityInvalidateAPIKey
	InvalidateToken             SecurityInvalidateToken
	OidcAuthenticate            SecurityOidcAuthenticate
	OidcLogout                  SecurityOidcLogout
	OidcPrepareAuthentication   SecurityOidcPrepareAuthentication
	PutPrivileges               SecurityPutPrivileges
	PutRoleMapping              SecurityPutRoleMapping
	PutRole                     SecurityPutRole
	PutUser                     SecurityPutUser
	QueryAPIKeys                SecurityQueryAPIKeys
	QueryRole                   SecurityQueryRole
	QueryUser                   SecurityQueryUser
	SamlAuthenticate            SecuritySamlAuthenticate
	SamlCompleteLogout          SecuritySamlCompleteLogout
	SamlInvalidate              SecuritySamlInvalidate
	SamlLogout                  SecuritySamlLogout
	SamlPrepareAuthentication   SecuritySamlPrepareAuthentication
	SamlServiceProviderMetadata SecuritySamlServiceProviderMetadata
	SuggestUserProfiles         SecuritySuggestUserProfiles
	UpdateAPIKey                SecurityUpdateAPIKey
	UpdateCrossClusterAPIKey    SecurityUpdateCrossClusterAPIKey
	UpdateSettings              SecurityUpdateSettings
	UpdateUserProfileData       SecurityUpdateUserProfileData
}

// SQL contains the SQL APIs
type SQL struct {
	ClearCursor    SQLClearCursor
	DeleteAsync    SQLDeleteAsync
	GetAsync       SQLGetAsync
	GetAsyncStatus SQLGetAsyncStatus
	Query          SQLQuery
	Translate      SQLTranslate
}

// SSL contains the SSL APIs
type SSL struct {
	Certificates SSLCertificates
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetSettings     WatcherGetSettings
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	QueryWatches    WatcherQueryWatches
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
	UpdateSettings  WatcherUpdateSettings
}

// XPack contains the XPack APIs
type XPack struct {
	Info  XPackInfo
	Usage XPackUsage
}

// New creates new API
func New(t Transport) *API {
	return &API{
		AutoscalingDeleteAutoscalingPolicy: newAutoscalingDeleteAutoscalingPolicyFunc(t),
		AutoscalingGetAutoscalingCapacity:  newAutoscalingGetAutoscalingCapacityFunc(t),
		AutoscalingGetAutoscalingPolicy:    newAutoscalingGetAutoscalingPolicyFunc(t),
		AutoscalingPutAutoscalingPolicy:    newAutoscalingPutAutoscalingPolicyFunc(t),
		Bulk:                               newBulkFunc(t),
		Capabilities:                       newCapabilitiesFunc(t),
		ClearScroll:                        newClearScrollFunc(t),
		ClosePointInTime:                   newClosePointInTimeFunc(t),
		ConnectorCheckIn:                   newConnectorCheckInFunc(t),
		ConnectorDelete:                    newConnectorDeleteFunc(t),
		ConnectorGet:                       newConnectorGetFunc(t),
		ConnectorLastSync:                  newConnectorLastSyncFunc(t),
		ConnectorList:                      newConnectorListFunc(t),
		ConnectorPost:                      newConnectorPostFunc(t),
		ConnectorPut:                       newConnectorPutFunc(t),
		ConnectorSecretDelete:              newConnectorSecretDeleteFunc(t),
		ConnectorSecretGet:                 newConnectorSecretGetFunc(t),
		ConnectorSecretPost:                newConnectorSecretPostFunc(t),
		ConnectorSecretPut:                 newConnectorSecretPutFunc(t),
		ConnectorSyncJobCancel:             newConnectorSyncJobCancelFunc(t),
		ConnectorSyncJobCheckIn:            newConnectorSyncJobCheckInFunc(t),
		ConnectorSyncJobClaim:              newConnectorSyncJobClaimFunc(t),
		ConnectorSyncJobDelete:             newConnectorSyncJobDeleteFunc(t),
		ConnectorSyncJobError:              newConnectorSyncJobErrorFunc(t),
		ConnectorSyncJobGet:                newConnectorSyncJobGetFunc(t),
		ConnectorSyncJobList:               newConnectorSyncJobListFunc(t),
		ConnectorSyncJobPost:               newConnectorSyncJobPostFunc(t),
		ConnectorSyncJobUpdateStats:        newConnectorSyncJobUpdateStatsFunc(t),
		ConnectorUpdateAPIKeyDocumentID:    newConnectorUpdateAPIKeyDocumentIDFunc(t),
		ConnectorUpdateActiveFiltering:     newConnectorUpdateActiveFilteringFunc(t),
		ConnectorUpdateConfiguration:       newConnectorUpdateConfigurationFunc(t),
		ConnectorUpdateError:               newConnectorUpdateErrorFunc(t),
		ConnectorUpdateFeatures:            newConnectorUpdateFeaturesFunc(t),
		ConnectorUpdateFiltering:           newConnectorUpdateFilteringFunc(t),
		ConnectorUpdateFilteringValidation: newConnectorUpdateFilteringValidationFunc(t),
		ConnectorUpdateIndexName:           newConnectorUpdateIndexNameFunc(t),
		ConnectorUpdateName:                newConnectorUpdateNameFunc(t),
		ConnectorUpdateNative:              newConnectorUpdateNativeFunc(t),
		ConnectorUpdatePipeline:            newConnectorUpdatePipelineFunc(t),
		ConnectorUpdateScheduling:          newConnectorUpdateSchedulingFunc(t),
		ConnectorUpdateServiceDocumentType: newConnectorUpdateServiceDocumentTypeFunc(t),
		ConnectorUpdateStatus:              newConnectorUpdateStatusFunc(t),
		Count:                              newCountFunc(t),
		Create:                             newCreateFunc(t),
		DanglingIndicesDeleteDanglingIndex: newDanglingIndicesDeleteDanglingIndexFunc(t),
		DanglingIndicesImportDanglingIndex: newDanglingIndicesImportDanglingIndexFunc(t),
		DanglingIndicesListDanglingIndices: newDanglingIndicesListDanglingIndicesFunc(t),
		DeleteByQuery:                      newDeleteByQueryFunc(t),
		DeleteByQueryRethrottle:            newDeleteByQueryRethrottleFunc(t),
		Delete:                             newDeleteFunc(t),
		DeleteScript:                       newDeleteScriptFunc(t),
		EnrichDeletePolicy:                 newEnrichDeletePolicyFunc(t),
		EnrichExecutePolicy:                newEnrichExecutePolicyFunc(t),
		EnrichGetPolicy:                    newEnrichGetPolicyFunc(t),
		EnrichPutPolicy:                    newEnrichPutPolicyFunc(t),
		EnrichStats:                        newEnrichStatsFunc(t),
		EqlDelete:                          newEqlDeleteFunc(t),
		EqlGet:                             newEqlGetFunc(t),
		EqlGetStatus:                       newEqlGetStatusFunc(t),
		EqlSearch:                          newEqlSearchFunc(t),
		EsqlAsyncQueryDelete:               newEsqlAsyncQueryDeleteFunc(t),
		EsqlAsyncQueryGet:                  newEsqlAsyncQueryGetFunc(t),
		EsqlAsyncQuery:                     newEsqlAsyncQueryFunc(t),
		EsqlAsyncQueryStop:                 newEsqlAsyncQueryStopFunc(t),
		EsqlGetQuery:                       newEsqlGetQueryFunc(t),
		EsqlListQueries:                    newEsqlListQueriesFunc(t),
		EsqlQuery:                          newEsqlQueryFunc(t),
		Exists:                             newExistsFunc(t),
		ExistsSource:                       newExistsSourceFunc(t),
		Explain:                            newExplainFunc(t),
		FeaturesGetFeatures:                newFeaturesGetFeaturesFunc(t),
		FeaturesResetFeatures:              newFeaturesResetFeaturesFunc(t),
		FieldCaps:                          newFieldCapsFunc(t),
		FleetDeleteSecret:                  newFleetDeleteSecretFunc(t),
		FleetGetSecret:                     newFleetGetSecretFunc(t),
		FleetGlobalCheckpoints:             newFleetGlobalCheckpointsFunc(t),
		FleetMsearch:                       newFleetMsearchFunc(t),
		FleetPostSecret:                    newFleetPostSecretFunc(t),
		FleetSearch:                        newFleetSearchFunc(t),
		Get:                                newGetFunc(t),
		GetScriptContext:                   newGetScriptContextFunc(t),
		GetScriptLanguages:                 newGetScriptLanguagesFunc(t),
		GetScript:                          newGetScriptFunc(t),
		GetSource:                          newGetSourceFunc(t),
		GraphExplore:                       newGraphExploreFunc(t),
		HealthReport:                       newHealthReportFunc(t),
		Index:                              newIndexFunc(t),
		InferenceChatCompletionUnified:     newInferenceChatCompletionUnifiedFunc(t),
		InferenceCompletion:                newInferenceCompletionFunc(t),
		InferenceDelete:                    newInferenceDeleteFunc(t),
		InferenceGet:                       newInferenceGetFunc(t),
		InferenceInference:                 newInferenceInferenceFunc(t),
		InferencePutAlibabacloud:           newInferencePutAlibabacloudFunc(t),
		InferencePutAmazonbedrock:          newInferencePutAmazonbedrockFunc(t),
		InferencePutAmazonsagemaker:        newInferencePutAmazonsagemakerFunc(t),
		InferencePutAnthropic:              newInferencePutAnthropicFunc(t),
		InferencePutAzureaistudio:          newInferencePutAzureaistudioFunc(t),
		InferencePutAzureopenai:            newInferencePutAzureopenaiFunc(t),
		InferencePutCohere:                 newInferencePutCohereFunc(t),
		InferencePutCustom:                 newInferencePutCustomFunc(t),
		InferencePutDeepseek:               newInferencePutDeepseekFunc(t),
		InferencePutElasticsearch:          newInferencePutElasticsearchFunc(t),
		InferencePutElser:                  newInferencePutElserFunc(t),
		InferencePutGoogleaistudio:         newInferencePutGoogleaistudioFunc(t),
		InferencePutGooglevertexai:         newInferencePutGooglevertexaiFunc(t),
		InferencePutHuggingFace:            newInferencePutHuggingFaceFunc(t),
		InferencePutJinaai:                 newInferencePutJinaaiFunc(t),
		InferencePutMistral:                newInferencePutMistralFunc(t),
		InferencePutOpenai:                 newInferencePutOpenaiFunc(t),
		InferencePut:                       newInferencePutFunc(t),
		InferencePutVoyageai:               newInferencePutVoyageaiFunc(t),
		InferencePutWatsonx:                newInferencePutWatsonxFunc(t),
		InferenceRerank:                    newInferenceRerankFunc(t),
		InferenceSparseEmbedding:           newInferenceSparseEmbeddingFunc(t),
		InferenceStreamCompletion:          newInferenceStreamCompletionFunc(t),
		InferenceTextEmbedding:             newInferenceTextEmbeddingFunc(t),
		InferenceUpdate:                    newInferenceUpdateFunc(t),
		Info:                               newInfoFunc(t),
		LogstashDeletePipeline:             newLogstashDeletePipelineFunc(t),
		LogstashGetPipeline:                newLogstashGetPipelineFunc(t),
		LogstashPutPipeline:                newLogstashPutPipelineFunc(t),
		Mget:                               newMgetFunc(t),
		Msearch:                            newMsearchFunc(t),
		MsearchTemplate:                    newMsearchTemplateFunc(t),
		Mtermvectors:                       newMtermvectorsFunc(t),
		OpenPointInTime:                    newOpenPointInTimeFunc(t),
		Ping:                               newPingFunc(t),
		ProfilingFlamegraph:                newProfilingFlamegraphFunc(t),
		ProfilingStacktraces:               newProfilingStacktracesFunc(t),
		ProfilingStatus:                    newProfilingStatusFunc(t),
		ProfilingTopnFunctions:             newProfilingTopnFunctionsFunc(t),
		PutScript:                          newPutScriptFunc(t),
		QueryRulesDeleteRule:               newQueryRulesDeleteRuleFunc(t),
		QueryRulesDeleteRuleset:            newQueryRulesDeleteRulesetFunc(t),
		QueryRulesGetRule:                  newQueryRulesGetRuleFunc(t),
		QueryRulesGetRuleset:               newQueryRulesGetRulesetFunc(t),
		QueryRulesListRulesets:             newQueryRulesListRulesetsFunc(t),
		QueryRulesPutRule:                  newQueryRulesPutRuleFunc(t),
		QueryRulesPutRuleset:               newQueryRulesPutRulesetFunc(t),
		QueryRulesTest:                     newQueryRulesTestFunc(t),
		RankEval:                           newRankEvalFunc(t),
		Reindex:                            newReindexFunc(t),
		ReindexRethrottle:                  newReindexRethrottleFunc(t),
		RenderSearchTemplate:               newRenderSearchTemplateFunc(t),
		ScriptsPainlessExecute:             newScriptsPainlessExecuteFunc(t),
		Scroll:                             newScrollFunc(t),
		SearchApplicationDeleteBehavioralAnalytics:    newSearchApplicationDeleteBehavioralAnalyticsFunc(t),
		SearchApplicationDelete:                       newSearchApplicationDeleteFunc(t),
		SearchApplicationGetBehavioralAnalytics:       newSearchApplicationGetBehavioralAnalyticsFunc(t),
		SearchApplicationGet:                          newSearchApplicationGetFunc(t),
		SearchApplicationList:                         newSearchApplicationListFunc(t),
		SearchApplicationPostBehavioralAnalyticsEvent: newSearchApplicationPostBehavioralAnalyticsEventFunc(t),
		SearchApplicationPutBehavioralAnalytics:       newSearchApplicationPutBehavioralAnalyticsFunc(t),
		SearchApplicationPut:                          newSearchApplicationPutFunc(t),
		SearchApplicationRenderQuery:                  newSearchApplicationRenderQueryFunc(t),
		SearchApplicationSearch:                       newSearchApplicationSearchFunc(t),
		SearchMvt:                                     newSearchMvtFunc(t),
		Search:                                        newSearchFunc(t),
		SearchShards:                                  newSearchShardsFunc(t),
		SearchTemplate:                                newSearchTemplateFunc(t),
		SearchableSnapshotsCacheStats:                 newSearchableSnapshotsCacheStatsFunc(t),
		SearchableSnapshotsClearCache:                 newSearchableSnapshotsClearCacheFunc(t),
		SearchableSnapshotsMount:                      newSearchableSnapshotsMountFunc(t),
		SearchableSnapshotsStats:                      newSearchableSnapshotsStatsFunc(t),
		ShutdownDeleteNode:                            newShutdownDeleteNodeFunc(t),
		ShutdownGetNode:                               newShutdownGetNodeFunc(t),
		ShutdownPutNode:                               newShutdownPutNodeFunc(t),
		SimulateIngest:                                newSimulateIngestFunc(t),
		SlmDeleteLifecycle:                            newSlmDeleteLifecycleFunc(t),
		SlmExecuteLifecycle:                           newSlmExecuteLifecycleFunc(t),
		SlmExecuteRetention:                           newSlmExecuteRetentionFunc(t),
		SlmGetLifecycle:                               newSlmGetLifecycleFunc(t),
		SlmGetStats:                                   newSlmGetStatsFunc(t),
		SlmGetStatus:                                  newSlmGetStatusFunc(t),
		SlmPutLifecycle:                               newSlmPutLifecycleFunc(t),
		SlmStart:                                      newSlmStartFunc(t),
		SlmStop:                                       newSlmStopFunc(t),
		StreamsLogsDisable:                            newStreamsLogsDisableFunc(t),
		StreamsLogsEnable:                             newStreamsLogsEnableFunc(t),
		StreamsStatus:                                 newStreamsStatusFunc(t),
		SynonymsDeleteSynonym:                         newSynonymsDeleteSynonymFunc(t),
		SynonymsDeleteSynonymRule:                     newSynonymsDeleteSynonymRuleFunc(t),
		SynonymsGetSynonym:                            newSynonymsGetSynonymFunc(t),
		SynonymsGetSynonymRule:                        newSynonymsGetSynonymRuleFunc(t),
		SynonymsGetSynonymsSets:                       newSynonymsGetSynonymsSetsFunc(t),
		SynonymsPutSynonym:                            newSynonymsPutSynonymFunc(t),
		SynonymsPutSynonymRule:                        newSynonymsPutSynonymRuleFunc(t),
		TermsEnum:                                     newTermsEnumFunc(t),
		Termvectors:                                   newTermvectorsFunc(t),
		TextStructureFindFieldStructure:               newTextStructureFindFieldStructureFunc(t),
		TextStructureFindMessageStructure:             newTextStructureFindMessageStructureFunc(t),
		TextStructureFindStructure:                    newTextStructureFindStructureFunc(t),
		TextStructureTestGrokPattern:                  newTextStructureTestGrokPatternFunc(t),
		TransformDeleteTransform:                      newTransformDeleteTransformFunc(t),
		TransformGetNodeStats:                         newTransformGetNodeStatsFunc(t),
		TransformGetTransform:                         newTransformGetTransformFunc(t),
		TransformGetTransformStats:                    newTransformGetTransformStatsFunc(t),
		TransformPreviewTransform:                     newTransformPreviewTransformFunc(t),
		TransformPutTransform:                         newTransformPutTransformFunc(t),
		TransformResetTransform:                       newTransformResetTransformFunc(t),
		TransformScheduleNowTransform:                 newTransformScheduleNowTransformFunc(t),
		TransformStartTransform:                       newTransformStartTransformFunc(t),
		TransformStopTransform:                        newTransformStopTransformFunc(t),
		TransformUpdateTransform:                      newTransformUpdateTransformFunc(t),
		TransformUpgradeTransforms:                    newTransformUpgradeTransformsFunc(t),
		UpdateByQuery:                                 newUpdateByQueryFunc(t),
		UpdateByQueryRethrottle:                       newUpdateByQueryRethrottleFunc(t),
		Update:                                        newUpdateFunc(t),
		Cat: &Cat{
			Aliases:              newCatAliasesFunc(t),
			Allocation:           newCatAllocationFunc(t),
			ComponentTemplates:   newCatComponentTemplatesFunc(t),
			Count:                newCatCountFunc(t),
			Fielddata:            newCatFielddataFunc(t),
			Health:               newCatHealthFunc(t),
			Help:                 newCatHelpFunc(t),
			Indices:              newCatIndicesFunc(t),
			MLDataFrameAnalytics: newCatMLDataFrameAnalyticsFunc(t),
			MLDatafeeds:          newCatMLDatafeedsFunc(t),
			MLJobs:               newCatMLJobsFunc(t),
			MLTrainedModels:      newCatMLTrainedModelsFunc(t),
			Master:               newCatMasterFunc(t),
			Nodeattrs:            newCatNodeattrsFunc(t),
			Nodes:                newCatNodesFunc(t),
			PendingTasks:         newCatPendingTasksFunc(t),
			Plugins:              newCatPluginsFunc(t),
			Recovery:             newCatRecoveryFunc(t),
			Repositories:         newCatRepositoriesFunc(t),
			Segments:             newCatSegmentsFunc(t),
			Shards:               newCatShardsFunc(t),
			Snapshots:            newCatSnapshotsFunc(t),
			Tasks:                newCatTasksFunc(t),
			Templates:            newCatTemplatesFunc(t),
			ThreadPool:           newCatThreadPoolFunc(t),
			Transforms:           newCatTransformsFunc(t),
		},
		Cluster: &Cluster{
			AllocationExplain:            newClusterAllocationExplainFunc(t),
			DeleteComponentTemplate:      newClusterDeleteComponentTemplateFunc(t),
			DeleteVotingConfigExclusions: newClusterDeleteVotingConfigExclusionsFunc(t),
			ExistsComponentTemplate:      newClusterExistsComponentTemplateFunc(t),
			GetComponentTemplate:         newClusterGetComponentTemplateFunc(t),
			GetSettings:                  newClusterGetSettingsFunc(t),
			Health:                       newClusterHealthFunc(t),
			Info:                         newClusterInfoFunc(t),
			PendingTasks:                 newClusterPendingTasksFunc(t),
			PostVotingConfigExclusions:   newClusterPostVotingConfigExclusionsFunc(t),
			PutComponentTemplate:         newClusterPutComponentTemplateFunc(t),
			PutSettings:                  newClusterPutSettingsFunc(t),
			RemoteInfo:                   newClusterRemoteInfoFunc(t),
			Reroute:                      newClusterRerouteFunc(t),
			State:                        newClusterStateFunc(t),
			Stats:                        newClusterStatsFunc(t),
		},
		Indices: &Indices{
			AddBlock:                newIndicesAddBlockFunc(t),
			Analyze:                 newIndicesAnalyzeFunc(t),
			CancelMigrateReindex:    newIndicesCancelMigrateReindexFunc(t),
			ClearCache:              newIndicesClearCacheFunc(t),
			Clone:                   newIndicesCloneFunc(t),
			Close:                   newIndicesCloseFunc(t),
			CreateDataStream:        newIndicesCreateDataStreamFunc(t),
			CreateFrom:              newIndicesCreateFromFunc(t),
			Create:                  newIndicesCreateFunc(t),
			DataStreamsStats:        newIndicesDataStreamsStatsFunc(t),
			DeleteAlias:             newIndicesDeleteAliasFunc(t),
			DeleteDataLifecycle:     newIndicesDeleteDataLifecycleFunc(t),
			DeleteDataStreamOptions: newIndicesDeleteDataStreamOptionsFunc(t),
			DeleteDataStream:        newIndicesDeleteDataStreamFunc(t),
			DeleteIndexTemplate:     newIndicesDeleteIndexTemplateFunc(t),
			Delete:                  newIndicesDeleteFunc(t),
			DeleteTemplate:          newIndicesDeleteTemplateFunc(t),
			DiskUsage:               newIndicesDiskUsageFunc(t),
			Downsample:              newIndicesDownsampleFunc(t),
			ExistsAlias:             newIndicesExistsAliasFunc(t),
			ExistsIndexTemplate:     newIndicesExistsIndexTemplateFunc(t),
			Exists:                  newIndicesExistsFunc(t),
			ExistsTemplate:          newIndicesExistsTemplateFunc(t),
			ExplainDataLifecycle:    newIndicesExplainDataLifecycleFunc(t),
			FieldUsageStats:         newIndicesFieldUsageStatsFunc(t),
			Flush:                   newIndicesFlushFunc(t),
			Forcemerge:              newIndicesForcemergeFunc(t),
			GetAlias:                newIndicesGetAliasFunc(t),
			GetDataLifecycle:        newIndicesGetDataLifecycleFunc(t),
			GetDataLifecycleStats:   newIndicesGetDataLifecycleStatsFunc(t),
			GetDataStreamOptions:    newIndicesGetDataStreamOptionsFunc(t),
			GetDataStream:           newIndicesGetDataStreamFunc(t),
			GetDataStreamSettings:   newIndicesGetDataStreamSettingsFunc(t),
			GetFieldMapping:         newIndicesGetFieldMappingFunc(t),
			GetIndexTemplate:        newIndicesGetIndexTemplateFunc(t),
			GetMapping:              newIndicesGetMappingFunc(t),
			GetMigrateReindexStatus: newIndicesGetMigrateReindexStatusFunc(t),
			Get:                     newIndicesGetFunc(t),
			GetSettings:             newIndicesGetSettingsFunc(t),
			GetTemplate:             newIndicesGetTemplateFunc(t),
			MigrateReindex:          newIndicesMigrateReindexFunc(t),
			MigrateToDataStream:     newIndicesMigrateToDataStreamFunc(t),
			ModifyDataStream:        newIndicesModifyDataStreamFunc(t),
			Open:                    newIndicesOpenFunc(t),
			PromoteDataStream:       newIndicesPromoteDataStreamFunc(t),
			PutAlias:                newIndicesPutAliasFunc(t),
			PutDataLifecycle:        newIndicesPutDataLifecycleFunc(t),
			PutDataStreamOptions:    newIndicesPutDataStreamOptionsFunc(t),
			PutDataStreamSettings:   newIndicesPutDataStreamSettingsFunc(t),
			PutIndexTemplate:        newIndicesPutIndexTemplateFunc(t),
			PutMapping:              newIndicesPutMappingFunc(t),
			PutSettings:             newIndicesPutSettingsFunc(t),
			PutTemplate:             newIndicesPutTemplateFunc(t),
			Recovery:                newIndicesRecoveryFunc(t),
			Refresh:                 newIndicesRefreshFunc(t),
			ReloadSearchAnalyzers:   newIndicesReloadSearchAnalyzersFunc(t),
			RemoveBlock:             newIndicesRemoveBlockFunc(t),
			ResolveCluster:          newIndicesResolveClusterFunc(t),
			ResolveIndex:            newIndicesResolveIndexFunc(t),
			Rollover:                newIndicesRolloverFunc(t),
			Segments:                newIndicesSegmentsFunc(t),
			ShardStores:             newIndicesShardStoresFunc(t),
			Shrink:                  newIndicesShrinkFunc(t),
			SimulateIndexTemplate:   newIndicesSimulateIndexTemplateFunc(t),
			SimulateTemplate:        newIndicesSimulateTemplateFunc(t),
			Split:                   newIndicesSplitFunc(t),
			Stats:                   newIndicesStatsFunc(t),
			UpdateAliases:           newIndicesUpdateAliasesFunc(t),
			ValidateQuery:           newIndicesValidateQueryFunc(t),
		},
		Ingest: &Ingest{
			DeleteGeoipDatabase:      newIngestDeleteGeoipDatabaseFunc(t),
			DeleteIPLocationDatabase: newIngestDeleteIPLocationDatabaseFunc(t),
			DeletePipeline:           newIngestDeletePipelineFunc(t),
			GeoIPStats:               newIngestGeoIPStatsFunc(t),
			GetGeoipDatabase:         newIngestGetGeoipDatabaseFunc(t),
			GetIPLocationDatabase:    newIngestGetIPLocationDatabaseFunc(t),
			GetPipeline:              newIngestGetPipelineFunc(t),
			ProcessorGrok:            newIngestProcessorGrokFunc(t),
			PutGeoipDatabase:         newIngestPutGeoipDatabaseFunc(t),
			PutIPLocationDatabase:    newIngestPutIPLocationDatabaseFunc(t),
			PutPipeline:              newIngestPutPipelineFunc(t),
			Simulate:                 newIngestSimulateFunc(t),
		},
		Nodes: &Nodes{
			ClearRepositoriesMeteringArchive: newNodesClearRepositoriesMeteringArchiveFunc(t),
			GetRepositoriesMeteringInfo:      newNodesGetRepositoriesMeteringInfoFunc(t),
			HotThreads:                       newNodesHotThreadsFunc(t),
			Info:                             newNodesInfoFunc(t),
			ReloadSecureSettings:             newNodesReloadSecureSettingsFunc(t),
			Stats:                            newNodesStatsFunc(t),
			Usage:                            newNodesUsageFunc(t),
		},
		Remote: &Remote{},
		Snapshot: &Snapshot{
			CleanupRepository:         newSnapshotCleanupRepositoryFunc(t),
			Clone:                     newSnapshotCloneFunc(t),
			CreateRepository:          newSnapshotCreateRepositoryFunc(t),
			Create:                    newSnapshotCreateFunc(t),
			DeleteRepository:          newSnapshotDeleteRepositoryFunc(t),
			Delete:                    newSnapshotDeleteFunc(t),
			GetRepository:             newSnapshotGetRepositoryFunc(t),
			Get:                       newSnapshotGetFunc(t),
			RepositoryAnalyze:         newSnapshotRepositoryAnalyzeFunc(t),
			RepositoryVerifyIntegrity: newSnapshotRepositoryVerifyIntegrityFunc(t),
			Restore:                   newSnapshotRestoreFunc(t),
			Status:                    newSnapshotStatusFunc(t),
			VerifyRepository:          newSnapshotVerifyRepositoryFunc(t),
		},
		Tasks: &Tasks{
			Cancel: newTasksCancelFunc(t),
			Get:    newTasksGetFunc(t),
			List:   newTasksListFunc(t),
		},
		AsyncSearch: &AsyncSearch{
			Delete: newAsyncSearchDeleteFunc(t),
			Get:    newAsyncSearchGetFunc(t),
			Status: newAsyncSearchStatusFunc(t),
			Submit: newAsyncSearchSubmitFunc(t),
		},
		CCR: &CCR{
			DeleteAutoFollowPattern: newCCRDeleteAutoFollowPatternFunc(t),
			FollowInfo:              newCCRFollowInfoFunc(t),
			Follow:                  newCCRFollowFunc(t),
			FollowStats:             newCCRFollowStatsFunc(t),
			ForgetFollower:          newCCRForgetFollowerFunc(t),
			GetAutoFollowPattern:    newCCRGetAutoFollowPatternFunc(t),
			PauseAutoFollowPattern:  newCCRPauseAutoFollowPatternFunc(t),
			PauseFollow:             newCCRPauseFollowFunc(t),
			PutAutoFollowPattern:    newCCRPutAutoFollowPatternFunc(t),
			ResumeAutoFollowPattern: newCCRResumeAutoFollowPatternFunc(t),
			ResumeFollow:            newCCRResumeFollowFunc(t),
			Stats:                   newCCRStatsFunc(t),
			Unfollow:                newCCRUnfollowFunc(t),
		},
		ILM: &ILM{
			DeleteLifecycle:    newILMDeleteLifecycleFunc(t),
			ExplainLifecycle:   newILMExplainLifecycleFunc(t),
			GetLifecycle:       newILMGetLifecycleFunc(t),
			GetStatus:          newILMGetStatusFunc(t),
			MigrateToDataTiers: newILMMigrateToDataTiersFunc(t),
			MoveToStep:         newILMMoveToStepFunc(t),
			PutLifecycle:       newILMPutLifecycleFunc(t),
			RemovePolicy:       newILMRemovePolicyFunc(t),
			Retry:              newILMRetryFunc(t),
			Start:              newILMStartFunc(t),
			Stop:               newILMStopFunc(t),
		},
		License: &License{
			Delete:         newLicenseDeleteFunc(t),
			GetBasicStatus: newLicenseGetBasicStatusFunc(t),
			Get:            newLicenseGetFunc(t),
			GetTrialStatus: newLicenseGetTrialStatusFunc(t),
			Post:           newLicensePostFunc(t),
			PostStartBasic: newLicensePostStartBasicFunc(t),
			PostStartTrial: newLicensePostStartTrialFunc(t),
		},
		Migration: &Migration{
			Deprecations:            newMigrationDeprecationsFunc(t),
			GetFeatureUpgradeStatus: newMigrationGetFeatureUpgradeStatusFunc(t),
			PostFeatureUpgrade:      newMigrationPostFeatureUpgradeFunc(t),
		},
		ML: &ML{
			ClearTrainedModelDeploymentCache: newMLClearTrainedModelDeploymentCacheFunc(t),
			CloseJob:                         newMLCloseJobFunc(t),
			DeleteCalendarEvent:              newMLDeleteCalendarEventFunc(t),
			DeleteCalendarJob:                newMLDeleteCalendarJobFunc(t),
			DeleteCalendar:                   newMLDeleteCalendarFunc(t),
			DeleteDataFrameAnalytics:         newMLDeleteDataFrameAnalyticsFunc(t),
			DeleteDatafeed:                   newMLDeleteDatafeedFunc(t),
			DeleteExpiredData:                newMLDeleteExpiredDataFunc(t),
			DeleteFilter:                     newMLDeleteFilterFunc(t),
			DeleteForecast:                   newMLDeleteForecastFunc(t),
			DeleteJob:                        newMLDeleteJobFunc(t),
			DeleteModelSnapshot:              newMLDeleteModelSnapshotFunc(t),
			DeleteTrainedModelAlias:          newMLDeleteTrainedModelAliasFunc(t),
			DeleteTrainedModel:               newMLDeleteTrainedModelFunc(t),
			EstimateModelMemory:              newMLEstimateModelMemoryFunc(t),
			EvaluateDataFrame:                newMLEvaluateDataFrameFunc(t),
			ExplainDataFrameAnalytics:        newMLExplainDataFrameAnalyticsFunc(t),
			FlushJob:                         newMLFlushJobFunc(t),
			Forecast:                         newMLForecastFunc(t),
			GetBuckets:                       newMLGetBucketsFunc(t),
			GetCalendarEvents:                newMLGetCalendarEventsFunc(t),
			GetCalendars:                     newMLGetCalendarsFunc(t),
			GetCategories:                    newMLGetCategoriesFunc(t),
			GetDataFrameAnalytics:            newMLGetDataFrameAnalyticsFunc(t),
			GetDataFrameAnalyticsStats:       newMLGetDataFrameAnalyticsStatsFunc(t),
			GetDatafeedStats:                 newMLGetDatafeedStatsFunc(t),
			GetDatafeeds:                     newMLGetDatafeedsFunc(t),
			GetFilters:                       newMLGetFiltersFunc(t),
			GetInfluencers:                   newMLGetInfluencersFunc(t),
			GetJobStats:                      newMLGetJobStatsFunc(t),
			GetJobs:                          newMLGetJobsFunc(t),
			GetMemoryStats:                   newMLGetMemoryStatsFunc(t),
			GetModelSnapshotUpgradeStats:     newMLGetModelSnapshotUpgradeStatsFunc(t),
			GetModelSnapshots:                newMLGetModelSnapshotsFunc(t),
			GetOverallBuckets:                newMLGetOverallBucketsFunc(t),
			GetRecords:                       newMLGetRecordsFunc(t),
			GetTrainedModels:                 newMLGetTrainedModelsFunc(t),
			GetTrainedModelsStats:            newMLGetTrainedModelsStatsFunc(t),
			InferTrainedModel:                newMLInferTrainedModelFunc(t),
			Info:                             newMLInfoFunc(t),
			OpenJob:                          newMLOpenJobFunc(t),
			PostCalendarEvents:               newMLPostCalendarEventsFunc(t),
			PostData:                         newMLPostDataFunc(t),
			PreviewDataFrameAnalytics:        newMLPreviewDataFrameAnalyticsFunc(t),
			PreviewDatafeed:                  newMLPreviewDatafeedFunc(t),
			PutCalendarJob:                   newMLPutCalendarJobFunc(t),
			PutCalendar:                      newMLPutCalendarFunc(t),
			PutDataFrameAnalytics:            newMLPutDataFrameAnalyticsFunc(t),
			PutDatafeed:                      newMLPutDatafeedFunc(t),
			PutFilter:                        newMLPutFilterFunc(t),
			PutJob:                           newMLPutJobFunc(t),
			PutTrainedModelAlias:             newMLPutTrainedModelAliasFunc(t),
			PutTrainedModelDefinitionPart:    newMLPutTrainedModelDefinitionPartFunc(t),
			PutTrainedModel:                  newMLPutTrainedModelFunc(t),
			PutTrainedModelVocabulary:        newMLPutTrainedModelVocabularyFunc(t),
			ResetJob:                         newMLResetJobFunc(t),
			RevertModelSnapshot:              newMLRevertModelSnapshotFunc(t),
			SetUpgradeMode:                   newMLSetUpgradeModeFunc(t),
			StartDataFrameAnalytics:          newMLStartDataFrameAnalyticsFunc(t),
			StartDatafeed:                    newMLStartDatafeedFunc(t),
			StartTrainedModelDeployment:      newMLStartTrainedModelDeploymentFunc(t),
			StopDataFrameAnalytics:           newMLStopDataFrameAnalyticsFunc(t),
			StopDatafeed:                     newMLStopDatafeedFunc(t),
			StopTrainedModelDeployment:       newMLStopTrainedModelDeploymentFunc(t),
			UpdateDataFrameAnalytics:         newMLUpdateDataFrameAnalyticsFunc(t),
			UpdateDatafeed:                   newMLUpdateDatafeedFunc(t),
			UpdateFilter:                     newMLUpdateFilterFunc(t),
			UpdateJob:                        newMLUpdateJobFunc(t),
			UpdateModelSnapshot:              newMLUpdateModelSnapshotFunc(t),
			UpdateTrainedModelDeployment:     newMLUpdateTrainedModelDeploymentFunc(t),
			UpgradeJobSnapshot:               newMLUpgradeJobSnapshotFunc(t),
			ValidateDetector:                 newMLValidateDetectorFunc(t),
			Validate:                         newMLValidateFunc(t),
		},
		Monitoring: &Monitoring{
			Bulk: newMonitoringBulkFunc(t),
		},
		Rollup: &Rollup{
			DeleteJob:    newRollupDeleteJobFunc(t),
			GetJobs:      newRollupGetJobsFunc(t),
			GetCaps:      newRollupGetRollupCapsFunc(t),
			GetIndexCaps: newRollupGetRollupIndexCapsFunc(t),
			PutJob:       newRollupPutJobFunc(t),
			Search:       newRollupRollupSearchFunc(t),
			StartJob:     newRollupStartJobFunc(t),
			StopJob:      newRollupStopJobFunc(t),
		},
		Security: &Security{
			ActivateUserProfile:         newSecurityActivateUserProfileFunc(t),
			Authenticate:                newSecurityAuthenticateFunc(t),
			BulkDeleteRole:              newSecurityBulkDeleteRoleFunc(t),
			BulkPutRole:                 newSecurityBulkPutRoleFunc(t),
			BulkUpdateAPIKeys:           newSecurityBulkUpdateAPIKeysFunc(t),
			ChangePassword:              newSecurityChangePasswordFunc(t),
			ClearAPIKeyCache:            newSecurityClearAPIKeyCacheFunc(t),
			ClearCachedPrivileges:       newSecurityClearCachedPrivilegesFunc(t),
			ClearCachedRealms:           newSecurityClearCachedRealmsFunc(t),
			ClearCachedRoles:            newSecurityClearCachedRolesFunc(t),
			ClearCachedServiceTokens:    newSecurityClearCachedServiceTokensFunc(t),
			CreateAPIKey:                newSecurityCreateAPIKeyFunc(t),
			CreateCrossClusterAPIKey:    newSecurityCreateCrossClusterAPIKeyFunc(t),
			CreateServiceToken:          newSecurityCreateServiceTokenFunc(t),
			DelegatePki:                 newSecurityDelegatePkiFunc(t),
			DeletePrivileges:            newSecurityDeletePrivilegesFunc(t),
			DeleteRoleMapping:           newSecurityDeleteRoleMappingFunc(t),
			DeleteRole:                  newSecurityDeleteRoleFunc(t),
			DeleteServiceToken:          newSecurityDeleteServiceTokenFunc(t),
			DeleteUser:                  newSecurityDeleteUserFunc(t),
			DisableUserProfile:          newSecurityDisableUserProfileFunc(t),
			DisableUser:                 newSecurityDisableUserFunc(t),
			EnableUserProfile:           newSecurityEnableUserProfileFunc(t),
			EnableUser:                  newSecurityEnableUserFunc(t),
			EnrollKibana:                newSecurityEnrollKibanaFunc(t),
			EnrollNode:                  newSecurityEnrollNodeFunc(t),
			GetAPIKey:                   newSecurityGetAPIKeyFunc(t),
			GetBuiltinPrivileges:        newSecurityGetBuiltinPrivilegesFunc(t),
			GetPrivileges:               newSecurityGetPrivilegesFunc(t),
			GetRoleMapping:              newSecurityGetRoleMappingFunc(t),
			GetRole:                     newSecurityGetRoleFunc(t),
			GetServiceAccounts:          newSecurityGetServiceAccountsFunc(t),
			GetServiceCredentials:       newSecurityGetServiceCredentialsFunc(t),
			GetSettings:                 newSecurityGetSettingsFunc(t),
			GetToken:                    newSecurityGetTokenFunc(t),
			GetUserPrivileges:           newSecurityGetUserPrivilegesFunc(t),
			GetUserProfile:              newSecurityGetUserProfileFunc(t),
			GetUser:                     newSecurityGetUserFunc(t),
			GrantAPIKey:                 newSecurityGrantAPIKeyFunc(t),
			HasPrivileges:               newSecurityHasPrivilegesFunc(t),
			HasPrivilegesUserProfile:    newSecurityHasPrivilegesUserProfileFunc(t),
			InvalidateAPIKey:            newSecurityInvalidateAPIKeyFunc(t),
			InvalidateToken:             newSecurityInvalidateTokenFunc(t),
			OidcAuthenticate:            newSecurityOidcAuthenticateFunc(t),
			OidcLogout:                  newSecurityOidcLogoutFunc(t),
			OidcPrepareAuthentication:   newSecurityOidcPrepareAuthenticationFunc(t),
			PutPrivileges:               newSecurityPutPrivilegesFunc(t),
			PutRoleMapping:              newSecurityPutRoleMappingFunc(t),
			PutRole:                     newSecurityPutRoleFunc(t),
			PutUser:                     newSecurityPutUserFunc(t),
			QueryAPIKeys:                newSecurityQueryAPIKeysFunc(t),
			QueryRole:                   newSecurityQueryRoleFunc(t),
			QueryUser:                   newSecurityQueryUserFunc(t),
			SamlAuthenticate:            newSecuritySamlAuthenticateFunc(t),
			SamlCompleteLogout:          newSecuritySamlCompleteLogoutFunc(t),
			SamlInvalidate:              newSecuritySamlInvalidateFunc(t),
			SamlLogout:                  newSecuritySamlLogoutFunc(t),
			SamlPrepareAuthentication:   newSecuritySamlPrepareAuthenticationFunc(t),
			SamlServiceProviderMetadata: newSecuritySamlServiceProviderMetadataFunc(t),
			SuggestUserProfiles:         newSecuritySuggestUserProfilesFunc(t),
			UpdateAPIKey:                newSecurityUpdateAPIKeyFunc(t),
			UpdateCrossClusterAPIKey:    newSecurityUpdateCrossClusterAPIKeyFunc(t),
			UpdateSettings:              newSecurityUpdateSettingsFunc(t),
			UpdateUserProfileData:       newSecurityUpdateUserProfileDataFunc(t),
		},
		SQL: &SQL{
			ClearCursor:    newSQLClearCursorFunc(t),
			DeleteAsync:    newSQLDeleteAsyncFunc(t),
			GetAsync:       newSQLGetAsyncFunc(t),
			GetAsyncStatus: newSQLGetAsyncStatusFunc(t),
			Query:          newSQLQueryFunc(t),
			Translate:      newSQLTranslateFunc(t),
		},
		SSL: &SSL{
			Certificates: newSSLCertificatesFunc(t),
		},
		Watcher: &Watcher{
			AckWatch:        newWatcherAckWatchFunc(t),
			ActivateWatch:   newWatcherActivateWatchFunc(t),
			DeactivateWatch: newWatcherDeactivateWatchFunc(t),
			DeleteWatch:     newWatcherDeleteWatchFunc(t),
			ExecuteWatch:    newWatcherExecuteWatchFunc(t),
			GetSettings:     newWatcherGetSettingsFunc(t),
			GetWatch:        newWatcherGetWatchFunc(t),
			PutWatch:        newWatcherPutWatchFunc(t),
			QueryWatches:    newWatcherQueryWatchesFunc(t),
			Start:           newWatcherStartFunc(t),
			Stats:           newWatcherStatsFunc(t),
			Stop:            newWatcherStopFunc(t),
			UpdateSettings:  newWatcherUpdateSettingsFunc(t),
		},
		XPack: &XPack{
			Info:  newXPackInfoFunc(t),
			Usage: newXPackUsageFunc(t),
		},
	}
}
