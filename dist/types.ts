declare interface SetUSDCentsToDeSoExchangeRateRequest {
  USDCentsPerDeSo: number;
  AdminPublicKey: string;
}

declare interface SetUSDCentsToDeSoExchangeRateResponse {
  USDCentsPerDeSo: number;
}

declare interface GetUSDCentsToDeSoExchangeRateResponse {
  USDCentsPerDeSo: number;
}

declare interface SetBuyDeSoFeeBasisPointsRequest {
  BuyDeSoFeeBasisPoints: number;
  AdminPublicKey: string;
}

declare interface SetBuyDeSoFeeBasisPointsResponse {
  BuyDeSoFeeBasisPoints: number;
}

declare interface GetBuyDeSoFeeBasisPointsResponse {
  BuyDeSoFeeBasisPoints: number;
}

declare interface AdminPinPostRequest {
  PostHashHex: string;
  UnpinPost: boolean;
}

declare interface AdminUpdateGlobalFeedRequest {
  PostHashHex: string;
  RemoveFromGlobalFeed: boolean;
}

declare interface AdminRemoveNilPostsRequest {
  NumPostsToSearch: number;
}

declare interface TransactionFee {
  PublicKeyBase58Check: string;
  ProfileEntryResponse?: ProfileEntryResponse;
  AmountNanos: number;
}

declare interface AdminSetTransactionFeeForTransactionTypeRequest {
  TransactionType: lib.TxnString;
  NewTransactionFees: TransactionFee[];
}

declare interface AdminSetTransactionFeeForTransactionTypeResponse {
  TransactionFeeMap: { [key: string]: TransactionFee[] };
}

declare interface AdminSetAllTransactionFeesRequest {
  NewTransactionFees: TransactionFee[];
}

declare interface AdminSetAllTransactionFeesResponse {
  TransactionFeeMap: { [key: string]: TransactionFee[] };
}

declare interface AdminGetTransactionFeeMapResponse {
  TransactionFeeMap: { [key: string]: TransactionFee[] };
}

declare interface AdminAddExemptPublicKey {
  PublicKeyBase58Check: string;
  IsRemoval: boolean;
}

declare interface AdminGetExemptPublicKeysResponse {
  ExemptPublicKeyMap: { [key: string]: ProfileEntryResponse | undefined };
}

declare interface AdminResetJumioRequest {
  PublicKeyBase58Check: string;
  Username: string;
  JWT: string;
}

declare interface AdminUpdateJumioDeSoRequest {
  JWT: string;
  DeSoNanos: number;
}

declare interface AdminUpdateJumioDeSoResponse {
  DeSoNanos: number;
}

declare interface AdminUpdateJumioUSDCentsRequest {
  JWT: string;
  USDCents: number;
}

declare interface AdminUpdateJumioUSDCentsResponse {
  USDCents: number;
}

declare interface AdminUpdateJumioKickbackUSDCentsRequest {
  JWT: string;
  USDCents: number;
}

declare interface AdminUpdateJumioKickbackUSDCentsResponse {
  USDCents: number;
}

declare interface AdminSetJumioVerifiedRequest {
  PublicKeyBase58Check: string;
  Username: string;
}

declare interface AdminJumioCallback {
  PublicKeyBase58Check: string;
  Username: string;
  CountryAlpha3: string;
}

declare interface AdminUpdateJumioCountrySignUpBonusRequest {
  CountryCode: string;
  CountryLevelSignUpBonus: CountryLevelSignUpBonus;
}

declare interface AdminGetNFTDropRequest {
  DropNumber: number;
}

declare interface AdminGetNFTDropResponse {
  DropEntry?: NFTDropEntry;
  Posts: (PostEntryResponse | undefined)[];
}

declare interface AdminUpdateNFTDropRequest {
  DropNumber: number;
  DropTstampNanos: number;
  IsActive: boolean;
  NFTHashHexToAdd: string;
  NFTHashHexToRemove: string;
}

declare interface AdminUpdateNFTDropResponse {
  DropEntry?: NFTDropEntry;
  Posts: (PostEntryResponse | undefined)[];
}

declare interface NodeControlRequest {
  Address: string;
  MinerPublicKeys: string;
  OperationType: string;
  JWT: string;
  AdminPublicKey: string;
}

declare interface NodeStatusResponse {
  State: string;
  LatestHeaderHeight: number;
  LatestHeaderHash: string;
  LatestHeaderTstampSecs: number;
  LatestBlockHeight: number;
  LatestBlockHash: string;
  LatestBlockTstampSecs: number;
  LatestTxIndexHeight: number;
  HeadersRemaining: number;
  BlocksRemaining: number;
}

declare interface PeerResponse {
  IP: string;
  ProtocolPort: number;
  IsSyncPeer: boolean;
}

declare interface NodeControlResponse {
  DeSoStatus?: NodeStatusResponse;
  DeSoOutboundPeers: (PeerResponse | undefined)[];
  DeSoInboundPeers: (PeerResponse | undefined)[];
  DeSoUnconnectedPeers: (PeerResponse | undefined)[];
  MinerPublicKeys: string[];
}

declare interface AdminGetMempoolStatsResponse {
  TransactionSummaryStats: { [key: string]: lib.SummaryStats | undefined };
}

declare interface AdminCreateReferralHashRequest {
  UserPublicKeyBase58Check: string;
  Username: string;
  ReferrerAmountUSDCents: number;
  RefereeAmountUSDCents: number;
  MaxReferrals: number;
  RequiresJumio: boolean;
  AdminPublicKey: string;
}

declare interface AdminCreateReferralHashResponse {
  ReferralInfoResponse: ReferralInfoResponse;
}

declare interface AdminUpdateReferralHashRequest {
  ReferralHashBase58: string;
  ReferrerAmountUSDCents: number;
  RefereeAmountUSDCents: number;
  MaxReferrals: number;
  RequiresJumio: boolean;
  IsActive: boolean;
  AdminPublicKey: string;
}

declare interface AdminUpdateReferralHashResponse {
  ReferralInfoResponse: ReferralInfoResponse;
}

declare interface ReferralInfoResponse {
  IsActive: boolean;
  Info: ReferralInfo;
  ReferredUsers: ProfileEntryResponse[];
}

declare interface SimpleReferralInfoResponse {
  IsActive: boolean;
  Info: SimpleReferralInfo;
}

declare interface AdminGetAllReferralInfoForUserRequest {
  UserPublicKeyBase58Check: string;
  Username: string;
  AdminPublicKey: string;
}

declare interface AdminGetAllReferralInfoForUserResponse {
  ReferralInfoResponses: ReferralInfoResponse[];
}

declare interface AdminDownloadReferralCSVResponse {
  CSVRows: string[][];
}

declare interface AdminUploadReferralCSVRequest {
  CSVRows: string[][];
}

declare interface AdminUploadReferralCSVResponse {
  LinksCreated: number;
  LinksUpdated: number;
}

declare interface AdminDownloadRefereeCSVResponse {
  CSVRows: string[][];
}

declare interface GetGlobalParamsRequest {}

declare interface GetGlobalParamsResponse {
  USDCentsPerBitcoin: number;
  CreateProfileFeeNanos: number;
  MinimumNetworkFeeNanosPerKB: number;
  CreateNFTFeeNanos: number;
  MaxCopiesPerNFT: number;
}

declare interface UpdateGlobalParamsRequest {
  UpdaterPublicKeyBase58Check: string;
  USDCentsPerBitcoin: number;
  CreateProfileFeeNanos: number;
  CreateNFTFeeNanos: number;
  MaxCopiesPerNFT: number;
  MinimumNetworkFeeNanosPerKB: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  Password: string;
  Sign: boolean;
  Validate: boolean;
  Broadcast: boolean;
}

declare interface UpdateGlobalParamsResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface SwapIdentityRequest {
  UpdaterPublicKeyBase58Check: string;
  FromUsernameOrPublicKeyBase58Check: string;
  ToUsernameOrPublicKeyBase58Check: string;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface SwapIdentityResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface TestSignTransactionWithDerivedKeyRequest {
  TransactionHex: string;
  DerivedKeySeedHex: string;
}

declare interface TestSignTransactionWithDerivedKeyResponse {
  TransactionHex: string;
}

declare interface AdminUpdateTutorialCreatorRequest {
  PublicKeyBase58Check: string;
  IsRemoval: boolean;
  IsWellKnown: boolean;
  JWT: string;
}

declare interface AdminResetTutorialStatusRequest {
  PublicKeyBase58Check: string;
}

declare interface AdminUpdateUserGlobalMetadataRequest {
  UserPublicKeyBase58Check: string;
  Username: string;
  IsBlacklistUpdate: boolean;
  RemoveEverywhere: boolean;
  RemoveFromLeaderboard: boolean;
  IsWhitelistUpdate: boolean;
  WhitelistPosts: boolean;
  RemovePhoneNumberMetadata: boolean;
  AdminPublicKey: string;
}

declare interface AdminGetAllUserGlobalMetadataRequest {
  NumToFetch: number;
}

declare interface AdminGetAllUserGlobalMetadataResponse {
  PubKeyToUserGlobalMetadata: { [key: string]: UserMetadata | undefined };
  PubKeyToUsername: { [key: string]: string };
}

declare interface AdminGetUserGlobalMetadataRequest {
  UserPublicKeyBase58Check: string;
}

declare interface AdminGetUserGlobalMetadataResponse {
  UserMetadata: UserMetadata;
  UserProfileEntryResponse?: ProfileEntryResponse;
}

declare interface VerifiedUsernameToPKID {
  VerifiedUsernameToPKID: { [key: string]: lib.PKID | undefined };
}

declare interface VerificationUsernameAuditLog {
  TimestampNanos: number;
  VerifierUsername: string;
  VerifierPKID?: lib.PKID;
  VerifiedUsername: string;
  VerifiedPKID?: lib.PKID;
  IsRemoval: boolean;
}

declare interface FilterAuditLog {
  TimestampNanos: number;
  Filter: FilterType;
  UpdaterUsername: string;
  UpdaterPKID?: lib.PKID;
  UpdatedUsername: string;
  UpdatedPKID?: lib.PKID;
  IsRemoval: boolean;
}

declare interface AdminGrantVerificationBadgeRequest {
  UsernameToVerify: string;
  AdminPublicKey: string;
}

declare interface AdminGrantVerificationBadgeResponse {
  Message: string;
}

declare interface AdminRemoveVerificationBadgeRequest {
  UsernameForWhomToRemoveVerification: string;
  AdminPublicKey: string;
}

declare interface AdminRemoveVerificationBadgeResponse {
  Message: string;
}

declare interface AdminGetVerifiedUsersResponse {
  VerifiedUsers: string[];
}

declare interface VerificationUsernameAuditLogResponse {
  TimestampNanos: number;
  VerifierUsername: string;
  VerifierPublicKeyBase58Check: string;
  VerifiedUsername: string;
  VerifiedPublicKeyBase58Check: string;
  IsRemoval: boolean;
}

declare interface AdminGetUsernameVerificationAuditLogsRequest {
  Username: string;
}

declare interface AdminGetUsernameVerificationAuditLogsResponse {
  VerificationAuditLogs: VerificationUsernameAuditLogResponse[];
}

declare interface AdminGetUserAdminDataRequest {
  UserPublicKeyBase58Check: string;
}

declare interface AdminGetUserAdminDataResponse {
  Username: string;
  IsVerified: boolean;
  LastVerifierPublicKey: string;
  LastVerifyRemoverPublicKey: string;
  IsWhitelisted: boolean;
  LastWhitelisterPublicKey: string;
  LastWhitelistRemoverPublicKey: string;
  IsGraylisted: boolean;
  LastGraylisterPublicKey: string;
  LastGraylistRemoverPublicKey: string;
  IsBlacklisted: boolean;
  LastBlacklisterPublicKey: string;
  LastBlacklistRemoverPublicKey: string;
  PhoneNumber: string;
  Email: string;
  ReferralHashBase58Check: string;
  JumioStarterDeSoTxnHashBase58Check: string;
  ReferrerDeSoTxnHashBase58Check: string;
}

declare interface GetExchangeRateResponse {
  SatoshisPerDeSoExchangeRate: number;
  USDCentsPerBitcoinExchangeRate: number;
  NanosPerETHExchangeRate: number;
  USDCentsPerETHExchangeRate: number;
  NanosSold: number;
  USDCentsPerDeSoExchangeRate: number;
  USDCentsPerDeSoReserveExchangeRate: number;
  BuyDeSoFeeBasisPoints: number;
  SatoshisPerBitCloutExchangeRate: number;
  USDCentsPerBitCloutExchangeRate: number;
  USDCentsPerBitCloutReserveExchangeRate: number;
}

declare interface BlockchainDeSoTickerResponse {
  symbol: string;
  price_24h: number;
  volume_24h: number;
  last_trade_price: number;
}

declare interface CoinbaseDeSoTickerResponse {
  data: {
    base: string;
    currency: string;
    amount: string;
  };
}

declare interface GetAppStateRequest {
  PublicKeyBase58Check: string;
}

declare interface GetAppStateResponse {
  MinSatoshisBurnedForProfileCreation: number;
  BlockHeight: number;
  IsTestnet: boolean;
  HasStarterDeSoSeed: boolean;
  HasTwilioAPIKey: boolean;
  CreateProfileFeeNanos: number;
  CompProfileCreation: boolean;
  DiamondLevelMap: { [key: number]: number };
  HasWyreIntegration: boolean;
  HasJumioIntegration: boolean;
  BuyWithETH: boolean;
  USDCentsPerDeSoExchangeRate: number;
  JumioDeSoNanos: number;
  JumioUSDCents: number;
  JumioKickbackUSDCents: number;
  CountrySignUpBonus: CountryLevelSignUpBonus;
  DefaultFeeRateNanosPerKB: number;
  TransactionFeeMap: { [key: string]: TransactionFee[] };
  BuyETHAddress: string;
  Nodes: { [key: number]: lib.DeSoNode };
  USDCentsPerBitCloutExchangeRate: number;
  JumioBitCloutNanos: number;
}

declare interface CoinbaseResponse {
  data: {
    amount: string;
  };
}

declare interface CoingeckoResponse {
  bitcoin: {
    usd: number;
  };
}

declare interface BlockchainDotcomResponse {
  USD: {
    "15m": number;
  };
}

declare interface GeminiResponse {
  last: string;
}

declare interface KrakenResponse {
  result: {
    XXBTZUSD: {
      c: string[];
    };
  };
}

declare interface ETHTx {
  nonce: string;
  value: string;
  chainId: string;
  to: string;
  r: string;
  s: string;
}

declare interface SubmitETHTxRequest {
  PublicKeyBase58Check: string;
  Tx: ETHTx;
  TxBytes: string;
  ToSign: string[];
  SignedHashes: string[];
}

declare interface SubmitETHTxResponse {
  DESOTxHash: string;
}

declare interface ETHTxLog {
  PublicKey: string;
  DESOTxHash: string;
}

declare interface AdminProcessETHTxRequest {
  ETHTxHash: string;
}

declare interface AdminProcessETHTxResponse {
  DESOTxHash: string;
}

declare interface InfuraRequest {
  jsonrpc: string;
  method: string;
  params: any[];
  id: number;
}

declare interface InfuraResponse {
  id: number;
  jsonrpc: string;
  result: any;
  error: {
    code: number;
    message: string;
  };
}

declare interface InfuraTx {
  blockHash?: string;
  blockNumber?: string;
  from: string;
  gas: string;
  gasPrice: string;
  hash: string;
  input: string;
  nonce: string;
  to?: string;
  transactionIndex?: string;
  value: string;
  v: string;
  r: string;
  s: string;
}

declare interface QueryETHRPCRequest {
  Method: string;
  Params: any[];
}

declare interface APIBaseResponse {
  Error: string;
  Header?: HeaderResponse;
  Transactions: (TransactionResponse | undefined)[];
}

declare interface APIKeyPairRequest {
  Mnemonic: string;
  ExtraText: string;
  Index: number;
}

declare interface APIKeyPairResponse {
  Error: string;
  PublicKeyBase58Check: string;
  PublicKeyHex: string;
  PrivateKeyBase58Check: string;
  PrivateKeyHex: string;
}

declare interface APIBalanceRequest {
  PublicKeyBase58Check: string;
  Confirmations: number;
}

declare interface UTXOEntryResponse {
  TransactionIDBase58Check: string;
  Index: number;
  AmountNanos: number;
  PublicKeyBase58Check: string;
  Confirmations: number;
  UtxoType: string;
  BlockHeight: number;
}

declare interface APIBalanceResponse {
  Error: string;
  ConfirmedBalanceNanos: number;
  UnconfirmedBalanceNanos: number;
  UTXOs: (UTXOEntryResponse | undefined)[];
}

declare interface InputResponse {
  TransactionIDBase58Check: string;
  Index: number;
}

declare interface OutputResponse {
  PublicKeyBase58Check: string;
  AmountNanos: number;
}

declare interface TransactionResponse {
  TransactionIDBase58Check: string;
  RawTransactionHex?: string;
  Inputs?: (InputResponse | undefined)[];
  Outputs?: (OutputResponse | undefined)[];
  SignatureHex?: string;
  TransactionType?: string;
  BlockHashHex?: string;
  TransactionMetadata?: lib.TransactionMetadata;
  ExtraData?: { [key: string]: string };
}

declare interface TransactionInfoResponse {
  TotalInputNanos: number;
  SpendAmountNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  FeeRateNanosPerKB: number;
  SenderPublicKeyBase58Check: string;
  RecipientPublicKeyBase58Check: string;
}

declare interface APITransferDeSoRequest {
  SenderPrivateKeyBase58Check: string;
  RecipientPublicKeyBase58Check: string;
  AmountNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  DryRun: boolean;
}

declare interface APITransferDeSoResponse {
  Error: string;
  Transaction?: TransactionResponse;
  TransactionInfo?: TransactionInfoResponse;
}

declare interface APITransactionInfoRequest {
  IsMempool: boolean;
  TransactionIDBase58Check: string;
  PublicKeyBase58Check: string;
  IDsOnly: boolean;
  LastTransactionIDBase58Check: string;
  LastPublicKeyTransactionIndex: number;
  Limit: number;
}

declare interface APITransactionInfoResponse {
  Error: string;
  Transactions: (TransactionResponse | undefined)[];
  LastTransactionIDBase58Check: string;
  LastPublicKeyTransactionIndex: number;
  BalanceNanos: number;
}

declare interface APINodeInfoRequest {}

declare interface APINodeInfoResponse {
  Error: string;
}

declare interface APIBlockRequest {
  Height: number;
  HashHex: string;
  FullBlock: boolean;
}

declare interface HeaderResponse {
  BlockHashHex: string;
  Version: number;
  PrevBlockHashHex: string;
  TransactionMerkleRootHex: string;
  TstampSecs: number;
  Height: number;
  Nonce: number;
  ExtraNonce: number;
}

declare interface APIBlockResponse {
  Error: string;
  Header?: HeaderResponse;
  Transactions: (TransactionResponse | undefined)[];
}

declare interface GlobalState {
  GlobalStateRemoteNode: string;
  GlobalStateRemoteSecret: string;
  GlobalStateDB?: badger.DB;
}

declare interface HotFeedApprovedPostOp {
  IsRemoval: boolean;
  Multiplier: number;
}

declare interface HotFeedPKIDMultiplierOp {
  InteractionMultiplier: number;
  PostsMultiplier: number;
}

declare interface ReferralInfo {
  ReferralHashBase58: string;
  ReferrerPKID?: lib.PKID;
  ReferrerAmountUSDCents: number;
  RefereeAmountUSDCents: number;
  MaxReferrals: number;
  RequiresJumio: boolean;
  NumJumioAttempts: number;
  NumJumioSuccesses: number;
  TotalReferrals: number;
  TotalReferrerDeSoNanos: number;
  TotalRefereeDeSoNanos: number;
  DateCreatedTStampNanos: number;
}

declare interface SimpleReferralInfo {
  ReferralHashBase58: string;
  RefereeAmountUSDCents: number;
  MaxReferrals: number;
  TotalReferrals: number;
}

declare interface NFTDropEntry {
  IsActive: boolean;
  DropNumber: number;
  DropTstampNanos: number;
  NFTHashes: (lib.BlockHash | undefined)[];
}

declare interface UserMetadata {
  PublicKey: string;
  RemoveEverywhere: boolean;
  RemoveFromLeaderboard: boolean;
  Email: string;
  EmailVerified: boolean;
  PhoneNumber: string;
  PhoneNumberCountryCode: string;
  MessageReadStateByContact: { [key: string]: number };
  NotificationLastSeenIndex: number;
  SatoshisBurnedSoFar: number;
  HasBurnedEnoughSatoshisToCreateProfile: boolean;
  BlockedPublicKeys: { [key: string]: {} };
  WhitelistPosts: boolean;
  JumioInternalReference: string;
  JumioFinishedTime: number;
  JumioVerified: boolean;
  JumioReturned: boolean;
  JumioTransactionID: string;
  JumioDocumentKey: string;
  RedoJumio: boolean;
  JumioStarterDeSoTxnHashHex: string;
  JumioShouldCompProfileCreation: boolean;
  MustCompleteTutorial: boolean;
  IsFeaturedTutorialWellKnownCreator: boolean;
  IsFeaturedTutorialUpAndComingCreator: boolean;
  TutorialStatus: TutorialStatus;
  CreatorPurchasedInTutorialPKID?: lib.PKID;
  CreatorCoinsPurchasedInTutorial: number;
  ReferralHashBase58Check: string;
  ReferrerDeSoTxnHash: string;
  UnreadNotifications: number;
  LatestUnreadNotificationIndex: number;
}

declare interface PhoneNumberMetadata {
  PublicKey: string;
  PhoneNumber: string;
  PhoneNumberCountryCode: string;
  ShouldCompProfileCreation: boolean;
  PublicKeyDeleted: boolean;
}

declare interface WyreWalletOrderMetadata {
  LatestWyreWalletOrderWebhookPayload: WyreWalletOrderWebhookPayload;
  LatestWyreTrackWalletOrderResponse?: WyreTrackOrderResponse;
  DeSoPurchasedNanos: number;
  BasicTransferTxnBlockHash?: lib.BlockHash;
}

declare interface CountryLevelSignUpBonus {
  AllowCustomReferralAmount: boolean;
  ReferralAmountOverrideUSDCents: number;
  AllowCustomKickbackAmount: boolean;
  KickbackAmountOverrideUSDCents: number;
}

declare interface PutRemoteRequest {
  Key: string;
  Value: string;
}

declare interface PutRemoteResponse {}

declare interface GetRemoteRequest {
  Key: string;
}

declare interface GetRemoteResponse {
  Value: string;
}

declare interface BatchGetRemoteRequest {
  KeyList: string[];
}

declare interface BatchGetRemoteResponse {
  ValueList: string[];
}

declare interface DeleteRemoteRequest {
  Key: string;
}

declare interface DeleteRemoteResponse {}

declare interface SeekRemoteRequest {
  StartPrefix: string;
  ValidForPrefix: string;
  MaxKeyLen: number;
  NumToFetch: number;
  Reverse: boolean;
  FetchValues: boolean;
}

declare interface SeekRemoteResponse {
  KeysFound: string[];
  ValsFound: string[];
}

declare interface HotFeedEntry {
  PostHash?: lib.BlockHash;
  PostHashHex: string;
  HotnessScore: number;
}

declare interface HotFeedInteractionKey {
  InteractionPKID: lib.PKID;
  InteractionPostHash: lib.BlockHash;
}

declare interface HotFeedPKIDMultiplier {
  InteractionMultiplier: number;
  PostsMultiplier: number;
}

declare interface HotnessPostInfo {
  PostBlockAge: number;
  HotnessScore: number;
}

declare interface HotFeedPageRequest {
  ReaderPublicKeyBase58Check: string;
  SeenPosts: string[];
  ResponseLimit: number;
}

declare interface HotFeedPageResponse {
  HotFeedPage: PostEntryResponse[];
}

declare interface AdminUpdateHotFeedAlgorithmRequest {
  InteractionCap: number;
  TimeDecayBlocks: number;
}

declare interface AdminGetHotFeedAlgorithmResponse {
  InteractionCap: number;
  TimeDecayBlocks: number;
}

declare interface AdminUpdateHotFeedPostMultiplierRequest {
  PostHashHex: string;
  Multiplier: number;
}

declare interface AdminUpdateHotFeedUserMultiplierRequest {
  Username: string;
  InteractionMultiplier: number;
  PostsMultiplier: number;
}

declare interface AdminGetHotFeedUserMultiplierRequest {
  Username: string;
}

declare interface AdminGetHotFeedUserMultiplierResponse {
  InteractionMultiplier: number;
  PostsMultiplier: number;
}

declare interface UploadImageResponse {
  ImageURL: string;
}

declare interface GetFullTikTokURLRequest {
  TikTokShortVideoID: string;
}

declare interface GetFullTikTokURLResponse {
  FullTikTokURL: string;
}

declare interface CFVideoDetailsResponse {
  result: { [key: string]: any };
  success: boolean;
  errors: any[];
  messages: any[];
}

declare interface GetVideoStatusResponse {
  ReadyToStream: boolean;
  Duration: number;
  Dimensions: { [key: string]: any };
}

declare interface GetMessagesStatelessRequest {
  PublicKeyBase58Check: string;
  FetchAfterPublicKeyBase58Check: string;
  NumToFetch: number;
  HoldersOnly: boolean;
  HoldingsOnly: boolean;
  FollowersOnly: boolean;
  FollowingOnly: boolean;
  SortAlgorithm: string;
}

declare interface GetMessagesResponse {
  PublicKeyToProfileEntry: { [key: string]: ProfileEntryResponse | undefined };
  OrderedContactsWithMessages: (MessageContactResponse | undefined)[];
  UnreadStateByContact: { [key: string]: boolean };
  NumberOfUnreadThreads: number;
  MessagingGroups: (MessagingGroupEntryResponse | undefined)[];
}

declare interface SendMessageStatelessRequest {
  SenderPublicKeyBase58Check: string;
  RecipientPublicKeyBase58Check: string;
  MessageText: string;
  EncryptedMessageText: string;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  SenderMessagingGroupKeyName: string;
  RecipientMessagingGroupKeyName: string;
}

declare interface SendMessageStatelessResponse {
  TstampNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface MarkContactMessagesReadRequest {
  JWT: string;
  UserPublicKeyBase58Check: string;
  ContactPublicKeyBase58Check: string;
}

declare interface MarkAllMessagesReadRequest {
  JWT: string;
  UserPublicKeyBase58Check: string;
}

declare interface RegisterMessagingGroupKeyRequest {
  OwnerPublicKeyBase58Check: string;
  MessagingPublicKeyBase58Check: string;
  MessagingGroupKeyName: string;
  MessagingKeySignatureHex: string;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface RegisterMessagingGroupKeyResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface GetAllMessagingGroupKeysRequest {
  OwnerPublicKeyBase58Check: string;
}

declare interface GetAllMessagingGroupKeysResponse {
  MessagingGroupEntries: (MessagingGroupEntryResponse | undefined)[];
}

declare interface CheckPartyMessagingKeysRequest {
  SenderPublicKeyBase58Check: string;
  SenderMessagingKeyName: string;
  RecipientPublicKeyBase58Check: string;
  RecipientMessagingKeyName: string;
}

declare interface CheckPartyMessagingKeysResponse {
  SenderMessagingPublicKeyBase58Check: string;
  SenderMessagingKeyName: string;
  IsSenderMessagingKey: boolean;
  RecipientMessagingPublicKeyBase58Check: string;
  RecipientMessagingKeyName: string;
  IsRecipientMessagingKey: boolean;
}

declare interface GetBlockTemplateRequest {
  PublicKeyBase58Check: string;
  NumHeaders: number;
  HeaderVersion: number;
}

declare interface GetBlockTemplateResponse {
  Headers: string[];
  ExtraNonces: number[];
  BlockID: string;
  DifficultyTargetHex: string;
  LatestBlockTemplateStats?: lib.BlockTemplateStats;
}

declare interface SubmitBlockRequest {
  PublicKeyBase58Check: string;
  Header: string;
  ExtraNonce: number;
  BlockID: string;
}

declare interface SubmitBlockResponse {
  IsMainChain: boolean;
  IsOrphan: boolean;
}

declare interface NFTEntryResponse {
  OwnerPublicKeyBase58Check: string;
  ProfileEntryResponse?: ProfileEntryResponse;
  PostEntryResponse?: PostEntryResponse;
  SerialNumber: number;
  IsForSale: boolean;
  IsPending: boolean;
  IsBuyNow: boolean;
  BuyNowPriceNanos: number;
  MinBidAmountNanos: number;
  LastAcceptedBidAmountNanos: number;
  HighestBidAmountNanos: number;
  LowestBidAmountNanos: number;
  LastOwnerPublicKeyBase58Check?: string;
  EncryptedUnlockableText?: string;
}

declare interface NFTCollectionResponse {
  ProfileEntryResponse?: ProfileEntryResponse;
  PostEntryResponse?: PostEntryResponse;
  HighestBidAmountNanos: number;
  LowestBidAmountNanos: number;
  HighestBuyNowPriceNanos?: number;
  LowestBuyNowPriceNanos?: number;
  NumCopiesForSale: number;
  NumCopiesBuyNow: number;
  AvailableSerialNumbers: number[];
}

declare interface NFTBidEntryResponse {
  PublicKeyBase58Check: string;
  ProfileEntryResponse?: ProfileEntryResponse;
  PostHashHex?: string;
  PostEntryResponse?: PostEntryResponse;
  SerialNumber: number;
  BidAmountNanos: number;
  HighestBidAmountNanos?: number;
  LowestBidAmountNanos?: number;
  AcceptedBlockHeight?: number;
  BidderBalanceNanos: number;
}

declare interface CreateNFTRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  NumCopies: number;
  NFTRoyaltyToCreatorBasisPoints: number;
  NFTRoyaltyToCoinBasisPoints: number;
  HasUnlockable: boolean;
  IsForSale: boolean;
  MinBidAmountNanos: number;
  IsBuyNow: boolean;
  BuyNowPriceNanos: number;
  AdditionalDESORoyaltiesMap: { [key: string]: number };
  AdditionalCoinRoyaltiesMap: { [key: string]: number };
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface CreateNFTResponse {
  NFTPostHashHex: string;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface UpdateNFTRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  IsForSale: boolean;
  MinBidAmountNanos: number;
  IsBuyNow: boolean;
  BuyNowPriceNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface UpdateNFTResponse {
  NFTPostHashHex: string;
  SerialNumber: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface CreateNFTBidRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  BidAmountNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface CreateNFTBidResponse {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  BidAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface AcceptNFTBidRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  BidderPublicKeyBase58Check: string;
  BidAmountNanos: number;
  EncryptedUnlockableText: string;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface AcceptNFTBidResponse {
  BidderPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  BidAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface GetNFTShowcaseRequest {
  ReaderPublicKeyBase58Check: string;
}

declare interface GetNFTShowcaseResponse {
  NFTCollections: (NFTCollectionResponse | undefined)[];
}

declare interface GetNextNFTShowcaseResponse {
  NextNFTShowcaseTstamp: number;
}

declare interface GetNFTsForUserRequest {
  UserPublicKeyBase58Check: string;
  ReaderPublicKeyBase58Check: string;
  IsForSale?: boolean;
  IsPending?: boolean;
}

declare interface NFTEntryAndPostEntryResponse {
  PostEntryResponse?: PostEntryResponse;
  NFTEntryResponses: (NFTEntryResponse | undefined)[];
}

declare interface GetNFTsForUserResponse {
  NFTsMap: { [key: string]: NFTEntryAndPostEntryResponse | undefined };
}

declare interface GetNFTBidsForUserRequest {
  UserPublicKeyBase58Check: string;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetNFTBidsForUserResponse {
  NFTBidEntries: (NFTBidEntryResponse | undefined)[];
  PublicKeyBase58CheckToProfileEntryResponse: {
    [key: string]: ProfileEntryResponse | undefined;
  };
  PostHashHexToPostEntryResponse: {
    [key: string]: PostEntryResponse | undefined;
  };
}

declare interface GetNFTBidsForNFTPostRequest {
  ReaderPublicKeyBase58Check: string;
  PostHashHex: string;
}

declare interface GetNFTBidsForNFTPostResponse {
  PostEntryResponse?: PostEntryResponse;
  NFTEntryResponses: (NFTEntryResponse | undefined)[];
  BidEntryResponses: (NFTBidEntryResponse | undefined)[];
}

declare interface GetNFTCollectionSummaryRequest {
  PostHashHex: string;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetNFTCollectionSummaryResponse {
  NFTCollectionResponse?: NFTCollectionResponse;
  SerialNumberToNFTEntryResponse: {
    [key: number]: NFTEntryResponse | undefined;
  };
}

declare interface GetNFTEntriesForPostHashRequest {
  PostHashHex: string;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetNFTEntriesForPostHashResponse {
  NFTEntryResponses: (NFTEntryResponse | undefined)[];
}

declare interface TransferNFTRequest {
  SenderPublicKeyBase58Check: string;
  ReceiverPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  EncryptedUnlockableText: string;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface TransferNFTResponse {
  SenderPublicKeyBase58Check: string;
  ReceiverPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface AcceptNFTTransferRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface AcceptNFTTransferResponse {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface BurnNFTRequest {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface BurnNFTResponse {
  UpdaterPublicKeyBase58Check: string;
  NFTPostHashHex: string;
  SerialNumber: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface GetNFTsCreatedByPublicKeyRequest {
  PublicKeyBase58Check: string;
  Username: string;
  ReaderPublicKeyBase58Check: string;
  LastPostHashHex: string;
  NumToFetch: number;
}

declare interface NFTDetails {
  NFTEntryResponses: (NFTEntryResponse | undefined)[];
  NFTCollectionResponse?: NFTCollectionResponse;
}

declare interface GetNFTsCreatedByPublicKeyResponse {
  NFTs: NFTDetails[];
  LastPostHashHex: string;
}

declare interface GetAcceptedBidHistoryResponse {
  AcceptedBidHistoryMap: { [key: number]: (NFTBidEntryResponse | undefined)[] };
}

declare interface GetPostsStatelessRequest {
  PostHashHex: string;
  ReaderPublicKeyBase58Check: string;
  OrderBy: string;
  StartTstampSecs: number;
  PostContent: string;
  NumToFetch: number;
  FetchSubcomments: boolean;
  GetPostsForFollowFeed: boolean;
  GetPostsForGlobalWhitelist: boolean;
  GetPostsByDESO: boolean;
  GetPostsByClout: boolean;
  MediaRequired: boolean;
  PostsByDESOMinutesLookback: number;
  AddGlobalFeedBool: boolean;
}

declare interface PostEntryResponse {
  PostHashHex: string;
  PosterPublicKeyBase58Check: string;
  ParentStakeID: string;
  Body: string;
  ImageURLs: string[];
  VideoURLs: string[];
  RepostedPostEntryResponse?: PostEntryResponse;
  CreatorBasisPoints: number;
  StakeMultipleBasisPoints: number;
  TimestampNanos: number;
  IsHidden: boolean;
  ConfirmationBlockHeight: number;
  InMempool: boolean;
  ProfileEntryResponse?: ProfileEntryResponse;
  Comments: (PostEntryResponse | undefined)[];
  LikeCount: number;
  DiamondCount: number;
  PostEntryReaderState?: lib.PostEntryReaderState;
  InGlobalFeed?: boolean;
  InHotFeed?: boolean;
  IsPinned?: boolean;
  PostExtraData: { [key: string]: string };
  CommentCount: number;
  RepostCount: number;
  QuoteRepostCount: number;
  ParentPosts: (PostEntryResponse | undefined)[];
  IsNFT: boolean;
  NumNFTCopies: number;
  NumNFTCopiesForSale: number;
  NumNFTCopiesBurned: number;
  HasUnlockable: boolean;
  NFTRoyaltyToCreatorBasisPoints: number;
  NFTRoyaltyToCoinBasisPoints: number;
  AdditionalDESORoyaltiesMap: { [key: string]: number };
  AdditionalCoinRoyaltiesMap: { [key: string]: number };
  DiamondsFromSender: number;
  HotnessScore: number;
  PostMultiplier: number;
  RecloutCount: number;
  QuoteRecloutCount: number;
  RecloutedPostEntryResponse?: PostEntryResponse;
}

declare interface GetPostsStatelessResponse {
  PostsFound: (PostEntryResponse | undefined)[];
}

declare interface GetSinglePostRequest {
  PostHashHex: string;
  FetchParents: boolean;
  CommentOffset: number;
  CommentLimit: number;
  ReaderPublicKeyBase58Check: string;
  ThreadLevelLimit: number;
  ThreadLeafLimit: number;
  LoadAuthorThread: boolean;
  AddGlobalFeedBool: boolean;
}

declare interface GetSinglePostResponse {
  PostFound?: PostEntryResponse;
}

declare interface CommentsPostEntryResponse {
  PostEntryResponse?: PostEntryResponse;
  PosterPublicKeyBytes: string;
}

declare interface GetPostsForPublicKeyRequest {
  PublicKeyBase58Check: string;
  Username: string;
  ReaderPublicKeyBase58Check: string;
  LastPostHashHex: string;
  NumToFetch: number;
  MediaRequired: boolean;
}

declare interface GetPostsForPublicKeyResponse {
  Posts: (PostEntryResponse | undefined)[];
  LastPostHashHex: string;
}

declare interface GetPostsDiamondedBySenderForReceiverRequest {
  ReceiverPublicKeyBase58Check: string;
  ReceiverUsername: string;
  SenderPublicKeyBase58Check: string;
  SenderUsername: string;
  ReaderPublicKeyBase58Check: string;
  StartPostHashHex: string;
  NumToFetch: number;
}

declare interface GetPostsDiamondedBySenderForReceiverResponse {
  DiamondedPosts: (PostEntryResponse | undefined)[];
  TotalDiamondsGiven: number;
  ReceiverProfileEntryResponse?: ProfileEntryResponse;
  SenderProfileEntryResponse?: ProfileEntryResponse;
}

declare interface GetLikesForPostRequest {
  PostHashHex: string;
  Offset: number;
  Limit: number;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetLikesForPostResponse {
  Likers: (ProfileEntryResponse | undefined)[];
}

declare interface GetDiamondsForPostRequest {
  PostHashHex: string;
  Offset: number;
  Limit: number;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetDiamondsForPostResponse {
  DiamondSenders: (DiamondSenderResponse | undefined)[];
}

declare interface DiamondSenderResponse {
  DiamondSenderProfile?: ProfileEntryResponse;
  DiamondLevel: number;
}

declare interface GetRepostsForPostRequest {
  PostHashHex: string;
  Offset: number;
  Limit: number;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetRepostsForPostResponse {
  Reposters: (ProfileEntryResponse | undefined)[];
  Reclouters: (ProfileEntryResponse | undefined)[];
}

declare interface GetQuoteRepostsForPostRequest {
  PostHashHex: string;
  Offset: number;
  Limit: number;
  ReaderPublicKeyBase58Check: string;
}

declare interface GetQuoteRepostsForPostResponse {
  QuoteReposts: (PostEntryResponse | undefined)[];
  QuoteReclouts: (PostEntryResponse | undefined)[];
}

declare interface GetReferralInfoForUserRequest {
  PublicKeyBase58Check: string;
  JWT: string;
}

declare interface GetReferralInfoForUserResponse {
  ReferralInfoResponses: ReferralInfoResponse[];
}

declare interface GetReferralInfoForReferralHashRequest {
  ReferralHash: string;
}

declare interface GetReferralInfoForReferralHashResponse {
  ReferralInfoResponse?: SimpleReferralInfoResponse;
  CountrySignUpBonus: CountryLevelSignUpBonus;
}

declare interface APIServer {
  Params?: lib.DeSoParams;
  Config?: config.Config;
  MinFeeRateNanosPerKB: number;
  TXIndex?: lib.TXIndex;
  GlobalState?: GlobalState;
  Twilio?: twilio.Client;
  BlockCypherAPIKey: string;
  UsdCentsPerDeSoExchangeRate: number;
  UsdCentsPerBitCoinExchangeRate: number;
  UsdCentsPerETHExchangeRate: number;
  LastTradeDeSoPriceHistory: LastTradePriceHistoryItem[];
  LastTradePriceLookback: number;
  PublicKeyBase58Prefix: string;
  HotFeedOrderedList: (HotFeedEntry | undefined)[];
  HotFeedBlockHeight: number;
  HotFeedApprovedPostsToMultipliers: { [key: lib.BlockHash]: number };
  LastHotFeedApprovedPostOpProcessedTstampNanos: number;
  HotFeedPKIDMultipliers: {
    [key: lib.PKID]: HotFeedPKIDMultiplier | undefined;
  };
  LastHotFeedPKIDMultiplierOpProcessedTstampNanos: number;
  HotFeedInteractionCap: number;
  HotFeedTimeDecayBlocks: number;
  HotFeedPostMultiplierUpdated: boolean;
  HotFeedPKIDMultiplierUpdated: boolean;
  TransactionFeeMap: { [key: lib.TxnType]: (lib.DeSoOutput | undefined)[] };
  ExemptPublicKeyMap: { [key: string]: any };
  VerifiedUsernameToPKIDMap: { [key: string]: lib.PKID | undefined };
  BlacklistedPKIDMap: { [key: lib.PKID]: string };
  BlacklistedResponseMap: { [key: string]: string };
  GraylistedPKIDMap: { [key: lib.PKID]: string };
  GraylistedResponseMap: { [key: string]: string };
  GlobalFeedPostHashes: (lib.BlockHash | undefined)[];
  TotalSupplyNanos: number;
  TotalSupplyDESO: number;
  RichList: RichListEntryResponse[];
  CountKeysWithDESO: number;
  USDCentsToDESOReserveExchangeRate: number;
  BuyDESOFeeBasisPoints: number;
  JumioUSDCents: number;
  JumioKickbackUSDCents: number;
}

declare interface LastTradePriceHistoryItem {
  LastTradePrice: number;
  Timestamp: number;
}

declare interface Route {
  Name: string;
  Method: string[];
  Pattern: string;
  HandlerFunc: http.HandlerFunc;
}

declare interface AdminRequest {
  JWT: string;
  AdminPublicKey: string;
}

declare interface AmplitudeUploadRequestBody {
  api_key: string;
  events: AmplitudeEvent[];
}

declare interface AmplitudeEvent {
  user_id: string;
  event_type: string;
  event_properties: { [key: string]: any };
}

declare interface TransactionInfo {
  TotalInputNanos: number;
  SpendAmountNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  TransactionIDBase58Check: string;
  RecipientPublicKeys: string[];
  RecipientAmountsNanos: number[];
  TransactionHex: string;
  TimeAdded: number;
}

declare interface MessageEntryResponse {
  SenderPublicKeyBase58Check: string;
  RecipientPublicKeyBase58Check: string;
  EncryptedText: string;
  TstampNanos: number;
  IsSender: boolean;
  V2: boolean;
  Version: number;
  SenderMessagingPublicKey: string;
  SenderMessagingGroupKeyName: string;
  RecipientMessagingPublicKey: string;
  RecipientMessagingGroupKeyName: string;
}

declare interface MessageContactResponse {
  PublicKeyBase58Check: string;
  Messages: (MessageEntryResponse | undefined)[];
  ProfileEntryResponse?: ProfileEntryResponse;
  NumMessagesRead: number;
}

declare interface MessagingGroupEntryResponse {
  GroupOwnerPublicKeyBase58Check: string;
  MessagingPublicKeyBase58Check: string;
  MessagingGroupKeyName: string;
  MessagingGroupMembers: (MessagingGroupMemberResponse | undefined)[];
  EncryptedKey: string;
}

declare interface MessagingGroupMemberResponse {
  GroupMemberPublicKeyBase58Check: string;
  GroupMemberKeyName: string;
  EncryptedKey: string;
}

declare interface User {
  PublicKeyBase58Check: string;
  ProfileEntryResponse?: ProfileEntryResponse;
  Utxos: (UTXOEntryResponse | undefined)[];
  BalanceNanos: number;
  UnminedBalanceNanos: number;
  PublicKeysBase58CheckFollowedByUser: string[];
  UsersYouHODL: (BalanceEntryResponse | undefined)[];
  UsersWhoHODLYouCount: number;
  HasPhoneNumber: boolean;
  CanCreateProfile: boolean;
  BlockedPubKeys: { [key: string]: {} };
  HasEmail: boolean;
  EmailVerified: boolean;
  JumioStartTime: number;
  JumioFinishedTime: number;
  JumioVerified: boolean;
  JumioReturned: boolean;
  IsAdmin: boolean;
  IsSuperAdmin: boolean;
  IsBlacklisted: boolean;
  IsGraylisted: boolean;
  TutorialStatus: TutorialStatus;
  CreatorPurchasedInTutorialUsername?: string;
  CreatorCoinsPurchasedInTutorial: number;
  MustCompleteTutorial: boolean;
}

declare interface BalanceEntryResponse {
  HODLerPublicKeyBase58Check: string;
  CreatorPublicKeyBase58Check: string;
  HasPurchased: boolean;
  BalanceNanos: number;
  BalanceNanosUint256: uint256.Int;
  NetBalanceInMempool: number;
  ProfileEntryResponse?: ProfileEntryResponse;
}

declare interface RichListEntry {
  KeyBytes: string;
  BalanceNanos: number;
}

declare interface RichListEntryResponse {
  PublicKeyBase58Check: string;
  BalanceNanos: number;
  BalanceDESO: number;
  Percentage: number;
  Value: number;
}

declare interface GetTxnRequest {
  TxnHashHex: string;
}

declare interface GetTxnResponse {
  TxnFound: boolean;
}

declare interface SubmitTransactionRequest {
  TransactionHex: string;
}

declare interface SubmitTransactionResponse {
  Transaction?: lib.MsgDeSoTxn;
  TxnHashHex: string;
  PostEntryResponse?: PostEntryResponse;
}

declare interface UpdateProfileRequest {
  UpdaterPublicKeyBase58Check: string;
  ProfilePublicKeyBase58Check: string;
  NewUsername: string;
  NewDescription: string;
  NewProfilePic: string;
  NewCreatorBasisPoints: number;
  NewStakeMultipleBasisPoints: number;
  IsHidden: boolean;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface UpdateProfileResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
  CompProfileCreationTxnHashHex: string;
}

declare interface ExchangeBitcoinRequest {
  PublicKeyBase58Check: string;
  BurnAmountSatoshis: number;
  FeeRateSatoshisPerKB: number;
  LatestBitcionAPIResponse?: lib.BlockCypherAPIFullAddressResponse;
  BTCDepositAddress: string;
  Broadcast: boolean;
  SignedHashes: string[];
}

declare interface ExchangeBitcoinResponse {
  TotalInputSatoshis: number;
  BurnAmountSatoshis: number;
  ChangeAmountSatoshis: number;
  FeeSatoshis: number;
  BitcoinTransaction?: wire.MsgTx;
  SerializedTxnHex: string;
  TxnHashHex: string;
  DeSoTxnHashHex: string;
  UnsignedHashes: string[];
}

declare interface SendDeSoRequest {
  SenderPublicKeyBase58Check: string;
  RecipientPublicKeyOrUsername: string;
  AmountNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface SendDeSoResponse {
  TotalInputNanos: number;
  SpendAmountNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  TransactionIDBase58Check: string;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface CreateLikeStatelessRequest {
  ReaderPublicKeyBase58Check: string;
  LikedPostHashHex: string;
  IsUnlike: boolean;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface CreateLikeStatelessResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface SubmitPostRequest {
  UpdaterPublicKeyBase58Check: string;
  PostHashHexToModify: string;
  ParentStakeID: string;
  BodyObj?: lib.DeSoBodySchema;
  RepostedPostHashHex: string;
  PostExtraData: { [key: string]: string };
  IsHidden: boolean;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  InTutorial: boolean;
}

declare interface SubmitPostResponse {
  TstampNanos: number;
  PostHashHex: string;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface CreateFollowTxnStatelessRequest {
  FollowerPublicKeyBase58Check: string;
  FollowedPublicKeyBase58Check: string;
  IsUnfollow: boolean;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface CreateFollowTxnStatelessResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
}

declare interface BuyOrSellCreatorCoinRequest {
  UpdaterPublicKeyBase58Check: string;
  CreatorPublicKeyBase58Check: string;
  OperationType: string;
  DeSoToSellNanos: number;
  CreatorCoinToSellNanos: number;
  DeSoToAddNanos: number;
  MinDeSoExpectedNanos: number;
  MinCreatorCoinExpectedNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  InTutorial: boolean;
  BitCloutToSellNanos: number;
  BitCloutToAddNanos: number;
  MinBitCloutExpectedNanos: number;
}

declare interface BuyOrSellCreatorCoinResponse {
  ExpectedDeSoReturnedNanos: number;
  ExpectedCreatorCoinReturnedNanos: number;
  FounderRewardGeneratedNanos: number;
  SpendAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface TransferCreatorCoinRequest {
  SenderPublicKeyBase58Check: string;
  CreatorPublicKeyBase58Check: string;
  ReceiverUsernameOrPublicKeyBase58Check: string;
  CreatorCoinToTransferNanos: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface TransferCreatorCoinResponse {
  SpendAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface SendDiamondsRequest {
  SenderPublicKeyBase58Check: string;
  ReceiverPublicKeyBase58Check: string;
  DiamondPostHashHex: string;
  DiamondLevel: number;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
  InTutorial: boolean;
}

declare interface SendDiamondsResponse {
  SpendAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface DAOCoinRequest {
  UpdaterPublicKeyBase58Check: string;
  ProfilePublicKeyBase58CheckOrUsername: string;
  OperationType: DAOCoinOperationTypeString;
  CoinsToMintNanos: uint256.Int;
  CoinsToBurnNanos: uint256.Int;
  TransferRestrictionStatus: TransferRestrictionStatusString;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface DAOCoinResponse {
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface TransferDAOCoinRequest {
  SenderPublicKeyBase58Check: string;
  ProfilePublicKeyBase58CheckOrUsername: string;
  ReceiverPublicKeyBase58CheckOrUsername: string;
  DAOCoinToTransferNanos: uint256.Int;
  MinFeeRateNanosPerKB: number;
  TransactionFees: TransactionFee[];
}

declare interface TransferDAOCoinResponse {
  SpendAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface AuthorizeDerivedKeyRequest {
  OwnerPublicKeyBase58Check: string;
  DerivedPublicKeyBase58Check: string;
  ExpirationBlock: number;
  AccessSignature: string;
  DeleteKey: boolean;
  DerivedKeySignature: boolean;
  TransactionFees: TransactionFee[];
  MinFeeRateNanosPerKB: number;
}

declare interface AuthorizeDerivedKeyResponse {
  SpendAmountNanos: number;
  TotalInputNanos: number;
  ChangeAmountNanos: number;
  FeeNanos: number;
  Transaction?: lib.MsgDeSoTxn;
  TransactionHex: string;
  TxnHashHex: string;
}

declare interface AppendExtraDataRequest {
  TransactionHex: string;
  ExtraData: { [key: string]: string };
}

declare interface AppendExtraDataResponse {
  TransactionHex: string;
}

declare interface GetTransactionSpendingRequest {
  TransactionHex: string;
}

declare interface GetTransactionSpendingResponse {
  TotalSpendingNanos: number;
}

declare interface GetTutorialCreatorsRequest {
  ResponseLimit: number;
}

declare interface UpdateTutorialStatusRequest {
  PublicKeyBase58Check: string;
  TutorialStatus: TutorialStatus;
  CreatorPurchasedInTutorialPublicKey: string;
  ClearCreatorCoinPurchasedInTutorial: boolean;
  JWT: string;
}

declare interface GetTutorialCreatorResponse {
  UpAndComingProfileEntryResponses: ProfileEntryResponse[];
  WellKnownProfileEntryResponses: ProfileEntryResponse[];
}

declare interface StartOrSkipTutorialRequest {
  PublicKeyBase58Check: string;
  JWT: string;
  IsSkip: boolean;
}

declare interface GetUsersStatelessRequest {
  PublicKeysBase58Check: string[];
  SkipForLeaderboard: boolean;
  GetUnminedBalance: boolean;
}

declare interface GetUsersResponse {
  UserList: (User | undefined)[];
  DefaultFeeRateNanosPerKB: number;
  ParamUpdaters: { [key: string]: boolean };
}

declare interface GetUserMetadataRequest {
  PublicKeyBase58Check: string;
}

declare interface GetUserMetadataResponse {
  HasPhoneNumber: boolean;
  CanCreateProfile: boolean;
  BlockedPubKeys: { [key: string]: {} };
  HasEmail: boolean;
  EmailVerified: boolean;
  JumioFinishedTime: number;
  JumioVerified: boolean;
  JumioReturned: boolean;
}

declare interface GetProfilesRequest {
  PublicKeyBase58Check: string;
  Username: string;
  UsernamePrefix: string;
  Description: string;
  OrderBy: string;
  NumToFetch: number;
  ReaderPublicKeyBase58Check: string;
  ModerationType: string;
  FetchUsersThatHODL: boolean;
  AddGlobalFeedBool: boolean;
}

declare interface GetProfilesResponse {
  ProfilesFound: (ProfileEntryResponse | undefined)[];
  NextPublicKey?: string;
}

declare interface ProfileEntryResponse {
  PublicKeyBase58Check: string;
  Username: string;
  Description: string;
  IsHidden: boolean;
  IsReserved: boolean;
  IsVerified: boolean;
  Comments: (PostEntryResponse | undefined)[];
  Posts: (PostEntryResponse | undefined)[];
  CoinEntry?: CoinEntryResponse;
  DAOCoinEntry?: DAOCoinEntryResponse;
  CoinPriceDeSoNanos: number;
  CoinPriceBitCloutNanos: number;
  UsersThatHODL: (BalanceEntryResponse | undefined)[];
  IsFeaturedTutorialWellKnownCreator: boolean;
  IsFeaturedTutorialUpAndComingCreator: boolean;
}

declare interface CoinEntryResponse {
  CreatorBasisPoints: number;
  DeSoLockedNanos: number;
  NumberOfHolders: number;
  CoinsInCirculationNanos: number;
  CoinWatermarkNanos: number;
  BitCloutLockedNanos: number;
}

declare interface DAOCoinEntryResponse {
  NumberOfHolders: number;
  CoinsInCirculationNanos: uint256.Int;
  MintingDisabled: boolean;
  TransferRestrictionStatus: TransferRestrictionStatusString;
}

declare interface GetSingleProfileRequest {
  PublicKeyBase58Check: string;
  Username: string;
  NoErrorOnMissing: boolean;
}

declare interface GetSingleProfileResponse {
  Profile?: ProfileEntryResponse;
  IsBlacklisted: boolean;
  IsGraylisted: boolean;
}

declare interface GetHodlersForPublicKeyRequest {
  PublicKeyBase58Check: string;
  Username: string;
  LastPublicKeyBase58Check: string;
  NumToFetch: number;
  IsDAOCoin: boolean;
  FetchHodlings: boolean;
  FetchAll: boolean;
}

declare interface GetHodlersForPublicKeyResponse {
  Hodlers: (BalanceEntryResponse | undefined)[];
  LastPublicKeyBase58Check: string;
}

declare interface DiamondSenderSummaryResponse {
  SenderPublicKeyBase58Check: string;
  ReceiverPublicKeyBase58Check: string;
  TotalDiamonds: number;
  HighestDiamondLevel: number;
  DiamondLevelMap: { [key: number]: number };
  ProfileEntryResponse?: ProfileEntryResponse;
}

declare interface GetDiamondsForPublicKeyRequest {
  PublicKeyBase58Check: string;
  FetchYouDiamonded: boolean;
}

declare interface GetDiamondsForPublicKeyResponse {
  DiamondSenderSummaryResponses: (DiamondSenderSummaryResponse | undefined)[];
  TotalDiamonds: number;
}

declare interface GetFollowsStatelessRequest {
  PublicKeyBase58Check: string;
  Username: string;
  GetEntriesFollowingUsername: boolean;
  LastPublicKeyBase58Check: string;
  NumToFetch: number;
}

declare interface GetFollowsResponse {
  PublicKeyToProfileEntry: { [key: string]: ProfileEntryResponse | undefined };
  NumFollowers: number;
}

declare interface GetUserGlobalMetadataRequest {
  UserPublicKeyBase58Check: string;
  JWT: string;
}

declare interface GetUserGlobalMetadataResponse {
  Email: string;
  PhoneNumber: string;
}

declare interface UpdateUserGlobalMetadataRequest {
  UserPublicKeyBase58Check: string;
  JWT: string;
  Email: string;
  MessageReadStateUpdatesByContact: { [key: string]: number };
}

declare interface GetNotificationsCountRequest {
  PublicKeyBase58Check: string;
}

declare interface GetNotificationsCountResponse {
  NotificationsCount: number;
  LastUnreadNotificationIndex: number;
  UpdateMetadata: boolean;
}

declare interface GetNotificationsRequest {
  PublicKeyBase58Check: string;
  FetchStartIndex: number;
  NumToFetch: number;
  FilteredOutNotificationCategories: { [key: string]: boolean };
}

declare interface GetNotificationsResponse {
  Notifications: (TransactionMetadataResponse | undefined)[];
  ProfilesByPublicKey: { [key: string]: ProfileEntryResponse | undefined };
  PostsByHash: { [key: string]: PostEntryResponse | undefined };
  LastSeenIndex: number;
}

declare interface SetNotificationMetadataRequest {
  PublicKeyBase58Check: string;
  LastSeenIndex: number;
  LastUnreadNotificationIndex: number;
  UnreadNotifications: number;
  JWT: string;
}

declare interface TransactionMetadataResponse {
  Metadata?: lib.TransactionMetadata;
  TxnOutputResponses: (OutputResponse | undefined)[];
  Txn?: TransactionResponse;
  Index: number;
}

declare interface BlockPublicKeyRequest {
  PublicKeyBase58Check: string;
  BlockPublicKeyBase58Check: string;
  Unblock: boolean;
  JWT: string;
}

declare interface BlockPublicKeyResponse {
  BlockedPublicKeys: { [key: string]: {} };
}

declare interface IsFollowingPublicKeyRequest {
  PublicKeyBase58Check: string;
  IsFollowingPublicKeyBase58Check: string;
}

declare interface IsFolllowingPublicKeyResponse {
  IsFollowing: boolean;
}

declare interface IsHodlingPublicKeyRequest {
  PublicKeyBase58Check: string;
  IsHodlingPublicKeyBase58Check: string;
  IsDAOCoin: boolean;
}

declare interface IsHodlingPublicKeyResponse {
  IsHodling: boolean;
  BalanceEntry?: BalanceEntryResponse;
}

declare interface GetUserDerivedKeysRequest {
  PublicKeyBase58Check: string;
}

declare interface UserDerivedKey {
  OwnerPublicKeyBase58Check: string;
  DerivedPublicKeyBase58Check: string;
  ExpirationBlock: number;
  IsValid: boolean;
}

declare interface GetUserDerivedKeysResponse {
  DerivedKeys: { [key: string]: UserDerivedKey | undefined };
}

declare interface DeletePIIRequest {
  PublicKeyBase58Check: string;
  JWT: string;
}

declare interface SendPhoneNumberVerificationTextRequest {
  PublicKeyBase58Check: string;
  PhoneNumber: string;
  JWT: string;
}

declare interface SendPhoneNumberVerificationTextResponse {}

declare interface SubmitPhoneNumberVerificationCodeRequest {
  JWT: string;
  PublicKeyBase58Check: string;
  PhoneNumber: string;
  VerificationCode: string;
}

declare interface SubmitPhoneNumberVerificationCodeResponse {
  TxnHashHex: string;
}

declare interface ResendVerifyEmailRequest {
  PublicKey: string;
  JWT: string;
}

declare interface VerifyEmailRequest {
  PublicKey: string;
  EmailHash: string;
}

declare interface JumioInitRequest {
  customerInternalReference: string;
  userReference: string;
  successUrl: string;
  errorUrl: string;
}

declare interface JumioInitResponse {
  redirectUrl: string;
  transactionReference: string;
}

declare interface JumioBeginRequest {
  PublicKey: string;
  ReferralHashBase58: string;
  SuccessURL: string;
  ErrorURL: string;
  JWT: string;
}

declare interface JumioBeginResponse {
  URL: string;
}

declare interface JumioFlowFinishedRequest {
  PublicKey: string;
  JumioInternalReference: string;
  JWT: string;
}

declare interface JumioIdentityVerification {
  similarity: string;
  validity: boolean;
  reason: string;
}

declare interface JumioRejectReason {
  rejectReasonCode: string;
  rejectReasonDescription: string;
  rejectReasonDetails: any;
}

declare interface GetJumioStatusForPublicKeyRequest {
  JWT: string;
  PublicKeyBase58Check: string;
}

declare interface GetJumioStatusForPublicKeyResponse {
  JumioFinishedTime: number;
  JumioReturned: boolean;
  JumioVerified: boolean;
  BalanceNanos?: number;
}

declare interface WyreWalletOrderWebhookPayload {
  referenceId: string;
  accountId: string;
  orderId: string;
  orderStatus: string;
  transferId: string;
  failedReason: string;
}

declare interface WyreWalletOrderFullDetails {
  id: string;
  createdAt: number;
  owner: string;
  status: string;
  orderType: string;
  sourceAmount: number;
  purchaseAmount: number;
  sourceCurrency: string;
  destCurrency: string;
  transferId: string;
  dest: string;
  authCodesRequested: boolean;
  errorCategory: string;
  errorCode: string;
  errorMessage: string;
  failureReason: string;
  accountId: string;
  paymentNetworkErrorCode: string;
  internalErrorCode: string;
}

declare interface WyreTransferDetails {
  owner: string;
  reversingSubStatus: any;
  source: string;
  pendingSubStatus: any;
  status: string;
  reversalReason: any;
  createdAt: number;
  sourceAmount: number;
  destCurrency: string;
  sourceCurrency: string;
  statusHistories: {
    id: string;
    transferId: string;
    createdAt: number;
    type: string;
    statusOrder: number;
    statusDetail: string;
    state: string;
    failedState: any;
  }[];
  blockchainTx: {
    id: string;
    networkTxId: string;
    createdAt: number;
    confirmations: number;
    timeObserved: number;
    blockTime: number;
    blockhash: string;
    amount: number;
    direction: string;
    networkFee: number;
    address: string;
    sourceAddress: any;
    currency: string;
    twinTxId: any;
  };
  expiresAt: number;
  completedAt: number;
  cancelledAt: any;
  failureReason: any;
  updatedAt: number;
  exchangeRate: number;
  destAmount: number;
  fees: {
    BTC: number;
    USD: number;
  };
  totalFees: number;
  customId: string;
  dest: string;
  message: any;
  id: string;
}

declare interface WyreTrackOrderResponse {
  transferId: string;
  feeCurrency: string;
  fee: number;
  fees: {
    BTC: number;
    USD: number;
  };
  sourceCurrency: string;
  destCurrency: string;
  sourceAmount: number;
  destAmount: number;
  destSrn: string;
  from: string;
  to: any;
  rate: number;
  customId: any;
  status: any;
  blockchainNetworkTx: any;
  message: any;
  transferHistoryEntryType: string;
  successTimeline: {
    statusDetails: string;
    state: string;
    createdAt: number;
  }[];
  failedTimeline: any[];
  failureReason: any;
  reversalReason: any;
}

declare interface WalletOrderQuotationRequest {
  SourceAmount: number;
  Country: string;
  SourceCurrency: string;
}

declare interface WyreWalletOrderQuotationPayload {
  sourceCurrency: string;
  dest: string;
  destCurrency: string;
  amountIncludeFees: boolean;
  country: string;
  sourceAmount: string;
  walletType: string;
  accountId: string;
}

declare interface WalletOrderReservationRequest {
  SourceAmount: number;
  ReferenceId: string;
  Country: string;
  SourceCurrency: string;
}

declare interface WyreWalletOrderReservationPayload {
  sourceCurrency: string;
  dest: string;
  destCurrency: string;
  country: string;
  amount: string;
  referrerAccountId: string;
  lockFields: string[];
  redirectUrl: string;
  referenceId: string;
}

declare interface GetWyreWalletOrderForPublicKeyRequest {
  PublicKeyBase58Check: string;
  Username: string;
  AdminPublicKey: string;
}

declare interface GetWyreWalletOrderForPublicKeyResponse {
  WyreWalletOrderMetadataResponses: (
    | WyreWalletOrderMetadataResponse
    | undefined
  )[];
}

declare interface WyreWalletOrderMetadataResponse {
  LatestWyreWalletOrderWebhookPayload: WyreWalletOrderWebhookPayload;
  LatestWyreTrackWalletOrderResponse?: WyreTrackOrderResponse;
  DeSoPurchasedNanos: number;
  BitCloutPurchasedNanos: number;
  BasicTransferTxnHash: string;
  Timestamp?: string;
}
