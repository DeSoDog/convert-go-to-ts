package types


type MiningSupplyIntervalStart struct {
	StartBlockHeight uint32
	BlockRewardNanos uint64
} 
type PurchaseSupplyIntervalStart struct {
	// How much each unit costs to purchase in Satoshis.
	SatoshisPerUnit uint64
	// The total supply cutoff at which this price applies.
	SupplyStartNanos uint64
}