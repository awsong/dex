// Code generated by ent, DO NOT EDIT.

package refreshtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/awsong/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldID, id))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClientID, v))
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldNonce, v))
}

// ClaimsUserID applies equality check predicate on the "claims_user_id" field. It's identical to ClaimsUserIDEQ.
func ClaimsUserID(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsUserID, v))
}

// ClaimsUsername applies equality check predicate on the "claims_username" field. It's identical to ClaimsUsernameEQ.
func ClaimsUsername(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsUsername, v))
}

// ClaimsEmail applies equality check predicate on the "claims_email" field. It's identical to ClaimsEmailEQ.
func ClaimsEmail(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsEmail, v))
}

// ClaimsEmailVerified applies equality check predicate on the "claims_email_verified" field. It's identical to ClaimsEmailVerifiedEQ.
func ClaimsEmailVerified(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsEmailVerified, v))
}

// ClaimsPreferredUsername applies equality check predicate on the "claims_preferred_username" field. It's identical to ClaimsPreferredUsernameEQ.
func ClaimsPreferredUsername(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsPreferredUsername, v))
}

// ConnectorID applies equality check predicate on the "connector_id" field. It's identical to ConnectorIDEQ.
func ConnectorID(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldConnectorID, v))
}

// ConnectorData applies equality check predicate on the "connector_data" field. It's identical to ConnectorDataEQ.
func ConnectorData(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldConnectorData, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldToken, v))
}

// ObsoleteToken applies equality check predicate on the "obsolete_token" field. It's identical to ObsoleteTokenEQ.
func ObsoleteToken(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldObsoleteToken, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldCreatedAt, v))
}

// LastUsed applies equality check predicate on the "last_used" field. It's identical to LastUsedEQ.
func LastUsed(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldLastUsed, v))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldClientID, v))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldClientID, v))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldClientID, v))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldClientID, v))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldClientID, v))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldClientID, v))
}

// ScopesIsNil applies the IsNil predicate on the "scopes" field.
func ScopesIsNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIsNull(FieldScopes))
}

// ScopesNotNil applies the NotNil predicate on the "scopes" field.
func ScopesNotNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotNull(FieldScopes))
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldNonce, v))
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldNonce, v))
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldNonce, vs...))
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldNonce, vs...))
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldNonce, v))
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldNonce, v))
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldNonce, v))
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldNonce, v))
}

// NonceContains applies the Contains predicate on the "nonce" field.
func NonceContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldNonce, v))
}

// NonceHasPrefix applies the HasPrefix predicate on the "nonce" field.
func NonceHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldNonce, v))
}

// NonceHasSuffix applies the HasSuffix predicate on the "nonce" field.
func NonceHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldNonce, v))
}

// NonceEqualFold applies the EqualFold predicate on the "nonce" field.
func NonceEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldNonce, v))
}

// NonceContainsFold applies the ContainsFold predicate on the "nonce" field.
func NonceContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldNonce, v))
}

// ClaimsUserIDEQ applies the EQ predicate on the "claims_user_id" field.
func ClaimsUserIDEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsUserID, v))
}

// ClaimsUserIDNEQ applies the NEQ predicate on the "claims_user_id" field.
func ClaimsUserIDNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClaimsUserID, v))
}

// ClaimsUserIDIn applies the In predicate on the "claims_user_id" field.
func ClaimsUserIDIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldClaimsUserID, vs...))
}

// ClaimsUserIDNotIn applies the NotIn predicate on the "claims_user_id" field.
func ClaimsUserIDNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldClaimsUserID, vs...))
}

// ClaimsUserIDGT applies the GT predicate on the "claims_user_id" field.
func ClaimsUserIDGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldClaimsUserID, v))
}

// ClaimsUserIDGTE applies the GTE predicate on the "claims_user_id" field.
func ClaimsUserIDGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldClaimsUserID, v))
}

// ClaimsUserIDLT applies the LT predicate on the "claims_user_id" field.
func ClaimsUserIDLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldClaimsUserID, v))
}

// ClaimsUserIDLTE applies the LTE predicate on the "claims_user_id" field.
func ClaimsUserIDLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldClaimsUserID, v))
}

// ClaimsUserIDContains applies the Contains predicate on the "claims_user_id" field.
func ClaimsUserIDContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldClaimsUserID, v))
}

// ClaimsUserIDHasPrefix applies the HasPrefix predicate on the "claims_user_id" field.
func ClaimsUserIDHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldClaimsUserID, v))
}

// ClaimsUserIDHasSuffix applies the HasSuffix predicate on the "claims_user_id" field.
func ClaimsUserIDHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldClaimsUserID, v))
}

// ClaimsUserIDEqualFold applies the EqualFold predicate on the "claims_user_id" field.
func ClaimsUserIDEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldClaimsUserID, v))
}

// ClaimsUserIDContainsFold applies the ContainsFold predicate on the "claims_user_id" field.
func ClaimsUserIDContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldClaimsUserID, v))
}

// ClaimsUsernameEQ applies the EQ predicate on the "claims_username" field.
func ClaimsUsernameEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsUsername, v))
}

// ClaimsUsernameNEQ applies the NEQ predicate on the "claims_username" field.
func ClaimsUsernameNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClaimsUsername, v))
}

// ClaimsUsernameIn applies the In predicate on the "claims_username" field.
func ClaimsUsernameIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldClaimsUsername, vs...))
}

// ClaimsUsernameNotIn applies the NotIn predicate on the "claims_username" field.
func ClaimsUsernameNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldClaimsUsername, vs...))
}

// ClaimsUsernameGT applies the GT predicate on the "claims_username" field.
func ClaimsUsernameGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldClaimsUsername, v))
}

// ClaimsUsernameGTE applies the GTE predicate on the "claims_username" field.
func ClaimsUsernameGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldClaimsUsername, v))
}

// ClaimsUsernameLT applies the LT predicate on the "claims_username" field.
func ClaimsUsernameLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldClaimsUsername, v))
}

// ClaimsUsernameLTE applies the LTE predicate on the "claims_username" field.
func ClaimsUsernameLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldClaimsUsername, v))
}

// ClaimsUsernameContains applies the Contains predicate on the "claims_username" field.
func ClaimsUsernameContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldClaimsUsername, v))
}

// ClaimsUsernameHasPrefix applies the HasPrefix predicate on the "claims_username" field.
func ClaimsUsernameHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldClaimsUsername, v))
}

// ClaimsUsernameHasSuffix applies the HasSuffix predicate on the "claims_username" field.
func ClaimsUsernameHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldClaimsUsername, v))
}

// ClaimsUsernameEqualFold applies the EqualFold predicate on the "claims_username" field.
func ClaimsUsernameEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldClaimsUsername, v))
}

// ClaimsUsernameContainsFold applies the ContainsFold predicate on the "claims_username" field.
func ClaimsUsernameContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldClaimsUsername, v))
}

// ClaimsEmailEQ applies the EQ predicate on the "claims_email" field.
func ClaimsEmailEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsEmail, v))
}

// ClaimsEmailNEQ applies the NEQ predicate on the "claims_email" field.
func ClaimsEmailNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClaimsEmail, v))
}

// ClaimsEmailIn applies the In predicate on the "claims_email" field.
func ClaimsEmailIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldClaimsEmail, vs...))
}

// ClaimsEmailNotIn applies the NotIn predicate on the "claims_email" field.
func ClaimsEmailNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldClaimsEmail, vs...))
}

// ClaimsEmailGT applies the GT predicate on the "claims_email" field.
func ClaimsEmailGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldClaimsEmail, v))
}

// ClaimsEmailGTE applies the GTE predicate on the "claims_email" field.
func ClaimsEmailGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldClaimsEmail, v))
}

// ClaimsEmailLT applies the LT predicate on the "claims_email" field.
func ClaimsEmailLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldClaimsEmail, v))
}

// ClaimsEmailLTE applies the LTE predicate on the "claims_email" field.
func ClaimsEmailLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldClaimsEmail, v))
}

// ClaimsEmailContains applies the Contains predicate on the "claims_email" field.
func ClaimsEmailContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldClaimsEmail, v))
}

// ClaimsEmailHasPrefix applies the HasPrefix predicate on the "claims_email" field.
func ClaimsEmailHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldClaimsEmail, v))
}

// ClaimsEmailHasSuffix applies the HasSuffix predicate on the "claims_email" field.
func ClaimsEmailHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldClaimsEmail, v))
}

// ClaimsEmailEqualFold applies the EqualFold predicate on the "claims_email" field.
func ClaimsEmailEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldClaimsEmail, v))
}

// ClaimsEmailContainsFold applies the ContainsFold predicate on the "claims_email" field.
func ClaimsEmailContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldClaimsEmail, v))
}

// ClaimsEmailVerifiedEQ applies the EQ predicate on the "claims_email_verified" field.
func ClaimsEmailVerifiedEQ(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsEmailVerified, v))
}

// ClaimsEmailVerifiedNEQ applies the NEQ predicate on the "claims_email_verified" field.
func ClaimsEmailVerifiedNEQ(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClaimsEmailVerified, v))
}

// ClaimsGroupsIsNil applies the IsNil predicate on the "claims_groups" field.
func ClaimsGroupsIsNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIsNull(FieldClaimsGroups))
}

// ClaimsGroupsNotNil applies the NotNil predicate on the "claims_groups" field.
func ClaimsGroupsNotNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotNull(FieldClaimsGroups))
}

// ClaimsPreferredUsernameEQ applies the EQ predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameNEQ applies the NEQ predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameIn applies the In predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldClaimsPreferredUsername, vs...))
}

// ClaimsPreferredUsernameNotIn applies the NotIn predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldClaimsPreferredUsername, vs...))
}

// ClaimsPreferredUsernameGT applies the GT predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameGTE applies the GTE predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameLT applies the LT predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameLTE applies the LTE predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameContains applies the Contains predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameHasPrefix applies the HasPrefix predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameHasSuffix applies the HasSuffix predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameEqualFold applies the EqualFold predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldClaimsPreferredUsername, v))
}

// ClaimsPreferredUsernameContainsFold applies the ContainsFold predicate on the "claims_preferred_username" field.
func ClaimsPreferredUsernameContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldClaimsPreferredUsername, v))
}

// ConnectorIDEQ applies the EQ predicate on the "connector_id" field.
func ConnectorIDEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldConnectorID, v))
}

// ConnectorIDNEQ applies the NEQ predicate on the "connector_id" field.
func ConnectorIDNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldConnectorID, v))
}

// ConnectorIDIn applies the In predicate on the "connector_id" field.
func ConnectorIDIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldConnectorID, vs...))
}

// ConnectorIDNotIn applies the NotIn predicate on the "connector_id" field.
func ConnectorIDNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldConnectorID, vs...))
}

// ConnectorIDGT applies the GT predicate on the "connector_id" field.
func ConnectorIDGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldConnectorID, v))
}

// ConnectorIDGTE applies the GTE predicate on the "connector_id" field.
func ConnectorIDGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldConnectorID, v))
}

// ConnectorIDLT applies the LT predicate on the "connector_id" field.
func ConnectorIDLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldConnectorID, v))
}

// ConnectorIDLTE applies the LTE predicate on the "connector_id" field.
func ConnectorIDLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldConnectorID, v))
}

// ConnectorIDContains applies the Contains predicate on the "connector_id" field.
func ConnectorIDContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldConnectorID, v))
}

// ConnectorIDHasPrefix applies the HasPrefix predicate on the "connector_id" field.
func ConnectorIDHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldConnectorID, v))
}

// ConnectorIDHasSuffix applies the HasSuffix predicate on the "connector_id" field.
func ConnectorIDHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldConnectorID, v))
}

// ConnectorIDEqualFold applies the EqualFold predicate on the "connector_id" field.
func ConnectorIDEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldConnectorID, v))
}

// ConnectorIDContainsFold applies the ContainsFold predicate on the "connector_id" field.
func ConnectorIDContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldConnectorID, v))
}

// ConnectorDataEQ applies the EQ predicate on the "connector_data" field.
func ConnectorDataEQ(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldConnectorData, v))
}

// ConnectorDataNEQ applies the NEQ predicate on the "connector_data" field.
func ConnectorDataNEQ(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldConnectorData, v))
}

// ConnectorDataIn applies the In predicate on the "connector_data" field.
func ConnectorDataIn(vs ...[]byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldConnectorData, vs...))
}

// ConnectorDataNotIn applies the NotIn predicate on the "connector_data" field.
func ConnectorDataNotIn(vs ...[]byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldConnectorData, vs...))
}

// ConnectorDataGT applies the GT predicate on the "connector_data" field.
func ConnectorDataGT(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldConnectorData, v))
}

// ConnectorDataGTE applies the GTE predicate on the "connector_data" field.
func ConnectorDataGTE(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldConnectorData, v))
}

// ConnectorDataLT applies the LT predicate on the "connector_data" field.
func ConnectorDataLT(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldConnectorData, v))
}

// ConnectorDataLTE applies the LTE predicate on the "connector_data" field.
func ConnectorDataLTE(v []byte) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldConnectorData, v))
}

// ConnectorDataIsNil applies the IsNil predicate on the "connector_data" field.
func ConnectorDataIsNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIsNull(FieldConnectorData))
}

// ConnectorDataNotNil applies the NotNil predicate on the "connector_data" field.
func ConnectorDataNotNil() predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotNull(FieldConnectorData))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldToken, v))
}

// ObsoleteTokenEQ applies the EQ predicate on the "obsolete_token" field.
func ObsoleteTokenEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldObsoleteToken, v))
}

// ObsoleteTokenNEQ applies the NEQ predicate on the "obsolete_token" field.
func ObsoleteTokenNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldObsoleteToken, v))
}

// ObsoleteTokenIn applies the In predicate on the "obsolete_token" field.
func ObsoleteTokenIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldObsoleteToken, vs...))
}

// ObsoleteTokenNotIn applies the NotIn predicate on the "obsolete_token" field.
func ObsoleteTokenNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldObsoleteToken, vs...))
}

// ObsoleteTokenGT applies the GT predicate on the "obsolete_token" field.
func ObsoleteTokenGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldObsoleteToken, v))
}

// ObsoleteTokenGTE applies the GTE predicate on the "obsolete_token" field.
func ObsoleteTokenGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldObsoleteToken, v))
}

// ObsoleteTokenLT applies the LT predicate on the "obsolete_token" field.
func ObsoleteTokenLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldObsoleteToken, v))
}

// ObsoleteTokenLTE applies the LTE predicate on the "obsolete_token" field.
func ObsoleteTokenLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldObsoleteToken, v))
}

// ObsoleteTokenContains applies the Contains predicate on the "obsolete_token" field.
func ObsoleteTokenContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldObsoleteToken, v))
}

// ObsoleteTokenHasPrefix applies the HasPrefix predicate on the "obsolete_token" field.
func ObsoleteTokenHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldObsoleteToken, v))
}

// ObsoleteTokenHasSuffix applies the HasSuffix predicate on the "obsolete_token" field.
func ObsoleteTokenHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldObsoleteToken, v))
}

// ObsoleteTokenEqualFold applies the EqualFold predicate on the "obsolete_token" field.
func ObsoleteTokenEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldObsoleteToken, v))
}

// ObsoleteTokenContainsFold applies the ContainsFold predicate on the "obsolete_token" field.
func ObsoleteTokenContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldObsoleteToken, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldCreatedAt, v))
}

// LastUsedEQ applies the EQ predicate on the "last_used" field.
func LastUsedEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldLastUsed, v))
}

// LastUsedNEQ applies the NEQ predicate on the "last_used" field.
func LastUsedNEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldLastUsed, v))
}

// LastUsedIn applies the In predicate on the "last_used" field.
func LastUsedIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldLastUsed, vs...))
}

// LastUsedNotIn applies the NotIn predicate on the "last_used" field.
func LastUsedNotIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldLastUsed, vs...))
}

// LastUsedGT applies the GT predicate on the "last_used" field.
func LastUsedGT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldLastUsed, v))
}

// LastUsedGTE applies the GTE predicate on the "last_used" field.
func LastUsedGTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldLastUsed, v))
}

// LastUsedLT applies the LT predicate on the "last_used" field.
func LastUsedLT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldLastUsed, v))
}

// LastUsedLTE applies the LTE predicate on the "last_used" field.
func LastUsedLTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldLastUsed, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
