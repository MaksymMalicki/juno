package feedergatewaysync

// GasPrice represents gas price details in different units.
type GasPrice struct {
	PriceInWei string `json:"price_in_wei"`
	PriceInFri string `json:"price_in_fri"`
}

// Transaction represents a single transaction.
type Transaction struct {
	TransactionHash string   `json:"transaction_hash"`
	Version         string   `json:"version"`
	MaxFee          string   `json:"max_fee"`
	Signature       []string `json:"signature"`
	Nonce           string   `json:"nonce"`
	SenderAddress   string   `json:"sender_address"`
	Calldata        []string `json:"calldata"`
	Type            string   `json:"type"`
}

// BuiltinInstanceCounter represents the number of times each builtin was used.
type BuiltinInstanceCounter struct {
	RangeCheckBuiltin int `json:"range_check_builtin"`
	EcdsaBuiltin      int `json:"ecdsa_builtin"`
	PedersenBuiltin   int `json:"pedersen_builtin"`
}

// ExecutionResources represents the computational resources used.
type ExecutionResources struct {
	Steps                  int                    `json:"n_steps"`
	BuiltinInstanceCounter BuiltinInstanceCounter `json:"builtin_instance_counter"`
	MemoryHoles            int                    `json:"n_memory_holes"`
}

// TransactionReceipt represents the receipt for a processed transaction.
type TransactionReceipt struct {
	ExecutionStatus    string             `json:"execution_status"`
	TransactionIndex   int                `json:"transaction_index"`
	TransactionHash    string             `json:"transaction_hash"`
	L2ToL1Messages     []interface{}      `json:"l2_to_l1_messages"` // Assuming empty array
	Events             []interface{}      `json:"events"`            // Assuming empty array
	ExecutionResources ExecutionResources `json:"execution_resources"`
	ActualFee          string             `json:"actual_fee"`
}

// Block represents the complete block structure.
type Block struct {
	BlockHash             string               `json:"block_hash"`
	ParentBlockHash       string               `json:"parent_block_hash"`
	BlockNumber           int                  `json:"block_number"`
	StateRoot             string               `json:"state_root"`
	TransactionCommitment string               `json:"transaction_commitment"`
	EventCommitment       string               `json:"event_commitment"`
	Status                string               `json:"status"`
	L1DaMode              string               `json:"l1_da_mode"`
	L1GasPrice            GasPrice             `json:"l1_gas_price"`
	L1DataGasPrice        GasPrice             `json:"l1_data_gas_price"`
	L2GasPrice            GasPrice             `json:"l2_gas_price"`
	Transactions          []Transaction        `json:"transactions"`
	Timestamp             int64                `json:"timestamp"`
	SequencerAddress      string               `json:"sequencer_address"`
	TransactionReceipts   []TransactionReceipt `json:"transaction_receipts"`
	StarknetVersion       string               `json:"starknet_version"`
}
