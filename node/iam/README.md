# CMesh IAM API

## Public Facing Handshake Related Functions

### Core Endpoints

#### DIDSession

Requests the current session token. If none exists, it attempts to authenticate with DIDAuth.

#### DIDAuth

Primary endpoint for authenticating. In current form, if no SSI exists, one is "generated", and if no DID exists, likewise one is "generated". Once a valid DID exists, the server calls `dIDSessionCall()` to begin the handshake.

#### DIDSessionHangup

As the name suggests, terminates any existing session associated with that DID.

#### DIDGen

Generates a new DID session for the given SSI and returns it. Note: the internal state hardcodes everything to index 0... so it won't do much good calling it outside of the first call through `DIDSession`/`DIDAuth`.

### Handshake Specific Endpoints

#### DIDSessionAnswer

Once the server has returned the callstring from `dIDSessionCall()`, the user calls `DIDSessionAnswer()` with the callstring and a signature to confirm they initiated the `DIDAuth()` call.

#### DIDSessionConsent

Once the server has returned the confirmation from `dIDSessionConfirm()`, the user calls `DIDSessionConsent()` with the confirmation string and a signature to consent to the confirmed answer, affirming they initiated the `DIDAuth()` call and establishing a session.

## Private Handshake Related Functions

### Call Functions

#### dIDSessionCall

Generates and returns the signed callstring, initiating the session handshake from the server's perspective.

#### genSignCallString

Generates, signs and returns the callstring.

#### genCallString

Generates the callstring.

#### signCall

Signs the call string.

### Answer Functions

#### genAnswerString

Generates, signs and returns the answer string.

#### signAnswer

Signs the answer string.

#### expectedAnswerSig

Calculates the expected answer signature for the handshake. Used to compare to the answer signature provided.

### Confirm Functions

#### dIDSessionConfirm

Generates and returns the signed confirmation string, acknowledging the user's answer and asking them for a second verification to be sure.

#### genConfirmString

Generates, signs and returns the confirmation string.

#### signConfirm

Signs the confirmation string.

### Consent Functions

#### signHandShake

Signs the consent string.

#### signConsent

Signs the... confirm string? Brain was mush, will tidy later.

## Public Facing Attribute Related Functions

Documentation to follow once actual cryptography in place to enable IRMA's attribute oriented featureset.
