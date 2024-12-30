// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by Cuelang code gen. DO NOT EDIT!
// Cuelang source: github.com/specterops/bloodhound/-/tree/main/packages/cue/schemas/

package ad

import (
	"errors"

	graph "github.com/specterops/bloodhound/dawgs/graph"
)

var (
	Entity                      = graph.StringKind("Base")
	User                        = graph.StringKind("User")
	Computer                    = graph.StringKind("Computer")
	Group                       = graph.StringKind("Group")
	GPO                         = graph.StringKind("GPO")
	OU                          = graph.StringKind("OU")
	Container                   = graph.StringKind("Container")
	Domain                      = graph.StringKind("Domain")
	LocalGroup                  = graph.StringKind("ADLocalGroup")
	LocalUser                   = graph.StringKind("ADLocalUser")
	AIACA                       = graph.StringKind("AIACA")
	RootCA                      = graph.StringKind("RootCA")
	EnterpriseCA                = graph.StringKind("EnterpriseCA")
	NTAuthStore                 = graph.StringKind("NTAuthStore")
	CertTemplate                = graph.StringKind("CertTemplate")
	IssuancePolicy              = graph.StringKind("IssuancePolicy")
	Owns                        = graph.StringKind("Owns")
	GenericAll                  = graph.StringKind("GenericAll")
	GenericWrite                = graph.StringKind("GenericWrite")
	WriteOwner                  = graph.StringKind("WriteOwner")
	WriteDACL                   = graph.StringKind("WriteDacl")
	MemberOf                    = graph.StringKind("MemberOf")
	ForceChangePassword         = graph.StringKind("ForceChangePassword")
	AllExtendedRights           = graph.StringKind("AllExtendedRights")
	AddMember                   = graph.StringKind("AddMember")
	HasSession                  = graph.StringKind("HasSession")
	Contains                    = graph.StringKind("Contains")
	GPLink                      = graph.StringKind("GPLink")
	AllowedToDelegate           = graph.StringKind("AllowedToDelegate")
	CoerceToTGT                 = graph.StringKind("CoerceToTGT")
	GetChanges                  = graph.StringKind("GetChanges")
	GetChangesAll               = graph.StringKind("GetChangesAll")
	GetChangesInFilteredSet     = graph.StringKind("GetChangesInFilteredSet")
	TrustedBy                   = graph.StringKind("TrustedBy")
	AllowedToAct                = graph.StringKind("AllowedToAct")
	AdminTo                     = graph.StringKind("AdminTo")
	CanPSRemote                 = graph.StringKind("CanPSRemote")
	CanRDP                      = graph.StringKind("CanRDP")
	ExecuteDCOM                 = graph.StringKind("ExecuteDCOM")
	HasSIDHistory               = graph.StringKind("HasSIDHistory")
	AddSelf                     = graph.StringKind("AddSelf")
	DCSync                      = graph.StringKind("DCSync")
	ReadLAPSPassword            = graph.StringKind("ReadLAPSPassword")
	ReadGMSAPassword            = graph.StringKind("ReadGMSAPassword")
	DumpSMSAPassword            = graph.StringKind("DumpSMSAPassword")
	SQLAdmin                    = graph.StringKind("SQLAdmin")
	AddAllowedToAct             = graph.StringKind("AddAllowedToAct")
	WriteSPN                    = graph.StringKind("WriteSPN")
	AddKeyCredentialLink        = graph.StringKind("AddKeyCredentialLink")
	LocalToComputer             = graph.StringKind("LocalToComputer")
	MemberOfLocalGroup          = graph.StringKind("MemberOfLocalGroup")
	RemoteInteractiveLogonRight = graph.StringKind("RemoteInteractiveLogonRight")
	SyncLAPSPassword            = graph.StringKind("SyncLAPSPassword")
	WriteAccountRestrictions    = graph.StringKind("WriteAccountRestrictions")
	WriteGPLink                 = graph.StringKind("WriteGPLink")
	RootCAFor                   = graph.StringKind("RootCAFor")
	DCFor                       = graph.StringKind("DCFor")
	PublishedTo                 = graph.StringKind("PublishedTo")
	ManageCertificates          = graph.StringKind("ManageCertificates")
	ManageCA                    = graph.StringKind("ManageCA")
	DelegatedEnrollmentAgent    = graph.StringKind("DelegatedEnrollmentAgent")
	Enroll                      = graph.StringKind("Enroll")
	HostsCAService              = graph.StringKind("HostsCAService")
	WritePKIEnrollmentFlag      = graph.StringKind("WritePKIEnrollmentFlag")
	WritePKINameFlag            = graph.StringKind("WritePKINameFlag")
	NTAuthStoreFor              = graph.StringKind("NTAuthStoreFor")
	TrustedForNTAuth            = graph.StringKind("TrustedForNTAuth")
	EnterpriseCAFor             = graph.StringKind("EnterpriseCAFor")
	IssuedSignedBy              = graph.StringKind("IssuedSignedBy")
	GoldenCert                  = graph.StringKind("GoldenCert")
	EnrollOnBehalfOf            = graph.StringKind("EnrollOnBehalfOf")
	OIDGroupLink                = graph.StringKind("OIDGroupLink")
	ExtendedByPolicy            = graph.StringKind("ExtendedByPolicy")
	ADCSESC1                    = graph.StringKind("ADCSESC1")
	ADCSESC3                    = graph.StringKind("ADCSESC3")
	ADCSESC4                    = graph.StringKind("ADCSESC4")
	ADCSESC6a                   = graph.StringKind("ADCSESC6a")
	ADCSESC6b                   = graph.StringKind("ADCSESC6b")
	ADCSESC9a                   = graph.StringKind("ADCSESC9a")
	ADCSESC9b                   = graph.StringKind("ADCSESC9b")
	ADCSESC10a                  = graph.StringKind("ADCSESC10a")
	ADCSESC10b                  = graph.StringKind("ADCSESC10b")
	ADCSESC13                   = graph.StringKind("ADCSESC13")
	SyncedToEntraUser           = graph.StringKind("SyncedToEntraUser")
	CoerceAndRelayNTLMToSMB     = graph.StringKind("CoerceAndRelayNTLMToSMB")
)

type Property string

const (
	AdminCount                              Property = "admincount"
	CASecurityCollected                     Property = "casecuritycollected"
	CAName                                  Property = "caname"
	CertChain                               Property = "certchain"
	CertName                                Property = "certname"
	CertThumbprint                          Property = "certthumbprint"
	CertThumbprints                         Property = "certthumbprints"
	HasEnrollmentAgentRestrictions          Property = "hasenrollmentagentrestrictions"
	EnrollmentAgentRestrictionsCollected    Property = "enrollmentagentrestrictionscollected"
	IsUserSpecifiesSanEnabled               Property = "isuserspecifiessanenabled"
	IsUserSpecifiesSanEnabledCollected      Property = "isuserspecifiessanenabledcollected"
	RoleSeparationEnabled                   Property = "roleseparationenabled"
	RoleSeparationEnabledCollected          Property = "roleseparationenabledcollected"
	HasBasicConstraints                     Property = "hasbasicconstraints"
	BasicConstraintPathLength               Property = "basicconstraintpathlength"
	UnresolvedPublishedTemplates            Property = "unresolvedpublishedtemplates"
	DNSHostname                             Property = "dnshostname"
	CrossCertificatePair                    Property = "crosscertificatepair"
	DistinguishedName                       Property = "distinguishedname"
	DomainFQDN                              Property = "domain"
	DomainSID                               Property = "domainsid"
	Sensitive                               Property = "sensitive"
	HighValue                               Property = "highvalue"
	BlocksInheritance                       Property = "blocksinheritance"
	IsACL                                   Property = "isacl"
	IsACLProtected                          Property = "isaclprotected"
	IsDeleted                               Property = "isdeleted"
	Enforced                                Property = "enforced"
	Department                              Property = "department"
	HasCrossCertificatePair                 Property = "hascrosscertificatepair"
	HasSPN                                  Property = "hasspn"
	UnconstrainedDelegation                 Property = "unconstraineddelegation"
	LastLogon                               Property = "lastlogon"
	LastLogonTimestamp                      Property = "lastlogontimestamp"
	IsPrimaryGroup                          Property = "isprimarygroup"
	HasLAPS                                 Property = "haslaps"
	DontRequirePreAuth                      Property = "dontreqpreauth"
	LogonType                               Property = "logontype"
	HasURA                                  Property = "hasura"
	PasswordNeverExpires                    Property = "pwdneverexpires"
	PasswordNotRequired                     Property = "passwordnotreqd"
	FunctionalLevel                         Property = "functionallevel"
	TrustType                               Property = "trusttype"
	SidFiltering                            Property = "sidfiltering"
	TrustedToAuth                           Property = "trustedtoauth"
	SamAccountName                          Property = "samaccountname"
	CertificateMappingMethodsRaw            Property = "certificatemappingmethodsraw"
	CertificateMappingMethods               Property = "certificatemappingmethods"
	StrongCertificateBindingEnforcementRaw  Property = "strongcertificatebindingenforcementraw"
	StrongCertificateBindingEnforcement     Property = "strongcertificatebindingenforcement"
	EKUs                                    Property = "ekus"
	SubjectAltRequireUPN                    Property = "subjectaltrequireupn"
	SubjectAltRequireDNS                    Property = "subjectaltrequiredns"
	SubjectAltRequireDomainDNS              Property = "subjectaltrequiredomaindns"
	SubjectAltRequireEmail                  Property = "subjectaltrequireemail"
	SubjectAltRequireSPN                    Property = "subjectaltrequirespn"
	SubjectRequireEmail                     Property = "subjectrequireemail"
	AuthorizedSignatures                    Property = "authorizedsignatures"
	ApplicationPolicies                     Property = "applicationpolicies"
	IssuancePolicies                        Property = "issuancepolicies"
	SchemaVersion                           Property = "schemaversion"
	RequiresManagerApproval                 Property = "requiresmanagerapproval"
	AuthenticationEnabled                   Property = "authenticationenabled"
	SchannelAuthenticationEnabled           Property = "schannelauthenticationenabled"
	EnrolleeSuppliesSubject                 Property = "enrolleesuppliessubject"
	CertificateApplicationPolicy            Property = "certificateapplicationpolicy"
	CertificateNameFlag                     Property = "certificatenameflag"
	EffectiveEKUs                           Property = "effectiveekus"
	EnrollmentFlag                          Property = "enrollmentflag"
	Flags                                   Property = "flags"
	NoSecurityExtension                     Property = "nosecurityextension"
	RenewalPeriod                           Property = "renewalperiod"
	ValidityPeriod                          Property = "validityperiod"
	OID                                     Property = "oid"
	HomeDirectory                           Property = "homedirectory"
	CertificatePolicy                       Property = "certificatepolicy"
	CertTemplateOID                         Property = "certtemplateoid"
	GroupLinkID                             Property = "grouplinkid"
	ObjectGUID                              Property = "objectguid"
	ExpirePasswordsOnSmartCardOnlyAccounts  Property = "expirepasswordsonsmartcardonlyaccounts"
	MachineAccountQuota                     Property = "machineaccountquota"
	SupportedKerberosEncryptionTypes        Property = "supportedencryptiontypes"
	TGTDelegationEnabled                    Property = "tgtdelegationenabled"
	PasswordStoredUsingReversibleEncryption Property = "encryptedtextpwdallowed"
	SmartcardRequired                       Property = "smartcardrequired"
	UseDESKeyOnly                           Property = "usedeskeyonly"
	LogonScriptEnabled                      Property = "logonscriptenabled"
	LockedOut                               Property = "lockedout"
	UserCannotChangePassword                Property = "passwordcantchange"
	PasswordExpired                         Property = "passwordexpired"
	DSHeuristics                            Property = "dsheuristics"
	UserAccountControl                      Property = "useraccountcontrol"
	TrustAttributes                         Property = "trustattributes"
	MinPwdLength                            Property = "minpwdlength"
	PwdProperties                           Property = "pwdproperties"
	PwdHistoryLength                        Property = "pwdhistorylength"
	LockoutThreshold                        Property = "lockoutthreshold"
	MinPwdAge                               Property = "minpwdage"
	MaxPwdAge                               Property = "maxpwdage"
	LockoutDuration                         Property = "lockoutduration"
	LockoutObservationWindow                Property = "lockoutobservationwindow"
	SmbSigning                              Property = "smbsigning"
)

func AllProperties() []Property {
	return []Property{AdminCount, CASecurityCollected, CAName, CertChain, CertName, CertThumbprint, CertThumbprints, HasEnrollmentAgentRestrictions, EnrollmentAgentRestrictionsCollected, IsUserSpecifiesSanEnabled, IsUserSpecifiesSanEnabledCollected, RoleSeparationEnabled, RoleSeparationEnabledCollected, HasBasicConstraints, BasicConstraintPathLength, UnresolvedPublishedTemplates, DNSHostname, CrossCertificatePair, DistinguishedName, DomainFQDN, DomainSID, Sensitive, HighValue, BlocksInheritance, IsACL, IsACLProtected, IsDeleted, Enforced, Department, HasCrossCertificatePair, HasSPN, UnconstrainedDelegation, LastLogon, LastLogonTimestamp, IsPrimaryGroup, HasLAPS, DontRequirePreAuth, LogonType, HasURA, PasswordNeverExpires, PasswordNotRequired, FunctionalLevel, TrustType, SidFiltering, TrustedToAuth, SamAccountName, CertificateMappingMethodsRaw, CertificateMappingMethods, StrongCertificateBindingEnforcementRaw, StrongCertificateBindingEnforcement, EKUs, SubjectAltRequireUPN, SubjectAltRequireDNS, SubjectAltRequireDomainDNS, SubjectAltRequireEmail, SubjectAltRequireSPN, SubjectRequireEmail, AuthorizedSignatures, ApplicationPolicies, IssuancePolicies, SchemaVersion, RequiresManagerApproval, AuthenticationEnabled, SchannelAuthenticationEnabled, EnrolleeSuppliesSubject, CertificateApplicationPolicy, CertificateNameFlag, EffectiveEKUs, EnrollmentFlag, Flags, NoSecurityExtension, RenewalPeriod, ValidityPeriod, OID, HomeDirectory, CertificatePolicy, CertTemplateOID, GroupLinkID, ObjectGUID, ExpirePasswordsOnSmartCardOnlyAccounts, MachineAccountQuota, SupportedKerberosEncryptionTypes, TGTDelegationEnabled, PasswordStoredUsingReversibleEncryption, SmartcardRequired, UseDESKeyOnly, LogonScriptEnabled, LockedOut, UserCannotChangePassword, PasswordExpired, DSHeuristics, UserAccountControl, TrustAttributes, MinPwdLength, PwdProperties, PwdHistoryLength, LockoutThreshold, MinPwdAge, MaxPwdAge, LockoutDuration, LockoutObservationWindow, SmbSigning}
}
func ParseProperty(source string) (Property, error) {
	switch source {
	case "admincount":
		return AdminCount, nil
	case "casecuritycollected":
		return CASecurityCollected, nil
	case "caname":
		return CAName, nil
	case "certchain":
		return CertChain, nil
	case "certname":
		return CertName, nil
	case "certthumbprint":
		return CertThumbprint, nil
	case "certthumbprints":
		return CertThumbprints, nil
	case "hasenrollmentagentrestrictions":
		return HasEnrollmentAgentRestrictions, nil
	case "enrollmentagentrestrictionscollected":
		return EnrollmentAgentRestrictionsCollected, nil
	case "isuserspecifiessanenabled":
		return IsUserSpecifiesSanEnabled, nil
	case "isuserspecifiessanenabledcollected":
		return IsUserSpecifiesSanEnabledCollected, nil
	case "roleseparationenabled":
		return RoleSeparationEnabled, nil
	case "roleseparationenabledcollected":
		return RoleSeparationEnabledCollected, nil
	case "hasbasicconstraints":
		return HasBasicConstraints, nil
	case "basicconstraintpathlength":
		return BasicConstraintPathLength, nil
	case "unresolvedpublishedtemplates":
		return UnresolvedPublishedTemplates, nil
	case "dnshostname":
		return DNSHostname, nil
	case "crosscertificatepair":
		return CrossCertificatePair, nil
	case "distinguishedname":
		return DistinguishedName, nil
	case "domain":
		return DomainFQDN, nil
	case "domainsid":
		return DomainSID, nil
	case "sensitive":
		return Sensitive, nil
	case "highvalue":
		return HighValue, nil
	case "blocksinheritance":
		return BlocksInheritance, nil
	case "isacl":
		return IsACL, nil
	case "isaclprotected":
		return IsACLProtected, nil
	case "isdeleted":
		return IsDeleted, nil
	case "enforced":
		return Enforced, nil
	case "department":
		return Department, nil
	case "hascrosscertificatepair":
		return HasCrossCertificatePair, nil
	case "hasspn":
		return HasSPN, nil
	case "unconstraineddelegation":
		return UnconstrainedDelegation, nil
	case "lastlogon":
		return LastLogon, nil
	case "lastlogontimestamp":
		return LastLogonTimestamp, nil
	case "isprimarygroup":
		return IsPrimaryGroup, nil
	case "haslaps":
		return HasLAPS, nil
	case "dontreqpreauth":
		return DontRequirePreAuth, nil
	case "logontype":
		return LogonType, nil
	case "hasura":
		return HasURA, nil
	case "pwdneverexpires":
		return PasswordNeverExpires, nil
	case "passwordnotreqd":
		return PasswordNotRequired, nil
	case "functionallevel":
		return FunctionalLevel, nil
	case "trusttype":
		return TrustType, nil
	case "sidfiltering":
		return SidFiltering, nil
	case "trustedtoauth":
		return TrustedToAuth, nil
	case "samaccountname":
		return SamAccountName, nil
	case "certificatemappingmethodsraw":
		return CertificateMappingMethodsRaw, nil
	case "certificatemappingmethods":
		return CertificateMappingMethods, nil
	case "strongcertificatebindingenforcementraw":
		return StrongCertificateBindingEnforcementRaw, nil
	case "strongcertificatebindingenforcement":
		return StrongCertificateBindingEnforcement, nil
	case "ekus":
		return EKUs, nil
	case "subjectaltrequireupn":
		return SubjectAltRequireUPN, nil
	case "subjectaltrequiredns":
		return SubjectAltRequireDNS, nil
	case "subjectaltrequiredomaindns":
		return SubjectAltRequireDomainDNS, nil
	case "subjectaltrequireemail":
		return SubjectAltRequireEmail, nil
	case "subjectaltrequirespn":
		return SubjectAltRequireSPN, nil
	case "subjectrequireemail":
		return SubjectRequireEmail, nil
	case "authorizedsignatures":
		return AuthorizedSignatures, nil
	case "applicationpolicies":
		return ApplicationPolicies, nil
	case "issuancepolicies":
		return IssuancePolicies, nil
	case "schemaversion":
		return SchemaVersion, nil
	case "requiresmanagerapproval":
		return RequiresManagerApproval, nil
	case "authenticationenabled":
		return AuthenticationEnabled, nil
	case "schannelauthenticationenabled":
		return SchannelAuthenticationEnabled, nil
	case "enrolleesuppliessubject":
		return EnrolleeSuppliesSubject, nil
	case "certificateapplicationpolicy":
		return CertificateApplicationPolicy, nil
	case "certificatenameflag":
		return CertificateNameFlag, nil
	case "effectiveekus":
		return EffectiveEKUs, nil
	case "enrollmentflag":
		return EnrollmentFlag, nil
	case "flags":
		return Flags, nil
	case "nosecurityextension":
		return NoSecurityExtension, nil
	case "renewalperiod":
		return RenewalPeriod, nil
	case "validityperiod":
		return ValidityPeriod, nil
	case "oid":
		return OID, nil
	case "homedirectory":
		return HomeDirectory, nil
	case "certificatepolicy":
		return CertificatePolicy, nil
	case "certtemplateoid":
		return CertTemplateOID, nil
	case "grouplinkid":
		return GroupLinkID, nil
	case "objectguid":
		return ObjectGUID, nil
	case "expirepasswordsonsmartcardonlyaccounts":
		return ExpirePasswordsOnSmartCardOnlyAccounts, nil
	case "machineaccountquota":
		return MachineAccountQuota, nil
	case "supportedencryptiontypes":
		return SupportedKerberosEncryptionTypes, nil
	case "tgtdelegationenabled":
		return TGTDelegationEnabled, nil
	case "encryptedtextpwdallowed":
		return PasswordStoredUsingReversibleEncryption, nil
	case "smartcardrequired":
		return SmartcardRequired, nil
	case "usedeskeyonly":
		return UseDESKeyOnly, nil
	case "logonscriptenabled":
		return LogonScriptEnabled, nil
	case "lockedout":
		return LockedOut, nil
	case "passwordcantchange":
		return UserCannotChangePassword, nil
	case "passwordexpired":
		return PasswordExpired, nil
	case "dsheuristics":
		return DSHeuristics, nil
	case "useraccountcontrol":
		return UserAccountControl, nil
	case "trustattributes":
		return TrustAttributes, nil
	case "minpwdlength":
		return MinPwdLength, nil
	case "pwdproperties":
		return PwdProperties, nil
	case "pwdhistorylength":
		return PwdHistoryLength, nil
	case "lockoutthreshold":
		return LockoutThreshold, nil
	case "minpwdage":
		return MinPwdAge, nil
	case "maxpwdage":
		return MaxPwdAge, nil
	case "lockoutduration":
		return LockoutDuration, nil
	case "lockoutobservationwindow":
		return LockoutObservationWindow, nil
	case "smbsigning":
		return SmbSigning, nil
	default:
		return "", errors.New("Invalid enumeration value: " + source)
	}
}
func (s Property) String() string {
	switch s {
	case AdminCount:
		return string(AdminCount)
	case CASecurityCollected:
		return string(CASecurityCollected)
	case CAName:
		return string(CAName)
	case CertChain:
		return string(CertChain)
	case CertName:
		return string(CertName)
	case CertThumbprint:
		return string(CertThumbprint)
	case CertThumbprints:
		return string(CertThumbprints)
	case HasEnrollmentAgentRestrictions:
		return string(HasEnrollmentAgentRestrictions)
	case EnrollmentAgentRestrictionsCollected:
		return string(EnrollmentAgentRestrictionsCollected)
	case IsUserSpecifiesSanEnabled:
		return string(IsUserSpecifiesSanEnabled)
	case IsUserSpecifiesSanEnabledCollected:
		return string(IsUserSpecifiesSanEnabledCollected)
	case RoleSeparationEnabled:
		return string(RoleSeparationEnabled)
	case RoleSeparationEnabledCollected:
		return string(RoleSeparationEnabledCollected)
	case HasBasicConstraints:
		return string(HasBasicConstraints)
	case BasicConstraintPathLength:
		return string(BasicConstraintPathLength)
	case UnresolvedPublishedTemplates:
		return string(UnresolvedPublishedTemplates)
	case DNSHostname:
		return string(DNSHostname)
	case CrossCertificatePair:
		return string(CrossCertificatePair)
	case DistinguishedName:
		return string(DistinguishedName)
	case DomainFQDN:
		return string(DomainFQDN)
	case DomainSID:
		return string(DomainSID)
	case Sensitive:
		return string(Sensitive)
	case HighValue:
		return string(HighValue)
	case BlocksInheritance:
		return string(BlocksInheritance)
	case IsACL:
		return string(IsACL)
	case IsACLProtected:
		return string(IsACLProtected)
	case IsDeleted:
		return string(IsDeleted)
	case Enforced:
		return string(Enforced)
	case Department:
		return string(Department)
	case HasCrossCertificatePair:
		return string(HasCrossCertificatePair)
	case HasSPN:
		return string(HasSPN)
	case UnconstrainedDelegation:
		return string(UnconstrainedDelegation)
	case LastLogon:
		return string(LastLogon)
	case LastLogonTimestamp:
		return string(LastLogonTimestamp)
	case IsPrimaryGroup:
		return string(IsPrimaryGroup)
	case HasLAPS:
		return string(HasLAPS)
	case DontRequirePreAuth:
		return string(DontRequirePreAuth)
	case LogonType:
		return string(LogonType)
	case HasURA:
		return string(HasURA)
	case PasswordNeverExpires:
		return string(PasswordNeverExpires)
	case PasswordNotRequired:
		return string(PasswordNotRequired)
	case FunctionalLevel:
		return string(FunctionalLevel)
	case TrustType:
		return string(TrustType)
	case SidFiltering:
		return string(SidFiltering)
	case TrustedToAuth:
		return string(TrustedToAuth)
	case SamAccountName:
		return string(SamAccountName)
	case CertificateMappingMethodsRaw:
		return string(CertificateMappingMethodsRaw)
	case CertificateMappingMethods:
		return string(CertificateMappingMethods)
	case StrongCertificateBindingEnforcementRaw:
		return string(StrongCertificateBindingEnforcementRaw)
	case StrongCertificateBindingEnforcement:
		return string(StrongCertificateBindingEnforcement)
	case EKUs:
		return string(EKUs)
	case SubjectAltRequireUPN:
		return string(SubjectAltRequireUPN)
	case SubjectAltRequireDNS:
		return string(SubjectAltRequireDNS)
	case SubjectAltRequireDomainDNS:
		return string(SubjectAltRequireDomainDNS)
	case SubjectAltRequireEmail:
		return string(SubjectAltRequireEmail)
	case SubjectAltRequireSPN:
		return string(SubjectAltRequireSPN)
	case SubjectRequireEmail:
		return string(SubjectRequireEmail)
	case AuthorizedSignatures:
		return string(AuthorizedSignatures)
	case ApplicationPolicies:
		return string(ApplicationPolicies)
	case IssuancePolicies:
		return string(IssuancePolicies)
	case SchemaVersion:
		return string(SchemaVersion)
	case RequiresManagerApproval:
		return string(RequiresManagerApproval)
	case AuthenticationEnabled:
		return string(AuthenticationEnabled)
	case SchannelAuthenticationEnabled:
		return string(SchannelAuthenticationEnabled)
	case EnrolleeSuppliesSubject:
		return string(EnrolleeSuppliesSubject)
	case CertificateApplicationPolicy:
		return string(CertificateApplicationPolicy)
	case CertificateNameFlag:
		return string(CertificateNameFlag)
	case EffectiveEKUs:
		return string(EffectiveEKUs)
	case EnrollmentFlag:
		return string(EnrollmentFlag)
	case Flags:
		return string(Flags)
	case NoSecurityExtension:
		return string(NoSecurityExtension)
	case RenewalPeriod:
		return string(RenewalPeriod)
	case ValidityPeriod:
		return string(ValidityPeriod)
	case OID:
		return string(OID)
	case HomeDirectory:
		return string(HomeDirectory)
	case CertificatePolicy:
		return string(CertificatePolicy)
	case CertTemplateOID:
		return string(CertTemplateOID)
	case GroupLinkID:
		return string(GroupLinkID)
	case ObjectGUID:
		return string(ObjectGUID)
	case ExpirePasswordsOnSmartCardOnlyAccounts:
		return string(ExpirePasswordsOnSmartCardOnlyAccounts)
	case MachineAccountQuota:
		return string(MachineAccountQuota)
	case SupportedKerberosEncryptionTypes:
		return string(SupportedKerberosEncryptionTypes)
	case TGTDelegationEnabled:
		return string(TGTDelegationEnabled)
	case PasswordStoredUsingReversibleEncryption:
		return string(PasswordStoredUsingReversibleEncryption)
	case SmartcardRequired:
		return string(SmartcardRequired)
	case UseDESKeyOnly:
		return string(UseDESKeyOnly)
	case LogonScriptEnabled:
		return string(LogonScriptEnabled)
	case LockedOut:
		return string(LockedOut)
	case UserCannotChangePassword:
		return string(UserCannotChangePassword)
	case PasswordExpired:
		return string(PasswordExpired)
	case DSHeuristics:
		return string(DSHeuristics)
	case UserAccountControl:
		return string(UserAccountControl)
	case TrustAttributes:
		return string(TrustAttributes)
	case MinPwdLength:
		return string(MinPwdLength)
	case PwdProperties:
		return string(PwdProperties)
	case PwdHistoryLength:
		return string(PwdHistoryLength)
	case LockoutThreshold:
		return string(LockoutThreshold)
	case MinPwdAge:
		return string(MinPwdAge)
	case MaxPwdAge:
		return string(MaxPwdAge)
	case LockoutDuration:
		return string(LockoutDuration)
	case LockoutObservationWindow:
		return string(LockoutObservationWindow)
	case SmbSigning:
		return string(SmbSigning)
	default:
		return "Invalid enumeration case: " + string(s)
	}
}
func (s Property) Name() string {
	switch s {
	case AdminCount:
		return "Admin Count"
	case CASecurityCollected:
		return "CA Security Collected"
	case CAName:
		return "CA Name"
	case CertChain:
		return "Certificate Chain"
	case CertName:
		return "Certificate Name"
	case CertThumbprint:
		return "Certificate Thumbprint"
	case CertThumbprints:
		return "Certificate Thumbprints"
	case HasEnrollmentAgentRestrictions:
		return "Has Enrollment Agent Restrictions"
	case EnrollmentAgentRestrictionsCollected:
		return "Enrollment Agent Restrictions Collected"
	case IsUserSpecifiesSanEnabled:
		return "Is User Specifies San Enabled"
	case IsUserSpecifiesSanEnabledCollected:
		return "Is User Specifies San Enabled Collected"
	case RoleSeparationEnabled:
		return "Role Separation Enabled"
	case RoleSeparationEnabledCollected:
		return "Role Separation Enabled Collected"
	case HasBasicConstraints:
		return "Has Basic Constraints"
	case BasicConstraintPathLength:
		return "Basic Constraint Path Length"
	case UnresolvedPublishedTemplates:
		return "Unresolved Published Certificate Templates"
	case DNSHostname:
		return "DNS Hostname"
	case CrossCertificatePair:
		return "Cross Certificate Pair"
	case DistinguishedName:
		return "Distinguished Name"
	case DomainFQDN:
		return "Domain FQDN"
	case DomainSID:
		return "Domain SID"
	case Sensitive:
		return "Marked Sensitive"
	case HighValue:
		return "High Value"
	case BlocksInheritance:
		return "Blocks GPO Inheritance"
	case IsACL:
		return "Is ACL"
	case IsACLProtected:
		return "ACL Inheritance Denied"
	case IsDeleted:
		return "Is Deleted"
	case Enforced:
		return "Enforced"
	case Department:
		return "Department"
	case HasCrossCertificatePair:
		return "Has Cross Certificate Pair"
	case HasSPN:
		return "Has SPN"
	case UnconstrainedDelegation:
		return "Allows Unconstrained Delegation"
	case LastLogon:
		return "Last Logon"
	case LastLogonTimestamp:
		return "Last Logon (Replicated)"
	case IsPrimaryGroup:
		return "Is Primary Group"
	case HasLAPS:
		return "LAPS Enabled"
	case DontRequirePreAuth:
		return "Do Not Require Pre-Authentication"
	case LogonType:
		return "Logon Type"
	case HasURA:
		return "Has User Rights Assignment Collection"
	case PasswordNeverExpires:
		return "Password Never Expires"
	case PasswordNotRequired:
		return "Password Not Required"
	case FunctionalLevel:
		return "Functional Level"
	case TrustType:
		return "Trust Type"
	case SidFiltering:
		return "SID Filtering Enabled"
	case TrustedToAuth:
		return "Trusted For Constrained Delegation"
	case SamAccountName:
		return "SAM Account Name"
	case CertificateMappingMethodsRaw:
		return "Certificate Mapping Methods (Raw)"
	case CertificateMappingMethods:
		return "Certificate Mapping Methods"
	case StrongCertificateBindingEnforcementRaw:
		return "Strong Certificate Binding Enforcement (Raw)"
	case StrongCertificateBindingEnforcement:
		return "Strong Certificate Binding Enforcement"
	case EKUs:
		return "Enhanced Key Usage"
	case SubjectAltRequireUPN:
		return "Subject Alternative Name Require UPN"
	case SubjectAltRequireDNS:
		return "Subject Alternative Name Require DNS"
	case SubjectAltRequireDomainDNS:
		return "Subject Alternative Name Require Domain DNS"
	case SubjectAltRequireEmail:
		return "Subject Alternative Name Require Email"
	case SubjectAltRequireSPN:
		return "Subject Alternative Name Require SPN"
	case SubjectRequireEmail:
		return "Subject Require Email"
	case AuthorizedSignatures:
		return "Authorized Signatures Required"
	case ApplicationPolicies:
		return "Application Policies Required"
	case IssuancePolicies:
		return "Issuance Policies Required"
	case SchemaVersion:
		return "Schema Version"
	case RequiresManagerApproval:
		return "Requires Manager Approval"
	case AuthenticationEnabled:
		return "Authentication Enabled"
	case SchannelAuthenticationEnabled:
		return "Schannel Authentication Enabled"
	case EnrolleeSuppliesSubject:
		return "Enrollee Supplies Subject"
	case CertificateApplicationPolicy:
		return "Application Policy Extensions"
	case CertificateNameFlag:
		return "Certificate Name Flags"
	case EffectiveEKUs:
		return "Effective EKUs"
	case EnrollmentFlag:
		return "Enrollment Flags"
	case Flags:
		return "Flags"
	case NoSecurityExtension:
		return "No Security Extension"
	case RenewalPeriod:
		return "Renewal Period"
	case ValidityPeriod:
		return "Validity Period"
	case OID:
		return "OID"
	case HomeDirectory:
		return "Home Directory"
	case CertificatePolicy:
		return "Issuance Policy Extensions"
	case CertTemplateOID:
		return "Certificate Template OID"
	case GroupLinkID:
		return "Group Link ID"
	case ObjectGUID:
		return "Object GUID"
	case ExpirePasswordsOnSmartCardOnlyAccounts:
		return "Expire Passwords on Smart Card only Accounts"
	case MachineAccountQuota:
		return "Machine Account Quota"
	case SupportedKerberosEncryptionTypes:
		return "Supported Kerberos Encryption Types"
	case TGTDelegationEnabled:
		return "TGT Delegation Enabled"
	case PasswordStoredUsingReversibleEncryption:
		return "Password Stored Using Reversible Encryption"
	case SmartcardRequired:
		return "Smartcard Required"
	case UseDESKeyOnly:
		return "Use DES Key Only"
	case LogonScriptEnabled:
		return "Logon Script Enabled"
	case LockedOut:
		return "Locked Out"
	case UserCannotChangePassword:
		return "User Cannot Change Password"
	case PasswordExpired:
		return "Password Expired"
	case DSHeuristics:
		return "DSHeuristics"
	case UserAccountControl:
		return "User Account Control"
	case TrustAttributes:
		return "Trust Attributes"
	case MinPwdLength:
		return "Minimum password length"
	case PwdProperties:
		return "Password Properties"
	case PwdHistoryLength:
		return "Password History Length"
	case LockoutThreshold:
		return "Lockout Threshold"
	case MinPwdAge:
		return "Minimum Password Age"
	case MaxPwdAge:
		return "Maximum Password Age"
	case LockoutDuration:
		return "Lockout Duration"
	case LockoutObservationWindow:
		return "Lockout Observation Window"
	case SmbSigning:
		return "SMB Signing"
	default:
		return "Invalid enumeration case: " + string(s)
	}
}
func (s Property) Is(others ...graph.Kind) bool {
	for _, other := range others {
		if value, err := ParseProperty(other.String()); err == nil && value == s {
			return true
		}
	}
	return false
}
func Nodes() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser, AIACA, RootCA, EnterpriseCA, NTAuthStore, CertTemplate, IssuancePolicy}
}
func Relationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, CoerceToTGT, GetChanges, GetChangesAll, GetChangesInFilteredSet, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, LocalToComputer, MemberOfLocalGroup, RemoteInteractiveLogonRight, SyncLAPSPassword, WriteAccountRestrictions, WriteGPLink, RootCAFor, DCFor, PublishedTo, ManageCertificates, ManageCA, DelegatedEnrollmentAgent, Enroll, HostsCAService, WritePKIEnrollmentFlag, WritePKINameFlag, NTAuthStoreFor, TrustedForNTAuth, EnterpriseCAFor, IssuedSignedBy, GoldenCert, EnrollOnBehalfOf, OIDGroupLink, ExtendedByPolicy, ADCSESC1, ADCSESC3, ADCSESC4, ADCSESC6a, ADCSESC6b, ADCSESC9a, ADCSESC9b, ADCSESC10a, ADCSESC10b, ADCSESC13, SyncedToEntraUser, CoerceAndRelayNTLMToSMB}
}
func ACLRelationships() []graph.Kind {
	return []graph.Kind{AllExtendedRights, ForceChangePassword, AddMember, AddAllowedToAct, GenericAll, WriteDACL, WriteOwner, GenericWrite, ReadLAPSPassword, ReadGMSAPassword, Owns, AddSelf, WriteSPN, AddKeyCredentialLink, GetChanges, GetChangesAll, GetChangesInFilteredSet, WriteAccountRestrictions, WriteGPLink, SyncLAPSPassword, DCSync, ManageCertificates, ManageCA, Enroll, WritePKIEnrollmentFlag, WritePKINameFlag}
}
func PathfindingRelationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, CoerceToTGT, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, SyncLAPSPassword, WriteAccountRestrictions, WriteGPLink, GoldenCert, ADCSESC1, ADCSESC3, ADCSESC4, ADCSESC6a, ADCSESC6b, ADCSESC9a, ADCSESC9b, ADCSESC10a, ADCSESC10b, ADCSESC13, DCFor, SyncedToEntraUser, CoerceAndRelayNTLMToSMB}
}
func IsACLKind(s graph.Kind) bool {
	for _, acl := range ACLRelationships() {
		if s == acl {
			return true
		}
	}
	return false
}
func NodeKinds() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser, AIACA, RootCA, EnterpriseCA, NTAuthStore, CertTemplate, IssuancePolicy}
}
