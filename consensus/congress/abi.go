package congress

import "github.com/ethereum/go-ethereum/common"

var (
	SystemEventObserverContractAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ConsensusContractAddr = common.HexToAddress("0x000000000000000000000000000000000000000A")
	GenesisMintContractAddr = common.HexToAddress("0x000000000000000000000000000000000000000B")
	CNSContractAddr = common.HexToAddress("0x000000000000000000000000000000000000000C")
	RelationsContractAddr = common.HexToAddress("0x000000000000000000000000000000000000000D")
	BridgeContractAddress = common.HexToAddress("0x000000000000000000000000000000000000000E")
	BridgeAnchorContractAddress = common.HexToAddress("0x0000000000000000000000000000000000000010")

)

const SystemEventObserverName = "SystemEventObserver"
const SystemEventObserverABI = `
[
    {
      "inputs": [],
      "name": "_systemEventNewBlockHandle",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address[]",
          "name": "vals",
          "type": "address[]"
        }
      ],
      "name": "_systemEventNewEpochHandle",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
]
`

const InitializeName = "Initialize"
const InitializeABI = `
[
	{
		"constant": false,
		"inputs": [{
			"internalType": "address[]",
			"name": "vals",
			"type": "address[]"
		}],
		"name": "initialize",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

const ConsensusName = "Consensus"
const ConsensusABI = `
[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "time",
          "type": "uint256"
        }
      ],
      "name": "OnPledge",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "time",
          "type": "uint256"
        }
      ],
      "name": "OnRedeem",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "Paused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "Unpaused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "vtype",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "quantity",
          "type": "uint256"
        }
      ],
      "name": "VotedEvent",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "uint8",
          "name": "ptype",
          "type": "uint8"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "reward",
          "type": "uint256"
        }
      ],
      "name": "WithdrawReward",
      "type": "event"
    },
    {
      "stateMutability": "payable",
      "type": "fallback"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "auther",
          "type": "address"
        }
      ],
      "name": "KAddAuthAddress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "KAuthAddresses",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "KContractOwners",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "auther",
          "type": "address"
        }
      ],
      "name": "KDelAuthAddress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "KImplementAddress",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "Pause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "Unpause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "node",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "vote",
          "type": "uint256"
        }
      ],
      "name": "approveVoteTo",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "blockReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "candidates",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "candidatesMaxLimit",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "doPledge",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "doRedeem",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "epochValidatorsNodeOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "pledge",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "extension",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "entrusted",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "inMining",
          "type": "bool"
        },
        {
          "internalType": "uint256",
          "name": "pledgeRewardLatestTime",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "everBlockHandle",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getTopValidators",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address[]",
          "name": "vals",
          "type": "address[]"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "minedBlockTotalOfDay",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "minerRewardsOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "nodeInfomationOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "pledge",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "extension",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "entrusted",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "inMining",
          "type": "bool"
        },
        {
          "internalType": "uint256",
          "name": "pledgeRewardLatestTime",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "paused",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "pledgeVoteOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "pledgeVoteTotal",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "nodeAddress",
          "type": "address"
        }
      ],
      "name": "powerOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "powerOfDays",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "pledge",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "extension",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "pledgeRewardTotal",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "receiptVoteMaxLimit",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "index",
          "type": "uint256"
        }
      ],
      "name": "rejectVote",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "updateActiveValidatorSet",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "validators",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "validatorsMaxLimit",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "withdrawNoderReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "withdrawPledgeReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "pledgeReward",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "stateMutability": "payable",
      "type": "receive"
    },
    {
      "inputs": [],
      "name": "enablePledge",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "enableCongress",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]
`
