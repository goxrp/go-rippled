package gorippled

import (
	"github.com/grokify/mogo/database"
)

func KeyGenerationMethods() []Method {
	return []Method{
		{
			Name:            MethodValidationCreate,
			Category:        Category{Name: CategoryKeyGeneration},
			Summary:         "Generate keys for a new rippled validator.",
			Description:     "Use the `validation_create` command to generate [cryptographic keys a rippled server can use to identify itself to the network](https://xrpl.org/peer-protocol.html#node-key-pair). Similar to the [`wallet_propose` method](https://xrpl.org/wallet_propose.html), this method only generates a set of keys in the proper format. It does not any makes changes to the XRP Ledger data or server configuration.\n\n*The `validation_create` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*\n\nYou can configure your server to use the generated key pair to sign validations (validation key pair) or regular peer-to-peer communications ([node key pair](https://xrpl.org/peer-protocol.html#node-key-pair)).\n\nTip: For configuring a robust validator, you should use the `validator-keys` tool (included in the `rippled` RPM) to generate validator tokens (which can be rotated) with an offline master key. For more information, see [Validator Setup](https://xrpl.org/run-rippled-as-a-validator.html#3-enable-validation-on-your-rippled-server).",
			FunctionType:    database.FunctionOther,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodWalletPropose,
			ExampleName:     "with key type",
			ExampleSlug:     "key_type",
			Category:        Category{Name: CategoryKeyGeneration},
			Summary:         "Generate keys for a new account.",
			Description:     "Use the `wallet_propose` method to generate a key pair and XRP Ledger address. This command only generates key and address values, and does not affect the XRP Ledger itself in any way. To become a funded address stored in the ledger, the address must [receive a Payment transaction](https://xrpl.org/accounts.html#creating-accounts) that provides enough XRP to meet the [reserve requirement](https://xrpl.org/reserves.html).\n\n*The `wallet_propose` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!* (This command is restricted to protect against people sniffing network traffic for account secrets, since admin commands are not usually transmitted over the outside network.).",
			FunctionType:    database.FunctionOther,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodWalletPropose,
			ExampleName:     "with passphrase",
			ExampleSlug:     "passphrase",
			Category:        Category{Name: CategoryKeyGeneration},
			Summary:         "Generate keys for a new account.",
			Description:     "Use the `wallet_propose` method to generate a key pair and XRP Ledger address. This command only generates key and address values, and does not affect the XRP Ledger itself in any way. To become a funded address stored in the ledger, the address must [receive a Payment transaction](https://xrpl.org/accounts.html#creating-accounts) that provides enough XRP to meet the [reserve requirement](https://xrpl.org/reserves.html).\n\n*The `wallet_propose` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!* (This command is restricted to protect against people sniffing network traffic for account secrets, since admin commands are not usually transmitted over the outside network.).",
			FunctionType:    database.FunctionOther,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
	}
}

func StatusAndDebuggingMethods() []Method {
	return []Method{
		{
			Name:            MethodConsensusInfo,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about the state of consensus as it happens.",
			Description:     "The `consensus_info` command provides information about the [consensus process](https://xrpl.org/consensus.html) for debugging purposes.\n\n*The `consensus_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodFeature,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about protocol amendments.",
			Description:     "The `feature` command returns information about [amendments](https://xrpl.org/amendments.html) this server knows about, including whether they are enabled and whether the server is voting in favor of those amendments in the [amendment process](https://xrpl.org/amendments.html#amendment-process).\n\nYou can use the `feature` command to configure the server to vote against or in favor of an amendment. This change persists even if you restart the server.\n\n*The `feature` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodFetchInfo,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about the server's sync with the network.",
			Description:     "The `fetch_info` command returns information about objects that this server is currently fetching from the network, and how many peers have that information. It can also be used to reset current fetches.\n\n*The `fetch_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodGetCounts,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get statistics about the server's internals and memory usage.",
			Description:     "The `get_counts` command provides various stats about the health of the server, mostly the number of objects of different types that it currently holds in memory.\n\n*The `get_counts` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodPrint,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about internal subsystems.",
			Description:     "The `print` command returns the current status of various internal subsystems, including peers, the ledger cleaner, and the resource manager.\n\n*The `print` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			IsOneOf:         true,
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   false,
			HasApiCli:       true,
		},
		{
			Name:            MethodValidatorInfo,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get the server's validation settings, if configured as a validator.",
			Description:     "The `validator_info` method returns the current validator settings of the server, if it is configured as a validator.\n\n*The `validator_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodValidatorListSites,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about sites that publish validator lists.",
			Description:     "The `validator_list_sites` command returns status information of sites serving validator lists.\n\n*The `validator_list_sites` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
		{
			Name:            MethodValidators,
			Category:        Category{Name: CategoryStatusAndDebugging},
			Summary:         "Get information about the current validators.",
			Description:     "The `validators` command returns human readable information about the current list of published and trusted validators used by the server.\n\n*The `validators` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			FunctionType:    database.FunctionRead,
			HasApiWebsocket: true,
			HasApiJsonRpc:   true,
			HasApiCli:       true,
		},
	}
}
