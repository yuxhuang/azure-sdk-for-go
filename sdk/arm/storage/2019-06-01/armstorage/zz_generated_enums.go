// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

// AccessTier - Required for storage accounts where kind = BlobStorage. The access tier used for billing.
type AccessTier string

const (
	AccessTierHot  AccessTier = "Hot"
	AccessTierCool AccessTier = "Cool"
)

func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{
		AccessTierHot,
		AccessTierCool,
	}
}

func (c AccessTier) ToPtr() *AccessTier {
	return &c
}

// AccountStatus - Gets the status indicating whether the primary location of the storage account is available or unavailable.
type AccountStatus string

const (
	AccountStatusAvailable   AccountStatus = "available"
	AccountStatusUnavailable AccountStatus = "unavailable"
)

func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{
		AccountStatusAvailable,
		AccountStatusUnavailable,
	}
}

func (c AccountStatus) ToPtr() *AccountStatus {
	return &c
}

// BlobRestoreProgressStatus - The status of blob restore progress. Possible values are: - InProgress: Indicates that blob restore is ongoing. - Complete:
// Indicates that blob restore has been completed successfully. - Failed:
// Indicates that blob restore is failed.
type BlobRestoreProgressStatus string

const (
	BlobRestoreProgressStatusComplete   BlobRestoreProgressStatus = "Complete"
	BlobRestoreProgressStatusFailed     BlobRestoreProgressStatus = "Failed"
	BlobRestoreProgressStatusInProgress BlobRestoreProgressStatus = "InProgress"
)

func PossibleBlobRestoreProgressStatusValues() []BlobRestoreProgressStatus {
	return []BlobRestoreProgressStatus{
		BlobRestoreProgressStatusComplete,
		BlobRestoreProgressStatusFailed,
		BlobRestoreProgressStatusInProgress,
	}
}

func (c BlobRestoreProgressStatus) ToPtr() *BlobRestoreProgressStatus {
	return &c
}

// Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices
// (For example, "Logging, Metrics"), or None to bypass none
// of those traffics.
type Bypass string

const (
	BypassAzureServices Bypass = "AzureServices"
	BypassLogging       Bypass = "Logging"
	BypassMetrics       Bypass = "Metrics"
	BypassNone          Bypass = "None"
)

func PossibleBypassValues() []Bypass {
	return []Bypass{
		BypassAzureServices,
		BypassLogging,
		BypassMetrics,
		BypassNone,
	}
}

func (c Bypass) ToPtr() *Bypass {
	return &c
}

type CorsRuleAllowedMethodsItem string

const (
	CorsRuleAllowedMethodsItemDelete  CorsRuleAllowedMethodsItem = "DELETE"
	CorsRuleAllowedMethodsItemGet     CorsRuleAllowedMethodsItem = "GET"
	CorsRuleAllowedMethodsItemHead    CorsRuleAllowedMethodsItem = "HEAD"
	CorsRuleAllowedMethodsItemMerge   CorsRuleAllowedMethodsItem = "MERGE"
	CorsRuleAllowedMethodsItemOptions CorsRuleAllowedMethodsItem = "OPTIONS"
	CorsRuleAllowedMethodsItemPost    CorsRuleAllowedMethodsItem = "POST"
	CorsRuleAllowedMethodsItemPut     CorsRuleAllowedMethodsItem = "PUT"
)

func PossibleCorsRuleAllowedMethodsItemValues() []CorsRuleAllowedMethodsItem {
	return []CorsRuleAllowedMethodsItem{
		CorsRuleAllowedMethodsItemDelete,
		CorsRuleAllowedMethodsItemGet,
		CorsRuleAllowedMethodsItemHead,
		CorsRuleAllowedMethodsItemMerge,
		CorsRuleAllowedMethodsItemOptions,
		CorsRuleAllowedMethodsItemPost,
		CorsRuleAllowedMethodsItemPut,
	}
}

func (c CorsRuleAllowedMethodsItem) ToPtr() *CorsRuleAllowedMethodsItem {
	return &c
}

// DefaultAction - Specifies the default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// DirectoryServiceOptions - Indicates the directory service used.
type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsAadds DirectoryServiceOptions = "AADDS"
	DirectoryServiceOptionsAd    DirectoryServiceOptions = "AD"
	DirectoryServiceOptionsNone  DirectoryServiceOptions = "None"
)

func PossibleDirectoryServiceOptionsValues() []DirectoryServiceOptions {
	return []DirectoryServiceOptions{
		DirectoryServiceOptionsAadds,
		DirectoryServiceOptionsAd,
		DirectoryServiceOptionsNone,
	}
}

func (c DirectoryServiceOptions) ToPtr() *DirectoryServiceOptions {
	return &c
}

// EnabledProtocols - The authentication protocol that is used for the file share. Can only be specified when creating a share.
type EnabledProtocols string

const (
	EnabledProtocolsNfs EnabledProtocols = "NFS"
	EnabledProtocolsSmb EnabledProtocols = "SMB"
)

func PossibleEnabledProtocolsValues() []EnabledProtocols {
	return []EnabledProtocols{
		EnabledProtocolsNfs,
		EnabledProtocolsSmb,
	}
}

func (c EnabledProtocols) ToPtr() *EnabledProtocols {
	return &c
}

// EncryptionScopeSource - The provider for the encryption scope. Possible values (case-insensitive): Microsoft.Storage, Microsoft.KeyVault.
type EncryptionScopeSource string

const (
	EncryptionScopeSourceMicrosoftKeyVault EncryptionScopeSource = "Microsoft.KeyVault"
	EncryptionScopeSourceMicrosoftStorage  EncryptionScopeSource = "Microsoft.Storage"
)

func PossibleEncryptionScopeSourceValues() []EncryptionScopeSource {
	return []EncryptionScopeSource{
		EncryptionScopeSourceMicrosoftKeyVault,
		EncryptionScopeSourceMicrosoftStorage,
	}
}

func (c EncryptionScopeSource) ToPtr() *EncryptionScopeSource {
	return &c
}

// EncryptionScopeState - The state of the encryption scope. Possible values (case-insensitive): Enabled, Disabled.
type EncryptionScopeState string

const (
	EncryptionScopeStateDisabled EncryptionScopeState = "Disabled"
	EncryptionScopeStateEnabled  EncryptionScopeState = "Enabled"
)

func PossibleEncryptionScopeStateValues() []EncryptionScopeState {
	return []EncryptionScopeState{
		EncryptionScopeStateDisabled,
		EncryptionScopeStateEnabled,
	}
}

func (c EncryptionScopeState) ToPtr() *EncryptionScopeState {
	return &c
}

// GeoReplicationStatus - The status of the secondary location. Possible values are: - Live: Indicates that the secondary location is active and operational.
// - Bootstrap: Indicates initial synchronization from the primary
// location to the secondary location is in progress.This typically occurs when replication is first enabled. - Unavailable: Indicates that the secondary
// location is temporarily unavailable.
type GeoReplicationStatus string

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = "Bootstrap"
	GeoReplicationStatusLive        GeoReplicationStatus = "Live"
	GeoReplicationStatusUnavailable GeoReplicationStatus = "Unavailable"
)

func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return []GeoReplicationStatus{
		GeoReplicationStatusBootstrap,
		GeoReplicationStatusLive,
		GeoReplicationStatusUnavailable,
	}
}

func (c GeoReplicationStatus) ToPtr() *GeoReplicationStatus {
	return &c
}

// HTTPProtocol - The protocol permitted for a request made with the account SAS.
type HTTPProtocol string

const (
	HTTPProtocolHTTPsHttp HTTPProtocol = "https,http"
	HTTPProtocolHTTPs     HTTPProtocol = "https"
)

func PossibleHTTPProtocolValues() []HTTPProtocol {
	return []HTTPProtocol{
		HTTPProtocolHTTPsHttp,
		HTTPProtocolHTTPs,
	}
}

func (c HTTPProtocol) ToPtr() *HTTPProtocol {
	return &c
}

// ImmutabilityPolicyState - The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
type ImmutabilityPolicyState string

const (
	ImmutabilityPolicyStateLocked   ImmutabilityPolicyState = "Locked"
	ImmutabilityPolicyStateUnlocked ImmutabilityPolicyState = "Unlocked"
)

func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return []ImmutabilityPolicyState{
		ImmutabilityPolicyStateLocked,
		ImmutabilityPolicyStateUnlocked,
	}
}

func (c ImmutabilityPolicyState) ToPtr() *ImmutabilityPolicyState {
	return &c
}

// ImmutabilityPolicyUpdateType - The ImmutabilityPolicy update type of a blob container, possible values include: put, lock and extend.
type ImmutabilityPolicyUpdateType string

const (
	ImmutabilityPolicyUpdateTypeExtend ImmutabilityPolicyUpdateType = "extend"
	ImmutabilityPolicyUpdateTypeLock   ImmutabilityPolicyUpdateType = "lock"
	ImmutabilityPolicyUpdateTypePut    ImmutabilityPolicyUpdateType = "put"
)

func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return []ImmutabilityPolicyUpdateType{
		ImmutabilityPolicyUpdateTypeExtend,
		ImmutabilityPolicyUpdateTypeLock,
		ImmutabilityPolicyUpdateTypePut,
	}
}

func (c ImmutabilityPolicyUpdateType) ToPtr() *ImmutabilityPolicyUpdateType {
	return &c
}

// KeyPermission - Permissions for the key -- read-only or full permissions.
type KeyPermission string

const (
	KeyPermissionRead KeyPermission = "Read"
	KeyPermissionFull KeyPermission = "Full"
)

func PossibleKeyPermissionValues() []KeyPermission {
	return []KeyPermission{
		KeyPermissionRead,
		KeyPermissionFull,
	}
}

func (c KeyPermission) ToPtr() *KeyPermission {
	return &c
}

// KeySource - The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Storage, Microsoft.Keyvault
type KeySource string

const (
	KeySourceMicrosoftKeyvault KeySource = "Microsoft.Keyvault"
	KeySourceMicrosoftStorage  KeySource = "Microsoft.Storage"
)

func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftKeyvault,
		KeySourceMicrosoftStorage,
	}
}

func (c KeySource) ToPtr() *KeySource {
	return &c
}

// KeyType - Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped encryption key will be used. 'Service'
// key type implies that a default service key is used.
type KeyType string

const (
	KeyTypeAccount KeyType = "Account"
	KeyTypeService KeyType = "Service"
)

func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypeAccount,
		KeyTypeService,
	}
}

func (c KeyType) ToPtr() *KeyType {
	return &c
}

// Kind - Indicates the type of storage account.
type Kind string

const (
	KindBlobStorage      Kind = "BlobStorage"
	KindBlockBlobStorage Kind = "BlockBlobStorage"
	KindFileStorage      Kind = "FileStorage"
	KindStorage          Kind = "Storage"
	KindStorageV2        Kind = "StorageV2"
)

func PossibleKindValues() []Kind {
	return []Kind{
		KindBlobStorage,
		KindBlockBlobStorage,
		KindFileStorage,
		KindStorage,
		KindStorageV2,
	}
}

func (c Kind) ToPtr() *Kind {
	return &c
}

// LargeFileSharesState - Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled LargeFileSharesState = "Disabled"
	LargeFileSharesStateEnabled  LargeFileSharesState = "Enabled"
)

func PossibleLargeFileSharesStateValues() []LargeFileSharesState {
	return []LargeFileSharesState{
		LargeFileSharesStateDisabled,
		LargeFileSharesStateEnabled,
	}
}

func (c LargeFileSharesState) ToPtr() *LargeFileSharesState {
	return &c
}

// LeaseContainerRequestAction - Specifies the lease action. Can be one of the available actions.
type LeaseContainerRequestAction string

const (
	LeaseContainerRequestActionAcquire LeaseContainerRequestAction = "Acquire"
	LeaseContainerRequestActionBreak   LeaseContainerRequestAction = "Break"
	LeaseContainerRequestActionChange  LeaseContainerRequestAction = "Change"
	LeaseContainerRequestActionRelease LeaseContainerRequestAction = "Release"
	LeaseContainerRequestActionRenew   LeaseContainerRequestAction = "Renew"
)

func PossibleLeaseContainerRequestActionValues() []LeaseContainerRequestAction {
	return []LeaseContainerRequestAction{
		LeaseContainerRequestActionAcquire,
		LeaseContainerRequestActionBreak,
		LeaseContainerRequestActionChange,
		LeaseContainerRequestActionRelease,
		LeaseContainerRequestActionRenew,
	}
}

func (c LeaseContainerRequestAction) ToPtr() *LeaseContainerRequestAction {
	return &c
}

// LeaseDuration - Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

func PossibleLeaseDurationValues() []LeaseDuration {
	return []LeaseDuration{
		LeaseDurationFixed,
		LeaseDurationInfinite,
	}
}

func (c LeaseDuration) ToPtr() *LeaseDuration {
	return &c
}

// LeaseState - Lease state of the container.
type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

func PossibleLeaseStateValues() []LeaseState {
	return []LeaseState{
		LeaseStateAvailable,
		LeaseStateBreaking,
		LeaseStateBroken,
		LeaseStateExpired,
		LeaseStateLeased,
	}
}

func (c LeaseState) ToPtr() *LeaseState {
	return &c
}

// LeaseStatus - The lease status of the container.
type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

func PossibleLeaseStatusValues() []LeaseStatus {
	return []LeaseStatus{
		LeaseStatusLocked,
		LeaseStatusUnlocked,
	}
}

func (c LeaseStatus) ToPtr() *LeaseStatus {
	return &c
}

type ListContainersInclude string

const (
	ListContainersIncludeDeleted ListContainersInclude = "deleted"
)

func PossibleListContainersIncludeValues() []ListContainersInclude {
	return []ListContainersInclude{
		ListContainersIncludeDeleted,
	}
}

func (c ListContainersInclude) ToPtr() *ListContainersInclude {
	return &c
}

type ManagementPolicyName string

const (
	ManagementPolicyNameDefault ManagementPolicyName = "default"
)

func PossibleManagementPolicyNameValues() []ManagementPolicyName {
	return []ManagementPolicyName{
		ManagementPolicyNameDefault,
	}
}

func (c ManagementPolicyName) ToPtr() *ManagementPolicyName {
	return &c
}

// MinimumTLSVersion - Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
type MinimumTLSVersion string

const (
	MinimumTLSVersionTLS10 MinimumTLSVersion = "TLS1_0"
	MinimumTLSVersionTLS11 MinimumTLSVersion = "TLS1_1"
	MinimumTLSVersionTLS12 MinimumTLSVersion = "TLS1_2"
)

func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return []MinimumTLSVersion{
		MinimumTLSVersionTLS10,
		MinimumTLSVersionTLS11,
		MinimumTLSVersionTLS12,
	}
}

func (c MinimumTLSVersion) ToPtr() *MinimumTLSVersion {
	return &c
}

// Permissions - The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update
// (u) and Process (p).
type Permissions string

const (
	PermissionsA Permissions = "a"
	PermissionsC Permissions = "c"
	PermissionsD Permissions = "d"
	PermissionsL Permissions = "l"
	PermissionsP Permissions = "p"
	PermissionsR Permissions = "r"
	PermissionsU Permissions = "u"
	PermissionsW Permissions = "w"
)

func PossiblePermissionsValues() []Permissions {
	return []Permissions{
		PermissionsA,
		PermissionsC,
		PermissionsD,
		PermissionsL,
		PermissionsP,
		PermissionsR,
		PermissionsU,
		PermissionsW,
	}
}

func (c Permissions) ToPtr() *Permissions {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

func (c PrivateEndpointServiceConnectionStatus) ToPtr() *PrivateEndpointServiceConnectionStatus {
	return &c
}

// ProvisioningState - Gets the status of the storage account at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateResolvingDNS ProvisioningState = "ResolvingDNS"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateResolvingDNS,
		ProvisioningStateSucceeded,
	}
}

func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// PublicAccess - Specifies whether data in the container may be accessed publicly and the level of access.
type PublicAccess string

const (
	PublicAccessContainer PublicAccess = "Container"
	PublicAccessBlob      PublicAccess = "Blob"
	PublicAccessNone      PublicAccess = "None"
)

func PossiblePublicAccessValues() []PublicAccess {
	return []PublicAccess{
		PublicAccessContainer,
		PublicAccessBlob,
		PublicAccessNone,
	}
}

func (c PublicAccess) ToPtr() *PublicAccess {
	return &c
}

// Reason - Gets the reason that a storage account name could not be used. The Reason element is only returned if NameAvailable is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

func (c Reason) ToPtr() *Reason {
	return &c
}

// ReasonCode - The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". Quota Id is set when the SKU has requiredQuotas
// parameter as the subscription does not belong to that
// quota. The "NotAvailableForSubscription" is related to capacity at DC.
type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaID                     ReasonCode = "QuotaId"
)

func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeQuotaID,
	}
}

func (c ReasonCode) ToPtr() *ReasonCode {
	return &c
}

// RootSquashType - The property is for NFS share only. The default is NoRootSquash.
type RootSquashType string

const (
	RootSquashTypeAllSquash    RootSquashType = "AllSquash"
	RootSquashTypeNoRootSquash RootSquashType = "NoRootSquash"
	RootSquashTypeRootSquash   RootSquashType = "RootSquash"
)

func PossibleRootSquashTypeValues() []RootSquashType {
	return []RootSquashType{
		RootSquashTypeAllSquash,
		RootSquashTypeNoRootSquash,
		RootSquashTypeRootSquash,
	}
}

func (c RootSquashType) ToPtr() *RootSquashType {
	return &c
}

// RoutingChoice - Routing Choice defines the kind of network routing opted by the user.
type RoutingChoice string

const (
	RoutingChoiceInternetRouting  RoutingChoice = "InternetRouting"
	RoutingChoiceMicrosoftRouting RoutingChoice = "MicrosoftRouting"
)

func PossibleRoutingChoiceValues() []RoutingChoice {
	return []RoutingChoice{
		RoutingChoiceInternetRouting,
		RoutingChoiceMicrosoftRouting,
	}
}

func (c RoutingChoice) ToPtr() *RoutingChoice {
	return &c
}

// RuleType - The valid value is Lifecycle
type RuleType string

const (
	RuleTypeLifecycle RuleType = "Lifecycle"
)

func PossibleRuleTypeValues() []RuleType {
	return []RuleType{
		RuleTypeLifecycle,
	}
}

func (c RuleType) ToPtr() *RuleType {
	return &c
}

// SKUName - The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called accountType.
type SKUName string

const (
	SKUNamePremiumLrs     SKUName = "Premium_LRS"
	SKUNamePremiumZrs     SKUName = "Premium_ZRS"
	SKUNameStandardGrs    SKUName = "Standard_GRS"
	SKUNameStandardGzrs   SKUName = "Standard_GZRS"
	SKUNameStandardLrs    SKUName = "Standard_LRS"
	SKUNameStandardRagrs  SKUName = "Standard_RAGRS"
	SKUNameStandardRagzrs SKUName = "Standard_RAGZRS"
	SKUNameStandardZrs    SKUName = "Standard_ZRS"
)

func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremiumLrs,
		SKUNamePremiumZrs,
		SKUNameStandardGrs,
		SKUNameStandardGzrs,
		SKUNameStandardLrs,
		SKUNameStandardRagrs,
		SKUNameStandardRagzrs,
		SKUNameStandardZrs,
	}
}

func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - The SKU tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierStandard,
		SKUTierPremium,
	}
}

func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// Services - The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
type Services string

const (
	ServicesB Services = "b"
	ServicesF Services = "f"
	ServicesQ Services = "q"
	ServicesT Services = "t"
)

func PossibleServicesValues() []Services {
	return []Services{
		ServicesB,
		ServicesF,
		ServicesQ,
		ServicesT,
	}
}

func (c Services) ToPtr() *Services {
	return &c
}

// ShareAccessTier - Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account
// can choose Premium.
type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierPremium              ShareAccessTier = "Premium"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

func PossibleShareAccessTierValues() []ShareAccessTier {
	return []ShareAccessTier{
		ShareAccessTierCool,
		ShareAccessTierHot,
		ShareAccessTierPremium,
		ShareAccessTierTransactionOptimized,
	}
}

func (c ShareAccessTier) ToPtr() *ShareAccessTier {
	return &c
}

// SignedResource - The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c), File (f), Share (s).
type SignedResource string

const (
	SignedResourceB SignedResource = "b"
	SignedResourceC SignedResource = "c"
	SignedResourceF SignedResource = "f"
	SignedResourceS SignedResource = "s"
)

func PossibleSignedResourceValues() []SignedResource {
	return []SignedResource{
		SignedResourceB,
		SignedResourceC,
		SignedResourceF,
		SignedResourceS,
	}
}

func (c SignedResource) ToPtr() *SignedResource {
	return &c
}

// SignedResourceTypes - The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access
// to container-level APIs; Object (o): Access to object-level APIs
// for blobs, queue messages, table entities, and files.
type SignedResourceTypes string

const (
	SignedResourceTypesC SignedResourceTypes = "c"
	SignedResourceTypesO SignedResourceTypes = "o"
	SignedResourceTypesS SignedResourceTypes = "s"
)

func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return []SignedResourceTypes{
		SignedResourceTypesC,
		SignedResourceTypesO,
		SignedResourceTypesS,
	}
}

func (c SignedResourceTypes) ToPtr() *SignedResourceTypes {
	return &c
}

// State - Gets the state of virtual network rule.
type State string

const (
	StateProvisioning         State = "provisioning"
	StateDeprovisioning       State = "deprovisioning"
	StateSucceeded            State = "succeeded"
	StateFailed               State = "failed"
	StateNetworkSourceDeleted State = "networkSourceDeleted"
)

func PossibleStateValues() []State {
	return []State{
		StateProvisioning,
		StateDeprovisioning,
		StateSucceeded,
		StateFailed,
		StateNetworkSourceDeleted,
	}
}

func (c State) ToPtr() *State {
	return &c
}

type StorageAccountExpand string

const (
	StorageAccountExpandGeoReplicationStats StorageAccountExpand = "geoReplicationStats"
	StorageAccountExpandBlobRestoreStatus   StorageAccountExpand = "blobRestoreStatus"
)

func PossibleStorageAccountExpandValues() []StorageAccountExpand {
	return []StorageAccountExpand{
		StorageAccountExpandGeoReplicationStats,
		StorageAccountExpandBlobRestoreStatus,
	}
}

func (c StorageAccountExpand) ToPtr() *StorageAccountExpand {
	return &c
}

// UsageUnit - Gets the unit of measurement.
type UsageUnit string

const (
	UsageUnitCount           UsageUnit = "Count"
	UsageUnitBytes           UsageUnit = "Bytes"
	UsageUnitSeconds         UsageUnit = "Seconds"
	UsageUnitPercent         UsageUnit = "Percent"
	UsageUnitCountsPerSecond UsageUnit = "CountsPerSecond"
	UsageUnitBytesPerSecond  UsageUnit = "BytesPerSecond"
)

func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
		UsageUnitBytes,
		UsageUnitSeconds,
		UsageUnitPercent,
		UsageUnitCountsPerSecond,
		UsageUnitBytesPerSecond,
	}
}

func (c UsageUnit) ToPtr() *UsageUnit {
	return &c
}
