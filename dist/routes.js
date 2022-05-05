"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getStructFiles = exports.getLib = exports.getCountryPackage = exports.getRoutes = void 0;
const getRoutes = () => {
    let routes = `admin_buy_deso.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago
dao_coin_exchange.go
admin_feed.go
Seperate out global state into it's own struct in order for it to be … (
3 months ago
admin_fees.go
run go fmt on backend project (#267)
2 months ago
admin_jumio.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago
admin_nft.go
Seperate out global state into it's own struct in order for it to be … (
3 months ago
admin_node.go
run go fmt on backend project (#267)
2 months ago
admin_referrals.go
update AdminUploadReferralCSV to take in a multipart form payload wit…
2 months ago
admin_transaction.go
run go fmt on backend project (#267)
2 months ago
admin_tutorial.go
run go fmt on backend project (#267)
2 months ago
admin_user.go
Move some types from core (#244)
2 months ago
base.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago
bitcoin_price.go
Buy clout with eth (#142)
5 months ago
eth.go
remove JWT check on query ETH rpc endpoint (#284)
24 days ago
exchange.go
Apply MediaRequired param for GetPostEntriesByDESOAfterTimePaginated …
20 hours ago
exchange_test.go
Ln/node monetization (#160)
5 months ago
expose_global_state.go
Move some types from core (#244)
2 months ago
extra_data_utils.go
Enable proper decoding for extra data values (#306)
22 hours ago
global_state.go
add endpoint to manage default kickback amount (#256)
2 months ago
global_state_test.go
Give different sign-up bonus amounts based on country of ID (#242)
2 months ago
hot_feed.go
Make backend work for core DAO changes (#269)
2 months ago
media.go
Add video dimensions to video status route (#304)
5 days ago
message.go
DeSo V3 Messages (#253)
last month
miner.go
Move vendored libraries to separate repos (#228)
3 months ago
nft.go
Include the AcceptedBlockHeight in the BidEntryResponse if it exists (#…
yesterday
post.go
Apply MediaRequired param for GetPostEntriesByDESOAfterTimePaginated …
20 hours ago
referrals.go
Use IP address to infer country sign up bonus config (#251)
2 months ago
server.go
add endpoint to fetch accepted bid history for a given NFT post hash …
2 days ago
shared.go
DeSo V3 Messages (#253)
last month
supply.go
update supply stats once at start up (#259)
2 months ago
transaction.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago
tutorial.go
Make backend work for core DAO changes (#269)
2 months ago
user.go
Add expiration check for derived keys in the backend API (#296)
10 days ago
verify.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago
wyre.go
Cache reserve exchange rate, buy deso fee basis points, and jumio def…
29 days ago`.matchAll(/[a-z0-9_]+?\.go/gs);
    routes = [...routes];
    return routes;
};
exports.getRoutes = getRoutes;
const getCountryPackage = () => {
    return [
        "https://raw.githubusercontent.com/deso-protocol/backend/main/countries/alpha-2-to-alpha-3-lookup.go",
        "https://raw.githubusercontent.com/deso-protocol/backend/main/countries/iso-3166-1-alpha-3-codes.go",
    ];
};
exports.getCountryPackage = getCountryPackage;
const getLib = () => {
    const lib = `
  base58.go
  snapshot.go
Ln/deso (#118)
5 months ago
bitcoin_burner.go
Factor the double-spend check into its own function (#165)
2 months ago
bitcoin_burner_test.go
Move vendored libraries to separate repos (#160)
3 months ago
block_producer.go
Move vendored libraries to separate repos (#160)
3 months ago
block_view.go
Fix stubborn NFT splits bug due to reordering of royalties in map (#188)
last month
block_view_balance_entry.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_bitcoin.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_bitcoin_test.go
Split block view and tests (#164)
2 months ago
block_view_creator_coin.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_creator_coin_test.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_dao_coin.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_dao_coin_test.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_flush.go
Fix instances where a pointer is taken of a map key iterator (#183)
last month
block_view_follow.go
Split block view and tests (#164)
2 months ago
block_view_follow_test.go
Split block view and tests (#164)
2 months ago
block_view_like.go
Split block view and tests (#164)
2 months ago
block_view_like_test.go
Split block view and tests (#164)
2 months ago
block_view_message.go
Fix get-all-messaging-keys & add base point test (#217)
2 days ago
block_view_message_test.go
Fix get-all-messaging-keys & add base point test (#217)
2 days ago
block_view_nft.go
add AcceptedBlockHeight to NFTBidEntry (REQUIRES RESYNC) (#224)
2 days ago
block_view_nft_test.go
Fix stubborn NFT splits bug due to reordering of royalties in map (#188)
last month
block_view_post.go
Split block view and tests (#164)
2 months ago
block_view_post_test.go
Move blockheights into DeSoParams and set values accordingly (#171)
last month
block_view_profile.go
Update Postgres handling of DAO Coins and NFTs (#194)
last month
block_view_profile_test.go
Messages V3 (#182)
last month
block_view_test.go
Buy Now NFTs with Buy Now Price (FORKING CHANGE) (#166)
2 months ago
block_view_types.go
add remaining UtxoTypes to UtxoType.String() (#229)
2 days ago
blockchain.go
Messages V3 (#182)
last month
blockchain_test.go
DAO coins (#174)
2 months ago
connection_manager.go
Move vendored libraries to separate repos (#160)
3 months ago
constants.go
Add new constant to constants file to represent the NodeSource map ke…
last month
db_utils.go
add txindex metadata for Burns and Accept NFT Transfers (#228)
2 days ago
db_utils_test.go
Messages V3 (#182)
last month
deso_math.go
Messages V3 (#182)
last month
diff.go
In retrospect, it was inevitable
9 months ago
errors.go
nft and dao coin notifications after fork (#185)
last month
event_manager.go
Historical balance lookup support (#141)
4 months ago
load_test.go
DAO coins (#174)
2 months ago
mempool.go
add txindex metadata for Burns and Accept NFT Transfers (#228)
2 days ago
mempool_test.go
Ln/deso (#118)
5 months ago
miner.go
Move vendored libraries to separate repos (#160)
3 months ago
network.go
nft and dao coin notifications after fork (#185)
last month
network_test.go
Messages V3 (#182)
last month
nodes.go
Add upcoming overclout node (#219)
2 days ago
notifier.go
Move vendored libraries to separate repos (#160)
3 months ago
peer.go
Move vendored libraries to separate repos (#160)
3 months ago
peer_test.go
In retrospect, it was inevitable
9 months ago
postgres.go
add AcceptedBlockHeight to NFTBidEntry (REQUIRES RESYNC) (#224)
2 days ago
reserved_usernames.go
DAO coins (#174)
2 months ago
seed_txns.go
Ln/deso (#118)
5 months ago
server.go
Fix instances where a pointer is taken of a map key iterator (#183)
last month
supply.go
DAO coins (#174)
2 months ago
supply_test.go
DAO coins (#174)
2 months ago
supplydata_test.go
In retrospect, it was inevitable
9 months ago
txindex.go
Return error if error during txindex db update (#220)
5 days ago
types.go
create maps for additional royalties when converting to/from PGPostEn…
19 days ago
utils.go
Ln/deso (#118)
5 months ago
varint.go
In retrospect, it was inevitable
  `.matchAll(/[a-z0-9_]+?\.go/gs);
    return [...lib];
};
exports.getLib = getLib;
const getStructFiles = () => {
    const structs = `admin_buy_deso.go
added go types and indvidual files
20 hours ago
admin_feed.go
added go types and indvidual files
20 hours ago
admin_fees.go
added go types and indvidual files
20 hours ago
admin_jumio.go
added go types and indvidual files
20 hours ago
admin_nft.go
added go types and indvidual files
20 hours ago
admin_node.go
added go types and indvidual files
20 hours ago
admin_referrals.go
added go types and indvidual files
20 hours ago
admin_transaction.go
added go types and indvidual files
20 hours ago
admin_tutorial.go
added go types and indvidual files
20 hours ago
admin_user.go
added go types and indvidual files
20 hours ago
base.go
added go types and indvidual files
20 hours ago
bitcoin_burner.go
added go types and indvidual files
20 hours ago
bitcoin_price.go
added go types and indvidual files
20 hours ago
block_producer.go
added go types and indvidual files
20 hours ago
block_view.go
added go types and indvidual files
20 hours ago
block_view_nft.go
added go types and indvidual files
20 hours ago
block_view_test.go
added go types and indvidual files
20 hours ago
block_view_types.go
added go types and indvidual files
20 hours ago
blockchain.go
added go types and indvidual files
20 hours ago
connection_manager.go
added go types and indvidual files
20 hours ago
constants.go
added go types and indvidual files
20 hours ago
db_utils.go
added go types and indvidual files
20 hours ago
errors.go
added go types and indvidual files
20 hours ago
eth.go
added go types and indvidual files
20 hours ago
event_manager.go
added go types and indvidual files
20 hours ago
exchange.go
added go types and indvidual files
20 hours ago
extra_data_utils.go
added go types and indvidual files
20 hours ago
global_state.go
added go types and indvidual files
20 hours ago
hot_feed.go
added go types and indvidual files
20 hours ago
iso-3166-1-alpha-3-codes.go
added go types and indvidual files
20 hours ago
media.go
added go types and indvidual files
20 hours ago
mempool.go
added go types and indvidual files
20 hours ago
message.go
added go types and indvidual files
20 hours ago
miner.go
added go types and indvidual files
20 hours ago
network.go
clean up
20 hours ago
nft.go
added go types and indvidual files
20 hours ago
nodes.go
added go types and indvidual files
20 hours ago
notifier.go
added go types and indvidual files
20 hours ago
peer.go
added go types and indvidual files
20 hours ago
post.go
added go types and indvidual files
20 hours ago
postgres.go
added go types and indvidual files
20 hours ago
referrals.go
added go types and indvidual files
20 hours ago
server.go
added go types and indvidual files
20 hours ago
shared.go
added go types and indvidual files
20 hours ago
supply.go
added go types and indvidual files
20 hours ago
transaction.go
added go types and indvidual files
20 hours ago
tutorial.go
added go types and indvidual files
20 hours ago
txindex.go
added go types and indvidual files
20 hours ago
types.go
added go types and indvidual files
20 hours ago
user.go
added go types and indvidual files
20 hours ago
verify.go
added go types and indvidual files
20 hours ago
wyre.go
added go types and indvidual files
20 hours ago`.matchAll(/[a-z0-9_]+?\.go/gs);
    return [...structs][0];
};
exports.getStructFiles = getStructFiles;
//# sourceMappingURL=routes.js.map